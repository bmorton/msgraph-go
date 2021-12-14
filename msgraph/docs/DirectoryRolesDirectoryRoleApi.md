# \DirectoryRolesDirectoryRoleApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesDirectoryRoleCreateDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleCreateDirectoryRole) | **Post** /directoryRoles | Add new entity to directoryRoles
[**DirectoryRolesDirectoryRoleDeleteDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleDeleteDirectoryRole) | **Delete** /directoryRoles/{directoryRole-id} | Delete entity from directoryRoles
[**DirectoryRolesDirectoryRoleGetDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleGetDirectoryRole) | **Get** /directoryRoles/{directoryRole-id} | Get entity from directoryRoles by key
[**DirectoryRolesDirectoryRoleListDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleListDirectoryRole) | **Get** /directoryRoles | Get entities from directoryRoles
[**DirectoryRolesDirectoryRoleUpdateDirectoryRole**](DirectoryRolesDirectoryRoleApi.md#DirectoryRolesDirectoryRoleUpdateDirectoryRole) | **Patch** /directoryRoles/{directoryRole-id} | Update entity in directoryRoles



## DirectoryRolesDirectoryRoleCreateDirectoryRole

> MicrosoftGraphDirectoryRole DirectoryRolesDirectoryRoleCreateDirectoryRole(ctx).MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole).Execute()

Add new entity to directoryRoles

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
    microsoftGraphDirectoryRole := *openapiclient.NewMicrosoftGraphDirectoryRole() // MicrosoftGraphDirectoryRole | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleCreateDirectoryRole(context.Background()).MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleCreateDirectoryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleCreateDirectoryRole`: MicrosoftGraphDirectoryRole
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleCreateDirectoryRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectoryRole** | [**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md) | New entity | 

### Return type

[**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleDeleteDirectoryRole

> DirectoryRolesDirectoryRoleDeleteDirectoryRole(ctx, directoryRoleId).IfMatch(ifMatch).Execute()

Delete entity from directoryRoles

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleDeleteDirectoryRole(context.Background(), directoryRoleId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleDeleteDirectoryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest struct via the builder pattern


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


## DirectoryRolesDirectoryRoleGetDirectoryRole

> MicrosoftGraphDirectoryRole DirectoryRolesDirectoryRoleGetDirectoryRole(ctx, directoryRoleId).Select_(select_).Expand(expand).Execute()

Get entity from directoryRoles by key

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleGetDirectoryRole(context.Background(), directoryRoleId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleGetDirectoryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleGetDirectoryRole`: MicrosoftGraphDirectoryRole
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleGetDirectoryRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleListDirectoryRole

> CollectionOfDirectoryRole DirectoryRolesDirectoryRoleListDirectoryRole(ctx).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from directoryRoles

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
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleListDirectoryRole(context.Background()).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleListDirectoryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleListDirectoryRole`: CollectionOfDirectoryRole
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleListDirectoryRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleListDirectoryRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDirectoryRole**](CollectionOfDirectoryRole.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRolesDirectoryRoleUpdateDirectoryRole

> DirectoryRolesDirectoryRoleUpdateDirectoryRole(ctx, directoryRoleId).MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole).Execute()

Update entity in directoryRoles

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
    microsoftGraphDirectoryRole := *openapiclient.NewMicrosoftGraphDirectoryRole() // MicrosoftGraphDirectoryRole | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleUpdateDirectoryRole(context.Background(), directoryRoleId).MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesDirectoryRoleApi.DirectoryRolesDirectoryRoleUpdateDirectoryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryRole** | [**MicrosoftGraphDirectoryRole**](MicrosoftGraphDirectoryRole.md) | New property values | 

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

