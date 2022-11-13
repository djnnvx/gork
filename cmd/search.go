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
    var results map[string][]googlesearch.Result = make(map[string][]googlesearch.Result)

    /* preparing google search according to user-defined settings */
    dorkURL := getDorkUrl(opts.Target, opts.Extensions)
    searchOpts := googlesearch.SearchOptions{
        UserAgent: opts.UserAgent,
        ProxyAddr: opts.Proxy,
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
    for e := range(opts.Extensions) {
        ext := opts.Extensions[e]
        results[ext] = getResultsByFiletype(googleResults, ext)
    }

    /*
        these must be done separately, because they have nothing to do with looking for files,
        they're more like, looking for links...
    */
    results["dir listing"] = runDirListing(opts.Target, searchOpts)
    results["project setup files"] = runSetupFiles(opts.Target, searchOpts)
    results["open redirects"] = runOpenRedirects(opts.Target, searchOpts)

    return results
}
