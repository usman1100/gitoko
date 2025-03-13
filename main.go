package main

import (
	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/git"
)

func main() {

	commits, err := git.GetAllCommits()

	if err != nil {
		panic(err)
	}

	commitOptions := make([]huh.Option[string], len(commits))

	for i, commit := range commits {
		commitOptions[i] = huh.NewOption(commit, commit)
	}

	var Selections []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Pick a commit: ").
				Options(
					commitOptions...,
				).
				Value(&Selections),
		),
	)

	err = form.Run()
	if err != nil {
		panic(err)
	}

	for _, selection := range Selections {
		println(selection)
	}

}
