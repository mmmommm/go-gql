// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
