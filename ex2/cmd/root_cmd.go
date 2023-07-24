package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sorter [sort_type] [element] [element]",
	Short: "A flexible sorting tool for numeric, string and mixed elements",
	Long:  `sorter [sort_type] [element] [element] ... A flexible sorting tool for numeric, string and mixed elements.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use sorter int or sorter string or sorter mixed")
	},
}

func Execute() {
	rootCmd.PersistentFlags().Bool("desc", false, "Sort in descending order")
	viper.BindPFlag("desc", rootCmd.PersistentFlags().Lookup("desc"))
	viper.SetDefault("desc", false)

	rootCmd.AddCommand(numericSortCmd)
	rootCmd.AddCommand(stringSortCmd)
	rootCmd.AddCommand(mixSortCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
