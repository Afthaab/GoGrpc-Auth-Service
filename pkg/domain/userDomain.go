package domain

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username string `json:"username" gorm:"unique" validate:"required,min=2,max=50"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"notnull;unique" validate:"email,required"`
}
