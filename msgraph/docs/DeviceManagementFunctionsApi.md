# \DeviceManagementFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest**](DeviceManagementFunctionsApi.md#DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest) | **Get** /deviceManagement/applePushNotificationCertificate/microsoft.graph.downloadApplePushNotificationCertificateSigningRequest() | Invoke function downloadApplePushNotificationCertificateSigningRequest
[**DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue**](DeviceManagementFunctionsApi.md#DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue) | **Get** /deviceManagement/deviceConfigurations/{deviceConfiguration-id}/microsoft.graph.getOmaSettingPlainTextValue(secretReferenceValueId&#x3D;&#39;{secretReferenceValueId}&#39;) | Invoke function getOmaSettingPlainTextValue
[**DeviceManagementGetEffectivePermissions**](DeviceManagementFunctionsApi.md#DeviceManagementGetEffectivePermissions) | **Get** /deviceManagement/microsoft.graph.getEffectivePermissions(scope&#x3D;&#39;{scope}&#39;) | Invoke function getEffectivePermissions
[**DeviceManagementVerifyWindowsEnrollmentAutoDiscovery**](DeviceManagementFunctionsApi.md#DeviceManagementVerifyWindowsEnrollmentAutoDiscovery) | **Get** /deviceManagement/microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName&#x3D;&#39;{domainName}&#39;) | Invoke function verifyWindowsEnrollmentAutoDiscovery



## DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest

> string DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest(ctx).Execute()

Invoke function downloadApplePushNotificationCertificateSigningRequest



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementFunctionsApi.DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementFunctionsApi.DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementFunctionsApi.DeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequest struct via the builder pattern


### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue

> string DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue(ctx, deviceConfigurationId, secretReferenceValueId).Execute()

Invoke function getOmaSettingPlainTextValue

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
    deviceConfigurationId := "deviceConfigurationId_example" // string | key: id of deviceConfiguration
    secretReferenceValueId := "secretReferenceValueId_example" // string | Usage: secretReferenceValueId={secretReferenceValueId}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementFunctionsApi.DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue(context.Background(), deviceConfigurationId, secretReferenceValueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementFunctionsApi.DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementFunctionsApi.DeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationId** | **string** | key: id of deviceConfiguration | 
**secretReferenceValueId** | **string** | Usage: secretReferenceValueId&#x3D;{secretReferenceValueId} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceConfigurationsDeviceConfigurationGetOmaSettingPlainTextValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetEffectivePermissions

> []*AnyOfmicrosoftGraphRolePermission DeviceManagementGetEffectivePermissions(ctx, scope).Execute()

Invoke function getEffectivePermissions



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
    scope := "scope_example" // string | Usage: scope={scope}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementFunctionsApi.DeviceManagementGetEffectivePermissions(context.Background(), scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementFunctionsApi.DeviceManagementGetEffectivePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetEffectivePermissions`: []*AnyOfmicrosoftGraphRolePermission
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementFunctionsApi.DeviceManagementGetEffectivePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | Usage: scope&#x3D;{scope} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetEffectivePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphRolePermission**](anyOf&lt;microsoft.graph.rolePermission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementVerifyWindowsEnrollmentAutoDiscovery

> bool DeviceManagementVerifyWindowsEnrollmentAutoDiscovery(ctx, domainName).Execute()

Invoke function verifyWindowsEnrollmentAutoDiscovery

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
    domainName := "domainName_example" // string | Usage: domainName={domainName}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementFunctionsApi.DeviceManagementVerifyWindowsEnrollmentAutoDiscovery(context.Background(), domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementFunctionsApi.DeviceManagementVerifyWindowsEnrollmentAutoDiscovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementVerifyWindowsEnrollmentAutoDiscovery`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementFunctionsApi.DeviceManagementVerifyWindowsEnrollmentAutoDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Usage: domainName&#x3D;{domainName} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementVerifyWindowsEnrollmentAutoDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

