package cmd

import (
	"github.com/reactiveNeon/jvc/internal/utils"
	"github.com/reactiveNeon/jvc/internal/vcs"
	"github.com/spf13/cobra"
)

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

		commitHash, err := vcs.StoreCommit(treeHash, "", "Initial commit")
		if err != nil {
			cmd.PrintErrf("Error storing commit: %v\n", err)
			return
		}

		cmd.Printf("Commit successful! Commit hash: %s\n", commitHash)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
