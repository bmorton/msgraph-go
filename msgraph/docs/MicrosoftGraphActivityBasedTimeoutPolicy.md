# MicrosoftGraphActivityBasedTimeoutPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | Description for this policy. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for this policy. Required. | [optional] 
**Definition** | Pointer to **[]string** | A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required. | [optional] 
**IsOrganizationDefault** | Pointer to **NullableBool** | If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false. | [optional] 
**AppliesTo** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 

## Methods

### NewMicrosoftGraphActivityBasedTimeoutPolicy

`func NewMicrosoftGraphActivityBasedTimeoutPolicy() *MicrosoftGraphActivityBasedTimeoutPolicy`

NewMicrosoftGraphActivityBasedTimeoutPolicy instantiates a new MicrosoftGraphActivityBasedTimeoutPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphActivityBasedTimeoutPolicyWithDefaults

`func NewMicrosoftGraphActivityBasedTimeoutPolicyWithDefaults() *MicrosoftGraphActivityBasedTimeoutPolicy`

NewMicrosoftGraphActivityBasedTimeoutPolicyWithDefaults instantiates a new MicrosoftGraphActivityBasedTimeoutPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDefinition

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDefinition() []string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetDefinitionOk() (*[]string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetDefinition(v []string)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetIsOrganizationDefault

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetIsOrganizationDefault() bool`

GetIsOrganizationDefault returns the IsOrganizationDefault field if non-nil, zero value otherwise.

### GetIsOrganizationDefaultOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetIsOrganizationDefaultOk() (*bool, bool)`

GetIsOrganizationDefaultOk returns a tuple with the IsOrganizationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationDefault

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetIsOrganizationDefault(v bool)`

SetIsOrganizationDefault sets IsOrganizationDefault field to given value.

### HasIsOrganizationDefault

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasIsOrganizationDefault() bool`

HasIsOrganizationDefault returns a boolean if a field has been set.

### SetIsOrganizationDefaultNil

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetIsOrganizationDefaultNil(b bool)`

 SetIsOrganizationDefaultNil sets the value for IsOrganizationDefault to be an explicit nil

### UnsetIsOrganizationDefault
`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) UnsetIsOrganizationDefault()`

UnsetIsOrganizationDefault ensures that no value is present for IsOrganizationDefault, not even an explicit nil
### GetAppliesTo

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *MicrosoftGraphActivityBasedTimeoutPolicy) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


