# MicrosoftGraphSchemaExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description for the schema extension. Supports $filter (eq). | [optional] 
**Owner** | Pointer to **string** | The appId of the application that is the owner of the schema extension. This property can be supplied on creation, to set the owner.  If not supplied, then the calling application&#39;s appId will be set as the owner. In either case, the signed-in user must be the owner of the application. So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner property. Once set, this property is read-only and cannot be changed. Supports $filter (eq). | [optional] 
**Properties** | Pointer to [**[]MicrosoftGraphExtensionSchemaProperty**](MicrosoftGraphExtensionSchemaProperty.md) | The collection of property names and types that make up the schema extension definition. | [optional] 
**Status** | Pointer to **string** | The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated. Automatically set to InDevelopment on creation. Schema extensions provides more information on the possible state transitions and behaviors. Supports $filter (eq). | [optional] 
**TargetTypes** | Pointer to **[]string** | Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from contact, device, event, group, message, organization, post, or user. | [optional] 

## Methods

### NewMicrosoftGraphSchemaExtension

`func NewMicrosoftGraphSchemaExtension() *MicrosoftGraphSchemaExtension`

NewMicrosoftGraphSchemaExtension instantiates a new MicrosoftGraphSchemaExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSchemaExtensionWithDefaults

`func NewMicrosoftGraphSchemaExtensionWithDefaults() *MicrosoftGraphSchemaExtension`

NewMicrosoftGraphSchemaExtensionWithDefaults instantiates a new MicrosoftGraphSchemaExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSchemaExtension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSchemaExtension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSchemaExtension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSchemaExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphSchemaExtension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSchemaExtension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphSchemaExtension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphSchemaExtension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphSchemaExtension) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphSchemaExtension) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOwner

`func (o *MicrosoftGraphSchemaExtension) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphSchemaExtension) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphSchemaExtension) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphSchemaExtension) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProperties

`func (o *MicrosoftGraphSchemaExtension) GetProperties() []MicrosoftGraphExtensionSchemaProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MicrosoftGraphSchemaExtension) GetPropertiesOk() (*[]MicrosoftGraphExtensionSchemaProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MicrosoftGraphSchemaExtension) SetProperties(v []MicrosoftGraphExtensionSchemaProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MicrosoftGraphSchemaExtension) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphSchemaExtension) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphSchemaExtension) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphSchemaExtension) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphSchemaExtension) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetTypes

`func (o *MicrosoftGraphSchemaExtension) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *MicrosoftGraphSchemaExtension) GetTargetTypesOk() (*[]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTypes

`func (o *MicrosoftGraphSchemaExtension) SetTargetTypes(v []string)`

SetTargetTypes sets TargetTypes field to given value.

### HasTargetTypes

`func (o *MicrosoftGraphSchemaExtension) HasTargetTypes() bool`

HasTargetTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


