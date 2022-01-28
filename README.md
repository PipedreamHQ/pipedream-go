# `pipedream-go`

```go
package main

import (
	"fmt"

	pd "github.com/PipedreamHQ/pipedream-go"
)

func main() {
	// Access previous step data using pd.Steps
	fmt.Println(pd.Steps)

	// Export data using pd.Export
	data := make(map[string]interface{})
	data["name"] = "Luke"
	pd.Export("data", data)
}
```

## Tests

Simulate the env vars present on the Pipedream execution environment when running `go test`

```bash
PIPEDREAM_STEPS=./test-step-data.json PIPEDREAM_EXPORTS=./test-exports-data go test
```
