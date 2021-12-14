# \DeviceManagementApplePushNotificationCertificateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteApplePushNotificationCertificate**](DeviceManagementApplePushNotificationCertificateApi.md#DeviceManagementDeleteApplePushNotificationCertificate) | **Delete** /deviceManagement/applePushNotificationCertificate | Delete navigation property applePushNotificationCertificate for deviceManagement
[**DeviceManagementGetApplePushNotificationCertificate**](DeviceManagementApplePushNotificationCertificateApi.md#DeviceManagementGetApplePushNotificationCertificate) | **Get** /deviceManagement/applePushNotificationCertificate | Get applePushNotificationCertificate from deviceManagement
[**DeviceManagementUpdateApplePushNotificationCertificate**](DeviceManagementApplePushNotificationCertificateApi.md#DeviceManagementUpdateApplePushNotificationCertificate) | **Patch** /deviceManagement/applePushNotificationCertificate | Update the navigation property applePushNotificationCertificate in deviceManagement



## DeviceManagementDeleteApplePushNotificationCertificate

> DeviceManagementDeleteApplePushNotificationCertificate(ctx).IfMatch(ifMatch).Execute()

Delete navigation property applePushNotificationCertificate for deviceManagement



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementApplePushNotificationCertificateApi.DeviceManagementDeleteApplePushNotificationCertificate(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementApplePushNotificationCertificateApi.DeviceManagementDeleteApplePushNotificationCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteApplePushNotificationCertificateRequest struct via the builder pattern


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


## DeviceManagementGetApplePushNotificationCertificate

> MicrosoftGraphApplePushNotificationCertificate DeviceManagementGetApplePushNotificationCertificate(ctx).Select_(select_).Expand(expand).Execute()

Get applePushNotificationCertificate from deviceManagement



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementApplePushNotificationCertificateApi.DeviceManagementGetApplePushNotificationCertificate(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementApplePushNotificationCertificateApi.DeviceManagementGetApplePushNotificationCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetApplePushNotificationCertificate`: MicrosoftGraphApplePushNotificationCertificate
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementApplePushNotificationCertificateApi.DeviceManagementGetApplePushNotificationCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetApplePushNotificationCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphApplePushNotificationCertificate**](MicrosoftGraphApplePushNotificationCertificate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateApplePushNotificationCertificate

> DeviceManagementUpdateApplePushNotificationCertificate(ctx).MicrosoftGraphApplePushNotificationCertificate(microsoftGraphApplePushNotificationCertificate).Execute()

Update the navigation property applePushNotificationCertificate in deviceManagement



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
    microsoftGraphApplePushNotificationCertificate := *openapiclient.NewMicrosoftGraphApplePushNotificationCertificate() // MicrosoftGraphApplePushNotificationCertificate | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementApplePushNotificationCertificateApi.DeviceManagementUpdateApplePushNotificationCertificate(context.Background()).MicrosoftGraphApplePushNotificationCertificate(microsoftGraphApplePushNotificationCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementApplePushNotificationCertificateApi.DeviceManagementUpdateApplePushNotificationCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateApplePushNotificationCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphApplePushNotificationCertificate** | [**MicrosoftGraphApplePushNotificationCertificate**](MicrosoftGraphApplePushNotificationCertificate.md) | New navigation property values | 

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

