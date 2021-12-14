# MicrosoftGraphDelegatedPermissionClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Classification** | Pointer to [**AnyOfmicrosoftGraphPermissionClassificationType**](anyOf&lt;microsoft.graph.permissionClassificationType&gt;.md) | The classification value being given. Possible value: low. Does not support $filter. | [optional] 
**PermissionId** | Pointer to **NullableString** | The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter. | [optional] 
**PermissionName** | Pointer to **NullableString** | The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter. | [optional] 

## Methods

### NewMicrosoftGraphDelegatedPermissionClassification

`func NewMicrosoftGraphDelegatedPermissionClassification() *MicrosoftGraphDelegatedPermissionClassification`

NewMicrosoftGraphDelegatedPermissionClassification instantiates a new MicrosoftGraphDelegatedPermissionClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDelegatedPermissionClassificationWithDefaults

`func NewMicrosoftGraphDelegatedPermissionClassificationWithDefaults() *MicrosoftGraphDelegatedPermissionClassification`

NewMicrosoftGraphDelegatedPermissionClassificationWithDefaults instantiates a new MicrosoftGraphDelegatedPermissionClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDelegatedPermissionClassification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClassification

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetClassification() AnyOfmicrosoftGraphPermissionClassificationType`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetClassificationOk() (*AnyOfmicrosoftGraphPermissionClassificationType, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetClassification(v AnyOfmicrosoftGraphPermissionClassificationType)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *MicrosoftGraphDelegatedPermissionClassification) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *MicrosoftGraphDelegatedPermissionClassification) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetPermissionId

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.

### HasPermissionId

`func (o *MicrosoftGraphDelegatedPermissionClassification) HasPermissionId() bool`

HasPermissionId returns a boolean if a field has been set.

### SetPermissionIdNil

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetPermissionIdNil(b bool)`

 SetPermissionIdNil sets the value for PermissionId to be an explicit nil

### UnsetPermissionId
`func (o *MicrosoftGraphDelegatedPermissionClassification) UnsetPermissionId()`

UnsetPermissionId ensures that no value is present for PermissionId, not even an explicit nil
### GetPermissionName

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *MicrosoftGraphDelegatedPermissionClassification) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.

### HasPermissionName

`func (o *MicrosoftGraphDelegatedPermissionClassification) HasPermissionName() bool`

HasPermissionName returns a boolean if a field has been set.

### SetPermissionNameNil

`func (o *MicrosoftGraphDelegatedPermissionClassification) SetPermissionNameNil(b bool)`

 SetPermissionNameNil sets the value for PermissionName to be an explicit nil

### UnsetPermissionName
`func (o *MicrosoftGraphDelegatedPermissionClassification) UnsetPermissionName()`

UnsetPermissionName ensures that no value is present for PermissionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


