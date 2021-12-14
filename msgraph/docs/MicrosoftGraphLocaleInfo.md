# MicrosoftGraphLocaleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | A name representing the user&#39;s locale in natural language, for example, &#39;English (United States)&#39;. | [optional] 
**Locale** | Pointer to **NullableString** | A locale representation for the user, which includes the user&#39;s preferred language and country/region. For example, &#39;en-us&#39;. The language component follows 2-letter codes as defined in ISO 639-1, and the country component follows 2-letter codes as defined in ISO 3166-1 alpha-2. | [optional] 

## Methods

### NewMicrosoftGraphLocaleInfo

`func NewMicrosoftGraphLocaleInfo() *MicrosoftGraphLocaleInfo`

NewMicrosoftGraphLocaleInfo instantiates a new MicrosoftGraphLocaleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLocaleInfoWithDefaults

`func NewMicrosoftGraphLocaleInfoWithDefaults() *MicrosoftGraphLocaleInfo`

NewMicrosoftGraphLocaleInfoWithDefaults instantiates a new MicrosoftGraphLocaleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphLocaleInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphLocaleInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphLocaleInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphLocaleInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphLocaleInfo) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphLocaleInfo) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLocale

`func (o *MicrosoftGraphLocaleInfo) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MicrosoftGraphLocaleInfo) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MicrosoftGraphLocaleInfo) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MicrosoftGraphLocaleInfo) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *MicrosoftGraphLocaleInfo) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *MicrosoftGraphLocaleInfo) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


