package main

import (
	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/git"
)

func main() {

	commitOptions, err := git.GetAllCommitsAsOptions()

	if err != nil {
		panic(err)
	}

	var Selections []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Pick a commit: ").
				Height(20).
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
