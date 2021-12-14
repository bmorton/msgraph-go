# \EducationFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationClassesDelta**](EducationFunctionsApi.md#EducationClassesDelta) | **Get** /education/classes/microsoft.graph.delta() | Invoke function delta
[**EducationClassesEducationClassMembersDelta**](EducationFunctionsApi.md#EducationClassesEducationClassMembersDelta) | **Get** /education/classes/{educationClass-id}/members/microsoft.graph.delta() | Invoke function delta
[**EducationClassesEducationClassSchoolsDelta**](EducationFunctionsApi.md#EducationClassesEducationClassSchoolsDelta) | **Get** /education/classes/{educationClass-id}/schools/microsoft.graph.delta() | Invoke function delta
[**EducationClassesEducationClassTeachersDelta**](EducationFunctionsApi.md#EducationClassesEducationClassTeachersDelta) | **Get** /education/classes/{educationClass-id}/teachers/microsoft.graph.delta() | Invoke function delta
[**EducationMeClassesDelta**](EducationFunctionsApi.md#EducationMeClassesDelta) | **Get** /education/me/classes/microsoft.graph.delta() | Invoke function delta
[**EducationMeSchoolsDelta**](EducationFunctionsApi.md#EducationMeSchoolsDelta) | **Get** /education/me/schools/microsoft.graph.delta() | Invoke function delta
[**EducationMeTaughtClassesDelta**](EducationFunctionsApi.md#EducationMeTaughtClassesDelta) | **Get** /education/me/taughtClasses/microsoft.graph.delta() | Invoke function delta
[**EducationSchoolsDelta**](EducationFunctionsApi.md#EducationSchoolsDelta) | **Get** /education/schools/microsoft.graph.delta() | Invoke function delta
[**EducationSchoolsEducationSchoolClassesDelta**](EducationFunctionsApi.md#EducationSchoolsEducationSchoolClassesDelta) | **Get** /education/schools/{educationSchool-id}/classes/microsoft.graph.delta() | Invoke function delta
[**EducationSchoolsEducationSchoolUsersDelta**](EducationFunctionsApi.md#EducationSchoolsEducationSchoolUsersDelta) | **Get** /education/schools/{educationSchool-id}/users/microsoft.graph.delta() | Invoke function delta
[**EducationUsersDelta**](EducationFunctionsApi.md#EducationUsersDelta) | **Get** /education/users/microsoft.graph.delta() | Invoke function delta
[**EducationUsersEducationUserClassesDelta**](EducationFunctionsApi.md#EducationUsersEducationUserClassesDelta) | **Get** /education/users/{educationUser-id}/classes/microsoft.graph.delta() | Invoke function delta
[**EducationUsersEducationUserSchoolsDelta**](EducationFunctionsApi.md#EducationUsersEducationUserSchoolsDelta) | **Get** /education/users/{educationUser-id}/schools/microsoft.graph.delta() | Invoke function delta
[**EducationUsersEducationUserTaughtClassesDelta**](EducationFunctionsApi.md#EducationUsersEducationUserTaughtClassesDelta) | **Get** /education/users/{educationUser-id}/taughtClasses/microsoft.graph.delta() | Invoke function delta



## EducationClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationClassesDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationClassesDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationClassesDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassMembersDelta

> []*AnyOfmicrosoftGraphEducationUser EducationClassesEducationClassMembersDelta(ctx, educationClassId).Execute()

Invoke function delta

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationClassesEducationClassMembersDelta(context.Background(), educationClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationClassesEducationClassMembersDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassMembersDelta`: []*AnyOfmicrosoftGraphEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationClassesEducationClassMembersDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassMembersDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassSchoolsDelta

> []*AnyOfmicrosoftGraphEducationSchool EducationClassesEducationClassSchoolsDelta(ctx, educationClassId).Execute()

Invoke function delta

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationClassesEducationClassSchoolsDelta(context.Background(), educationClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationClassesEducationClassSchoolsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassSchoolsDelta`: []*AnyOfmicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationClassesEducationClassSchoolsDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassSchoolsDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationSchool**](anyOf&lt;microsoft.graph.educationSchool&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassTeachersDelta

> []*AnyOfmicrosoftGraphEducationUser EducationClassesEducationClassTeachersDelta(ctx, educationClassId).Execute()

Invoke function delta

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationClassesEducationClassTeachersDelta(context.Background(), educationClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationClassesEducationClassTeachersDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassTeachersDelta`: []*AnyOfmicrosoftGraphEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationClassesEducationClassTeachersDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassTeachersDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationMeClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationMeClassesDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationMeClassesDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationMeClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationMeClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationMeClassesDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationMeClassesDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationMeSchoolsDelta

> []*AnyOfmicrosoftGraphEducationSchool EducationMeSchoolsDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationMeSchoolsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationMeSchoolsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationMeSchoolsDelta`: []*AnyOfmicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationMeSchoolsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationMeSchoolsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationSchool**](anyOf&lt;microsoft.graph.educationSchool&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationMeTaughtClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationMeTaughtClassesDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationMeTaughtClassesDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationMeTaughtClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationMeTaughtClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationMeTaughtClassesDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationMeTaughtClassesDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsDelta

> []*AnyOfmicrosoftGraphEducationSchool EducationSchoolsDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationSchoolsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationSchoolsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsDelta`: []*AnyOfmicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationSchoolsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationSchool**](anyOf&lt;microsoft.graph.educationSchool&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsEducationSchoolClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationSchoolsEducationSchoolClassesDelta(ctx, educationSchoolId).Execute()

Invoke function delta

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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationSchoolsEducationSchoolClassesDelta(context.Background(), educationSchoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationSchoolsEducationSchoolClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsEducationSchoolClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationSchoolsEducationSchoolClassesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsEducationSchoolClassesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsEducationSchoolUsersDelta

> []*AnyOfmicrosoftGraphEducationUser EducationSchoolsEducationSchoolUsersDelta(ctx, educationSchoolId).Execute()

Invoke function delta

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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationSchoolsEducationSchoolUsersDelta(context.Background(), educationSchoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationSchoolsEducationSchoolUsersDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsEducationSchoolUsersDelta`: []*AnyOfmicrosoftGraphEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationSchoolsEducationSchoolUsersDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsEducationSchoolUsersDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUsersDelta

> []*AnyOfmicrosoftGraphEducationUser EducationUsersDelta(ctx).Execute()

Invoke function delta

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationUsersDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationUsersDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationUsersDelta`: []*AnyOfmicrosoftGraphEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationUsersDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEducationUsersDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUsersEducationUserClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationUsersEducationUserClassesDelta(ctx, educationUserId).Execute()

Invoke function delta

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
    educationUserId := "educationUserId_example" // string | key: id of educationUser

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationUsersEducationUserClassesDelta(context.Background(), educationUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationUsersEducationUserClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationUsersEducationUserClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationUsersEducationUserClassesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string** | key: id of educationUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationUsersEducationUserClassesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUsersEducationUserSchoolsDelta

> []*AnyOfmicrosoftGraphEducationSchool EducationUsersEducationUserSchoolsDelta(ctx, educationUserId).Execute()

Invoke function delta

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
    educationUserId := "educationUserId_example" // string | key: id of educationUser

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationUsersEducationUserSchoolsDelta(context.Background(), educationUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationUsersEducationUserSchoolsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationUsersEducationUserSchoolsDelta`: []*AnyOfmicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationUsersEducationUserSchoolsDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string** | key: id of educationUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationUsersEducationUserSchoolsDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationSchool**](anyOf&lt;microsoft.graph.educationSchool&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationUsersEducationUserTaughtClassesDelta

> []*AnyOfmicrosoftGraphEducationClass EducationUsersEducationUserTaughtClassesDelta(ctx, educationUserId).Execute()

Invoke function delta

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
    educationUserId := "educationUserId_example" // string | key: id of educationUser

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationFunctionsApi.EducationUsersEducationUserTaughtClassesDelta(context.Background(), educationUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationFunctionsApi.EducationUsersEducationUserTaughtClassesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationUsersEducationUserTaughtClassesDelta`: []*AnyOfmicrosoftGraphEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationFunctionsApi.EducationUsersEducationUserTaughtClassesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationUserId** | **string** | key: id of educationUser | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationUsersEducationUserTaughtClassesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphEducationClass**](anyOf&lt;microsoft.graph.educationClass&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

