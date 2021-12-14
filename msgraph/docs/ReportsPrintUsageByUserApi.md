# \ReportsPrintUsageByUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsCreateDailyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsCreateDailyPrintUsageByUser) | **Post** /reports/dailyPrintUsageByUser | Create new navigation property to dailyPrintUsageByUser for reports
[**ReportsCreateMonthlyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsCreateMonthlyPrintUsageByUser) | **Post** /reports/monthlyPrintUsageByUser | Create new navigation property to monthlyPrintUsageByUser for reports
[**ReportsDeleteDailyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsDeleteDailyPrintUsageByUser) | **Delete** /reports/dailyPrintUsageByUser/{printUsageByUser-id} | Delete navigation property dailyPrintUsageByUser for reports
[**ReportsDeleteMonthlyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsDeleteMonthlyPrintUsageByUser) | **Delete** /reports/monthlyPrintUsageByUser/{printUsageByUser-id} | Delete navigation property monthlyPrintUsageByUser for reports
[**ReportsGetDailyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsGetDailyPrintUsageByUser) | **Get** /reports/dailyPrintUsageByUser/{printUsageByUser-id} | Get dailyPrintUsageByUser from reports
[**ReportsGetMonthlyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsGetMonthlyPrintUsageByUser) | **Get** /reports/monthlyPrintUsageByUser/{printUsageByUser-id} | Get monthlyPrintUsageByUser from reports
[**ReportsListDailyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsListDailyPrintUsageByUser) | **Get** /reports/dailyPrintUsageByUser | Get dailyPrintUsageByUser from reports
[**ReportsListMonthlyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsListMonthlyPrintUsageByUser) | **Get** /reports/monthlyPrintUsageByUser | Get monthlyPrintUsageByUser from reports
[**ReportsUpdateDailyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsUpdateDailyPrintUsageByUser) | **Patch** /reports/dailyPrintUsageByUser/{printUsageByUser-id} | Update the navigation property dailyPrintUsageByUser in reports
[**ReportsUpdateMonthlyPrintUsageByUser**](ReportsPrintUsageByUserApi.md#ReportsUpdateMonthlyPrintUsageByUser) | **Patch** /reports/monthlyPrintUsageByUser/{printUsageByUser-id} | Update the navigation property monthlyPrintUsageByUser in reports



## ReportsCreateDailyPrintUsageByUser

> MicrosoftGraphPrintUsageByUser ReportsCreateDailyPrintUsageByUser(ctx).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()

Create new navigation property to dailyPrintUsageByUser for reports

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
    microsoftGraphPrintUsageByUser := *openapiclient.NewMicrosoftGraphPrintUsageByUser() // MicrosoftGraphPrintUsageByUser | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsCreateDailyPrintUsageByUser(context.Background()).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsCreateDailyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCreateDailyPrintUsageByUser`: MicrosoftGraphPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsCreateDailyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCreateDailyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintUsageByUser** | [**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsCreateMonthlyPrintUsageByUser

> MicrosoftGraphPrintUsageByUser ReportsCreateMonthlyPrintUsageByUser(ctx).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()

Create new navigation property to monthlyPrintUsageByUser for reports

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
    microsoftGraphPrintUsageByUser := *openapiclient.NewMicrosoftGraphPrintUsageByUser() // MicrosoftGraphPrintUsageByUser | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsCreateMonthlyPrintUsageByUser(context.Background()).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsCreateMonthlyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCreateMonthlyPrintUsageByUser`: MicrosoftGraphPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsCreateMonthlyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsCreateMonthlyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintUsageByUser** | [**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsDeleteDailyPrintUsageByUser

> ReportsDeleteDailyPrintUsageByUser(ctx, printUsageByUserId).IfMatch(ifMatch).Execute()

Delete navigation property dailyPrintUsageByUser for reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsDeleteDailyPrintUsageByUser(context.Background(), printUsageByUserId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsDeleteDailyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDeleteDailyPrintUsageByUserRequest struct via the builder pattern


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


## ReportsDeleteMonthlyPrintUsageByUser

> ReportsDeleteMonthlyPrintUsageByUser(ctx, printUsageByUserId).IfMatch(ifMatch).Execute()

Delete navigation property monthlyPrintUsageByUser for reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsDeleteMonthlyPrintUsageByUser(context.Background(), printUsageByUserId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsDeleteMonthlyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDeleteMonthlyPrintUsageByUserRequest struct via the builder pattern


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


## ReportsGetDailyPrintUsageByUser

> MicrosoftGraphPrintUsageByUser ReportsGetDailyPrintUsageByUser(ctx, printUsageByUserId).Select_(select_).Expand(expand).Execute()

Get dailyPrintUsageByUser from reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsGetDailyPrintUsageByUser(context.Background(), printUsageByUserId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsGetDailyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetDailyPrintUsageByUser`: MicrosoftGraphPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsGetDailyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetDailyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGetMonthlyPrintUsageByUser

> MicrosoftGraphPrintUsageByUser ReportsGetMonthlyPrintUsageByUser(ctx, printUsageByUserId).Select_(select_).Expand(expand).Execute()

Get monthlyPrintUsageByUser from reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsGetMonthlyPrintUsageByUser(context.Background(), printUsageByUserId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsGetMonthlyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetMonthlyPrintUsageByUser`: MicrosoftGraphPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsGetMonthlyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetMonthlyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsListDailyPrintUsageByUser

> CollectionOfPrintUsageByUser ReportsListDailyPrintUsageByUser(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get dailyPrintUsageByUser from reports

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
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsListDailyPrintUsageByUser(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsListDailyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsListDailyPrintUsageByUser`: CollectionOfPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsListDailyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsListDailyPrintUsageByUserRequest struct via the builder pattern


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

[**CollectionOfPrintUsageByUser**](CollectionOfPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsListMonthlyPrintUsageByUser

> CollectionOfPrintUsageByUser ReportsListMonthlyPrintUsageByUser(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get monthlyPrintUsageByUser from reports

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
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsListMonthlyPrintUsageByUser(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsListMonthlyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsListMonthlyPrintUsageByUser`: CollectionOfPrintUsageByUser
    fmt.Fprintf(os.Stdout, "Response from `ReportsPrintUsageByUserApi.ReportsListMonthlyPrintUsageByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsListMonthlyPrintUsageByUserRequest struct via the builder pattern


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

[**CollectionOfPrintUsageByUser**](CollectionOfPrintUsageByUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUpdateDailyPrintUsageByUser

> ReportsUpdateDailyPrintUsageByUser(ctx, printUsageByUserId).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()

Update the navigation property dailyPrintUsageByUser in reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    microsoftGraphPrintUsageByUser := *openapiclient.NewMicrosoftGraphPrintUsageByUser() // MicrosoftGraphPrintUsageByUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsUpdateDailyPrintUsageByUser(context.Background(), printUsageByUserId).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsUpdateDailyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateDailyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintUsageByUser** | [**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) | New navigation property values | 

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


## ReportsUpdateMonthlyPrintUsageByUser

> ReportsUpdateMonthlyPrintUsageByUser(ctx, printUsageByUserId).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()

Update the navigation property monthlyPrintUsageByUser in reports

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
    printUsageByUserId := "printUsageByUserId_example" // string | key: id of printUsageByUser
    microsoftGraphPrintUsageByUser := *openapiclient.NewMicrosoftGraphPrintUsageByUser() // MicrosoftGraphPrintUsageByUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsPrintUsageByUserApi.ReportsUpdateMonthlyPrintUsageByUser(context.Background(), printUsageByUserId).MicrosoftGraphPrintUsageByUser(microsoftGraphPrintUsageByUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsPrintUsageByUserApi.ReportsUpdateMonthlyPrintUsageByUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printUsageByUserId** | **string** | key: id of printUsageByUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateMonthlyPrintUsageByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintUsageByUser** | [**MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) | New navigation property values | 

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

