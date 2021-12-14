# \UsersAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateAppRoleAssignments**](UsersAppRoleAssignmentApi.md#UsersCreateAppRoleAssignments) | **Post** /users/{user-id}/appRoleAssignments | Create new navigation property to appRoleAssignments for users
[**UsersDeleteAppRoleAssignments**](UsersAppRoleAssignmentApi.md#UsersDeleteAppRoleAssignments) | **Delete** /users/{user-id}/appRoleAssignments/{appRoleAssignment-id} | Delete navigation property appRoleAssignments for users
[**UsersGetAppRoleAssignments**](UsersAppRoleAssignmentApi.md#UsersGetAppRoleAssignments) | **Get** /users/{user-id}/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from users
[**UsersListAppRoleAssignments**](UsersAppRoleAssignmentApi.md#UsersListAppRoleAssignments) | **Get** /users/{user-id}/appRoleAssignments | Get appRoleAssignments from users
[**UsersUpdateAppRoleAssignments**](UsersAppRoleAssignmentApi.md#UsersUpdateAppRoleAssignments) | **Patch** /users/{user-id}/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in users



## UsersCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment UsersCreateAppRoleAssignments(ctx, userId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Create new navigation property to appRoleAssignments for users



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
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAppRoleAssignmentApi.UsersCreateAppRoleAssignments(context.Background(), userId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAppRoleAssignmentApi.UsersCreateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UsersAppRoleAssignmentApi.UsersCreateAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteAppRoleAssignments

> UsersDeleteAppRoleAssignments(ctx, userId, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property appRoleAssignments for users



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAppRoleAssignmentApi.UsersDeleteAppRoleAssignments(context.Background(), userId, appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAppRoleAssignmentApi.UsersDeleteAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteAppRoleAssignmentsRequest struct via the builder pattern


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


## UsersGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment UsersGetAppRoleAssignments(ctx, userId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from users



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAppRoleAssignmentApi.UsersGetAppRoleAssignments(context.Background(), userId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAppRoleAssignmentApi.UsersGetAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UsersAppRoleAssignmentApi.UsersGetAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListAppRoleAssignments

> CollectionOfAppRoleAssignment UsersListAppRoleAssignments(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from users



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
    resp, r, err := api_client.UsersAppRoleAssignmentApi.UsersListAppRoleAssignments(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAppRoleAssignmentApi.UsersListAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListAppRoleAssignments`: CollectionOfAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UsersAppRoleAssignmentApi.UsersListAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListAppRoleAssignmentsRequest struct via the builder pattern


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

[**CollectionOfAppRoleAssignment**](CollectionOfAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateAppRoleAssignments

> UsersUpdateAppRoleAssignments(ctx, userId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Update the navigation property appRoleAssignments in users



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAppRoleAssignmentApi.UsersUpdateAppRoleAssignments(context.Background(), userId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAppRoleAssignmentApi.UsersUpdateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property values | 

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

