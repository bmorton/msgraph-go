# MicrosoftGraphAppliedConditionalAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Refers to the Name of the conditional access policy (example: &#39;Require MFA for Salesforce&#39;). | [optional] 
**EnforcedGrantControls** | Pointer to **[]string** | Refers to the grant controls enforced by the conditional access policy (example: &#39;Require multi-factor authentication&#39;). | [optional] 
**EnforcedSessionControls** | Pointer to **[]string** | Refers to the session controls enforced by the conditional access policy (example: &#39;Require app enforced controls&#39;). | [optional] 
**Id** | Pointer to **NullableString** | An identifier of the conditional access policy. | [optional] 
**Result** | Pointer to [**AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult**](anyOf&lt;microsoft.graph.appliedConditionalAccessPolicyResult&gt;.md) | Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn&#39;t applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphAppliedConditionalAccessPolicy

`func NewMicrosoftGraphAppliedConditionalAccessPolicy() *MicrosoftGraphAppliedConditionalAccessPolicy`

NewMicrosoftGraphAppliedConditionalAccessPolicy instantiates a new MicrosoftGraphAppliedConditionalAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppliedConditionalAccessPolicyWithDefaults

`func NewMicrosoftGraphAppliedConditionalAccessPolicyWithDefaults() *MicrosoftGraphAppliedConditionalAccessPolicy`

NewMicrosoftGraphAppliedConditionalAccessPolicyWithDefaults instantiates a new MicrosoftGraphAppliedConditionalAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedGrantControls() []*string`

GetEnforcedGrantControls returns the EnforcedGrantControls field if non-nil, zero value otherwise.

### GetEnforcedGrantControlsOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedGrantControlsOk() (*[]*string, bool)`

GetEnforcedGrantControlsOk returns a tuple with the EnforcedGrantControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetEnforcedGrantControls(v []*string)`

SetEnforcedGrantControls sets EnforcedGrantControls field to given value.

### HasEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasEnforcedGrantControls() bool`

HasEnforcedGrantControls returns a boolean if a field has been set.

### GetEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedSessionControls() []*string`

GetEnforcedSessionControls returns the EnforcedSessionControls field if non-nil, zero value otherwise.

### GetEnforcedSessionControlsOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedSessionControlsOk() (*[]*string, bool)`

GetEnforcedSessionControlsOk returns a tuple with the EnforcedSessionControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetEnforcedSessionControls(v []*string)`

SetEnforcedSessionControls sets EnforcedSessionControls field to given value.

### HasEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasEnforcedSessionControls() bool`

HasEnforcedSessionControls returns a boolean if a field has been set.

### GetId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetResult() AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetResultOk() (*AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetResult(v AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


