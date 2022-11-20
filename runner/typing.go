package gork

// Result represents a single result from Google Search.
type Result struct {
	// URL of result.
	URL string `json:"url"`

	// Title of result.
	Title string `json:"title"`
}


// SearchOptions modifies how the Search function behaves.
type SearchOptions struct {
	// UserAgent sets the UserAgent of the http request.
	// Default: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	UserAgent string

	// ProxyAddr sets a proxy address to avoid IP blocking.
	ProxyAddr string

	// follow links
	FollowLinks bool
}

