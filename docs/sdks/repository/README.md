# Repository
(*.Repository*)

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
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/shared"
	"tlan-trust-spaces/v2/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DownloadDocumentByIDRequest](../../models/operations/downloaddocumentbyidrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.DownloadDocumentByIDResponse](../../models/operations/downloaddocumentbyidresponse.md), error**


## GetDocumentByID

Retrieve document

### Example Usage

```go
package main

import(
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/shared"
	"tlan-trust-spaces/v2/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDocumentByIDRequest](../../models/operations/getdocumentbyidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetDocumentByIDResponse](../../models/operations/getdocumentbyidresponse.md), error**


## GetFolderByID

List folder content

### Example Usage

```go
package main

import(
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/shared"
	"tlan-trust-spaces/v2/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetFolderByIDRequest](../../models/operations/getfolderbyidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetFolderByIDResponse](../../models/operations/getfolderbyidresponse.md), error**


## GetSpaces

List all spaces with access

### Example Usage

```go
package main

import(
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/shared"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
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

**[*operations.GetSpacesResponse](../../models/operations/getspacesresponse.md), error**


## UploadDocument

Upload new document

### Example Usage

```go
package main

import(
	"context"
	"log"
	tlantrustspaces "tlan-trust-spaces/v2"
	"tlan-trust-spaces/v2/pkg/models/shared"
	"tlan-trust-spaces/v2/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UploadDocumentRequestBody](../../models/operations/uploaddocumentrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UploadDocumentResponse](../../models/operations/uploaddocumentresponse.md), error**

