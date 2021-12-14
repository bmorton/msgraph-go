# SharePointIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginName** | Pointer to **NullableString** | The sign in name of the SharePoint identity. | [optional] 

## Methods

### NewSharePointIdentity

`func NewSharePointIdentity() *SharePointIdentity`

NewSharePointIdentity instantiates a new SharePointIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePointIdentityWithDefaults

`func NewSharePointIdentityWithDefaults() *SharePointIdentity`

NewSharePointIdentityWithDefaults instantiates a new SharePointIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginName

`func (o *SharePointIdentity) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *SharePointIdentity) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *SharePointIdentity) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *SharePointIdentity) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### SetLoginNameNil

`func (o *SharePointIdentity) SetLoginNameNil(b bool)`

 SetLoginNameNil sets the value for LoginName to be an explicit nil

### UnsetLoginName
`func (o *SharePointIdentity) UnsetLoginName()`

UnsetLoginName ensures that no value is present for LoginName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


