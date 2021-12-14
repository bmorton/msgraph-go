# MicrosoftGraphMessageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Actions** | Pointer to [**AnyOfmicrosoftGraphMessageRuleActions**](anyOf&lt;microsoft.graph.messageRuleActions&gt;.md) | Actions to be taken on a message when the corresponding conditions are fulfilled. | [optional] 
**Conditions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) | Conditions that when fulfilled, will trigger the corresponding actions for that rule. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the rule. | [optional] 
**Exceptions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) | Exception conditions for the rule. | [optional] 
**HasError** | Pointer to **NullableBool** | Indicates whether the rule is in an error condition. Read-only. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether the rule is enabled to be applied to messages. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. | [optional] 
**Sequence** | Pointer to **NullableInt32** | Indicates the order in which the rule is executed, among other rules. | [optional] 

## Methods

### NewMicrosoftGraphMessageRule

`func NewMicrosoftGraphMessageRule() *MicrosoftGraphMessageRule`

NewMicrosoftGraphMessageRule instantiates a new MicrosoftGraphMessageRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMessageRuleWithDefaults

`func NewMicrosoftGraphMessageRuleWithDefaults() *MicrosoftGraphMessageRule`

NewMicrosoftGraphMessageRuleWithDefaults instantiates a new MicrosoftGraphMessageRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMessageRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMessageRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMessageRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMessageRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActions

`func (o *MicrosoftGraphMessageRule) GetActions() AnyOfmicrosoftGraphMessageRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MicrosoftGraphMessageRule) GetActionsOk() (*AnyOfmicrosoftGraphMessageRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MicrosoftGraphMessageRule) SetActions(v AnyOfmicrosoftGraphMessageRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MicrosoftGraphMessageRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *MicrosoftGraphMessageRule) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *MicrosoftGraphMessageRule) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetConditions

`func (o *MicrosoftGraphMessageRule) GetConditions() AnyOfmicrosoftGraphMessageRulePredicates`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MicrosoftGraphMessageRule) GetConditionsOk() (*AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MicrosoftGraphMessageRule) SetConditions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MicrosoftGraphMessageRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *MicrosoftGraphMessageRule) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *MicrosoftGraphMessageRule) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphMessageRule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMessageRule) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphMessageRule) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphMessageRule) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphMessageRule) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphMessageRule) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExceptions

`func (o *MicrosoftGraphMessageRule) GetExceptions() AnyOfmicrosoftGraphMessageRulePredicates`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *MicrosoftGraphMessageRule) GetExceptionsOk() (*AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *MicrosoftGraphMessageRule) SetExceptions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *MicrosoftGraphMessageRule) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### SetExceptionsNil

`func (o *MicrosoftGraphMessageRule) SetExceptionsNil(b bool)`

 SetExceptionsNil sets the value for Exceptions to be an explicit nil

### UnsetExceptions
`func (o *MicrosoftGraphMessageRule) UnsetExceptions()`

UnsetExceptions ensures that no value is present for Exceptions, not even an explicit nil
### GetHasError

`func (o *MicrosoftGraphMessageRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *MicrosoftGraphMessageRule) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *MicrosoftGraphMessageRule) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *MicrosoftGraphMessageRule) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### SetHasErrorNil

`func (o *MicrosoftGraphMessageRule) SetHasErrorNil(b bool)`

 SetHasErrorNil sets the value for HasError to be an explicit nil

### UnsetHasError
`func (o *MicrosoftGraphMessageRule) UnsetHasError()`

UnsetHasError ensures that no value is present for HasError, not even an explicit nil
### GetIsEnabled

`func (o *MicrosoftGraphMessageRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphMessageRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphMessageRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphMessageRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MicrosoftGraphMessageRule) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MicrosoftGraphMessageRule) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsReadOnly

`func (o *MicrosoftGraphMessageRule) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MicrosoftGraphMessageRule) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MicrosoftGraphMessageRule) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MicrosoftGraphMessageRule) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *MicrosoftGraphMessageRule) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *MicrosoftGraphMessageRule) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetSequence

`func (o *MicrosoftGraphMessageRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MicrosoftGraphMessageRule) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MicrosoftGraphMessageRule) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MicrosoftGraphMessageRule) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *MicrosoftGraphMessageRule) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *MicrosoftGraphMessageRule) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


