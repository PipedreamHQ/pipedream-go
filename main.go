package pd

import (
	"encoding/json"
	"os"
)  

func getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var Steps map[string]interface{}

func Export(name string, value interface{}) {
	export, _ := json.Marshal(value)
	env := getenv("PIPEDREAM_EXPORTS", "null")
	os.Setenv("PIPEDREAM_EXPORTS", env + name + ":json=" + string(export) + "\n")
}

func init() {
	pdSteps := []byte(getenv("PIPEDREAM_STEPS", "null"))
	if err := json.Unmarshal(pdSteps, &Steps); err != nil {
		panic(err)
	}
}