package git

import (
	"errors"
	"os/exec"
)

func Checkout(branch string) error {
	_, err := exec.Command("git", "checkout", branch).Output()

	if branch == "" {
		defaultBranch, err := GetCurrentBranch()
		if err != nil {
			return errors.New("could not get current branch")
		}
		branch = defaultBranch
	}

	if err != nil {
		return errors.New("could not switch to branch " + branch)
	}

	return nil
}

func IsCurrentDirectoryARepo() bool {
	cmd, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		return false
	}

	return string(cmd) == "true\n"
}

func GetCurrentBranch() (string, error) {
	cmd, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", errors.New("could not get current branch")
	}

	return string(cmd), nil
}
