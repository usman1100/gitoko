package ui

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func InputBranchName() (string, error) {

	var branchName string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Branch Name: ").
				Validate(func(s string) error {
					if len(s) == 0 {
						return errors.New("branch name cannot be empty")
					}
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
