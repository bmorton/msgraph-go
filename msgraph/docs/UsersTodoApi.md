# \UsersTodoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteTodo**](UsersTodoApi.md#UsersDeleteTodo) | **Delete** /users/{user-id}/todo | Delete navigation property todo for users
[**UsersGetTodo**](UsersTodoApi.md#UsersGetTodo) | **Get** /users/{user-id}/todo | Get todo from users
[**UsersTodoCreateLists**](UsersTodoApi.md#UsersTodoCreateLists) | **Post** /users/{user-id}/todo/lists | Create new navigation property to lists for users
[**UsersTodoDeleteLists**](UsersTodoApi.md#UsersTodoDeleteLists) | **Delete** /users/{user-id}/todo/lists/{todoTaskList-id} | Delete navigation property lists for users
[**UsersTodoGetLists**](UsersTodoApi.md#UsersTodoGetLists) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id} | Get lists from users
[**UsersTodoListLists**](UsersTodoApi.md#UsersTodoListLists) | **Get** /users/{user-id}/todo/lists | Get lists from users
[**UsersTodoListsCreateExtensions**](UsersTodoApi.md#UsersTodoListsCreateExtensions) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions | Create new navigation property to extensions for users
[**UsersTodoListsCreateTasks**](UsersTodoApi.md#UsersTodoListsCreateTasks) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks | Create new navigation property to tasks for users
[**UsersTodoListsDeleteExtensions**](UsersTodoApi.md#UsersTodoListsDeleteExtensions) | **Delete** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersTodoListsDeleteTasks**](UsersTodoApi.md#UsersTodoListsDeleteTasks) | **Delete** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Delete navigation property tasks for users
[**UsersTodoListsGetExtensions**](UsersTodoApi.md#UsersTodoListsGetExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Get extensions from users
[**UsersTodoListsGetTasks**](UsersTodoApi.md#UsersTodoListsGetTasks) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Get tasks from users
[**UsersTodoListsListExtensions**](UsersTodoApi.md#UsersTodoListsListExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions | Get extensions from users
[**UsersTodoListsListTasks**](UsersTodoApi.md#UsersTodoListsListTasks) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks | Get tasks from users
[**UsersTodoListsTasksCreateExtensions**](UsersTodoApi.md#UsersTodoListsTasksCreateExtensions) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Create new navigation property to extensions for users
[**UsersTodoListsTasksCreateLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksCreateLinkedResources) | **Post** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Create new navigation property to linkedResources for users
[**UsersTodoListsTasksDeleteExtensions**](UsersTodoApi.md#UsersTodoListsTasksDeleteExtensions) | **Delete** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersTodoListsTasksDeleteLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksDeleteLinkedResources) | **Delete** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Delete navigation property linkedResources for users
[**UsersTodoListsTasksGetExtensions**](UsersTodoApi.md#UsersTodoListsTasksGetExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Get extensions from users
[**UsersTodoListsTasksGetLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksGetLinkedResources) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Get linkedResources from users
[**UsersTodoListsTasksListExtensions**](UsersTodoApi.md#UsersTodoListsTasksListExtensions) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions | Get extensions from users
[**UsersTodoListsTasksListLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksListLinkedResources) | **Get** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources | Get linkedResources from users
[**UsersTodoListsTasksUpdateExtensions**](UsersTodoApi.md#UsersTodoListsTasksUpdateExtensions) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersTodoListsTasksUpdateLinkedResources**](UsersTodoApi.md#UsersTodoListsTasksUpdateLinkedResources) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id}/linkedResources/{linkedResource-id} | Update the navigation property linkedResources in users
[**UsersTodoListsUpdateExtensions**](UsersTodoApi.md#UsersTodoListsUpdateExtensions) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersTodoListsUpdateTasks**](UsersTodoApi.md#UsersTodoListsUpdateTasks) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id}/tasks/{todoTask-id} | Update the navigation property tasks in users
[**UsersTodoUpdateLists**](UsersTodoApi.md#UsersTodoUpdateLists) | **Patch** /users/{user-id}/todo/lists/{todoTaskList-id} | Update the navigation property lists in users
[**UsersUpdateTodo**](UsersTodoApi.md#UsersUpdateTodo) | **Patch** /users/{user-id}/todo | Update the navigation property todo in users



## UsersDeleteTodo

> UsersDeleteTodo(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property todo for users



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
    userId := "userId_example" // string | key: id of user
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersDeleteTodo(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersDeleteTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteTodoRequest struct via the builder pattern


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


## UsersGetTodo

> MicrosoftGraphTodo UsersGetTodo(ctx, userId).Select_(select_).Expand(expand).Execute()

Get todo from users



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
    userId := "userId_example" // string | key: id of user
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersGetTodo(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersGetTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetTodo`: MicrosoftGraphTodo
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersGetTodo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetTodoRequest struct via the builder pattern


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


## UsersTodoCreateLists

> MicrosoftGraphTodoTaskList UsersTodoCreateLists(ctx, userId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()

Create new navigation property to lists for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphTodoTaskList := *openapiclient.NewMicrosoftGraphTodoTaskList() // MicrosoftGraphTodoTaskList | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoCreateLists(context.Background(), userId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoCreateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoCreateLists`: MicrosoftGraphTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoCreateLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoCreateListsRequest struct via the builder pattern


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


## UsersTodoDeleteLists

> UsersTodoDeleteLists(ctx, userId, todoTaskListId).IfMatch(ifMatch).Execute()

Delete navigation property lists for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoDeleteLists(context.Background(), userId, todoTaskListId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoDeleteLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoDeleteListsRequest struct via the builder pattern


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


## UsersTodoGetLists

> MicrosoftGraphTodoTaskList UsersTodoGetLists(ctx, userId, todoTaskListId).Select_(select_).Expand(expand).Execute()

Get lists from users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoGetLists(context.Background(), userId, todoTaskListId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoGetLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoGetLists`: MicrosoftGraphTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoGetLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoGetListsRequest struct via the builder pattern


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


## UsersTodoListLists

> CollectionOfTodoTaskList UsersTodoListLists(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get lists from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersTodoApi.UsersTodoListLists(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListLists`: CollectionOfTodoTaskList
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListListsRequest struct via the builder pattern


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


## UsersTodoListsCreateExtensions

> MicrosoftGraphExtension UsersTodoListsCreateExtensions(ctx, userId, todoTaskListId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsCreateExtensions(context.Background(), userId, todoTaskListId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsCreateExtensionsRequest struct via the builder pattern


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


## UsersTodoListsCreateTasks

> MicrosoftGraphTodoTask UsersTodoListsCreateTasks(ctx, userId, todoTaskListId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()

Create new navigation property to tasks for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphTodoTask := *openapiclient.NewMicrosoftGraphTodoTask() // MicrosoftGraphTodoTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsCreateTasks(context.Background(), userId, todoTaskListId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsCreateTasks`: MicrosoftGraphTodoTask
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsCreateTasksRequest struct via the builder pattern


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


## UsersTodoListsDeleteExtensions

> UsersTodoListsDeleteExtensions(ctx, userId, todoTaskListId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsDeleteExtensions(context.Background(), userId, todoTaskListId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsDeleteExtensionsRequest struct via the builder pattern


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


## UsersTodoListsDeleteTasks

> UsersTodoListsDeleteTasks(ctx, userId, todoTaskListId, todoTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsDeleteTasks(context.Background(), userId, todoTaskListId, todoTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsDeleteTasksRequest struct via the builder pattern


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


## UsersTodoListsGetExtensions

> MicrosoftGraphExtension UsersTodoListsGetExtensions(ctx, userId, todoTaskListId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsGetExtensions(context.Background(), userId, todoTaskListId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsGetExtensionsRequest struct via the builder pattern


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


## UsersTodoListsGetTasks

> MicrosoftGraphTodoTask UsersTodoListsGetTasks(ctx, userId, todoTaskListId, todoTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsGetTasks(context.Background(), userId, todoTaskListId, todoTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsGetTasks`: MicrosoftGraphTodoTask
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsGetTasksRequest struct via the builder pattern


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


## UsersTodoListsListExtensions

> CollectionOfExtension UsersTodoListsListExtensions(ctx, userId, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsListExtensions(context.Background(), userId, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsListExtensionsRequest struct via the builder pattern


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


## UsersTodoListsListTasks

> CollectionOfTodoTask UsersTodoListsListTasks(ctx, userId, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsListTasks(context.Background(), userId, todoTaskListId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsListTasks`: CollectionOfTodoTask
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsListTasksRequest struct via the builder pattern


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


## UsersTodoListsTasksCreateExtensions

> MicrosoftGraphExtension UsersTodoListsTasksCreateExtensions(ctx, userId, todoTaskListId, todoTaskId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksCreateExtensions(context.Background(), userId, todoTaskListId, todoTaskId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksCreateExtensionsRequest struct via the builder pattern


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


## UsersTodoListsTasksCreateLinkedResources

> MicrosoftGraphLinkedResource UsersTodoListsTasksCreateLinkedResources(ctx, userId, todoTaskListId, todoTaskId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()

Create new navigation property to linkedResources for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphLinkedResource := *openapiclient.NewMicrosoftGraphLinkedResource() // MicrosoftGraphLinkedResource | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksCreateLinkedResources(context.Background(), userId, todoTaskListId, todoTaskId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksCreateLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksCreateLinkedResources`: MicrosoftGraphLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksCreateLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksCreateLinkedResourcesRequest struct via the builder pattern


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


## UsersTodoListsTasksDeleteExtensions

> UsersTodoListsTasksDeleteExtensions(ctx, userId, todoTaskListId, todoTaskId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksDeleteExtensions(context.Background(), userId, todoTaskListId, todoTaskId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksDeleteExtensionsRequest struct via the builder pattern


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


## UsersTodoListsTasksDeleteLinkedResources

> UsersTodoListsTasksDeleteLinkedResources(ctx, userId, todoTaskListId, todoTaskId, linkedResourceId).IfMatch(ifMatch).Execute()

Delete navigation property linkedResources for users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksDeleteLinkedResources(context.Background(), userId, todoTaskListId, todoTaskId, linkedResourceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksDeleteLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksDeleteLinkedResourcesRequest struct via the builder pattern


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


## UsersTodoListsTasksGetExtensions

> MicrosoftGraphExtension UsersTodoListsTasksGetExtensions(ctx, userId, todoTaskListId, todoTaskId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksGetExtensions(context.Background(), userId, todoTaskListId, todoTaskId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksGetExtensionsRequest struct via the builder pattern


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


## UsersTodoListsTasksGetLinkedResources

> MicrosoftGraphLinkedResource UsersTodoListsTasksGetLinkedResources(ctx, userId, todoTaskListId, todoTaskId, linkedResourceId).Select_(select_).Expand(expand).Execute()

Get linkedResources from users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksGetLinkedResources(context.Background(), userId, todoTaskListId, todoTaskId, linkedResourceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksGetLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksGetLinkedResources`: MicrosoftGraphLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksGetLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksGetLinkedResourcesRequest struct via the builder pattern


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


## UsersTodoListsTasksListExtensions

> CollectionOfExtension UsersTodoListsTasksListExtensions(ctx, userId, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksListExtensions(context.Background(), userId, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksListExtensionsRequest struct via the builder pattern


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


## UsersTodoListsTasksListLinkedResources

> CollectionOfLinkedResource UsersTodoListsTasksListLinkedResources(ctx, userId, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get linkedResources from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksListLinkedResources(context.Background(), userId, todoTaskListId, todoTaskId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksListLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTodoListsTasksListLinkedResources`: CollectionOfLinkedResource
    fmt.Fprintf(os.Stdout, "Response from `UsersTodoApi.UsersTodoListsTasksListLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksListLinkedResourcesRequest struct via the builder pattern


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


## UsersTodoListsTasksUpdateExtensions

> UsersTodoListsTasksUpdateExtensions(ctx, userId, todoTaskListId, todoTaskId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksUpdateExtensions(context.Background(), userId, todoTaskListId, todoTaskId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksUpdateExtensionsRequest struct via the builder pattern


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


## UsersTodoListsTasksUpdateLinkedResources

> UsersTodoListsTasksUpdateLinkedResources(ctx, userId, todoTaskListId, todoTaskId, linkedResourceId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()

Update the navigation property linkedResources in users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    linkedResourceId := "linkedResourceId_example" // string | key: id of linkedResource
    microsoftGraphLinkedResource := *openapiclient.NewMicrosoftGraphLinkedResource() // MicrosoftGraphLinkedResource | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsTasksUpdateLinkedResources(context.Background(), userId, todoTaskListId, todoTaskId, linkedResourceId).MicrosoftGraphLinkedResource(microsoftGraphLinkedResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsTasksUpdateLinkedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 
**linkedResourceId** | **string** | key: id of linkedResource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsTasksUpdateLinkedResourcesRequest struct via the builder pattern


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


## UsersTodoListsUpdateExtensions

> UsersTodoListsUpdateExtensions(ctx, userId, todoTaskListId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsUpdateExtensions(context.Background(), userId, todoTaskListId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsUpdateExtensionsRequest struct via the builder pattern


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


## UsersTodoListsUpdateTasks

> UsersTodoListsUpdateTasks(ctx, userId, todoTaskListId, todoTaskId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()

Update the navigation property tasks in users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    todoTaskId := "todoTaskId_example" // string | key: id of todoTask
    microsoftGraphTodoTask := *openapiclient.NewMicrosoftGraphTodoTask() // MicrosoftGraphTodoTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoListsUpdateTasks(context.Background(), userId, todoTaskListId, todoTaskId).MicrosoftGraphTodoTask(microsoftGraphTodoTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoListsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 
**todoTaskId** | **string** | key: id of todoTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoListsUpdateTasksRequest struct via the builder pattern


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


## UsersTodoUpdateLists

> UsersTodoUpdateLists(ctx, userId, todoTaskListId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()

Update the navigation property lists in users



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
    userId := "userId_example" // string | key: id of user
    todoTaskListId := "todoTaskListId_example" // string | key: id of todoTaskList
    microsoftGraphTodoTaskList := *openapiclient.NewMicrosoftGraphTodoTaskList() // MicrosoftGraphTodoTaskList | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersTodoUpdateLists(context.Background(), userId, todoTaskListId).MicrosoftGraphTodoTaskList(microsoftGraphTodoTaskList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersTodoUpdateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**todoTaskListId** | **string** | key: id of todoTaskList | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTodoUpdateListsRequest struct via the builder pattern


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


## UsersUpdateTodo

> UsersUpdateTodo(ctx, userId).MicrosoftGraphTodo(microsoftGraphTodo).Execute()

Update the navigation property todo in users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphTodo := *openapiclient.NewMicrosoftGraphTodo() // MicrosoftGraphTodo | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTodoApi.UsersUpdateTodo(context.Background(), userId).MicrosoftGraphTodo(microsoftGraphTodo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTodoApi.UsersUpdateTodo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateTodoRequest struct via the builder pattern


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

