# \UsersPersonApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreatePeople**](UsersPersonApi.md#UsersCreatePeople) | **Post** /users/{user-id}/people | Create new navigation property to people for users
[**UsersDeletePeople**](UsersPersonApi.md#UsersDeletePeople) | **Delete** /users/{user-id}/people/{person-id} | Delete navigation property people for users
[**UsersGetPeople**](UsersPersonApi.md#UsersGetPeople) | **Get** /users/{user-id}/people/{person-id} | Get people from users
[**UsersListPeople**](UsersPersonApi.md#UsersListPeople) | **Get** /users/{user-id}/people | Get people from users
[**UsersUpdatePeople**](UsersPersonApi.md#UsersUpdatePeople) | **Patch** /users/{user-id}/people/{person-id} | Update the navigation property people in users



## UsersCreatePeople

> MicrosoftGraphPerson UsersCreatePeople(ctx, userId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()

Create new navigation property to people for users



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
    microsoftGraphPerson := *openapiclient.NewMicrosoftGraphPerson() // MicrosoftGraphPerson | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPersonApi.UsersCreatePeople(context.Background(), userId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPersonApi.UsersCreatePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreatePeople`: MicrosoftGraphPerson
    fmt.Fprintf(os.Stdout, "Response from `UsersPersonApi.UsersCreatePeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreatePeopleRequest struct via the builder pattern


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


## UsersDeletePeople

> UsersDeletePeople(ctx, userId, personId).IfMatch(ifMatch).Execute()

Delete navigation property people for users



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
    personId := "personId_example" // string | key: id of person
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPersonApi.UsersDeletePeople(context.Background(), userId, personId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPersonApi.UsersDeletePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePeopleRequest struct via the builder pattern


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


## UsersGetPeople

> MicrosoftGraphPerson UsersGetPeople(ctx, userId, personId).Select_(select_).Execute()

Get people from users



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
    personId := "personId_example" // string | key: id of person
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPersonApi.UsersGetPeople(context.Background(), userId, personId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPersonApi.UsersGetPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPeople`: MicrosoftGraphPerson
    fmt.Fprintf(os.Stdout, "Response from `UsersPersonApi.UsersGetPeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPeopleRequest struct via the builder pattern


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


## UsersListPeople

> CollectionOfPerson UsersListPeople(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get people from users



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPersonApi.UsersListPeople(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPersonApi.UsersListPeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListPeople`: CollectionOfPerson
    fmt.Fprintf(os.Stdout, "Response from `UsersPersonApi.UsersListPeople`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListPeopleRequest struct via the builder pattern


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


## UsersUpdatePeople

> UsersUpdatePeople(ctx, userId, personId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()

Update the navigation property people in users



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
    personId := "personId_example" // string | key: id of person
    microsoftGraphPerson := *openapiclient.NewMicrosoftGraphPerson() // MicrosoftGraphPerson | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPersonApi.UsersUpdatePeople(context.Background(), userId, personId).MicrosoftGraphPerson(microsoftGraphPerson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPersonApi.UsersUpdatePeople``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**personId** | **string** | key: id of person | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePeopleRequest struct via the builder pattern


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

