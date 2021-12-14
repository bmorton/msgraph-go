# \AdminAdminApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminGetAdmin**](AdminAdminApi.md#AdminAdminGetAdmin) | **Get** /admin | Get admin
[**AdminAdminUpdateAdmin**](AdminAdminApi.md#AdminAdminUpdateAdmin) | **Patch** /admin | Update admin



## AdminAdminGetAdmin

> MicrosoftGraphAdmin AdminAdminGetAdmin(ctx).Select_(select_).Expand(expand).Execute()

Get admin

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
    resp, r, err := api_client.AdminAdminApi.AdminAdminGetAdmin(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAdminApi.AdminAdminGetAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminAdminGetAdmin`: MicrosoftGraphAdmin
    fmt.Fprintf(os.Stdout, "Response from `AdminAdminApi.AdminAdminGetAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminGetAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAdmin**](MicrosoftGraphAdmin.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminUpdateAdmin

> AdminAdminUpdateAdmin(ctx).MicrosoftGraphAdmin(microsoftGraphAdmin).Execute()

Update admin

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
    microsoftGraphAdmin := *openapiclient.NewMicrosoftGraphAdmin() // MicrosoftGraphAdmin | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminAdminApi.AdminAdminUpdateAdmin(context.Background()).MicrosoftGraphAdmin(microsoftGraphAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAdminApi.AdminAdminUpdateAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminUpdateAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAdmin** | [**MicrosoftGraphAdmin**](MicrosoftGraphAdmin.md) | New property values | 

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

