package main

import (
	"fmt"
	"os"

	"github.com/usman1100/gitoko/core/git"

	"github.com/manifoldco/promptui"
)

func main() {
	if !git.IsCurrentDirectoryARepo() {
		fmt.Println("Current directory is not a git repository")
		os.Exit(1)
	}
	prompt := promptui.Prompt{
		Label: "Branch name",
	}

	inputBranchName, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if inputBranchName == "" {
		inputBranchName = "main"
	}

	err = git.Checkout(inputBranchName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Switched to " + inputBranchName + " branch")

}
