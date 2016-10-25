package ferror

import (
	"fmt"
	"net/http"
)

type FuzzerError struct {
	code     int
	msg      string
	bag      error
	resp     *http.Response
	respBody string
}

func (e FuzzerError) Error() string {
	str := "%v: %v \n\n" +
		"     bag: %v \n" +
		"     response: %v \n" +
		"          body-> %v\n"

	return fmt.Sprintf(str, e.code, e.msg, e.bag, e.resp, e.respBody)
}

func SetupErr() FuzzerError {
	return FuzzerError{
		msg: "Error: You havn't setup the required information, please refer to srv config.",
	}
}

func RequestErr(err error, c int) FuzzerError {
	return FuzzerError{
		code: c,
		msg:  "Error: Impossible to send request.",
		bag:  err,
	}
}

func BuildRequestErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Impossible to create request",
		bag: err,
	}
}

func FileErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Encounter a problem with file",
		bag: err,
	}
}

func PartErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Can't create part",
		bag: err,
	}
}

func NormalizeErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Impossible to normalize the string",
		bag: err,
	}
}

func TestErr(r *http.Response, b string) FuzzerError {
	return FuzzerError{
		msg:      "Error: Server doesn't respond well to test",
		resp:     r,
		respBody: b,
	}
}
