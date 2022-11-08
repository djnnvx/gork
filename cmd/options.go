package gork

type Options struct {
    /* HTTP client settings */
    proxy           string
    userAgent       string

    /* Out settings */
    outfile         string

    /* Target settings */
    target          string
    extensions      []string
}

