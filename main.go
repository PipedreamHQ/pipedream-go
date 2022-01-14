package pd

import (
 "fmt"
 "os"
)  

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

Steps := getenv("PIPEDREAM_STEPS", "null")