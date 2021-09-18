package model

type AddPermision struct {
	Permissiontype string `json:"p_type,omitempty"`
	Subject    string `json:"subject,omitempty" binding:"required"`
	Object    string `json:"object,omitempty" binding:"required"`
}