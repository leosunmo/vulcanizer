package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdMapping)
	cmdMapping.AddCommand(cmdMappingShow)
	cmdMapping.AddCommand(cmdMappingsDiff)

}

var cmdMapping = &cobra.Command{
	Use:   "mapping",
	Short: "Display and Diff index Mappings.",
	Long:  `Use the show and diff subcommand.`,
}

var cmdMappingShow = &cobra.Command{
	Use:   "show",
	Short: "Display index mapping.",
	Long:  `This command will print out the mapping for the given index.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		v := getClient()

		mappings, err := v.GetPrettyIndexMappings(args[0])
		if err != nil {
			fmt.Printf("Error getting mappings for index %s - %s\n", args[0], err)
			os.Exit(1)
		}
		fmt.Println(mappings)
	},
}

var cmdMappingsDiff = &cobra.Command{
	Use:   "diff",
	Short: "Diff index mappings.",
	Long:  "Get the difference between index A's and index B's mappings.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		v := getClient()

		mappingsDiff, err := v.GetMappingDiff(args[0], args[1])
		if err != nil {
			fmt.Printf("Error getting mappings for indices %s & %s - %s\n", args[0], args[1], err)
			os.Exit(1)
		}
		fmt.Println(mappingsDiff)
	},
}
