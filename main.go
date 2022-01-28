package pd

import (
	"encoding/json"
	"os"
)

func get(key string, fallback []byte) []byte {
	value, _ := os.ReadFile(os.Getenv(key))
	if value == nil || len(value) == 0 {
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
	pdSteps := []byte(get("PIPEDREAM_STEPS", []byte("null")))
	if err := json.Unmarshal(pdSteps, &Steps); err != nil {
		panic(err)
	}
}
