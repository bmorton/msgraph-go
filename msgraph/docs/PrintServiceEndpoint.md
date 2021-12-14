# PrintServiceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A human-readable display name for the endpoint. | [optional] 
**Uri** | Pointer to **string** | The URI that can be used to access the service. | [optional] 

## Methods

### NewPrintServiceEndpoint

`func NewPrintServiceEndpoint() *PrintServiceEndpoint`

NewPrintServiceEndpoint instantiates a new PrintServiceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintServiceEndpointWithDefaults

`func NewPrintServiceEndpointWithDefaults() *PrintServiceEndpoint`

NewPrintServiceEndpointWithDefaults instantiates a new PrintServiceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PrintServiceEndpoint) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrintServiceEndpoint) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrintServiceEndpoint) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrintServiceEndpoint) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUri

`func (o *PrintServiceEndpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PrintServiceEndpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PrintServiceEndpoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PrintServiceEndpoint) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


