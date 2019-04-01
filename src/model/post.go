package model

// Post model
type Post struct {
	Common

	Title    string
	Content  string
	Author   string
	Tags     []*Tag
	Comments []*Comment
}
