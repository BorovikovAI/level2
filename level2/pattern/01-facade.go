package main

import "fmt"

// Scenario: The login registration function is integrated together.

// General login registration interface
type IUser interface {
	Login(phone, code int) (*User, error)
	Register(phone, code int) error
}

// Facial mode interface
type IUserFacade interface {
	LoginOrRegister(phone int, code int) (*User, error)
}

type User struct {
	name  string
	phone int
}

type UserService struct {
}

func (u *UserService) Login(phone, code int) (*User, error) {
	// A series of operations
	fmt.Println("Login-ok")
	return &User{name: "Alex", phone: phone}, nil
}

func (u *UserService) Register(phone, code int) error {
	// A series of operations
	fmt.Println("Register-ok")
	return nil
}

func (u *UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user == nil {
		err := u.Register(phone, code)
		return nil, err
	}

	return user, nil
}

func main() {
	var User IUserFacade = &UserService{}
	fmt.Println(User.LoginOrRegister(8916, 111))
}
