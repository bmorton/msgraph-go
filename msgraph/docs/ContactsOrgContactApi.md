# \ContactsOrgContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsOrgContactCreateOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactCreateOrgContact) | **Post** /contacts | Add new entity to contacts
[**ContactsOrgContactDeleteOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactDeleteOrgContact) | **Delete** /contacts/{orgContact-id} | Delete entity from contacts
[**ContactsOrgContactGetOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactGetOrgContact) | **Get** /contacts/{orgContact-id} | Get entity from contacts by key
[**ContactsOrgContactListOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactListOrgContact) | **Get** /contacts | Get entities from contacts
[**ContactsOrgContactUpdateOrgContact**](ContactsOrgContactApi.md#ContactsOrgContactUpdateOrgContact) | **Patch** /contacts/{orgContact-id} | Update entity in contacts



## ContactsOrgContactCreateOrgContact

> MicrosoftGraphOrgContact ContactsOrgContactCreateOrgContact(ctx).MicrosoftGraphOrgContact(microsoftGraphOrgContact).Execute()

Add new entity to contacts

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
    microsoftGraphOrgContact := *openapiclient.NewMicrosoftGraphOrgContact() // MicrosoftGraphOrgContact | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsOrgContactApi.ContactsOrgContactCreateOrgContact(context.Background()).MicrosoftGraphOrgContact(microsoftGraphOrgContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsOrgContactApi.ContactsOrgContactCreateOrgContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactCreateOrgContact`: MicrosoftGraphOrgContact
    fmt.Fprintf(os.Stdout, "Response from `ContactsOrgContactApi.ContactsOrgContactCreateOrgContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactCreateOrgContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOrgContact** | [**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md) | New entity | 

### Return type

[**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactDeleteOrgContact

> ContactsOrgContactDeleteOrgContact(ctx, orgContactId).IfMatch(ifMatch).Execute()

Delete entity from contacts

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
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsOrgContactApi.ContactsOrgContactDeleteOrgContact(context.Background(), orgContactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsOrgContactApi.ContactsOrgContactDeleteOrgContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactDeleteOrgContactRequest struct via the builder pattern


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


## ContactsOrgContactGetOrgContact

> MicrosoftGraphOrgContact ContactsOrgContactGetOrgContact(ctx, orgContactId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()

Get entity from contacts by key

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
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsOrgContactApi.ContactsOrgContactGetOrgContact(context.Background(), orgContactId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsOrgContactApi.ContactsOrgContactGetOrgContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactGetOrgContact`: MicrosoftGraphOrgContact
    fmt.Fprintf(os.Stdout, "Response from `ContactsOrgContactApi.ContactsOrgContactGetOrgContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactGetOrgContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactListOrgContact

> CollectionOfOrgContact ContactsOrgContactListOrgContact(ctx).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from contacts

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
    resp, r, err := api_client.ContactsOrgContactApi.ContactsOrgContactListOrgContact(context.Background()).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsOrgContactApi.ContactsOrgContactListOrgContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactListOrgContact`: CollectionOfOrgContact
    fmt.Fprintf(os.Stdout, "Response from `ContactsOrgContactApi.ContactsOrgContactListOrgContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactListOrgContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfOrgContact**](CollectionOfOrgContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactUpdateOrgContact

> ContactsOrgContactUpdateOrgContact(ctx, orgContactId).MicrosoftGraphOrgContact(microsoftGraphOrgContact).Execute()

Update entity in contacts

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
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    microsoftGraphOrgContact := *openapiclient.NewMicrosoftGraphOrgContact() // MicrosoftGraphOrgContact | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsOrgContactApi.ContactsOrgContactUpdateOrgContact(context.Background(), orgContactId).MicrosoftGraphOrgContact(microsoftGraphOrgContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsOrgContactApi.ContactsOrgContactUpdateOrgContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactUpdateOrgContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOrgContact** | [**MicrosoftGraphOrgContact**](MicrosoftGraphOrgContact.md) | New property values | 

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

