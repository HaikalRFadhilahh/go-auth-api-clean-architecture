package pkg

import (
	"net/http"
)

type HandleFuncWithError func(http.ResponseWriter, *http.Request) error

func ConvertToHttpHandleFunc(f HandleFuncWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			HttpErrorResponse(w, err)
			return
		}
	}
}
