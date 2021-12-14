# \DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries) | **Post** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries | Create new navigation property to windowsInformationProtectionNetworkLearningSummaries for deviceManagement
[**DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries) | **Delete** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary-id} | Delete navigation property windowsInformationProtectionNetworkLearningSummaries for deviceManagement
[**DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary-id} | Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement
[**DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries | Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement
[**DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries**](DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.md#DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries) | **Patch** /deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary-id} | Update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement



## DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries

> MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries(ctx).MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary(microsoftGraphWindowsInformationProtectionNetworkLearningSummary).Execute()

Create new navigation property to windowsInformationProtectionNetworkLearningSummaries for deviceManagement



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
    microsoftGraphWindowsInformationProtectionNetworkLearningSummary := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionNetworkLearningSummary() // MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries(context.Background()).MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary(microsoftGraphWindowsInformationProtectionNetworkLearningSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries`: MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionNetworkLearningSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateWindowsInformationProtectionNetworkLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWindowsInformationProtectionNetworkLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md) | New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries

> DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries(ctx, windowsInformationProtectionNetworkLearningSummaryId).IfMatch(ifMatch).Execute()

Delete navigation property windowsInformationProtectionNetworkLearningSummaries for deviceManagement



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
    windowsInformationProtectionNetworkLearningSummaryId := "windowsInformationProtectionNetworkLearningSummaryId_example" // string | key: id of windowsInformationProtectionNetworkLearningSummary
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries(context.Background(), windowsInformationProtectionNetworkLearningSummaryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionNetworkLearningSummaryId** | **string** | key: id of windowsInformationProtectionNetworkLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteWindowsInformationProtectionNetworkLearningSummariesRequest struct via the builder pattern


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


## DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries

> MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries(ctx, windowsInformationProtectionNetworkLearningSummaryId).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement



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
    windowsInformationProtectionNetworkLearningSummaryId := "windowsInformationProtectionNetworkLearningSummaryId_example" // string | key: id of windowsInformationProtectionNetworkLearningSummary
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries(context.Background(), windowsInformationProtectionNetworkLearningSummaryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries`: MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionNetworkLearningSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionNetworkLearningSummaryId** | **string** | key: id of windowsInformationProtectionNetworkLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetWindowsInformationProtectionNetworkLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries

> CollectionOfWindowsInformationProtectionNetworkLearningSummary DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionNetworkLearningSummaries from deviceManagement



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
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries`: CollectionOfWindowsInformationProtectionNetworkLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementListWindowsInformationProtectionNetworkLearningSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListWindowsInformationProtectionNetworkLearningSummariesRequest struct via the builder pattern


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

[**CollectionOfWindowsInformationProtectionNetworkLearningSummary**](CollectionOfWindowsInformationProtectionNetworkLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries

> DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries(ctx, windowsInformationProtectionNetworkLearningSummaryId).MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary(microsoftGraphWindowsInformationProtectionNetworkLearningSummary).Execute()

Update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement



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
    windowsInformationProtectionNetworkLearningSummaryId := "windowsInformationProtectionNetworkLearningSummaryId_example" // string | key: id of windowsInformationProtectionNetworkLearningSummary
    microsoftGraphWindowsInformationProtectionNetworkLearningSummary := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionNetworkLearningSummary() // MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries(context.Background(), windowsInformationProtectionNetworkLearningSummaryId).MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary(microsoftGraphWindowsInformationProtectionNetworkLearningSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionNetworkLearningSummaryApi.DeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionNetworkLearningSummaryId** | **string** | key: id of windowsInformationProtectionNetworkLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateWindowsInformationProtectionNetworkLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWindowsInformationProtectionNetworkLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md) | New navigation property values | 

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

