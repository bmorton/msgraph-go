# \ReportsPrintUsageByPrinterApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsCreateDailyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsCreateDailyPrintUsageByPrinter) | **Post** /reports/dailyPrintUsageByPrinter | Create new navigation property to dailyPrintUsageByPrinter for reports
[**ReportsCreateMonthlyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsCreateMonthlyPrintUsageByPrinter) | **Post** /reports/monthlyPrintUsageByPrinter | Create new navigation property to monthlyPrintUsageByPrinter for reports
[**ReportsDeleteDailyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsDeleteDailyPrintUsageByPrinter) | **Delete** /reports/dailyPrintUsageByPrinter/{printUsageByPrinter-id} | Delete navigation property dailyPrintUsageByPrinter for reports
[**ReportsDeleteMonthlyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsDeleteMonthlyPrintUsageByPrinter) | **Delete** /reports/monthlyPrintUsageByPrinter/{printUsageByPrinter-id} | Delete navigation property monthlyPrintUsageByPrinter for reports
[**ReportsGetDailyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsGetDailyPrintUsageByPrinter) | **Get** /reports/dailyPrintUsageByPrinter/{printUsageByPrinter-id} | Get dailyPrintUsageByPrinter from reports
[**ReportsGetMonthlyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsGetMonthlyPrintUsageByPrinter) | **Get** /reports/monthlyPrintUsageByPrinter/{printUsageByPrinter-id} | Get monthlyPrintUsageByPrinter from reports
[**ReportsListDailyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsListDailyPrintUsageByPrinter) | **Get** /reports/dailyPrintUsageByPrinter | Get dailyPrintUsageByPrinter from reports
[**ReportsListMonthlyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsListMonthlyPrintUsageByPrinter) | **Get** /reports/monthlyPrintUsageByPrinter | Get monthlyPrintUsageByPrinter from reports
[**ReportsUpdateDailyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsUpdateDailyPrintUsageByPrinter) | **Patch** /reports/dailyPrintUsageByPrinter/{printUsageByPrinter-id} | Update the navigation property dailyPrintUsageByPrinter in reports
[**ReportsUpdateMonthlyPrintUsageByPrinter**](ReportsPrintUsageByPrinterApi.md#ReportsUpdateMonthlyPrintUsageByPrinter) | **Patch** /reports/monthlyPrintUsageByPrinter/{printUsageByPrinter-id} | Update the navigation property monthlyPrintUsageByPrinter in reports



## ReportsCreateDailyPrintUsageByPrinter

> MicrosoftGraphPrintUsageByPrinter ReportsCreateDailyPrintUsageByPrinter(ctx).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()

Create new navigation property to dailyPrintUsageByPrinter for reports

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
    microsoftGraphPrintUsageByPrinter := *openapiclient.NewMicrosoftGraphPrintUsageByPrinter() // MicrosoftGraphPrintUsageByPrinter | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsCreateDailyPrintUsageByPrinter(context.Background()).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsCreateDailyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCreateDailyPrintUsageByPrinter`: MicrosoftGraphPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsCreateDailyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCreateDailyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintUsageByPrinter** | [**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsCreateMonthlyPrintUsageByPrinter

> MicrosoftGraphPrintUsageByPrinter ReportsCreateMonthlyPrintUsageByPrinter(ctx).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()

Create new navigation property to monthlyPrintUsageByPrinter for reports

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
    microsoftGraphPrintUsageByPrinter := *openapiclient.NewMicrosoftGraphPrintUsageByPrinter() // MicrosoftGraphPrintUsageByPrinter | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsCreateMonthlyPrintUsageByPrinter(context.Background()).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsCreateMonthlyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCreateMonthlyPrintUsageByPrinter`: MicrosoftGraphPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsCreateMonthlyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCreateMonthlyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintUsageByPrinter** | [**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsDeleteDailyPrintUsageByPrinter

> ReportsDeleteDailyPrintUsageByPrinter(ctx, printUsageByPrinterId).IfMatch(ifMatch).Execute()

Delete navigation property dailyPrintUsageByPrinter for reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsDeleteDailyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsDeleteDailyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDeleteDailyPrintUsageByPrinterRequest struct via the builder pattern


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


## ReportsDeleteMonthlyPrintUsageByPrinter

> ReportsDeleteMonthlyPrintUsageByPrinter(ctx, printUsageByPrinterId).IfMatch(ifMatch).Execute()

Delete navigation property monthlyPrintUsageByPrinter for reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsDeleteMonthlyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsDeleteMonthlyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDeleteMonthlyPrintUsageByPrinterRequest struct via the builder pattern


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


## ReportsGetDailyPrintUsageByPrinter

> MicrosoftGraphPrintUsageByPrinter ReportsGetDailyPrintUsageByPrinter(ctx, printUsageByPrinterId).Select_(select_).Expand(expand).Execute()

Get dailyPrintUsageByPrinter from reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsGetDailyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsGetDailyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetDailyPrintUsageByPrinter`: MicrosoftGraphPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsGetDailyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetDailyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGetMonthlyPrintUsageByPrinter

> MicrosoftGraphPrintUsageByPrinter ReportsGetMonthlyPrintUsageByPrinter(ctx, printUsageByPrinterId).Select_(select_).Expand(expand).Execute()

Get monthlyPrintUsageByPrinter from reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsGetMonthlyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsGetMonthlyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetMonthlyPrintUsageByPrinter`: MicrosoftGraphPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsGetMonthlyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetMonthlyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsListDailyPrintUsageByPrinter

> CollectionOfPrintUsageByPrinter ReportsListDailyPrintUsageByPrinter(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get dailyPrintUsageByPrinter from reports

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
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsListDailyPrintUsageByPrinter(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsListDailyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsListDailyPrintUsageByPrinter`: CollectionOfPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsListDailyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsListDailyPrintUsageByPrinterRequest struct via the builder pattern


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

[**CollectionOfPrintUsageByPrinter**](CollectionOfPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsListMonthlyPrintUsageByPrinter

> CollectionOfPrintUsageByPrinter ReportsListMonthlyPrintUsageByPrinter(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get monthlyPrintUsageByPrinter from reports

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
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsListMonthlyPrintUsageByPrinter(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsListMonthlyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsListMonthlyPrintUsageByPrinter`: CollectionOfPrintUsageByPrinter
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByPrinterApi.ReportsListMonthlyPrintUsageByPrinter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsListMonthlyPrintUsageByPrinterRequest struct via the builder pattern


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

[**CollectionOfPrintUsageByPrinter**](CollectionOfPrintUsageByPrinter.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUpdateDailyPrintUsageByPrinter

> ReportsUpdateDailyPrintUsageByPrinter(ctx, printUsageByPrinterId).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()

Update the navigation property dailyPrintUsageByPrinter in reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    microsoftGraphPrintUsageByPrinter := *openapiclient.NewMicrosoftGraphPrintUsageByPrinter() // MicrosoftGraphPrintUsageByPrinter | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsUpdateDailyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsUpdateDailyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateDailyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintUsageByPrinter** | [**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) | New navigation property values | 

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


## ReportsUpdateMonthlyPrintUsageByPrinter

> ReportsUpdateMonthlyPrintUsageByPrinter(ctx, printUsageByPrinterId).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()

Update the navigation property monthlyPrintUsageByPrinter in reports

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
    printUsageByPrinterId := "printUsageByPrinterId_example" // string | key: id of printUsageByPrinter
    microsoftGraphPrintUsageByPrinter := *openapiclient.NewMicrosoftGraphPrintUsageByPrinter() // MicrosoftGraphPrintUsageByPrinter | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByPrinterApi.ReportsUpdateMonthlyPrintUsageByPrinter(context.Background(), printUsageByPrinterId).MicrosoftGraphPrintUsageByPrinter(microsoftGraphPrintUsageByPrinter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByPrinterApi.ReportsUpdateMonthlyPrintUsageByPrinter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByPrinterId** | **string** | key: id of printUsageByPrinter | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateMonthlyPrintUsageByPrinterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintUsageByPrinter** | [**MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) | New navigation property values | 

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

