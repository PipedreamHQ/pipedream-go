package pd

import (
	"encoding/json"
	"os"
)

var (

	// Steps provides access to previous step data
	// Unmarshal should be preferred with an explicit type
	Steps map[string]interface{}

	// StepsEnv is the environment variable name for previous pipedream steps
	StepsEnv = "PIPEDREAM_STEPS"
	// ExportsEnv is the environment variable name for exporting data to a file
	ExportsEnv = "PIPEDREAM_EXPORTS"
)

// MustUnmarshal unmarshals the previous step data, panicking on error
func MustUnmarshal(in interface{}) {
	if err := Unmarshal(in); err != nil {
		panic(err)
	}
}

// Unmarshal unmarshals the previous step data, returning any errors encountered
func Unmarshal(in interface{}) error {
	pdSteps := getFromFile(StepsEnv, []byte("null"))
	if err := json.Unmarshal(pdSteps, in); err != nil {
		return err
	}
	return nil
}

// Export exports data for subsequent steps to use
func Export(name string, value interface{}) {
	export, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(os.Getenv(ExportsEnv), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(name + ":json=" + string(export) + "\n"); err != nil {
		panic(err)
	}
}

func getFromFile(key string, fallback []byte) []byte {
	value, _ := os.ReadFile(os.Getenv(key))
	if len(value) == 0 {
		return fallback
	}
	return value
}

func init() {
	MustUnmarshal(&Steps)
}
