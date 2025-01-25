package safe_error

import (
	"fmt"
	"strings"
)

var API_KEY string

func Return(error_msg string, e error) error {
	error := fmt.Sprintf(error_msg, e)
	error = strings.ReplaceAll(error, API_KEY, "[API_KEY HIDDEN]")
	return fmt.Errorf("%v", error)
}
