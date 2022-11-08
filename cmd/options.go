package gork

type Options struct {
    /* HTTP client settings */
    proxy           string
    userAgent       string

    /* HTTPS settings */
    useHTTPS        bool
    skipTLScheck    bool

    /* Out settings */
    outfile         string

    /* Target settings */
    target          string
    extensions      []string
}

