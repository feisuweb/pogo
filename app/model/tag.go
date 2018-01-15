package model

import "strings"

// Tag is tag struct of post
type Tag struct {
	Name string
	URL  string
}

// NewTag returns new Tag with name and proper url
func NewTag(name string) *Tag {
	name = strings.TrimSpace(name)
	return &Tag{
		Name: name,
		URL:  "/tags/" + name + ".html",
	}
}
