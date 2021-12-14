# \PrintPrintTaskDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreateTaskDefinitions**](PrintPrintTaskDefinitionApi.md#PrintCreateTaskDefinitions) | **Post** /print/taskDefinitions | Create new navigation property to taskDefinitions for print
[**PrintDeleteTaskDefinitions**](PrintPrintTaskDefinitionApi.md#PrintDeleteTaskDefinitions) | **Delete** /print/taskDefinitions/{printTaskDefinition-id} | Delete navigation property taskDefinitions for print
[**PrintGetTaskDefinitions**](PrintPrintTaskDefinitionApi.md#PrintGetTaskDefinitions) | **Get** /print/taskDefinitions/{printTaskDefinition-id} | Get taskDefinitions from print
[**PrintListTaskDefinitions**](PrintPrintTaskDefinitionApi.md#PrintListTaskDefinitions) | **Get** /print/taskDefinitions | Get taskDefinitions from print
[**PrintTaskDefinitionsCreateTasks**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsCreateTasks) | **Post** /print/taskDefinitions/{printTaskDefinition-id}/tasks | Create new navigation property to tasks for print
[**PrintTaskDefinitionsDeleteTasks**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsDeleteTasks) | **Delete** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id} | Delete navigation property tasks for print
[**PrintTaskDefinitionsGetTasks**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsGetTasks) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id} | Get tasks from print
[**PrintTaskDefinitionsListTasks**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsListTasks) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks | Get tasks from print
[**PrintTaskDefinitionsTasksDeleteRefDefinition**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksDeleteRefDefinition) | **Delete** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/definition/$ref | Delete ref of navigation property definition for print
[**PrintTaskDefinitionsTasksDeleteRefTrigger**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksDeleteRefTrigger) | **Delete** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/trigger/$ref | Delete ref of navigation property trigger for print
[**PrintTaskDefinitionsTasksGetDefinition**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksGetDefinition) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/definition | Get definition from print
[**PrintTaskDefinitionsTasksGetRefDefinition**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksGetRefDefinition) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/definition/$ref | Get ref of definition from print
[**PrintTaskDefinitionsTasksGetRefTrigger**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksGetRefTrigger) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/trigger/$ref | Get ref of trigger from print
[**PrintTaskDefinitionsTasksGetTrigger**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksGetTrigger) | **Get** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/trigger | Get trigger from print
[**PrintTaskDefinitionsTasksUpdateRefDefinition**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksUpdateRefDefinition) | **Put** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/definition/$ref | Update the ref of navigation property definition in print
[**PrintTaskDefinitionsTasksUpdateRefTrigger**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsTasksUpdateRefTrigger) | **Put** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id}/trigger/$ref | Update the ref of navigation property trigger in print
[**PrintTaskDefinitionsUpdateTasks**](PrintPrintTaskDefinitionApi.md#PrintTaskDefinitionsUpdateTasks) | **Patch** /print/taskDefinitions/{printTaskDefinition-id}/tasks/{printTask-id} | Update the navigation property tasks in print
[**PrintUpdateTaskDefinitions**](PrintPrintTaskDefinitionApi.md#PrintUpdateTaskDefinitions) | **Patch** /print/taskDefinitions/{printTaskDefinition-id} | Update the navigation property taskDefinitions in print



## PrintCreateTaskDefinitions

> MicrosoftGraphPrintTaskDefinition PrintCreateTaskDefinitions(ctx).MicrosoftGraphPrintTaskDefinition(microsoftGraphPrintTaskDefinition).Execute()

Create new navigation property to taskDefinitions for print



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
    microsoftGraphPrintTaskDefinition := *openapiclient.NewMicrosoftGraphPrintTaskDefinition() // MicrosoftGraphPrintTaskDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintCreateTaskDefinitions(context.Background()).MicrosoftGraphPrintTaskDefinition(microsoftGraphPrintTaskDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintCreateTaskDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreateTaskDefinitions`: MicrosoftGraphPrintTaskDefinition
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintCreateTaskDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreateTaskDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintTaskDefinition** | [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeleteTaskDefinitions

> PrintDeleteTaskDefinitions(ctx, printTaskDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property taskDefinitions for print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintDeleteTaskDefinitions(context.Background(), printTaskDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintDeleteTaskDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeleteTaskDefinitionsRequest struct via the builder pattern


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


## PrintGetTaskDefinitions

> MicrosoftGraphPrintTaskDefinition PrintGetTaskDefinitions(ctx, printTaskDefinitionId).Select_(select_).Expand(expand).Execute()

Get taskDefinitions from print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintGetTaskDefinitions(context.Background(), printTaskDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintGetTaskDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetTaskDefinitions`: MicrosoftGraphPrintTaskDefinition
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintGetTaskDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetTaskDefinitionsRequest struct via the builder pattern


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


## PrintListTaskDefinitions

> CollectionOfPrintTaskDefinition PrintListTaskDefinitions(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get taskDefinitions from print



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
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintListTaskDefinitions(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintListTaskDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListTaskDefinitions`: CollectionOfPrintTaskDefinition
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintListTaskDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListTaskDefinitionsRequest struct via the builder pattern


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

[**CollectionOfPrintTaskDefinition**](CollectionOfPrintTaskDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintTaskDefinitionsCreateTasks

> MicrosoftGraphPrintTask PrintTaskDefinitionsCreateTasks(ctx, printTaskDefinitionId).MicrosoftGraphPrintTask(microsoftGraphPrintTask).Execute()

Create new navigation property to tasks for print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    microsoftGraphPrintTask := *openapiclient.NewMicrosoftGraphPrintTask() // MicrosoftGraphPrintTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsCreateTasks(context.Background(), printTaskDefinitionId).MicrosoftGraphPrintTask(microsoftGraphPrintTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsCreateTasks`: MicrosoftGraphPrintTask
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsCreateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintTask** | [**MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintTaskDefinitionsDeleteTasks

> PrintTaskDefinitionsDeleteTasks(ctx, printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsDeleteTasks(context.Background(), printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsDeleteTasksRequest struct via the builder pattern


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


## PrintTaskDefinitionsGetTasks

> MicrosoftGraphPrintTask PrintTaskDefinitionsGetTasks(ctx, printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsGetTasks(context.Background(), printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsGetTasks`: MicrosoftGraphPrintTask
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintTaskDefinitionsListTasks

> CollectionOfPrintTask PrintTaskDefinitionsListTasks(ctx, printTaskDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
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
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsListTasks(context.Background(), printTaskDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsListTasks`: CollectionOfPrintTask
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsListTasksRequest struct via the builder pattern


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

[**CollectionOfPrintTask**](CollectionOfPrintTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintTaskDefinitionsTasksDeleteRefDefinition

> PrintTaskDefinitionsTasksDeleteRefDefinition(ctx, printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()

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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksDeleteRefDefinition(context.Background(), printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksDeleteRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksDeleteRefDefinitionRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksDeleteRefTrigger

> PrintTaskDefinitionsTasksDeleteRefTrigger(ctx, printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()

Delete ref of navigation property trigger for print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksDeleteRefTrigger(context.Background(), printTaskDefinitionId, printTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksDeleteRefTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksDeleteRefTriggerRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksGetDefinition

> MicrosoftGraphPrintTaskDefinition PrintTaskDefinitionsTasksGetDefinition(ctx, printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()

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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetDefinition(context.Background(), printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsTasksGetDefinition`: MicrosoftGraphPrintTaskDefinition
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksGetDefinitionRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksGetRefDefinition

> string PrintTaskDefinitionsTasksGetRefDefinition(ctx, printTaskDefinitionId, printTaskId).Execute()

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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefDefinition(context.Background(), printTaskDefinitionId, printTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsTasksGetRefDefinition`: string
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksGetRefDefinitionRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksGetRefTrigger

> string PrintTaskDefinitionsTasksGetRefTrigger(ctx, printTaskDefinitionId, printTaskId).Execute()

Get ref of trigger from print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefTrigger(context.Background(), printTaskDefinitionId, printTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsTasksGetRefTrigger`: string
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetRefTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksGetRefTriggerRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksGetTrigger

> MicrosoftGraphPrintTaskTrigger PrintTaskDefinitionsTasksGetTrigger(ctx, printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()

Get trigger from print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetTrigger(context.Background(), printTaskDefinitionId, printTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintTaskDefinitionsTasksGetTrigger`: MicrosoftGraphPrintTaskTrigger
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksGetTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksGetTriggerRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksUpdateRefDefinition

> PrintTaskDefinitionsTasksUpdateRefDefinition(ctx, printTaskDefinitionId, printTaskId).RequestBody(requestBody).Execute()

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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksUpdateRefDefinition(context.Background(), printTaskDefinitionId, printTaskId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksUpdateRefDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksUpdateRefDefinitionRequest struct via the builder pattern


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


## PrintTaskDefinitionsTasksUpdateRefTrigger

> PrintTaskDefinitionsTasksUpdateRefTrigger(ctx, printTaskDefinitionId, printTaskId).RequestBody(requestBody).Execute()

Update the ref of navigation property trigger in print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksUpdateRefTrigger(context.Background(), printTaskDefinitionId, printTaskId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsTasksUpdateRefTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsTasksUpdateRefTriggerRequest struct via the builder pattern


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


## PrintTaskDefinitionsUpdateTasks

> PrintTaskDefinitionsUpdateTasks(ctx, printTaskDefinitionId, printTaskId).MicrosoftGraphPrintTask(microsoftGraphPrintTask).Execute()

Update the navigation property tasks in print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    printTaskId := "printTaskId_example" // string | key: id of printTask
    microsoftGraphPrintTask := *openapiclient.NewMicrosoftGraphPrintTask() // MicrosoftGraphPrintTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintTaskDefinitionsUpdateTasks(context.Background(), printTaskDefinitionId, printTaskId).MicrosoftGraphPrintTask(microsoftGraphPrintTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintTaskDefinitionsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 
**printTaskId** | **string** | key: id of printTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintTaskDefinitionsUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPrintTask** | [**MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md) | New navigation property values | 

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


## PrintUpdateTaskDefinitions

> PrintUpdateTaskDefinitions(ctx, printTaskDefinitionId).MicrosoftGraphPrintTaskDefinition(microsoftGraphPrintTaskDefinition).Execute()

Update the navigation property taskDefinitions in print



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
    printTaskDefinitionId := "printTaskDefinitionId_example" // string | key: id of printTaskDefinition
    microsoftGraphPrintTaskDefinition := *openapiclient.NewMicrosoftGraphPrintTaskDefinition() // MicrosoftGraphPrintTaskDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintTaskDefinitionApi.PrintUpdateTaskDefinitions(context.Background(), printTaskDefinitionId).MicrosoftGraphPrintTaskDefinition(microsoftGraphPrintTaskDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintTaskDefinitionApi.PrintUpdateTaskDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printTaskDefinitionId** | **string** | key: id of printTaskDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdateTaskDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintTaskDefinition** | [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) | New navigation property values | 

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

