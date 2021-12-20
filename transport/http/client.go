package http

import (
	"bytes"
	"encoding/json"
	"github.com/zedisdog/cola/errx"
	"io"
	"net/http"
	urlpkg "net/url"
)

func buildBody(data interface{}) (body io.ReadCloser, err error) {
	switch data.(type) {
	case []byte:
		body = io.NopCloser(bytes.NewBuffer(data.([]byte)))
	case string:
		body = io.NopCloser(bytes.NewBufferString(data.(string)))
	default:
		if data == nil {
			return
		}
		tmp, err := json.Marshal(data)
		if err != nil {
			return nil, errx.Wrap(err, "covert interface{} to json bytes error")
		}
		body = io.NopCloser(bytes.NewBuffer(tmp))
	}
	return
}

func buildRequest(method string, url string, data interface{}) (request *http.Request, err error) {
	u, err := urlpkg.Parse(url)
	if err != nil {
		err = errx.Wrap(err, "parse url error")
		return
	}

	body, err := buildBody(data)
	if err != nil {
		err = errx.Wrap(err, "build body error")
		return
	}

	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
		},
		Method: method,
		Body:   body,
		URL:    u,
	}

	return
}

func PutJSON(url string, data interface{}) (response []byte, err *errx.HttpError) {
	request, e := buildRequest(http.MethodPut, url, data)
	if e != nil {
		err = errx.NewHttpError(0, e.Error())
		return
	}
	return Request(request)
}

//PostJSON post json
//	url is the target to post.
//
//	data is to be posted.it can be string, []byte and struct, also nil.
func PostJSON(url string, data interface{}) (response []byte, err *errx.HttpError) {
	request, e := buildRequest(http.MethodPost, url, data)
	if e != nil {
		err = errx.NewHttpError(0, e.Error())
		return
	}
	return Request(request)
}

//GetJSON get json
func GetJSON(url string) (response []byte, err *errx.HttpError) {
	request, e := buildRequest(http.MethodGet, url, nil)
	if e != nil {
		err = errx.NewHttpError(0, e.Error())
		return
	}
	return Request(request)
}

func Request(request *http.Request) (response []byte, err *errx.HttpError) {
	resp, e := http.DefaultClient.Do(request)
	if e != nil {
		err = errx.NewHttpError(0, err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		content, _ := io.ReadAll(resp.Body)
		err = errx.NewHttpError(resp.StatusCode, string(content))
		return
	}

	response, e = io.ReadAll(resp.Body)
	if e != nil {
		err = errx.NewHttpError(0, err.Error())
		return
	}
	return
}
