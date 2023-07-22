package main

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func handleLoginRequest(loginURL, password string) (string, error ) {

	
	return "", nil
}
