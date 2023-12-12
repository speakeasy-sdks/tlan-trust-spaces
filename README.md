# tlan-trust-spaces

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tlan-trust-spaces
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Repository](docs/sdks/repository/README.md)

* [DownloadDocumentByID](docs/sdks/repository/README.md#downloaddocumentbyid) - Download document
* [GetDocumentByID](docs/sdks/repository/README.md#getdocumentbyid) - Retrieve document
* [GetFolderByID](docs/sdks/repository/README.md#getfolderbyid) - List folder content
* [GetSpaces](docs/sdks/repository/README.md#getspaces) - List all spaces with access
* [UploadDocument](docs/sdks/repository/README.md#uploaddocument) - Upload new document

### [Inbox](docs/sdks/inbox/README.md)

* [DeleteMessageByID](docs/sdks/inbox/README.md#deletemessagebyid) - Delete message
* [GetConversationByID](docs/sdks/inbox/README.md#getconversationbyid) - Get conversation content
* [GetMessageByID](docs/sdks/inbox/README.md#getmessagebyid) - Get message details
* [ListConversationsBySpaceID](docs/sdks/inbox/README.md#listconversationsbyspaceid) - List conversations of a space
* [ReplyToConversation](docs/sdks/inbox/README.md#replytoconversation) - Reply to conversation
* [StartConversation](docs/sdks/inbox/README.md#startconversation) - Start new conversation
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/operations"
	"tlan-trust-spaces/v2/pkg/models/sdkerrors"
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

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://repository.int.otc.fts.d-3.de/api/v1` | None |
| 1 | `https://trustspaces.d-velop/api/v1` | None |

#### Example

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
		tlantrustspaces.WithServerIndex(1),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		tlantrustspaces.WithServerURL("https://repository.int.otc.fts.d-3.de/api/v1"),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/operations"
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
