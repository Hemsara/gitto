package git

import (
	"os/exec"
	"strings"
)

func GetGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return TruncateDiff(string(output)), nil
}

func TruncateDiff(diff string) string {
	lines := strings.Split(diff, "\n")
	var result []string
	for _, line := range lines {
		if len(result) > 100 {
			break
		}

		if strings.HasPrefix(line, "+") || strings.HasPrefix(line, "-") || strings.Contains(line, "diff") {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}
