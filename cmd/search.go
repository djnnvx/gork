package gork

import (
	"context"
	"fmt"
	"sync"

	"github.com/rocketlaunchr/google-search"
)


func runDirListing(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {

    ctx := context.Background()

    term := fmt.Sprintf("site:%s intitle:index.of", target)
    results, err := googlesearch.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform Dir Listing dork: %s", err.Error())
        return []googlesearch.Result{}
    }

    return results
}

func runSetupFiles(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    ctx := context.Background()

    term := fmt.Sprintf("site:%s inurl:readme | inurl:license | inurl:install | inurl:setup | inurl:config", target)
    results, err := googlesearch.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform Dir Listing dork: %s", err.Error())
        return []googlesearch.Result{}
    }

    return results
}

func runOpenRedirects(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    ctx := context.Background()

    term := fmt.Sprintf("site:%s  inurl:redir | inurl:url | inurl:redirect | inurl:return | inurl:src=http | inurl:r=http", target)
    results, err := googlesearch.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform Dir Listing dork: %s", err.Error())
        return []googlesearch.Result{}
    }

    return results
}

func RunSearch(opts *Options) map[string][]googlesearch.Result {
    var wg sync.WaitGroup
    searchOpts := googlesearch.SearchOptions{
        UserAgent: opts.UserAgent,
        ProxyAddr: opts.Proxy,
    }

    ctx := context.Background()
    var results map[string][]googlesearch.Result = make(map[string][]googlesearch.Result)

    for ext := range opts.Extensions {
        extension := opts.Extensions[ext]

        /*
            TODO:
            have all the extensions in only one request, and then filter
            by filetype the request results

        */

        wg.Add(1)
        go func() {
            defer wg.Done()

            term := fmt.Sprintf("site:%s ext:%s", opts.Target, extension)
            r, err := googlesearch.Search(ctx, term, searchOpts)

            if (err != nil) {
                fmt.Printf("[!] could not perform dork %s: %s", term, err.Error())
                return
            }

            results[extension] = r
        }()
    }
    wg.Wait()

    results["dir listing"] = runDirListing(opts.Target, searchOpts)
    results["project setup files"] = runSetupFiles(opts.Target, searchOpts)
    results["open redirects"] = runOpenRedirects(opts.Target, searchOpts)

    return results
}
