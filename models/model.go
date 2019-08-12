package models

//Category model
type Category struct {
	ID   string `form:"id" json:"id"`
	Name string `form:"Name" json:"Name"`
}

//List Model
type List struct {
	ID           string `form:"id" json:"id"`
	Name         string `form:"name" json:"name"`
	CategoryName string `form:"CategoryID" json:"CategoryID"`
}

//ResponseCategory for response
type ResponseCategory struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Category
}

//ResponseList from Response
type ResponseList struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []List
}
