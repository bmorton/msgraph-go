# MicrosoftGraphShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The identity of the owner of the shared item. Read-only. | [optional] 
**Scope** | Pointer to **NullableString** | Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only. | [optional] 
**SharedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The identity of the user who shared the item. Read-only. | [optional] 
**SharedDateTime** | Pointer to **NullableTime** | The UTC date and time when the item was shared. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphShared

`func NewMicrosoftGraphShared() *MicrosoftGraphShared`

NewMicrosoftGraphShared instantiates a new MicrosoftGraphShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharedWithDefaults

`func NewMicrosoftGraphSharedWithDefaults() *MicrosoftGraphShared`

NewMicrosoftGraphSharedWithDefaults instantiates a new MicrosoftGraphShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *MicrosoftGraphShared) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphShared) GetOwnerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphShared) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphShared) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MicrosoftGraphShared) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MicrosoftGraphShared) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetScope

`func (o *MicrosoftGraphShared) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphShared) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphShared) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphShared) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphShared) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphShared) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSharedBy

`func (o *MicrosoftGraphShared) GetSharedBy() AnyOfmicrosoftGraphIdentitySet`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *MicrosoftGraphShared) GetSharedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *MicrosoftGraphShared) SetSharedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *MicrosoftGraphShared) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### SetSharedByNil

`func (o *MicrosoftGraphShared) SetSharedByNil(b bool)`

 SetSharedByNil sets the value for SharedBy to be an explicit nil

### UnsetSharedBy
`func (o *MicrosoftGraphShared) UnsetSharedBy()`

UnsetSharedBy ensures that no value is present for SharedBy, not even an explicit nil
### GetSharedDateTime

`func (o *MicrosoftGraphShared) GetSharedDateTime() time.Time`

GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.

### GetSharedDateTimeOk

`func (o *MicrosoftGraphShared) GetSharedDateTimeOk() (*time.Time, bool)`

GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDateTime

`func (o *MicrosoftGraphShared) SetSharedDateTime(v time.Time)`

SetSharedDateTime sets SharedDateTime field to given value.

### HasSharedDateTime

`func (o *MicrosoftGraphShared) HasSharedDateTime() bool`

HasSharedDateTime returns a boolean if a field has been set.

### SetSharedDateTimeNil

`func (o *MicrosoftGraphShared) SetSharedDateTimeNil(b bool)`

 SetSharedDateTimeNil sets the value for SharedDateTime to be an explicit nil

### UnsetSharedDateTime
`func (o *MicrosoftGraphShared) UnsetSharedDateTime()`

UnsetSharedDateTime ensures that no value is present for SharedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


