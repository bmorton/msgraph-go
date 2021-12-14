# MicrosoftGraphConditionalAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Conditions** | Pointer to [**MicrosoftGraphConditionalAccessConditionSet**](MicrosoftGraphConditionalAccessConditionSet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly. | [optional] 
**Description** | Pointer to **NullableString** | Not used. | [optional] 
**DisplayName** | Pointer to **string** | Specifies a display name for the conditionalAccessPolicy object. | [optional] 
**GrantControls** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessGrantControls**](anyOf&lt;microsoft.graph.conditionalAccessGrantControls&gt;.md) | Specifies the grant controls that must be fulfilled to pass the policy. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly. | [optional] 
**SessionControls** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessSessionControls**](anyOf&lt;microsoft.graph.conditionalAccessSessionControls&gt;.md) | Specifies the session controls that are enforced after sign-in. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessPolicyState**](anyOf&lt;microsoft.graph.conditionalAccessPolicyState&gt;.md) | Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessPolicy

`func NewMicrosoftGraphConditionalAccessPolicy() *MicrosoftGraphConditionalAccessPolicy`

NewMicrosoftGraphConditionalAccessPolicy instantiates a new MicrosoftGraphConditionalAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessPolicyWithDefaults

`func NewMicrosoftGraphConditionalAccessPolicyWithDefaults() *MicrosoftGraphConditionalAccessPolicy`

NewMicrosoftGraphConditionalAccessPolicyWithDefaults instantiates a new MicrosoftGraphConditionalAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphConditionalAccessPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphConditionalAccessPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphConditionalAccessPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConditions

`func (o *MicrosoftGraphConditionalAccessPolicy) GetConditions() MicrosoftGraphConditionalAccessConditionSet`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetConditionsOk() (*MicrosoftGraphConditionalAccessConditionSet, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MicrosoftGraphConditionalAccessPolicy) SetConditions(v MicrosoftGraphConditionalAccessConditionSet)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MicrosoftGraphConditionalAccessPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphConditionalAccessPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphConditionalAccessPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphConditionalAccessPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphConditionalAccessPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphConditionalAccessPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphConditionalAccessPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGrantControls

`func (o *MicrosoftGraphConditionalAccessPolicy) GetGrantControls() AnyOfmicrosoftGraphConditionalAccessGrantControls`

GetGrantControls returns the GrantControls field if non-nil, zero value otherwise.

### GetGrantControlsOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetGrantControlsOk() (*AnyOfmicrosoftGraphConditionalAccessGrantControls, bool)`

GetGrantControlsOk returns a tuple with the GrantControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantControls

`func (o *MicrosoftGraphConditionalAccessPolicy) SetGrantControls(v AnyOfmicrosoftGraphConditionalAccessGrantControls)`

SetGrantControls sets GrantControls field to given value.

### HasGrantControls

`func (o *MicrosoftGraphConditionalAccessPolicy) HasGrantControls() bool`

HasGrantControls returns a boolean if a field has been set.

### SetGrantControlsNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetGrantControlsNil(b bool)`

 SetGrantControlsNil sets the value for GrantControls to be an explicit nil

### UnsetGrantControls
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetGrantControls()`

UnsetGrantControls ensures that no value is present for GrantControls, not even an explicit nil
### GetModifiedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *MicrosoftGraphConditionalAccessPolicy) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetSessionControls

`func (o *MicrosoftGraphConditionalAccessPolicy) GetSessionControls() AnyOfmicrosoftGraphConditionalAccessSessionControls`

GetSessionControls returns the SessionControls field if non-nil, zero value otherwise.

### GetSessionControlsOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetSessionControlsOk() (*AnyOfmicrosoftGraphConditionalAccessSessionControls, bool)`

GetSessionControlsOk returns a tuple with the SessionControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionControls

`func (o *MicrosoftGraphConditionalAccessPolicy) SetSessionControls(v AnyOfmicrosoftGraphConditionalAccessSessionControls)`

SetSessionControls sets SessionControls field to given value.

### HasSessionControls

`func (o *MicrosoftGraphConditionalAccessPolicy) HasSessionControls() bool`

HasSessionControls returns a boolean if a field has been set.

### SetSessionControlsNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetSessionControlsNil(b bool)`

 SetSessionControlsNil sets the value for SessionControls to be an explicit nil

### UnsetSessionControls
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetSessionControls()`

UnsetSessionControls ensures that no value is present for SessionControls, not even an explicit nil
### GetState

`func (o *MicrosoftGraphConditionalAccessPolicy) GetState() AnyOfmicrosoftGraphConditionalAccessPolicyState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphConditionalAccessPolicy) GetStateOk() (*AnyOfmicrosoftGraphConditionalAccessPolicyState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphConditionalAccessPolicy) SetState(v AnyOfmicrosoftGraphConditionalAccessPolicyState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphConditionalAccessPolicy) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphConditionalAccessPolicy) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphConditionalAccessPolicy) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


