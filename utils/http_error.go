package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// ErrorType represent valid error types.
type ErrorType string

// Keys for ErrorType
const (
	ErrorTypeBadRequest       ErrorType = "BadRequest"
	ErrorTypeBinding          ErrorType = "Binding"
	ErrorTypeValidation       ErrorType = "Validation"
	ErrorTypeResourceNotFound ErrorType = "ResourceNotFound"
	ErrorTypeURLNotFound      ErrorType = "URLNotFound"
	ErrorTypeDatabase         ErrorType = "Database"
	ErrorTypeInternalServer   ErrorType = "InternalServer"
	ErrorTypeBadGateway       ErrorType = "BadGateway"
	ErrorTypeUnauthorized     ErrorType = "Unauthorized"
	ErrorTypeForbidden        ErrorType = "Forbidden"
	ErrorTypeConflict         ErrorType = "Conflict"
)

var codeTypeMapping = map[int]ErrorType{
	http.StatusBadRequest:          ErrorTypeBadRequest,
	http.StatusNotFound:            ErrorTypeURLNotFound,
	http.StatusInternalServerError: ErrorTypeInternalServer,
	http.StatusBadGateway:          ErrorTypeBadGateway,
	http.StatusUnauthorized:        ErrorTypeUnauthorized,
	http.StatusForbidden:           ErrorTypeForbidden,
	http.StatusConflict:            ErrorTypeConflict,
}

// HasHTTPStatus Error Interface which contains an HTTP Status and a specific error type
type HasHTTPStatus interface {
	HTTPStatusCode() int
	ErrorType() ErrorType
}

// CommonError Parent struct for errors which are used in microservices
// Each child error contains an error type, http status code and the previous 'thrown' error
// It implements the HasHTTPStatus Interface
type CommonError struct {
	Err   error
	Code  int
	XType ErrorType
}

// errorResponse is used to handle an error response from IPC calls.
type errorResponse struct {
	Message string
	XType   ErrorType `json:"type"`
}

func (e *CommonError) Error() string {
	return e.Err.Error()
}
func (e *CommonError) HTTPStatusCode() int {
	return e.Code
}
func (e *CommonError) ErrorType() ErrorType {
	return e.XType
}

// ErrorBadRequest Error for 400 responses
type ErrorBadRequest struct {
	CommonError
}

// ErrorUnauthorized Error for 401 responses
type ErrorUnauthorized struct {
	CommonError
}

// ErrorBinding Error for 400 responses when echo.Bind fails
type ErrorBinding struct {
	CommonError
}

// ErrorValidation Error for 401 responses when validation fails
type ErrorValidation struct {
	CommonError
}

// ErrorDatabase Error for 500 responses when database functions fails
type ErrorDatabase struct {
	CommonError
}

// ErrorResourceNotFound Error for 404 responses when a resource is not found
type ErrorResourceNotFound struct {
	CommonError
}

// ErrorURLNotFound Error for 404 responses when a URL is not available
type ErrorURLNotFound struct {
	CommonError
}

// ErrorInternalServer Error for 500 responses when server runs into an error.
// Also used for errors, which are not clearly specified
type ErrorInternalServer struct {
	CommonError
}

// ErrorBadGateway Error for 502 responses when request cannot be served.
type ErrorBadGateway struct {
	CommonError
}

// ErrorForbidden Error for 403 responses when access to the requested resource is forbidden.
type ErrorForbidden struct {
	CommonError
}

// ErrorConflict Error for 409 responses when the request conflicts with the current state of the server.
type ErrorConflict struct {
	CommonError
}

// Errorf creates an error using Sprintf to build the error message.
func Errorf(xtype ErrorType, message string, args ...interface{}) error {
	return Error(fmt.Sprintf(message, args...), xtype)
}

// Error Factory method for creating CommonErrors
func Error(i interface{}, xtype ErrorType) error {
	var err error
	if str, ok := i.(string); ok {
		err = errors.New(str)
	} else if commonError, ok := i.(HasHTTPStatus); ok {
		return commonError.(error)
	} else if err, ok = i.(error); !ok {
		panic(fmt.Sprintf("I don't know how to handle that type: %v", i))
	}

	switch xtype {
	case ErrorTypeBadRequest:
		return &ErrorBadRequest{CommonError{err, http.StatusBadRequest, xtype}}
	case ErrorTypeBinding:
		return &ErrorBinding{CommonError{err, http.StatusBadRequest, xtype}}
	case ErrorTypeValidation:
		return &ErrorValidation{CommonError{err, http.StatusBadRequest, xtype}}
	case ErrorTypeDatabase:
		return &ErrorDatabase{CommonError{err, http.StatusInternalServerError, xtype}}
	case ErrorTypeResourceNotFound:
		return &ErrorResourceNotFound{CommonError{err, http.StatusNotFound, xtype}}
	case ErrorTypeURLNotFound:
		return &ErrorURLNotFound{CommonError{err, http.StatusNotFound, xtype}}
	case ErrorTypeUnauthorized:
		return &ErrorUnauthorized{CommonError{err, http.StatusUnauthorized, xtype}}
	case ErrorTypeBadGateway:
		return &ErrorBadGateway{CommonError{err, http.StatusBadGateway, xtype}}
	case ErrorTypeForbidden:
		return &ErrorForbidden{CommonError{err, http.StatusForbidden, xtype}}
	case ErrorTypeConflict:
		return &ErrorConflict{CommonError{err, http.StatusConflict, xtype}}
	default:
		return &ErrorInternalServer{CommonError{err, http.StatusInternalServerError, xtype}}
	}
}

