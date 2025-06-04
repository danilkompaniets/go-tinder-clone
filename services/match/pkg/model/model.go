package model

type Match struct {
	FromId       string `json:"from_id"`
	ToId         string `json:"to_id"`
	FromDecision *bool  `json:"from_decision"`
	ToDecision   *bool  `json:"to_decision"`
}
