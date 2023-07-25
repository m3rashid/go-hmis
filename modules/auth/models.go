package auth

type User struct {
	Name          string  `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Email         string  `bson:"email" json:"email" validate:"required,min=2,max=100"`
	Password      string  `bson:"password" json:"password" validate:"required,min=6"`
	EmailVerified bool    `bson:"email_verified" json:"email_verified"`
	Origin        string  `bson:"origin" json:"origin"`
	Role          Role    `bson:"role" json:"role"`
	Profile       Profile `bson:"profile" json:"profile"`
}

type Profile struct {
	Age            int    `bson:"age" json:"age"`
	Sex            string `bson:"sex" json:"sex"`
	BloodGroup     string `bson:"bloodGroup" json:"bloodGroup"`
	ProfilePicture string `bson:"profilePicture" json:"profilePicture"`
	Designation    string `bson:"designation" json:"designation"`
	Department     string `bson:"department" json:"department"`
}

type Role struct {
	Name        string    `bson:"name" json:"name" validate:"required"`
	Description string    `bson:"description" json:"description"`
	Workspace   Workspace `bson:"workspace" json:"workspace" validate:"required"`
}

type Workspace struct {
	Name        string `bson:"name" json:"name" validate:"required"`
	Description string `bson:"description" json:"description" validate:"required"`
	ColorCode   string `bson:"color_code" json:"color_code" validate:"required"`
}

type Permission struct {
	Name       string `bson:"name" json:"name" validate:"required"`
	Identifier string `bson:"identifier" json:"identifier" validate:"required"`
}
