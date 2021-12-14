# MicrosoftGraphPermissionGrantPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | Description for this policy. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for this policy. Required. | [optional] 
**Excludes** | Pointer to [**[]MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | Condition sets which are excluded in this permission grant policy. Automatically expanded on GET. | [optional] 
**Includes** | Pointer to [**[]MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | Condition sets which are included in this permission grant policy. Automatically expanded on GET. | [optional] 

## Methods

### NewMicrosoftGraphPermissionGrantPolicy

`func NewMicrosoftGraphPermissionGrantPolicy() *MicrosoftGraphPermissionGrantPolicy`

NewMicrosoftGraphPermissionGrantPolicy instantiates a new MicrosoftGraphPermissionGrantPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPermissionGrantPolicyWithDefaults

`func NewMicrosoftGraphPermissionGrantPolicyWithDefaults() *MicrosoftGraphPermissionGrantPolicy`

NewMicrosoftGraphPermissionGrantPolicyWithDefaults instantiates a new MicrosoftGraphPermissionGrantPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPermissionGrantPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPermissionGrantPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPermissionGrantPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphPermissionGrantPolicy) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphPermissionGrantPolicy) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPermissionGrantPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphPermissionGrantPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPermissionGrantPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphPermissionGrantPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphPermissionGrantPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExcludes

`func (o *MicrosoftGraphPermissionGrantPolicy) GetExcludes() []MicrosoftGraphPermissionGrantConditionSet`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetExcludesOk() (*[]MicrosoftGraphPermissionGrantConditionSet, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *MicrosoftGraphPermissionGrantPolicy) SetExcludes(v []MicrosoftGraphPermissionGrantConditionSet)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *MicrosoftGraphPermissionGrantPolicy) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetIncludes

`func (o *MicrosoftGraphPermissionGrantPolicy) GetIncludes() []MicrosoftGraphPermissionGrantConditionSet`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *MicrosoftGraphPermissionGrantPolicy) GetIncludesOk() (*[]MicrosoftGraphPermissionGrantConditionSet, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *MicrosoftGraphPermissionGrantPolicy) SetIncludes(v []MicrosoftGraphPermissionGrantConditionSet)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *MicrosoftGraphPermissionGrantPolicy) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


