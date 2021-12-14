# MicrosoftGraphAdminConsentRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**IsEnabled** | Pointer to **bool** | Specifies whether the admin consent request feature is enabled or disabled. Required. | [optional] 
**NotifyReviewers** | Pointer to **bool** | Specifies whether reviewers will receive notifications. Required. | [optional] 
**RemindersEnabled** | Pointer to **bool** | Specifies whether reviewers will receive reminder emails. Required. | [optional] 
**RequestDurationInDays** | Pointer to **int32** | Specifies the duration the request is active before it automatically expires if no decision is applied. | [optional] 
**Reviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | The list of reviewers for the admin consent. Required. | [optional] 
**Version** | Pointer to **int32** | Specifies the version of this policy. When the policy is updated, this version is updated. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphAdminConsentRequestPolicy

`func NewMicrosoftGraphAdminConsentRequestPolicy() *MicrosoftGraphAdminConsentRequestPolicy`

NewMicrosoftGraphAdminConsentRequestPolicy instantiates a new MicrosoftGraphAdminConsentRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAdminConsentRequestPolicyWithDefaults

`func NewMicrosoftGraphAdminConsentRequestPolicyWithDefaults() *MicrosoftGraphAdminConsentRequestPolicy`

NewMicrosoftGraphAdminConsentRequestPolicyWithDefaults instantiates a new MicrosoftGraphAdminConsentRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetNotifyReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetNotifyReviewers() bool`

GetNotifyReviewers returns the NotifyReviewers field if non-nil, zero value otherwise.

### GetNotifyReviewersOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetNotifyReviewersOk() (*bool, bool)`

GetNotifyReviewersOk returns a tuple with the NotifyReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetNotifyReviewers(v bool)`

SetNotifyReviewers sets NotifyReviewers field to given value.

### HasNotifyReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasNotifyReviewers() bool`

HasNotifyReviewers returns a boolean if a field has been set.

### GetRemindersEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetRemindersEnabled() bool`

GetRemindersEnabled returns the RemindersEnabled field if non-nil, zero value otherwise.

### GetRemindersEnabledOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetRemindersEnabledOk() (*bool, bool)`

GetRemindersEnabledOk returns a tuple with the RemindersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindersEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetRemindersEnabled(v bool)`

SetRemindersEnabled sets RemindersEnabled field to given value.

### HasRemindersEnabled

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasRemindersEnabled() bool`

HasRemindersEnabled returns a boolean if a field has been set.

### GetRequestDurationInDays

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetRequestDurationInDays() int32`

GetRequestDurationInDays returns the RequestDurationInDays field if non-nil, zero value otherwise.

### GetRequestDurationInDaysOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetRequestDurationInDaysOk() (*int32, bool)`

GetRequestDurationInDaysOk returns a tuple with the RequestDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDurationInDays

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetRequestDurationInDays(v int32)`

SetRequestDurationInDays sets RequestDurationInDays field to given value.

### HasRequestDurationInDays

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasRequestDurationInDays() bool`

HasRequestDurationInDays returns a boolean if a field has been set.

### GetReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAdminConsentRequestPolicy) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphAdminConsentRequestPolicy) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphAdminConsentRequestPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


