# IdentityUserFlowAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeDataType&gt;.md) | The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime. | [optional] 
**Description** | Pointer to **NullableString** | The description of the user flow attribute that&#39;s shown to the user at the time of sign-up. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the user flow attribute. | [optional] 
**UserFlowAttributeType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeType&gt;.md) | The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required. | [optional] 

## Methods

### NewIdentityUserFlowAttribute

`func NewIdentityUserFlowAttribute() *IdentityUserFlowAttribute`

NewIdentityUserFlowAttribute instantiates a new IdentityUserFlowAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserFlowAttributeWithDefaults

`func NewIdentityUserFlowAttributeWithDefaults() *IdentityUserFlowAttribute`

NewIdentityUserFlowAttributeWithDefaults instantiates a new IdentityUserFlowAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *IdentityUserFlowAttribute) GetDataType() AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *IdentityUserFlowAttribute) GetDataTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *IdentityUserFlowAttribute) SetDataType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *IdentityUserFlowAttribute) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *IdentityUserFlowAttribute) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *IdentityUserFlowAttribute) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetDescription

`func (o *IdentityUserFlowAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityUserFlowAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityUserFlowAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityUserFlowAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IdentityUserFlowAttribute) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IdentityUserFlowAttribute) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *IdentityUserFlowAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityUserFlowAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityUserFlowAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityUserFlowAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityUserFlowAttribute) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityUserFlowAttribute) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetUserFlowAttributeType

`func (o *IdentityUserFlowAttribute) GetUserFlowAttributeType() AnyOfmicrosoftGraphIdentityUserFlowAttributeType`

GetUserFlowAttributeType returns the UserFlowAttributeType field if non-nil, zero value otherwise.

### GetUserFlowAttributeTypeOk

`func (o *IdentityUserFlowAttribute) GetUserFlowAttributeTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeType, bool)`

GetUserFlowAttributeTypeOk returns a tuple with the UserFlowAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowAttributeType

`func (o *IdentityUserFlowAttribute) SetUserFlowAttributeType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeType)`

SetUserFlowAttributeType sets UserFlowAttributeType field to given value.

### HasUserFlowAttributeType

`func (o *IdentityUserFlowAttribute) HasUserFlowAttributeType() bool`

HasUserFlowAttributeType returns a boolean if a field has been set.

### SetUserFlowAttributeTypeNil

`func (o *IdentityUserFlowAttribute) SetUserFlowAttributeTypeNil(b bool)`

 SetUserFlowAttributeTypeNil sets the value for UserFlowAttributeType to be an explicit nil

### UnsetUserFlowAttributeType
`func (o *IdentityUserFlowAttribute) UnsetUserFlowAttributeType()`

UnsetUserFlowAttributeType ensures that no value is present for UserFlowAttributeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


