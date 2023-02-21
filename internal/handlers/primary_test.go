package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		haveNum int
		wantRes bool
	}{
		{
			name:    "1",
			haveNum: 1,
			wantRes: false,
		},
		{
			name:    "2",
			haveNum: 2,
			wantRes: true,
		},
		{
			name:    "3",
			haveNum: 3,
			wantRes: true,
		},
		{
			name:    "4",
			haveNum: 4,
			wantRes: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isPrime(test.haveNum); got != test.wantRes {
				t.Errorf("Error have num: %v want result: %v", test.haveNum, test.wantRes)
			}
		})
	}
}

func testServer(addr string, mutex *sync.Mutex) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := NewApplication(errorLog, addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	log.Println("Starting Hosting on ", addr)
	mutex.Unlock()
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func TestHandler(t *testing.T) {
	addr := "localhost:8080"
	mutex := sync.Mutex{}

	mutex.Lock()
	go testServer(addr, &mutex)
	mutex.Lock()

	tests := []struct {
		name            string
		haveRequestBody []string
		wantResponse    []string
	}{
		{
			name:            "Test requests",
			haveRequestBody: []string{"1", "2", "3", "4", "5", "6"},
			wantResponse:    []string{"false", "true", "true", "false", "true", "false"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonData, err := json.Marshal(test.haveRequestBody)
			if err != nil {
				t.Error(err)
			}

			bodyReader := bytes.NewReader(jsonData)
			response, err := http.Post("http://"+addr+"/check", "text/json", bodyReader)
			if err != nil {
				t.Error(err)
			}

			bodyData, err := io.ReadAll(response.Body)
			if err != nil {
				t.Error(err)
			}

			responseArray := make([]string, 0)
			if err := json.Unmarshal(bodyData, &responseArray); err != nil {
				t.Error(err)
			}

			if strings.Join(responseArray, " ") != strings.Join(test.wantResponse, " ") {
				t.Errorf("Got response: \n%v\n Want response: \n%v", responseArray, test.wantResponse)
			}
		})
	}
}
