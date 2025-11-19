package open

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// OpenURI opens a URI.
// Supported schemes are: file, http & https.
func OpenURI(uri string) (io.ReadCloser, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		resp, err := http.Get(uri)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("%q: bad status - %s", uri, resp.Status)
		}
		return resp.Body, nil
	case "file":
		file, err := os.Open(u.Path)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return nil, fmt.Errorf("unknown scheme: %s", u.Scheme)
}
