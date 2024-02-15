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
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Inbox.DeleteMessageByID(ctx, operations.DeleteMessageByIDRequest{
        MessageID: "<value>",
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
| `request`                                                                                      | [operations.DeleteMessageByIDRequest](../../pkg/models/operations/deletemessagebyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteMessageByIDResponse](../../pkg/models/operations/deletemessagebyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetConversationByID

Get conversation content

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
    res, err := s.Inbox.GetConversationByID(ctx, operations.GetConversationByIDRequest{
        ConversationID: "<value>",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetConversationByIDRequest](../../pkg/models/operations/getconversationbyidrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetConversationByIDResponse](../../pkg/models/operations/getconversationbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMessageByID

Get message details

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
    res, err := s.Inbox.GetMessageByID(ctx, operations.GetMessageByIDRequest{
        MessageID: "<value>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetMessageByIDRequest](../../pkg/models/operations/getmessagebyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetMessageByIDResponse](../../pkg/models/operations/getmessagebyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListConversationsBySpaceID

List conversations of a space

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
    res, err := s.Inbox.ListConversationsBySpaceID(ctx, operations.ListConversationsBySpaceIDRequest{
        SpaceID: "<value>",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListConversationsBySpaceIDRequest](../../pkg/models/operations/listconversationsbyspaceidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.ListConversationsBySpaceIDResponse](../../pkg/models/operations/listconversationsbyspaceidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReplyToConversation

Reply to conversation

### Example Usage

```go
package main

import(
	"tlan-trust-spaces/v3/pkg/models/shared"
	tlantrustspaces "tlan-trust-spaces/v3"
	"context"
	"tlan-trust-spaces/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := tlantrustspaces.New(
        tlantrustspaces.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Inbox.ReplyToConversation(ctx, operations.ReplyToConversationRequest{
        ConversationID: "<value>",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ReplyToConversationRequest](../../pkg/models/operations/replytoconversationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ReplyToConversationResponse](../../pkg/models/operations/replytoconversationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## StartConversation

Start new conversation

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
    res, err := s.Inbox.StartConversation(ctx, operations.StartConversationRequest{
        SpaceID: "<value>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.StartConversationRequest](../../pkg/models/operations/startconversationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.StartConversationResponse](../../pkg/models/operations/startconversationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
