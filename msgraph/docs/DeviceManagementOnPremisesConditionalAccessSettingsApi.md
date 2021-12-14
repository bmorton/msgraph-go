# \DeviceManagementOnPremisesConditionalAccessSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteConditionalAccessSettings**](DeviceManagementOnPremisesConditionalAccessSettingsApi.md#DeviceManagementDeleteConditionalAccessSettings) | **Delete** /deviceManagement/conditionalAccessSettings | Delete navigation property conditionalAccessSettings for deviceManagement
[**DeviceManagementGetConditionalAccessSettings**](DeviceManagementOnPremisesConditionalAccessSettingsApi.md#DeviceManagementGetConditionalAccessSettings) | **Get** /deviceManagement/conditionalAccessSettings | Get conditionalAccessSettings from deviceManagement
[**DeviceManagementUpdateConditionalAccessSettings**](DeviceManagementOnPremisesConditionalAccessSettingsApi.md#DeviceManagementUpdateConditionalAccessSettings) | **Patch** /deviceManagement/conditionalAccessSettings | Update the navigation property conditionalAccessSettings in deviceManagement



## DeviceManagementDeleteConditionalAccessSettings

> DeviceManagementDeleteConditionalAccessSettings(ctx).IfMatch(ifMatch).Execute()

Delete navigation property conditionalAccessSettings for deviceManagement



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
    resp, r, err := api_client.DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementDeleteConditionalAccessSettings(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementDeleteConditionalAccessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteConditionalAccessSettingsRequest struct via the builder pattern


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


## DeviceManagementGetConditionalAccessSettings

> MicrosoftGraphOnPremisesConditionalAccessSettings DeviceManagementGetConditionalAccessSettings(ctx).Select_(select_).Expand(expand).Execute()

Get conditionalAccessSettings from deviceManagement



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
    resp, r, err := api_client.DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementGetConditionalAccessSettings(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementGetConditionalAccessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetConditionalAccessSettings`: MicrosoftGraphOnPremisesConditionalAccessSettings
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementGetConditionalAccessSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetConditionalAccessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOnPremisesConditionalAccessSettings**](MicrosoftGraphOnPremisesConditionalAccessSettings.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateConditionalAccessSettings

> DeviceManagementUpdateConditionalAccessSettings(ctx).MicrosoftGraphOnPremisesConditionalAccessSettings(microsoftGraphOnPremisesConditionalAccessSettings).Execute()

Update the navigation property conditionalAccessSettings in deviceManagement



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
    microsoftGraphOnPremisesConditionalAccessSettings := *openapiclient.NewMicrosoftGraphOnPremisesConditionalAccessSettings() // MicrosoftGraphOnPremisesConditionalAccessSettings | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementUpdateConditionalAccessSettings(context.Background()).MicrosoftGraphOnPremisesConditionalAccessSettings(microsoftGraphOnPremisesConditionalAccessSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementOnPremisesConditionalAccessSettingsApi.DeviceManagementUpdateConditionalAccessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateConditionalAccessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOnPremisesConditionalAccessSettings** | [**MicrosoftGraphOnPremisesConditionalAccessSettings**](MicrosoftGraphOnPremisesConditionalAccessSettings.md) | New navigation property values | 

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

