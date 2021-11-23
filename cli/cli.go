package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A generator for Cobra based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runE,
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("jobs", "j", []string{"a", "b", "c"}, "for, well, things")
}

func runE(cmd *cobra.Command, args []string) error {
	argsss, err := cmd.PersistentFlags().GetStringSlice("jobs")
	if err != nil {
		return err
	}

	log.Printf("%v\n", argsss)

	return nil
}

func main() {
	rootCmd.Execute()
}
