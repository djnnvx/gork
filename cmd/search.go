package gork

import (
	"context"
	"fmt"
	"sync"

	"github.com/rocketlaunchr/google-search"
)

func RunSearch(opts *Options) map[string][]googlesearch.Result {
    var wg sync.WaitGroup
    searchOpts := googlesearch.SearchOptions{
        UserAgent: opts.userAgent,
        ProxyAddr: opts.proxy,
    }

    ctx := context.Background()
    var results map[string][]googlesearch.Result = make(map[string][]googlesearch.Result)

    for ext := range opts.extensions {
        extension := opts.extensions[ext]

        /*
            TODO:
            have all the extensions in only one request, and then filter
            by filetype the request results

        */

        wg.Add(1)
        go func() {
            defer wg.Done()

            term := fmt.Sprintf("site:%s ext:%s", opts.target, extension)
            r, err := googlesearch.Search(ctx, term, searchOpts)

            if (err != nil) {
                fmt.Printf("[!] could not perform dork %s: %s", term, err.Error())
                return
            }

            results[extension] = r
        }()
    }

    wg.Wait()
    return results
}
