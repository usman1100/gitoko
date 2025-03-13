package sanitize

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateBranchName(name string) error {
	// Git branch names must not be empty
	if len(name) == 0 {
		return errors.New("branch name cannot be empty")
	}

	// Must not exceed 255 characters
	if len(name) > 255 {
		return errors.New("branch name cannot exceed 255 characters")
	}

	// Must not be "." or ".."
	if name == "." || name == ".." {
		return errors.New("branch name cannot be '.' or '..'")
	}

	// Must not contain ASCII control characters or spaces
	if regexp.MustCompile(`[\x00-\x20\x7F]`).MatchString(name) {
		return errors.New("branch name cannot contain spaces or control characters")
	}

	// Must not start or end with certain characters: '/', '.', '~'
	if strings.HasPrefix(name, "/") {
		return errors.New("branch name cannot start with '/'")
	}
	if strings.HasSuffix(name, "/") {
		return errors.New("branch name cannot end with '/'")
	}
	if strings.HasSuffix(name, ".") {
		return errors.New("branch name cannot end with '.'")
	}
	if strings.HasSuffix(name, "~") {
		return errors.New("branch name cannot end with '~'")
	}

	// Must not contain certain special characters
	if strings.Contains(name, "^") {
		return errors.New("branch name cannot contain '^'")
	}
	if strings.Contains(name, ":") {
		return errors.New("branch name cannot contain ':'")
	}

	// Must not contain consecutive slashes
	if strings.Contains(name, "//") {
		return errors.New("branch name cannot contain consecutive slashes ('//')")
	}

	// Must not end with ".lock" (Git restriction)
	if strings.HasSuffix(name, ".lock") {
		return errors.New("branch name cannot end with '.lock'")
	}

	// Must not contain "@{" (used for Git ref syntax)
	if strings.Contains(name, "@{") {
		return errors.New("branch name cannot contain '@{'")
	}

	// If no errors, the branch name is valid
	return nil
}

func SanitizeCommit(commit string) string {
	// "\"8faf324dbd933b7c1d5f8f76d4e4327744081dbf chore: use promptui for tui\"\n\"0fa5bc4736a731ac50d7438b9acbd516e72daa3f init\""
	// '8faf324dbd933b7c1d5f8f76d4e4327744081dbf chore: use promptui for tui'

	cleaned := strings.Replace(commit, "\"", "", -1)
	cleaned = strings.Replace(cleaned, "\\n", "", -1)
	cleaned = strings.TrimSpace(cleaned)
	return cleaned
}

func SanitizeCommits(commits []string) []string {
	var cleaned []string
	for _, commit := range commits {
		cleaned = append(cleaned, SanitizeCommit(commit))
	}
	return cleaned
}

func SanitizeBranchName(name string) string {
	return strings.TrimSpace(name)
}
