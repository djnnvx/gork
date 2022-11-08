package gork

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GetCmdParser(opts *Options) *cobra.Command {
    var version = "0.0.2";
    var rootCmd = &cobra.Command{
        Use:   "gork",
        Version: version,
        DisableSuggestions : true,
        Short: "gork - a simple CLI to perform Google dorks",
        Long: `gork is a CLI to perform Google dorks in order to retrieve cool files :)~`,
        Run: func(cmd *cobra.Command, args []string) {

            if (opts.target == "") {
                println("[!] Please specify target with -t")
                return
            }

            println("[+] Running gork on " + opts.target)
            dorks := RunSearch(opts)

            fileOpts := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
            if (opts.appendResults) {
                fileOpts = os.O_CREATE | os.O_WRONLY | os.O_APPEND
            }

            fs, err := os.OpenFile(opts.outfile, fileOpts, 0644)
            if (err != nil) {
                fmt.Printf("[!] could not open file %s: %s", opts.outfile, err.Error())
                return
            }

            fs.WriteString("\t-== GORK RESULTS FOR " + opts.target + " ==-\n\n")
            for extensions, results := range dorks {
                fs.WriteString("\t--==== " + extensions + " ===-\n")
                for idx := range results {
                    fs.WriteString(results[idx].URL + " " + results[idx].Title + "\n")
                }
                fs.WriteString("\n")
            }

            println("[+] done.")
        },
    }

    defaultUserAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36";
    defaultExtensions := []string{"doc", "docx", "csv", "pdf", "txt", "log", "bak", "json", "xlsx"}

    rootCmd.PersistentFlags().StringVarP(&opts.target, "target", "t", "", "target site for your dorks")
    rootCmd.PersistentFlags().StringVarP(&opts.outfile, "outfile", "o", "./gork.txt", "directory storing dorks results")
    rootCmd.PersistentFlags().StringVarP(&opts.proxy, "proxy", "p", "", "proxy URL")
    rootCmd.PersistentFlags().StringVarP(&opts.userAgent, "user-agent", "u", defaultUserAgent, "Which user-agent gork should use")
    rootCmd.PersistentFlags().StringArrayVarP(&opts.extensions, "extensions", "e", defaultExtensions, "filetype extensions")
    rootCmd.PersistentFlags().BoolVarP(&opts.appendResults, "append-results", "a", false, "append dork results to out file")

    return rootCmd
}
