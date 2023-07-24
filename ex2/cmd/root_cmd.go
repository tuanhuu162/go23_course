package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sorter.go [sort_type] [element] [element] ...",
	Short: "A flexible sorting tool for numeric, string and mixed elements",
	Long:  `A flexible sorting tool for numeric, string and mixed elements.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var sortType string = args[0]
		var element []string = args[1:]
		sortingFactory := SortingFactory{}
		sorter := sortingFactory.GetSorter(sortType)
		var finalValue string = sorter.SortElement(element)
		fmt.Println("Output:", finalValue)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
