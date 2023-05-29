package entity

type ResponseModel struct {
	Successful bool        `json:"successful"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta"`
}
