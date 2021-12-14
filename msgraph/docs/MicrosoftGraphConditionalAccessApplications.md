# MicrosoftGraphConditionalAccessApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeApplications** | Pointer to **[]string** | The list of application IDs explicitly excluded from the policy. | [optional] 
**IncludeApplications** | Pointer to **[]string** | The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All. | [optional] 
**IncludeAuthenticationContextClassReferences** | Pointer to **[]string** | Authentication context class references include. Supported values are c1 through c25. | [optional] 
**IncludeUserActions** | Pointer to **[]string** | User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessApplications

`func NewMicrosoftGraphConditionalAccessApplications() *MicrosoftGraphConditionalAccessApplications`

NewMicrosoftGraphConditionalAccessApplications instantiates a new MicrosoftGraphConditionalAccessApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessApplicationsWithDefaults

`func NewMicrosoftGraphConditionalAccessApplicationsWithDefaults() *MicrosoftGraphConditionalAccessApplications`

NewMicrosoftGraphConditionalAccessApplicationsWithDefaults instantiates a new MicrosoftGraphConditionalAccessApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) GetExcludeApplications() []string`

GetExcludeApplications returns the ExcludeApplications field if non-nil, zero value otherwise.

### GetExcludeApplicationsOk

`func (o *MicrosoftGraphConditionalAccessApplications) GetExcludeApplicationsOk() (*[]string, bool)`

GetExcludeApplicationsOk returns a tuple with the ExcludeApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) SetExcludeApplications(v []string)`

SetExcludeApplications sets ExcludeApplications field to given value.

### HasExcludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) HasExcludeApplications() bool`

HasExcludeApplications returns a boolean if a field has been set.

### GetIncludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeApplications() []string`

GetIncludeApplications returns the IncludeApplications field if non-nil, zero value otherwise.

### GetIncludeApplicationsOk

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeApplicationsOk() (*[]string, bool)`

GetIncludeApplicationsOk returns a tuple with the IncludeApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeApplications(v []string)`

SetIncludeApplications sets IncludeApplications field to given value.

### HasIncludeApplications

`func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeApplications() bool`

HasIncludeApplications returns a boolean if a field has been set.

### GetIncludeAuthenticationContextClassReferences

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeAuthenticationContextClassReferences() []string`

GetIncludeAuthenticationContextClassReferences returns the IncludeAuthenticationContextClassReferences field if non-nil, zero value otherwise.

### GetIncludeAuthenticationContextClassReferencesOk

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeAuthenticationContextClassReferencesOk() (*[]string, bool)`

GetIncludeAuthenticationContextClassReferencesOk returns a tuple with the IncludeAuthenticationContextClassReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthenticationContextClassReferences

`func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeAuthenticationContextClassReferences(v []string)`

SetIncludeAuthenticationContextClassReferences sets IncludeAuthenticationContextClassReferences field to given value.

### HasIncludeAuthenticationContextClassReferences

`func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeAuthenticationContextClassReferences() bool`

HasIncludeAuthenticationContextClassReferences returns a boolean if a field has been set.

### GetIncludeUserActions

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeUserActions() []string`

GetIncludeUserActions returns the IncludeUserActions field if non-nil, zero value otherwise.

### GetIncludeUserActionsOk

`func (o *MicrosoftGraphConditionalAccessApplications) GetIncludeUserActionsOk() (*[]string, bool)`

GetIncludeUserActionsOk returns a tuple with the IncludeUserActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUserActions

`func (o *MicrosoftGraphConditionalAccessApplications) SetIncludeUserActions(v []string)`

SetIncludeUserActions sets IncludeUserActions field to given value.

### HasIncludeUserActions

`func (o *MicrosoftGraphConditionalAccessApplications) HasIncludeUserActions() bool`

HasIncludeUserActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


