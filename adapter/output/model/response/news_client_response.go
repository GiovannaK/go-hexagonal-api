package response

type NewsClientResponse struct {
	Status       string
	TotalResults int
	Articles     []ArticleResponse
}

type ArticleResponse struct {
	Source      ArticleSourceResponse
	Id          string
	Name        string
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponse struct {
	ID *string
	Name string
}