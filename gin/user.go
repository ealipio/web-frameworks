package main

type userType struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

var users = []userType{
	{Id: "1", Name: "Jon Connor", IsActive: true},
	{Id: "2", Name: "Sara Connor", IsActive: true},
	{Id: "3", Name: "Genesis", IsActive: false},
	{Id: "4", Name: "Smith Agent", IsActive: false},
}
