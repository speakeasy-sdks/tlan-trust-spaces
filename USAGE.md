<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/operations"
	"tlan-trust-spaces/pkg/models/shared"
)

func main() {
	s := tlantrustspaces.New(
		tlantrustspaces.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Inbox.DeleteMessageByID(ctx, operations.DeleteMessageByIDRequest{
		MessageID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->