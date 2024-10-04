package handlers

import (
	// "fmt"
	// "context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	model "project1/models"
	"project1/mongodb"

	// "project1/mongodb"
	"reflect"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Reaction struct {
	Likes		int		`json:"likes" bson:"likes"`
	Dislikes	int		`json:"dislikes" bson:"dislikes"`
}

type post struct {
	ID 		int			`json:"id" bson:"_id"`
	Title	string		`json:"title" bson:"title"`	
	Body	string		`json:"body" bson:"body"`
	Tags	[]string	`json:"tags" bson:"tags"`
	Recations Reaction  `json:"reactions" bson:"reactions"`
	Views	int			`json:"views" bson:"views"`
	UserId  int			`json:"userId" bson:"userId"`
}

type ApiResponse struct {

	Posts []post `json:"posts" bson:"posts"`
	Total int	`json:"total" bson:"total"`
	Skip  int	`json:"skip" bson:"skip"`
	Limit int	`json:"limit" bson:"limit"`
	AA    int	`json:"aa,omitempty" bson:"aa"`
	BB    int	`json:"bb" binding:"required" bson:"bb"`

}


func FetchPosts(){
	url := "https://dummyjson.com/posts?skip=0&limit=251"

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

	var apiResponse model.ApiResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Panic("Error in fetching the posts", err)
	}

	// fmt.Println(apiResponse)
	// fmt.Println(apiResponse.Posts[0])
	fmt.Println(reflect.TypeOf(apiResponse.Posts[0]))
	// productCollection := mongodb.MongoClient.Database("testdb").Collection("post")
	// // singlePost := apiResponse.Posts[0]
	// result, err := productCollection.InsertOne(context.Background(), apiResponse)
    // if err != nil {
    //    fmt.Println("error inserting document: ", err)
    // }
	// fmt.Println(result)
    // return result.InsertedID, nil
	for _, singlePost := range apiResponse.Posts {
		mongodb.InsertOneDocument("post",singlePost)
	}
	
	// jsonData, err := json.MarshalIndent(apiResponse, "", "  ")
    // if err != nil {
    //     fmt.Println("Error marshalling to JSON:", err)
    //     return
    // }

    // // Print the pretty JSON
    // fmt.Println(string(jsonData))
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	filter := bson.M{
		"_id": id,
	}

	
}