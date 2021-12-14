# \DeviceManagementDeviceCompliancePolicySettingStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries) | **Post** /deviceManagement/deviceCompliancePolicySettingStateSummaries | Create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement
[**DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries) | **Delete** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id} | Delete navigation property deviceCompliancePolicySettingStateSummaries for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates) | **Post** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates | Create new navigation property to deviceComplianceSettingStates for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates) | **Delete** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates/{deviceComplianceSettingState-id} | Delete navigation property deviceComplianceSettingStates for deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates/{deviceComplianceSettingState-id} | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates | Get deviceComplianceSettingStates from deviceManagement
[**DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates) | **Patch** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id}/deviceComplianceSettingStates/{deviceComplianceSettingState-id} | Update the navigation property deviceComplianceSettingStates in deviceManagement
[**DeviceManagementGetDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementGetDeviceCompliancePolicySettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id} | Get deviceCompliancePolicySettingStateSummaries from deviceManagement
[**DeviceManagementListDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementListDeviceCompliancePolicySettingStateSummaries) | **Get** /deviceManagement/deviceCompliancePolicySettingStateSummaries | Get deviceCompliancePolicySettingStateSummaries from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries**](DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.md#DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries) | **Patch** /deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary-id} | Update the navigation property deviceCompliancePolicySettingStateSummaries in deviceManagement



## DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries

> MicrosoftGraphDeviceCompliancePolicySettingStateSummary DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries(ctx).MicrosoftGraphDeviceCompliancePolicySettingStateSummary(microsoftGraphDeviceCompliancePolicySettingStateSummary).Execute()

Create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement



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
    microsoftGraphDeviceCompliancePolicySettingStateSummary := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicySettingStateSummary() // MicrosoftGraphDeviceCompliancePolicySettingStateSummary | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries(context.Background()).MicrosoftGraphDeviceCompliancePolicySettingStateSummary(microsoftGraphDeviceCompliancePolicySettingStateSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries`: MicrosoftGraphDeviceCompliancePolicySettingStateSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementCreateDeviceCompliancePolicySettingStateSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateDeviceCompliancePolicySettingStateSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceCompliancePolicySettingStateSummary** | [**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries

> DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries(ctx, deviceCompliancePolicySettingStateSummaryId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCompliancePolicySettingStateSummaries for deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries(context.Background(), deviceCompliancePolicySettingStateSummaryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeleteDeviceCompliancePolicySettingStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceCompliancePolicySettingStateSummariesRequest struct via the builder pattern


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


## DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId).MicrosoftGraphDeviceComplianceSettingState(microsoftGraphDeviceComplianceSettingState).Execute()

Create new navigation property to deviceComplianceSettingStates for deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    microsoftGraphDeviceComplianceSettingState := *openapiclient.NewMicrosoftGraphDeviceComplianceSettingState() // MicrosoftGraphDeviceComplianceSettingState | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates(context.Background(), deviceCompliancePolicySettingStateSummaryId).MicrosoftGraphDeviceComplianceSettingState(microsoftGraphDeviceComplianceSettingState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates`: MicrosoftGraphDeviceComplianceSettingState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceCompliancePolicySettingStateSummariesCreateDeviceComplianceSettingStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates

> DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceComplianceSettingStates for deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    deviceComplianceSettingStateId := "deviceComplianceSettingStateId_example" // string | key: id of deviceComplianceSettingState
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates(context.Background(), deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string** | key: id of deviceComplianceSettingState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceCompliancePolicySettingStateSummariesDeleteDeviceComplianceSettingStatesRequest struct via the builder pattern


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


## DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates

> MicrosoftGraphDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).Select_(select_).Expand(expand).Execute()

Get deviceComplianceSettingStates from deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    deviceComplianceSettingStateId := "deviceComplianceSettingStateId_example" // string | key: id of deviceComplianceSettingState
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates(context.Background(), deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates`: MicrosoftGraphDeviceComplianceSettingState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string** | key: id of deviceComplianceSettingState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceCompliancePolicySettingStateSummariesGetDeviceComplianceSettingStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates

> CollectionOfDeviceComplianceSettingState DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceComplianceSettingStates from deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
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
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates(context.Background(), deviceCompliancePolicySettingStateSummaryId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates`: CollectionOfDeviceComplianceSettingState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceCompliancePolicySettingStateSummariesListDeviceComplianceSettingStatesRequest struct via the builder pattern


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

[**CollectionOfDeviceComplianceSettingState**](CollectionOfDeviceComplianceSettingState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates

> DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates(ctx, deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).MicrosoftGraphDeviceComplianceSettingState(microsoftGraphDeviceComplianceSettingState).Execute()

Update the navigation property deviceComplianceSettingStates in deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    deviceComplianceSettingStateId := "deviceComplianceSettingStateId_example" // string | key: id of deviceComplianceSettingState
    microsoftGraphDeviceComplianceSettingState := *openapiclient.NewMicrosoftGraphDeviceComplianceSettingState() // MicrosoftGraphDeviceComplianceSettingState | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates(context.Background(), deviceCompliancePolicySettingStateSummaryId, deviceComplianceSettingStateId).MicrosoftGraphDeviceComplianceSettingState(microsoftGraphDeviceComplianceSettingState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 
**deviceComplianceSettingStateId** | **string** | key: id of deviceComplianceSettingState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceCompliancePolicySettingStateSummariesUpdateDeviceComplianceSettingStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDeviceComplianceSettingState** | [**MicrosoftGraphDeviceComplianceSettingState**](MicrosoftGraphDeviceComplianceSettingState.md) | New navigation property values | 

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


## DeviceManagementGetDeviceCompliancePolicySettingStateSummaries

> MicrosoftGraphDeviceCompliancePolicySettingStateSummary DeviceManagementGetDeviceCompliancePolicySettingStateSummaries(ctx, deviceCompliancePolicySettingStateSummaryId).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicySettingStateSummaries from deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementGetDeviceCompliancePolicySettingStateSummaries(context.Background(), deviceCompliancePolicySettingStateSummaryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementGetDeviceCompliancePolicySettingStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceCompliancePolicySettingStateSummaries`: MicrosoftGraphDeviceCompliancePolicySettingStateSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementGetDeviceCompliancePolicySettingStateSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceCompliancePolicySettingStateSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceCompliancePolicySettingStateSummaries

> CollectionOfDeviceCompliancePolicySettingStateSummary DeviceManagementListDeviceCompliancePolicySettingStateSummaries(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicySettingStateSummaries from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementListDeviceCompliancePolicySettingStateSummaries(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementListDeviceCompliancePolicySettingStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListDeviceCompliancePolicySettingStateSummaries`: CollectionOfDeviceCompliancePolicySettingStateSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementListDeviceCompliancePolicySettingStateSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListDeviceCompliancePolicySettingStateSummariesRequest struct via the builder pattern


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

[**CollectionOfDeviceCompliancePolicySettingStateSummary**](CollectionOfDeviceCompliancePolicySettingStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries

> DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries(ctx, deviceCompliancePolicySettingStateSummaryId).MicrosoftGraphDeviceCompliancePolicySettingStateSummary(microsoftGraphDeviceCompliancePolicySettingStateSummary).Execute()

Update the navigation property deviceCompliancePolicySettingStateSummaries in deviceManagement



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
    deviceCompliancePolicySettingStateSummaryId := "deviceCompliancePolicySettingStateSummaryId_example" // string | key: id of deviceCompliancePolicySettingStateSummary
    microsoftGraphDeviceCompliancePolicySettingStateSummary := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicySettingStateSummary() // MicrosoftGraphDeviceCompliancePolicySettingStateSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries(context.Background(), deviceCompliancePolicySettingStateSummaryId).MicrosoftGraphDeviceCompliancePolicySettingStateSummary(microsoftGraphDeviceCompliancePolicySettingStateSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicySettingStateSummaryApi.DeviceManagementUpdateDeviceCompliancePolicySettingStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCompliancePolicySettingStateSummaryId** | **string** | key: id of deviceCompliancePolicySettingStateSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceCompliancePolicySettingStateSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceCompliancePolicySettingStateSummary** | [**MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md) | New navigation property values | 

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

