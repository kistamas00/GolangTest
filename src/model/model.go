package model

type Vote struct {
	Id string
	Question string
	Options []string
	Votes []int64
}

func (vote *Vote) IdIsInArray(ids []string) bool {

	for i := 0; i < len(ids); i++ {
		if ids[i] == vote.Id {
			return true
		}
	}
	return false
}