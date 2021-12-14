# \InvitationsInvitationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationsInvitationCreateInvitation**](InvitationsInvitationApi.md#InvitationsInvitationCreateInvitation) | **Post** /invitations | Add new entity to invitations
[**InvitationsInvitationDeleteInvitation**](InvitationsInvitationApi.md#InvitationsInvitationDeleteInvitation) | **Delete** /invitations/{invitation-id} | Delete entity from invitations
[**InvitationsInvitationGetInvitation**](InvitationsInvitationApi.md#InvitationsInvitationGetInvitation) | **Get** /invitations/{invitation-id} | Get entity from invitations by key
[**InvitationsInvitationListInvitation**](InvitationsInvitationApi.md#InvitationsInvitationListInvitation) | **Get** /invitations | Get entities from invitations
[**InvitationsInvitationUpdateInvitation**](InvitationsInvitationApi.md#InvitationsInvitationUpdateInvitation) | **Patch** /invitations/{invitation-id} | Update entity in invitations



## InvitationsInvitationCreateInvitation

> MicrosoftGraphInvitation InvitationsInvitationCreateInvitation(ctx).MicrosoftGraphInvitation(microsoftGraphInvitation).Execute()

Add new entity to invitations

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
    microsoftGraphInvitation := *openapiclient.NewMicrosoftGraphInvitation() // MicrosoftGraphInvitation | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsInvitationApi.InvitationsInvitationCreateInvitation(context.Background()).MicrosoftGraphInvitation(microsoftGraphInvitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsInvitationApi.InvitationsInvitationCreateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsInvitationCreateInvitation`: MicrosoftGraphInvitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsInvitationApi.InvitationsInvitationCreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsInvitationCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphInvitation** | [**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md) | New entity | 

### Return type

[**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationDeleteInvitation

> InvitationsInvitationDeleteInvitation(ctx, invitationId).IfMatch(ifMatch).Execute()

Delete entity from invitations

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
    invitationId := "invitationId_example" // string | key: id of invitation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsInvitationApi.InvitationsInvitationDeleteInvitation(context.Background(), invitationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsInvitationApi.InvitationsInvitationDeleteInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsInvitationDeleteInvitationRequest struct via the builder pattern


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


## InvitationsInvitationGetInvitation

> MicrosoftGraphInvitation InvitationsInvitationGetInvitation(ctx, invitationId).Select_(select_).Expand(expand).Execute()

Get entity from invitations by key

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
    invitationId := "invitationId_example" // string | key: id of invitation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsInvitationApi.InvitationsInvitationGetInvitation(context.Background(), invitationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsInvitationApi.InvitationsInvitationGetInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsInvitationGetInvitation`: MicrosoftGraphInvitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsInvitationApi.InvitationsInvitationGetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsInvitationGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationListInvitation

> CollectionOfInvitation InvitationsInvitationListInvitation(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from invitations

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
    resp, r, err := api_client.InvitationsInvitationApi.InvitationsInvitationListInvitation(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsInvitationApi.InvitationsInvitationListInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsInvitationListInvitation`: CollectionOfInvitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsInvitationApi.InvitationsInvitationListInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsInvitationListInvitationRequest struct via the builder pattern


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

[**CollectionOfInvitation**](CollectionOfInvitation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsInvitationUpdateInvitation

> InvitationsInvitationUpdateInvitation(ctx, invitationId).MicrosoftGraphInvitation(microsoftGraphInvitation).Execute()

Update entity in invitations

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
    invitationId := "invitationId_example" // string | key: id of invitation
    microsoftGraphInvitation := *openapiclient.NewMicrosoftGraphInvitation() // MicrosoftGraphInvitation | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsInvitationApi.InvitationsInvitationUpdateInvitation(context.Background(), invitationId).MicrosoftGraphInvitation(microsoftGraphInvitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsInvitationApi.InvitationsInvitationUpdateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsInvitationUpdateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphInvitation** | [**MicrosoftGraphInvitation**](MicrosoftGraphInvitation.md) | New property values | 

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

