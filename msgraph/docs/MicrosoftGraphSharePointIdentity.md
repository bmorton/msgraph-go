# MicrosoftGraphSharePointIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the identity. | [optional] 
**LoginName** | Pointer to **NullableString** | The sign in name of the SharePoint identity. | [optional] 

## Methods

### NewMicrosoftGraphSharePointIdentity

`func NewMicrosoftGraphSharePointIdentity() *MicrosoftGraphSharePointIdentity`

NewMicrosoftGraphSharePointIdentity instantiates a new MicrosoftGraphSharePointIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharePointIdentityWithDefaults

`func NewMicrosoftGraphSharePointIdentityWithDefaults() *MicrosoftGraphSharePointIdentity`

NewMicrosoftGraphSharePointIdentityWithDefaults instantiates a new MicrosoftGraphSharePointIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphSharePointIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSharePointIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSharePointIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSharePointIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSharePointIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSharePointIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphSharePointIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSharePointIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSharePointIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSharePointIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphSharePointIdentity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphSharePointIdentity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLoginName

`func (o *MicrosoftGraphSharePointIdentity) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MicrosoftGraphSharePointIdentity) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MicrosoftGraphSharePointIdentity) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MicrosoftGraphSharePointIdentity) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### SetLoginNameNil

`func (o *MicrosoftGraphSharePointIdentity) SetLoginNameNil(b bool)`

 SetLoginNameNil sets the value for LoginName to be an explicit nil

### UnsetLoginName
`func (o *MicrosoftGraphSharePointIdentity) UnsetLoginName()`

UnsetLoginName ensures that no value is present for LoginName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


