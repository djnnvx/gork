package gork

type Options struct {
    /* HTTP client settings */
    proxy           string
    userAgent       string

    /* HTTPS settings */
    useHTTPS        bool
    skipTLScheck    bool

    /* Out settings */
    outdir          string
}

