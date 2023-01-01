package gork

import (
	"github.com/spf13/cobra"
)

func DefaultOutfile() string {
    return "./gork.txt"
}

func DefaultUserAgent() string {
    return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
}

func DefaultExclusions() []string {
    return []string{
        "html",
    }
}

func DefaultFileExtensions() []string {
    return []string{
        "asp",
        "aspx",
        "backup",
        "bak",
        "bkf",
        "bkp",
        "cfg",
        "cgi",
        "cnf",
        "conf",
        "csv",
        "dbf",
        "doc",
        "docx",
        "fla",
        "inf",
        "ini",
        "json",
        "jsp",
        "jspx",
        "log",
        "mdb",
        "odt",
        "old",
        "ora",
        "pdf",
        "php",
        "ppt",
        "pptx",
        "rdp",
        "reg",
        "rtf",
        "sql",
        "sxt",
        "txt",
        "xlsx",
        "xml",
    }
}

func GetCmdParser(opts *Options) *cobra.Command {
    var version = "0.0.5";
    var rootCmd = &cobra.Command{
        Use:   "gork",
        Version: version,
        DisableSuggestions : true,
        Short: "gork - a simple CLI to perform Google dorks",
        Long: `gork is a CLI to perform Google dorks on a target domain :)~ (Example: ./gork -t nmap.org)`,
        Run: func(cmd *cobra.Command, args []string) {
            Run(opts)
        },
    }

    rootCmd.PersistentFlags().StringVarP(&opts.Target, "target", "t", "", "target site for your dorks")
    rootCmd.PersistentFlags().StringVarP(&opts.Outfile, "outfile", "o", DefaultOutfile(), "directory storing dorks results")
    rootCmd.PersistentFlags().StringVarP(&opts.Proxy, "proxy", "p", "", "proxy URL")
    rootCmd.PersistentFlags().StringVarP(&opts.UserAgent, "user-agent", "u", DefaultUserAgent(), "Which user-agent gork should use")
    rootCmd.PersistentFlags().StringArrayVarP(&opts.Extensions, "extensions", "e", DefaultFileExtensions(), "filetype extensions")
    rootCmd.PersistentFlags().StringArrayVarP(&opts.Exclusions, "exclude", "x", DefaultExclusions(), "exclude specific filetype")
    rootCmd.PersistentFlags().BoolVarP(&opts.AppendResults, "append-results", "a", false, "append dork results to out file")
    return rootCmd
}
