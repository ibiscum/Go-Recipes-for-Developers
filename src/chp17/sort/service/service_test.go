package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	if testing.Short() {
		t.Skip("Service")
	}
	mux := GetServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	rsp, err := http.Post(server.URL+"/sort/asc", "application/json", strings.NewReader("test"))
	if err != nil {
		t.Error(err)
		return
	}
	// Must return http error
	if rsp.StatusCode/100 == 2 {
		t.Errorf("Error was expected")
		return
	}

	data, err := json.Marshal([]time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	})
	if err != nil {
		t.Error(err)
		return
	}
	rsp, err = http.Post(server.URL+"/sort/asc", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", rsp.StatusCode)
		return
	}
	var output []time.Time
	if err := json.NewDecoder(rsp.Body).Decode(&output); err != nil {
		t.Error(err)
		return
	}
	for i := 1; i < len(output); i++ {
		if !output[i-1].Before(output[i]) {
			t.Errorf("Wrong order")
		}
	}

}

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	data, err := json.Marshal([]time.Time{
		time.Date(2023, 2, 1, 12, 8, 37, 0, time.Local),
		time.Date(2021, 5, 6, 9, 48, 11, 0, time.Local),
		time.Date(2022, 11, 13, 17, 13, 54, 0, time.Local),
		time.Date(2022, 6, 23, 22, 29, 28, 0, time.Local),
		time.Date(2023, 3, 17, 4, 5, 9, 0, time.Local),
	})
	if err != nil {
		t.Error(err)
		return
	}
	req, _ := http.NewRequest("POST", "localhost/sort/asc", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	HandleSort(w, req, true)
	if w.Result().StatusCode != 200 {
		t.Errorf("Expecting HTTP 200, got %d", w.Result().StatusCode)
		return
	}
	var output []time.Time
	if err := json.NewDecoder(w.Result().Body).Decode(&output); err != nil {
		t.Error(err)
		return
	}
	for i := 1; i < len(output); i++ {
		if !output[i-1].Before(output[i]) {
			t.Errorf("Wrong order")
		}
	}
}
