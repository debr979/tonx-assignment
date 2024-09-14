package models

type MgrLoginRequest struct {
	ManagerName string `json:"manager_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type GetBodyRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
