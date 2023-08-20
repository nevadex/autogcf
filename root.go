package main

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "autogcf",
	Short: "Automatically host a website on Google Cloud Functions",
	Long: `AutoGCF automatically reads a static website, writes functions, and
deploys your website and associated functions to Google Cloud Functions.

Your static website can be up and running in minutes!
github.com/nevadex/autogcf`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := AutoGCF(); err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var (
	SourceDir      string
	FileExtensions string
	AutoGenDir     string
	ModulePath     string
	Region         string
	Memory         string
	MaxInstances   string
	MaxAgeHTML     string
	MaxAgeOther    string

	DeleteGenFiles    bool
	SkipInteraction   bool
	ReadableFunctions bool
)

func init() {
	rootCmd.Flags().StringVarP(&SourceDir, "source", "s", ".", "Source directory to read files from")
	rootCmd.Flags().StringVarP(&FileExtensions, "filetypes", "f", "html,css,js", "Comma-separated list of file extensions without the dot that will be public-facing on your website")
	rootCmd.Flags().StringVarP(&AutoGenDir, "generated", "g", "autogcf-gen", "Directory to create and fill with code and website")
	rootCmd.Flags().StringVarP(&ModulePath, "modulePath", "p", "autogcf.gen/", "Path to module location in generated go files")

	rootCmd.Flags().StringVarP(&Region, "region", "r", "us-central1", "Google Cloud region to deploy functions to")
	rootCmd.Flags().StringVarP(&Memory, "memory", "m", "128Mi", "Memory limit for each function, corresponding to the vCPU limit as well")
	rootCmd.Flags().StringVarP(&MaxInstances, "maxInstances", "i", "10", "Maximum number of instances of each function")
	rootCmd.Flags().StringVar(&MaxAgeHTML, "maxAgeHTML", "", "Specify Max-Age in seconds to reduce bandwidth on only HTML files")
	rootCmd.Flags().StringVar(&MaxAgeOther, "maxAgeOther", "", "Specify Max-Age in seconds to reduce bandwidth on all non-HTML files")

	rootCmd.Flags().BoolVarP(&DeleteGenFiles, "delete", "d", false, "Delete autogenerated files after successful deployment")
	rootCmd.Flags().BoolVarP(&SkipInteraction, "no-interactive", "y", false, "Skip all interactive portions of the deployment")
	rootCmd.Flags().BoolVarP(&ReadableFunctions, "format-functions", "c", false, "Format generated code to be readable")
}