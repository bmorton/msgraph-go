# MicrosoftGraphPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **NullableString** |  | [optional] 
**Number** | Pointer to **NullableString** | The phone number. | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphPhoneType**](anyOf&lt;microsoft.graph.phoneType&gt;.md) | The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio. | [optional] 

## Methods

### NewMicrosoftGraphPhone

`func NewMicrosoftGraphPhone() *MicrosoftGraphPhone`

NewMicrosoftGraphPhone instantiates a new MicrosoftGraphPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPhoneWithDefaults

`func NewMicrosoftGraphPhoneWithDefaults() *MicrosoftGraphPhone`

NewMicrosoftGraphPhoneWithDefaults instantiates a new MicrosoftGraphPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *MicrosoftGraphPhone) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphPhone) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MicrosoftGraphPhone) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MicrosoftGraphPhone) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MicrosoftGraphPhone) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MicrosoftGraphPhone) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetNumber

`func (o *MicrosoftGraphPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *MicrosoftGraphPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *MicrosoftGraphPhone) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *MicrosoftGraphPhone) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *MicrosoftGraphPhone) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *MicrosoftGraphPhone) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetRegion

`func (o *MicrosoftGraphPhone) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MicrosoftGraphPhone) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MicrosoftGraphPhone) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MicrosoftGraphPhone) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MicrosoftGraphPhone) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MicrosoftGraphPhone) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetType

`func (o *MicrosoftGraphPhone) GetType() AnyOfmicrosoftGraphPhoneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphPhone) GetTypeOk() (*AnyOfmicrosoftGraphPhoneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphPhone) SetType(v AnyOfmicrosoftGraphPhoneType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphPhone) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphPhone) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphPhone) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


