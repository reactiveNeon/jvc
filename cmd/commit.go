package cmd

import (
	"os"

	"github.com/reactiveNeon/jvc/internal/utils"
	"github.com/reactiveNeon/jvc/internal/vcs"
	"github.com/spf13/cobra"
)

var commitMessage string

var commitCmd = &cobra.Command{
	Use:   "commit [file.json]",
	Short: "Commit a JSON file",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		jsonData, err := utils.ReadJsonFile(filePath)
		if err != nil {
			cmd.PrintErrf("Error reading JSON file: %v\n", err)
			return
		}

		treeHash, err := vcs.StoreJson(jsonData)
		if err != nil {
			cmd.PrintErrf("Error storing JSON data: %v\n", err)
			return
		}

		parentHash, err := utils.ReadHead()
		if err != nil {
			if !os.IsNotExist(err) {
				cmd.PrintErrf("Error reading HEAD: %v\n", err)
				return
			}
		}

		if parentHash != "" {
			prevTreeHash, err := utils.GetTreeHashFromCommitHash(parentHash)
			if err != nil {
				cmd.PrintErrf("Error getting previous tree hash: %v\n", err)
				return
			}

			if treeHash == prevTreeHash {
				cmd.PrintErrf("No changes detected. Nothing to commit.\n")
				return
			}
		}

		commitHash, err := vcs.StoreCommit(treeHash, parentHash, commitMessage)
		if err != nil {
			cmd.PrintErrf("Error storing commit: %v\n", err)
			return
		}

		cmd.Printf("Commit successful! Commit hash: %s\n", commitHash)

		utils.WriteHead(commitHash)
		if err != nil {
			cmd.PrintErrf("Error writing HEAD: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
	commitCmd.MarkFlagRequired("message")
}
