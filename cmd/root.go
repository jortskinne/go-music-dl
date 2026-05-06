package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	outputDir string
	quality   string
	verbose   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-music-dl",
	Short: "A music downloader supporting multiple platforms",
	Long: `go-music-dl is a command-line tool for downloading music from
various streaming platforms including NetEase Cloud Music, QQ Music,
Kugou, Kuwo, and more.

Example:
  go-music-dl search "周杰伦 稻香"
  go-music-dl download --id 123456 --source netease
  go-music-dl batch --file songs.txt`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $HOME/.go-music-dl.yaml)")
	// Default output directory changed to ~/Music for convenience
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "~/Music", "output directory for downloaded files")
	rootCmd.PersistentFlags().StringVarP(&quality, "quality", "q", "flac", "preferred audio quality (mp3_128, mp3_320, flac, try)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	// Bind flags to viper
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("quality", rootCmd.PersistentFlags().Lookup("quality"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-music-dl"
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-music-dl")
	}

	// Read in environment variables that match
	viper.SetEnvPrefix("MUSIC_DL")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
