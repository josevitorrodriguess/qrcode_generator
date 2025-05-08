package main

import (
	"github.com/joho/godotenv"
	"github.com/josevitorrodriguess/qrcode_generator/internal/storage/awsS3"
)

func main() {
	godotenv.Load()

	s3 := awsS3.S3Storage{}

	s3.InitSession()
	
}
