# \ContactsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsCreateRefDirectReports**](ContactsDirectoryObjectApi.md#ContactsCreateRefDirectReports) | **Post** /contacts/{orgContact-id}/directReports/$ref | Create new navigation property ref to directReports for contacts
[**ContactsCreateRefMemberOf**](ContactsDirectoryObjectApi.md#ContactsCreateRefMemberOf) | **Post** /contacts/{orgContact-id}/memberOf/$ref | Create new navigation property ref to memberOf for contacts
[**ContactsCreateRefTransitiveMemberOf**](ContactsDirectoryObjectApi.md#ContactsCreateRefTransitiveMemberOf) | **Post** /contacts/{orgContact-id}/transitiveMemberOf/$ref | Create new navigation property ref to transitiveMemberOf for contacts
[**ContactsDeleteRefManager**](ContactsDirectoryObjectApi.md#ContactsDeleteRefManager) | **Delete** /contacts/{orgContact-id}/manager/$ref | Delete ref of navigation property manager for contacts
[**ContactsGetManager**](ContactsDirectoryObjectApi.md#ContactsGetManager) | **Get** /contacts/{orgContact-id}/manager | Get manager from contacts
[**ContactsGetRefManager**](ContactsDirectoryObjectApi.md#ContactsGetRefManager) | **Get** /contacts/{orgContact-id}/manager/$ref | Get ref of manager from contacts
[**ContactsListDirectReports**](ContactsDirectoryObjectApi.md#ContactsListDirectReports) | **Get** /contacts/{orgContact-id}/directReports | Get directReports from contacts
[**ContactsListMemberOf**](ContactsDirectoryObjectApi.md#ContactsListMemberOf) | **Get** /contacts/{orgContact-id}/memberOf | Get memberOf from contacts
[**ContactsListRefDirectReports**](ContactsDirectoryObjectApi.md#ContactsListRefDirectReports) | **Get** /contacts/{orgContact-id}/directReports/$ref | Get ref of directReports from contacts
[**ContactsListRefMemberOf**](ContactsDirectoryObjectApi.md#ContactsListRefMemberOf) | **Get** /contacts/{orgContact-id}/memberOf/$ref | Get ref of memberOf from contacts
[**ContactsListRefTransitiveMemberOf**](ContactsDirectoryObjectApi.md#ContactsListRefTransitiveMemberOf) | **Get** /contacts/{orgContact-id}/transitiveMemberOf/$ref | Get ref of transitiveMemberOf from contacts
[**ContactsListTransitiveMemberOf**](ContactsDirectoryObjectApi.md#ContactsListTransitiveMemberOf) | **Get** /contacts/{orgContact-id}/transitiveMemberOf | Get transitiveMemberOf from contacts
[**ContactsUpdateRefManager**](ContactsDirectoryObjectApi.md#ContactsUpdateRefManager) | **Put** /contacts/{orgContact-id}/manager/$ref | Update the ref of navigation property manager in contacts



## ContactsCreateRefDirectReports

> map[string]interface{} ContactsCreateRefDirectReports(ctx, orgContactId).RequestBody(requestBody).Execute()

Create new navigation property ref to directReports for contacts



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsCreateRefDirectReports(context.Background(), orgContactId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsCreateRefDirectReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsCreateRefDirectReports`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsCreateRefDirectReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsCreateRefDirectReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsCreateRefMemberOf

> map[string]interface{} ContactsCreateRefMemberOf(ctx, orgContactId).RequestBody(requestBody).Execute()

Create new navigation property ref to memberOf for contacts



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsCreateRefMemberOf(context.Background(), orgContactId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsCreateRefMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsCreateRefMemberOf`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsCreateRefMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsCreateRefMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsCreateRefTransitiveMemberOf

> map[string]interface{} ContactsCreateRefTransitiveMemberOf(ctx, orgContactId).RequestBody(requestBody).Execute()

Create new navigation property ref to transitiveMemberOf for contacts

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsCreateRefTransitiveMemberOf(context.Background(), orgContactId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsCreateRefTransitiveMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsCreateRefTransitiveMemberOf`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsCreateRefTransitiveMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsCreateRefTransitiveMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsDeleteRefManager

> ContactsDeleteRefManager(ctx, orgContactId).IfMatch(ifMatch).Execute()

Delete ref of navigation property manager for contacts



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
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsDeleteRefManager(context.Background(), orgContactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsDeleteRefManager``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContactsDeleteRefManagerRequest struct via the builder pattern


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


## ContactsGetManager

> MicrosoftGraphDirectoryObject ContactsGetManager(ctx, orgContactId).Select_(select_).Expand(expand).Execute()

Get manager from contacts



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsGetManager(context.Background(), orgContactId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsGetManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsGetManager`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsGetManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsGetManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsGetRefManager

> string ContactsGetRefManager(ctx, orgContactId).Execute()

Get ref of manager from contacts



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsGetRefManager(context.Background(), orgContactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsGetRefManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsGetRefManager`: string
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsGetRefManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsGetRefManagerRequest struct via the builder pattern


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


## ContactsListDirectReports

> CollectionOfDirectoryObject ContactsListDirectReports(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get directReports from contacts



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
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListDirectReports(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListDirectReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListDirectReports`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListDirectReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListDirectReportsRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListMemberOf

> CollectionOfDirectoryObject ContactsListMemberOf(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get memberOf from contacts



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
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListMemberOf(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListMemberOf`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListMemberOfRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListRefDirectReports

> CollectionOfLinksOfDirectoryObject ContactsListRefDirectReports(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of directReports from contacts



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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListRefDirectReports(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListRefDirectReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListRefDirectReports`: CollectionOfLinksOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListRefDirectReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListRefDirectReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfDirectoryObject**](CollectionOfLinksOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListRefMemberOf

> CollectionOfLinksOfDirectoryObject ContactsListRefMemberOf(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of memberOf from contacts



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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListRefMemberOf(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListRefMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListRefMemberOf`: CollectionOfLinksOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListRefMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListRefMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfDirectoryObject**](CollectionOfLinksOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListRefTransitiveMemberOf

> CollectionOfLinksOfDirectoryObject ContactsListRefTransitiveMemberOf(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of transitiveMemberOf from contacts

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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListRefTransitiveMemberOf(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListRefTransitiveMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListRefTransitiveMemberOf`: CollectionOfLinksOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListRefTransitiveMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListRefTransitiveMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfDirectoryObject**](CollectionOfLinksOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsListTransitiveMemberOf

> CollectionOfDirectoryObject ContactsListTransitiveMemberOf(ctx, orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get transitiveMemberOf from contacts

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
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsListTransitiveMemberOf(context.Background(), orgContactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsListTransitiveMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsListTransitiveMemberOf`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsDirectoryObjectApi.ContactsListTransitiveMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsListTransitiveMemberOfRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsUpdateRefManager

> ContactsUpdateRefManager(ctx, orgContactId).RequestBody(requestBody).Execute()

Update the ref of navigation property manager in contacts



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsDirectoryObjectApi.ContactsUpdateRefManager(context.Background(), orgContactId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsDirectoryObjectApi.ContactsUpdateRefManager``: %v\n", err)
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

Other parameters are passed through a pointer to a apiContactsUpdateRefManagerRequest struct via the builder pattern


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

