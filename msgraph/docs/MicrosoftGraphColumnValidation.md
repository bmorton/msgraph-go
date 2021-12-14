# MicrosoftGraphColumnValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLanguage** | Pointer to **NullableString** | Default BCP 47 language tag for the description. | [optional] 
**Descriptions** | Pointer to [**[]AnyOfmicrosoftGraphDisplayNameLocalization**](AnyOfmicrosoftGraphDisplayNameLocalization.md) | Localized messages that explain what is needed for this column&#39;s value to be considered valid. User will be prompted with this message if validation fails. | [optional] 
**Formula** | Pointer to **NullableString** | The formula to validate column value. For examples, see Examples of common formulas in lists. | [optional] 

## Methods

### NewMicrosoftGraphColumnValidation

`func NewMicrosoftGraphColumnValidation() *MicrosoftGraphColumnValidation`

NewMicrosoftGraphColumnValidation instantiates a new MicrosoftGraphColumnValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphColumnValidationWithDefaults

`func NewMicrosoftGraphColumnValidationWithDefaults() *MicrosoftGraphColumnValidation`

NewMicrosoftGraphColumnValidationWithDefaults instantiates a new MicrosoftGraphColumnValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLanguage

`func (o *MicrosoftGraphColumnValidation) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *MicrosoftGraphColumnValidation) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *MicrosoftGraphColumnValidation) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *MicrosoftGraphColumnValidation) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### SetDefaultLanguageNil

`func (o *MicrosoftGraphColumnValidation) SetDefaultLanguageNil(b bool)`

 SetDefaultLanguageNil sets the value for DefaultLanguage to be an explicit nil

### UnsetDefaultLanguage
`func (o *MicrosoftGraphColumnValidation) UnsetDefaultLanguage()`

UnsetDefaultLanguage ensures that no value is present for DefaultLanguage, not even an explicit nil
### GetDescriptions

`func (o *MicrosoftGraphColumnValidation) GetDescriptions() []*AnyOfmicrosoftGraphDisplayNameLocalization`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *MicrosoftGraphColumnValidation) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphDisplayNameLocalization, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *MicrosoftGraphColumnValidation) SetDescriptions(v []*AnyOfmicrosoftGraphDisplayNameLocalization)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *MicrosoftGraphColumnValidation) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetFormula

`func (o *MicrosoftGraphColumnValidation) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *MicrosoftGraphColumnValidation) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *MicrosoftGraphColumnValidation) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *MicrosoftGraphColumnValidation) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### SetFormulaNil

`func (o *MicrosoftGraphColumnValidation) SetFormulaNil(b bool)`

 SetFormulaNil sets the value for Formula to be an explicit nil

### UnsetFormula
`func (o *MicrosoftGraphColumnValidation) UnsetFormula()`

UnsetFormula ensures that no value is present for Formula, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


