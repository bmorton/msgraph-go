# \DeviceManagementWindowsInformationProtectionAppLearningSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries) | **Post** /deviceManagement/windowsInformationProtectionAppLearningSummaries | Create new navigation property to windowsInformationProtectionAppLearningSummaries for deviceManagement
[**DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries) | **Delete** /deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary-id} | Delete navigation property windowsInformationProtectionAppLearningSummaries for deviceManagement
[**DeviceManagementGetWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementGetWindowsInformationProtectionAppLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary-id} | Get windowsInformationProtectionAppLearningSummaries from deviceManagement
[**DeviceManagementListWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementListWindowsInformationProtectionAppLearningSummaries) | **Get** /deviceManagement/windowsInformationProtectionAppLearningSummaries | Get windowsInformationProtectionAppLearningSummaries from deviceManagement
[**DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries**](DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.md#DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries) | **Patch** /deviceManagement/windowsInformationProtectionAppLearningSummaries/{windowsInformationProtectionAppLearningSummary-id} | Update the navigation property windowsInformationProtectionAppLearningSummaries in deviceManagement



## DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries

> MicrosoftGraphWindowsInformationProtectionAppLearningSummary DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries(ctx).MicrosoftGraphWindowsInformationProtectionAppLearningSummary(microsoftGraphWindowsInformationProtectionAppLearningSummary).Execute()

Create new navigation property to windowsInformationProtectionAppLearningSummaries for deviceManagement



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
    microsoftGraphWindowsInformationProtectionAppLearningSummary := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionAppLearningSummary() // MicrosoftGraphWindowsInformationProtectionAppLearningSummary | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries(context.Background()).MicrosoftGraphWindowsInformationProtectionAppLearningSummary(microsoftGraphWindowsInformationProtectionAppLearningSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries`: MicrosoftGraphWindowsInformationProtectionAppLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementCreateWindowsInformationProtectionAppLearningSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateWindowsInformationProtectionAppLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWindowsInformationProtectionAppLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md) | New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries

> DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries(ctx, windowsInformationProtectionAppLearningSummaryId).IfMatch(ifMatch).Execute()

Delete navigation property windowsInformationProtectionAppLearningSummaries for deviceManagement



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
    windowsInformationProtectionAppLearningSummaryId := "windowsInformationProtectionAppLearningSummaryId_example" // string | key: id of windowsInformationProtectionAppLearningSummary
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries(context.Background(), windowsInformationProtectionAppLearningSummaryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementDeleteWindowsInformationProtectionAppLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionAppLearningSummaryId** | **string** | key: id of windowsInformationProtectionAppLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteWindowsInformationProtectionAppLearningSummariesRequest struct via the builder pattern


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


## DeviceManagementGetWindowsInformationProtectionAppLearningSummaries

> MicrosoftGraphWindowsInformationProtectionAppLearningSummary DeviceManagementGetWindowsInformationProtectionAppLearningSummaries(ctx, windowsInformationProtectionAppLearningSummaryId).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionAppLearningSummaries from deviceManagement



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
    windowsInformationProtectionAppLearningSummaryId := "windowsInformationProtectionAppLearningSummaryId_example" // string | key: id of windowsInformationProtectionAppLearningSummary
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionAppLearningSummaries(context.Background(), windowsInformationProtectionAppLearningSummaryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionAppLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetWindowsInformationProtectionAppLearningSummaries`: MicrosoftGraphWindowsInformationProtectionAppLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementGetWindowsInformationProtectionAppLearningSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionAppLearningSummaryId** | **string** | key: id of windowsInformationProtectionAppLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetWindowsInformationProtectionAppLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListWindowsInformationProtectionAppLearningSummaries

> CollectionOfWindowsInformationProtectionAppLearningSummary DeviceManagementListWindowsInformationProtectionAppLearningSummaries(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionAppLearningSummaries from deviceManagement



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
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementListWindowsInformationProtectionAppLearningSummaries(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementListWindowsInformationProtectionAppLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListWindowsInformationProtectionAppLearningSummaries`: CollectionOfWindowsInformationProtectionAppLearningSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementListWindowsInformationProtectionAppLearningSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListWindowsInformationProtectionAppLearningSummariesRequest struct via the builder pattern


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

[**CollectionOfWindowsInformationProtectionAppLearningSummary**](CollectionOfWindowsInformationProtectionAppLearningSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries

> DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries(ctx, windowsInformationProtectionAppLearningSummaryId).MicrosoftGraphWindowsInformationProtectionAppLearningSummary(microsoftGraphWindowsInformationProtectionAppLearningSummary).Execute()

Update the navigation property windowsInformationProtectionAppLearningSummaries in deviceManagement



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
    windowsInformationProtectionAppLearningSummaryId := "windowsInformationProtectionAppLearningSummaryId_example" // string | key: id of windowsInformationProtectionAppLearningSummary
    microsoftGraphWindowsInformationProtectionAppLearningSummary := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionAppLearningSummary() // MicrosoftGraphWindowsInformationProtectionAppLearningSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries(context.Background(), windowsInformationProtectionAppLearningSummaryId).MicrosoftGraphWindowsInformationProtectionAppLearningSummary(microsoftGraphWindowsInformationProtectionAppLearningSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsInformationProtectionAppLearningSummaryApi.DeviceManagementUpdateWindowsInformationProtectionAppLearningSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionAppLearningSummaryId** | **string** | key: id of windowsInformationProtectionAppLearningSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateWindowsInformationProtectionAppLearningSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWindowsInformationProtectionAppLearningSummary** | [**MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md) | New navigation property values | 

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

