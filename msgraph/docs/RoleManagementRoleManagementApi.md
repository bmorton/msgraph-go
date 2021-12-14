# \RoleManagementRoleManagementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleManagementRoleManagementGetRoleManagement**](RoleManagementRoleManagementApi.md#RoleManagementRoleManagementGetRoleManagement) | **Get** /roleManagement | Get roleManagement
[**RoleManagementRoleManagementUpdateRoleManagement**](RoleManagementRoleManagementApi.md#RoleManagementRoleManagementUpdateRoleManagement) | **Patch** /roleManagement | Update roleManagement



## RoleManagementRoleManagementGetRoleManagement

> MicrosoftGraphRoleManagement RoleManagementRoleManagementGetRoleManagement(ctx).Select_(select_).Expand(expand).Execute()

Get roleManagement

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
    resp, r, err := api_client.RoleManagementRoleManagementApi.RoleManagementRoleManagementGetRoleManagement(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementRoleManagementApi.RoleManagementRoleManagementGetRoleManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoleManagementRoleManagementGetRoleManagement`: MicrosoftGraphRoleManagement
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementRoleManagementApi.RoleManagementRoleManagementGetRoleManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleManagementRoleManagementGetRoleManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRoleManagement**](MicrosoftGraphRoleManagement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleManagementRoleManagementUpdateRoleManagement

> RoleManagementRoleManagementUpdateRoleManagement(ctx).MicrosoftGraphRoleManagement(microsoftGraphRoleManagement).Execute()

Update roleManagement

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
    microsoftGraphRoleManagement := *openapiclient.NewMicrosoftGraphRoleManagement() // MicrosoftGraphRoleManagement | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleManagementRoleManagementApi.RoleManagementRoleManagementUpdateRoleManagement(context.Background()).MicrosoftGraphRoleManagement(microsoftGraphRoleManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementRoleManagementApi.RoleManagementRoleManagementUpdateRoleManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleManagementRoleManagementUpdateRoleManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphRoleManagement** | [**MicrosoftGraphRoleManagement**](MicrosoftGraphRoleManagement.md) | New property values | 

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

