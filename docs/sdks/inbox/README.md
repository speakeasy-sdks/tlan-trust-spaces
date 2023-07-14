# Inbox

### Available Operations

* [DeleteMessageByID](#deletemessagebyid) - Delete message
* [GetConversationByID](#getconversationbyid) - Get conversation content
* [GetMessageByID](#getmessagebyid) - Get message details
* [ListConversationsBySpaceID](#listconversationsbyspaceid) - List conversations of a space
* [ReplyToConversation](#replytoconversation) - Reply to conversation
* [StartConversation](#startconversation) - Start new conversation

## DeleteMessageByID

Delete message

### Example Usage

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
        MessageID: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteMessageByIDRequest](../../models/operations/deletemessagebyidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteMessageByIDResponse](../../models/operations/deletemessagebyidresponse.md), error**


## GetConversationByID

Get conversation content

### Example Usage

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
    res, err := s.Inbox.GetConversationByID(ctx, operations.GetConversationByIDRequest{
        ConversationID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetConversationByIDRequest](../../models/operations/getconversationbyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetConversationByIDResponse](../../models/operations/getconversationbyidresponse.md), error**


## GetMessageByID

Get message details

### Example Usage

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
    res, err := s.Inbox.GetMessageByID(ctx, operations.GetMessageByIDRequest{
        MessageID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Message != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetMessageByIDRequest](../../models/operations/getmessagebyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetMessageByIDResponse](../../models/operations/getmessagebyidresponse.md), error**


## ListConversationsBySpaceID

List conversations of a space

### Example Usage

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
    res, err := s.Inbox.ListConversationsBySpaceID(ctx, operations.ListConversationsBySpaceIDRequest{
        SpaceID: "unde",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Inbox != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListConversationsBySpaceIDRequest](../../models/operations/listconversationsbyspaceidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.ListConversationsBySpaceIDResponse](../../models/operations/listconversationsbyspaceidresponse.md), error**


## ReplyToConversation

Reply to conversation

### Example Usage

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
    res, err := s.Inbox.ReplyToConversation(ctx, operations.ReplyToConversationRequest{
        Reply: &shared.Reply{
            Documents: []shared.MessageDocument{
                shared.MessageDocument{
                    ID: "8d69a674-e0f4-467c-8879-6ed151a05dfc",
                    Name: "Teri Strosin",
                },
                shared.MessageDocument{
                    ID: "cc78ca1b-a928-4fc8-9674-2cb739205929",
                    Name: "Faye Howe",
                },
                shared.MessageDocument{
                    ID: "a7596eb1-0faa-4a23-92c5-955907aff1a3",
                    Name: "Carlos Ziemann",
                },
                shared.MessageDocument{
                    ID: "46773925-1aa5-42c3-b5ad-019da1ffe78f",
                    Name: "Geneva Klein Jr.",
                },
            },
            Message: tlantrustspaces.String("reprehenderit"),
        },
        ConversationID: "ut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ReplyToConversationRequest](../../models/operations/replytoconversationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ReplyToConversationResponse](../../models/operations/replytoconversationresponse.md), error**


## StartConversation

Start new conversation

### Example Usage

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
    res, err := s.Inbox.StartConversation(ctx, operations.StartConversationRequest{
        ConversationStart: &shared.ConversationStart{
            Documents: []shared.ConversationStartDocuments{
                shared.ConversationStartDocuments{
                    ID: tlantrustspaces.String("15471b5e-6e13-4b99-9488-e1e91e450ad2"),
                },
                shared.ConversationStartDocuments{
                    ID: tlantrustspaces.String("abd44269-802d-4502-a94b-b4f63c969e9a"),
                },
                shared.ConversationStartDocuments{
                    ID: tlantrustspaces.String("3efa77df-b14c-4d66-ae39-5efb9ba88f3a"),
                },
                shared.ConversationStartDocuments{
                    ID: tlantrustspaces.String("66997074-ba44-469b-ae21-41959890afa5"),
                },
            },
            Message: tlantrustspaces.String("eum"),
            Recipients: []string{
                "necessitatibus",
            },
            Subject: "odit",
        },
        SpaceID: "nemo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationStartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.StartConversationRequest](../../models/operations/startconversationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.StartConversationResponse](../../models/operations/startconversationresponse.md), error**

