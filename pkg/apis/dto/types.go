package dto

type AvengerDto struct {
	Id                 string `json:"id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	ImageName          string `json:"imageName"`
	Image []byte `json:"image"`
}