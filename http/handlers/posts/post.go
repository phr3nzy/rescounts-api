package posts

type Post struct {
	ID         int64    `json:"id"`
	Author     string   `json:"author"`
	AuthorID   int64    `json:"authorId"`
	Likes      int64    `json:"likes"`
	Popularity float64  `json:"popularity"`
	Reads      int64    `json:"reads"`
	Tags       []string `json:"tags"`
}

type ApiJsonResponse struct {
	Posts []Post `json:"posts"`
}
