package cmd

import (
	"fmt"
  "os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "qbt-exporter",
  Short: "QBT Exporter is a Prometheus exporter for Qbittorrent",
  Long: `A simple exporter that polls qbittorrent for statistics
        and publishes them as Prometheus data`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
