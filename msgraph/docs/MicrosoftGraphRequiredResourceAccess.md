# MicrosoftGraphRequiredResourceAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceAccess** | Pointer to [**[]MicrosoftGraphResourceAccess**](MicrosoftGraphResourceAccess.md) | The list of OAuth2.0 permission scopes and app roles that the application requires from the specified resource. | [optional] 
**ResourceAppId** | Pointer to **string** | The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application. | [optional] 

## Methods

### NewMicrosoftGraphRequiredResourceAccess

`func NewMicrosoftGraphRequiredResourceAccess() *MicrosoftGraphRequiredResourceAccess`

NewMicrosoftGraphRequiredResourceAccess instantiates a new MicrosoftGraphRequiredResourceAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRequiredResourceAccessWithDefaults

`func NewMicrosoftGraphRequiredResourceAccessWithDefaults() *MicrosoftGraphRequiredResourceAccess`

NewMicrosoftGraphRequiredResourceAccessWithDefaults instantiates a new MicrosoftGraphRequiredResourceAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceAccess

`func (o *MicrosoftGraphRequiredResourceAccess) GetResourceAccess() []MicrosoftGraphResourceAccess`

GetResourceAccess returns the ResourceAccess field if non-nil, zero value otherwise.

### GetResourceAccessOk

`func (o *MicrosoftGraphRequiredResourceAccess) GetResourceAccessOk() (*[]MicrosoftGraphResourceAccess, bool)`

GetResourceAccessOk returns a tuple with the ResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAccess

`func (o *MicrosoftGraphRequiredResourceAccess) SetResourceAccess(v []MicrosoftGraphResourceAccess)`

SetResourceAccess sets ResourceAccess field to given value.

### HasResourceAccess

`func (o *MicrosoftGraphRequiredResourceAccess) HasResourceAccess() bool`

HasResourceAccess returns a boolean if a field has been set.

### GetResourceAppId

`func (o *MicrosoftGraphRequiredResourceAccess) GetResourceAppId() string`

GetResourceAppId returns the ResourceAppId field if non-nil, zero value otherwise.

### GetResourceAppIdOk

`func (o *MicrosoftGraphRequiredResourceAccess) GetResourceAppIdOk() (*string, bool)`

GetResourceAppIdOk returns a tuple with the ResourceAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAppId

`func (o *MicrosoftGraphRequiredResourceAccess) SetResourceAppId(v string)`

SetResourceAppId sets ResourceAppId field to given value.

### HasResourceAppId

`func (o *MicrosoftGraphRequiredResourceAccess) HasResourceAppId() bool`

HasResourceAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


