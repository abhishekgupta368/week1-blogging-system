package main

type BlogPayload struct {
	RequestType string      `json:"type"`
	Body        interface{} `json:"body"`
	CallBackUrl string      `json:"callback"`
}

type Blog struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type DeleteBody struct {
	Name string `json:"name"`
}
