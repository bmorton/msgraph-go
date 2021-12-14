# \InvitationsUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationsDeleteRefInvitedUser**](InvitationsUserApi.md#InvitationsDeleteRefInvitedUser) | **Delete** /invitations/{invitation-id}/invitedUser/$ref | Delete ref of navigation property invitedUser for invitations
[**InvitationsGetInvitedUser**](InvitationsUserApi.md#InvitationsGetInvitedUser) | **Get** /invitations/{invitation-id}/invitedUser | Get invitedUser from invitations
[**InvitationsGetRefInvitedUser**](InvitationsUserApi.md#InvitationsGetRefInvitedUser) | **Get** /invitations/{invitation-id}/invitedUser/$ref | Get ref of invitedUser from invitations
[**InvitationsUpdateRefInvitedUser**](InvitationsUserApi.md#InvitationsUpdateRefInvitedUser) | **Put** /invitations/{invitation-id}/invitedUser/$ref | Update the ref of navigation property invitedUser in invitations



## InvitationsDeleteRefInvitedUser

> InvitationsDeleteRefInvitedUser(ctx, invitationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property invitedUser for invitations



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
    resp, r, err := api_client.InvitationsUserApi.InvitationsDeleteRefInvitedUser(context.Background(), invitationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsUserApi.InvitationsDeleteRefInvitedUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInvitationsDeleteRefInvitedUserRequest struct via the builder pattern


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


## InvitationsGetInvitedUser

> MicrosoftGraphUser InvitationsGetInvitedUser(ctx, invitationId).Select_(select_).Expand(expand).Execute()

Get invitedUser from invitations



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
    resp, r, err := api_client.InvitationsUserApi.InvitationsGetInvitedUser(context.Background(), invitationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsUserApi.InvitationsGetInvitedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsGetInvitedUser`: MicrosoftGraphUser
    fmt.Fprintf(os.Stdout, "Response from `InvitationsUserApi.InvitationsGetInvitedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsGetInvitedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## InvitationsGetRefInvitedUser

> string InvitationsGetRefInvitedUser(ctx, invitationId).Execute()

Get ref of invitedUser from invitations



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsUserApi.InvitationsGetRefInvitedUser(context.Background(), invitationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsUserApi.InvitationsGetRefInvitedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvitationsGetRefInvitedUser`: string
    fmt.Fprintf(os.Stdout, "Response from `InvitationsUserApi.InvitationsGetRefInvitedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsGetRefInvitedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitationsUpdateRefInvitedUser

> InvitationsUpdateRefInvitedUser(ctx, invitationId).RequestBody(requestBody).Execute()

Update the ref of navigation property invitedUser in invitations



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsUserApi.InvitationsUpdateRefInvitedUser(context.Background(), invitationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsUserApi.InvitationsUpdateRefInvitedUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInvitationsUpdateRefInvitedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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

