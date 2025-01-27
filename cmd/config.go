package cmd

import (
	// "github.com/m-tsuru/tt/lib"
	"fmt"
	"os"

	"github.com/m-tsuru/tt/lib"
	"github.com/spf13/cobra"
)

var global, edit bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and Set repository or global options of tt",
	Long:  `Get and Set repository or global options of tt`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var path string
		if global == true {
			err, path = lib.GetGlobalConfigPath(nil)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		} else {
			curDir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			_, err = os.Stat(curDir + "/.tt")
			if os.IsNotExist(err) {
				os.Mkdir(curDir+"/.tt", 0755)
			}
			path = curDir + "/.tt/config"
			f, err := os.Create(path)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			defer f.Close()
		}
		if edit == true {
			lib.EditGlobalConfig(path)
		} else {
			if len(args) == 1 {
				err, res := lib.GetConfig(args[0], path)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				fmt.Println(*res)
			} else if len(args) == 2 {
				err, res := lib.SetConfig(args[0], &args[1], path)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				fmt.Println(*res)
			} else {
				err := "Wrong Arguments."
				fmt.Printf("Error: %s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&global, "global", "g", false, "Get or Set Global Configuration")
	configCmd.Flags().BoolVarP(&edit, "edit", "e", false, "Edit Configuration with your text editor")
}
