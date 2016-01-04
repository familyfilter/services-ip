package rest

import (
     "encoding/json"
)

type ProcessError struct {
    Error string `json:"error"`
}

// When an error occurs processig JSON this method is called to format the error on the
// return JSON.  If this method fails it will panic becuase it assumed that an unrecoverable
// syntax error has been encounted,
func ProcessingError(err error) (string) {
    
    processError := ProcessError{Error: err.Error()}
    b, err := json.Marshal(&processError)
    if err != nil {
        panic(err)
    }
    return string(b)
}