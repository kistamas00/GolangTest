package test

import (
	"testing"
	"net/http/httptest"
	"controller/engine"
	"model"
	"net/http"
	"encoding/json"
	"log"
)

type mockedDB struct {
}

func (db *mockedDB) Init() {}
func (db *mockedDB) GetUsers() map[string]string { return map[string]string{ "testuser" : "testpass" } }
func (db *mockedDB) GetVotes(filterIds []string) []model.Vote {
	return []model.Vote{
		{
			Id: "testid",
			Question: "testquestion",
			Options: []string{
				"testoption1",
				"testoption2",
			},
			Votes: []int64{
				1,
				2,
			},
		},
	}
}
func (db *mockedDB) AddVote(vote model.Vote) {}
func (db *mockedDB) EditVote(vote model.Vote) bool { return false }
func (db *mockedDB) RemoveVote(id string) bool { return false }
func (db *mockedDB) IncreaseVoteCount(id string, optionIndex int) bool { return false }

var db = mockedDB{}

func TestPublicGetVotes(t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/votes/testid")

	if err != nil {
		log.Println("HTTP get error!")
		t.Fail()
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	var votes []model.Vote
	json.NewDecoder(resp.Body).Decode(&votes)

	if len(votes) != 1 {
		log.Println("Wrong vote length!")
		t.Fail()
	}

	vote := votes[0]

	if vote.Id != "testid" || vote.Question != "testquestion" || len(vote.Options) != 2 || len(vote.Votes) != 2 {
		log.Println("Wrong vote field value(s)!")
		t.Fail()
	}
}
