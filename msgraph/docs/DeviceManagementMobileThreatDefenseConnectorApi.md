# \DeviceManagementMobileThreatDefenseConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementCreateMobileThreatDefenseConnectors) | **Post** /deviceManagement/mobileThreatDefenseConnectors | Create new navigation property to mobileThreatDefenseConnectors for deviceManagement
[**DeviceManagementDeleteMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementDeleteMobileThreatDefenseConnectors) | **Delete** /deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id} | Delete navigation property mobileThreatDefenseConnectors for deviceManagement
[**DeviceManagementGetMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementGetMobileThreatDefenseConnectors) | **Get** /deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id} | Get mobileThreatDefenseConnectors from deviceManagement
[**DeviceManagementListMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementListMobileThreatDefenseConnectors) | **Get** /deviceManagement/mobileThreatDefenseConnectors | Get mobileThreatDefenseConnectors from deviceManagement
[**DeviceManagementUpdateMobileThreatDefenseConnectors**](DeviceManagementMobileThreatDefenseConnectorApi.md#DeviceManagementUpdateMobileThreatDefenseConnectors) | **Patch** /deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id} | Update the navigation property mobileThreatDefenseConnectors in deviceManagement



## DeviceManagementCreateMobileThreatDefenseConnectors

> MicrosoftGraphMobileThreatDefenseConnector DeviceManagementCreateMobileThreatDefenseConnectors(ctx).MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector).Execute()

Create new navigation property to mobileThreatDefenseConnectors for deviceManagement



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
    microsoftGraphMobileThreatDefenseConnector := *openapiclient.NewMicrosoftGraphMobileThreatDefenseConnector() // MicrosoftGraphMobileThreatDefenseConnector | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementCreateMobileThreatDefenseConnectors(context.Background()).MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementCreateMobileThreatDefenseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateMobileThreatDefenseConnectors`: MicrosoftGraphMobileThreatDefenseConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementCreateMobileThreatDefenseConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateMobileThreatDefenseConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMobileThreatDefenseConnector** | [**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md) | New navigation property | 

### Return type

[**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteMobileThreatDefenseConnectors

> DeviceManagementDeleteMobileThreatDefenseConnectors(ctx, mobileThreatDefenseConnectorId).IfMatch(ifMatch).Execute()

Delete navigation property mobileThreatDefenseConnectors for deviceManagement



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
    mobileThreatDefenseConnectorId := "mobileThreatDefenseConnectorId_example" // string | key: id of mobileThreatDefenseConnector
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementDeleteMobileThreatDefenseConnectors(context.Background(), mobileThreatDefenseConnectorId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementDeleteMobileThreatDefenseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileThreatDefenseConnectorId** | **string** | key: id of mobileThreatDefenseConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest struct via the builder pattern


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


## DeviceManagementGetMobileThreatDefenseConnectors

> MicrosoftGraphMobileThreatDefenseConnector DeviceManagementGetMobileThreatDefenseConnectors(ctx, mobileThreatDefenseConnectorId).Select_(select_).Expand(expand).Execute()

Get mobileThreatDefenseConnectors from deviceManagement



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
    mobileThreatDefenseConnectorId := "mobileThreatDefenseConnectorId_example" // string | key: id of mobileThreatDefenseConnector
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementGetMobileThreatDefenseConnectors(context.Background(), mobileThreatDefenseConnectorId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementGetMobileThreatDefenseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetMobileThreatDefenseConnectors`: MicrosoftGraphMobileThreatDefenseConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementGetMobileThreatDefenseConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileThreatDefenseConnectorId** | **string** | key: id of mobileThreatDefenseConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetMobileThreatDefenseConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListMobileThreatDefenseConnectors

> CollectionOfMobileThreatDefenseConnector DeviceManagementListMobileThreatDefenseConnectors(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get mobileThreatDefenseConnectors from deviceManagement



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
    resp, r, err := api_client.DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementListMobileThreatDefenseConnectors(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementListMobileThreatDefenseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListMobileThreatDefenseConnectors`: CollectionOfMobileThreatDefenseConnector
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementListMobileThreatDefenseConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListMobileThreatDefenseConnectorsRequest struct via the builder pattern


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

[**CollectionOfMobileThreatDefenseConnector**](CollectionOfMobileThreatDefenseConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateMobileThreatDefenseConnectors

> DeviceManagementUpdateMobileThreatDefenseConnectors(ctx, mobileThreatDefenseConnectorId).MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector).Execute()

Update the navigation property mobileThreatDefenseConnectors in deviceManagement



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
    mobileThreatDefenseConnectorId := "mobileThreatDefenseConnectorId_example" // string | key: id of mobileThreatDefenseConnector
    microsoftGraphMobileThreatDefenseConnector := *openapiclient.NewMicrosoftGraphMobileThreatDefenseConnector() // MicrosoftGraphMobileThreatDefenseConnector | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementUpdateMobileThreatDefenseConnectors(context.Background(), mobileThreatDefenseConnectorId).MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementMobileThreatDefenseConnectorApi.DeviceManagementUpdateMobileThreatDefenseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileThreatDefenseConnectorId** | **string** | key: id of mobileThreatDefenseConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMobileThreatDefenseConnector** | [**MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md) | New navigation property values | 

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

