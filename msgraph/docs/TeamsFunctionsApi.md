# \TeamsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGetAllMessages**](TeamsFunctionsApi.md#TeamsGetAllMessages) | **Get** /teams/microsoft.graph.getAllMessages() | Invoke function getAllMessages
[**TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta**](TeamsFunctionsApi.md#TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies/microsoft.graph.delta() | Invoke function delta
[**TeamsTeamChannelsChannelMessagesDelta**](TeamsFunctionsApi.md#TeamsTeamChannelsChannelMessagesDelta) | **Get** /teams/{team-id}/channels/{channel-id}/messages/microsoft.graph.delta() | Invoke function delta
[**TeamsTeamChannelsGetAllMessages**](TeamsFunctionsApi.md#TeamsTeamChannelsGetAllMessages) | **Get** /teams/{team-id}/channels/microsoft.graph.getAllMessages() | Invoke function getAllMessages
[**TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta**](TeamsFunctionsApi.md#TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies/microsoft.graph.delta() | Invoke function delta
[**TeamsTeamPrimaryChannelMessagesDelta**](TeamsFunctionsApi.md#TeamsTeamPrimaryChannelMessagesDelta) | **Get** /teams/{team-id}/primaryChannel/messages/microsoft.graph.delta() | Invoke function delta



## TeamsGetAllMessages

> []*AnyOfmicrosoftGraphChatMessage TeamsGetAllMessages(ctx).Execute()

Invoke function getAllMessages

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsGetAllMessages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsGetAllMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetAllMessages`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsGetAllMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetAllMessagesRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta

> []*AnyOfmicrosoftGraphChatMessage TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta(ctx, teamId, channelId, chatMessageId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := "teamId_example" // string | key: id of team
    channelId := "channelId_example" // string | key: id of channel
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta(context.Background(), teamId, channelId, chatMessageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesChatMessageRepliesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelMessagesChatMessageRepliesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamChannelsChannelMessagesDelta

> []*AnyOfmicrosoftGraphChatMessage TeamsTeamChannelsChannelMessagesDelta(ctx, teamId, channelId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := "teamId_example" // string | key: id of team
    channelId := "channelId_example" // string | key: id of channel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesDelta(context.Background(), teamId, channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamChannelsChannelMessagesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsTeamChannelsChannelMessagesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelMessagesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamChannelsGetAllMessages

> []*AnyOfmicrosoftGraphChatMessage TeamsTeamChannelsGetAllMessages(ctx, teamId).Execute()

Invoke function getAllMessages

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := "teamId_example" // string | key: id of team

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsTeamChannelsGetAllMessages(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsTeamChannelsGetAllMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamChannelsGetAllMessages`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsTeamChannelsGetAllMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsGetAllMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta

> []*AnyOfmicrosoftGraphChatMessage TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta(ctx, teamId, chatMessageId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := "teamId_example" // string | key: id of team
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta(context.Background(), teamId, chatMessageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesChatMessageRepliesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelMessagesChatMessageRepliesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamPrimaryChannelMessagesDelta

> []*AnyOfmicrosoftGraphChatMessage TeamsTeamPrimaryChannelMessagesDelta(ctx, teamId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    teamId := "teamId_example" // string | key: id of team

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesDelta(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamPrimaryChannelMessagesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsFunctionsApi.TeamsTeamPrimaryChannelMessagesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelMessagesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

