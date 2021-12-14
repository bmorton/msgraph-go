# ConditionalAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**MicrosoftGraphConditionalAccessConditionSet**](MicrosoftGraphConditionalAccessConditionSet.md) |  | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly. | [optional] 
**Description** | Pointer to **NullableString** | Not used. | [optional] 
**DisplayName** | Pointer to **string** | Specifies a display name for the conditionalAccessPolicy object. | [optional] 
**GrantControls** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessGrantControls**](anyOf&lt;microsoft.graph.conditionalAccessGrantControls&gt;.md) | Specifies the grant controls that must be fulfilled to pass the policy. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly. | [optional] 
**SessionControls** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessSessionControls**](anyOf&lt;microsoft.graph.conditionalAccessSessionControls&gt;.md) | Specifies the session controls that are enforced after sign-in. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessPolicyState**](anyOf&lt;microsoft.graph.conditionalAccessPolicyState&gt;.md) | Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required. | [optional] 

## Methods

### NewConditionalAccessPolicy

`func NewConditionalAccessPolicy() *ConditionalAccessPolicy`

NewConditionalAccessPolicy instantiates a new ConditionalAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalAccessPolicyWithDefaults

`func NewConditionalAccessPolicyWithDefaults() *ConditionalAccessPolicy`

NewConditionalAccessPolicyWithDefaults instantiates a new ConditionalAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ConditionalAccessPolicy) GetConditions() MicrosoftGraphConditionalAccessConditionSet`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ConditionalAccessPolicy) GetConditionsOk() (*MicrosoftGraphConditionalAccessConditionSet, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ConditionalAccessPolicy) SetConditions(v MicrosoftGraphConditionalAccessConditionSet)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ConditionalAccessPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *ConditionalAccessPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ConditionalAccessPolicy) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ConditionalAccessPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ConditionalAccessPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *ConditionalAccessPolicy) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *ConditionalAccessPolicy) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *ConditionalAccessPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConditionalAccessPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConditionalAccessPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConditionalAccessPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConditionalAccessPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConditionalAccessPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ConditionalAccessPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConditionalAccessPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConditionalAccessPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConditionalAccessPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGrantControls

`func (o *ConditionalAccessPolicy) GetGrantControls() AnyOfmicrosoftGraphConditionalAccessGrantControls`

GetGrantControls returns the GrantControls field if non-nil, zero value otherwise.

### GetGrantControlsOk

`func (o *ConditionalAccessPolicy) GetGrantControlsOk() (*AnyOfmicrosoftGraphConditionalAccessGrantControls, bool)`

GetGrantControlsOk returns a tuple with the GrantControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantControls

`func (o *ConditionalAccessPolicy) SetGrantControls(v AnyOfmicrosoftGraphConditionalAccessGrantControls)`

SetGrantControls sets GrantControls field to given value.

### HasGrantControls

`func (o *ConditionalAccessPolicy) HasGrantControls() bool`

HasGrantControls returns a boolean if a field has been set.

### SetGrantControlsNil

`func (o *ConditionalAccessPolicy) SetGrantControlsNil(b bool)`

 SetGrantControlsNil sets the value for GrantControls to be an explicit nil

### UnsetGrantControls
`func (o *ConditionalAccessPolicy) UnsetGrantControls()`

UnsetGrantControls ensures that no value is present for GrantControls, not even an explicit nil
### GetModifiedDateTime

`func (o *ConditionalAccessPolicy) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *ConditionalAccessPolicy) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *ConditionalAccessPolicy) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *ConditionalAccessPolicy) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *ConditionalAccessPolicy) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *ConditionalAccessPolicy) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetSessionControls

`func (o *ConditionalAccessPolicy) GetSessionControls() AnyOfmicrosoftGraphConditionalAccessSessionControls`

GetSessionControls returns the SessionControls field if non-nil, zero value otherwise.

### GetSessionControlsOk

`func (o *ConditionalAccessPolicy) GetSessionControlsOk() (*AnyOfmicrosoftGraphConditionalAccessSessionControls, bool)`

GetSessionControlsOk returns a tuple with the SessionControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionControls

`func (o *ConditionalAccessPolicy) SetSessionControls(v AnyOfmicrosoftGraphConditionalAccessSessionControls)`

SetSessionControls sets SessionControls field to given value.

### HasSessionControls

`func (o *ConditionalAccessPolicy) HasSessionControls() bool`

HasSessionControls returns a boolean if a field has been set.

### SetSessionControlsNil

`func (o *ConditionalAccessPolicy) SetSessionControlsNil(b bool)`

 SetSessionControlsNil sets the value for SessionControls to be an explicit nil

### UnsetSessionControls
`func (o *ConditionalAccessPolicy) UnsetSessionControls()`

UnsetSessionControls ensures that no value is present for SessionControls, not even an explicit nil
### GetState

`func (o *ConditionalAccessPolicy) GetState() AnyOfmicrosoftGraphConditionalAccessPolicyState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConditionalAccessPolicy) GetStateOk() (*AnyOfmicrosoftGraphConditionalAccessPolicyState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConditionalAccessPolicy) SetState(v AnyOfmicrosoftGraphConditionalAccessPolicyState)`

SetState sets State field to given value.

### HasState

`func (o *ConditionalAccessPolicy) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ConditionalAccessPolicy) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ConditionalAccessPolicy) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


