# MicrosoftGraphPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | A string indicating the type of package. While oneNote is the only currently defined value, you should expect other package types to be returned and handle them accordingly. | [optional] 

## Methods

### NewMicrosoftGraphPackage

`func NewMicrosoftGraphPackage() *MicrosoftGraphPackage`

NewMicrosoftGraphPackage instantiates a new MicrosoftGraphPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPackageWithDefaults

`func NewMicrosoftGraphPackageWithDefaults() *MicrosoftGraphPackage`

NewMicrosoftGraphPackageWithDefaults instantiates a new MicrosoftGraphPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MicrosoftGraphPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphPackage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphPackage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


