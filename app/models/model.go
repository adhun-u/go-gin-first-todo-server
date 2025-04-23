package models

type LoginInfo struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
}

type Idonly struct {
	Id int `json:"id"`
}
type UpdateTodo struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
}
