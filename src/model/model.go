package model

type Vote struct {
	Id string
	Question string
	Options []string
	Votes []int64
}