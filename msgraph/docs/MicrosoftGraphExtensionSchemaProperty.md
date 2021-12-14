# MicrosoftGraphExtensionSchemaProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the strongly-typed property defined as part of a schema extension. | [optional] 
**Type** | Pointer to **NullableString** | The type of the property that is defined as part of a schema extension.  Allowed values are Binary, Boolean, DateTime, Integer or String.  See the table below for more details. | [optional] 

## Methods

### NewMicrosoftGraphExtensionSchemaProperty

`func NewMicrosoftGraphExtensionSchemaProperty() *MicrosoftGraphExtensionSchemaProperty`

NewMicrosoftGraphExtensionSchemaProperty instantiates a new MicrosoftGraphExtensionSchemaProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExtensionSchemaPropertyWithDefaults

`func NewMicrosoftGraphExtensionSchemaPropertyWithDefaults() *MicrosoftGraphExtensionSchemaProperty`

NewMicrosoftGraphExtensionSchemaPropertyWithDefaults instantiates a new MicrosoftGraphExtensionSchemaProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MicrosoftGraphExtensionSchemaProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphExtensionSchemaProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphExtensionSchemaProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphExtensionSchemaProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphExtensionSchemaProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphExtensionSchemaProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MicrosoftGraphExtensionSchemaProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExtensionSchemaProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExtensionSchemaProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExtensionSchemaProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExtensionSchemaProperty) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExtensionSchemaProperty) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


