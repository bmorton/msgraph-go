# \DeviceManagementImportedWindowsAutopilotDeviceIdentityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities**](DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.md#DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities) | **Post** /deviceManagement/importedWindowsAutopilotDeviceIdentities | Create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
[**DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities**](DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.md#DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities) | **Delete** /deviceManagement/importedWindowsAutopilotDeviceIdentities/{importedWindowsAutopilotDeviceIdentity-id} | Delete navigation property importedWindowsAutopilotDeviceIdentities for deviceManagement
[**DeviceManagementGetImportedWindowsAutopilotDeviceIdentities**](DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.md#DeviceManagementGetImportedWindowsAutopilotDeviceIdentities) | **Get** /deviceManagement/importedWindowsAutopilotDeviceIdentities/{importedWindowsAutopilotDeviceIdentity-id} | Get importedWindowsAutopilotDeviceIdentities from deviceManagement
[**DeviceManagementListImportedWindowsAutopilotDeviceIdentities**](DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.md#DeviceManagementListImportedWindowsAutopilotDeviceIdentities) | **Get** /deviceManagement/importedWindowsAutopilotDeviceIdentities | Get importedWindowsAutopilotDeviceIdentities from deviceManagement
[**DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities**](DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.md#DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities) | **Patch** /deviceManagement/importedWindowsAutopilotDeviceIdentities/{importedWindowsAutopilotDeviceIdentity-id} | Update the navigation property importedWindowsAutopilotDeviceIdentities in deviceManagement



## DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities

> MicrosoftGraphImportedWindowsAutopilotDeviceIdentity DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities(ctx).MicrosoftGraphImportedWindowsAutopilotDeviceIdentity(microsoftGraphImportedWindowsAutopilotDeviceIdentity).Execute()

Create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement



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
    microsoftGraphImportedWindowsAutopilotDeviceIdentity := *openapiclient.NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentity() // MicrosoftGraphImportedWindowsAutopilotDeviceIdentity | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities(context.Background()).MicrosoftGraphImportedWindowsAutopilotDeviceIdentity(microsoftGraphImportedWindowsAutopilotDeviceIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities`: MicrosoftGraphImportedWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateImportedWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateImportedWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphImportedWindowsAutopilotDeviceIdentity** | [**MicrosoftGraphImportedWindowsAutopilotDeviceIdentity**](MicrosoftGraphImportedWindowsAutopilotDeviceIdentity.md) | New navigation property | 

### Return type

[**MicrosoftGraphImportedWindowsAutopilotDeviceIdentity**](MicrosoftGraphImportedWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities

> DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities(ctx, importedWindowsAutopilotDeviceIdentityId).IfMatch(ifMatch).Execute()

Delete navigation property importedWindowsAutopilotDeviceIdentities for deviceManagement



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
    importedWindowsAutopilotDeviceIdentityId := "importedWindowsAutopilotDeviceIdentityId_example" // string | key: id of importedWindowsAutopilotDeviceIdentity
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities(context.Background(), importedWindowsAutopilotDeviceIdentityId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementDeleteImportedWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importedWindowsAutopilotDeviceIdentityId** | **string** | key: id of importedWindowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteImportedWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


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


## DeviceManagementGetImportedWindowsAutopilotDeviceIdentities

> MicrosoftGraphImportedWindowsAutopilotDeviceIdentity DeviceManagementGetImportedWindowsAutopilotDeviceIdentities(ctx, importedWindowsAutopilotDeviceIdentityId).Select_(select_).Expand(expand).Execute()

Get importedWindowsAutopilotDeviceIdentities from deviceManagement



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
    importedWindowsAutopilotDeviceIdentityId := "importedWindowsAutopilotDeviceIdentityId_example" // string | key: id of importedWindowsAutopilotDeviceIdentity
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementGetImportedWindowsAutopilotDeviceIdentities(context.Background(), importedWindowsAutopilotDeviceIdentityId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementGetImportedWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetImportedWindowsAutopilotDeviceIdentities`: MicrosoftGraphImportedWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementGetImportedWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importedWindowsAutopilotDeviceIdentityId** | **string** | key: id of importedWindowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetImportedWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphImportedWindowsAutopilotDeviceIdentity**](MicrosoftGraphImportedWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListImportedWindowsAutopilotDeviceIdentities

> CollectionOfImportedWindowsAutopilotDeviceIdentity DeviceManagementListImportedWindowsAutopilotDeviceIdentities(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get importedWindowsAutopilotDeviceIdentities from deviceManagement



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
    resp, r, err := api_client.DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementListImportedWindowsAutopilotDeviceIdentities(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementListImportedWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListImportedWindowsAutopilotDeviceIdentities`: CollectionOfImportedWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementListImportedWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListImportedWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


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

[**CollectionOfImportedWindowsAutopilotDeviceIdentity**](CollectionOfImportedWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities

> DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities(ctx, importedWindowsAutopilotDeviceIdentityId).MicrosoftGraphImportedWindowsAutopilotDeviceIdentity(microsoftGraphImportedWindowsAutopilotDeviceIdentity).Execute()

Update the navigation property importedWindowsAutopilotDeviceIdentities in deviceManagement



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
    importedWindowsAutopilotDeviceIdentityId := "importedWindowsAutopilotDeviceIdentityId_example" // string | key: id of importedWindowsAutopilotDeviceIdentity
    microsoftGraphImportedWindowsAutopilotDeviceIdentity := *openapiclient.NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentity() // MicrosoftGraphImportedWindowsAutopilotDeviceIdentity | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities(context.Background(), importedWindowsAutopilotDeviceIdentityId).MicrosoftGraphImportedWindowsAutopilotDeviceIdentity(microsoftGraphImportedWindowsAutopilotDeviceIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementImportedWindowsAutopilotDeviceIdentityApi.DeviceManagementUpdateImportedWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importedWindowsAutopilotDeviceIdentityId** | **string** | key: id of importedWindowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateImportedWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphImportedWindowsAutopilotDeviceIdentity** | [**MicrosoftGraphImportedWindowsAutopilotDeviceIdentity**](MicrosoftGraphImportedWindowsAutopilotDeviceIdentity.md) | New navigation property values | 

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

