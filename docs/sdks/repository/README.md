# Repository
(*Repository*)

### Available Operations

* [DownloadDocumentByID](#downloaddocumentbyid) - Download document
* [GetDocumentByID](#getdocumentbyid) - Retrieve document
* [GetFolderByID](#getfolderbyid) - List folder content
* [GetSpaces](#getspaces) - List all spaces with access
* [UploadDocument](#uploaddocument) - Upload new document

## DownloadDocumentByID

Download document

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DownloadDocumentByIDRequest](../../pkg/models/operations/downloaddocumentbyidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DownloadDocumentByIDResponse](../../pkg/models/operations/downloaddocumentbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetDocumentByID

Retrieve document

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Repository.GetDocumentByID(ctx, operations.GetDocumentByIDRequest{
        DocumentID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetDocumentByIDRequest](../../pkg/models/operations/getdocumentbyidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetDocumentByIDResponse](../../pkg/models/operations/getdocumentbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetFolderByID

List folder content

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Repository.GetFolderByID(ctx, operations.GetFolderByIDRequest{
        FolderID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FolderListing != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetFolderByIDRequest](../../pkg/models/operations/getfolderbyidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetFolderByIDResponse](../../pkg/models/operations/getfolderbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSpaces

List all spaces with access

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"log"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Repository.GetSpaces(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SpaceList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSpacesResponse](../../pkg/models/operations/getspacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UploadDocument

Upload new document

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Repository.UploadDocument(ctx, &operations.UploadDocumentRequestBody{
        Attributes: operations.Attributes{
            FolderID: "string",
        },
        Document: operations.Document{
            Content: []byte("0xaAFa6d8bD5"),
            FileName: "electronic_hempstead_fresh.mpga",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UploadDocumentRequestBody](../../pkg/models/operations/uploaddocumentrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UploadDocumentResponse](../../pkg/models/operations/uploaddocumentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
