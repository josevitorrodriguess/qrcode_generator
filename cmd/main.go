package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/qrcode_generator/internal/api"
	"github.com/josevitorrodriguess/qrcode_generator/internal/storage/awsS3"
)

func main() {
	godotenv.Load()

	s3 := awsS3.S3Storage{}

	s3.InitSession()
	api := api.Api{
		S3: s3,
	}

	http.HandleFunc("/qrcode", api.GenerateQrCodeHandler)

	http.ListenAndServe(":3050", nil)
}
