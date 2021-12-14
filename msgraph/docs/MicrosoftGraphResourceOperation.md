# MicrosoftGraphResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActionName** | Pointer to **NullableString** | Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible. | [optional] 
**Description** | Pointer to **NullableString** | Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal. | [optional] 
**ResourceName** | Pointer to **NullableString** | Name of the Resource this operation is performed on. | [optional] 

## Methods

### NewMicrosoftGraphResourceOperation

`func NewMicrosoftGraphResourceOperation() *MicrosoftGraphResourceOperation`

NewMicrosoftGraphResourceOperation instantiates a new MicrosoftGraphResourceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphResourceOperationWithDefaults

`func NewMicrosoftGraphResourceOperationWithDefaults() *MicrosoftGraphResourceOperation`

NewMicrosoftGraphResourceOperationWithDefaults instantiates a new MicrosoftGraphResourceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphResourceOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphResourceOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphResourceOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphResourceOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionName

`func (o *MicrosoftGraphResourceOperation) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphResourceOperation) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *MicrosoftGraphResourceOperation) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *MicrosoftGraphResourceOperation) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionNameNil

`func (o *MicrosoftGraphResourceOperation) SetActionNameNil(b bool)`

 SetActionNameNil sets the value for ActionName to be an explicit nil

### UnsetActionName
`func (o *MicrosoftGraphResourceOperation) UnsetActionName()`

UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphResourceOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphResourceOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphResourceOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphResourceOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphResourceOperation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphResourceOperation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetResourceName

`func (o *MicrosoftGraphResourceOperation) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MicrosoftGraphResourceOperation) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *MicrosoftGraphResourceOperation) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *MicrosoftGraphResourceOperation) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *MicrosoftGraphResourceOperation) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *MicrosoftGraphResourceOperation) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


