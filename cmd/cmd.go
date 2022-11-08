package gork

import (
    "github.com/spf13/cobra"
)

func GetCmdParser(opts *Options) *cobra.Command {
    var version = "0.0.1";
    var rootCmd = &cobra.Command{
        Use:   "gork",
        Version: version,
        DisableSuggestions : false,
        Short: "gork - a simple CLI to perform Google dorks",
        Long: `gork is a Golang CLI module to perform Google dorks using google's chromedriver`,
        Args: cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            target := args[0];
            println("[+] Running gork on " + target)

            /* TODO: send option to chromedriver interface and execute */

        },
    }

    var defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36";

    rootCmd.PersistentFlags().StringVarP(&opts.outdir, "outdir", "o", "./grok_out", "directory storing dorks results")
    rootCmd.PersistentFlags().StringVarP(&opts.proxy, "proxy", "p", "", "proxy URL")
    rootCmd.PersistentFlags().StringVarP(&opts.userAgent, "user-agent", "u", defaultUserAgent, "Which user-agent gork should use")
    rootCmd.PersistentFlags().BoolVar(&opts.useHTTPS, "use-https", true, "toggle HTTPS usage")
    rootCmd.PersistentFlags().BoolVar(&opts.skipTLScheck, "skip-tls", false, "skip TLS certificate check")

    return rootCmd
}
