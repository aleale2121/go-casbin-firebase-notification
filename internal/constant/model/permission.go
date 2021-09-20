package model

type Permision struct {
	Subject    string `json:"subject" validate:"required"`
	Object    string `json:"object" validate:"required"`
	Action   string `json:"action" validate:"required"`
}