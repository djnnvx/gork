package gork

import (
    "os"
    "fmt"
)

func Run(opts *Options) {
    if (opts.Target == "") {
        println("[!] please specify target with -t")
        return
    }

    println("[+] running gork on " + opts.Target)
    dorks := RunSearch(opts)

    fileOpts := os.O_CREATE | os.O_TRUNC | os.O_WRONLY
    if (opts.AppendResults) {
        fileOpts = os.O_CREATE | os.O_WRONLY | os.O_APPEND
    }

    fs, err := os.OpenFile(opts.Outfile, fileOpts, 0644)
    if (err != nil) {
        fmt.Printf("[!] could not open file %s: %s", opts.Outfile, err.Error())
        return
    }

    fs.WriteString("\t-== GORK RESULTS FOR " + opts.Target + " ==-\n\n")
    for extensions, results := range dorks {
        if len(results) == 0 {
            continue
        }

        fmt.Printf("\tfound result(s) for %s\n", extensions)
        fs.WriteString("\t    --==== " + extensions + " ===-\n")
        for idx := range results {
            fs.WriteString(results[idx].URL + " " + results[idx].Title + "\n")
        }
        fs.WriteString("\n")
    }

    println("[+] scan completed")
}
