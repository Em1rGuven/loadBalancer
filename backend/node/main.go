package main

import (
	"github.com/spf13/cobra"
)

const API_KEY = "2a007225"

var port int
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "API sunucusu",
	Run: func(cmd *cobra.Command, args []string) {
		server := newServer()
		server.start(port)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 8081, "Sunucunun dinleyeceği port")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
