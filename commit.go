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

func GetCommits(oid *git.Oid, revWalk *git.RevWalk) []Commit {
	var commit Commit
	commits := make([]Commit, 0)

	err := revWalk.Push(oid)
	if err != nil {
		return commits
	}
	f := func(c *git.Commit) bool {
		commit = Commit{Message: c.Summary(), Id: c.Id().String(), ObjectType: c.Type().String(), Author: c.Author()}
		commits = append(commits, commit)
		return true
	}
	_ = revWalk.Iterate(f)
	return commits
}
