package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (app *Application) checkPrimaries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	inputData, err := io.ReadAll(r.Body)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	inputNums := make([]string, 0)
	if err := json.Unmarshal(inputData, &inputNums); err != nil {
		app.serverError(w, err)
	}

	respBool := make([]string, 0, len(inputNums))
	for i := range inputNums {
		if num, err := strconv.Atoi(inputNums[i]); err != nil {
			app.clientError(w, http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Element on index %v is not a number", i+1)))
			return
		} else {
			if isPrime(num) {
				respBool = append(respBool, "true")
				continue
			}

			respBool = append(respBool, "false")
		}
	}

	respData, err := json.Marshal(respBool)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write(respData)
}
