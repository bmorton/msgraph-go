# MicrosoftGraphDisplayNameLocalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | If present, the value of this field contains the displayName string that has been set for the language present in the languageTag field. | [optional] 
**LanguageTag** | Pointer to **NullableString** | Provides the language culture-code and friendly name of the language that the displayName field has been provided in. | [optional] 

## Methods

### NewMicrosoftGraphDisplayNameLocalization

`func NewMicrosoftGraphDisplayNameLocalization() *MicrosoftGraphDisplayNameLocalization`

NewMicrosoftGraphDisplayNameLocalization instantiates a new MicrosoftGraphDisplayNameLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDisplayNameLocalizationWithDefaults

`func NewMicrosoftGraphDisplayNameLocalizationWithDefaults() *MicrosoftGraphDisplayNameLocalization`

NewMicrosoftGraphDisplayNameLocalizationWithDefaults instantiates a new MicrosoftGraphDisplayNameLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphDisplayNameLocalization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDisplayNameLocalization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDisplayNameLocalization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDisplayNameLocalization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDisplayNameLocalization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDisplayNameLocalization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLanguageTag

`func (o *MicrosoftGraphDisplayNameLocalization) GetLanguageTag() string`

GetLanguageTag returns the LanguageTag field if non-nil, zero value otherwise.

### GetLanguageTagOk

`func (o *MicrosoftGraphDisplayNameLocalization) GetLanguageTagOk() (*string, bool)`

GetLanguageTagOk returns a tuple with the LanguageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageTag

`func (o *MicrosoftGraphDisplayNameLocalization) SetLanguageTag(v string)`

SetLanguageTag sets LanguageTag field to given value.

### HasLanguageTag

`func (o *MicrosoftGraphDisplayNameLocalization) HasLanguageTag() bool`

HasLanguageTag returns a boolean if a field has been set.

### SetLanguageTagNil

`func (o *MicrosoftGraphDisplayNameLocalization) SetLanguageTagNil(b bool)`

 SetLanguageTagNil sets the value for LanguageTag to be an explicit nil

### UnsetLanguageTag
`func (o *MicrosoftGraphDisplayNameLocalization) UnsetLanguageTag()`

UnsetLanguageTag ensures that no value is present for LanguageTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


