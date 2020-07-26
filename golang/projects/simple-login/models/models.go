package models

// User ...
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

// ResponseResult ...
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
