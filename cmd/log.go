package cmd

import (
	"fmt"
	"time"

	"github.com/reactiveNeon/jvc/internal/utils"
	"github.com/spf13/cobra"
)

var showAll bool

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit history",
	Run: func(cmd *cobra.Command, args []string) {
		if showAll {
			commitHashes, err := utils.GetAllCommitHashes()
			if err != nil {
				cmd.PrintErrf("Error finding commits: %v\n", err)
				return
			}

			// TODO: Sort commit hashes by timestamp
			for _, commitHash := range commitHashes {
				obj, err := utils.LoadObject(commitHash)
				if err != nil {
					continue
				}
				message := obj["message"]
				if ts, ok := obj["timestamp"].(float64); ok {
					t := time.Unix(int64(ts), 0).UTC()
					fmt.Printf("commit %s\nDate:   %s\n\n    %s\n\n", commitHash, t.Format(time.RFC3339), message)
				}
			}

			return
		}

		commitHash, err := utils.ReadHead()
		if err != nil {
			cmd.PrintErrf("Error reading HEAD: %v\n", err)
			return
		}

		for commitHash != "" {
			obj, err := utils.LoadObject(commitHash)
			if err != nil {
				cmd.PrintErrf("Error loading commit %s: %v\n", commitHash, err)
				break
			}

			message := obj["message"]
			if ts, ok := obj["timestamp"].(float64); ok {
				t := time.Unix(int64(ts), 0).UTC()
				fmt.Printf("commit %s\nDate:   %s\n\n    %s\n\n", commitHash, t.Format(time.RFC3339), message)
			}

			parent, ok := obj["parent"].(string)
			if ok {
				commitHash = parent
			} else {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVar(&showAll, "all", false, "Show all commits, not just from HEAD")
}
