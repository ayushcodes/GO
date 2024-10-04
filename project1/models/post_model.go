package model


type Reaction struct {
	Likes		int		`json:"likes" bson:"likes"`
	Dislikes	int		`json:"dislikes" bson:"dislikes"`
}

type Post struct {
	ID 		int			`json:"id" bson:"_id"`
	Title	string		`json:"title" bson:"title"`	
	Body	string		`json:"body" bson:"body"`
	Tags	[]string	`json:"tags" bson:"tags"`
	Recations Reaction  `json:"reactions" bson:"reactions"`
	Views	int			`json:"views" bson:"views"`
	UserId  int			`json:"userId" bson:"userId"`
}

type ApiResponse struct {

	Posts []Post `json:"posts" bson:"posts"`
	Total int	`json:"total" bson:"total"`
	Skip  int	`json:"skip" bson:"skip"`
	Limit int	`json:"limit" bson:"limit"`
	AA    int	`json:"aa,omitempty" bson:"aa"`
	BB    int	`json:"bb" binding:"required" bson:"bb"`

}