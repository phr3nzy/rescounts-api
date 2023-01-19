package posts

// Post is the model for a single post.
type Post struct {
	ID         int64    `json:"id"`
	Author     string   `json:"author"`
	AuthorID   int64    `json:"authorId"`
	Likes      int64    `json:"likes"`
	Popularity float64  `json:"popularity"`
	Reads      int64    `json:"reads"`
	Tags       []string `json:"tags"`
}

// ApiJsonResponse is the model for the JSON response from the API.
type ApiJsonResponse struct {
	Posts []Post `json:"posts"`
}
