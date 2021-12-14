# \PrintPrintConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreateConnectors**](PrintPrintConnectorApi.md#PrintCreateConnectors) | **Post** /print/connectors | Create new navigation property to connectors for print
[**PrintDeleteConnectors**](PrintPrintConnectorApi.md#PrintDeleteConnectors) | **Delete** /print/connectors/{printConnector-id} | Delete navigation property connectors for print
[**PrintGetConnectors**](PrintPrintConnectorApi.md#PrintGetConnectors) | **Get** /print/connectors/{printConnector-id} | Get connectors from print
[**PrintListConnectors**](PrintPrintConnectorApi.md#PrintListConnectors) | **Get** /print/connectors | Get connectors from print
[**PrintUpdateConnectors**](PrintPrintConnectorApi.md#PrintUpdateConnectors) | **Patch** /print/connectors/{printConnector-id} | Update the navigation property connectors in print



## PrintCreateConnectors

> MicrosoftGraphPrintConnector PrintCreateConnectors(ctx).MicrosoftGraphPrintConnector(microsoftGraphPrintConnector).Execute()

Create new navigation property to connectors for print



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
    microsoftGraphPrintConnector := *openapiclient.NewMicrosoftGraphPrintConnector() // MicrosoftGraphPrintConnector | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintConnectorApi.PrintCreateConnectors(context.Background()).MicrosoftGraphPrintConnector(microsoftGraphPrintConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintConnectorApi.PrintCreateConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreateConnectors`: MicrosoftGraphPrintConnector
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintConnectorApi.PrintCreateConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreateConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintConnector** | [**MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeleteConnectors

> PrintDeleteConnectors(ctx, printConnectorId).IfMatch(ifMatch).Execute()

Delete navigation property connectors for print



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
    printConnectorId := "printConnectorId_example" // string | key: id of printConnector
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintConnectorApi.PrintDeleteConnectors(context.Background(), printConnectorId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintConnectorApi.PrintDeleteConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printConnectorId** | **string** | key: id of printConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeleteConnectorsRequest struct via the builder pattern


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


## PrintGetConnectors

> MicrosoftGraphPrintConnector PrintGetConnectors(ctx, printConnectorId).Select_(select_).Expand(expand).Execute()

Get connectors from print



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
    printConnectorId := "printConnectorId_example" // string | key: id of printConnector
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintConnectorApi.PrintGetConnectors(context.Background(), printConnectorId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintConnectorApi.PrintGetConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetConnectors`: MicrosoftGraphPrintConnector
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintConnectorApi.PrintGetConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printConnectorId** | **string** | key: id of printConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintListConnectors

> CollectionOfPrintConnector PrintListConnectors(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get connectors from print



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
    resp, r, err := api_client.PrintPrintConnectorApi.PrintListConnectors(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintConnectorApi.PrintListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListConnectors`: CollectionOfPrintConnector
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintConnectorApi.PrintListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListConnectorsRequest struct via the builder pattern


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

[**CollectionOfPrintConnector**](CollectionOfPrintConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintUpdateConnectors

> PrintUpdateConnectors(ctx, printConnectorId).MicrosoftGraphPrintConnector(microsoftGraphPrintConnector).Execute()

Update the navigation property connectors in print



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
    printConnectorId := "printConnectorId_example" // string | key: id of printConnector
    microsoftGraphPrintConnector := *openapiclient.NewMicrosoftGraphPrintConnector() // MicrosoftGraphPrintConnector | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintConnectorApi.PrintUpdateConnectors(context.Background(), printConnectorId).MicrosoftGraphPrintConnector(microsoftGraphPrintConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintConnectorApi.PrintUpdateConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printConnectorId** | **string** | key: id of printConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdateConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintConnector** | [**MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md) | New navigation property values | 

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

