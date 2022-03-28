package entity

// Comment represents a comment object as returned by the JSON Placeholder API:
// https://jsonplaceholder.typicode.com/comments
type Comment struct {
	ID     intProperty `json:"id"`
	PostID intProperty `json:"postId"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Body   string      `json:"body"`
}
