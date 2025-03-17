package git

import "os/exec"

func GetGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
