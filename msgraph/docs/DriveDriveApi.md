# \DriveDriveApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveDriveGetDrive**](DriveDriveApi.md#DriveDriveGetDrive) | **Get** /drive | Get drive
[**DriveDriveUpdateDrive**](DriveDriveApi.md#DriveDriveUpdateDrive) | **Patch** /drive | Update drive



## DriveDriveGetDrive

> MicrosoftGraphDrive DriveDriveGetDrive(ctx).Select_(select_).Expand(expand).Execute()

Get drive

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
    resp, r, err := api_client.DriveDriveApi.DriveDriveGetDrive(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveApi.DriveDriveGetDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveDriveGetDrive`: MicrosoftGraphDrive
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveApi.DriveDriveGetDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveDriveGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveDriveUpdateDrive

> DriveDriveUpdateDrive(ctx).MicrosoftGraphDrive(microsoftGraphDrive).Execute()

Update drive

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
    microsoftGraphDrive := *openapiclient.NewMicrosoftGraphDrive() // MicrosoftGraphDrive | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveApi.DriveDriveUpdateDrive(context.Background()).MicrosoftGraphDrive(microsoftGraphDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveApi.DriveDriveUpdateDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveDriveUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | New property values | 

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

