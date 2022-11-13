package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"test/pkg/customerror"
	errorCustomStatus "test/pkg/error"
)

type restError struct {
	Error string `json:"error"`
}

// Success Serve Data With Success Status
func Success(w http.ResponseWriter, data interface{}) {
	responses, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		Error(w, customerror.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responses)
	return
}

// Error Serve Data With Error Status
func Error(w http.ResponseWriter, err error) {
	var result = restError{}
	var status int
	customErrors, _ := err.(*errorCustomStatus.Error)

	if customErrors == nil {
		result = restError{
			Error: customerror.ErrInternalServerError.Detail,
		}
	} else {
		status, err = strconv.Atoi(customErrors.Status)
		if err != nil {
			log.Println(err)
			Error(w, customerror.ErrInternalServerError)
			return
		}
		result = restError{
			Error: customErrors.Detail,
		}
	}

	responses, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		Error(w, customerror.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responses)
}
