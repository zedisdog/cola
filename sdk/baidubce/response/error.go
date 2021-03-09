package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponseWithSuccess struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func ParseError(response *http.Response) (err error) {
	var (
		e       ErrorResponse
		errData []byte
	)
	defer response.Body.Close()
	errData, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(errData, &e)
	if err != nil {
		return
	}
	err = errors.New(fmt.Sprintf("%s error: %s error_description: %s", response.Status, e.Error, e.ErrorDescription))
	return
}
