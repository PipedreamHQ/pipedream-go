package pd

import (
	"encoding/json"
	"os"
)  

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

var Steps map[string]interface{}

func init() {
	pdSteps := []byte(getenv("PIPEDREAM_STEPS", "null"))
	if err := json.Unmarshal(pdSteps, &Steps); err != nil {
        panic(err)
    }
}