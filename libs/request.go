package libs

import (
	"net/http"
	"io"

	c "../conf"
)

func CreateRequest(URL string) (io.ReadCloser){
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", c.UA)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
	}

	return res.Body
}
