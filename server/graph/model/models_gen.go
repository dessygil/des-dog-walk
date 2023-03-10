// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Pet struct {
	ID      string `json:"id"`
	PetName string `json:"petName"`
}

type User struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Pets           []*Pet `json:"Pets"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	SocialLogin    bool   `json:"socialLogin"`
	SocialProvider string `json:"socialProvider"`
}
