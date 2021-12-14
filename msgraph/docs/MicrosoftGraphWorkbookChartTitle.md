# MicrosoftGraphWorkbookChartTitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Overlay** | Pointer to **NullableBool** | Boolean value representing if the chart title will overlay the chart or not. | [optional] 
**Text** | Pointer to **NullableString** | Represents the title text of a chart. | [optional] 
**Visible** | Pointer to **bool** | A boolean value the represents the visibility of a chart title object. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartTitleFormat**](anyOf&lt;microsoft.graph.workbookChartTitleFormat&gt;.md) | Represents the formatting of a chart title, which includes fill and font formatting. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartTitle

`func NewMicrosoftGraphWorkbookChartTitle() *MicrosoftGraphWorkbookChartTitle`

NewMicrosoftGraphWorkbookChartTitle instantiates a new MicrosoftGraphWorkbookChartTitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartTitleWithDefaults

`func NewMicrosoftGraphWorkbookChartTitleWithDefaults() *MicrosoftGraphWorkbookChartTitle`

NewMicrosoftGraphWorkbookChartTitleWithDefaults instantiates a new MicrosoftGraphWorkbookChartTitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartTitle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartTitle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartTitle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartTitle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOverlay

`func (o *MicrosoftGraphWorkbookChartTitle) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *MicrosoftGraphWorkbookChartTitle) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *MicrosoftGraphWorkbookChartTitle) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *MicrosoftGraphWorkbookChartTitle) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### SetOverlayNil

`func (o *MicrosoftGraphWorkbookChartTitle) SetOverlayNil(b bool)`

 SetOverlayNil sets the value for Overlay to be an explicit nil

### UnsetOverlay
`func (o *MicrosoftGraphWorkbookChartTitle) UnsetOverlay()`

UnsetOverlay ensures that no value is present for Overlay, not even an explicit nil
### GetText

`func (o *MicrosoftGraphWorkbookChartTitle) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphWorkbookChartTitle) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MicrosoftGraphWorkbookChartTitle) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MicrosoftGraphWorkbookChartTitle) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MicrosoftGraphWorkbookChartTitle) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MicrosoftGraphWorkbookChartTitle) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetVisible

`func (o *MicrosoftGraphWorkbookChartTitle) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MicrosoftGraphWorkbookChartTitle) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MicrosoftGraphWorkbookChartTitle) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MicrosoftGraphWorkbookChartTitle) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFormat

`func (o *MicrosoftGraphWorkbookChartTitle) GetFormat() AnyOfmicrosoftGraphWorkbookChartTitleFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartTitle) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartTitleFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartTitle) SetFormat(v AnyOfmicrosoftGraphWorkbookChartTitleFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartTitle) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartTitle) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartTitle) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


