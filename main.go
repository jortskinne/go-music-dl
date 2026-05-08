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
	// Quality options: 128 (low), 192 (medium), 320 (high), 0 = lossless/flac where available
	rootCmd.PersistentFlags().IntP("quality", "q", 320, "Audio quality in kbps (128, 192, 320, 0 for lossless)")
	// Increased concurrency to 5 — my connection handles it fine and it speeds up playlist downloads
	rootCmd.PersistentFlags().IntP("concurrency", "c", 5, "Number of concurrent downloads")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
