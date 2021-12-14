# \MeUserTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteTeamwork**](MeUserTeamworkApi.md#MeDeleteTeamwork) | **Delete** /me/teamwork | Delete navigation property teamwork for me
[**MeGetTeamwork**](MeUserTeamworkApi.md#MeGetTeamwork) | **Get** /me/teamwork | Get teamwork from me
[**MeTeamworkCreateInstalledApps**](MeUserTeamworkApi.md#MeTeamworkCreateInstalledApps) | **Post** /me/teamwork/installedApps | Create new navigation property to installedApps for me
[**MeTeamworkDeleteInstalledApps**](MeUserTeamworkApi.md#MeTeamworkDeleteInstalledApps) | **Delete** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Delete navigation property installedApps for me
[**MeTeamworkGetInstalledApps**](MeUserTeamworkApi.md#MeTeamworkGetInstalledApps) | **Get** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Get installedApps from me
[**MeTeamworkInstalledAppsDeleteRefChat**](MeUserTeamworkApi.md#MeTeamworkInstalledAppsDeleteRefChat) | **Delete** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Delete ref of navigation property chat for me
[**MeTeamworkInstalledAppsGetChat**](MeUserTeamworkApi.md#MeTeamworkInstalledAppsGetChat) | **Get** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat | Get chat from me
[**MeTeamworkInstalledAppsGetRefChat**](MeUserTeamworkApi.md#MeTeamworkInstalledAppsGetRefChat) | **Get** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Get ref of chat from me
[**MeTeamworkInstalledAppsUpdateRefChat**](MeUserTeamworkApi.md#MeTeamworkInstalledAppsUpdateRefChat) | **Put** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Update the ref of navigation property chat in me
[**MeTeamworkListInstalledApps**](MeUserTeamworkApi.md#MeTeamworkListInstalledApps) | **Get** /me/teamwork/installedApps | Get installedApps from me
[**MeTeamworkUpdateInstalledApps**](MeUserTeamworkApi.md#MeTeamworkUpdateInstalledApps) | **Patch** /me/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Update the navigation property installedApps in me
[**MeUpdateTeamwork**](MeUserTeamworkApi.md#MeUpdateTeamwork) | **Patch** /me/teamwork | Update the navigation property teamwork in me



## MeDeleteTeamwork

> MeDeleteTeamwork(ctx).IfMatch(ifMatch).Execute()

Delete navigation property teamwork for me



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeDeleteTeamwork(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeDeleteTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteTeamworkRequest struct via the builder pattern


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


## MeGetTeamwork

> MicrosoftGraphUserTeamwork MeGetTeamwork(ctx).Select_(select_).Expand(expand).Execute()

Get teamwork from me



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeGetTeamwork(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeGetTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetTeamwork`: MicrosoftGraphUserTeamwork
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeGetTeamwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetTeamworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUserTeamwork**](MicrosoftGraphUserTeamwork.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTeamworkCreateInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation MeTeamworkCreateInstalledApps(ctx).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()

Create new navigation property to installedApps for me



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
    microsoftGraphUserScopeTeamsAppInstallation := *openapiclient.NewMicrosoftGraphUserScopeTeamsAppInstallation() // MicrosoftGraphUserScopeTeamsAppInstallation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkCreateInstalledApps(context.Background()).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkCreateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTeamworkCreateInstalledApps`: MicrosoftGraphUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeTeamworkCreateInstalledApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkCreateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphUserScopeTeamsAppInstallation** | [**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md) | New navigation property | 

### Return type

[**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTeamworkDeleteInstalledApps

> MeTeamworkDeleteInstalledApps(ctx, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete navigation property installedApps for me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkDeleteInstalledApps(context.Background(), userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkDeleteInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkDeleteInstalledAppsRequest struct via the builder pattern


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


## MeTeamworkGetInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation MeTeamworkGetInstalledApps(ctx, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get installedApps from me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkGetInstalledApps(context.Background(), userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkGetInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTeamworkGetInstalledApps`: MicrosoftGraphUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeTeamworkGetInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkGetInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTeamworkInstalledAppsDeleteRefChat

> MeTeamworkInstalledAppsDeleteRefChat(ctx, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property chat for me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkInstalledAppsDeleteRefChat(context.Background(), userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkInstalledAppsDeleteRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkInstalledAppsDeleteRefChatRequest struct via the builder pattern


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


## MeTeamworkInstalledAppsGetChat

> MicrosoftGraphChat MeTeamworkInstalledAppsGetChat(ctx, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get chat from me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkInstalledAppsGetChat(context.Background(), userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkInstalledAppsGetChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTeamworkInstalledAppsGetChat`: MicrosoftGraphChat
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeTeamworkInstalledAppsGetChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkInstalledAppsGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChat**](MicrosoftGraphChat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTeamworkInstalledAppsGetRefChat

> string MeTeamworkInstalledAppsGetRefChat(ctx, userScopeTeamsAppInstallationId).Execute()

Get ref of chat from me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkInstalledAppsGetRefChat(context.Background(), userScopeTeamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkInstalledAppsGetRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTeamworkInstalledAppsGetRefChat`: string
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeTeamworkInstalledAppsGetRefChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkInstalledAppsGetRefChatRequest struct via the builder pattern


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


## MeTeamworkInstalledAppsUpdateRefChat

> MeTeamworkInstalledAppsUpdateRefChat(ctx, userScopeTeamsAppInstallationId).RequestBody(requestBody).Execute()

Update the ref of navigation property chat in me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkInstalledAppsUpdateRefChat(context.Background(), userScopeTeamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkInstalledAppsUpdateRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkInstalledAppsUpdateRefChatRequest struct via the builder pattern


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


## MeTeamworkListInstalledApps

> CollectionOfUserScopeTeamsAppInstallation MeTeamworkListInstalledApps(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get installedApps from me



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
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkListInstalledApps(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkListInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTeamworkListInstalledApps`: CollectionOfUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `MeUserTeamworkApi.MeTeamworkListInstalledApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkListInstalledAppsRequest struct via the builder pattern


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

[**CollectionOfUserScopeTeamsAppInstallation**](CollectionOfUserScopeTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTeamworkUpdateInstalledApps

> MeTeamworkUpdateInstalledApps(ctx, userScopeTeamsAppInstallationId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()

Update the navigation property installedApps in me



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
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    microsoftGraphUserScopeTeamsAppInstallation := *openapiclient.NewMicrosoftGraphUserScopeTeamsAppInstallation() // MicrosoftGraphUserScopeTeamsAppInstallation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeTeamworkUpdateInstalledApps(context.Background(), userScopeTeamsAppInstallationId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeTeamworkUpdateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTeamworkUpdateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphUserScopeTeamsAppInstallation** | [**MicrosoftGraphUserScopeTeamsAppInstallation**](MicrosoftGraphUserScopeTeamsAppInstallation.md) | New navigation property values | 

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


## MeUpdateTeamwork

> MeUpdateTeamwork(ctx).MicrosoftGraphUserTeamwork(microsoftGraphUserTeamwork).Execute()

Update the navigation property teamwork in me



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
    microsoftGraphUserTeamwork := *openapiclient.NewMicrosoftGraphUserTeamwork() // MicrosoftGraphUserTeamwork | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserTeamworkApi.MeUpdateTeamwork(context.Background()).MicrosoftGraphUserTeamwork(microsoftGraphUserTeamwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserTeamworkApi.MeUpdateTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateTeamworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphUserTeamwork** | [**MicrosoftGraphUserTeamwork**](MicrosoftGraphUserTeamwork.md) | New navigation property values | 

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

