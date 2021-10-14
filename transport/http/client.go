package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//PostJSON post json
//	url: url to post data
//	data: the json string as []byte
func PostJSON(url string, data []byte) ([]byte, error) {
	body := bytes.NewBuffer(data)
	response, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : url=%v , statusCode=%v", url, response.StatusCode)
	}
	v, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return v, nil
}

//GetJSON get json
func GetJSON(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : url=%v , statusCode=%v", url, response.StatusCode)
	}
	v, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return v, nil
}
