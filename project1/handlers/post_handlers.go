package handlers

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Reaction struct {
	Likes		int		`json:"likes"`
	Dislikes	int		`json:"dislikes"`
}

type Post struct {
	ID 		int			`json:"id"`
	Title	string		`json:"title"`	
	Body	string		`json:"body"`
	Tags	[]string	`json:"tags"`
	Recations Reaction  `json:"reactions"`
	Views	int			`json:"views"`
	UserId  int			`json:"userId"`
}

type ApiResponse struct {

	Posts []Post `json:"posts"`
	Total int	`json:"total"`
	Skip  int	`json:"skip"`
	Limit int	`json:"limit"`

}


func FetchPosts(){
	url := "https://dummyjson.com/posts"

	resp, err := http.Get(url)
	if err !=nil {
		log.Panic("Error in calling fetch post api", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panic("Error: Fetching the posts")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error in fetching the posts")
	}

	var apiResponse ApiResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Panic("Error in fetching the posts", err)
	}

	fmt.Println(apiResponse)
}