package http

import (
	"bytes"
	"encoding/json"
	"github.com/zedisdog/cola/errx"
	"io"
	"net/http"
	urlpkg "net/url"
)

//PostJSON post json
//	url is the target to post.
//
//	data is to be posted.it can be string, []byte and struct, also nil.
func PostJSON(url string, data interface{}) ([]byte, error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, errx.Wrap(err, "parse url error")
	}

	var body io.ReadCloser
	switch data.(type) {
	case []byte:
		body = io.NopCloser(bytes.NewBuffer(data.([]byte)))
	case string:
		body = io.NopCloser(bytes.NewBufferString(data.(string)))
	default:
		if data == nil {
			break
		}
		tmp, err := json.Marshal(data)
		if err != nil {
			return nil, errx.Wrap(err, "covert interface{} to json bytes error")
		}
		body = io.NopCloser(bytes.NewBuffer(tmp))
	}

	req := http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
		},
		Method: http.MethodPost,
		Body:   body,
		URL:    u,
	}

	return Request(&req)
}

//GetJSON get json
func GetJSON(url string) ([]byte, error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		return nil, errx.Wrap(err, "parse url error")
	}
	request := http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
		},
	}
	return Request(&request)
}

func Request(request *http.Request) (response []byte, err error) {
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		content, _ := io.ReadAll(resp.Body)
		return nil, errx.NewHttpError(resp.StatusCode, string(content))
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		err = errx.Wrap(err, "read body err")
	}
	return
}
