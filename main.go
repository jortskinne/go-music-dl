package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "go-music-dl",
	Short: "A music downloader written in Go",
	Long: `go-music-dl is a command-line tool for downloading music
from various streaming platforms including NetEase Cloud Music,
QQ Music, Kugou, Kuwo, and more.`,
	SilenceUsage: true,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-music-dl %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Global flags
	// Changed default output dir to ~/Music for convenience
	rootCmd.PersistentFlags().StringP("output", "o", "~/Music", "Output directory for downloaded files")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	// Default to 128 kbps to save disk space for personal use
	rootCmd.PersistentFlags().IntP("quality", "q", 128, "Audio quality in kbps (128, 192, 320)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
