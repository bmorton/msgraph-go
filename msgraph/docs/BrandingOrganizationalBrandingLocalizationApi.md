# \BrandingOrganizationalBrandingLocalizationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrandingCreateLocalizations**](BrandingOrganizationalBrandingLocalizationApi.md#BrandingCreateLocalizations) | **Post** /branding/localizations | Create new navigation property to localizations for branding
[**BrandingDeleteLocalizations**](BrandingOrganizationalBrandingLocalizationApi.md#BrandingDeleteLocalizations) | **Delete** /branding/localizations/{organizationalBrandingLocalization-id} | Delete navigation property localizations for branding
[**BrandingGetLocalizations**](BrandingOrganizationalBrandingLocalizationApi.md#BrandingGetLocalizations) | **Get** /branding/localizations/{organizationalBrandingLocalization-id} | Get localizations from branding
[**BrandingListLocalizations**](BrandingOrganizationalBrandingLocalizationApi.md#BrandingListLocalizations) | **Get** /branding/localizations | Get localizations from branding
[**BrandingUpdateLocalizations**](BrandingOrganizationalBrandingLocalizationApi.md#BrandingUpdateLocalizations) | **Patch** /branding/localizations/{organizationalBrandingLocalization-id} | Update the navigation property localizations in branding



## BrandingCreateLocalizations

> MicrosoftGraphOrganizationalBrandingLocalization BrandingCreateLocalizations(ctx).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()

Create new navigation property to localizations for branding



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
    microsoftGraphOrganizationalBrandingLocalization := *openapiclient.NewMicrosoftGraphOrganizationalBrandingLocalization() // MicrosoftGraphOrganizationalBrandingLocalization | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandingOrganizationalBrandingLocalizationApi.BrandingCreateLocalizations(context.Background()).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingOrganizationalBrandingLocalizationApi.BrandingCreateLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandingCreateLocalizations`: MicrosoftGraphOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `BrandingOrganizationalBrandingLocalizationApi.BrandingCreateLocalizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandingCreateLocalizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOrganizationalBrandingLocalization** | [**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md) | New navigation property | 

### Return type

[**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingDeleteLocalizations

> BrandingDeleteLocalizations(ctx, organizationalBrandingLocalizationId).IfMatch(ifMatch).Execute()

Delete navigation property localizations for branding



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
    organizationalBrandingLocalizationId := "organizationalBrandingLocalizationId_example" // string | key: id of organizationalBrandingLocalization
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandingOrganizationalBrandingLocalizationApi.BrandingDeleteLocalizations(context.Background(), organizationalBrandingLocalizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingOrganizationalBrandingLocalizationApi.BrandingDeleteLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationalBrandingLocalizationId** | **string** | key: id of organizationalBrandingLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingDeleteLocalizationsRequest struct via the builder pattern


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


## BrandingGetLocalizations

> MicrosoftGraphOrganizationalBrandingLocalization BrandingGetLocalizations(ctx, organizationalBrandingLocalizationId).Select_(select_).Expand(expand).Execute()

Get localizations from branding



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
    organizationalBrandingLocalizationId := "organizationalBrandingLocalizationId_example" // string | key: id of organizationalBrandingLocalization
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandingOrganizationalBrandingLocalizationApi.BrandingGetLocalizations(context.Background(), organizationalBrandingLocalizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingOrganizationalBrandingLocalizationApi.BrandingGetLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandingGetLocalizations`: MicrosoftGraphOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `BrandingOrganizationalBrandingLocalizationApi.BrandingGetLocalizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationalBrandingLocalizationId** | **string** | key: id of organizationalBrandingLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingGetLocalizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingListLocalizations

> CollectionOfOrganizationalBrandingLocalization BrandingListLocalizations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get localizations from branding



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
    resp, r, err := api_client.BrandingOrganizationalBrandingLocalizationApi.BrandingListLocalizations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingOrganizationalBrandingLocalizationApi.BrandingListLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrandingListLocalizations`: CollectionOfOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `BrandingOrganizationalBrandingLocalizationApi.BrandingListLocalizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrandingListLocalizationsRequest struct via the builder pattern


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

[**CollectionOfOrganizationalBrandingLocalization**](CollectionOfOrganizationalBrandingLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrandingUpdateLocalizations

> BrandingUpdateLocalizations(ctx, organizationalBrandingLocalizationId).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()

Update the navigation property localizations in branding



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
    organizationalBrandingLocalizationId := "organizationalBrandingLocalizationId_example" // string | key: id of organizationalBrandingLocalization
    microsoftGraphOrganizationalBrandingLocalization := *openapiclient.NewMicrosoftGraphOrganizationalBrandingLocalization() // MicrosoftGraphOrganizationalBrandingLocalization | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrandingOrganizationalBrandingLocalizationApi.BrandingUpdateLocalizations(context.Background(), organizationalBrandingLocalizationId).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingOrganizationalBrandingLocalizationApi.BrandingUpdateLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationalBrandingLocalizationId** | **string** | key: id of organizationalBrandingLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrandingUpdateLocalizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOrganizationalBrandingLocalization** | [**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md) | New navigation property values | 

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

