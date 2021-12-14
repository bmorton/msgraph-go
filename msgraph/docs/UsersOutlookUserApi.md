# \UsersOutlookUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteOutlook**](UsersOutlookUserApi.md#UsersDeleteOutlook) | **Delete** /users/{user-id}/outlook | Delete navigation property outlook for users
[**UsersGetOutlook**](UsersOutlookUserApi.md#UsersGetOutlook) | **Get** /users/{user-id}/outlook | Get outlook from users
[**UsersOutlookCreateMasterCategories**](UsersOutlookUserApi.md#UsersOutlookCreateMasterCategories) | **Post** /users/{user-id}/outlook/masterCategories | Create new navigation property to masterCategories for users
[**UsersOutlookDeleteMasterCategories**](UsersOutlookUserApi.md#UsersOutlookDeleteMasterCategories) | **Delete** /users/{user-id}/outlook/masterCategories/{outlookCategory-id} | Delete navigation property masterCategories for users
[**UsersOutlookGetMasterCategories**](UsersOutlookUserApi.md#UsersOutlookGetMasterCategories) | **Get** /users/{user-id}/outlook/masterCategories/{outlookCategory-id} | Get masterCategories from users
[**UsersOutlookListMasterCategories**](UsersOutlookUserApi.md#UsersOutlookListMasterCategories) | **Get** /users/{user-id}/outlook/masterCategories | Get masterCategories from users
[**UsersOutlookUpdateMasterCategories**](UsersOutlookUserApi.md#UsersOutlookUpdateMasterCategories) | **Patch** /users/{user-id}/outlook/masterCategories/{outlookCategory-id} | Update the navigation property masterCategories in users
[**UsersUpdateOutlook**](UsersOutlookUserApi.md#UsersUpdateOutlook) | **Patch** /users/{user-id}/outlook | Update the navigation property outlook in users



## UsersDeleteOutlook

> UsersDeleteOutlook(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property outlook for users



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
    resp, r, err := api_client.UsersOutlookUserApi.UsersDeleteOutlook(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersDeleteOutlook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersDeleteOutlookRequest struct via the builder pattern


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


## UsersGetOutlook

> MicrosoftGraphOutlookUser UsersGetOutlook(ctx, userId).Select_(select_).Execute()

Get outlook from users



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersGetOutlook(context.Background(), userId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersGetOutlook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetOutlook`: MicrosoftGraphOutlookUser
    fmt.Fprintf(os.Stdout, "Response from `UsersOutlookUserApi.UsersGetOutlook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetOutlookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOutlookCreateMasterCategories

> MicrosoftGraphOutlookCategory UsersOutlookCreateMasterCategories(ctx, userId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()

Create new navigation property to masterCategories for users



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
    microsoftGraphOutlookCategory := *openapiclient.NewMicrosoftGraphOutlookCategory() // MicrosoftGraphOutlookCategory | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersOutlookCreateMasterCategories(context.Background(), userId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersOutlookCreateMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersOutlookCreateMasterCategories`: MicrosoftGraphOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `UsersOutlookUserApi.UsersOutlookCreateMasterCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOutlookCreateMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md) | New navigation property | 

### Return type

[**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOutlookDeleteMasterCategories

> UsersOutlookDeleteMasterCategories(ctx, userId, outlookCategoryId).IfMatch(ifMatch).Execute()

Delete navigation property masterCategories for users



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersOutlookDeleteMasterCategories(context.Background(), userId, outlookCategoryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersOutlookDeleteMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOutlookDeleteMasterCategoriesRequest struct via the builder pattern


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


## UsersOutlookGetMasterCategories

> MicrosoftGraphOutlookCategory UsersOutlookGetMasterCategories(ctx, userId, outlookCategoryId).Select_(select_).Execute()

Get masterCategories from users



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersOutlookGetMasterCategories(context.Background(), userId, outlookCategoryId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersOutlookGetMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersOutlookGetMasterCategories`: MicrosoftGraphOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `UsersOutlookUserApi.UsersOutlookGetMasterCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOutlookGetMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOutlookListMasterCategories

> CollectionOfOutlookCategory UsersOutlookListMasterCategories(ctx, userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get masterCategories from users



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
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersOutlookListMasterCategories(context.Background(), userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersOutlookListMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersOutlookListMasterCategories`: CollectionOfOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `UsersOutlookUserApi.UsersOutlookListMasterCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOutlookListMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfOutlookCategory**](CollectionOfOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOutlookUpdateMasterCategories

> UsersOutlookUpdateMasterCategories(ctx, userId, outlookCategoryId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()

Update the navigation property masterCategories in users



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    microsoftGraphOutlookCategory := *openapiclient.NewMicrosoftGraphOutlookCategory() // MicrosoftGraphOutlookCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersOutlookUpdateMasterCategories(context.Background(), userId, outlookCategoryId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersOutlookUpdateMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersOutlookUpdateMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md) | New navigation property values | 

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


## UsersUpdateOutlook

> UsersUpdateOutlook(ctx, userId).MicrosoftGraphOutlookUser(microsoftGraphOutlookUser).Execute()

Update the navigation property outlook in users



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
    microsoftGraphOutlookUser := *openapiclient.NewMicrosoftGraphOutlookUser() // MicrosoftGraphOutlookUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOutlookUserApi.UsersUpdateOutlook(context.Background(), userId).MicrosoftGraphOutlookUser(microsoftGraphOutlookUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOutlookUserApi.UsersUpdateOutlook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersUpdateOutlookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOutlookUser** | [**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md) | New navigation property values | 

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

