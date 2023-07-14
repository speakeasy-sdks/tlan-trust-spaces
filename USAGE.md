<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Inbox.DeleteMessageByID(ctx, operations.DeleteMessageByIDRequest{
        MessageID: "corrupti",
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