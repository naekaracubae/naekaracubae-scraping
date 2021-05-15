package _struct

type SecretManager struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

type Subscriber struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
