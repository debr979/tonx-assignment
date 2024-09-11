package models

type MgrLoginRequest struct {
	ManagerName string `json:"manager_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
