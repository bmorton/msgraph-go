# \GroupsAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsCreateAppRoleAssignments) | **Post** /groups/{group-id}/appRoleAssignments | Create new navigation property to appRoleAssignments for groups
[**GroupsDeleteAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsDeleteAppRoleAssignments) | **Delete** /groups/{group-id}/appRoleAssignments/{appRoleAssignment-id} | Delete navigation property appRoleAssignments for groups
[**GroupsGetAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsGetAppRoleAssignments) | **Get** /groups/{group-id}/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from groups
[**GroupsListAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsListAppRoleAssignments) | **Get** /groups/{group-id}/appRoleAssignments | Get appRoleAssignments from groups
[**GroupsUpdateAppRoleAssignments**](GroupsAppRoleAssignmentApi.md#GroupsUpdateAppRoleAssignments) | **Patch** /groups/{group-id}/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in groups



## GroupsCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment GroupsCreateAppRoleAssignments(ctx, groupId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Create new navigation property to appRoleAssignments for groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsAppRoleAssignmentApi.GroupsCreateAppRoleAssignments(context.Background(), groupId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAppRoleAssignmentApi.GroupsCreateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GroupsAppRoleAssignmentApi.GroupsCreateAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateAppRoleAssignmentsRequest struct via the builder pattern


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


## GroupsDeleteAppRoleAssignments

> GroupsDeleteAppRoleAssignments(ctx, groupId, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property appRoleAssignments for groups



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
    groupId := "groupId_example" // string | key: id of group
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsAppRoleAssignmentApi.GroupsDeleteAppRoleAssignments(context.Background(), groupId, appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAppRoleAssignmentApi.GroupsDeleteAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteAppRoleAssignmentsRequest struct via the builder pattern


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


## GroupsGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment GroupsGetAppRoleAssignments(ctx, groupId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from groups



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
    groupId := "groupId_example" // string | key: id of group
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsAppRoleAssignmentApi.GroupsGetAppRoleAssignments(context.Background(), groupId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAppRoleAssignmentApi.GroupsGetAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GroupsAppRoleAssignmentApi.GroupsGetAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetAppRoleAssignmentsRequest struct via the builder pattern


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


## GroupsListAppRoleAssignments

> CollectionOfAppRoleAssignment GroupsListAppRoleAssignments(ctx, groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from groups



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
    groupId := "groupId_example" // string | key: id of group
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
    resp, r, err := api_client.GroupsAppRoleAssignmentApi.GroupsListAppRoleAssignments(context.Background(), groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAppRoleAssignmentApi.GroupsListAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListAppRoleAssignments`: CollectionOfAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GroupsAppRoleAssignmentApi.GroupsListAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListAppRoleAssignmentsRequest struct via the builder pattern


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


## GroupsUpdateAppRoleAssignments

> GroupsUpdateAppRoleAssignments(ctx, groupId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Update the navigation property appRoleAssignments in groups



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
    groupId := "groupId_example" // string | key: id of group
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsAppRoleAssignmentApi.GroupsUpdateAppRoleAssignments(context.Background(), groupId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsAppRoleAssignmentApi.GroupsUpdateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateAppRoleAssignmentsRequest struct via the builder pattern


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

