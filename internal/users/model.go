package users

type Users struct {
	Id           string `json:"id" bson:"_id, omitempty" ` // DTO
	Email        string `json:"email" bson:"email"`
	Username     string ` json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUsers struct { // <-  DTO
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
