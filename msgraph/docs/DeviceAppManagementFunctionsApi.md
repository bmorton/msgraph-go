# \DeviceAppManagementFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration**](DeviceAppManagementFunctionsApi.md#DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration) | **Get** /deviceAppManagement/managedAppRegistrations/microsoft.graph.getUserIdsWithFlaggedAppRegistration() | Invoke function getUserIdsWithFlaggedAppRegistration



## DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration

> []*string DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration(ctx).Execute()

Invoke function getUserIdsWithFlaggedAppRegistration

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
    resp, r, err := api_client.DeviceAppManagementFunctionsApi.DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementFunctionsApi.DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration`: []*string
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementFunctionsApi.DeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementManagedAppRegistrationsGetUserIdsWithFlaggedAppRegistrationRequest struct via the builder pattern


### Return type

**[]*string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

