package util

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
}

func (err *ServiceError) Error() string {
	return err.Message
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

func NewServiceError(message string) error {
	return &ServiceError{Message: message}
}
