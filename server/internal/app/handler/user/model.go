package user

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission int    `json:"permission"`
}

type JWTToken struct {
	Token string `json:"token"`
}

//var jwtKey = []byte("my_secret_key")
