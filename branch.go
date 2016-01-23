package main

type Branch struct {
	Name    string   `json:"name"`
	IsHead  bool     `json:"isHead"`
	Commits []Commit `json:"commits"`
}
