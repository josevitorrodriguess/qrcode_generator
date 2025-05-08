package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josevitorrodriguess/qrcode_generator/internal/qrcode"
	"github.com/josevitorrodriguess/qrcode_generator/internal/storage/awsS3"
)

type Api struct {
	S3 awsS3.S3Storage
}
type Request struct {
	Text string `json:"text"`
}

func (api *Api) GenerateQrCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var text Request
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		http.Error(w, "error to decode json", http.StatusBadRequest)
		return
	}

	qrcode, err := qrcode.GenerateQrCode(text.Text)
	if err != nil {
		http.Error(w, "failed to generate qrcode", http.StatusInternalServerError)
		return
	}
	
	url, err := api.S3.UploadFile(qrcode, "image", ".png")
	if err != nil {
		http.Error(w, "failed to save qrcode on bucket", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	response := map[string]string{"qrcode": url}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
