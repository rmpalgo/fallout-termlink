package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/rmpalgo/fallout-termlink/internal/config"
	"github.com/rmpalgo/fallout-termlink/pkg/termlink"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "termlink",
	Short: "Run termlink",
	Long:  "Run terminal recreation of Fallout game's RobCo Termlink",
	Example: `termlink
termlink -v`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfg.Version {
			fmt.Printf("Termlink Version: %s\n", "1.0.0")
			return
		}

		err := termlink.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&cfg.Version, "version", "v", false, "Display version information")
}
