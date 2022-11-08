package gork

import (
    "github.com/spf13/cobra"
)

func GetCmdParser(opts *Options) *cobra.Command {
    var version = "0.0.1";
    var rootCmd = &cobra.Command{
        Use:   "gork",
        Version: version,
        DisableSuggestions : true,
        Short: "gork - a simple CLI to perform Google dorks",
        Long: `gork is a Golang CLI module to perform Google dorks :)~`,
        Run: func(cmd *cobra.Command, args []string) {
            println("[+] Running gork on " + opts.target)


            /* TODO: send option to chromedriver interface and execute */

        },
    }

    defaultUserAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36";
    defaultExtensions := []string{"doc", "docx", "csv", "pdf", "txt", "log", "bak"}

    rootCmd.PersistentFlags().StringVarP(&opts.target, "target", "t", "", "target site for your dorks")
    rootCmd.PersistentFlags().StringVarP(&opts.outfile, "outfile", "o", "./grok.txt", "directory storing dorks results")
    rootCmd.PersistentFlags().StringVarP(&opts.proxy, "proxy", "p", "", "proxy URL")
    rootCmd.PersistentFlags().StringVarP(&opts.userAgent, "user-agent", "u", defaultUserAgent, "Which user-agent gork should use")
    rootCmd.PersistentFlags().BoolVar(&opts.useHTTPS, "use-https", true, "toggle HTTPS usage")
    rootCmd.PersistentFlags().BoolVar(&opts.skipTLScheck, "skip-tls", false, "skip TLS certificate check")
    rootCmd.PersistentFlags().StringArrayVarP(&opts.extensions, "extensions", "e", defaultExtensions, "filetype extensions")

    return rootCmd
}
