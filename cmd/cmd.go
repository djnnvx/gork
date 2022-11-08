package gork

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gork",
	DisableSuggestions : false,
}
