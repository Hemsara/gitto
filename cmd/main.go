package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Hemsara/gitto/config"
	"github.com/Hemsara/gitto/internal/ai"

	"github.com/Hemsara/gitto/internal/git"
	"github.com/urfave/cli/v3"
)

func main() {
	config.LoadConfig()

	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}

func Execute() error {
	cmd := &cli.Command{
		Name:   "gitto",
		Usage:  "Git commits on steroids",
		Action: PerformCommit,
	}

	return cmd.Run(context.Background(), []string{"gitto"})
}

func PerformCommit(ctx context.Context, cmd *cli.Command) error {
	if _, err := git.IsGitRepo(); err != nil {
		log.Fatalf("❌ Not a git repository: %v", err)
	}

	gitDiff, err := git.GetGitDiff()
	if err != nil {
		log.Fatalf("❌ Failed to get git diff: %v", err)
	}

	if strings.TrimSpace(gitDiff) == "" {
		log.Println("⚠️ No changes to commit.")
		return nil
	}

	commitMessage, err := ai.GenerateCommitMessage(gitDiff)
	if err != nil {
		log.Fatalf("❌ Failed to generate commit message: %v", err)
	}

	fmt.Printf("\n💡 Generated commit message:\n%s\n", commitMessage)

	fmt.Print("\n✅ Commit? (y/N): ")
	var confirm string
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) != "y" {
		log.Println("❌ Commit canceled.")
		return nil
	}

	cmdExecAdd := exec.Command("git", "add", "-A")
	if err := cmdExecAdd.Run(); err != nil {
		log.Fatalf("❌ Failed to stage files: %v", err)
	}

	output, err := git.Commit(commitMessage)
	if err != nil {
		log.Fatalf("❌ Failed to commit: %v", err)
	}

	fmt.Println("✅ Commit successful!")
	fmt.Println(output)

	return nil
}
