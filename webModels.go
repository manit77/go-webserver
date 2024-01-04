package main

type loginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResult struct {
	AuthToken string `json:"authtoken"`
}
