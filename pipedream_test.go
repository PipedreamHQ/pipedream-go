package pd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var STEPS = "PIPEDREAM_STEPS"
var EXPORTS = "PIPEDREAM_EXPORTS"

func teardown() {
	err := os.Remove(os.Getenv(EXPORTS))
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(os.Getenv(EXPORTS), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(""); err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	var tests = []struct {
		key           string
		fallback      []byte
		expectedValue []byte
	}{
		{STEPS, []byte("null"), []byte(`{"step1": { "data": "foo" }}`)},
		{"FOO", []byte("null"), []byte("null")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s/%s", tt.key, tt.fallback)
		t.Run(testname, func(t *testing.T) {
			pdSteps := []byte(get(tt.key, tt.fallback))
			if bytes.Compare(pdSteps, tt.expectedValue) != 0 {
				t.Errorf("got %d, want %d", pdSteps, tt.expectedValue)
			}
		})
	}
}

func TestExport(t *testing.T) {
	testname := fmt.Sprintf("$PIPEDREAM_EXPORTS file holds the correct data")
	t.Run(testname, func(t *testing.T) {
		value := "value"
		key := "key"
		Export(key, value)

		fileData, _ := os.ReadFile(os.Getenv(EXPORTS))

		expectedValue := fmt.Sprintf("%s:json=\"%s\"\n", key, value)
		if string(fileData) != expectedValue {
			t.Errorf("got %v, want %v", string(fileData), expectedValue)
		}

		teardown()
	})
}

func TestSteps(t *testing.T) {
	testname := fmt.Sprintf("Steps holds the value of $PIPEDREAM_STEPS")
	t.Run(testname, func(t *testing.T) {
		var envJSON map[string]interface{}
		envData := []byte(get(STEPS, []byte("null")))
		if err := json.Unmarshal(envData, &envJSON); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(Steps, envJSON) {
			t.Errorf("got %v, want %v", Steps, envData)
		}
	})
}
