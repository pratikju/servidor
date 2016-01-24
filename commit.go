package main

import (
	"github.com/libgit2/git2go"
)

type Commit struct {
	Message    string         `json:"message"`
	ID         string         `json:"id"`
	ObjectType string         `json:"object_type"`
	Author     *git.Signature `json:"author"`
}

func GetCommits(oid *git.Oid, revWalk *git.RevWalk) []Commit {
	var commit Commit
	var commits []Commit

	err := revWalk.Push(oid)
	if err != nil {
		return commits
	}
	f := func(c *git.Commit) bool {
		commit = Commit{Message: c.Summary(), ID: c.Id().String(), ObjectType: c.Type().String(), Author: c.Author()}
		commits = append(commits, commit)
		return true
	}
	_ = revWalk.Iterate(f)
	return commits
}
