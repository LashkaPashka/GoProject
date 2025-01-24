package link


type CreateLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UploadLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksResponse struct {
	Links []Link `json:"links"`
	Count int64 `json:"count"`
}