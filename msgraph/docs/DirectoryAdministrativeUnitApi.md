# \DirectoryAdministrativeUnitApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryAdministrativeUnitsCreateExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsCreateExtensions) | **Post** /directory/administrativeUnits/{administrativeUnit-id}/extensions | Create new navigation property to extensions for directory
[**DirectoryAdministrativeUnitsCreateRefMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsCreateRefMembers) | **Post** /directory/administrativeUnits/{administrativeUnit-id}/members/$ref | Create new navigation property ref to members for directory
[**DirectoryAdministrativeUnitsCreateScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsCreateScopedRoleMembers) | **Post** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers | Create new navigation property to scopedRoleMembers for directory
[**DirectoryAdministrativeUnitsDeleteExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsDeleteExtensions) | **Delete** /directory/administrativeUnits/{administrativeUnit-id}/extensions/{extension-id} | Delete navigation property extensions for directory
[**DirectoryAdministrativeUnitsDeleteScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsDeleteScopedRoleMembers) | **Delete** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers/{scopedRoleMembership-id} | Delete navigation property scopedRoleMembers for directory
[**DirectoryAdministrativeUnitsGetExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsGetExtensions) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/extensions/{extension-id} | Get extensions from directory
[**DirectoryAdministrativeUnitsGetScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsGetScopedRoleMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers/{scopedRoleMembership-id} | Get scopedRoleMembers from directory
[**DirectoryAdministrativeUnitsListExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListExtensions) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/extensions | Get extensions from directory
[**DirectoryAdministrativeUnitsListMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/members | Get members from directory
[**DirectoryAdministrativeUnitsListRefMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListRefMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/members/$ref | Get ref of members from directory
[**DirectoryAdministrativeUnitsListScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsListScopedRoleMembers) | **Get** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers | Get scopedRoleMembers from directory
[**DirectoryAdministrativeUnitsUpdateExtensions**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsUpdateExtensions) | **Patch** /directory/administrativeUnits/{administrativeUnit-id}/extensions/{extension-id} | Update the navigation property extensions in directory
[**DirectoryAdministrativeUnitsUpdateScopedRoleMembers**](DirectoryAdministrativeUnitApi.md#DirectoryAdministrativeUnitsUpdateScopedRoleMembers) | **Patch** /directory/administrativeUnits/{administrativeUnit-id}/scopedRoleMembers/{scopedRoleMembership-id} | Update the navigation property scopedRoleMembers in directory
[**DirectoryCreateAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryCreateAdministrativeUnits) | **Post** /directory/administrativeUnits | Create new navigation property to administrativeUnits for directory
[**DirectoryDeleteAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryDeleteAdministrativeUnits) | **Delete** /directory/administrativeUnits/{administrativeUnit-id} | Delete navigation property administrativeUnits for directory
[**DirectoryGetAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryGetAdministrativeUnits) | **Get** /directory/administrativeUnits/{administrativeUnit-id} | Get administrativeUnits from directory
[**DirectoryListAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryListAdministrativeUnits) | **Get** /directory/administrativeUnits | Get administrativeUnits from directory
[**DirectoryUpdateAdministrativeUnits**](DirectoryAdministrativeUnitApi.md#DirectoryUpdateAdministrativeUnits) | **Patch** /directory/administrativeUnits/{administrativeUnit-id} | Update the navigation property administrativeUnits in directory



## DirectoryAdministrativeUnitsCreateExtensions

> MicrosoftGraphExtension DirectoryAdministrativeUnitsCreateExtensions(ctx, administrativeUnitId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateExtensions(context.Background(), administrativeUnitId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsCreateExtensionsRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsCreateRefMembers

> map[string]interface{} DirectoryAdministrativeUnitsCreateRefMembers(ctx, administrativeUnitId).RequestBody(requestBody).Execute()

Create new navigation property ref to members for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateRefMembers(context.Background(), administrativeUnitId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateRefMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsCreateRefMembers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateRefMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsCreateRefMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsCreateScopedRoleMembers

> MicrosoftGraphScopedRoleMembership DirectoryAdministrativeUnitsCreateScopedRoleMembers(ctx, administrativeUnitId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Create new navigation property to scopedRoleMembers for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateScopedRoleMembers(context.Background(), administrativeUnitId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateScopedRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsCreateScopedRoleMembers`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsCreateScopedRoleMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsCreateScopedRoleMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsDeleteExtensions

> DirectoryAdministrativeUnitsDeleteExtensions(ctx, administrativeUnitId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsDeleteExtensions(context.Background(), administrativeUnitId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsDeleteExtensionsRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsDeleteScopedRoleMembers

> DirectoryAdministrativeUnitsDeleteScopedRoleMembers(ctx, administrativeUnitId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()

Delete navigation property scopedRoleMembers for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsDeleteScopedRoleMembers(context.Background(), administrativeUnitId, scopedRoleMembershipId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsDeleteScopedRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsDeleteScopedRoleMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsGetExtensions

> MicrosoftGraphExtension DirectoryAdministrativeUnitsGetExtensions(ctx, administrativeUnitId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetExtensions(context.Background(), administrativeUnitId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsGetExtensionsRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsGetScopedRoleMembers

> MicrosoftGraphScopedRoleMembership DirectoryAdministrativeUnitsGetScopedRoleMembers(ctx, administrativeUnitId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()

Get scopedRoleMembers from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetScopedRoleMembers(context.Background(), administrativeUnitId, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetScopedRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsGetScopedRoleMembers`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsGetScopedRoleMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsGetScopedRoleMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsListExtensions

> CollectionOfExtension DirectoryAdministrativeUnitsListExtensions(ctx, administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
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
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListExtensions(context.Background(), administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsListExtensionsRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsListMembers

> CollectionOfDirectoryObject DirectoryAdministrativeUnitsListMembers(ctx, administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
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
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListMembers(context.Background(), administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsListMembers`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsListMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsListRefMembers

> CollectionOfLinksOfDirectoryObject DirectoryAdministrativeUnitsListRefMembers(ctx, administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of members from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListRefMembers(context.Background(), administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListRefMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsListRefMembers`: CollectionOfLinksOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListRefMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsListRefMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsListScopedRoleMembers

> CollectionOfScopedRoleMembership DirectoryAdministrativeUnitsListScopedRoleMembers(ctx, administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get scopedRoleMembers from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
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
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListScopedRoleMembers(context.Background(), administrativeUnitId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListScopedRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsListScopedRoleMembers`: CollectionOfScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsListScopedRoleMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsListScopedRoleMembersRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsUpdateExtensions

> DirectoryAdministrativeUnitsUpdateExtensions(ctx, administrativeUnitId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsUpdateExtensions(context.Background(), administrativeUnitId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsUpdateExtensionsRequest struct via the builder pattern


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


## DirectoryAdministrativeUnitsUpdateScopedRoleMembers

> DirectoryAdministrativeUnitsUpdateScopedRoleMembers(ctx, administrativeUnitId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Update the navigation property scopedRoleMembers in directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsUpdateScopedRoleMembers(context.Background(), administrativeUnitId, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryAdministrativeUnitsUpdateScopedRoleMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsUpdateScopedRoleMembersRequest struct via the builder pattern


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


## DirectoryCreateAdministrativeUnits

> MicrosoftGraphAdministrativeUnit DirectoryCreateAdministrativeUnits(ctx).MicrosoftGraphAdministrativeUnit(microsoftGraphAdministrativeUnit).Execute()

Create new navigation property to administrativeUnits for directory



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
    microsoftGraphAdministrativeUnit := *openapiclient.NewMicrosoftGraphAdministrativeUnit() // MicrosoftGraphAdministrativeUnit | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryCreateAdministrativeUnits(context.Background()).MicrosoftGraphAdministrativeUnit(microsoftGraphAdministrativeUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryCreateAdministrativeUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryCreateAdministrativeUnits`: MicrosoftGraphAdministrativeUnit
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryCreateAdministrativeUnits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryCreateAdministrativeUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAdministrativeUnit** | [**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md) | New navigation property | 

### Return type

[**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryDeleteAdministrativeUnits

> DirectoryDeleteAdministrativeUnits(ctx, administrativeUnitId).IfMatch(ifMatch).Execute()

Delete navigation property administrativeUnits for directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryDeleteAdministrativeUnits(context.Background(), administrativeUnitId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryDeleteAdministrativeUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryDeleteAdministrativeUnitsRequest struct via the builder pattern


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


## DirectoryGetAdministrativeUnits

> MicrosoftGraphAdministrativeUnit DirectoryGetAdministrativeUnits(ctx, administrativeUnitId).Select_(select_).Expand(expand).Execute()

Get administrativeUnits from directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryGetAdministrativeUnits(context.Background(), administrativeUnitId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryGetAdministrativeUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryGetAdministrativeUnits`: MicrosoftGraphAdministrativeUnit
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryGetAdministrativeUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryGetAdministrativeUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryListAdministrativeUnits

> CollectionOfAdministrativeUnit DirectoryListAdministrativeUnits(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get administrativeUnits from directory



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
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryListAdministrativeUnits(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryListAdministrativeUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryListAdministrativeUnits`: CollectionOfAdministrativeUnit
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAdministrativeUnitApi.DirectoryListAdministrativeUnits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryListAdministrativeUnitsRequest struct via the builder pattern


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

[**CollectionOfAdministrativeUnit**](CollectionOfAdministrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryUpdateAdministrativeUnits

> DirectoryUpdateAdministrativeUnits(ctx, administrativeUnitId).MicrosoftGraphAdministrativeUnit(microsoftGraphAdministrativeUnit).Execute()

Update the navigation property administrativeUnits in directory



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
    administrativeUnitId := "administrativeUnitId_example" // string | key: id of administrativeUnit
    microsoftGraphAdministrativeUnit := *openapiclient.NewMicrosoftGraphAdministrativeUnit() // MicrosoftGraphAdministrativeUnit | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryAdministrativeUnitApi.DirectoryUpdateAdministrativeUnits(context.Background(), administrativeUnitId).MicrosoftGraphAdministrativeUnit(microsoftGraphAdministrativeUnit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAdministrativeUnitApi.DirectoryUpdateAdministrativeUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administrativeUnitId** | **string** | key: id of administrativeUnit | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryUpdateAdministrativeUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAdministrativeUnit** | [**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md) | New navigation property values | 

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

