# StsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to **[]string** | A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required. | [optional] 
**IsOrganizationDefault** | Pointer to **NullableBool** | If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false. | [optional] 
**AppliesTo** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 

## Methods

### NewStsPolicy

`func NewStsPolicy() *StsPolicy`

NewStsPolicy instantiates a new StsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsPolicyWithDefaults

`func NewStsPolicyWithDefaults() *StsPolicy`

NewStsPolicyWithDefaults instantiates a new StsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *StsPolicy) GetDefinition() []string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *StsPolicy) GetDefinitionOk() (*[]string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *StsPolicy) SetDefinition(v []string)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *StsPolicy) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetIsOrganizationDefault

`func (o *StsPolicy) GetIsOrganizationDefault() bool`

GetIsOrganizationDefault returns the IsOrganizationDefault field if non-nil, zero value otherwise.

### GetIsOrganizationDefaultOk

`func (o *StsPolicy) GetIsOrganizationDefaultOk() (*bool, bool)`

GetIsOrganizationDefaultOk returns a tuple with the IsOrganizationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationDefault

`func (o *StsPolicy) SetIsOrganizationDefault(v bool)`

SetIsOrganizationDefault sets IsOrganizationDefault field to given value.

### HasIsOrganizationDefault

`func (o *StsPolicy) HasIsOrganizationDefault() bool`

HasIsOrganizationDefault returns a boolean if a field has been set.

### SetIsOrganizationDefaultNil

`func (o *StsPolicy) SetIsOrganizationDefaultNil(b bool)`

 SetIsOrganizationDefaultNil sets the value for IsOrganizationDefault to be an explicit nil

### UnsetIsOrganizationDefault
`func (o *StsPolicy) UnsetIsOrganizationDefault()`

UnsetIsOrganizationDefault ensures that no value is present for IsOrganizationDefault, not even an explicit nil
### GetAppliesTo

`func (o *StsPolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *StsPolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *StsPolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *StsPolicy) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


