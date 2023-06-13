package domain

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username string `json:"username" validate:"required,min=8,max=24"`
	Password string `json:"password" validate:"required,min=8,max=16"`
	Email    string `json:"email" validate:"email,required"`
	Otp      string `json:"otp"`
}