// ErrorTypeByStatusCode returns the ErrorType by status code.
func ErrorTypeByStatusCode(code int) ErrorType {
	if xtype, ok := codeTypeMapping[code]; ok {
		return xtype
	}
	return ErrorTypeInternalServer
}

// ValidationErrorStructure - Representation of a validation error
type ValidationErrorStructure struct {
	Class     string `json:"class,omitempty"`
	Field     string `json:"field"`
	Validator string `json:"validator"`
	Message   string `json:"message,omitempty"`
}

// HTTPError - Basic Implementation of ClientError
type HTTPError struct {
	Code             string                     `json:"code,omitempty"`
	Type             ErrorType                  `json:"type"`
	Message          string                     `json:"message"`
	MessageID        string                     `json:"messageId,omitempty"`
	ValidationErrors []ValidationErrorStructure `json:"validationErrors,omitempty"`
	Status           int                        `json:"-"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

// HTTPErrorHandler Error handler to use in echo's middleware
// Example:
// 			router := echo.New()
// 			router.HTTPErrorHandler = commons.HTTPErrorHandler
func HTTPErrorHandler(err error, c echo.Context) {
	// Because echo.NotFoundHandler returns echo.HTTPError and to avoid duplicate of code at router initialization
	// the error given by echo, we overwrite it by an error of our own type
	if (reflect.TypeOf(err) == reflect.TypeOf(&echo.HTTPError{})) && (err.(*echo.HTTPError).Code == http.StatusNotFound) {
		err = Error(err, ErrorTypeURLNotFound)
	}

	// Because echo.MethodNotAllowedHandler returns echo.HTTPError and to avoid duplicate of code at router initialization
	// the error given by echo, we overwrite it by an error of our own type
	if (reflect.TypeOf(err) == reflect.TypeOf(&echo.HTTPError{})) && (err.(*echo.HTTPError).Code == http.StatusUnauthorized) {
		err = Error(err, ErrorTypeUnauthorized)
	}

	requestID := c.Response().Header().Get(echo.HeaderXRequestID)

	switch err.(type) {
	case *ErrorValidation:
		err = c.JSON(err.(*ErrorValidation).Code, validationErrorToHTTPError(err.(*ErrorValidation), requestID))
	case HasHTTPStatus:
		httpError := newHTTPError(err.(HasHTTPStatus), nil, requestID)
		err = c.JSON(err.(HasHTTPStatus).HTTPStatusCode(), httpError)
	default:
		httpError := newHTTPError(Error(err, ErrorTypeInternalServer).(HasHTTPStatus), nil, requestID)
		err = c.JSON(http.StatusInternalServerError, httpError)
	}
}

// HandleErrorResponse parses an error response and returns a CommonError with the parsed information.
func HandleErrorResponse(resp *http.Response) error {
	errorRespJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error("could not read authentication response", ErrorTypeInternalServer)
	}

	var errorResp errorResponse

	err = json.Unmarshal(errorRespJSON, &errorResp)
	if err != nil {
		return Error("unknown error format", ErrorTypeBadGateway)
	}

	return Error(errorResp.Message, errorResp.XType)
}

// newHTTPError - Create a new HTTPError instance
func newHTTPError(err HasHTTPStatus, valErrors []ValidationErrorStructure, requestID string) error {

	errMsg := err.(error).Error()
	if Environment() == ENV_PROD {
		switch err.(type) {
		case *ErrorDatabase:
			errMsg = http.StatusText(err.HTTPStatusCode())
		}
	}

	return &HTTPError{
		Code:             "",
		Type:             err.ErrorType(),
		Message:          errMsg,
		MessageID:        requestID,
		ValidationErrors: valErrors,
		Status:           err.HTTPStatusCode(),
	}
}

// validationErrorToHTTPError - Generate a HTTPError for a bad request including validation errors
func validationErrorToHTTPError(err *ErrorValidation, requestID string) error {

	// extract validation error information from validation structs
	valErrors := []ValidationErrorStructure{}
	if verr, ok := err.Err.(validator.ValidationErrors); ok {
		for i := 0; i < len(verr); i++ {
			structNames := strings.Split(verr[i].StructNamespace(), ".")
			structNS := structNames[0]
			valError := ValidationErrorStructure{structNS, verr[i].StructField(), verr[i].ActualTag(), verr[i].Translate(nil)}
			valErrors = append(valErrors, valError)
		}
	}

	// and paste it to http error struct
	httpError := newHTTPError(err, valErrors, requestID)

	return httpError
}
