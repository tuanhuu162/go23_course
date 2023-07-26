package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tuanhuu162/go23_course/ex2/source"
)

func Init() {
	var input_list_int []string
	var input_list_string []string
	var input_list_mix []string

	var otherCmd = &cobra.Command{
		Use:   "sorter [sort_type] [element] [element]",
		Short: "A flexible sorting tool for numeric, string and mixed elements",
		Long:  `sorter [sort_type] [element] [element] ... A flexible sorting tool for numeric, string and mixed elements.`,
		Run: func(cmd *cobra.Command, args []string) {
			var sortType string
			if cmd.Flags().NFlag() > 1 {
				fmt.Println("Please use only one sorter '-int' or sorter '-string' or sorter '-mixed'")
				os.Exit(1)
			}
			if cmd.Flags().Changed("-int") {
				sortType = "int"
			} else if cmd.Flags().Changed("-string") {
				sortType = "string"
			} else if cmd.Flags().Changed("-mix") {
				sortType = "mix"
			} else {
				fmt.Println("Please use sorter '-int' or sorter '-string' or sorter '-mixed'")
				os.Exit(1)
			}
			is_desc := false
			source.PerformSort(sortType, args, is_desc)
		},
	}
	otherCmd.Flags().StringArrayVarP(&input_list_int, "-int", "i", []string{}, "Sort for integer/float type")
	otherCmd.Flags().StringArrayVarP(&input_list_string, "-string", "s", []string{}, "Sort for string type")
	otherCmd.Flags().StringArrayVarP(&input_list_mix, "-mix", "m", []string{}, "Sort for mixed type")

	if err := otherCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
