package cmd

import (
	"encoding/json"
	"os"

	"github.com/reactiveNeon/jvc/internal/utils"
	"github.com/reactiveNeon/jvc/internal/vcs"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout [hash] [output.json]",
	Short: "Checkout a commit by hash",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hash := args[0]
		outputFile := args[1]

		commit, err := utils.LoadObject(hash)
		if err != nil {
			cmd.PrintErrf("Error loading commit: %v\n", err)
			return
		}

		treeHash, ok := commit["tree"].(string)
		if !ok {
			cmd.PrintErrf("Invalid commit object: %v\n", commit)
			return
		}

		jsonData, err := vcs.CheckoutJson(treeHash)
		if err != nil {
			cmd.PrintErrf("Error checking out JSON data: %v\n", err)
			return
		}

		prettyJson, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			cmd.PrintErrf("Error formatting JSON data: %v\n", err)
			return
		}

		if err := os.WriteFile(outputFile, prettyJson, 0644); err != nil {
			cmd.PrintErrf("Error writing JSON file: %v\n", err)
			return
		}

		cmd.Printf("Checked out commit %s to %s\n", hash, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
