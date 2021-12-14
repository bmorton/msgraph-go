# \PrintPrinterApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreatePrinters**](PrintPrinterApi.md#PrintCreatePrinters) | **Post** /print/printers | Create new navigation property to printers for print
[**PrintDeletePrinters**](PrintPrinterApi.md#PrintDeletePrinters) | **Delete** /print/printers/{printer-id} | Delete navigation property printers for print
[**PrintGetPrinters**](PrintPrinterApi.md#PrintGetPrinters) | **Get** /print/printers/{printer-id} | Get printers from print
[**PrintListPrinters**](PrintPrinterApi.md#PrintListPrinters) | **Get** /print/printers | Get printers from print
[**PrintPrintersCreateRefConnectors**](PrintPrinterApi.md#PrintPrintersCreateRefConnectors) | **Post** /print/printers/{printer-id}/connectors/$ref | Create new navigation property ref to connectors for print
[**PrintPrintersCreateRefShares**](PrintPrinterApi.md#PrintPrintersCreateRefShares) | **Post** /print/printers/{printer-id}/shares/$ref | Create new navigation property ref to shares for print
[**PrintPrintersCreateTaskTriggers**](PrintPrinterApi.md#PrintPrintersCreateTaskTriggers) | **Post** /print/printers/{printer-id}/taskTriggers | Create new navigation property to taskTriggers for print
[**PrintPrintersDeleteTaskTriggers**](PrintPrinterApi.md#PrintPrintersDeleteTaskTriggers) | **Delete** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id} | Delete navigation property taskTriggers for print
[**PrintPrintersGetTaskTriggers**](PrintPrinterApi.md#PrintPrintersGetTaskTriggers) | **Get** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id} | Get taskTriggers from print
[**PrintPrintersListConnectors**](PrintPrinterApi.md#PrintPrintersListConnectors) | **Get** /print/printers/{printer-id}/connectors | Get connectors from print
[**PrintPrintersListRefConnectors**](PrintPrinterApi.md#PrintPrintersListRefConnectors) | **Get** /print/printers/{printer-id}/connectors/$ref | Get ref of connectors from print
[**PrintPrintersListRefShares**](PrintPrinterApi.md#PrintPrintersListRefShares) | **Get** /print/printers/{printer-id}/shares/$ref | Get ref of shares from print
[**PrintPrintersListShares**](PrintPrinterApi.md#PrintPrintersListShares) | **Get** /print/printers/{printer-id}/shares | Get shares from print
[**PrintPrintersListTaskTriggers**](PrintPrinterApi.md#PrintPrintersListTaskTriggers) | **Get** /print/printers/{printer-id}/taskTriggers | Get taskTriggers from print
[**PrintPrintersTaskTriggersDeleteRefDefinition**](PrintPrinterApi.md#PrintPrintersTaskTriggersDeleteRefDefinition) | **Delete** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id}/definition/$ref | Delete ref of navigation property definition for print
[**PrintPrintersTaskTriggersGetDefinition**](PrintPrinterApi.md#PrintPrintersTaskTriggersGetDefinition) | **Get** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id}/definition | Get definition from print
[**PrintPrintersTaskTriggersGetRefDefinition**](PrintPrinterApi.md#PrintPrintersTaskTriggersGetRefDefinition) | **Get** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id}/definition/$ref | Get ref of definition from print
[**PrintPrintersTaskTriggersUpdateRefDefinition**](PrintPrinterApi.md#PrintPrintersTaskTriggersUpdateRefDefinition) | **Put** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id}/definition/$ref | Update the ref of navigation property definition in print
[**PrintPrintersUpdateTaskTriggers**](PrintPrinterApi.md#PrintPrintersUpdateTaskTriggers) | **Patch** /print/printers/{printer-id}/taskTriggers/{printTaskTrigger-id} | Update the navigation property taskTriggers in print
[**PrintUpdatePrinters**](PrintPrinterApi.md#PrintUpdatePrinters) | **Patch** /print/printers/{printer-id} | Update the navigation property printers in print



## PrintCreatePrinters

> MicrosoftGraphPrinter PrintCreatePrinters(ctx).MicrosoftGraphPrinter(microsoftGraphPrinter).Execute()

Create new navigation property to printers for print



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
    microsoftGraphPrinter := *openapiclient.NewMicrosoftGraphPrinter() // MicrosoftGraphPrinter | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintCreatePrinters(context.Background()).MicrosoftGraphPrinter(microsoftGraphPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintCreatePrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreatePrinters`: MicrosoftGraphPrinter
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintCreatePrinters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreatePrintersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrinter** | [**MicrosoftGraphPrinter**](MicrosoftGraphPrinter.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrinter**](MicrosoftGraphPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeletePrinters

> PrintDeletePrinters(ctx, printerId).IfMatch(ifMatch).Execute()

Delete navigation property printers for print



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
    printerId := "printerId_example" // string | key: id of printer
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintDeletePrinters(context.Background(), printerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintDeletePrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeletePrintersRequest struct via the builder pattern


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


## PrintGetPrinters

> MicrosoftGraphPrinter PrintGetPrinters(ctx, printerId).Select_(select_).Expand(expand).Execute()

Get printers from print



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
    printerId := "printerId_example" // string | key: id of printer
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintGetPrinters(context.Background(), printerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintGetPrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetPrinters`: MicrosoftGraphPrinter
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintGetPrinters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetPrintersRequest struct via the builder pattern


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


## PrintListPrinters

> CollectionOfPrinter PrintListPrinters(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get printers from print



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
    resp, r, err := api_client.PrintPrinterApi.PrintListPrinters(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintListPrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListPrinters`: CollectionOfPrinter
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintListPrinters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListPrintersRequest struct via the builder pattern


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

[**CollectionOfPrinter**](CollectionOfPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersCreateRefConnectors

> map[string]interface{} PrintPrintersCreateRefConnectors(ctx, printerId).RequestBody(requestBody).Execute()

Create new navigation property ref to connectors for print



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
    printerId := "printerId_example" // string | key: id of printer
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersCreateRefConnectors(context.Background(), printerId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersCreateRefConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersCreateRefConnectors`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersCreateRefConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersCreateRefConnectorsRequest struct via the builder pattern


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


## PrintPrintersCreateRefShares

> map[string]interface{} PrintPrintersCreateRefShares(ctx, printerId).RequestBody(requestBody).Execute()

Create new navigation property ref to shares for print



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
    printerId := "printerId_example" // string | key: id of printer
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersCreateRefShares(context.Background(), printerId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersCreateRefShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersCreateRefShares`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersCreateRefShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersCreateRefSharesRequest struct via the builder pattern


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


## PrintPrintersCreateTaskTriggers

> MicrosoftGraphPrintTaskTrigger PrintPrintersCreateTaskTriggers(ctx, printerId).MicrosoftGraphPrintTaskTrigger(microsoftGraphPrintTaskTrigger).Execute()

Create new navigation property to taskTriggers for print



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
    printerId := "printerId_example" // string | key: id of printer
    microsoftGraphPrintTaskTrigger := *openapiclient.NewMicrosoftGraphPrintTaskTrigger() // MicrosoftGraphPrintTaskTrigger | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersCreateTaskTriggers(context.Background(), printerId).MicrosoftGraphPrintTaskTrigger(microsoftGraphPrintTaskTrigger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersCreateTaskTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersCreateTaskTriggers`: MicrosoftGraphPrintTaskTrigger
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersCreateTaskTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersCreateTaskTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintTaskTrigger** | [**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersDeleteTaskTriggers

> PrintPrintersDeleteTaskTriggers(ctx, printerId, printTaskTriggerId).IfMatch(ifMatch).Execute()

Delete navigation property taskTriggers for print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersDeleteTaskTriggers(context.Background(), printerId, printTaskTriggerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersDeleteTaskTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersDeleteTaskTriggersRequest struct via the builder pattern


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


## PrintPrintersGetTaskTriggers

> MicrosoftGraphPrintTaskTrigger PrintPrintersGetTaskTriggers(ctx, printerId, printTaskTriggerId).Select_(select_).Expand(expand).Execute()

Get taskTriggers from print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersGetTaskTriggers(context.Background(), printerId, printTaskTriggerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersGetTaskTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersGetTaskTriggers`: MicrosoftGraphPrintTaskTrigger
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersGetTaskTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersGetTaskTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersListConnectors

> CollectionOfPrintConnector PrintPrintersListConnectors(ctx, printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    printerId := "printerId_example" // string | key: id of printer
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
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersListConnectors(context.Background(), printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersListConnectors`: CollectionOfPrintConnector
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersListConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersListConnectorsRequest struct via the builder pattern


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


## PrintPrintersListRefConnectors

> CollectionOfLinksOfPrintConnector PrintPrintersListRefConnectors(ctx, printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of connectors from print



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
    printerId := "printerId_example" // string | key: id of printer
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersListRefConnectors(context.Background(), printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersListRefConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersListRefConnectors`: CollectionOfLinksOfPrintConnector
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersListRefConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersListRefConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfPrintConnector**](CollectionOfLinksOfPrintConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersListRefShares

> CollectionOfLinksOfPrinterShare PrintPrintersListRefShares(ctx, printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of shares from print



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
    printerId := "printerId_example" // string | key: id of printer
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersListRefShares(context.Background(), printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersListRefShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersListRefShares`: CollectionOfLinksOfPrinterShare
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersListRefShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersListRefSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfPrinterShare**](CollectionOfLinksOfPrinterShare.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersListShares

> CollectionOfPrinterShare PrintPrintersListShares(ctx, printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    printerId := "printerId_example" // string | key: id of printer
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
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersListShares(context.Background(), printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersListShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersListShares`: CollectionOfPrinterShare
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersListShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersListSharesRequest struct via the builder pattern


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


## PrintPrintersListTaskTriggers

> CollectionOfPrintTaskTrigger PrintPrintersListTaskTriggers(ctx, printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get taskTriggers from print



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
    printerId := "printerId_example" // string | key: id of printer
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
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersListTaskTriggers(context.Background(), printerId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersListTaskTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersListTaskTriggers`: CollectionOfPrintTaskTrigger
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersListTaskTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersListTaskTriggersRequest struct via the builder pattern


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

[**CollectionOfPrintTaskTrigger**](CollectionOfPrintTaskTrigger.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersTaskTriggersDeleteRefDefinition

> PrintPrintersTaskTriggersDeleteRefDefinition(ctx, printerId, printTaskTriggerId).IfMatch(ifMatch).Execute()

Delete ref of navigation property definition for print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersTaskTriggersDeleteRefDefinition(context.Background(), printerId, printTaskTriggerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersTaskTriggersDeleteRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersTaskTriggersDeleteRefDefinitionRequest struct via the builder pattern


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


## PrintPrintersTaskTriggersGetDefinition

> MicrosoftGraphPrintTaskDefinition PrintPrintersTaskTriggersGetDefinition(ctx, printerId, printTaskTriggerId).Select_(select_).Expand(expand).Execute()

Get definition from print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersTaskTriggersGetDefinition(context.Background(), printerId, printTaskTriggerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersTaskTriggersGetDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersTaskTriggersGetDefinition`: MicrosoftGraphPrintTaskDefinition
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersTaskTriggersGetDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersTaskTriggersGetDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintPrintersTaskTriggersGetRefDefinition

> string PrintPrintersTaskTriggersGetRefDefinition(ctx, printerId, printTaskTriggerId).Execute()

Get ref of definition from print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersTaskTriggersGetRefDefinition(context.Background(), printerId, printTaskTriggerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersTaskTriggersGetRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintPrintersTaskTriggersGetRefDefinition`: string
    fmt.Fprintf(os.Stdout, "Response from `PrintPrinterApi.PrintPrintersTaskTriggersGetRefDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersTaskTriggersGetRefDefinitionRequest struct via the builder pattern


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


## PrintPrintersTaskTriggersUpdateRefDefinition

> PrintPrintersTaskTriggersUpdateRefDefinition(ctx, printerId, printTaskTriggerId).RequestBody(requestBody).Execute()

Update the ref of navigation property definition in print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersTaskTriggersUpdateRefDefinition(context.Background(), printerId, printTaskTriggerId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersTaskTriggersUpdateRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersTaskTriggersUpdateRefDefinitionRequest struct via the builder pattern


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


## PrintPrintersUpdateTaskTriggers

> PrintPrintersUpdateTaskTriggers(ctx, printerId, printTaskTriggerId).MicrosoftGraphPrintTaskTrigger(microsoftGraphPrintTaskTrigger).Execute()

Update the navigation property taskTriggers in print



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
    printerId := "printerId_example" // string | key: id of printer
    printTaskTriggerId := "printTaskTriggerId_example" // string | key: id of printTaskTrigger
    microsoftGraphPrintTaskTrigger := *openapiclient.NewMicrosoftGraphPrintTaskTrigger() // MicrosoftGraphPrintTaskTrigger | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintPrintersUpdateTaskTriggers(context.Background(), printerId, printTaskTriggerId).MicrosoftGraphPrintTaskTrigger(microsoftGraphPrintTaskTrigger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintPrintersUpdateTaskTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 
**printTaskTriggerId** | **string** | key: id of printTaskTrigger | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersUpdateTaskTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPrintTaskTrigger** | [**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) | New navigation property values | 

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


## PrintUpdatePrinters

> PrintUpdatePrinters(ctx, printerId).MicrosoftGraphPrinter(microsoftGraphPrinter).Execute()

Update the navigation property printers in print



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
    printerId := "printerId_example" // string | key: id of printer
    microsoftGraphPrinter := *openapiclient.NewMicrosoftGraphPrinter() // MicrosoftGraphPrinter | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrinterApi.PrintUpdatePrinters(context.Background(), printerId).MicrosoftGraphPrinter(microsoftGraphPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrinterApi.PrintUpdatePrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerId** | **string** | key: id of printer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdatePrintersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrinter** | [**MicrosoftGraphPrinter**](MicrosoftGraphPrinter.md) | New navigation property values | 

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

