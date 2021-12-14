# \MeLicenseDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateLicenseDetails**](MeLicenseDetailsApi.md#MeCreateLicenseDetails) | **Post** /me/licenseDetails | Create new navigation property to licenseDetails for me
[**MeDeleteLicenseDetails**](MeLicenseDetailsApi.md#MeDeleteLicenseDetails) | **Delete** /me/licenseDetails/{licenseDetails-id} | Delete navigation property licenseDetails for me
[**MeGetLicenseDetails**](MeLicenseDetailsApi.md#MeGetLicenseDetails) | **Get** /me/licenseDetails/{licenseDetails-id} | Get licenseDetails from me
[**MeListLicenseDetails**](MeLicenseDetailsApi.md#MeListLicenseDetails) | **Get** /me/licenseDetails | Get licenseDetails from me
[**MeUpdateLicenseDetails**](MeLicenseDetailsApi.md#MeUpdateLicenseDetails) | **Patch** /me/licenseDetails/{licenseDetails-id} | Update the navigation property licenseDetails in me



## MeCreateLicenseDetails

> MicrosoftGraphLicenseDetails MeCreateLicenseDetails(ctx).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()

Create new navigation property to licenseDetails for me



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
    microsoftGraphLicenseDetails := *openapiclient.NewMicrosoftGraphLicenseDetails() // MicrosoftGraphLicenseDetails | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeLicenseDetailsApi.MeCreateLicenseDetails(context.Background()).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeLicenseDetailsApi.MeCreateLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateLicenseDetails`: MicrosoftGraphLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `MeLicenseDetailsApi.MeCreateLicenseDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateLicenseDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphLicenseDetails** | [**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md) | New navigation property | 

### Return type

[**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteLicenseDetails

> MeDeleteLicenseDetails(ctx, licenseDetailsId).IfMatch(ifMatch).Execute()

Delete navigation property licenseDetails for me



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeLicenseDetailsApi.MeDeleteLicenseDetails(context.Background(), licenseDetailsId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeLicenseDetailsApi.MeDeleteLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteLicenseDetailsRequest struct via the builder pattern


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


## MeGetLicenseDetails

> MicrosoftGraphLicenseDetails MeGetLicenseDetails(ctx, licenseDetailsId).Select_(select_).Expand(expand).Execute()

Get licenseDetails from me



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeLicenseDetailsApi.MeGetLicenseDetails(context.Background(), licenseDetailsId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeLicenseDetailsApi.MeGetLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetLicenseDetails`: MicrosoftGraphLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `MeLicenseDetailsApi.MeGetLicenseDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetLicenseDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListLicenseDetails

> CollectionOfLicenseDetails MeListLicenseDetails(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get licenseDetails from me



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
    resp, r, err := api_client.MeLicenseDetailsApi.MeListLicenseDetails(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeLicenseDetailsApi.MeListLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListLicenseDetails`: CollectionOfLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `MeLicenseDetailsApi.MeListLicenseDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListLicenseDetailsRequest struct via the builder pattern


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

[**CollectionOfLicenseDetails**](CollectionOfLicenseDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateLicenseDetails

> MeUpdateLicenseDetails(ctx, licenseDetailsId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()

Update the navigation property licenseDetails in me



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    microsoftGraphLicenseDetails := *openapiclient.NewMicrosoftGraphLicenseDetails() // MicrosoftGraphLicenseDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeLicenseDetailsApi.MeUpdateLicenseDetails(context.Background(), licenseDetailsId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeLicenseDetailsApi.MeUpdateLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateLicenseDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphLicenseDetails** | [**MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md) | New navigation property values | 

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

