# \MeInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteInferenceClassification**](MeInferenceClassificationApi.md#MeDeleteInferenceClassification) | **Delete** /me/inferenceClassification | Delete navigation property inferenceClassification for me
[**MeGetInferenceClassification**](MeInferenceClassificationApi.md#MeGetInferenceClassification) | **Get** /me/inferenceClassification | Get inferenceClassification from me
[**MeInferenceClassificationCreateOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationCreateOverrides) | **Post** /me/inferenceClassification/overrides | Create new navigation property to overrides for me
[**MeInferenceClassificationDeleteOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationDeleteOverrides) | **Delete** /me/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Delete navigation property overrides for me
[**MeInferenceClassificationGetOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationGetOverrides) | **Get** /me/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Get overrides from me
[**MeInferenceClassificationListOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationListOverrides) | **Get** /me/inferenceClassification/overrides | Get overrides from me
[**MeInferenceClassificationUpdateOverrides**](MeInferenceClassificationApi.md#MeInferenceClassificationUpdateOverrides) | **Patch** /me/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Update the navigation property overrides in me
[**MeUpdateInferenceClassification**](MeInferenceClassificationApi.md#MeUpdateInferenceClassification) | **Patch** /me/inferenceClassification | Update the navigation property inferenceClassification in me



## MeDeleteInferenceClassification

> MeDeleteInferenceClassification(ctx).IfMatch(ifMatch).Execute()

Delete navigation property inferenceClassification for me



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
    resp, r, err := api_client.MeInferenceClassificationApi.MeDeleteInferenceClassification(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeDeleteInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteInferenceClassificationRequest struct via the builder pattern


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


## MeGetInferenceClassification

> MicrosoftGraphInferenceClassification MeGetInferenceClassification(ctx).Select_(select_).Execute()

Get inferenceClassification from me



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
    resp, r, err := api_client.MeInferenceClassificationApi.MeGetInferenceClassification(context.Background()).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeGetInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetInferenceClassification`: MicrosoftGraphInferenceClassification
    fmt.Fprintf(os.Stdout, "Response from `MeInferenceClassificationApi.MeGetInferenceClassification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetInferenceClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphInferenceClassification**](MicrosoftGraphInferenceClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInferenceClassificationCreateOverrides

> MicrosoftGraphInferenceClassificationOverride MeInferenceClassificationCreateOverrides(ctx).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()

Create new navigation property to overrides for me



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
    microsoftGraphInferenceClassificationOverride := *openapiclient.NewMicrosoftGraphInferenceClassificationOverride() // MicrosoftGraphInferenceClassificationOverride | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeInferenceClassificationApi.MeInferenceClassificationCreateOverrides(context.Background()).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeInferenceClassificationCreateOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInferenceClassificationCreateOverrides`: MicrosoftGraphInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `MeInferenceClassificationApi.MeInferenceClassificationCreateOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInferenceClassificationCreateOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphInferenceClassificationOverride** | [**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md) | New navigation property | 

### Return type

[**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInferenceClassificationDeleteOverrides

> MeInferenceClassificationDeleteOverrides(ctx, inferenceClassificationOverrideId).IfMatch(ifMatch).Execute()

Delete navigation property overrides for me



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
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeInferenceClassificationApi.MeInferenceClassificationDeleteOverrides(context.Background(), inferenceClassificationOverrideId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeInferenceClassificationDeleteOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInferenceClassificationDeleteOverridesRequest struct via the builder pattern


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


## MeInferenceClassificationGetOverrides

> MicrosoftGraphInferenceClassificationOverride MeInferenceClassificationGetOverrides(ctx, inferenceClassificationOverrideId).Select_(select_).Execute()

Get overrides from me



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
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeInferenceClassificationApi.MeInferenceClassificationGetOverrides(context.Background(), inferenceClassificationOverrideId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeInferenceClassificationGetOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInferenceClassificationGetOverrides`: MicrosoftGraphInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `MeInferenceClassificationApi.MeInferenceClassificationGetOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInferenceClassificationGetOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInferenceClassificationListOverrides

> CollectionOfInferenceClassificationOverride MeInferenceClassificationListOverrides(ctx).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get overrides from me



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
    resp, r, err := api_client.MeInferenceClassificationApi.MeInferenceClassificationListOverrides(context.Background()).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeInferenceClassificationListOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInferenceClassificationListOverrides`: CollectionOfInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `MeInferenceClassificationApi.MeInferenceClassificationListOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInferenceClassificationListOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfInferenceClassificationOverride**](CollectionOfInferenceClassificationOverride.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInferenceClassificationUpdateOverrides

> MeInferenceClassificationUpdateOverrides(ctx, inferenceClassificationOverrideId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()

Update the navigation property overrides in me



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
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    microsoftGraphInferenceClassificationOverride := *openapiclient.NewMicrosoftGraphInferenceClassificationOverride() // MicrosoftGraphInferenceClassificationOverride | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeInferenceClassificationApi.MeInferenceClassificationUpdateOverrides(context.Background(), inferenceClassificationOverrideId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeInferenceClassificationUpdateOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInferenceClassificationUpdateOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphInferenceClassificationOverride** | [**MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md) | New navigation property values | 

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


## MeUpdateInferenceClassification

> MeUpdateInferenceClassification(ctx).MicrosoftGraphInferenceClassification(microsoftGraphInferenceClassification).Execute()

Update the navigation property inferenceClassification in me



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
    microsoftGraphInferenceClassification := *openapiclient.NewMicrosoftGraphInferenceClassification() // MicrosoftGraphInferenceClassification | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeInferenceClassificationApi.MeUpdateInferenceClassification(context.Background()).MicrosoftGraphInferenceClassification(microsoftGraphInferenceClassification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeInferenceClassificationApi.MeUpdateInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateInferenceClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphInferenceClassification** | [**MicrosoftGraphInferenceClassification**](MicrosoftGraphInferenceClassification.md) | New navigation property values | 

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

