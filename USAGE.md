<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/operations"
	"tlan-trust-spaces/v2/pkg/models/shared"
)

func main() {
	s := tlantrustspaces.New(
		tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Repository.DownloadDocumentByID(ctx, operations.DownloadDocumentByIDRequest{
		DocumentID: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Stream != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->