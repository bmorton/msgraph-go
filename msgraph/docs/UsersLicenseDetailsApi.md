# \UsersLicenseDetailsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateLicenseDetails**](UsersLicenseDetailsApi.md#UsersCreateLicenseDetails) | **Post** /users/{user-id}/licenseDetails | Create new navigation property to licenseDetails for users
[**UsersDeleteLicenseDetails**](UsersLicenseDetailsApi.md#UsersDeleteLicenseDetails) | **Delete** /users/{user-id}/licenseDetails/{licenseDetails-id} | Delete navigation property licenseDetails for users
[**UsersGetLicenseDetails**](UsersLicenseDetailsApi.md#UsersGetLicenseDetails) | **Get** /users/{user-id}/licenseDetails/{licenseDetails-id} | Get licenseDetails from users
[**UsersListLicenseDetails**](UsersLicenseDetailsApi.md#UsersListLicenseDetails) | **Get** /users/{user-id}/licenseDetails | Get licenseDetails from users
[**UsersUpdateLicenseDetails**](UsersLicenseDetailsApi.md#UsersUpdateLicenseDetails) | **Patch** /users/{user-id}/licenseDetails/{licenseDetails-id} | Update the navigation property licenseDetails in users



## UsersCreateLicenseDetails

> MicrosoftGraphLicenseDetails UsersCreateLicenseDetails(ctx, userId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()

Create new navigation property to licenseDetails for users



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
    microsoftGraphLicenseDetails := *openapiclient.NewMicrosoftGraphLicenseDetails() // MicrosoftGraphLicenseDetails | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersLicenseDetailsApi.UsersCreateLicenseDetails(context.Background(), userId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLicenseDetailsApi.UsersCreateLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateLicenseDetails`: MicrosoftGraphLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersLicenseDetailsApi.UsersCreateLicenseDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateLicenseDetailsRequest struct via the builder pattern


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


## UsersDeleteLicenseDetails

> UsersDeleteLicenseDetails(ctx, userId, licenseDetailsId).IfMatch(ifMatch).Execute()

Delete navigation property licenseDetails for users



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersLicenseDetailsApi.UsersDeleteLicenseDetails(context.Background(), userId, licenseDetailsId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLicenseDetailsApi.UsersDeleteLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteLicenseDetailsRequest struct via the builder pattern


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


## UsersGetLicenseDetails

> MicrosoftGraphLicenseDetails UsersGetLicenseDetails(ctx, userId, licenseDetailsId).Select_(select_).Expand(expand).Execute()

Get licenseDetails from users



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersLicenseDetailsApi.UsersGetLicenseDetails(context.Background(), userId, licenseDetailsId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLicenseDetailsApi.UsersGetLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetLicenseDetails`: MicrosoftGraphLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersLicenseDetailsApi.UsersGetLicenseDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetLicenseDetailsRequest struct via the builder pattern


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


## UsersListLicenseDetails

> CollectionOfLicenseDetails UsersListLicenseDetails(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get licenseDetails from users



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
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersLicenseDetailsApi.UsersListLicenseDetails(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLicenseDetailsApi.UsersListLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListLicenseDetails`: CollectionOfLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersLicenseDetailsApi.UsersListLicenseDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListLicenseDetailsRequest struct via the builder pattern


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


## UsersUpdateLicenseDetails

> UsersUpdateLicenseDetails(ctx, userId, licenseDetailsId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()

Update the navigation property licenseDetails in users



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
    licenseDetailsId := "licenseDetailsId_example" // string | key: id of licenseDetails
    microsoftGraphLicenseDetails := *openapiclient.NewMicrosoftGraphLicenseDetails() // MicrosoftGraphLicenseDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersLicenseDetailsApi.UsersUpdateLicenseDetails(context.Background(), userId, licenseDetailsId).MicrosoftGraphLicenseDetails(microsoftGraphLicenseDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersLicenseDetailsApi.UsersUpdateLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**licenseDetailsId** | **string** | key: id of licenseDetails | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateLicenseDetailsRequest struct via the builder pattern


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

