# \MePersonApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreatePeople**](MePersonApi.md#MeCreatePeople) | **Post** /me/people | Create new navigation property to people for me
[**MeDeletePeople**](MePersonApi.md#MeDeletePeople) | **Delete** /me/people/{person-id} | Delete navigation property people for me
[**MeGetPeople**](MePersonApi.md#MeGetPeople) | **Get** /me/people/{person-id} | Get people from me
[**MeListPeople**](MePersonApi.md#MeListPeople) | **Get** /me/people | Get people from me
[**MeUpdatePeople**](MePersonApi.md#MeUpdatePeople) | **Patch** /me/people/{person-id} | Update the navigation property people in me



## MeCreatePeople

> MicrosoftGraphPerson MeCreatePeople(ctx).MicrosoftGraphPerson(microsoftGraphPerson).Execute()

Create new navigation property to people for me



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
    microsoftGraphPerson := *openapiclient.NewMicrosoftGraphPerson() // MicrosoftGraphPerson | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePersonApi.MeCreatePeople(context.Background()).MicrosoftGraphPerson(microsoftGraphPerson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePersonApi.MeCreatePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreatePeople`: MicrosoftGraphPerson
    fmt.Fprintf(os.Stdout, "Response from `MePersonApi.MeCreatePeople`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreatePeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPerson** | [**MicrosoftGraphPerson**](MicrosoftGraphPerson.md) | New navigation property | 

### Return type

[**MicrosoftGraphPerson**](MicrosoftGraphPerson.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeletePeople

> MeDeletePeople(ctx, personId).IfMatch(ifMatch).Execute()

Delete navigation property people for me



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
    personId := "personId_example" // string | key: id of person
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePersonApi.MeDeletePeople(context.Background(), personId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePersonApi.MeDeletePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeletePeopleRequest struct via the builder pattern


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


## MeGetPeople

> MicrosoftGraphPerson MeGetPeople(ctx, personId).Select_(select_).Execute()

Get people from me



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
    personId := "personId_example" // string | key: id of person
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePersonApi.MeGetPeople(context.Background(), personId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePersonApi.MeGetPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPeople`: MicrosoftGraphPerson
    fmt.Fprintf(os.Stdout, "Response from `MePersonApi.MeGetPeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphPerson**](MicrosoftGraphPerson.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListPeople

> CollectionOfPerson MeListPeople(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get people from me



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePersonApi.MeListPeople(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePersonApi.MeListPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListPeople`: CollectionOfPerson
    fmt.Fprintf(os.Stdout, "Response from `MePersonApi.MeListPeople`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListPeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfPerson**](CollectionOfPerson.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdatePeople

> MeUpdatePeople(ctx, personId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()

Update the navigation property people in me



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
    personId := "personId_example" // string | key: id of person
    microsoftGraphPerson := *openapiclient.NewMicrosoftGraphPerson() // MicrosoftGraphPerson | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePersonApi.MeUpdatePeople(context.Background(), personId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePersonApi.MeUpdatePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPerson** | [**MicrosoftGraphPerson**](MicrosoftGraphPerson.md) | New navigation property values | 

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

