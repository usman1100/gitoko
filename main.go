package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/usman1100/gitoko/core/git"
	"github.com/usman1100/gitoko/core/ui"
)

func main() {

	branchName, err := ui.InputBranchName()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var commits []string

	if branchName == "" {
		commits, err = git.GetAllCommits()
	} else {
		commits, err = git.GetOnlyBranchCommits(branchName)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	commitOptions := git.CommitsToOptions(commits)

	selections, err := ui.InputCommitSelection(commitOptions)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("You selected " + strconv.Itoa(len(selections)) + " commits\n")
	for _, selection := range selections {
		println(selection)
	}
	{
		var input string
		fmt.Print("Press Enter to start cherry-picking: ")
		fmt.Println("or to abort, press ctrl+c")
		fmt.Println()
		fmt.Scanln(&input)
	}

	err = ui.InuptMultipleCherryPickingPrompts(selections)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Cherry-picking completed successfully")

}
