package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/git"
)

func InputBranchName() (string, error) {

	var branchName string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Pick a branch name to pick commits from or leave empty for ALL commits ").
				Validate(func(s string) error {
					if len(s) > 255 {
						return errors.New("branch name cannot exceed 255 characters")
					}
					return nil
				},
				).
				Placeholder("Enter branch name").
				Value(&branchName),
		),
	)
	err := form.Run()
	if err != nil {
		return "", errors.New("invalid branch name")
	}

	return branchName, nil
}

func InputCommitSelection(options []huh.Option[string]) ([]string, error) {

	var Selections []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Pick a commit: ").
				Height(20).
				Options(
					options...,
				).
				Value(&Selections),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, errors.New("invalid commit selection")
	}

	return Selections, nil
}

func InuptMultipleCherryPickingPrompts(commits []string) error {

	var input string
	for _, commit := range commits {
		println("Cherry-picking commit: " + commit)
		println("Press Enter to start cherry-picking: ")
		println("or to abort, press ctrl+c")
		println("or to skip this commit, press s")
		println("or to abort all cherry-picking, press a")

		fmt.Scanln(&input)

		if input == "s" {
			continue
		} else if input == "a" {
			return errors.New("cherry-picking")
		} else {
			commitHash := strings.Split(commit, " ")[0]
			err := git.CherryPick(commitHash)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
