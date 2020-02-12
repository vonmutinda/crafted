package cmd

import (
	"github.com/spf13/cobra" 
)

var rootCmd = &cobra.Command{
	Use:     "Crafted",
	Aliases: []string{"kr"},
	Short:   "Crafted Backend Service",
	Long: `Crafted is an awesome pet project. It does absolutely nothing!`,
}
 

// Execute executes the root command.
func Execute() error { 
	return rootCmd.Execute()
}

