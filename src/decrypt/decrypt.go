package decrypt

import (
	"context"
	"encoding/base64"
	"example/hello/src/data_info"
	"example/hello/src/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/gorilla/mux"
)

type KMSDecryptAPI interface {
	Decrypt(ctx context.Context,
		params *kms.DecryptInput,
		optFns ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

func DecodeData(c context.Context, api KMSDecryptAPI, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	return api.Decrypt(c, input)
}

func DecryptKey(w http.ResponseWriter, r *http.Request) {
	// awsEndpoint := os.Getenv("AWS_ENDPOINT")
	// awsRegion := os.Getenv("AWS_REGION")
	vars := mux.Vars(r)
	keyID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		fmt.Println("error", err)
	}

	db, err := database.Connection()
	if err != nil {
		fmt.Println("error", err)
	}

	dataInfo := data_info.DataInfo{}

	requestD := db.Table("data_info").Where("id = ?", keyID).Take(&dataInfo)

	fmt.Println(requestD)
	ctx := context.TODO()

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://localhost:4566",
			SigningRegion: "eu-west-1",
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		fmt.Println(err)
	}
	client := kms.NewFromConfig(cfg)

	text := dataInfo.ApiKey

	blob, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic("error converting string to blob, " + err.Error())
	}

	input2 := kms.DecryptInput{
		CiphertextBlob: blob,
	}

	result2, err := DecodeData(context.TODO(), client, &input2)

	if err != nil {
		fmt.Println("Got error decrypting data: ", err)
	}

	fmt.Println(string(result2.Plaintext))
}
