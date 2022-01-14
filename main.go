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
	f, err := os.OpenFile(os.Getenv("PIPEDREAM_EXPORTS"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(name + ":json=" + string(export) + "\n"); err != nil {
		panic(err)
	}
}

func init() {
	pdSteps := []byte(getenv("PIPEDREAM_STEPS", "null"))
	if err := json.Unmarshal(pdSteps, &Steps); err != nil {
		panic(err)
	}
}
