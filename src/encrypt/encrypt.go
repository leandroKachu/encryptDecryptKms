package encrypt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"example/hello/src/data_info"
	"example/hello/src/database"
	"flag"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSEncryptAPI interface {
	Encrypt(ctx context.Context,
		params *kms.EncryptInput,
		optFns ...func(*kms.Options)) (*kms.EncryptOutput, error)
}

func EncryptText(c context.Context, api KMSEncryptAPI, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
	return api.Encrypt(c, input)
}

var (
// awsEndpoint = os.Getenv("AWS_ENDPOINT")
// awsRegion = os.Getenv("AWS_REGION")
// KeyIdenv  = os.Getenv("KeyId")
)

func EncryptKey(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	var dataInfo data_info.DataInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dataInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	text := dataInfo.ApiKey
	fmt.Println(text)
	flag.Parse()
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
	input := &kms.EncryptInput{
		KeyId:     aws.String("ca6e1fc0-e30f-4295-a182-4004859ad3fd"),
		Plaintext: []byte(text),
	}

	result, err := client.Encrypt(context.TODO(), input)

	if err != nil {
		fmt.Println("Ocorreu um erro ao criptografar:", err)
	}

	encryptedValue := base64.StdEncoding.EncodeToString(result.CiphertextBlob)
	fmt.Println("Texto criptografado em base64:", encryptedValue)
	dataInfo.ApiKey = encryptedValue
	db, err := database.Connection()
	if err != nil {
		fmt.Println("error")
	}

	requestD := db.Table("data_info").Create(&dataInfo)
	fmt.Println("requestDB", requestD)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}

// CREATE TABLE data_info (
//     ID INT AUTO_INCREMENT PRIMARY KEY,
//     api_key VARCHAR(255) NOT NULL,
//     secret_key VARCHAR(255) NOT NULL
// );
