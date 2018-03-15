package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Session struct {
	SessionID int
	Imsi      string `json:"IMSI"`

UpfTeid   int
NrTeid    int
Qfis      []Qfi
}

type Qfi struct {
	QfiID    int
	Priority int `json:"PRI"`
	ARP      int
	FilterID int
}
//var SessionS map[string]Session

func main() {
	Sessions := make(map[string]Session)
	Sessions["4600020"] = Session{1, "4600020", 0x3456, 0x4325, []Qfi{{1, 2, 3, 5}, {2, 3, 6, 7}}}
	Sessions["4600010"] = Session{2, "4600010", 0x4563, 0x8970, []Qfi{{4, 5, 6, 7}, {1, 5, 4, 3}, {7, 5, 4, 6}}}
	data, err := json.MarshalIndent(Sessions, "", "  ")
	if err != nil {
		log.Fatal("JSON Marshal is failure.")
	}
	fmt.Println(Sessions)
	fmt.Printf("%s\n", data)

	UnmarshalSessions := make(map[string]Session)
	if err := json.Unmarshal(data, &UnmarshalSessions); err != nil {
		log.Fatalf("JSON Unmarshal failure: %s", err)
	}
	fmt.Println(UnmarshalSessions)
}
