package link


type CreateLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UploadLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}