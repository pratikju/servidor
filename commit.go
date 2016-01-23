package main

import (
	"github.com/libgit2/git2go"
)

type Commit struct {
	Message    string         `json:"message"`
	Id         string         `json:"id"`
	ObjectType string         `json:"object_type"`
	Author     *git.Signature `json:"author"`
}
