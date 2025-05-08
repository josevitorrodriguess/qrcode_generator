package storage

type Storage interface {
	UploadFile(fileData []byte, fileName, contentType string)
}
