package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of qbt-exporter",
  Long:  `This is the version number of qbt-exporter`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("QBT-Exporter 0.01 -- HEAD")
  },
}
