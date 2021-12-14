# \ChatsTeamsAppInstallationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsCreateInstalledApps**](ChatsTeamsAppInstallationApi.md#ChatsCreateInstalledApps) | **Post** /chats/{chat-id}/installedApps | Create new navigation property to installedApps for chats
[**ChatsDeleteInstalledApps**](ChatsTeamsAppInstallationApi.md#ChatsDeleteInstalledApps) | **Delete** /chats/{chat-id}/installedApps/{teamsAppInstallation-id} | Delete navigation property installedApps for chats
[**ChatsGetInstalledApps**](ChatsTeamsAppInstallationApi.md#ChatsGetInstalledApps) | **Get** /chats/{chat-id}/installedApps/{teamsAppInstallation-id} | Get installedApps from chats
[**ChatsInstalledAppsDeleteRefTeamsApp**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsDeleteRefTeamsApp) | **Delete** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Delete ref of navigation property teamsApp for chats
[**ChatsInstalledAppsDeleteRefTeamsAppDefinition**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsDeleteRefTeamsAppDefinition) | **Delete** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Delete ref of navigation property teamsAppDefinition for chats
[**ChatsInstalledAppsGetRefTeamsApp**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsGetRefTeamsApp) | **Get** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Get ref of teamsApp from chats
[**ChatsInstalledAppsGetRefTeamsAppDefinition**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsGetRefTeamsAppDefinition) | **Get** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Get ref of teamsAppDefinition from chats
[**ChatsInstalledAppsGetTeamsApp**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsGetTeamsApp) | **Get** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsApp | Get teamsApp from chats
[**ChatsInstalledAppsGetTeamsAppDefinition**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsGetTeamsAppDefinition) | **Get** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition | Get teamsAppDefinition from chats
[**ChatsInstalledAppsUpdateRefTeamsApp**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsUpdateRefTeamsApp) | **Put** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Update the ref of navigation property teamsApp in chats
[**ChatsInstalledAppsUpdateRefTeamsAppDefinition**](ChatsTeamsAppInstallationApi.md#ChatsInstalledAppsUpdateRefTeamsAppDefinition) | **Put** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Update the ref of navigation property teamsAppDefinition in chats
[**ChatsListInstalledApps**](ChatsTeamsAppInstallationApi.md#ChatsListInstalledApps) | **Get** /chats/{chat-id}/installedApps | Get installedApps from chats
[**ChatsUpdateInstalledApps**](ChatsTeamsAppInstallationApi.md#ChatsUpdateInstalledApps) | **Patch** /chats/{chat-id}/installedApps/{teamsAppInstallation-id} | Update the navigation property installedApps in chats



## ChatsCreateInstalledApps

> MicrosoftGraphTeamsAppInstallation ChatsCreateInstalledApps(ctx, chatId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()

Create new navigation property to installedApps for chats



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
    microsoftGraphTeamsAppInstallation := *openapiclient.NewMicrosoftGraphTeamsAppInstallation() // MicrosoftGraphTeamsAppInstallation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsCreateInstalledApps(context.Background(), chatId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsCreateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsCreateInstalledApps`: MicrosoftGraphTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsCreateInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsCreateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsDeleteInstalledApps

> ChatsDeleteInstalledApps(ctx, chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete navigation property installedApps for chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsDeleteInstalledApps(context.Background(), chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsDeleteInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsDeleteInstalledAppsRequest struct via the builder pattern


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


## ChatsGetInstalledApps

> MicrosoftGraphTeamsAppInstallation ChatsGetInstalledApps(ctx, chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get installedApps from chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsGetInstalledApps(context.Background(), chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsGetInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsGetInstalledApps`: MicrosoftGraphTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsGetInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsGetInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsInstalledAppsDeleteRefTeamsApp

> ChatsInstalledAppsDeleteRefTeamsApp(ctx, chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsDeleteRefTeamsApp(context.Background(), chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsDeleteRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsDeleteRefTeamsAppRequest struct via the builder pattern


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


## ChatsInstalledAppsDeleteRefTeamsAppDefinition

> ChatsInstalledAppsDeleteRefTeamsAppDefinition(ctx, chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property teamsAppDefinition for chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsDeleteRefTeamsAppDefinition(context.Background(), chatId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsDeleteRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsDeleteRefTeamsAppDefinitionRequest struct via the builder pattern


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


## ChatsInstalledAppsGetRefTeamsApp

> string ChatsInstalledAppsGetRefTeamsApp(ctx, chatId, teamsAppInstallationId).Execute()

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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsApp(context.Background(), chatId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsInstalledAppsGetRefTeamsApp`: string
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsGetRefTeamsAppRequest struct via the builder pattern


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


## ChatsInstalledAppsGetRefTeamsAppDefinition

> string ChatsInstalledAppsGetRefTeamsAppDefinition(ctx, chatId, teamsAppInstallationId).Execute()

Get ref of teamsAppDefinition from chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsAppDefinition(context.Background(), chatId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsInstalledAppsGetRefTeamsAppDefinition`: string
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetRefTeamsAppDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsGetRefTeamsAppDefinitionRequest struct via the builder pattern


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


## ChatsInstalledAppsGetTeamsApp

> MicrosoftGraphTeamsApp ChatsInstalledAppsGetTeamsApp(ctx, chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsApp(context.Background(), chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsInstalledAppsGetTeamsApp`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsGetTeamsAppRequest struct via the builder pattern


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


## ChatsInstalledAppsGetTeamsAppDefinition

> MicrosoftGraphTeamsAppDefinition ChatsInstalledAppsGetTeamsAppDefinition(ctx, chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get teamsAppDefinition from chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsAppDefinition(context.Background(), chatId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsInstalledAppsGetTeamsAppDefinition`: MicrosoftGraphTeamsAppDefinition
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsInstalledAppsGetTeamsAppDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsGetTeamsAppDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsInstalledAppsUpdateRefTeamsApp

> ChatsInstalledAppsUpdateRefTeamsApp(ctx, chatId, teamsAppInstallationId).RequestBody(requestBody).Execute()

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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsUpdateRefTeamsApp(context.Background(), chatId, teamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsUpdateRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsUpdateRefTeamsAppRequest struct via the builder pattern


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


## ChatsInstalledAppsUpdateRefTeamsAppDefinition

> ChatsInstalledAppsUpdateRefTeamsAppDefinition(ctx, chatId, teamsAppInstallationId).RequestBody(requestBody).Execute()

Update the ref of navigation property teamsAppDefinition in chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsInstalledAppsUpdateRefTeamsAppDefinition(context.Background(), chatId, teamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsInstalledAppsUpdateRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsInstalledAppsUpdateRefTeamsAppDefinitionRequest struct via the builder pattern


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


## ChatsListInstalledApps

> CollectionOfTeamsAppInstallation ChatsListInstalledApps(ctx, chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get installedApps from chats



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
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsListInstalledApps(context.Background(), chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsListInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsListInstalledApps`: CollectionOfTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `ChatsTeamsAppInstallationApi.ChatsListInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsListInstalledAppsRequest struct via the builder pattern


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

[**CollectionOfTeamsAppInstallation**](CollectionOfTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsUpdateInstalledApps

> ChatsUpdateInstalledApps(ctx, chatId, teamsAppInstallationId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()

Update the navigation property installedApps in chats



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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    microsoftGraphTeamsAppInstallation := *openapiclient.NewMicrosoftGraphTeamsAppInstallation() // MicrosoftGraphTeamsAppInstallation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsTeamsAppInstallationApi.ChatsUpdateInstalledApps(context.Background(), chatId, teamsAppInstallationId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsTeamsAppInstallationApi.ChatsUpdateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsUpdateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | New navigation property values | 

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

