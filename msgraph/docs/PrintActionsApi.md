# \PrintActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintPrintersCreate**](PrintActionsApi.md#PrintPrintersCreate) | **Post** /print/printers/microsoft.graph.create | Invoke action create
[**PrintPrintersPrinterRestoreFactoryDefaults**](PrintActionsApi.md#PrintPrintersPrinterRestoreFactoryDefaults) | **Post** /print/printers/{printer-id}/microsoft.graph.restoreFactoryDefaults | Invoke action restoreFactoryDefaults
[**PrintSharesPrinterSharePrinterRestoreFactoryDefaults**](PrintActionsApi.md#PrintSharesPrinterSharePrinterRestoreFactoryDefaults) | **Post** /print/shares/{printerShare-id}/printer/microsoft.graph.restoreFactoryDefaults | Invoke action restoreFactoryDefaults



## PrintPrintersCreate

> PrintPrintersCreate(ctx).InlineObject713(inlineObject713).Execute()

Invoke action create

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
    inlineObject713 := *openapiclient.NewInlineObject713() // InlineObject713 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintActionsApi.PrintPrintersCreate(context.Background()).InlineObject713(inlineObject713).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintActionsApi.PrintPrintersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintPrintersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject713** | [**InlineObject713**](InlineObject713.md) |  | 

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


## PrintPrintersPrinterRestoreFactoryDefaults

> PrintPrintersPrinterRestoreFactoryDefaults(ctx, printerId).Execute()

Invoke action restoreFactoryDefaults

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintActionsApi.PrintPrintersPrinterRestoreFactoryDefaults(context.Background(), printerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintActionsApi.PrintPrintersPrinterRestoreFactoryDefaults``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPrintPrintersPrinterRestoreFactoryDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PrintSharesPrinterSharePrinterRestoreFactoryDefaults

> PrintSharesPrinterSharePrinterRestoreFactoryDefaults(ctx, printerShareId).Execute()

Invoke action restoreFactoryDefaults

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
    resp, r, err := api_client.PrintActionsApi.PrintSharesPrinterSharePrinterRestoreFactoryDefaults(context.Background(), printerShareId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintActionsApi.PrintSharesPrinterSharePrinterRestoreFactoryDefaults``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

