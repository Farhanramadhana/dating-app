package dto

type Signup struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
