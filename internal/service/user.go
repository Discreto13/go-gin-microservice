package user

type User struct {
	Id       uint   `json: "-"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Birthday string `json: "birthday"`
}
