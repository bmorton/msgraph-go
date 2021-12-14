# \MeAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeCreateAppRoleAssignments) | **Post** /me/appRoleAssignments | Create new navigation property to appRoleAssignments for me
[**MeDeleteAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeDeleteAppRoleAssignments) | **Delete** /me/appRoleAssignments/{appRoleAssignment-id} | Delete navigation property appRoleAssignments for me
[**MeGetAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeGetAppRoleAssignments) | **Get** /me/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from me
[**MeListAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeListAppRoleAssignments) | **Get** /me/appRoleAssignments | Get appRoleAssignments from me
[**MeUpdateAppRoleAssignments**](MeAppRoleAssignmentApi.md#MeUpdateAppRoleAssignments) | **Patch** /me/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in me



## MeCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment MeCreateAppRoleAssignments(ctx).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Create new navigation property to appRoleAssignments for me



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
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAppRoleAssignmentApi.MeCreateAppRoleAssignments(context.Background()).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAppRoleAssignmentApi.MeCreateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `MeAppRoleAssignmentApi.MeCreateAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateAppRoleAssignmentsRequest struct via the builder pattern


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


## MeDeleteAppRoleAssignments

> MeDeleteAppRoleAssignments(ctx, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property appRoleAssignments for me



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAppRoleAssignmentApi.MeDeleteAppRoleAssignments(context.Background(), appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAppRoleAssignmentApi.MeDeleteAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteAppRoleAssignmentsRequest struct via the builder pattern


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


## MeGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment MeGetAppRoleAssignments(ctx, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from me



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAppRoleAssignmentApi.MeGetAppRoleAssignments(context.Background(), appRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAppRoleAssignmentApi.MeGetAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `MeAppRoleAssignmentApi.MeGetAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetAppRoleAssignmentsRequest struct via the builder pattern


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


## MeListAppRoleAssignments

> CollectionOfAppRoleAssignment MeListAppRoleAssignments(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from me



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
    resp, r, err := api_client.MeAppRoleAssignmentApi.MeListAppRoleAssignments(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAppRoleAssignmentApi.MeListAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListAppRoleAssignments`: CollectionOfAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `MeAppRoleAssignmentApi.MeListAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListAppRoleAssignmentsRequest struct via the builder pattern


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


## MeUpdateAppRoleAssignments

> MeUpdateAppRoleAssignments(ctx, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Update the navigation property appRoleAssignments in me



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
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAppRoleAssignmentApi.MeUpdateAppRoleAssignments(context.Background(), appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAppRoleAssignmentApi.MeUpdateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateAppRoleAssignmentsRequest struct via the builder pattern


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

