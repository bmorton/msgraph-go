# \EducationEducationRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationEducationRootGetEducationRoot**](EducationEducationRootApi.md#EducationEducationRootGetEducationRoot) | **Get** /education | Get education
[**EducationEducationRootUpdateEducationRoot**](EducationEducationRootApi.md#EducationEducationRootUpdateEducationRoot) | **Patch** /education | Update education



## EducationEducationRootGetEducationRoot

> MicrosoftGraphEducationRoot EducationEducationRootGetEducationRoot(ctx).Select_(select_).Expand(expand).Execute()

Get education

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
    resp, r, err := api_client.EducationEducationRootApi.EducationEducationRootGetEducationRoot(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationRootApi.EducationEducationRootGetEducationRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationEducationRootGetEducationRoot`: MicrosoftGraphEducationRoot
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationRootApi.EducationEducationRootGetEducationRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEducationEducationRootGetEducationRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEducationRoot**](MicrosoftGraphEducationRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationEducationRootUpdateEducationRoot

> EducationEducationRootUpdateEducationRoot(ctx).MicrosoftGraphEducationRoot(microsoftGraphEducationRoot).Execute()

Update education

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
    microsoftGraphEducationRoot := *openapiclient.NewMicrosoftGraphEducationRoot() // MicrosoftGraphEducationRoot | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationRootApi.EducationEducationRootUpdateEducationRoot(context.Background()).MicrosoftGraphEducationRoot(microsoftGraphEducationRoot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationRootApi.EducationEducationRootUpdateEducationRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEducationEducationRootUpdateEducationRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphEducationRoot** | [**MicrosoftGraphEducationRoot**](MicrosoftGraphEducationRoot.md) | New property values | 

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

