# \DeviceAppManagementVppTokenApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementCreateVppTokens) | **Post** /deviceAppManagement/vppTokens | Create new navigation property to vppTokens for deviceAppManagement
[**DeviceAppManagementDeleteVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementDeleteVppTokens) | **Delete** /deviceAppManagement/vppTokens/{vppToken-id} | Delete navigation property vppTokens for deviceAppManagement
[**DeviceAppManagementGetVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementGetVppTokens) | **Get** /deviceAppManagement/vppTokens/{vppToken-id} | Get vppTokens from deviceAppManagement
[**DeviceAppManagementListVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementListVppTokens) | **Get** /deviceAppManagement/vppTokens | Get vppTokens from deviceAppManagement
[**DeviceAppManagementUpdateVppTokens**](DeviceAppManagementVppTokenApi.md#DeviceAppManagementUpdateVppTokens) | **Patch** /deviceAppManagement/vppTokens/{vppToken-id} | Update the navigation property vppTokens in deviceAppManagement



## DeviceAppManagementCreateVppTokens

> MicrosoftGraphVppToken DeviceAppManagementCreateVppTokens(ctx).MicrosoftGraphVppToken(microsoftGraphVppToken).Execute()

Create new navigation property to vppTokens for deviceAppManagement



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
    microsoftGraphVppToken := *openapiclient.NewMicrosoftGraphVppToken() // MicrosoftGraphVppToken | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementVppTokenApi.DeviceAppManagementCreateVppTokens(context.Background()).MicrosoftGraphVppToken(microsoftGraphVppToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementVppTokenApi.DeviceAppManagementCreateVppTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateVppTokens`: MicrosoftGraphVppToken
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementVppTokenApi.DeviceAppManagementCreateVppTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateVppTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphVppToken** | [**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md) | New navigation property | 

### Return type

[**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteVppTokens

> DeviceAppManagementDeleteVppTokens(ctx, vppTokenId).IfMatch(ifMatch).Execute()

Delete navigation property vppTokens for deviceAppManagement



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
    vppTokenId := "vppTokenId_example" // string | key: id of vppToken
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementVppTokenApi.DeviceAppManagementDeleteVppTokens(context.Background(), vppTokenId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementVppTokenApi.DeviceAppManagementDeleteVppTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string** | key: id of vppToken | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteVppTokensRequest struct via the builder pattern


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


## DeviceAppManagementGetVppTokens

> MicrosoftGraphVppToken DeviceAppManagementGetVppTokens(ctx, vppTokenId).Select_(select_).Expand(expand).Execute()

Get vppTokens from deviceAppManagement



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
    vppTokenId := "vppTokenId_example" // string | key: id of vppToken
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementVppTokenApi.DeviceAppManagementGetVppTokens(context.Background(), vppTokenId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementVppTokenApi.DeviceAppManagementGetVppTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetVppTokens`: MicrosoftGraphVppToken
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementVppTokenApi.DeviceAppManagementGetVppTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string** | key: id of vppToken | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetVppTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListVppTokens

> CollectionOfVppToken DeviceAppManagementListVppTokens(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get vppTokens from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementVppTokenApi.DeviceAppManagementListVppTokens(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementVppTokenApi.DeviceAppManagementListVppTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListVppTokens`: CollectionOfVppToken
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementVppTokenApi.DeviceAppManagementListVppTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListVppTokensRequest struct via the builder pattern


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

[**CollectionOfVppToken**](CollectionOfVppToken.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateVppTokens

> DeviceAppManagementUpdateVppTokens(ctx, vppTokenId).MicrosoftGraphVppToken(microsoftGraphVppToken).Execute()

Update the navigation property vppTokens in deviceAppManagement



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
    vppTokenId := "vppTokenId_example" // string | key: id of vppToken
    microsoftGraphVppToken := *openapiclient.NewMicrosoftGraphVppToken() // MicrosoftGraphVppToken | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementVppTokenApi.DeviceAppManagementUpdateVppTokens(context.Background(), vppTokenId).MicrosoftGraphVppToken(microsoftGraphVppToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementVppTokenApi.DeviceAppManagementUpdateVppTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vppTokenId** | **string** | key: id of vppToken | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateVppTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphVppToken** | [**MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md) | New navigation property values | 

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

