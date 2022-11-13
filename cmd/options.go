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
}

