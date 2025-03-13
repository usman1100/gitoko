package git

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/sanitize"
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

func GetAllCommits() ([]string, error) {
	// git log --all --pretty=format:"%H %s"
	cmd, err := exec.Command("git", "log", "--all", "--pretty=format:\"%H %s\"").Output()
	if err != nil {
		return nil, errors.New("could not get commits")
	}

	commits := strings.Split(string(cmd), "\n")
	return sanitize.SanitizeCommits(commits), nil

}

func GetAllCommitsAsOptions() ([]huh.Option[string], error) {
	commits, err := GetAllCommits()
	if err != nil {
		return nil, err
	}
	commitOptions := make([]huh.Option[string], len(commits))

	for i, commit := range commits {
		commitOptions[i] = huh.NewOption(commit, commit)
	}

	return commitOptions, nil
}
