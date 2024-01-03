package staff

type CreateDTO struct {
	Chapa        string `json:"chapa" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	FunctionName string `json:"function_name" binding:"required"`
	Whatsapp     string `json:"whatsapp" binding:"required"`
}

type SignInDTO struct {
	Chapa    string `json:"chapa" binding:"required"`
	Password string `json:"password" binding:"required"`
}
