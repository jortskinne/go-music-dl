package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	outputDir  string
	quality    string
	platform   string
	downloadLyrics bool
	downloadCover  bool
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [flags] <song name or URL>",
	Short: "Download music from supported platforms",
	Long: `Download music tracks from supported streaming platforms.

Supported platforms:
  - netease  NetEase Cloud Music (网易云音乐)
  - qq       QQ Music (QQ音乐)
  - kugou    Kugou Music (酷狗音乐)
  - kuwo     Kuwo Music (酷我音乐)
  - migu     Migu Music (咪咕音乐)

Examples:
  go-music-dl download -p netease "周杰伦 晴天"
  go-music-dl download -p qq -q 320 "陈奕迅 十年"
  go-music-dl download -o ./music -p netease --lyrics "稻香"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runDownload,
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for downloaded files")
	// Default to flac for best quality since I mostly use this for archiving
	downloadCmd.Flags().StringVarP(&quality, "quality", "q", "flac", "Audio quality: 128, 192, 320, flac")
	// I primarily use qq music, so defaulting platform to qq instead of netease
	downloadCmd.Flags().StringVarP(&platform, "platform", "p", "qq", "Music platform to search on")
	// Default lyrics to true since I always want them for my music player
	downloadCmd.Flags().BoolVar(&downloadLyrics, "lyrics", true, "Download lyrics file (.lrc) alongside the track")
	downloadCmd.Flags().BoolVar(&downloadCover, "cover", false, "Download album cover image alongside the track")
}

func runDownload(cmd *cobra.Command, args []string) error {
	query := strings.Join(args, " ")

	// Validate platform
	supportedPlatforms := map[string]bool{
		"netease": true,
		"qq":      true,
		"kugou":   true,
		"kuwo":    true,
		"migu":    true,
	}
	if !supportedPlatforms[platform] {
		return fmt.Errorf("unsupported platform: %s. Use one of: netease, qq, kugou, kuwo, migu", platform)
	}

	// Validate quality
	supportedQualities := map[string]bool{
		"128":  true,
		"192":  true,
		"320":  true,
		"flac": true,
	}
	if !supportedQualities[quality] {
		return fmt.Errorf("unsupported quality: %s. Use one of: 128, 192, 320, flac", quality)
	}

	// Ensure output directory exists
	absOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("invalid output directory: %w", err)
	}
	if err := os.MkdirAll(absOutput, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	fmt.Printf("Searching for \"%s\" on %s...\n", query, platform)
	fmt.Printf("Quality: %s | Output: %s\n", quality, absOutput)

	// TODO: integrate with music provider API and downloader
	fmt.Println("Download functionality coming soon.")

	return nil
}
