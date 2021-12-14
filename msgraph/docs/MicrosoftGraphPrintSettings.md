# MicrosoftGraphPrintSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentConversionEnabled** | Pointer to **bool** | Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print service will automatically convert documents into a format compatible with the printer (xps to pdf) when needed. | [optional] 

## Methods

### NewMicrosoftGraphPrintSettings

`func NewMicrosoftGraphPrintSettings() *MicrosoftGraphPrintSettings`

NewMicrosoftGraphPrintSettings instantiates a new MicrosoftGraphPrintSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintSettingsWithDefaults

`func NewMicrosoftGraphPrintSettingsWithDefaults() *MicrosoftGraphPrintSettings`

NewMicrosoftGraphPrintSettingsWithDefaults instantiates a new MicrosoftGraphPrintSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentConversionEnabled

`func (o *MicrosoftGraphPrintSettings) GetDocumentConversionEnabled() bool`

GetDocumentConversionEnabled returns the DocumentConversionEnabled field if non-nil, zero value otherwise.

### GetDocumentConversionEnabledOk

`func (o *MicrosoftGraphPrintSettings) GetDocumentConversionEnabledOk() (*bool, bool)`

GetDocumentConversionEnabledOk returns a tuple with the DocumentConversionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentConversionEnabled

`func (o *MicrosoftGraphPrintSettings) SetDocumentConversionEnabled(v bool)`

SetDocumentConversionEnabled sets DocumentConversionEnabled field to given value.

### HasDocumentConversionEnabled

`func (o *MicrosoftGraphPrintSettings) HasDocumentConversionEnabled() bool`

HasDocumentConversionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


