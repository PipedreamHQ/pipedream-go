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
	
	// Alternatively, unmarshal into a type
	var steps Steps
	pd.MustUnmarshal(&steps)
	fmt.Println(steps.Trigger.Event)

	// Export data using pd.Export
	data := make(map[string]interface{})
	data["name"] = "Luke"
	pd.Export("data", data)
}

type Steps struct {
	Trigger struct {
		Event struct {
			// ...
		} `json:"event"`
	} `json:"trigger"`
}
```

## Tests

```bash
go test ./...
```
