# \UsersScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersCreateScopedRoleMemberOf) | **Post** /users/{user-id}/scopedRoleMemberOf | Create new navigation property to scopedRoleMemberOf for users
[**UsersDeleteScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersDeleteScopedRoleMemberOf) | **Delete** /users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id} | Delete navigation property scopedRoleMemberOf for users
[**UsersGetScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersGetScopedRoleMemberOf) | **Get** /users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id} | Get scopedRoleMemberOf from users
[**UsersListScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersListScopedRoleMemberOf) | **Get** /users/{user-id}/scopedRoleMemberOf | Get scopedRoleMemberOf from users
[**UsersUpdateScopedRoleMemberOf**](UsersScopedRoleMembershipApi.md#UsersUpdateScopedRoleMemberOf) | **Patch** /users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id} | Update the navigation property scopedRoleMemberOf in users



## UsersCreateScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership UsersCreateScopedRoleMemberOf(ctx, userId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Create new navigation property to scopedRoleMemberOf for users



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
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersScopedRoleMembershipApi.UsersCreateScopedRoleMemberOf(context.Background(), userId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersScopedRoleMembershipApi.UsersCreateScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateScopedRoleMemberOf`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `UsersScopedRoleMembershipApi.UsersCreateScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New navigation property | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteScopedRoleMemberOf

> UsersDeleteScopedRoleMemberOf(ctx, userId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()

Delete navigation property scopedRoleMemberOf for users



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersScopedRoleMembershipApi.UsersDeleteScopedRoleMemberOf(context.Background(), userId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersScopedRoleMembershipApi.UsersDeleteScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteScopedRoleMemberOfRequest struct via the builder pattern


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


## UsersGetScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership UsersGetScopedRoleMemberOf(ctx, userId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()

Get scopedRoleMemberOf from users



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersScopedRoleMembershipApi.UsersGetScopedRoleMemberOf(context.Background(), userId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersScopedRoleMembershipApi.UsersGetScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetScopedRoleMemberOf`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `UsersScopedRoleMembershipApi.UsersGetScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListScopedRoleMemberOf

> CollectionOfScopedRoleMembership UsersListScopedRoleMemberOf(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get scopedRoleMemberOf from users



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
    resp, r, err := api_client.UsersScopedRoleMembershipApi.UsersListScopedRoleMemberOf(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersScopedRoleMembershipApi.UsersListScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListScopedRoleMemberOf`: CollectionOfScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `UsersScopedRoleMembershipApi.UsersListScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListScopedRoleMemberOfRequest struct via the builder pattern


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

[**CollectionOfScopedRoleMembership**](CollectionOfScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateScopedRoleMemberOf

> UsersUpdateScopedRoleMemberOf(ctx, userId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Update the navigation property scopedRoleMemberOf in users



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersScopedRoleMembershipApi.UsersUpdateScopedRoleMemberOf(context.Background(), userId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersScopedRoleMembershipApi.UsersUpdateScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New navigation property values | 

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

