# MessageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AnyOfmicrosoftGraphMessageRuleActions**](anyOf&lt;microsoft.graph.messageRuleActions&gt;.md) | Actions to be taken on a message when the corresponding conditions are fulfilled. | [optional] 
**Conditions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) | Conditions that when fulfilled, will trigger the corresponding actions for that rule. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the rule. | [optional] 
**Exceptions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) | Exception conditions for the rule. | [optional] 
**HasError** | Pointer to **NullableBool** | Indicates whether the rule is in an error condition. Read-only. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether the rule is enabled to be applied to messages. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. | [optional] 
**Sequence** | Pointer to **NullableInt32** | Indicates the order in which the rule is executed, among other rules. | [optional] 

## Methods

### NewMessageRule

`func NewMessageRule() *MessageRule`

NewMessageRule instantiates a new MessageRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageRuleWithDefaults

`func NewMessageRuleWithDefaults() *MessageRule`

NewMessageRuleWithDefaults instantiates a new MessageRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *MessageRule) GetActions() AnyOfmicrosoftGraphMessageRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MessageRule) GetActionsOk() (*AnyOfmicrosoftGraphMessageRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MessageRule) SetActions(v AnyOfmicrosoftGraphMessageRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *MessageRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *MessageRule) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *MessageRule) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetConditions

`func (o *MessageRule) GetConditions() AnyOfmicrosoftGraphMessageRulePredicates`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MessageRule) GetConditionsOk() (*AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MessageRule) SetConditions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MessageRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *MessageRule) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *MessageRule) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetDisplayName

`func (o *MessageRule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MessageRule) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MessageRule) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MessageRule) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MessageRule) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MessageRule) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExceptions

`func (o *MessageRule) GetExceptions() AnyOfmicrosoftGraphMessageRulePredicates`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *MessageRule) GetExceptionsOk() (*AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *MessageRule) SetExceptions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *MessageRule) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### SetExceptionsNil

`func (o *MessageRule) SetExceptionsNil(b bool)`

 SetExceptionsNil sets the value for Exceptions to be an explicit nil

### UnsetExceptions
`func (o *MessageRule) UnsetExceptions()`

UnsetExceptions ensures that no value is present for Exceptions, not even an explicit nil
### GetHasError

`func (o *MessageRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *MessageRule) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *MessageRule) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *MessageRule) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### SetHasErrorNil

`func (o *MessageRule) SetHasErrorNil(b bool)`

 SetHasErrorNil sets the value for HasError to be an explicit nil

### UnsetHasError
`func (o *MessageRule) UnsetHasError()`

UnsetHasError ensures that no value is present for HasError, not even an explicit nil
### GetIsEnabled

`func (o *MessageRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MessageRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MessageRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MessageRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MessageRule) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MessageRule) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsReadOnly

`func (o *MessageRule) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MessageRule) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *MessageRule) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *MessageRule) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *MessageRule) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *MessageRule) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetSequence

`func (o *MessageRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MessageRule) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MessageRule) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MessageRule) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *MessageRule) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *MessageRule) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


