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



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
