package baidubce

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zedisdog/cola/sdk/baidubce/auth"
	"github.com/zedisdog/cola/sdk/baidubce/response"
	"github.com/zedisdog/cola/sdk/baidubce/store"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// todo: 5,423

const Host = "aip.baidubce.com"

var s *store.VerifyToken

func getStore() *store.VerifyToken {
	if s == nil {
		s = new(store.VerifyToken)
	}
	return s
}

func New(clientId string, clientSecret string) *Client {
	a := auth.NewAuth(clientId, clientSecret, Host)
	return &Client{
		auth: a,
	}
}

// Client 百度ai平台sdk客户端
type Client struct {
	auth         *auth.Auth
	verifyPlanId int
}

// SetVerifyPlanId 设置方案id，不要忘记
func (c *Client) SetVerifyPlanId(id int) {
	c.verifyPlanId = id
}

// VerifyToken 获取verify_token
func (c Client) VerifyToken() (token string, err error) {
	if c.verifyPlanId == 0 {
		err = errors.New("invalid plan id")
	}
	u, err := c.genUrl("rpc/2.0/brain/solution/faceprint/verifyToken/generate")
	if err != nil {
		return
	}
	res, err := c.Post(u, map[string]interface{}{
		"plan_id": c.verifyPlanId,
	}, false)
	if err != nil {
		return
	}
	var result VerifyResponse
	err = c.Read(res, &result)
	token = result.Result.VerifyToken
	return
}

// VerifyTokenUsingStore 获取verify_token,并且存到store中
//  key 保存的verify_token的唯一键名
func (c Client) VerifyTokenUsingStore(key string) (token string, err error) {
	s := getStore()
	token = s.Pull(key)
	if token == "" {
		token, err = c.VerifyToken()
		if err != nil {
			return
		}
		s.Put(key, token)
	}
	return
}

type VerifyResponse struct {
	Success bool `json:"success"`
	Result  struct {
		VerifyToken string `json:"verify_token"`
	} `json:"result"`
	LogId int64 `json:"log_id"`
}

// GenVerifyUrl 生成人脸实名认证H5跳转链接
//  successUrl 成功回调链接
//  failedUrl 失败回调链接
//  verifyToken verify_token
func (c Client) GenVerifyUrl(successUrl string, failedUrl string, verifyToken string) (u string, err error) {
	tmp := url.URL{
		Scheme: "https",
		Host:   "brain.baidu.com",
		Path:   "face/print/",
		RawQuery: fmt.Sprintf(
			"token=%s&successUrl=%s&failedUrl=%s",
			verifyToken,
			url.QueryEscape(successUrl),
			url.QueryEscape(failedUrl),
		),
	}
	return tmp.String(), nil
}

// GenVerifyUrlUsingStore 从store中获取verify_token，并生成人脸实名认证H5跳转链接
//  successUrl 成功回调链接
//  failedUrl 失败回调链接
//  key 保存的verify_token的唯一键名
func (c Client) GenVerifyUrlUsingStore(successUrl string, failedUrl string, key string) (u string, err error) {
	verifyToken, err := c.VerifyTokenUsingStore(key)
	if err != nil {
		return
	}
	return c.GenVerifyUrl(successUrl, failedUrl, verifyToken)
}

func (c Client) GetVerifyResult(verifyToken string) (res VerifyResultResponse, err error) {
	u, err := c.genUrl("rpc/2.0/brain/solution/faceprint/result/detail")
	if err != nil {
		return
	}
	r, err := c.Post(u, map[string]interface{}{
		"verify_token": verifyToken,
	}, false)
	if err != nil {
		return
	}
	err = c.Read(r, &res)
	return
}

func (c Client) GetVerifyResultUsingStore(key string) (res VerifyResultResponse, err error) {
	verifyToken, err := c.VerifyTokenUsingStore(key)
	if err != nil {
		return
	}
	return c.GetVerifyResult(verifyToken)
}

