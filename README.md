# Qr Code Generator
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Amazon S3](https://img.shields.io/badge/Amazon%20S3-FF9900?style=for-the-badge&logo=amazons3&logoColor=white)

A golang application that generates Qr Codes and stores them in AWS S3. This project  demonstrate the integration of skip2 library for QR code generation and AWS S3 for storage.

## Table of Contents

- [How to Use](#how-to-use)
    - [Prerequisites](#prerequisites)
    - [Running the application](#running-the-application)
    - [AWS S3 Configuration](#aws-s3-configuration)

- [API Endpoints](#api-endpoints)
= [License](#license)

## How to Use

This section provides comprehensive instructions for setting up and running the QR Code Generator application.

### Prerequisites

- Golang 1.24.3
<!-- - Docker -->
- AWS Account with S3 access
- AWS CLI configured with appropriate credentials

### Environment Variables

Create a `.env` file in the project root with the following variables:

```env
AWS_SECRET_KEY=your_secret_key
AWS_ACCESS_KEY=your_acess_ley
AWS_S3_BUCKET=your_bucket_name
AWS_REGION=your_region
```

### Running the Application
1. Create the `.env` file as described above
2. Donwload dependencies with:   
`go mod tidy`
3. Run the application:  
 `go run cmd/main.go`
    #### AWS S3 Configuration
    1. Create a S3 bucket in your AWS account
    2. Update the `AWS_S3_BUCKET` in your `.env`
    3. Ensure your AWS credentials have appropriate permissions to acess the S3 bucket

## API Endpoints
POST /qrcode  
Generate a QR code from the provided text and store it in AWS S3. The QR code will be generated as a PNG image.

**Parameters**

| Name | Required | Type | Description |
|------|----------|------|-------------|
| `text` | required | string | The text content to be encoded in the QR code. This can be any string value that you want to convert into a QR code. |

**Response**

```json
{
    "url": "https://your-bucket.s3.your-region.amazonaws.com/random-uuid"
}
```

**Error Response**

If an error occurs during QR code generation or S3 upload, the API will return a 500 Internal Server Error.

**Example Usage**

```bash
curl -X POST http://localhost:3050/qrcode \
     -H "Content-Type: application/json" \
     -d '{"text": "https://example.com"}'
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 