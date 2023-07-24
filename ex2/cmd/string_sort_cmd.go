package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tuanhuu162/go23_course/ex2/source"
)

var stringSortCmd = &cobra.Command{
	Use:   "string",
	Short: "Sorter for string elements",
	Long:  `sorter string [element] [element] ... A flexible sorting tool for string elements.`,
	Run: func(cmd *cobra.Command, args []string) {
		is_desc := viper.GetBool("desc")
		var sortType string = "string"
		source.PerformSort(sortType, args, is_desc)
	},
}
