# \SecuritySecureScoreControlProfileApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityCreateSecureScoreControlProfiles) | **Post** /security/secureScoreControlProfiles | Create new navigation property to secureScoreControlProfiles for security
[**SecurityDeleteSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityDeleteSecureScoreControlProfiles) | **Delete** /security/secureScoreControlProfiles/{secureScoreControlProfile-id} | Delete navigation property secureScoreControlProfiles for security
[**SecurityGetSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityGetSecureScoreControlProfiles) | **Get** /security/secureScoreControlProfiles/{secureScoreControlProfile-id} | Get secureScoreControlProfiles from security
[**SecurityListSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityListSecureScoreControlProfiles) | **Get** /security/secureScoreControlProfiles | Get secureScoreControlProfiles from security
[**SecurityUpdateSecureScoreControlProfiles**](SecuritySecureScoreControlProfileApi.md#SecurityUpdateSecureScoreControlProfiles) | **Patch** /security/secureScoreControlProfiles/{secureScoreControlProfile-id} | Update the navigation property secureScoreControlProfiles in security



## SecurityCreateSecureScoreControlProfiles

> MicrosoftGraphSecureScoreControlProfile SecurityCreateSecureScoreControlProfiles(ctx).MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile).Execute()

Create new navigation property to secureScoreControlProfiles for security

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
    microsoftGraphSecureScoreControlProfile := *openapiclient.NewMicrosoftGraphSecureScoreControlProfile() // MicrosoftGraphSecureScoreControlProfile | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreControlProfileApi.SecurityCreateSecureScoreControlProfiles(context.Background()).MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreControlProfileApi.SecurityCreateSecureScoreControlProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityCreateSecureScoreControlProfiles`: MicrosoftGraphSecureScoreControlProfile
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreControlProfileApi.SecurityCreateSecureScoreControlProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCreateSecureScoreControlProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSecureScoreControlProfile** | [**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md) | New navigation property | 

### Return type

[**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityDeleteSecureScoreControlProfiles

> SecurityDeleteSecureScoreControlProfiles(ctx, secureScoreControlProfileId).IfMatch(ifMatch).Execute()

Delete navigation property secureScoreControlProfiles for security

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
    secureScoreControlProfileId := "secureScoreControlProfileId_example" // string | key: id of secureScoreControlProfile
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreControlProfileApi.SecurityDeleteSecureScoreControlProfiles(context.Background(), secureScoreControlProfileId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreControlProfileApi.SecurityDeleteSecureScoreControlProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreControlProfileId** | **string** | key: id of secureScoreControlProfile | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityDeleteSecureScoreControlProfilesRequest struct via the builder pattern


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


## SecurityGetSecureScoreControlProfiles

> MicrosoftGraphSecureScoreControlProfile SecurityGetSecureScoreControlProfiles(ctx, secureScoreControlProfileId).Select_(select_).Expand(expand).Execute()

Get secureScoreControlProfiles from security

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
    secureScoreControlProfileId := "secureScoreControlProfileId_example" // string | key: id of secureScoreControlProfile
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreControlProfileApi.SecurityGetSecureScoreControlProfiles(context.Background(), secureScoreControlProfileId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreControlProfileApi.SecurityGetSecureScoreControlProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityGetSecureScoreControlProfiles`: MicrosoftGraphSecureScoreControlProfile
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreControlProfileApi.SecurityGetSecureScoreControlProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreControlProfileId** | **string** | key: id of secureScoreControlProfile | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGetSecureScoreControlProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListSecureScoreControlProfiles

> CollectionOfSecureScoreControlProfile SecurityListSecureScoreControlProfiles(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get secureScoreControlProfiles from security

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
    resp, r, err := api_client.SecuritySecureScoreControlProfileApi.SecurityListSecureScoreControlProfiles(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreControlProfileApi.SecurityListSecureScoreControlProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityListSecureScoreControlProfiles`: CollectionOfSecureScoreControlProfile
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreControlProfileApi.SecurityListSecureScoreControlProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityListSecureScoreControlProfilesRequest struct via the builder pattern


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

[**CollectionOfSecureScoreControlProfile**](CollectionOfSecureScoreControlProfile.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateSecureScoreControlProfiles

> SecurityUpdateSecureScoreControlProfiles(ctx, secureScoreControlProfileId).MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile).Execute()

Update the navigation property secureScoreControlProfiles in security

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
    secureScoreControlProfileId := "secureScoreControlProfileId_example" // string | key: id of secureScoreControlProfile
    microsoftGraphSecureScoreControlProfile := *openapiclient.NewMicrosoftGraphSecureScoreControlProfile() // MicrosoftGraphSecureScoreControlProfile | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreControlProfileApi.SecurityUpdateSecureScoreControlProfiles(context.Background(), secureScoreControlProfileId).MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreControlProfileApi.SecurityUpdateSecureScoreControlProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreControlProfileId** | **string** | key: id of secureScoreControlProfile | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityUpdateSecureScoreControlProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSecureScoreControlProfile** | [**MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md) | New navigation property values | 

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

