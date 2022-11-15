package gork

import (
	"fmt"
    "strings"
    "context"

    "golang.org/x/time/rate"

	"github.com/rocketlaunchr/google-search"
)

func getFileExtensionSearchUrl(target string, extensions []string) string {
    var result string = fmt.Sprintf("site:%s ", target)
    nbExtensions := len(extensions)

    for e := range(extensions) {

        /* if there are no extensions left after this one, don't add OR clause */
        sprintfFormat := "ext:%s"
        if (nbExtensions != e + 1) {
            sprintfFormat = "ext:%s | "
        }

        /* append newly created dork */
        ext := fmt.Sprintf(sprintfFormat, extensions[e])
        result += ext
    }

    return result
}

func filterByFiletype(searchResults []googlesearch.Result, extension string) []googlesearch.Result {
    result := []googlesearch.Result{};

    /* look if the URL finishes with filetype. */
    for s := range(searchResults) {
        if (strings.HasSuffix(searchResults[s].URL, extension)) {
            result = append(result, searchResults[s])
        }
    }
    return result
}

func runDorkWrapper(term string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    ctx := context.Background()

    results, err := googlesearch.Search(ctx, term, searchOpts)
    if err != nil {
        fmt.Printf("[!] could not perform dork: %s\n", err.Error())
        fmt.Printf("\t[URL]: %s\n", term)
        return []googlesearch.Result{}
    }

    return results
}

func runDirListing(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    term := fmt.Sprintf("site:%s intitle:index.of", target)
    return runDorkWrapper(term, searchOpts)
}

func runSetupFiles(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    term := fmt.Sprintf("site:%s inurl:readme | inurl:license | inurl:install | inurl:setup | inurl:config", target)
    return runDorkWrapper(term, searchOpts)
}

func runOpenRedirects(target string, searchOpts googlesearch.SearchOptions) []googlesearch.Result {
    term := fmt.Sprintf("site:%s  inurl:redir | inurl:url | inurl:redirect | inurl:return | inurl:src=http | inurl:r=http", target)
    return runDorkWrapper(term, searchOpts)
}

func RunSearch(opts *Options) map[string][]googlesearch.Result {

    /* results is map where keys are extension-names (or dork-type) & values are googlesearch.Result objects */
    var results map[string][]googlesearch.Result = make(map[string][]googlesearch.Result)

    /* preparing google search according to user-defined settings */
    searchOpts := googlesearch.SearchOptions{
        UserAgent: opts.UserAgent,
        ProxyAddr: opts.Proxy,
        // FollowLinks: true,
        /* Waiting for https://github.com/rocketlaunchr/google-search/pull/18 to be merged */
    }

    /*
        Setting rate-limiter:
        https://pkg.go.dev/golang.org/x/time/rate?utm_source=godoc#NewLimiter
    */
    var RateLimit = rate.NewLimiter(5, 0)
    _ = RateLimit

    /* running google dork to look for interesting files */
    filesTerm := getFileExtensionSearchUrl(opts.Target, opts.Extensions)
    extResults := runDorkWrapper(filesTerm, searchOpts)

    /*
        build a map where the results are mapped to the filetype (or extension, if you will)
        it would probably be faster to loop only once through the results and append to the correct value,
        so this is subject to change, if it doesn't increase code readability too much
    */
    for e := range(opts.Extensions) {
        ext := opts.Extensions[e]
        results[ext] = filterByFiletype(extResults, ext)
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
