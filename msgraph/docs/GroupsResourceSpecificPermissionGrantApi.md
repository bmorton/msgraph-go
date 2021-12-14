# \GroupsResourceSpecificPermissionGrantApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreatePermissionGrants**](GroupsResourceSpecificPermissionGrantApi.md#GroupsCreatePermissionGrants) | **Post** /groups/{group-id}/permissionGrants | Create new navigation property to permissionGrants for groups
[**GroupsDeletePermissionGrants**](GroupsResourceSpecificPermissionGrantApi.md#GroupsDeletePermissionGrants) | **Delete** /groups/{group-id}/permissionGrants/{resourceSpecificPermissionGrant-id} | Delete navigation property permissionGrants for groups
[**GroupsGetPermissionGrants**](GroupsResourceSpecificPermissionGrantApi.md#GroupsGetPermissionGrants) | **Get** /groups/{group-id}/permissionGrants/{resourceSpecificPermissionGrant-id} | Get permissionGrants from groups
[**GroupsListPermissionGrants**](GroupsResourceSpecificPermissionGrantApi.md#GroupsListPermissionGrants) | **Get** /groups/{group-id}/permissionGrants | Get permissionGrants from groups
[**GroupsUpdatePermissionGrants**](GroupsResourceSpecificPermissionGrantApi.md#GroupsUpdatePermissionGrants) | **Patch** /groups/{group-id}/permissionGrants/{resourceSpecificPermissionGrant-id} | Update the navigation property permissionGrants in groups



## GroupsCreatePermissionGrants

> MicrosoftGraphResourceSpecificPermissionGrant GroupsCreatePermissionGrants(ctx, groupId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()

Create new navigation property to permissionGrants for groups



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
    microsoftGraphResourceSpecificPermissionGrant := *openapiclient.NewMicrosoftGraphResourceSpecificPermissionGrant() // MicrosoftGraphResourceSpecificPermissionGrant | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsResourceSpecificPermissionGrantApi.GroupsCreatePermissionGrants(context.Background(), groupId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsResourceSpecificPermissionGrantApi.GroupsCreatePermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreatePermissionGrants`: MicrosoftGraphResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `GroupsResourceSpecificPermissionGrantApi.GroupsCreatePermissionGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreatePermissionGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphResourceSpecificPermissionGrant** | [**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md) | New navigation property | 

### Return type

[**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeletePermissionGrants

> GroupsDeletePermissionGrants(ctx, groupId, resourceSpecificPermissionGrantId).IfMatch(ifMatch).Execute()

Delete navigation property permissionGrants for groups



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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsResourceSpecificPermissionGrantApi.GroupsDeletePermissionGrants(context.Background(), groupId, resourceSpecificPermissionGrantId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsResourceSpecificPermissionGrantApi.GroupsDeletePermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeletePermissionGrantsRequest struct via the builder pattern


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


## GroupsGetPermissionGrants

> MicrosoftGraphResourceSpecificPermissionGrant GroupsGetPermissionGrants(ctx, groupId, resourceSpecificPermissionGrantId).Select_(select_).Expand(expand).Execute()

Get permissionGrants from groups



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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsResourceSpecificPermissionGrantApi.GroupsGetPermissionGrants(context.Background(), groupId, resourceSpecificPermissionGrantId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsResourceSpecificPermissionGrantApi.GroupsGetPermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPermissionGrants`: MicrosoftGraphResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `GroupsResourceSpecificPermissionGrantApi.GroupsGetPermissionGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPermissionGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListPermissionGrants

> CollectionOfResourceSpecificPermissionGrant GroupsListPermissionGrants(ctx, groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get permissionGrants from groups



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
    resp, r, err := api_client.GroupsResourceSpecificPermissionGrantApi.GroupsListPermissionGrants(context.Background(), groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsResourceSpecificPermissionGrantApi.GroupsListPermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListPermissionGrants`: CollectionOfResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `GroupsResourceSpecificPermissionGrantApi.GroupsListPermissionGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListPermissionGrantsRequest struct via the builder pattern


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

[**CollectionOfResourceSpecificPermissionGrant**](CollectionOfResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdatePermissionGrants

> GroupsUpdatePermissionGrants(ctx, groupId, resourceSpecificPermissionGrantId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()

Update the navigation property permissionGrants in groups



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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    microsoftGraphResourceSpecificPermissionGrant := *openapiclient.NewMicrosoftGraphResourceSpecificPermissionGrant() // MicrosoftGraphResourceSpecificPermissionGrant | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsResourceSpecificPermissionGrantApi.GroupsUpdatePermissionGrants(context.Background(), groupId, resourceSpecificPermissionGrantId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsResourceSpecificPermissionGrantApi.GroupsUpdatePermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePermissionGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphResourceSpecificPermissionGrant** | [**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md) | New navigation property values | 

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

