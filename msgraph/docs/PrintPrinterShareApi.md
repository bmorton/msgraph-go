# \PrintPrinterShareApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreateShares**](PrintPrinterShareApi.md#PrintCreateShares) | **Post** /print/shares | Create new navigation property to shares for print
[**PrintDeleteShares**](PrintPrinterShareApi.md#PrintDeleteShares) | **Delete** /print/shares/{printerShare-id} | Delete navigation property shares for print
[**PrintGetShares**](PrintPrinterShareApi.md#PrintGetShares) | **Get** /print/shares/{printerShare-id} | Get shares from print
[**PrintListShares**](PrintPrinterShareApi.md#PrintListShares) | **Get** /print/shares | Get shares from print
[**PrintSharesCreateRefAllowedGroups**](PrintPrinterShareApi.md#PrintSharesCreateRefAllowedGroups) | **Post** /print/shares/{printerShare-id}/allowedGroups/$ref | Create new navigation property ref to allowedGroups for print
[**PrintSharesCreateRefAllowedUsers**](PrintPrinterShareApi.md#PrintSharesCreateRefAllowedUsers) | **Post** /print/shares/{printerShare-id}/allowedUsers/$ref | Create new navigation property ref to allowedUsers for print
[**PrintSharesDeleteRefPrinter**](PrintPrinterShareApi.md#PrintSharesDeleteRefPrinter) | **Delete** /print/shares/{printerShare-id}/printer/$ref | Delete ref of navigation property printer for print
[**PrintSharesGetPrinter**](PrintPrinterShareApi.md#PrintSharesGetPrinter) | **Get** /print/shares/{printerShare-id}/printer | Get printer from print
[**PrintSharesGetRefPrinter**](PrintPrinterShareApi.md#PrintSharesGetRefPrinter) | **Get** /print/shares/{printerShare-id}/printer/$ref | Get ref of printer from print
[**PrintSharesListAllowedGroups**](PrintPrinterShareApi.md#PrintSharesListAllowedGroups) | **Get** /print/shares/{printerShare-id}/allowedGroups | Get allowedGroups from print
[**PrintSharesListAllowedUsers**](PrintPrinterShareApi.md#PrintSharesListAllowedUsers) | **Get** /print/shares/{printerShare-id}/allowedUsers | Get allowedUsers from print
[**PrintSharesListRefAllowedGroups**](PrintPrinterShareApi.md#PrintSharesListRefAllowedGroups) | **Get** /print/shares/{printerShare-id}/allowedGroups/$ref | Get ref of allowedGroups from print
[**PrintSharesListRefAllowedUsers**](PrintPrinterShareApi.md#PrintSharesListRefAllowedUsers) | **Get** /print/shares/{printerShare-id}/allowedUsers/$ref | Get ref of allowedUsers from print
[**PrintSharesUpdateRefPrinter**](PrintPrinterShareApi.md#PrintSharesUpdateRefPrinter) | **Put** /print/shares/{printerShare-id}/printer/$ref | Update the ref of navigation property printer in print
[**PrintUpdateShares**](PrintPrinterShareApi.md#PrintUpdateShares) | **Patch** /print/shares/{printerShare-id} | Update the navigation property shares in print



## PrintCreateShares

> MicrosoftGraphPrinterShare PrintCreateShares(ctx).MicrosoftGraphPrinterShare(microsoftGraphPrinterShare).Execute()

Create new navigation property to shares for print



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
    microsoftGraphPrinterShare := *openapiclient.NewMicrosoftGraphPrinterShare() // MicrosoftGraphPrinterShare | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintCreateShares(context.Background()).MicrosoftGraphPrinterShare(microsoftGraphPrinterShare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintCreateShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreateShares`: MicrosoftGraphPrinterShare
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintCreateShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreateSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrinterShare** | [**MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeleteShares

> PrintDeleteShares(ctx, printerShareId).IfMatch(ifMatch).Execute()

Delete navigation property shares for print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintDeleteShares(context.Background(), printerShareId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintDeleteShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeleteSharesRequest struct via the builder pattern


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


## PrintGetShares

> MicrosoftGraphPrinterShare PrintGetShares(ctx, printerShareId).Select_(select_).Expand(expand).Execute()

Get shares from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintGetShares(context.Background(), printerShareId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintGetShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetShares`: MicrosoftGraphPrinterShare
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintGetShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintListShares

> CollectionOfPrinterShare PrintListShares(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get shares from print



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
    resp, r, err := api_client.PrintPrinterShareApi.PrintListShares(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintListShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListShares`: CollectionOfPrinterShare
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintListShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListSharesRequest struct via the builder pattern


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

[**CollectionOfPrinterShare**](CollectionOfPrinterShare.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesCreateRefAllowedGroups

> map[string]interface{} PrintSharesCreateRefAllowedGroups(ctx, printerShareId).RequestBody(requestBody).Execute()

Create new navigation property ref to allowedGroups for print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesCreateRefAllowedGroups(context.Background(), printerShareId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesCreateRefAllowedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesCreateRefAllowedGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesCreateRefAllowedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesCreateRefAllowedGroupsRequest struct via the builder pattern


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


## PrintSharesCreateRefAllowedUsers

> map[string]interface{} PrintSharesCreateRefAllowedUsers(ctx, printerShareId).RequestBody(requestBody).Execute()

Create new navigation property ref to allowedUsers for print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesCreateRefAllowedUsers(context.Background(), printerShareId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesCreateRefAllowedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesCreateRefAllowedUsers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesCreateRefAllowedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesCreateRefAllowedUsersRequest struct via the builder pattern


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


## PrintSharesDeleteRefPrinter

> PrintSharesDeleteRefPrinter(ctx, printerShareId).IfMatch(ifMatch).Execute()

Delete ref of navigation property printer for print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesDeleteRefPrinter(context.Background(), printerShareId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesDeleteRefPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesDeleteRefPrinterRequest struct via the builder pattern


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


## PrintSharesGetPrinter

> MicrosoftGraphPrinter PrintSharesGetPrinter(ctx, printerShareId).Select_(select_).Expand(expand).Execute()

Get printer from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesGetPrinter(context.Background(), printerShareId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesGetPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesGetPrinter`: MicrosoftGraphPrinter
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesGetPrinter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesGetPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrinter**](MicrosoftGraphPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesGetRefPrinter

> string PrintSharesGetRefPrinter(ctx, printerShareId).Execute()

Get ref of printer from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesGetRefPrinter(context.Background(), printerShareId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesGetRefPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesGetRefPrinter`: string
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesGetRefPrinter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesGetRefPrinterRequest struct via the builder pattern


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


## PrintSharesListAllowedGroups

> CollectionOfGroup PrintSharesListAllowedGroups(ctx, printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get allowedGroups from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
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
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesListAllowedGroups(context.Background(), printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesListAllowedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesListAllowedGroups`: CollectionOfGroup
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesListAllowedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesListAllowedGroupsRequest struct via the builder pattern


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

[**CollectionOfGroup**](CollectionOfGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesListAllowedUsers

> CollectionOfUser PrintSharesListAllowedUsers(ctx, printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get allowedUsers from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
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
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesListAllowedUsers(context.Background(), printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesListAllowedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesListAllowedUsers`: CollectionOfUser
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesListAllowedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesListAllowedUsersRequest struct via the builder pattern


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

[**CollectionOfUser**](CollectionOfUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesListRefAllowedGroups

> CollectionOfLinksOfGroup PrintSharesListRefAllowedGroups(ctx, printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of allowedGroups from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesListRefAllowedGroups(context.Background(), printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesListRefAllowedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesListRefAllowedGroups`: CollectionOfLinksOfGroup
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesListRefAllowedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesListRefAllowedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfGroup**](CollectionOfLinksOfGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesListRefAllowedUsers

> CollectionOfLinksOfUser PrintSharesListRefAllowedUsers(ctx, printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of allowedUsers from print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesListRefAllowedUsers(context.Background(), printerShareId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesListRefAllowedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintSharesListRefAllowedUsers`: CollectionOfLinksOfUser
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterShareApi.PrintSharesListRefAllowedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesListRefAllowedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfUser**](CollectionOfLinksOfUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintSharesUpdateRefPrinter

> PrintSharesUpdateRefPrinter(ctx, printerShareId).RequestBody(requestBody).Execute()

Update the ref of navigation property printer in print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintSharesUpdateRefPrinter(context.Background(), printerShareId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintSharesUpdateRefPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintSharesUpdateRefPrinterRequest struct via the builder pattern


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


## PrintUpdateShares

> PrintUpdateShares(ctx, printerShareId).MicrosoftGraphPrinterShare(microsoftGraphPrinterShare).Execute()

Update the navigation property shares in print



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
    printerShareId := "printerShareId_example" // string | key: id of printerShare
    microsoftGraphPrinterShare := *openapiclient.NewMicrosoftGraphPrinterShare() // MicrosoftGraphPrinterShare | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterShareApi.PrintUpdateShares(context.Background(), printerShareId).MicrosoftGraphPrinterShare(microsoftGraphPrinterShare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterShareApi.PrintUpdateShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerShareId** | **string** | key: id of printerShare | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdateSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrinterShare** | [**MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md) | New navigation property values | 

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

