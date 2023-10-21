# Inbox
(*Inbox*)

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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Inbox.GetConversationByID(ctx, operations.GetConversationByIDRequest{
        ConversationID: "string",
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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Inbox.GetMessageByID(ctx, operations.GetMessageByIDRequest{
        MessageID: "string",
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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Inbox.ListConversationsBySpaceID(ctx, operations.ListConversationsBySpaceIDRequest{
        SpaceID: "string",
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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Inbox.ReplyToConversation(ctx, operations.ReplyToConversationRequest{
        Reply: &shared.Reply{
            Documents: []shared.MessageDocument{
                shared.MessageDocument{
                    ID: "<ID>",
                    Name: "string",
                },
            },
        },
        ConversationID: "string",
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
	tlantrustspaces "tlan-trust-spaces"
	"tlan-trust-spaces/pkg/models/shared"
	"tlan-trust-spaces/pkg/models/operations"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Inbox.StartConversation(ctx, operations.StartConversationRequest{
        ConversationStart: &shared.ConversationStart{
            Documents: []shared.ConversationStartDocuments{
                shared.ConversationStartDocuments{},
            },
            Recipients: []string{
                "string",
            },
            Subject: "string",
        },
        SpaceID: "string",
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

