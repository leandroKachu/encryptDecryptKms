FROM debian:latest
RUN apt-get update
RUN apt-get install -y locales locales-all tzdata ca-certificates
ENV LC_ALL pt_BR.UTF-8
ENV LANG pt_BR.UTF-8
ENV LANGUAGE pt_BR.UTF-8
ENV TZ="America/Sao_Paulo"

RUN  update-ca-certificates


ADD /build /app
CMD ["/app/main"]