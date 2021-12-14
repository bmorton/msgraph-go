# \LocalizationsOrganizationalBrandingLocalizationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization**](LocalizationsOrganizationalBrandingLocalizationApi.md#LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization) | **Post** /localizations | Add new entity to localizations
[**LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization**](LocalizationsOrganizationalBrandingLocalizationApi.md#LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization) | **Delete** /localizations/{organizationalBrandingLocalization-id} | Delete entity from localizations
[**LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization**](LocalizationsOrganizationalBrandingLocalizationApi.md#LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization) | **Get** /localizations/{organizationalBrandingLocalization-id} | Get entity from localizations by key
[**LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization**](LocalizationsOrganizationalBrandingLocalizationApi.md#LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization) | **Get** /localizations | Get entities from localizations
[**LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization**](LocalizationsOrganizationalBrandingLocalizationApi.md#LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization) | **Patch** /localizations/{organizationalBrandingLocalization-id} | Update entity in localizations



## LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization

> MicrosoftGraphOrganizationalBrandingLocalization LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization(ctx).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()

Add new entity to localizations

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
    microsoftGraphOrganizationalBrandingLocalization := *openapiclient.NewMicrosoftGraphOrganizationalBrandingLocalization() // MicrosoftGraphOrganizationalBrandingLocalization | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization(context.Background()).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization`: MicrosoftGraphOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocalizationsOrganizationalBrandingLocalizationCreateOrganizationalBrandingLocalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOrganizationalBrandingLocalization** | [**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md) | New entity | 

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


## LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization

> LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization(ctx, organizationalBrandingLocalizationId).IfMatch(ifMatch).Execute()

Delete entity from localizations

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
    resp, r, err := api_client.LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization(context.Background(), organizationalBrandingLocalizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLocalizationsOrganizationalBrandingLocalizationDeleteOrganizationalBrandingLocalizationRequest struct via the builder pattern


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


## LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization

> MicrosoftGraphOrganizationalBrandingLocalization LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization(ctx, organizationalBrandingLocalizationId).Select_(select_).Expand(expand).Execute()

Get entity from localizations by key

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
    resp, r, err := api_client.LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization(context.Background(), organizationalBrandingLocalizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization`: MicrosoftGraphOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationalBrandingLocalizationId** | **string** | key: id of organizationalBrandingLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocalizationsOrganizationalBrandingLocalizationGetOrganizationalBrandingLocalizationRequest struct via the builder pattern


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


## LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization

> CollectionOfOrganizationalBrandingLocalization LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from localizations

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
    resp, r, err := api_client.LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization`: CollectionOfOrganizationalBrandingLocalization
    fmt.Fprintf(os.Stdout, "Response from `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocalizationsOrganizationalBrandingLocalizationListOrganizationalBrandingLocalizationRequest struct via the builder pattern


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


## LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization

> LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization(ctx, organizationalBrandingLocalizationId).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()

Update entity in localizations

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
    microsoftGraphOrganizationalBrandingLocalization := *openapiclient.NewMicrosoftGraphOrganizationalBrandingLocalization() // MicrosoftGraphOrganizationalBrandingLocalization | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization(context.Background(), organizationalBrandingLocalizationId).MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalizationsOrganizationalBrandingLocalizationApi.LocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLocalizationsOrganizationalBrandingLocalizationUpdateOrganizationalBrandingLocalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOrganizationalBrandingLocalization** | [**MicrosoftGraphOrganizationalBrandingLocalization**](MicrosoftGraphOrganizationalBrandingLocalization.md) | New property values | 

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

