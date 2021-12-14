# \MeTodoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteTodo**](MeTodoApi.md#MeDeleteTodo) | **Delete** /me/todo | Delete navigation property todo for me
[**MeGetTodo**](MeTodoApi.md#MeGetTodo) | **Get** /me/todo | Get todo from me
[**MeTodoCreateLists**](MeTodoApi.md#MeTodoCreateLists) | **Post** /me/todo/lists | Create new navigation property to lists for me
[**MeTodoDeleteLists**](MeTodoApi.md#MeTodoDeleteLists) | **Delete** /me/todo/lists/{todoTaskList-id} | Delete navigation property lists for me
[**MeTodoGetLists**](MeTodoApi.md#MeTodoGetLists) | **Get** /me/todo/lists/{todoTaskList-id} | Get lists from me
[**MeTodoListLists**](MeTodoApi.md#MeTodoListLists) | **Get** /me/todo/lists | Get lists from me
[**MeTodoListsCreateExtensions**](MeTodoApi.md#MeTodoListsCreateExtensions) | **Post** /me/todo/lists/{todoTaskList-id}/extensions | Create new navigation property to extensions for me
[**MeTodoListsCreateTasks**](MeTodoApi.md#MeTodoListsCreateTasks) | **Post** /me/todo/lists/{todoTaskList-id}/tasks | Create new navigation property to tasks for me
[**MeTodoListsDeleteExtensions**](MeTodoApi.md#MeTodoListsDeleteExtensions) | **Delete** /me/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Delete navigation property extensions for me
[**MeTodoListsDeleteTasks**](MeTodoApi.md#MeTodoListsDeleteTasks) | **Delete** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Delete navigation property tasks for me
[**MeTodoListsGetExtensions**](MeTodoApi.md#MeTodoListsGetExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Get extensions from me
[**MeTodoListsGetTasks**](MeTodoApi.md#MeTodoListsGetTasks) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Get tasks from me
[**MeTodoListsListExtensions**](MeTodoApi.md#MeTodoListsListExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/extensions | Get extensions from me
[**MeTodoListsListTasks**](MeTodoApi.md#MeTodoListsListTasks) | **Get** /me/todo/lists/{todoTaskList-id}/tasks | Get tasks from me
[**MeTodoListsTasksCreateExtensions**](MeTodoApi.md#MeTodoListsTasksCreateExtensions) | **Post** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Create new navigation property to extensions for me
[**MeTodoListsTasksCreateLinkedResources**](MeTodoApi.md#MeTodoListsTasksCreateLinkedResources) | **Post** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Create new navigation property to linkedResources for me
[**MeTodoListsTasksDeleteExtensions**](MeTodoApi.md#MeTodoListsTasksDeleteExtensions) | **Delete** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Delete navigation property extensions for me
[**MeTodoListsTasksDeleteLinkedResources**](MeTodoApi.md#MeTodoListsTasksDeleteLinkedResources) | **Delete** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Delete navigation property linkedResources for me
[**MeTodoListsTasksGetExtensions**](MeTodoApi.md#MeTodoListsTasksGetExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Get extensions from me
[**MeTodoListsTasksGetLinkedResources**](MeTodoApi.md#MeTodoListsTasksGetLinkedResources) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Get linkedResources from me
[**MeTodoListsTasksListExtensions**](MeTodoApi.md#MeTodoListsTasksListExtensions) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Get extensions from me
[**MeTodoListsTasksListLinkedResources**](MeTodoApi.md#MeTodoListsTasksListLinkedResources) | **Get** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Get linkedResources from me
[**MeTodoListsTasksUpdateExtensions**](MeTodoApi.md#MeTodoListsTasksUpdateExtensions) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeTodoListsTasksUpdateLinkedResources**](MeTodoApi.md#MeTodoListsTasksUpdateLinkedResources) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Update the navigation property linkedResources in me
[**MeTodoListsUpdateExtensions**](MeTodoApi.md#MeTodoListsUpdateExtensions) | **Patch** /me/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeTodoListsUpdateTasks**](MeTodoApi.md#MeTodoListsUpdateTasks) | **Patch** /me/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Update the navigation property tasks in me
[**MeTodoUpdateLists**](MeTodoApi.md#MeTodoUpdateLists) | **Patch** /me/todo/lists/{todoTaskList-id} | Update the navigation property lists in me
[**MeUpdateTodo**](MeTodoApi.md#MeUpdateTodo) | **Patch** /me/todo | Update the navigation property todo in me



## MeDeleteTodo

> MeDeleteTodo(ctx).IfMatch(ifMatch).Execute()

Delete navigation property todo for me



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeDeleteTodo(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeDeleteTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteTodoRequest struct via the builder pattern


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


## MeGetTodo

> MicrosoftGraphTodo MeGetTodo(ctx).Select_(select_).Expand(expand).Execute()

Get todo from me



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeGetTodo(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeGetTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetTodo`: MicrosoftGraphTodo
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeGetTodo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetTodoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTodo**](MicrosoftGraphTodo.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoCreateLists

> MicrosoftGraphTodoTaskList MeTodoCreateLists(ctx).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()

Create new navigation property to lists for me



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
    microsoftGraphTodoTaskList := *openapiclient.NewMicrosoftGraphTodoTaskList() // MicrosoftGraphTodoTaskList | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoCreateLists(context.Background()).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoCreateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoCreateLists`: MicrosoftGraphTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoCreateLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoCreateListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTodoTaskList** | [**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md) | New navigation property | 

### Return type

[**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoDeleteLists

> MeTodoDeleteLists(ctx, todoTaskListId).IfMatch(ifMatch).Execute()

Delete navigation property lists for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoDeleteLists(context.Background(), todoTaskListId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoDeleteLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoDeleteListsRequest struct via the builder pattern


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


## MeTodoGetLists

> MicrosoftGraphTodoTaskList MeTodoGetLists(ctx, todoTaskListId).Select_(select_).Expand(expand).Execute()

Get lists from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoGetLists(context.Background(), todoTaskListId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoGetLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoGetLists`: MicrosoftGraphTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoGetLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListLists

> CollectionOfTodoTaskList MeTodoListLists(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get lists from me



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
    resp, r, err := api_client.MeTodoApi.MeTodoListLists(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListLists`: CollectionOfTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListListsRequest struct via the builder pattern


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

[**CollectionOfTodoTaskList**](CollectionOfTodoTaskList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsCreateExtensions

> MicrosoftGraphExtension MeTodoListsCreateExtensions(ctx, todoTaskListId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsCreateExtensions(context.Background(), todoTaskListId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsCreateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsCreateTasks

> MicrosoftGraphTodoTask MeTodoListsCreateTasks(ctx, todoTaskListId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()

Create new navigation property to tasks for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphTodoTask := *openapiclient.NewMicrosoftGraphTodoTask() // MicrosoftGraphTodoTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsCreateTasks(context.Background(), todoTaskListId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsCreateTasks`: MicrosoftGraphTodoTask
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsCreateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTodoTask** | [**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md) | New navigation property | 

### Return type

[**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsDeleteExtensions

> MeTodoListsDeleteExtensions(ctx, todoTaskListId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsDeleteExtensions(context.Background(), todoTaskListId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsDeleteExtensionsRequest struct via the builder pattern


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


## MeTodoListsDeleteTasks

> MeTodoListsDeleteTasks(ctx, todoTaskListId, todoTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsDeleteTasks(context.Background(), todoTaskListId, todoTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsDeleteTasksRequest struct via the builder pattern


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


## MeTodoListsGetExtensions

> MicrosoftGraphExtension MeTodoListsGetExtensions(ctx, todoTaskListId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsGetExtensions(context.Background(), todoTaskListId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsGetTasks

> MicrosoftGraphTodoTask MeTodoListsGetTasks(ctx, todoTaskListId, todoTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsGetTasks(context.Background(), todoTaskListId, todoTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsGetTasks`: MicrosoftGraphTodoTask
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsListExtensions

> CollectionOfExtension MeTodoListsListExtensions(ctx, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
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
    resp, r, err := api_client.MeTodoApi.MeTodoListsListExtensions(context.Background(), todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsListExtensionsRequest struct via the builder pattern


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

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsListTasks

> CollectionOfTodoTask MeTodoListsListTasks(ctx, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
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
    resp, r, err := api_client.MeTodoApi.MeTodoListsListTasks(context.Background(), todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsListTasks`: CollectionOfTodoTask
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsListTasksRequest struct via the builder pattern


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

[**CollectionOfTodoTask**](CollectionOfTodoTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksCreateExtensions

> MicrosoftGraphExtension MeTodoListsTasksCreateExtensions(ctx, todoTaskListId, todoTaskId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksCreateExtensions(context.Background(), todoTaskListId, todoTaskId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksCreateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksCreateLinkedResources

> MicrosoftGraphLinkedResource MeTodoListsTasksCreateLinkedResources(ctx, todoTaskListId, todoTaskId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()

Create new navigation property to linkedResources for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphLinkedResource := *openapiclient.NewMicrosoftGraphLinkedResource() // MicrosoftGraphLinkedResource | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksCreateLinkedResources(context.Background(), todoTaskListId, todoTaskId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksCreateLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksCreateLinkedResources`: MicrosoftGraphLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksCreateLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksCreateLinkedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphLinkedResource** | [**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md) | New navigation property | 

### Return type

[**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksDeleteExtensions

> MeTodoListsTasksDeleteExtensions(ctx, todoTaskListId, todoTaskId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksDeleteExtensions(context.Background(), todoTaskListId, todoTaskId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksDeleteExtensionsRequest struct via the builder pattern


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


## MeTodoListsTasksDeleteLinkedResources

> MeTodoListsTasksDeleteLinkedResources(ctx, todoTaskListId, todoTaskId, linkedResourceId).IfMatch(ifMatch).Execute()

Delete navigation property linkedResources for me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksDeleteLinkedResources(context.Background(), todoTaskListId, todoTaskId, linkedResourceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksDeleteLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksDeleteLinkedResourcesRequest struct via the builder pattern


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


## MeTodoListsTasksGetExtensions

> MicrosoftGraphExtension MeTodoListsTasksGetExtensions(ctx, todoTaskListId, todoTaskId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksGetExtensions(context.Background(), todoTaskListId, todoTaskId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksGetLinkedResources

> MicrosoftGraphLinkedResource MeTodoListsTasksGetLinkedResources(ctx, todoTaskListId, todoTaskId, linkedResourceId).Select_(select_).Expand(expand).Execute()

Get linkedResources from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksGetLinkedResources(context.Background(), todoTaskListId, todoTaskId, linkedResourceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksGetLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksGetLinkedResources`: MicrosoftGraphLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksGetLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksGetLinkedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksListExtensions

> CollectionOfExtension MeTodoListsTasksListExtensions(ctx, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
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
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksListExtensions(context.Background(), todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksListExtensionsRequest struct via the builder pattern


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

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksListLinkedResources

> CollectionOfLinkedResource MeTodoListsTasksListLinkedResources(ctx, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get linkedResources from me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
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
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksListLinkedResources(context.Background(), todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksListLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeTodoListsTasksListLinkedResources`: CollectionOfLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `MeTodoApi.MeTodoListsTasksListLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksListLinkedResourcesRequest struct via the builder pattern


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

[**CollectionOfLinkedResource**](CollectionOfLinkedResource.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeTodoListsTasksUpdateExtensions

> MeTodoListsTasksUpdateExtensions(ctx, todoTaskListId, todoTaskId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksUpdateExtensions(context.Background(), todoTaskListId, todoTaskId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksUpdateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property values | 

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


## MeTodoListsTasksUpdateLinkedResources

> MeTodoListsTasksUpdateLinkedResources(ctx, todoTaskListId, todoTaskId, linkedResourceId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()

Update the navigation property linkedResources in me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    microsoftGraphLinkedResource := *openapiclient.NewMicrosoftGraphLinkedResource() // MicrosoftGraphLinkedResource | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsTasksUpdateLinkedResources(context.Background(), todoTaskListId, todoTaskId, linkedResourceId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsTasksUpdateLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsTasksUpdateLinkedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphLinkedResource** | [**MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md) | New navigation property values | 

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


## MeTodoListsUpdateExtensions

> MeTodoListsUpdateExtensions(ctx, todoTaskListId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsUpdateExtensions(context.Background(), todoTaskListId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsUpdateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property values | 

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


## MeTodoListsUpdateTasks

> MeTodoListsUpdateTasks(ctx, todoTaskListId, todoTaskId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()

Update the navigation property tasks in me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphTodoTask := *openapiclient.NewMicrosoftGraphTodoTask() // MicrosoftGraphTodoTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoListsUpdateTasks(context.Background(), todoTaskListId, todoTaskId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoListsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoListsUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTodoTask** | [**MicrosoftGraphTodoTask**](MicrosoftGraphTodoTask.md) | New navigation property values | 

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


## MeTodoUpdateLists

> MeTodoUpdateLists(ctx, todoTaskListId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()

Update the navigation property lists in me



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
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphTodoTaskList := *openapiclient.NewMicrosoftGraphTodoTaskList() // MicrosoftGraphTodoTaskList | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeTodoUpdateLists(context.Background(), todoTaskListId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeTodoUpdateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeTodoUpdateListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTodoTaskList** | [**MicrosoftGraphTodoTaskList**](MicrosoftGraphTodoTaskList.md) | New navigation property values | 

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


## MeUpdateTodo

> MeUpdateTodo(ctx).MicrosoftGraphTodo(microsoftGraphTodo).Execute()

Update the navigation property todo in me



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
    microsoftGraphTodo := *openapiclient.NewMicrosoftGraphTodo() // MicrosoftGraphTodo | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTodoApi.MeUpdateTodo(context.Background()).MicrosoftGraphTodo(microsoftGraphTodo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTodoApi.MeUpdateTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateTodoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTodo** | [**MicrosoftGraphTodo**](MicrosoftGraphTodo.md) | New navigation property values | 

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

