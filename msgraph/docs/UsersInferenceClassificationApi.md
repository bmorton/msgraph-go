# \UsersInferenceClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteInferenceClassification**](UsersInferenceClassificationApi.md#UsersDeleteInferenceClassification) | **Delete** /users/{user-id}/inferenceClassification | Delete navigation property inferenceClassification for users
[**UsersGetInferenceClassification**](UsersInferenceClassificationApi.md#UsersGetInferenceClassification) | **Get** /users/{user-id}/inferenceClassification | Get inferenceClassification from users
[**UsersInferenceClassificationCreateOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationCreateOverrides) | **Post** /users/{user-id}/inferenceClassification/overrides | Create new navigation property to overrides for users
[**UsersInferenceClassificationDeleteOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationDeleteOverrides) | **Delete** /users/{user-id}/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Delete navigation property overrides for users
[**UsersInferenceClassificationGetOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationGetOverrides) | **Get** /users/{user-id}/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Get overrides from users
[**UsersInferenceClassificationListOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationListOverrides) | **Get** /users/{user-id}/inferenceClassification/overrides | Get overrides from users
[**UsersInferenceClassificationUpdateOverrides**](UsersInferenceClassificationApi.md#UsersInferenceClassificationUpdateOverrides) | **Patch** /users/{user-id}/inferenceClassification/overrides/{inferenceClassificationOverride-id} | Update the navigation property overrides in users
[**UsersUpdateInferenceClassification**](UsersInferenceClassificationApi.md#UsersUpdateInferenceClassification) | **Patch** /users/{user-id}/inferenceClassification | Update the navigation property inferenceClassification in users



## UsersDeleteInferenceClassification

> UsersDeleteInferenceClassification(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property inferenceClassification for users



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
    userId := "userId_example" // string | key: id of user
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersDeleteInferenceClassification(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersDeleteInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteInferenceClassificationRequest struct via the builder pattern


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


## UsersGetInferenceClassification

> MicrosoftGraphInferenceClassification UsersGetInferenceClassification(ctx, userId).Select_(select_).Execute()

Get inferenceClassification from users



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
    userId := "userId_example" // string | key: id of user
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersGetInferenceClassification(context.Background(), userId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersGetInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetInferenceClassification`: MicrosoftGraphInferenceClassification
    fmt.Fprintf(os.Stdout, "Response from `UsersInferenceClassificationApi.UsersGetInferenceClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetInferenceClassificationRequest struct via the builder pattern


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


## UsersInferenceClassificationCreateOverrides

> MicrosoftGraphInferenceClassificationOverride UsersInferenceClassificationCreateOverrides(ctx, userId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()

Create new navigation property to overrides for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphInferenceClassificationOverride := *openapiclient.NewMicrosoftGraphInferenceClassificationOverride() // MicrosoftGraphInferenceClassificationOverride | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersInferenceClassificationCreateOverrides(context.Background(), userId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersInferenceClassificationCreateOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersInferenceClassificationCreateOverrides`: MicrosoftGraphInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `UsersInferenceClassificationApi.UsersInferenceClassificationCreateOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersInferenceClassificationCreateOverridesRequest struct via the builder pattern


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


## UsersInferenceClassificationDeleteOverrides

> UsersInferenceClassificationDeleteOverrides(ctx, userId, inferenceClassificationOverrideId).IfMatch(ifMatch).Execute()

Delete navigation property overrides for users



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
    userId := "userId_example" // string | key: id of user
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersInferenceClassificationDeleteOverrides(context.Background(), userId, inferenceClassificationOverrideId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersInferenceClassificationDeleteOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersInferenceClassificationDeleteOverridesRequest struct via the builder pattern


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


## UsersInferenceClassificationGetOverrides

> MicrosoftGraphInferenceClassificationOverride UsersInferenceClassificationGetOverrides(ctx, userId, inferenceClassificationOverrideId).Select_(select_).Execute()

Get overrides from users



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
    userId := "userId_example" // string | key: id of user
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersInferenceClassificationGetOverrides(context.Background(), userId, inferenceClassificationOverrideId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersInferenceClassificationGetOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersInferenceClassificationGetOverrides`: MicrosoftGraphInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `UsersInferenceClassificationApi.UsersInferenceClassificationGetOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersInferenceClassificationGetOverridesRequest struct via the builder pattern


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


## UsersInferenceClassificationListOverrides

> CollectionOfInferenceClassificationOverride UsersInferenceClassificationListOverrides(ctx, userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get overrides from users



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
    userId := "userId_example" // string | key: id of user
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersInferenceClassificationListOverrides(context.Background(), userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersInferenceClassificationListOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersInferenceClassificationListOverrides`: CollectionOfInferenceClassificationOverride
    fmt.Fprintf(os.Stdout, "Response from `UsersInferenceClassificationApi.UsersInferenceClassificationListOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersInferenceClassificationListOverridesRequest struct via the builder pattern


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


## UsersInferenceClassificationUpdateOverrides

> UsersInferenceClassificationUpdateOverrides(ctx, userId, inferenceClassificationOverrideId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()

Update the navigation property overrides in users



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
    userId := "userId_example" // string | key: id of user
    inferenceClassificationOverrideId := "inferenceClassificationOverrideId_example" // string | key: id of inferenceClassificationOverride
    microsoftGraphInferenceClassificationOverride := *openapiclient.NewMicrosoftGraphInferenceClassificationOverride() // MicrosoftGraphInferenceClassificationOverride | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersInferenceClassificationUpdateOverrides(context.Background(), userId, inferenceClassificationOverrideId).MicrosoftGraphInferenceClassificationOverride(microsoftGraphInferenceClassificationOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersInferenceClassificationUpdateOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**inferenceClassificationOverrideId** | **string** | key: id of inferenceClassificationOverride | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersInferenceClassificationUpdateOverridesRequest struct via the builder pattern


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


## UsersUpdateInferenceClassification

> UsersUpdateInferenceClassification(ctx, userId).MicrosoftGraphInferenceClassification(microsoftGraphInferenceClassification).Execute()

Update the navigation property inferenceClassification in users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphInferenceClassification := *openapiclient.NewMicrosoftGraphInferenceClassification() // MicrosoftGraphInferenceClassification | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersInferenceClassificationApi.UsersUpdateInferenceClassification(context.Background(), userId).MicrosoftGraphInferenceClassification(microsoftGraphInferenceClassification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersInferenceClassificationApi.UsersUpdateInferenceClassification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateInferenceClassificationRequest struct via the builder pattern


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

