# \UsersExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateExtensions**](UsersExtensionApi.md#UsersCreateExtensions) | **Post** /users/{user-id}/extensions | Create new navigation property to extensions for users
[**UsersDeleteExtensions**](UsersExtensionApi.md#UsersDeleteExtensions) | **Delete** /users/{user-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersGetExtensions**](UsersExtensionApi.md#UsersGetExtensions) | **Get** /users/{user-id}/extensions/{extension-id} | Get extensions from users
[**UsersListExtensions**](UsersExtensionApi.md#UsersListExtensions) | **Get** /users/{user-id}/extensions | Get extensions from users
[**UsersUpdateExtensions**](UsersExtensionApi.md#UsersUpdateExtensions) | **Patch** /users/{user-id}/extensions/{extension-id} | Update the navigation property extensions in users



## UsersCreateExtensions

> MicrosoftGraphExtension UsersCreateExtensions(ctx, userId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for users



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
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersExtensionApi.UsersCreateExtensions(context.Background(), userId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersExtensionApi.UsersCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersExtensionApi.UsersCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteExtensions

> UsersDeleteExtensions(ctx, userId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for users



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
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersExtensionApi.UsersDeleteExtensions(context.Background(), userId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersExtensionApi.UsersDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteExtensionsRequest struct via the builder pattern


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


## UsersGetExtensions

> MicrosoftGraphExtension UsersGetExtensions(ctx, userId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersExtensionApi.UsersGetExtensions(context.Background(), userId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersExtensionApi.UsersGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersExtensionApi.UsersGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListExtensions

> CollectionOfExtension UsersListExtensions(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    resp, r, err := api_client.UsersExtensionApi.UsersListExtensions(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersExtensionApi.UsersListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersExtensionApi.UsersListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListExtensionsRequest struct via the builder pattern


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

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateExtensions

> UsersUpdateExtensions(ctx, userId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in users



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
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersExtensionApi.UsersUpdateExtensions(context.Background(), userId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersExtensionApi.UsersUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property values | 

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

