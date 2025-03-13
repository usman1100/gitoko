package git

import (
	"fmt"
	"os/exec"
)

func Checkout(branch string) {
	cmd, err := exec.Command("git", "checkout", branch).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(cmd))
}

func IsCurrentDirectoryARepo() bool {
	cmd, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil {
		return false
	}

	return string(cmd) == "true\n"
}
