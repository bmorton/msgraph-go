# \OrganizationOrganizationalBrandingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationDeleteBranding**](OrganizationOrganizationalBrandingApi.md#OrganizationDeleteBranding) | **Delete** /organization/{organization-id}/branding | Delete navigation property branding for organization
[**OrganizationGetBranding**](OrganizationOrganizationalBrandingApi.md#OrganizationGetBranding) | **Get** /organization/{organization-id}/branding | Get branding from organization
[**OrganizationUpdateBranding**](OrganizationOrganizationalBrandingApi.md#OrganizationUpdateBranding) | **Patch** /organization/{organization-id}/branding | Update the navigation property branding in organization



## OrganizationDeleteBranding

> OrganizationDeleteBranding(ctx, organizationId).IfMatch(ifMatch).Execute()

Delete navigation property branding for organization

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
    organizationId := "organizationId_example" // string | key: id of organization
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationalBrandingApi.OrganizationDeleteBranding(context.Background(), organizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationalBrandingApi.OrganizationDeleteBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationDeleteBrandingRequest struct via the builder pattern


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


## OrganizationGetBranding

> MicrosoftGraphOrganizationalBranding OrganizationGetBranding(ctx, organizationId).Select_(select_).Expand(expand).Execute()

Get branding from organization

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
    organizationId := "organizationId_example" // string | key: id of organization
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationalBrandingApi.OrganizationGetBranding(context.Background(), organizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationalBrandingApi.OrganizationGetBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationGetBranding`: MicrosoftGraphOrganizationalBranding
    fmt.Fprintf(os.Stdout, "Response from `OrganizationOrganizationalBrandingApi.OrganizationGetBranding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGetBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOrganizationalBranding**](MicrosoftGraphOrganizationalBranding.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUpdateBranding

> OrganizationUpdateBranding(ctx, organizationId).MicrosoftGraphOrganizationalBranding(microsoftGraphOrganizationalBranding).Execute()

Update the navigation property branding in organization

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
    organizationId := "organizationId_example" // string | key: id of organization
    microsoftGraphOrganizationalBranding := *openapiclient.NewMicrosoftGraphOrganizationalBranding() // MicrosoftGraphOrganizationalBranding | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationalBrandingApi.OrganizationUpdateBranding(context.Background(), organizationId).MicrosoftGraphOrganizationalBranding(microsoftGraphOrganizationalBranding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationalBrandingApi.OrganizationUpdateBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUpdateBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOrganizationalBranding** | [**MicrosoftGraphOrganizationalBranding**](MicrosoftGraphOrganizationalBranding.md) | New navigation property values | 

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

