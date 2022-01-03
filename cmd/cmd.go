package cmd

import (
	"fmt"
	"os"

	//"github.com/hartfordfive/csv-to-openmetrics/version"

	"github.com/hartfordfive/csv-to-openmetrics/generator"
	"github.com/hartfordfive/csv-to-openmetrics/version"
	"github.com/spf13/cobra"
)

// Execute executes the root command.
func Execute() error {
	return entry.Execute()
}

var (
	FlagConfigPath string
	FlagOutputFile string
	FlagLogLevel   string
)

var (
	entry = &cobra.Command{
		Use:   "csv-to-openmetrics",
		Short: "Application to generate OpenTSDB formated metrics files from CSV input",
		Long: `Long multiline description
to go here.`,
	}
)

func init() {
	GenerateCmd.Flags().StringVarP(&FlagConfigPath, "config", "c", "", "Path to the configuration file")
	GenerateCmd.Flags().StringVarP(&FlagOutputFile, "output", "o", "", "Path OpenTSDB output file")
	GenerateCmd.Flags().StringVarP(&FlagLogLevel, "log-level", "l", "", "Set log level (DEBUG,INFO,WARN,ERROR)")
	entry.AddCommand(GenerateCmd, VersionCmd)
}

// GenerateCmd is used to initialize the "run" sub-command under the n2p-script-executor
var GenerateCmd = &cobra.Command{
	Use:   "generate ",
	Short: "Run the conversion",
	Long:  `Runs the script execution, which will run all scripts in the specified directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		//logging.SetLogLevel(FlagLogLevel)
		g, err := generator.New(FlagConfigPath)
		if err != nil {
			fmt.Errorf("[ERROR] Could not initiate CSV to OpenTSDB generator: %v", err)
			os.Exit(1)
		}
		for _, f := range g.Config.Files {
			if err := g.ConvertToOpenMetricsFormat(f, FlagOutputFile); err != nil {
				fmt.Errorf("[ERROR] Could not initiate CSV to OpenTSDB generator: %v", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	},
}

// VersionCmd is used to initialize the "version" sub-command under the n2p-script-executor
var VersionCmd = &cobra.Command{
	Use:   "version ",
	Short: "Show version",
	Long:  `Show the version and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion()
		os.Exit(0)
	},
}
