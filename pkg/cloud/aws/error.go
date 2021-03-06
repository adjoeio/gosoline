package aws

import (
	"errors"
	"fmt"
	"github.com/applike/gosoline/pkg/exec"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

func init() {
	exec.AddRequestCancelCheck(IsAwsErrorCodeRequestCanceled)
}

type InvalidStatusError struct {
	Status int
}

func (e *InvalidStatusError) Error() string {
	return fmt.Sprintf("http status code: %d", e.Status)
}

func (e *InvalidStatusError) Is(err error) bool {
	_, ok := err.(*InvalidStatusError)

	return ok
}

func (e *InvalidStatusError) As(target interface{}) bool {
	err, ok := target.(*InvalidStatusError)

	if ok && err != nil {
		*err = *e
	}

	return ok
}

func IsInvalidStatusError(err error) bool {
	return errors.Is(err, &InvalidStatusError{})
}

func CheckInvalidStatusError(_ interface{}, err error) exec.ErrorType {
	if IsInvalidStatusError(err) {
		return exec.ErrorRetryable
	}

	return exec.ErrorUnknown
}

func IsAwsError(err error, awsCode string) bool {
	var awsErr awserr.Error

	if errors.As(err, &awsErr) {
		return awsErr.Code() == awsCode
	}

	return false
}

func IsAwsErrorCodeRequestCanceled(err error) bool {
	if IsAwsError(err, request.CanceledErrorCode) {
		return true
	}

	return false
}

func CheckConnectionError(_ interface{}, err error) exec.ErrorType {
	if IsConnectionError(err) {
		return exec.ErrorRetryable
	}

	return exec.ErrorUnknown
}

func IsConnectionError(err error) bool {
	var awsErr awserr.Error

	if errors.As(err, &awsErr) {
		err = awsErr.OrigErr()
	}

	return exec.IsConnectionError(awsErr)
}

func CheckErrorRetryable(_ interface{}, err error) exec.ErrorType {
	if request.IsErrorRetryable(err) {
		return exec.ErrorRetryable
	}

	return exec.ErrorUnknown
}

func CheckErrorThrottle(_ interface{}, err error) exec.ErrorType {
	if request.IsErrorThrottle(err) {
		return exec.ErrorRetryable
	}

	return exec.ErrorUnknown
}
