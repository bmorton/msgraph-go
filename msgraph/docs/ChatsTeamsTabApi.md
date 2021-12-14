# \ChatsTeamsTabApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsCreateTabs**](ChatsTeamsTabApi.md#ChatsCreateTabs) | **Post** /chats/{chat-id}/tabs | Create new navigation property to tabs for chats
[**ChatsDeleteTabs**](ChatsTeamsTabApi.md#ChatsDeleteTabs) | **Delete** /chats/{chat-id}/tabs/{teamsTab-id} | Delete navigation property tabs for chats
[**ChatsGetTabs**](ChatsTeamsTabApi.md#ChatsGetTabs) | **Get** /chats/{chat-id}/tabs/{teamsTab-id} | Get tabs from chats
[**ChatsListTabs**](ChatsTeamsTabApi.md#ChatsListTabs) | **Get** /chats/{chat-id}/tabs | Get tabs from chats
[**ChatsTabsDeleteRefTeamsApp**](ChatsTeamsTabApi.md#ChatsTabsDeleteRefTeamsApp) | **Delete** /chats/{chat-id}/tabs/{teamsTab-id}/teamsApp/$ref | Delete ref of navigation property teamsApp for chats
[**ChatsTabsGetRefTeamsApp**](ChatsTeamsTabApi.md#ChatsTabsGetRefTeamsApp) | **Get** /chats/{chat-id}/tabs/{teamsTab-id}/teamsApp/$ref | Get ref of teamsApp from chats
[**ChatsTabsGetTeamsApp**](ChatsTeamsTabApi.md#ChatsTabsGetTeamsApp) | **Get** /chats/{chat-id}/tabs/{teamsTab-id}/teamsApp | Get teamsApp from chats
[**ChatsTabsUpdateRefTeamsApp**](ChatsTeamsTabApi.md#ChatsTabsUpdateRefTeamsApp) | **Put** /chats/{chat-id}/tabs/{teamsTab-id}/teamsApp/$ref | Update the ref of navigation property teamsApp in chats
[**ChatsUpdateTabs**](ChatsTeamsTabApi.md#ChatsUpdateTabs) | **Patch** /chats/{chat-id}/tabs/{teamsTab-id} | Update the navigation property tabs in chats



## ChatsCreateTabs

> MicrosoftGraphTeamsTab ChatsCreateTabs(ctx, chatId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Create new navigation property to tabs for chats

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
    chatId := "chatId_example" // string | key: id of chat
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsCreateTabs(context.Background(), chatId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsCreateTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsCreateTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsTabApi.ChatsCreateTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsCreateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsDeleteTabs

> ChatsDeleteTabs(ctx, chatId, teamsTabId).IfMatch(ifMatch).Execute()

Delete navigation property tabs for chats

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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsDeleteTabs(context.Background(), chatId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsDeleteTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsDeleteTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsGetTabs

> MicrosoftGraphTeamsTab ChatsGetTabs(ctx, chatId, teamsTabId).Select_(select_).Expand(expand).Execute()

Get tabs from chats

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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsGetTabs(context.Background(), chatId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsGetTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsGetTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsTabApi.ChatsGetTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsGetTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsListTabs

> CollectionOfTeamsTab ChatsListTabs(ctx, chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tabs from chats

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
    chatId := "chatId_example" // string | key: id of chat
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsListTabs(context.Background(), chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsListTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsListTabs`: CollectionOfTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsTabApi.ChatsListTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsListTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfTeamsTab**](CollectionOfTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsTabsDeleteRefTeamsApp

> ChatsTabsDeleteRefTeamsApp(ctx, chatId, teamsTabId).IfMatch(ifMatch).Execute()

Delete ref of navigation property teamsApp for chats



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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsTabsDeleteRefTeamsApp(context.Background(), chatId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsTabsDeleteRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsTabsDeleteRefTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsTabsGetRefTeamsApp

> string ChatsTabsGetRefTeamsApp(ctx, chatId, teamsTabId).Execute()

Get ref of teamsApp from chats



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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsTabsGetRefTeamsApp(context.Background(), chatId, teamsTabId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsTabsGetRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsTabsGetRefTeamsApp`: string
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsTabApi.ChatsTabsGetRefTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsTabsGetRefTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsTabsGetTeamsApp

> MicrosoftGraphTeamsApp ChatsTabsGetTeamsApp(ctx, chatId, teamsTabId).Select_(select_).Expand(expand).Execute()

Get teamsApp from chats



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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsTabsGetTeamsApp(context.Background(), chatId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsTabsGetTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsTabsGetTeamsApp`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsTabApi.ChatsTabsGetTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsTabsGetTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsTabsUpdateRefTeamsApp

> ChatsTabsUpdateRefTeamsApp(ctx, chatId, teamsTabId).RequestBody(requestBody).Execute()

Update the ref of navigation property teamsApp in chats



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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsTabsUpdateRefTeamsApp(context.Background(), chatId, teamsTabId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsTabsUpdateRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsTabsUpdateRefTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsUpdateTabs

> ChatsUpdateTabs(ctx, chatId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Update the navigation property tabs in chats

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
    chatId := "chatId_example" // string | key: id of chat
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsTabApi.ChatsUpdateTabs(context.Background(), chatId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsTabApi.ChatsUpdateTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsUpdateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

