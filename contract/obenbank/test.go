package main

import (
	"encoding/json"
	"fmt"
)

type UserRating struct {
	User    string  `json:"user"`
	Average float64 `json:"average"`
	Rates   []Rate  `json:"rates"`
}

type Rate struct {
	ProjectTitle string  `json:"projecttitle`
	Score        float64 `json:"score"`
}

func main() {

	var user = UserRating{User: "quotia72@naver.com", Average: 0}
	fmt.Println(user)
	userAsBytes, _ := json.Marshal(user)
	fmt.Println(userAsBytes)
	fmt.Println("Hello World")
	var res = UserRating{}
	fmt.Println(json.Unmarshal(userAsBytes, &res))
}
