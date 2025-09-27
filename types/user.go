package types

type User struct{
	StudentId string `json:"-"`
	Name string `json:"name"`
	Grade string `json:"-"`
}
