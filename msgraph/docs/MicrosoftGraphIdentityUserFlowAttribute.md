# MicrosoftGraphIdentityUserFlowAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DataType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeDataType&gt;.md) | The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime. | [optional] 
**Description** | Pointer to **NullableString** | The description of the user flow attribute that&#39;s shown to the user at the time of sign-up. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the user flow attribute. | [optional] 
**UserFlowAttributeType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeType&gt;.md) | The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required. | [optional] 

## Methods

### NewMicrosoftGraphIdentityUserFlowAttribute

`func NewMicrosoftGraphIdentityUserFlowAttribute() *MicrosoftGraphIdentityUserFlowAttribute`

NewMicrosoftGraphIdentityUserFlowAttribute instantiates a new MicrosoftGraphIdentityUserFlowAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityUserFlowAttributeWithDefaults

`func NewMicrosoftGraphIdentityUserFlowAttributeWithDefaults() *MicrosoftGraphIdentityUserFlowAttribute`

NewMicrosoftGraphIdentityUserFlowAttributeWithDefaults instantiates a new MicrosoftGraphIdentityUserFlowAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentityUserFlowAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDataType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDataType() AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDataTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDataType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *MicrosoftGraphIdentityUserFlowAttribute) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphIdentityUserFlowAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphIdentityUserFlowAttribute) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphIdentityUserFlowAttribute) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUserFlowAttributeType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetUserFlowAttributeType() AnyOfmicrosoftGraphIdentityUserFlowAttributeType`

GetUserFlowAttributeType returns the UserFlowAttributeType field if non-nil, zero value otherwise.

### GetUserFlowAttributeTypeOk

`func (o *MicrosoftGraphIdentityUserFlowAttribute) GetUserFlowAttributeTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeType, bool)`

GetUserFlowAttributeTypeOk returns a tuple with the UserFlowAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowAttributeType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetUserFlowAttributeType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeType)`

SetUserFlowAttributeType sets UserFlowAttributeType field to given value.

### HasUserFlowAttributeType

`func (o *MicrosoftGraphIdentityUserFlowAttribute) HasUserFlowAttributeType() bool`

HasUserFlowAttributeType returns a boolean if a field has been set.

### SetUserFlowAttributeTypeNil

`func (o *MicrosoftGraphIdentityUserFlowAttribute) SetUserFlowAttributeTypeNil(b bool)`

 SetUserFlowAttributeTypeNil sets the value for UserFlowAttributeType to be an explicit nil

### UnsetUserFlowAttributeType
`func (o *MicrosoftGraphIdentityUserFlowAttribute) UnsetUserFlowAttributeType()`

UnsetUserFlowAttributeType ensures that no value is present for UserFlowAttributeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


