package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tuanhuu162/go23_course/ex2/source"
)

var numericSortCmd = &cobra.Command{
	Use:   "int",
	Short: "Sorter for numeric elements",
	Long:  `sorter int [element] [element] ... A flexible sorting tool for numeric elements.`,
	Run: func(cmd *cobra.Command, args []string) {
		is_desc := viper.GetBool("desc")
		var sortType string = "int"
		source.PerformSort(sortType, args, is_desc)
	},
}
