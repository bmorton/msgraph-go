# ResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **NullableString** | Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible. | [optional] 
**Description** | Pointer to **NullableString** | Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal. | [optional] 
**ResourceName** | Pointer to **NullableString** | Name of the Resource this operation is performed on. | [optional] 

## Methods

### NewResourceOperation

`func NewResourceOperation() *ResourceOperation`

NewResourceOperation instantiates a new ResourceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOperationWithDefaults

`func NewResourceOperationWithDefaults() *ResourceOperation`

NewResourceOperationWithDefaults instantiates a new ResourceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionName

`func (o *ResourceOperation) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *ResourceOperation) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *ResourceOperation) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *ResourceOperation) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionNameNil

`func (o *ResourceOperation) SetActionNameNil(b bool)`

 SetActionNameNil sets the value for ActionName to be an explicit nil

### UnsetActionName
`func (o *ResourceOperation) UnsetActionName()`

UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
### GetDescription

`func (o *ResourceOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourceOperation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceOperation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetResourceName

`func (o *ResourceOperation) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceOperation) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceOperation) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceOperation) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ResourceOperation) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ResourceOperation) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


