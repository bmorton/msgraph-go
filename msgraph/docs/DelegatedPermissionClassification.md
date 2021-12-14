# DelegatedPermissionClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to [**AnyOfmicrosoftGraphPermissionClassificationType**](anyOf&lt;microsoft.graph.permissionClassificationType&gt;.md) | The classification value being given. Possible value: low. Does not support $filter. | [optional] 
**PermissionId** | Pointer to **NullableString** | The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter. | [optional] 
**PermissionName** | Pointer to **NullableString** | The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter. | [optional] 

## Methods

### NewDelegatedPermissionClassification

`func NewDelegatedPermissionClassification() *DelegatedPermissionClassification`

NewDelegatedPermissionClassification instantiates a new DelegatedPermissionClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedPermissionClassificationWithDefaults

`func NewDelegatedPermissionClassificationWithDefaults() *DelegatedPermissionClassification`

NewDelegatedPermissionClassificationWithDefaults instantiates a new DelegatedPermissionClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *DelegatedPermissionClassification) GetClassification() AnyOfmicrosoftGraphPermissionClassificationType`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *DelegatedPermissionClassification) GetClassificationOk() (*AnyOfmicrosoftGraphPermissionClassificationType, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *DelegatedPermissionClassification) SetClassification(v AnyOfmicrosoftGraphPermissionClassificationType)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *DelegatedPermissionClassification) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *DelegatedPermissionClassification) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *DelegatedPermissionClassification) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetPermissionId

`func (o *DelegatedPermissionClassification) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *DelegatedPermissionClassification) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *DelegatedPermissionClassification) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.

### HasPermissionId

`func (o *DelegatedPermissionClassification) HasPermissionId() bool`

HasPermissionId returns a boolean if a field has been set.

### SetPermissionIdNil

`func (o *DelegatedPermissionClassification) SetPermissionIdNil(b bool)`

 SetPermissionIdNil sets the value for PermissionId to be an explicit nil

### UnsetPermissionId
`func (o *DelegatedPermissionClassification) UnsetPermissionId()`

UnsetPermissionId ensures that no value is present for PermissionId, not even an explicit nil
### GetPermissionName

`func (o *DelegatedPermissionClassification) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *DelegatedPermissionClassification) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *DelegatedPermissionClassification) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.

### HasPermissionName

`func (o *DelegatedPermissionClassification) HasPermissionName() bool`

HasPermissionName returns a boolean if a field has been set.

### SetPermissionNameNil

`func (o *DelegatedPermissionClassification) SetPermissionNameNil(b bool)`

 SetPermissionNameNil sets the value for PermissionName to be an explicit nil

### UnsetPermissionName
`func (o *DelegatedPermissionClassification) UnsetPermissionName()`

UnsetPermissionName ensures that no value is present for PermissionName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


