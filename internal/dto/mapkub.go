package dto

type MapKub struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Cert      string `json:"cert"`
	Key       string `json:"key"`
}
