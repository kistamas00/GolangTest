package db

import "model"

type DataBase interface {
	Init()
	GetUsers() map[string]string
	GetVotes(filterIds []string) []model.Vote
	AddVote(vote model.Vote)
	EditVote(vote model.Vote) bool
	RemoveVote(id string) bool
	IncreaseVoteCount(id string, optionIndex int) bool
}
