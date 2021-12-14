# \DeviceManagementDeviceManagementExchangeConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementCreateExchangeConnectors) | **Post** /deviceManagement/exchangeConnectors | Create new navigation property to exchangeConnectors for deviceManagement
[**DeviceManagementDeleteExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementDeleteExchangeConnectors) | **Delete** /deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector-id} | Delete navigation property exchangeConnectors for deviceManagement
[**DeviceManagementGetExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementGetExchangeConnectors) | **Get** /deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector-id} | Get exchangeConnectors from deviceManagement
[**DeviceManagementListExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementListExchangeConnectors) | **Get** /deviceManagement/exchangeConnectors | Get exchangeConnectors from deviceManagement
[**DeviceManagementUpdateExchangeConnectors**](DeviceManagementDeviceManagementExchangeConnectorApi.md#DeviceManagementUpdateExchangeConnectors) | **Patch** /deviceManagement/exchangeConnectors/{deviceManagementExchangeConnector-id} | Update the navigation property exchangeConnectors in deviceManagement



## DeviceManagementCreateExchangeConnectors

> MicrosoftGraphDeviceManagementExchangeConnector DeviceManagementCreateExchangeConnectors(ctx).MicrosoftGraphDeviceManagementExchangeConnector(microsoftGraphDeviceManagementExchangeConnector).Execute()

Create new navigation property to exchangeConnectors for deviceManagement



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
    microsoftGraphDeviceManagementExchangeConnector := *openapiclient.NewMicrosoftGraphDeviceManagementExchangeConnector() // MicrosoftGraphDeviceManagementExchangeConnector | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementCreateExchangeConnectors(context.Background()).MicrosoftGraphDeviceManagementExchangeConnector(microsoftGraphDeviceManagementExchangeConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementCreateExchangeConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateExchangeConnectors`: MicrosoftGraphDeviceManagementExchangeConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementCreateExchangeConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateExchangeConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagementExchangeConnector** | [**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteExchangeConnectors

> DeviceManagementDeleteExchangeConnectors(ctx, deviceManagementExchangeConnectorId).IfMatch(ifMatch).Execute()

Delete navigation property exchangeConnectors for deviceManagement



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
    deviceManagementExchangeConnectorId := "deviceManagementExchangeConnectorId_example" // string | key: id of deviceManagementExchangeConnector
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementDeleteExchangeConnectors(context.Background(), deviceManagementExchangeConnectorId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementDeleteExchangeConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExchangeConnectorId** | **string** | key: id of deviceManagementExchangeConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteExchangeConnectorsRequest struct via the builder pattern


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


## DeviceManagementGetExchangeConnectors

> MicrosoftGraphDeviceManagementExchangeConnector DeviceManagementGetExchangeConnectors(ctx, deviceManagementExchangeConnectorId).Select_(select_).Expand(expand).Execute()

Get exchangeConnectors from deviceManagement



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
    deviceManagementExchangeConnectorId := "deviceManagementExchangeConnectorId_example" // string | key: id of deviceManagementExchangeConnector
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementGetExchangeConnectors(context.Background(), deviceManagementExchangeConnectorId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementGetExchangeConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetExchangeConnectors`: MicrosoftGraphDeviceManagementExchangeConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementGetExchangeConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExchangeConnectorId** | **string** | key: id of deviceManagementExchangeConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetExchangeConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListExchangeConnectors

> CollectionOfDeviceManagementExchangeConnector DeviceManagementListExchangeConnectors(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get exchangeConnectors from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementListExchangeConnectors(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementListExchangeConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListExchangeConnectors`: CollectionOfDeviceManagementExchangeConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementListExchangeConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListExchangeConnectorsRequest struct via the builder pattern


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

[**CollectionOfDeviceManagementExchangeConnector**](CollectionOfDeviceManagementExchangeConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateExchangeConnectors

> DeviceManagementUpdateExchangeConnectors(ctx, deviceManagementExchangeConnectorId).MicrosoftGraphDeviceManagementExchangeConnector(microsoftGraphDeviceManagementExchangeConnector).Execute()

Update the navigation property exchangeConnectors in deviceManagement



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
    deviceManagementExchangeConnectorId := "deviceManagementExchangeConnectorId_example" // string | key: id of deviceManagementExchangeConnector
    microsoftGraphDeviceManagementExchangeConnector := *openapiclient.NewMicrosoftGraphDeviceManagementExchangeConnector() // MicrosoftGraphDeviceManagementExchangeConnector | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementUpdateExchangeConnectors(context.Background(), deviceManagementExchangeConnectorId).MicrosoftGraphDeviceManagementExchangeConnector(microsoftGraphDeviceManagementExchangeConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementExchangeConnectorApi.DeviceManagementUpdateExchangeConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExchangeConnectorId** | **string** | key: id of deviceManagementExchangeConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateExchangeConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceManagementExchangeConnector** | [**MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md) | New navigation property values | 

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

