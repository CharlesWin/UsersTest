package parser

import (
	. "UsersTest/internal/config"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

type SpecialDate struct {
	time.Time
}

type User struct {
	Id              int         `json:"Id"`
	UserName        string      `json:"UserName"`
	FullName        string      `json:"FullName"`
	City            string      `json:"City"`
	BirthDate       SpecialDate `json:"BirthDate"`
	Department      string      `json:"Department"`
	Gender          string      `json:"Gender"`
	ExperienceYears int         `json:"ExperienceYears"`
}

type ShortInfo struct {
	Id       int    `json:"Id"`
	UserName string `json:"UserName"`
}

var users map[int]*User

func GetUsers() map[int]*User {
	if users == nil {
		readUsers()
	}
	return users
}

func readUsers() {
	users = make(map[int]*User)
	var tmp []*User
	data, err := ioutil.ReadFile(GetInstance().InMemory.Path)
	if err != nil {
		Log.Fatal(err)
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		Log.Fatal(err)
	}
	for _, user := range tmp {
		GetUsers()[user.Id] = user
	}
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strTime := string(input)
	strTime = strings.Trim(strTime, `"`)
	newTime, err := time.Parse("Monday 2 Jan 2006", strTime)
	if err != nil {
		return err
	}
	sd.Time = newTime
	return err
}
