package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

const (
	SuccessCode      = 1000
	ServiceErrorCode = 2000
	SystemErrorCode  = 2000
)

type ApiResult struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}
type ServiceError struct {
	Message string
	Err     error
}

func (err *ServiceError) Error() string {
	return fmt.Sprint(err.Message, err.Err)
}

func BaseResult(status int, message string, description string, value interface{}) *ApiResult {
	return &ApiResult{
		Status:      status,
		Message:     message,
		Description: description,
		Value:       value,
	}
}
func SuccessResult(value interface{}) (apiResult *ApiResult) {
	apiResult = &ApiResult{Status: SuccessCode, Value: value}
	return BaseResult(SuccessCode, "", "", value)
}

func ServiceErrorResult(msg string, des string) (apiResult *ApiResult) {
	return BaseResult(ServiceErrorCode, msg, des, nil)
}

func SystemErrorResult() (apiResult *ApiResult) {
	return BaseResult(SystemErrorCode, "系统异常", "", nil)
}

func AssertError(err error, message string) {
	assertExpError(err == nil, err, message)
}
func assertExpError(exp bool, err error, message string) {
	if !exp {
		panic(&ServiceError{Message: message, Err: err})
	}
}
func ValidateError(exp bool, message string) {
	assertExpError(exp, nil, message)
}

func Body2Json(request *http.Request, v interface{}) {
	err := json.Unmarshal(Body2Bytes(request), v)
	AssertError(err, "json转换失败")
}

func Body2Bytes(request *http.Request) []byte {
	bytes, err := ioutil.ReadAll(request.Body)
	AssertError(err, "读取body失败")
	return bytes
}

type RegexMux struct {
	mu sync.RWMutex
	m  map[string]*regexp.Regexp
}

var regexMux RegexMux = RegexMux{m: map[string]*regexp.Regexp{}}

func compileRegex(pattern string) *regexp.Regexp {
	r := regexMux.m[pattern]
	if r != nil {
		return r
	}
	regexMux.mu.Lock()
	defer regexMux.mu.Unlock()
	if regexMux.m[pattern] == nil {
		r, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		regexMux.m[pattern] = r
	}
	return regexMux.m[pattern]
}

//
func RegexMatch(pattern string, src string) bool {
	return compileRegex(pattern).Match([]byte(src))
}
