package gork

type Options struct {
    /* HTTP client settings */
    Proxy           string
    UserAgent       string

    /* Out settings */
    Outfile         string
    AppendResults   bool

    /* Target settings */
    Target          string
    Extensions      []string
    Exclusions      []string
}

/*
    This function should be used by programs wanting to integrate gork as a module.
    That way, only relevant values should be updated by the developer.
*/
func DefaultSearchOptions() Options {
    result := Options{
        Proxy: "",
        UserAgent: DefaultUserAgent(),
        Extensions: DefaultFileExtensions(),
        Exclusions: DefaultExclusions(),
        Outfile: DefaultOutfile(),
        AppendResults: false,
        Target: "",
    }

    return result
}
