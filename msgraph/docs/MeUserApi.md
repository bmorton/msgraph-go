# \MeUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeUserGetUser**](MeUserApi.md#MeUserGetUser) | **Get** /me | Get me
[**MeUserUpdateUser**](MeUserApi.md#MeUserUpdateUser) | **Patch** /me | Update me



## MeUserGetUser

> MicrosoftGraphUser MeUserGetUser(ctx).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()

Get me

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
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserApi.MeUserGetUser(context.Background()).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserApi.MeUserGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeUserGetUser`: MicrosoftGraphUser
    fmt.Fprintf(os.Stdout, "Response from `MeUserApi.MeUserGetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUserGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUser**](MicrosoftGraphUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUserUpdateUser

> MeUserUpdateUser(ctx).MicrosoftGraphUser(microsoftGraphUser).Execute()

Update me

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
    microsoftGraphUser := *openapiclient.NewMicrosoftGraphUser() // MicrosoftGraphUser | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserApi.MeUserUpdateUser(context.Background()).MicrosoftGraphUser(microsoftGraphUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserApi.MeUserUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUserUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphUser** | [**MicrosoftGraphUser**](MicrosoftGraphUser.md) | New property values | 

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

