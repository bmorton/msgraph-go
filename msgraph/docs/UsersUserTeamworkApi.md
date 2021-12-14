# \UsersUserTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteTeamwork**](UsersUserTeamworkApi.md#UsersDeleteTeamwork) | **Delete** /users/{user-id}/teamwork | Delete navigation property teamwork for users
[**UsersGetTeamwork**](UsersUserTeamworkApi.md#UsersGetTeamwork) | **Get** /users/{user-id}/teamwork | Get teamwork from users
[**UsersTeamworkCreateInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkCreateInstalledApps) | **Post** /users/{user-id}/teamwork/installedApps | Create new navigation property to installedApps for users
[**UsersTeamworkDeleteInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkDeleteInstalledApps) | **Delete** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Delete navigation property installedApps for users
[**UsersTeamworkGetInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkGetInstalledApps) | **Get** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Get installedApps from users
[**UsersTeamworkInstalledAppsDeleteRefChat**](UsersUserTeamworkApi.md#UsersTeamworkInstalledAppsDeleteRefChat) | **Delete** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Delete ref of navigation property chat for users
[**UsersTeamworkInstalledAppsGetChat**](UsersUserTeamworkApi.md#UsersTeamworkInstalledAppsGetChat) | **Get** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat | Get chat from users
[**UsersTeamworkInstalledAppsGetRefChat**](UsersUserTeamworkApi.md#UsersTeamworkInstalledAppsGetRefChat) | **Get** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Get ref of chat from users
[**UsersTeamworkInstalledAppsUpdateRefChat**](UsersUserTeamworkApi.md#UsersTeamworkInstalledAppsUpdateRefChat) | **Put** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id}/chat/$ref | Update the ref of navigation property chat in users
[**UsersTeamworkListInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkListInstalledApps) | **Get** /users/{user-id}/teamwork/installedApps | Get installedApps from users
[**UsersTeamworkUpdateInstalledApps**](UsersUserTeamworkApi.md#UsersTeamworkUpdateInstalledApps) | **Patch** /users/{user-id}/teamwork/installedApps/{userScopeTeamsAppInstallation-id} | Update the navigation property installedApps in users
[**UsersUpdateTeamwork**](UsersUserTeamworkApi.md#UsersUpdateTeamwork) | **Patch** /users/{user-id}/teamwork | Update the navigation property teamwork in users



## UsersDeleteTeamwork

> UsersDeleteTeamwork(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property teamwork for users



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
    userId := "userId_example" // string | key: id of user
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersDeleteTeamwork(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersDeleteTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteTeamworkRequest struct via the builder pattern


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


## UsersGetTeamwork

> MicrosoftGraphUserTeamwork UsersGetTeamwork(ctx, userId).Select_(select_).Expand(expand).Execute()

Get teamwork from users



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
    userId := "userId_example" // string | key: id of user
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersGetTeamwork(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersGetTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetTeamwork`: MicrosoftGraphUserTeamwork
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersGetTeamwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetTeamworkRequest struct via the builder pattern


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


## UsersTeamworkCreateInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation UsersTeamworkCreateInstalledApps(ctx, userId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()

Create new navigation property to installedApps for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphUserScopeTeamsAppInstallation := *openapiclient.NewMicrosoftGraphUserScopeTeamsAppInstallation() // MicrosoftGraphUserScopeTeamsAppInstallation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkCreateInstalledApps(context.Background(), userId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkCreateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTeamworkCreateInstalledApps`: MicrosoftGraphUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersTeamworkCreateInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkCreateInstalledAppsRequest struct via the builder pattern


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


## UsersTeamworkDeleteInstalledApps

> UsersTeamworkDeleteInstalledApps(ctx, userId, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete navigation property installedApps for users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkDeleteInstalledApps(context.Background(), userId, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkDeleteInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkDeleteInstalledAppsRequest struct via the builder pattern


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


## UsersTeamworkGetInstalledApps

> MicrosoftGraphUserScopeTeamsAppInstallation UsersTeamworkGetInstalledApps(ctx, userId, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get installedApps from users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkGetInstalledApps(context.Background(), userId, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkGetInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTeamworkGetInstalledApps`: MicrosoftGraphUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersTeamworkGetInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkGetInstalledAppsRequest struct via the builder pattern


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


## UsersTeamworkInstalledAppsDeleteRefChat

> UsersTeamworkInstalledAppsDeleteRefChat(ctx, userId, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property chat for users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkInstalledAppsDeleteRefChat(context.Background(), userId, userScopeTeamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkInstalledAppsDeleteRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkInstalledAppsDeleteRefChatRequest struct via the builder pattern


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


## UsersTeamworkInstalledAppsGetChat

> MicrosoftGraphChat UsersTeamworkInstalledAppsGetChat(ctx, userId, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get chat from users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetChat(context.Background(), userId, userScopeTeamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTeamworkInstalledAppsGetChat`: MicrosoftGraphChat
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkInstalledAppsGetChatRequest struct via the builder pattern


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


## UsersTeamworkInstalledAppsGetRefChat

> string UsersTeamworkInstalledAppsGetRefChat(ctx, userId, userScopeTeamsAppInstallationId).Execute()

Get ref of chat from users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetRefChat(context.Background(), userId, userScopeTeamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTeamworkInstalledAppsGetRefChat`: string
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersTeamworkInstalledAppsGetRefChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkInstalledAppsGetRefChatRequest struct via the builder pattern


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


## UsersTeamworkInstalledAppsUpdateRefChat

> UsersTeamworkInstalledAppsUpdateRefChat(ctx, userId, userScopeTeamsAppInstallationId).RequestBody(requestBody).Execute()

Update the ref of navigation property chat in users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkInstalledAppsUpdateRefChat(context.Background(), userId, userScopeTeamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkInstalledAppsUpdateRefChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkInstalledAppsUpdateRefChatRequest struct via the builder pattern


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


## UsersTeamworkListInstalledApps

> CollectionOfUserScopeTeamsAppInstallation UsersTeamworkListInstalledApps(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get installedApps from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkListInstalledApps(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkListInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTeamworkListInstalledApps`: CollectionOfUserScopeTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `UsersUserTeamworkApi.UsersTeamworkListInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkListInstalledAppsRequest struct via the builder pattern


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


## UsersTeamworkUpdateInstalledApps

> UsersTeamworkUpdateInstalledApps(ctx, userId, userScopeTeamsAppInstallationId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()

Update the navigation property installedApps in users



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
    userId := "userId_example" // string | key: id of user
    userScopeTeamsAppInstallationId := "userScopeTeamsAppInstallationId_example" // string | key: id of userScopeTeamsAppInstallation
    microsoftGraphUserScopeTeamsAppInstallation := *openapiclient.NewMicrosoftGraphUserScopeTeamsAppInstallation() // MicrosoftGraphUserScopeTeamsAppInstallation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersTeamworkUpdateInstalledApps(context.Background(), userId, userScopeTeamsAppInstallationId).MicrosoftGraphUserScopeTeamsAppInstallation(microsoftGraphUserScopeTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersTeamworkUpdateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userScopeTeamsAppInstallationId** | **string** | key: id of userScopeTeamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTeamworkUpdateInstalledAppsRequest struct via the builder pattern


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


## UsersUpdateTeamwork

> UsersUpdateTeamwork(ctx, userId).MicrosoftGraphUserTeamwork(microsoftGraphUserTeamwork).Execute()

Update the navigation property teamwork in users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphUserTeamwork := *openapiclient.NewMicrosoftGraphUserTeamwork() // MicrosoftGraphUserTeamwork | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserTeamworkApi.UsersUpdateTeamwork(context.Background(), userId).MicrosoftGraphUserTeamwork(microsoftGraphUserTeamwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserTeamworkApi.UsersUpdateTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateTeamworkRequest struct via the builder pattern


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

