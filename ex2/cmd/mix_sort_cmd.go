package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tuanhuu162/go23_course/ex2/source"
)

var mixSortCmd = &cobra.Command{
	Use:   "mix",
	Short: "Sorter for mixed elements including numeric and string",
	Long:  `sorter mix [element] [element] ... A flexible sorting tool for mixed elements including numeric and string.`,
	Run: func(cmd *cobra.Command, args []string) {
		is_desc := viper.GetBool("desc")
		var sortType string = "mix"
		source.PerformSort(sortType, args, is_desc)
	},
}
