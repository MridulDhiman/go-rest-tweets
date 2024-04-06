package main


type Tweet struct {
	Id string `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

type UpdateRequestBody struct {
	Content string `json:"content"`
}