# Gitoko

Gitoko is an interactive terminal user interface (TUI) that simplifies Git cherry-picking. Instead of manually hunting down commit hashes and running multiple Git commands, Gitoko provides a streamlined, visual way to browse, select, and apply commits.

## Installtion

## Mac/Linux One-liner

```bash
curl -fsSL https://raw.githubusercontent.com/usman1100/gitoko/refs/heads/install-script-docs/docs/install-unix.sh | bash
```

## Windows
```bash
irm https://raw.githubusercontent.com/usman1100/gitoko/refs/heads/install-script-docs/docs/install-windows.ps1 | iex
```

## Features

- Interactive UI – Navigate through commit history with an easy-to-use terminal interface.

- Batch Cherry-Picking – Select multiple commits at once and apply them in one go.

- Commit Preview – View commit details before applying changes.

- Branch Selection – Easily switch between branches to pick commits from different sources.

- Minimal Setup – Works seamlessly with existing Git repositories.



## Usage

Launch Gitoko inside any Git repository to start cherry-picking interactively:


- Navigate to the Git repository in your terminal.
- Launch gitoko by typing `gitoko` in the terminal.
- Enter a branch name to pick commits from. Or skip to use all the commits in the current branch.
- Browse through the commit history using arrow keys and search filter.
- You can search via commit message or hash
- Select commits to cherry-pick.
- Start the interative cherry-picking process by pressing `Enter`.
- The app then ask you for confirmation for each commit.
- Once done, the app wil exit.


## Why Use Gitoko?

Cherry-picking in Git can be a tedious task, especially when dealing with multiple commits. Gitoko eliminates the hassle by providing a user-friendly interface, reducing errors, and improving workflow efficiency. Whether you're managing hotfixes, backporting changes, or selectively applying commits, Gitoko makes the process smoother and faster.



## License

Gitoko is open-source and available under the MIT License.

