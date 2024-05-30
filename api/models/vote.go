package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Vote struct {
	Id       int     `json:"id"`
	UserId   *int    `json:"user_id"`
	OptionId int     `json:"option_id"`
	Guest    *string `json:"guest"`
	PollId   int     `json:"poll_id"`

	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
}

func (v *Vote) calculateHash() string {
	record := fmt.Sprint(v.OptionId) + fmt.Sprint(v.PollId) + fmt.Sprint(v.Timestamp)
	if v.UserId != nil {
		record += fmt.Sprint(*v.UserId)
	}
	if v.Guest != nil {
		record += *v.Guest
	}
	record += v.PrevHash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func (v *Vote) IsVoteValid(previous *Vote) bool {
	if previous.Hash != v.PrevHash {
		return false
	}
	
	if v.calculateHash() != v.Hash {
		return false
	}
	
	return true
}

func (v *Vote) CompleteVote(previous *Vote) {
	v.Timestamp = time.Now().Unix()
	v.PrevHash = previous.Hash
	
	v.Hash = v.calculateHash()
}

type VotePoll struct {
	Id        int    `json:"id"`
	IsGuest   bool   `json:"is_guest"`
	User      string `json:"user"`
	OptionId  int    `json:"option_id"`
	Timestamp int64  `json:"timestamp"`
}