type VerifyResultResponse struct {
	Success bool `json:"success"`
	Result  struct {
		VerifyResult struct {
			Score         float64 `json:"score"`
			LivenessScore float64 `json:"liveness_score"`
			Spoofing      float64 `json:"spoofing"`
		} `json:"verify_result"`
		IdcardOcrResult struct {
			Birthday       string `json:"birthday"`
			IssueAuthority string `json:"issue_authority"`
			Address        string `json:"address"`
			Gender         string `json:"gender"`
			Nation         string `json:"nation"`
			ExpireTime     string `json:"expire_time"`
			Name           string `json:"name"`
			IssueTime      string `json:"issue_time"`
			IdCardNumber   string `json:"id_card_number"`
		} `json:"idcard_ocr_result"`
		IdcardImages struct {
			FrontBase64 string `json:"front_base_64"`
			BackBase64  string `json:"back_base_64"`
		} `json:"idcard_images"`
		IdcardConfirm struct {
			IdcardNumber string `json:"idcard_number"`
			Name         string `json:"name"`
		} `json:"idcard_confirm"`
	} `json:"result"`
	LogId int64 `json:"log_id"`
}

type HandwritingConfig struct {
	Image                []byte
	RecognizeGranularity string
	Probability          string
	DetectDirection      string
}

func (c Client) Handwriting(config HandwritingConfig) (result HandwritingResponse, err error) {
	if config.RecognizeGranularity == "" {
		config.RecognizeGranularity = "big"
	}
	if config.Probability == "" {
		config.Probability = "false"
	}
	if config.DetectDirection == "" {
		config.DetectDirection = "false"
	}
	if config.Image == nil || len(config.Image) < 1 {
		err = errors.New("image is required")
		return
	}
	u, err := c.genUrl("rest/2.0/ocr/v1/handwriting")
	if err != nil {
		return
	}
	image := base64.StdEncoding.EncodeToString(config.Image)
	r, err := c.Post(u, map[string]interface{}{
		"image":                 image,
		"recognize_granularity": config.RecognizeGranularity,
		"probability":           config.Probability,
		"detect_direction":      config.DetectDirection,
	}, true)
	if err != nil {
		return
	}
	err = c.Read(r, &result)
	return
}

type HandwritingResponse struct {
	LogId       int64 `json:"log_id"`
	WordsResult []struct {
		Location struct {
			Left   int `json:"left"`
			Top    int `json:"top"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"location"`
		Words string `json:"words"`
		Chars []struct {
			Location struct {
				Left   int `json:"left"`
				Top    int `json:"top"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"location"`
			Char string `json:"char"`
		} `json:"chars"`
	} `json:"words_result"`
	Direction      int `json:"direction"`
	WordsResultNum int `json:"words_result_num"`
}

func (c Client) Post(u string, data map[string]interface{}, formData bool) (res *http.Response, err error) {
	d, err := json.Marshal(data)
	if err != nil {
		return
	}
	var req *http.Request
	if !formData {
		req, err = http.NewRequest("POST", u, bytes.NewReader(d))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		body := new(bytes.Buffer)
		f := multipart.NewWriter(body)
		for key, value := range data {
			err = f.WriteField(key, value.(string))
			if err != nil {
				return
			}
		}
		req, err = http.NewRequest("POST", u, body)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", f.FormDataContentType())
	}
	client := http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode >= 400 {
		return nil, response.ParseError(res)
	}
	return
}

func (c Client) Read(res *http.Response, result interface{}) (err error) {
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	println(string(content))
	if err != nil {
		return
	}
	var errorWithSuccess response.ErrorResponseWithSuccess
	err = json.Unmarshal(content, &errorWithSuccess)
	if err != nil {
		return
	}
	if errorWithSuccess.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("error_code: %d, error_msg: %s", errorWithSuccess.ErrorCode, errorWithSuccess.ErrorMsg))
	}
	err = json.Unmarshal(content, result)
	return
}

func (c Client) genUrl(path string) (u string, err error) {
	tmp := url.URL{
		Scheme: "https",
		Host:   Host,
		Path:   path,
	}
	accessToken, err := c.auth.GetAccessToken()
	if err != nil {
		return
	}
	query := tmp.Query()
	query.Set("access_token", accessToken)
	tmp.RawQuery = query.Encode()
	u = tmp.String()
	return
}
