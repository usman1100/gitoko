package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/git"
	"github.com/usman1100/gitoko/core/ui"
)

func main() {

	branchName, err := ui.InputBranchName()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	commits, err := git.GetOnlyBranchCommits(branchName)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	commitOptions := git.CommitsToOptions(commits)

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
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("You selected " + strconv.Itoa(len(Selections)) + " commits\n")
	for _, selection := range Selections {
		println(selection)
	}

}
