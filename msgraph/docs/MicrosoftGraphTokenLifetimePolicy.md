# MicrosoftGraphTokenLifetimePolicy

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

### NewMicrosoftGraphTokenLifetimePolicy

`func NewMicrosoftGraphTokenLifetimePolicy() *MicrosoftGraphTokenLifetimePolicy`

NewMicrosoftGraphTokenLifetimePolicy instantiates a new MicrosoftGraphTokenLifetimePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTokenLifetimePolicyWithDefaults

`func NewMicrosoftGraphTokenLifetimePolicyWithDefaults() *MicrosoftGraphTokenLifetimePolicy`

NewMicrosoftGraphTokenLifetimePolicyWithDefaults instantiates a new MicrosoftGraphTokenLifetimePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTokenLifetimePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTokenLifetimePolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTokenLifetimePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphTokenLifetimePolicy) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphTokenLifetimePolicy) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTokenLifetimePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTokenLifetimePolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTokenLifetimePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTokenLifetimePolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDefinition

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDefinition() []string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetDefinitionOk() (*[]string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *MicrosoftGraphTokenLifetimePolicy) SetDefinition(v []string)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *MicrosoftGraphTokenLifetimePolicy) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetIsOrganizationDefault

`func (o *MicrosoftGraphTokenLifetimePolicy) GetIsOrganizationDefault() bool`

GetIsOrganizationDefault returns the IsOrganizationDefault field if non-nil, zero value otherwise.

### GetIsOrganizationDefaultOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetIsOrganizationDefaultOk() (*bool, bool)`

GetIsOrganizationDefaultOk returns a tuple with the IsOrganizationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationDefault

`func (o *MicrosoftGraphTokenLifetimePolicy) SetIsOrganizationDefault(v bool)`

SetIsOrganizationDefault sets IsOrganizationDefault field to given value.

### HasIsOrganizationDefault

`func (o *MicrosoftGraphTokenLifetimePolicy) HasIsOrganizationDefault() bool`

HasIsOrganizationDefault returns a boolean if a field has been set.

### SetIsOrganizationDefaultNil

`func (o *MicrosoftGraphTokenLifetimePolicy) SetIsOrganizationDefaultNil(b bool)`

 SetIsOrganizationDefaultNil sets the value for IsOrganizationDefault to be an explicit nil

### UnsetIsOrganizationDefault
`func (o *MicrosoftGraphTokenLifetimePolicy) UnsetIsOrganizationDefault()`

UnsetIsOrganizationDefault ensures that no value is present for IsOrganizationDefault, not even an explicit nil
### GetAppliesTo

`func (o *MicrosoftGraphTokenLifetimePolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MicrosoftGraphTokenLifetimePolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MicrosoftGraphTokenLifetimePolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *MicrosoftGraphTokenLifetimePolicy) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


