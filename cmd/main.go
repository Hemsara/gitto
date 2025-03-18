package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	// "github.com/Hemsara/gitto/config"
	"github.com/Hemsara/gitto/internal/ai"
	"github.com/Hemsara/gitto/internal/keys"

	"github.com/Hemsara/gitto/internal/git"
	"github.com/urfave/cli/v3"
)

func main() {

	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}

func Execute() error {
	cmd := &cli.Command{
		Name:  "gitto",
		Usage: "Git commits on steroids",

		Commands: []*cli.Command{
			{
				Name:  "config",
				Usage: "Configure gitto",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "apikey",
						Aliases: []string{"k"},
						Usage:   "API key for authentication",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					apiKey := cmd.String("apikey")
					if apiKey == "" {
						return fmt.Errorf("API key is required")
					}

					err := keys.SaveAPIKey(apiKey)
					if err != nil {
						return fmt.Errorf("failed to save API key: %v", err)
					}

					return nil
				},
			},
			{
				Name:   "commit",
				Usage:  "Commit changes",
				Action: PerformCommit,
			},
		},
	}

	// This is the main issue - you're passing "gitto" as an argument which isn't needed
	// When running "go run main.go commit", the os.Args will be ["main.go", "commit"]
	// So we should use os.Args directly
	return cmd.Run(context.Background(), os.Args)
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
