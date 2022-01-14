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

func Export(name, value string) {
	var export string
	val := []byte(value)
	if err := json.Unmarshal(val, &export); err != nil {
        panic(err)
    }
	os.Setenv("PIPEDREAM_EXPORTS", name + ":json=" + export + "\n")
}

func init() {
	pdSteps := []byte(getenv("PIPEDREAM_STEPS", "null"))
	if err := json.Unmarshal(pdSteps, &Steps); err != nil {
        panic(err)
    }
}