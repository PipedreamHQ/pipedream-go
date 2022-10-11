package pd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var (
	stepData = []byte(`{"step1": { "data": "foo" }}`)
	tmpDir   string
)

func TestMain(m *testing.M) {
	var err error
	tmpDir, err = os.MkdirTemp(os.TempDir(), "pipedream")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			fmt.Println("could not clean up temp test dir")
		}
	}()

	stepsFile := filepath.Join(tmpDir, "test-step-data.json")
	fi, err := os.Create(stepsFile)
	if err != nil {
		panic(err)
	}
	fi.Write(stepData)
	fi.Close()

	os.Setenv(StepsEnv, stepsFile)
	os.Setenv(ExportsEnv, filepath.Join(tmpDir, "test-exports-data"))
	MustUnmarshal(&Steps)

	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	tests := []struct {
		key           string
		fallback      []byte
		expectedValue []byte
	}{
		{StepsEnv, []byte("null"), []byte(`{"step1": { "data": "foo" }}`)},
		{"FOO", []byte("null"), []byte("null")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s/%s", tt.key, tt.fallback)
		t.Run(testname, func(t *testing.T) {
			pdSteps := getFromFile(tt.key, tt.fallback)
			if !bytes.Equal(pdSteps, tt.expectedValue) {
				t.Errorf("got %d, want %d", pdSteps, tt.expectedValue)
			}
		})
	}
}

func TestExport(t *testing.T) {
	testname := "$PIPEDREAM_EXPORTS file holds the correct data"
	t.Run(testname, func(t *testing.T) {
		value := "value"
		key := "key"
		Export(key, value)

		fileData, _ := os.ReadFile(os.Getenv(ExportsEnv))

		expectedValue := fmt.Sprintf("%s:json=\"%s\"\n", key, value)
		if string(fileData) != expectedValue {
			t.Errorf("got %v, want %v", string(fileData), expectedValue)
		}
	})
}

func TestSteps(t *testing.T) {
	testname := "Steps holds the value of $PIPEDREAM_STEPS"
	t.Run(testname, func(t *testing.T) {
		var envJSON map[string]interface{}
		envData := getFromFile(StepsEnv, []byte("null"))
		if err := json.Unmarshal(envData, &envJSON); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(Steps, envJSON) {
			t.Errorf("got %v, want %v", Steps, envJSON)
		}
	})
}
