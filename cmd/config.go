package cmd

import (
	// "github.com/m-tsuru/tt/lib"
	"fmt"

	"github.com/m-tsuru/tt/lib"
	"github.com/spf13/cobra"
)

var global, edit bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and Set repository or global options of tt",
	Long:  `Get and Set repository or global options of tt`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if global == true {
			if edit == true {
				err, confPath := lib.GetGlobalConfigPath(nil)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				lib.EditGlobalConfig(confPath)
			} else {
				if len(args) == 1 {
					err, res := lib.GetGlobalConfigFromRaw(args[0], nil)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					}
					fmt.Println(*res)
				} else if len(args) == 2 {
					err, res := lib.SetGlobalConfigFromRaw(args[0], &args[1], nil)
					if err != nil {
						fmt.Printf("Error: %s\n", err)
					}
					fmt.Println(*res)
				}
			}
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&global, "global", "g", false, "Get or Set Global Configuration")
	configCmd.Flags().BoolVarP(&edit, "edit", "e", false, "Edit Configuration with your text editor")
}
