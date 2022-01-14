package pd

import (
 "os"
)  

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

var Steps string

func init() {
	Steps = getenv("PIPEDREAM_STEPS", "null")
}