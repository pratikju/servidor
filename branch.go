package main

import (
	"github.com/libgit2/git2go"
	"log"
)

type Branch struct {
	Name    string   `json:"name"`
	IsHead  bool     `json:"isHead"`
	Commits []Commit `json:"commits"`
}

func GetBranches(repo *git.Repository) ([]Branch, error) {
	var branch Branch
	branches := make([]Branch, 0)

	itr, _ := repo.NewReferenceIterator()
	refs := getReferences(itr)

	revWalk, err := repo.Walk()
	if err != nil {
		log.Println(err)
		return branches, err
	}

	for i := 0; i < len(refs); i++ {
		branch = getBranch(refs[i], revWalk)
		branches = append(branches, branch)
	}

	return branches, nil
}

func getReferences(itr *git.ReferenceIterator) []*git.Reference {
	var ref *git.Reference
	refs := make([]*git.Reference, 0)
	var err error
	for {
		ref, err = itr.Next()
		if err != nil {
			break
		}
		refs = append(refs, ref)
	}
	return refs
}

func getBranch(ref *git.Reference, revWalk *git.RevWalk) Branch {
	var branch Branch
	b := ref.Branch()
	name, err := b.Name()
	if err != nil {
		log.Println(err)
	}
	isHead, err := b.IsHead()
	if err != nil {
		log.Println(err)
	}
	commits := GetCommits(ref.Target(), revWalk)
	branch = Branch{Name: name, IsHead: isHead, Commits: commits}
	return branch
}

func GetBranchByName(name string, repo *git.Repository) (Branch, bool) {
	var branch Branch
	gitBranch, err := repo.LookupBranch(name, git.BranchLocal)
	if err != nil {
		return branch, false
	}

	revWalk, err := repo.Walk()
	if err != nil {
		log.Println(err)
		return branch, false
	}
	return getBranch(gitBranch.Reference, revWalk), true
}
