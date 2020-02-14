package data

type User struct {
	Name        string `json:"username" form:"username"    binding:"required"`
	Password    string `json:"password" form:"password"    binding:"required"`
	Phonenumber string `json:"phonenumber"  form:"phonenumber"`
}
