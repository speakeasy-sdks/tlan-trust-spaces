# tlan-trust-spaces

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tlan-trust-spaces
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Inbox](docs/sdks/inbox/README.md)

* [DeleteMessageByID](docs/sdks/inbox/README.md#deletemessagebyid) - Delete message
* [GetConversationByID](docs/sdks/inbox/README.md#getconversationbyid) - Get conversation content
* [GetMessageByID](docs/sdks/inbox/README.md#getmessagebyid) - Get message details
* [ListConversationsBySpaceID](docs/sdks/inbox/README.md#listconversationsbyspaceid) - List conversations of a space
* [ReplyToConversation](docs/sdks/inbox/README.md#replytoconversation) - Reply to conversation
* [StartConversation](docs/sdks/inbox/README.md#startconversation) - Start new conversation

### [Repository](docs/sdks/repository/README.md)

* [DownloadDocumentByID](docs/sdks/repository/README.md#downloaddocumentbyid) - Download document
* [GetDocumentByID](docs/sdks/repository/README.md#getdocumentbyid) - Retrieve document
* [GetFolderByID](docs/sdks/repository/README.md#getfolderbyid) - List folder content
* [GetSpaces](docs/sdks/repository/README.md#getspaces) - List all spaces with access
* [UploadDocument](docs/sdks/repository/README.md#uploaddocument) - Upload new document
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
