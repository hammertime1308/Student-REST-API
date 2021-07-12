package model

type Student struct {
	Uid     uint64 `json:"uid"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Class   int    `json:"class"`
	Subject string `json:"subject"`
}
