# \DirectoryRolesScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesCreateScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesCreateScopedMembers) | **Post** /directoryRoles/{directoryRole-id}/scopedMembers | Create new navigation property to scopedMembers for directoryRoles
[**DirectoryRolesDeleteScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesDeleteScopedMembers) | **Delete** /directoryRoles/{directoryRole-id}/scopedMembers/{scopedRoleMembership-id} | Delete navigation property scopedMembers for directoryRoles
[**DirectoryRolesGetScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesGetScopedMembers) | **Get** /directoryRoles/{directoryRole-id}/scopedMembers/{scopedRoleMembership-id} | Get scopedMembers from directoryRoles
[**DirectoryRolesListScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesListScopedMembers) | **Get** /directoryRoles/{directoryRole-id}/scopedMembers | Get scopedMembers from directoryRoles
[**DirectoryRolesUpdateScopedMembers**](DirectoryRolesScopedRoleMembershipApi.md#DirectoryRolesUpdateScopedMembers) | **Patch** /directoryRoles/{directoryRole-id}/scopedMembers/{scopedRoleMembership-id} | Update the navigation property scopedMembers in directoryRoles



## DirectoryRolesCreateScopedMembers

> MicrosoftGraphScopedRoleMembership DirectoryRolesCreateScopedMembers(ctx, directoryRoleId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Create new navigation property to scopedMembers for directoryRoles



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
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesScopedRoleMembershipApi.DirectoryRolesCreateScopedMembers(context.Background(), directoryRoleId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesCreateScopedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesCreateScopedMembers`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesCreateScopedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesCreateScopedMembersRequest struct via the builder pattern


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


## DirectoryRolesDeleteScopedMembers

> DirectoryRolesDeleteScopedMembers(ctx, directoryRoleId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()

Delete navigation property scopedMembers for directoryRoles



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesScopedRoleMembershipApi.DirectoryRolesDeleteScopedMembers(context.Background(), directoryRoleId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesDeleteScopedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDeleteScopedMembersRequest struct via the builder pattern


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


## DirectoryRolesGetScopedMembers

> MicrosoftGraphScopedRoleMembership DirectoryRolesGetScopedMembers(ctx, directoryRoleId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()

Get scopedMembers from directoryRoles



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesScopedRoleMembershipApi.DirectoryRolesGetScopedMembers(context.Background(), directoryRoleId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesGetScopedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesGetScopedMembers`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesGetScopedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesGetScopedMembersRequest struct via the builder pattern


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


## DirectoryRolesListScopedMembers

> CollectionOfScopedRoleMembership DirectoryRolesListScopedMembers(ctx, directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get scopedMembers from directoryRoles



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
    resp, r, err := api_client.DirectoryRolesScopedRoleMembershipApi.DirectoryRolesListScopedMembers(context.Background(), directoryRoleId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesListScopedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesListScopedMembers`: CollectionOfScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesListScopedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesListScopedMembersRequest struct via the builder pattern


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


## DirectoryRolesUpdateScopedMembers

> DirectoryRolesUpdateScopedMembers(ctx, directoryRoleId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Update the navigation property scopedMembers in directoryRoles



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesScopedRoleMembershipApi.DirectoryRolesUpdateScopedMembers(context.Background(), directoryRoleId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesScopedRoleMembershipApi.DirectoryRolesUpdateScopedMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesUpdateScopedMembersRequest struct via the builder pattern


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

