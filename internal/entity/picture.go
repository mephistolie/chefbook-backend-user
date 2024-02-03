package entity

type PictureUpload struct {
	PictureLink string
	UploadUrl   string
	FormData    map[string]string
	MaxSize     int64
}
