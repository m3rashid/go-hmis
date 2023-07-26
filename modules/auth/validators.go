package auth

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type LoginBody struct {
	Email    string `bson:"email" json:"email" validate:"required,min=2,max=100"`
	Password string `bson:"password" json:"password" validate:"required,min=6"`
}

type CreateAccountBody struct {
	Name          string `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Email         string `bson:"email" json:"email" validate:"required,min=2,max=100"`
	Password      string `bson:"password" json:"password" validate:"required,min=6"`
	EmailVerified bool   `bson:"emailVerified" json:"emailVerified"`
	Origin        string `bson:"origin" json:"origin"`
	Role          string `bson:"role" json:"role,omitempty" validate:"mongodb"`
	Profile       string `bson:"profile" json:"profile,omitempty" validate:"mongodb"`
}

func (c CreateAccountBody) Validate() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

type ProfileBody struct {
	Age            int    `bson:"age" json:"age"`
	Sex            string `bson:"sex" json:"sex"`
	BloodGroup     string `bson:"bloodGroup" json:"bloodGroup"`
	ProfilePicture string `bson:"profilePicture" json:"profilePicture"`
	Designation    string `bson:"designation" json:"designation"`
	Department     string `bson:"department" json:"department"`
}

type RoleBody struct {
	Name        string `bson:"name" json:"name" validate:"required"`
	Description string `bson:"description" json:"description"`
	Workspace   string `bson:"workspace" json:"workspace" validate:"required"`
}

type WorkspaceBody struct {
	Name        string `bson:"name" json:"name" validate:"required"`
	Description string `bson:"description" json:"description" validate:"required"`
	ColorCode   string `bson:"colorCode" json:"colorCode" validate:"required"`
}

type PermissionBody struct {
	Name       string `bson:"name" json:"name" validate:"required"`
	Identifier string `bson:"identifier" json:"identifier" validate:"required"`
}
