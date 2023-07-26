package auth

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	EmailVerified bool     `json:"emailVerified"`
	Origin        string   `json:"origin"`
	Roles         []string `json:"roles,omitempty"`
	Profile       string   `json:"profile,omitempty"`
}

type ProfileModel struct {
	gorm.Model
	Age            int    `json:"age"`
	Sex            string `json:"sex"`
	BloodGroup     string `json:"bloodGroup"`
	ProfilePicture string `json:"profilePicture"`
	Designation    string `json:"designation"`
	Department     string `json:"department"`
}

type RoleModel struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Workspace   string `json:"workspace"`
}

type WorkspaceModel struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ColorCode   string `json:"colorCode"`
}

type PermissionModel struct {
	gorm.Model
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}
