package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/usman1100/gitoko/core/sanitize"
)

func Checkout(branch string) error {
	out := exec.Command("git", "checkout", branch)
	out.Dir = "."
	_, err := out.Output()

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
	out := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	out.Dir = "."
	cmd, err := out.Output()
	if err != nil {
		return false
	}

	return string(cmd) == "true\n"
}

func GetCurrentBranch() (string, error) {
	out := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out.Dir = "."
	cmd, err := out.Output()
	if err != nil {
		return "", errors.New("could not get current branch")
	}

	return string(cmd), nil
}

func GetAllCommits() ([]string, error) {
	// git log --all --pretty=format:"%H %s"
	out := exec.Command("git", "log", "--all", "--pretty=format:\"%H %s\"")
	out.Dir = "."
	cmd, err := out.Output()
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

func CommitsToOptions(commits []string) []huh.Option[string] {
	commitOptions := make([]huh.Option[string], len(commits))

	for i, commit := range commits {
		commitOptions[i] = huh.NewOption(commit, commit)
	}

	return commitOptions
}

func FetchBranch(branch string) error {
	out := exec.Command("git", "fetch", "origin", branch)
	out.Dir = "."
	_, err := out.Output()
	if err != nil {
		return errors.New("could not fetch branch " + branch)
	}

	return nil
}

func GetOnlyBranchCommits(branch string) ([]string, error) {
	// git log --pretty=format:"%H %s" branch
	out := exec.Command("git", "log", "--pretty=format:\"%H %s\"", branch)
	out.Dir = "."
	cmd, err := out.Output()
	if err != nil {
		fmt.Println("branch not found locally, do you want to fetch it?")
		fmt.Println("Press Enter to fetch or type 'no' and press Enter to continue without fetching")

		var input string
		fmt.Scanln(&input)

		if input == "no" {
			return nil, errors.New("branch not found")
		}

		fmt.Println("Fetching branch " + branch)
		err := FetchBranch(branch)
		if err != nil {
			return nil, errors.New("could not fetch branch")
		}
	}

	commits := strings.Split(string(cmd), "\n")
	return sanitize.SanitizeCommits(commits), nil
}

func GetAllLocalBranches() ([]string, error) {
	out := exec.Command("git", "branch", "--list")
	out.Dir = "."
	cmd, err := out.Output()
	if err != nil {
		return nil, errors.New("could not get branches")
	}

	branches := strings.Split(string(cmd), "\n")
	for i, branch := range branches {
		branches[i] = sanitize.SanitizeBranchName(branch)
	}
	return branches, nil
}

func CherryPick(commit string) error {
	out := exec.Command("git", "cherry-pick", commit)
	out.Dir = "."
	_, err := out.Output()
	if err != nil {
		fmt.Println("\nPlease check for conflicts")
		fmt.Println("If all conflicts are resolved, press Enter to continue")
		fmt.Println("Or to abort, type 'abort' and press Enter")

		var input string
		fmt.Scanln(&input)

		if input == "" {
			_, err := exec.Command("git", "cherry-pick", "--continue").Output()
			if err != nil {
				return errors.New("could not continue cherry-pick")
			}
			fmt.Println("Cherry-pick continued")
		} else if input == "abort" {
			_, err := exec.Command("git", "cherry-pick", "--abort").Output()
			if err != nil {
				return errors.New("could not abort cherry-pick")
			}
			fmt.Println("Cherry-pick aborted")
		} else {
			return errors.New("invalid input")
		}

		return nil

	}
	return nil
}
