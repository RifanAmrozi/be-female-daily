package model

type QRRequest struct {
	Data string `json:"data" binding:"required"`
}

type QRResponse struct {
	ImageBase64 string `json:"image_base64"`
}
