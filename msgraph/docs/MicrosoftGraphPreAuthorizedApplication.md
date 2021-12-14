# MicrosoftGraphPreAuthorizedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableString** | The unique identifier for the application. | [optional] 
**DelegatedPermissionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMicrosoftGraphPreAuthorizedApplication

`func NewMicrosoftGraphPreAuthorizedApplication() *MicrosoftGraphPreAuthorizedApplication`

NewMicrosoftGraphPreAuthorizedApplication instantiates a new MicrosoftGraphPreAuthorizedApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPreAuthorizedApplicationWithDefaults

`func NewMicrosoftGraphPreAuthorizedApplicationWithDefaults() *MicrosoftGraphPreAuthorizedApplication`

NewMicrosoftGraphPreAuthorizedApplicationWithDefaults instantiates a new MicrosoftGraphPreAuthorizedApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *MicrosoftGraphPreAuthorizedApplication) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphPreAuthorizedApplication) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphPreAuthorizedApplication) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphPreAuthorizedApplication) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphPreAuthorizedApplication) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphPreAuthorizedApplication) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetDelegatedPermissionIds

`func (o *MicrosoftGraphPreAuthorizedApplication) GetDelegatedPermissionIds() []string`

GetDelegatedPermissionIds returns the DelegatedPermissionIds field if non-nil, zero value otherwise.

### GetDelegatedPermissionIdsOk

`func (o *MicrosoftGraphPreAuthorizedApplication) GetDelegatedPermissionIdsOk() (*[]string, bool)`

GetDelegatedPermissionIdsOk returns a tuple with the DelegatedPermissionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedPermissionIds

`func (o *MicrosoftGraphPreAuthorizedApplication) SetDelegatedPermissionIds(v []string)`

SetDelegatedPermissionIds sets DelegatedPermissionIds field to given value.

### HasDelegatedPermissionIds

`func (o *MicrosoftGraphPreAuthorizedApplication) HasDelegatedPermissionIds() bool`

HasDelegatedPermissionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


