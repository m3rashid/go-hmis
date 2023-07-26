package auth

type UserModel struct {
	Name          string       `bson:"name" json:"name"`
	Email         string       `bson:"email" json:"email"`
	Password      string       `bson:"password" json:"password"`
	EmailVerified bool         `bson:"emailVerified" json:"emailVerified"`
	Origin        string       `bson:"origin" json:"origin"`
	Role          RoleModel    `bson:"role" json:"role,omitempty"`
	Profile       ProfileModel `bson:"profile" json:"profile,omitempty"`
}

type ProfileModel struct {
	Age            int    `bson:"age" json:"age"`
	Sex            string `bson:"sex" json:"sex"`
	BloodGroup     string `bson:"bloodGroup" json:"bloodGroup"`
	ProfilePicture string `bson:"profilePicture" json:"profilePicture"`
	Designation    string `bson:"designation" json:"designation"`
	Department     string `bson:"department" json:"department"`
}

type RoleModel struct {
	Name        string         `bson:"name" json:"name"`
	Description string         `bson:"description" json:"description"`
	Workspace   WorkspaceModel `bson:"workspace" json:"workspace"`
}

type WorkspaceModel struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	ColorCode   string `bson:"colorCode" json:"colorCode"`
}

type PermissionModel struct {
	Name       string `bson:"name" json:"name"`
	Identifier string `bson:"identifier" json:"identifier"`
}
