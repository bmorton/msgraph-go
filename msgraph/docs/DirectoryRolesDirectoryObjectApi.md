# \DirectoryRolesDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesCreateRefMembers**](DirectoryRolesDirectoryObjectApi.md#DirectoryRolesCreateRefMembers) | **Post** /directoryRoles/{directoryRole-id}/members/$ref | Create new navigation property ref to members for directoryRoles
[**DirectoryRolesListMembers**](DirectoryRolesDirectoryObjectApi.md#DirectoryRolesListMembers) | **Get** /directoryRoles/{directoryRole-id}/members | Get members from directoryRoles
[**DirectoryRolesListRefMembers**](DirectoryRolesDirectoryObjectApi.md#DirectoryRolesListRefMembers) | **Get** /directoryRoles/{directoryRole-id}/members/$ref | Get ref of members from directoryRoles



## DirectoryRolesCreateRefMembers

> map[string]interface{} DirectoryRolesCreateRefMembers(ctx, directoryRoleId).RequestBody(requestBody).Execute()

Create new navigation property ref to members for directoryRoles



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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryObjectApi.DirectoryRolesCreateRefMembers(context.Background(), directoryRoleId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryObjectApi.DirectoryRolesCreateRefMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesCreateRefMembers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryObjectApi.DirectoryRolesCreateRefMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesCreateRefMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesListMembers

> CollectionOfDirectoryObject DirectoryRolesListMembers(ctx, directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from directoryRoles



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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
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
    resp, r, err := api_client.DirectoryRolesDirectoryObjectApi.DirectoryRolesListMembers(context.Background(), directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryObjectApi.DirectoryRolesListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesListMembers`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryObjectApi.DirectoryRolesListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesListMembersRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesListRefMembers

> CollectionOfLinksOfDirectoryObject DirectoryRolesListRefMembers(ctx, directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of members from directoryRoles



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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryObjectApi.DirectoryRolesListRefMembers(context.Background(), directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryObjectApi.DirectoryRolesListRefMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesListRefMembers`: CollectionOfLinksOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryObjectApi.DirectoryRolesListRefMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesListRefMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfDirectoryObject**](CollectionOfLinksOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

