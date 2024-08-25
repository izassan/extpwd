/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/izassan/extpwd/lib"
	"github.com/spf13/cobra"
)
const flagName = "replace-rules"

var rootCmd = &cobra.Command{
	Use:   "extpwd",
	Short: "extend pwd",
	RunE: func(cmd *cobra.Command, args []string) error{
		replaceRuleStr, err := cmd.Flags().GetString(flagName)
		if err != nil{
			return err
		}
		currentDir, err := os.Getwd()
		if err != nil{
			return err
		}
		if replaceRuleStr != ""{
			rules := lib.ParseOptionString(replaceRuleStr)
			for _, r := range rules{
				currentDir = r.ReplaceString(currentDir)
			}
		}
		fmt.Print(currentDir)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("replace-rules", "r", "", "specify rules for replacing path string")
}
