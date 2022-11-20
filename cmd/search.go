package gork

import (
	"context"
	"fmt"
	"sync"

	runner "github.com/bogdzn/gork/runner"
)


func runDirListing(target string, searchOpts runner.SearchOptions) []runner.Result {

    ctx := context.Background()

    term := fmt.Sprintf("site:%s intitle:index.of", target)
    results, err := runner.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform Dir Listing dork: %s\n", err.Error())
        return []runner.Result{}
    }

    return results
}

func runSetupFiles(target string, searchOpts runner.SearchOptions) []runner.Result {
    ctx := context.Background()

    term := fmt.Sprintf("site:%s inurl:readme | inurl:license | inurl:install | inurl:setup | inurl:config", target)
    results, err := runner.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform SetupFiles dork: %s\n", err.Error())
        return []runner.Result{}
    }

    return results
}

func runOpenRedirects(target string, searchOpts runner.SearchOptions) []runner.Result {
    ctx := context.Background()

    term := fmt.Sprintf("site:%s  inurl:redir | inurl:url | inurl:redirect | inurl:return | inurl:src=http | inurl:r=http", target)
    results, err := runner.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] Could not perform OpenRedirects dork: %s\n", err.Error())
        return []runner.Result{}
    }

    return results
}

func RunSearch(opts *Options) map[string][]runner.Result {
    var wg sync.WaitGroup
    searchOpts := runner.SearchOptions{
        UserAgent: opts.UserAgent,
        ProxyAddr: opts.Proxy,
        FollowLinks: true,
    }

    ctx := context.Background()
    var results map[string][]runner.Result = make(map[string][]runner.Result)

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
            r, err := runner.Search(ctx, term, searchOpts)

            if (err != nil) {
                fmt.Printf("[!] could not perform dork %s: %s\n", term, err.Error())
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
