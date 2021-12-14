# \MeOutlookUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteOutlook**](MeOutlookUserApi.md#MeDeleteOutlook) | **Delete** /me/outlook | Delete navigation property outlook for me
[**MeGetOutlook**](MeOutlookUserApi.md#MeGetOutlook) | **Get** /me/outlook | Get outlook from me
[**MeOutlookCreateMasterCategories**](MeOutlookUserApi.md#MeOutlookCreateMasterCategories) | **Post** /me/outlook/masterCategories | Create new navigation property to masterCategories for me
[**MeOutlookDeleteMasterCategories**](MeOutlookUserApi.md#MeOutlookDeleteMasterCategories) | **Delete** /me/outlook/masterCategories/{outlookCategory-id} | Delete navigation property masterCategories for me
[**MeOutlookGetMasterCategories**](MeOutlookUserApi.md#MeOutlookGetMasterCategories) | **Get** /me/outlook/masterCategories/{outlookCategory-id} | Get masterCategories from me
[**MeOutlookListMasterCategories**](MeOutlookUserApi.md#MeOutlookListMasterCategories) | **Get** /me/outlook/masterCategories | Get masterCategories from me
[**MeOutlookUpdateMasterCategories**](MeOutlookUserApi.md#MeOutlookUpdateMasterCategories) | **Patch** /me/outlook/masterCategories/{outlookCategory-id} | Update the navigation property masterCategories in me
[**MeUpdateOutlook**](MeOutlookUserApi.md#MeUpdateOutlook) | **Patch** /me/outlook | Update the navigation property outlook in me



## MeDeleteOutlook

> MeDeleteOutlook(ctx).IfMatch(ifMatch).Execute()

Delete navigation property outlook for me



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
    resp, r, err := api_client.MeOutlookUserApi.MeDeleteOutlook(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeDeleteOutlook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteOutlookRequest struct via the builder pattern


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


## MeGetOutlook

> MicrosoftGraphOutlookUser MeGetOutlook(ctx).Select_(select_).Execute()

Get outlook from me



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeGetOutlook(context.Background()).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeGetOutlook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetOutlook`: MicrosoftGraphOutlookUser
    fmt.Fprintf(os.Stdout, "Response from `MeOutlookUserApi.MeGetOutlook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetOutlookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookCreateMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookCreateMasterCategories(ctx).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()

Create new navigation property to masterCategories for me



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
    microsoftGraphOutlookCategory := *openapiclient.NewMicrosoftGraphOutlookCategory() // MicrosoftGraphOutlookCategory | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeOutlookCreateMasterCategories(context.Background()).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeOutlookCreateMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeOutlookCreateMasterCategories`: MicrosoftGraphOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `MeOutlookUserApi.MeOutlookCreateMasterCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeOutlookCreateMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md) | New navigation property | 

### Return type

[**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookDeleteMasterCategories

> MeOutlookDeleteMasterCategories(ctx, outlookCategoryId).IfMatch(ifMatch).Execute()

Delete navigation property masterCategories for me



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeOutlookDeleteMasterCategories(context.Background(), outlookCategoryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeOutlookDeleteMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeOutlookDeleteMasterCategoriesRequest struct via the builder pattern


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


## MeOutlookGetMasterCategories

> MicrosoftGraphOutlookCategory MeOutlookGetMasterCategories(ctx, outlookCategoryId).Select_(select_).Execute()

Get masterCategories from me



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeOutlookGetMasterCategories(context.Background(), outlookCategoryId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeOutlookGetMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeOutlookGetMasterCategories`: MicrosoftGraphOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `MeOutlookUserApi.MeOutlookGetMasterCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeOutlookGetMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookListMasterCategories

> CollectionOfOutlookCategory MeOutlookListMasterCategories(ctx).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get masterCategories from me



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
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeOutlookListMasterCategories(context.Background()).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeOutlookListMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeOutlookListMasterCategories`: CollectionOfOutlookCategory
    fmt.Fprintf(os.Stdout, "Response from `MeOutlookUserApi.MeOutlookListMasterCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeOutlookListMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfOutlookCategory**](CollectionOfOutlookCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeOutlookUpdateMasterCategories

> MeOutlookUpdateMasterCategories(ctx, outlookCategoryId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()

Update the navigation property masterCategories in me



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
    outlookCategoryId := "outlookCategoryId_example" // string | key: id of outlookCategory
    microsoftGraphOutlookCategory := *openapiclient.NewMicrosoftGraphOutlookCategory() // MicrosoftGraphOutlookCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeOutlookUpdateMasterCategories(context.Background(), outlookCategoryId).MicrosoftGraphOutlookCategory(microsoftGraphOutlookCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeOutlookUpdateMasterCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outlookCategoryId** | **string** | key: id of outlookCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeOutlookUpdateMasterCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOutlookCategory** | [**MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md) | New navigation property values | 

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


## MeUpdateOutlook

> MeUpdateOutlook(ctx).MicrosoftGraphOutlookUser(microsoftGraphOutlookUser).Execute()

Update the navigation property outlook in me



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
    microsoftGraphOutlookUser := *openapiclient.NewMicrosoftGraphOutlookUser() // MicrosoftGraphOutlookUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOutlookUserApi.MeUpdateOutlook(context.Background()).MicrosoftGraphOutlookUser(microsoftGraphOutlookUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOutlookUserApi.MeUpdateOutlook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateOutlookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOutlookUser** | [**MicrosoftGraphOutlookUser**](MicrosoftGraphOutlookUser.md) | New navigation property values | 

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

