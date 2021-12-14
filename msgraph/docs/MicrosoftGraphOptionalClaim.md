# MicrosoftGraphOptionalClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalPropertiesField** | Pointer to **[]string** | Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional claim specified in the name property. | [optional] 
**Essential** | Pointer to **bool** | If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for the specific task requested by the end user. The default value is false. | [optional] 
**Name** | Pointer to **string** | The name of the optional claim. | [optional] 
**Source** | Pointer to **NullableString** | The source (directory object) of the claim. There are predefined claims and user-defined claims from extension properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the value in the name property is the extension property from the user object. | [optional] 

## Methods

### NewMicrosoftGraphOptionalClaim

`func NewMicrosoftGraphOptionalClaim() *MicrosoftGraphOptionalClaim`

NewMicrosoftGraphOptionalClaim instantiates a new MicrosoftGraphOptionalClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOptionalClaimWithDefaults

`func NewMicrosoftGraphOptionalClaimWithDefaults() *MicrosoftGraphOptionalClaim`

NewMicrosoftGraphOptionalClaimWithDefaults instantiates a new MicrosoftGraphOptionalClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalPropertiesField

`func (o *MicrosoftGraphOptionalClaim) GetAdditionalPropertiesField() []*string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *MicrosoftGraphOptionalClaim) GetAdditionalPropertiesFieldOk() (*[]*string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *MicrosoftGraphOptionalClaim) SetAdditionalPropertiesField(v []*string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *MicrosoftGraphOptionalClaim) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetEssential

`func (o *MicrosoftGraphOptionalClaim) GetEssential() bool`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *MicrosoftGraphOptionalClaim) GetEssentialOk() (*bool, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *MicrosoftGraphOptionalClaim) SetEssential(v bool)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *MicrosoftGraphOptionalClaim) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphOptionalClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphOptionalClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphOptionalClaim) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphOptionalClaim) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *MicrosoftGraphOptionalClaim) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MicrosoftGraphOptionalClaim) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MicrosoftGraphOptionalClaim) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MicrosoftGraphOptionalClaim) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MicrosoftGraphOptionalClaim) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MicrosoftGraphOptionalClaim) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


