package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go01/util"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	ApplicationJSONValue = "application/json"
	ContentType          = "content-type"
)

type HttpFunc func(http.ResponseWriter, *http.Request) interface{}
type RequestHandler struct {
	Path         string
	Method       string
	ResponseBody bool
	Handler      HttpFunc
}
type Handler interface {
	Handles() []RequestHandler
	BasePath() string
}

func (requestHandler *RequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()
	defer func() {
		subTime := time.Now().Sub(startTime)
		log.Println(request.RemoteAddr, request.Method, request.RequestURI, " end request,spend time :", subTime.String())
		handlerError(writer, recover())
	}()
	checkRequest(request, requestHandler)
	log.Println(request.RemoteAddr, request.Method, request.RequestURI, " start request!")
	val := requestHandler.Handler(writer, request)
	if requestHandler.ResponseBody {
		writeJson(writer, util.SuccessResult(val))
	}
}
func RegisterHandlers() {
	var handlerMap = map[string]*RequestHandler{
		"/ping": {
			Path:    "/ping",
			Method:  http.MethodGet,
			Handler: index,
		},
	}
	putAllHandlers(handlerMap, new(GenHandler))
	for path, handler := range handlerMap {
		log.Println(path, " bound func", runtime.FuncForPC(reflect.ValueOf(handler.Handler).Pointer()).Name())
		http.Handle(path, handler)
	}
}

func handlerError(writer http.ResponseWriter, recoverError interface{}) {
	if recoverError == nil {
		return
	}
	errLog(recoverError)
	err, ok := recoverError.(error)
	var errResult *util.ApiResult
	if !ok {
		errResult = util.SystemErrorResult()
	} else {
		var serviceError *util.ServiceError
		if errors.As(err, &serviceError) {
			errResult = util.ServiceErrorResult(serviceError.Message, "")
		} else {
			errResult = util.SystemErrorResult()
		}

	}
	writeJson(writer, errResult)
}

func checkRequest(request *http.Request, handler *RequestHandler) {
	if handler.Method != "" && request.Method != handler.Method {
		panic(util.ServiceErrorResult(request.Method+" not support!", ""))
	}

}

func errLog(err interface{}) {
	log.Println(err)
	skip := 0
	_, curFile, _, ok := runtime.Caller(skip)
	if !ok {
		log.Println("get current file name fail!")
		return
	}
	for {
		skip++
		_, file, _, ok := runtime.Caller(skip)
		if !ok || file != curFile {
			break
		}
	}
	for {
		skip++
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		log.Println(file, ":", line)
	}
}
func writeJson(writer http.ResponseWriter, result *util.ApiResult) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set(ContentType, ApplicationJSONValue)
	var bytes []byte
	stop := false
	var err error
	for !stop {
		bytes, err = json.Marshal(result)
		if err == nil {
			break
		}
		result = util.SystemErrorResult()
		stop = true
	}
	_, _ = writer.Write(bytes)
}

func getFuncName(handler HttpFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
}

func putAllHandlers(baseHandlers map[string]*RequestHandler, handler Handler) {
	basePath := handler.BasePath()
	for i, v := range handler.Handles() {
		if v.Handler == nil {
			panic(reflect.ValueOf(handler).String() + " handlers[" + strconv.Itoa(i) + "]" + " is nil!")
		}
		path := basePath + v.Path
		if baseHandlers[path] != nil {
			panic(fmt.Sprint(path, " bound duplicate! ", getFuncName(baseHandlers[path].Handler), " ", getFuncName(v.Handler)))
		}
		if v.Method != "" {
			v.Method = strings.ToUpper(v.Method)
		}
		baseHandlers[path] = &v

	}
}

func index(writer http.ResponseWriter, _ *http.Request) interface{} {
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("welcome!"))
	return nil
}
