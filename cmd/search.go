package gork

import (
	"context"
	"fmt"
    "strings"

	"github.com/rocketlaunchr/google-search"
)

/*
    TODO:

    this is not the fastest way to do this at all, but it will do for now
    - using calls to Sprintf is costly, perhaps should use a StringBuilder instead
    - instead of filtering results from extensions to extensions, we should get the
        current file's extension & loop only once through the searchResults array

    - important issue: we need to crawl all of the results, because a single page of results
    is not enough
*/

func getDorkUrl(target string, extensions []string) string {
    var result string = fmt.Sprintf("site:%s ", target)
    nbExtensions := len(extensions)

    for e := range(extensions) {
        sprintfFormat := "ext:%s"
        if (nbExtensions != e + 1) {
            sprintfFormat = "ext:%s OR "
        }

        ext := fmt.Sprintf(sprintfFormat, extensions[e])
        result += ext
    }
    return result
}

func getResultsByFiletype(searchResults []googlesearch.Result, extension string) []googlesearch.Result {
    result := []googlesearch.Result{};

    for s := range(searchResults) {
        if (strings.HasSuffix(searchResults[s].URL, extension)) {
            result = append(result, searchResults[s])
        }
    }
    return result
}

func RunSearch(opts *Options) map[string][]googlesearch.Result {
    var results map[string][]googlesearch.Result = make(map[string][]googlesearch.Result)

    /* preparing google search according to user-defined settings */
    dorkURL := getDorkUrl(opts.target, opts.extensions)
    searchOpts := googlesearch.SearchOptions{
        UserAgent: opts.userAgent,
        ProxyAddr: opts.proxy,
    }

    /* running the actual google-search */
    ctx := context.Background()
    googleResults, err := googlesearch.Search(ctx, dorkURL, searchOpts)

    if (err != nil) {
        fmt.Printf("[!] could not perform dork %s: %s", dorkURL, err.Error())
        return results
    }

    /*
        build a map where the results are mapped to the filetype (or extension, if you will)
        it would probably be faster to loop only once through the results and append to the correct value,
        so this is subject to change, if it doesn't increase code readability too much
    */
    for e := range(opts.extensions) {
        ext := opts.extensions[e]
        results[ext] = getResultsByFiletype(googleResults, ext)
    }

    return results
}
