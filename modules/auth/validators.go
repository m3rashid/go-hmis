package auth

type Login struct {
	Email    string `bson:"email" json:"email" validate:"required,min=2,max=100"`
	Password string `bson:"password" json:"password" validate:"required,min=6"`
}
