package gork

import (
	"github.com/spf13/cobra"
)

func DefaultUserAgent() string {
    return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
}

func DefaultExtensions() []string {
    return []string{
        "doc",
        "docx",
        "csv",
        "pdf",
        "txt",
        "log",
        "bak",
        "json",
        "xlsx",
        "xml",
        "conf",
        "cnf",
        "reg",
        "inf",
        "rdp",
        "cfg",
        "ora",
        "ini",
        "sql",
        "mdb",
        "dbf",
        "bkf",
        "bkp",
        "old",
        "backup",
        "rtf",
        "odt",
        "ppt",
        "sxt",
        "pptx",
    }
}

func GetCmdParser(opts *Options) *cobra.Command {
    var version = "0.0.4";
    var rootCmd = &cobra.Command{
        Use:   "gork",
        Version: version,
        DisableSuggestions : true,
        Short: "gork - a simple CLI to perform Google dorks",
        Long: `gork is a CLI to perform Google dorks in order to retrieve cool files :)~`,
        Run: func(cmd *cobra.Command, args []string) {
            Run(opts)
        },
    }

    rootCmd.PersistentFlags().StringVarP(&opts.Target, "target", "t", "", "target site for your dorks")
    rootCmd.PersistentFlags().StringVarP(&opts.Outfile, "outfile", "o", "./gork.txt", "directory storing dorks results")
    rootCmd.PersistentFlags().StringVarP(&opts.Proxy, "proxy", "p", "", "proxy URL")
    rootCmd.PersistentFlags().StringVarP(&opts.UserAgent, "user-agent", "u", DefaultUserAgent(), "Which user-agent gork should use")
    rootCmd.PersistentFlags().StringArrayVarP(&opts.Extensions, "extensions", "e", DefaultExtensions(), "filetype extensions")
    rootCmd.PersistentFlags().BoolVarP(&opts.AppendResults, "append-results", "a", false, "append dork results to out file")

    return rootCmd
}
