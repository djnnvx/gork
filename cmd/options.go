package gork

type Options struct {
    /* HTTP client settings */
    proxy           string
    userAgent       string

    /* Out settings */
    outfile         string
    appendResults   bool

    /* Target settings */
    target          string
    extensions      []string
}

