package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID    int
	first string
}

type MockDataStore struct {
	user map[int]User
}

func (m MockDataStore) getUser(id int) (User, error) {
	u, ok := m.user[id]
	if !ok {
		return User{}, errors.New("User not found")
	}
	return u, nil
}

func (m MockDataStore) saveUser(u User) error {
	_, ok := m.user[u.ID]
	if ok {
		return fmt.Errorf("User already exists")

	}
	m.user[u.ID] = u
	return nil
}

type DataStore interface {
	getUser(id int) (User, error)
	saveUser(u User) error
}

type service struct {
	ds DataStore
}

func (s service) getUser(id int) (User, error) {
	return s.ds.getUser(id)
}

func (s service) saveUser(u User) error {
	return s.ds.saveUser(u)
}

func main() {
	db := MockDataStore{
		user: make(map[int]User),
	}
	servc := service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		first: "Bhumit",
	}
	err := servc.saveUser(u1)
	if err != nil {
		fmt.Println(err)
	}
	userRequest1, er := servc.getUser(1)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("The user Name is ", userRequest1.first)

	}
	userRequest2, er := servc.getUser(2)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("The user Name is ", userRequest2.first)

	}
}
