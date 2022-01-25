package models

type Message struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
	User User   `json:"user"`
	Date Time   `json:"date"`
}

type Status struct {
	Online bool `json:"online"`
	User   User `json:"user"`
}
