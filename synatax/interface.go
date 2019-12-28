 package main

import (
	"fmt"
	"os"
)

type GitController interface {
	// Should return status as bool and link to repo
	Push() (bool, string)
}

type Github struct {
	repo, branch, owner string
}

type Gitlab struct {
	repo, branch, owner string
}

type Localhost struct {
	repo, branch, owner string
}

func (gh Github) Push() (bool, string) {
	changes := fmt.Sprintf("https://github.com/%s/%s/tree/%s", gh.owner, gh.repo, gh.branch)
	return true, changes 
}

func (gl Gitlab) Push() (bool, string) {
	changes := fmt.Sprintf("https://gitlab.com/%s/%s/tree/%s", gl.owner, gl.repo, gl.branch) 
	return true, changes 
}

func (gl Localhost) Push() (bool, string) {
	changes := fmt.Sprintf("http://localhost/%s/%s/tree/%s", gl.owner, gl.repo, gl.branch)

	// failing to test realMain() os.Exit
	return false, changes 
}

func realMain() int {
	var github GitController = Github {
		owner: "pwlae",
		repo: "golang-examples",
		branch: "master",
	}

	var gitlab GitController = Gitlab {
		owner: "pwlae",
		repo: "golang-examples",
		branch: "master",
	}

	var localhost GitController = Localhost {
		owner: "pwlae",
		repo: "golang-examples",
		branch: "master",
	}

	// Should return true and repo link
	status, repo := github.Push()
	if status != true {
		return 1
	}
	fmt.Println(repo)

	// Should return true and repo link
	status, repo = gitlab.Push()
	if status != true {
		return 1
	}
	fmt.Println(repo)

	// Should return false and repo link
	status, repo = localhost.Push()
	if status != true {
		return 1
	}
	fmt.Println(repo)

	return 0
}

func main() {
	os.Exit(realMain())
}