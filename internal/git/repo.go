package git

import (
	"os/exec"
)

func IsGitRepo() (bool, error) {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false, err
	}
	return true, nil
}



