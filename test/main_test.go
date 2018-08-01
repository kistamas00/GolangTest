package test

import (
	"testing"
	"net/http/httptest"
	"controller/engine"
	"model"
	"net/http"
	"encoding/json"
	"log"
	"strings"
)

var voteCount int
var addedVote model.Vote
var editedVote model.Vote
var voteDeleted bool

type mockedDB struct {
}

func (db *mockedDB) Init() {}
func (db *mockedDB) ClearDbAndInsertSamples() {}
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
func (db *mockedDB) AddVote(vote model.Vote) { addedVote = vote }
func (db *mockedDB) EditVote(vote model.Vote) bool { editedVote = vote; return true }
func (db *mockedDB) RemoveVote(id string) bool { voteDeleted = true; return true }
func (db *mockedDB) IncreaseVoteCount(id string, optionIndex int) bool { voteCount++; return true}

var db = mockedDB{}

func TestPublicGetVotes(t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/votes/testid")

	if err != nil {
		log.Println("HTTP get error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	var votes []model.Vote
	json.NewDecoder(res.Body).Decode(&votes)

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
func TestPublicPutVotes(t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL + "/votes/testid/inc/0", nil)

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP put error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	if voteCount != 1 {
		log.Println("Counter not increased code!")
		t.Fail()
	}
}

func TestAdminGetVotes (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/admin/votes")

	if err != nil {
		log.Println("HTTP get error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusUnauthorized {
		log.Println("Wrong status code!")
		t.Fail()
	}

	req, err := http.NewRequest("GET", ts.URL + "/admin/votes", nil)

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	req.SetBasicAuth("testuser", "testpass")
	client := &http.Client{}
	res, err = client.Do(req)

	if err != nil {
		log.Println("HTTP get error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	var votes []model.Vote
	json.NewDecoder(res.Body).Decode(&votes)

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
func TestAdminPostVotes (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL + "/admin/votes",
		strings.NewReader("question=testquestion&options=testoption1&options=testoption2"))

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP post error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusUnauthorized {
		log.Println("Wrong status code!")
		t.Fail()
	}

	req.SetBasicAuth("testuser", "testpass")
	res, err = client.Do(req)

	if err != nil {
		log.Println("HTTP post error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	if addedVote.Question != "testquestion" || addedVote.Options[0] != "testoption1" ||
		addedVote.Options[1] != "testoption2" || len(addedVote.Options) != 2 {
		log.Println("Wrong vote field value(s)!")
		t.Fail()
	}
}
func TestAdminPutVotes (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL + "/admin/votes/testid",
		strings.NewReader("question=testquestion&options=testoption1&options=testoption2"))

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP put error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusUnauthorized {
		log.Println("Wrong status code!")
		t.Fail()
	}

	req.SetBasicAuth("testuser", "testpass")
	res, err = client.Do(req)

	if err != nil {
		log.Println("HTTP put error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	if editedVote.Question != "testquestion" || editedVote.Options[0] != "testoption1" ||
		editedVote.Options[1] != "testoption2" || len(editedVote.Options) != 2 {
		log.Println("Wrong vote field value(s)!")
		t.Fail()
	}
}
func TestAdminDeleteVotes (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("DELETE", ts.URL + "/admin/votes/testid", nil)

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP delete error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusUnauthorized {
		log.Println("Wrong status code!")
		t.Fail()
	}

	req.SetBasicAuth("testuser", "testpass")
	res, err = client.Do(req)

	if err != nil {
		log.Println("HTTP delete error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Wrong status code!")
		t.Fail()
	}

	if voteDeleted != true {
		log.Println("Delete failed!")
		t.Fail()
	}
}
func TestAdminPostVotesFormValidator (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL + "/admin/votes",
		strings.NewReader("question=testquestion&options=testoption"))

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("testuser", "testpass")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP post error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusBadRequest {
		log.Println("Wrong status code!")
		t.Fail()
	}
}
func TestAdminPutVotesFormValidator (t *testing.T) {

	ts := httptest.NewServer(engine.NewEngine(&db))
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL + "/admin/votes/testid",
		strings.NewReader("options=testoption1&options=testoption2"))

	if err != nil {
		log.Println("HTTP request creation error!")
		t.Fail()
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("testuser", "testpass")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("HTTP put error!")
		t.Fail()
	}

	if res.StatusCode != http.StatusBadRequest {
		log.Println("Wrong status code!")
		t.Fail()
	}
}