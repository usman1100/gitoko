package main

import (
	"fmt"
	"os"

	"github.com/usman1100/gitoko/core/git"
)

func main() {
	if !git.IsCurrentDirectoryARepo() {
		fmt.Println("Current directory is not a git repository")
		os.Exit(1)
	}
	git.Checkout("test")

}
