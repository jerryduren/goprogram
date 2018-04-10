package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Session struct {
	SessionID int
	Imsi      string `json:"IMSI"`

	UpfTeid int
	NrTeid  int
	Qfis    []Qfi
}

type Qfi struct {
	QfiID    int
	Priority int `json:"PRI"`
	ARP      int
	FilterID int
}

//var SessionS map[string]Session

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HtmlUrl   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in markdown format
}
type User struct {
	Login   string
	HtmlUrl string `json:"html_url"`
}

/* search issues from github.com */
func SearchIssues(sterm []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(sterm, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		log.Fatalf("Get issues is failure from %s\n, And the error is: \"%s\"", IssueURL+"?q="+q, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Search Issues Failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { //有多个json实体的时候需要用json.Decode方法来解码
		log.Fatalf("Resutl of Search is error, the erros is %s\n", err)
		return nil, err
	}

	fmt.Printf("%s\n", resp.Body)
	var f interface{}
	f = json.NewDecoder(resp.Body).Decode(&f)
	PrintUnknownJson(f)

	return &result, nil
}

func PrintUnknownJson(f interface{}) {
	if f == nil {
		fmt.Println("the variable is nil.")
		return
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}

	}
}

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

	result, err := SearchIssues([]string{"repo:golang/go", "is:open", "json decoder"})
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Search Issues %d results\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
