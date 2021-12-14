# MicrosoftGraphWorkbookChartDataLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Position** | Pointer to **NullableString** | DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout. | [optional] 
**Separator** | Pointer to **NullableString** | String representing the separator used for the data labels on a chart. | [optional] 
**ShowBubbleSize** | Pointer to **NullableBool** | Boolean value representing if the data label bubble size is visible or not. | [optional] 
**ShowCategoryName** | Pointer to **NullableBool** | Boolean value representing if the data label category name is visible or not. | [optional] 
**ShowLegendKey** | Pointer to **NullableBool** | Boolean value representing if the data label legend key is visible or not. | [optional] 
**ShowPercentage** | Pointer to **NullableBool** | Boolean value representing if the data label percentage is visible or not. | [optional] 
**ShowSeriesName** | Pointer to **NullableBool** | Boolean value representing if the data label series name is visible or not. | [optional] 
**ShowValue** | Pointer to **NullableBool** | Boolean value representing if the data label value is visible or not. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartDataLabelFormat**](anyOf&lt;microsoft.graph.workbookChartDataLabelFormat&gt;.md) | Represents the format of chart data labels, which includes fill and font formatting. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartDataLabels

`func NewMicrosoftGraphWorkbookChartDataLabels() *MicrosoftGraphWorkbookChartDataLabels`

NewMicrosoftGraphWorkbookChartDataLabels instantiates a new MicrosoftGraphWorkbookChartDataLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartDataLabelsWithDefaults

`func NewMicrosoftGraphWorkbookChartDataLabelsWithDefaults() *MicrosoftGraphWorkbookChartDataLabels`

NewMicrosoftGraphWorkbookChartDataLabelsWithDefaults instantiates a new MicrosoftGraphWorkbookChartDataLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetSeparator

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetShowBubbleSize

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowBubbleSize() bool`

GetShowBubbleSize returns the ShowBubbleSize field if non-nil, zero value otherwise.

### GetShowBubbleSizeOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowBubbleSizeOk() (*bool, bool)`

GetShowBubbleSizeOk returns a tuple with the ShowBubbleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBubbleSize

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowBubbleSize(v bool)`

SetShowBubbleSize sets ShowBubbleSize field to given value.

### HasShowBubbleSize

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowBubbleSize() bool`

HasShowBubbleSize returns a boolean if a field has been set.

### SetShowBubbleSizeNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowBubbleSizeNil(b bool)`

 SetShowBubbleSizeNil sets the value for ShowBubbleSize to be an explicit nil

### UnsetShowBubbleSize
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowBubbleSize()`

UnsetShowBubbleSize ensures that no value is present for ShowBubbleSize, not even an explicit nil
### GetShowCategoryName

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowCategoryName() bool`

GetShowCategoryName returns the ShowCategoryName field if non-nil, zero value otherwise.

### GetShowCategoryNameOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowCategoryNameOk() (*bool, bool)`

GetShowCategoryNameOk returns a tuple with the ShowCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCategoryName

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowCategoryName(v bool)`

SetShowCategoryName sets ShowCategoryName field to given value.

### HasShowCategoryName

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowCategoryName() bool`

HasShowCategoryName returns a boolean if a field has been set.

### SetShowCategoryNameNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowCategoryNameNil(b bool)`

 SetShowCategoryNameNil sets the value for ShowCategoryName to be an explicit nil

### UnsetShowCategoryName
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowCategoryName()`

UnsetShowCategoryName ensures that no value is present for ShowCategoryName, not even an explicit nil
### GetShowLegendKey

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowLegendKey() bool`

GetShowLegendKey returns the ShowLegendKey field if non-nil, zero value otherwise.

### GetShowLegendKeyOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowLegendKeyOk() (*bool, bool)`

GetShowLegendKeyOk returns a tuple with the ShowLegendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegendKey

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowLegendKey(v bool)`

SetShowLegendKey sets ShowLegendKey field to given value.

### HasShowLegendKey

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowLegendKey() bool`

HasShowLegendKey returns a boolean if a field has been set.

### SetShowLegendKeyNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowLegendKeyNil(b bool)`

 SetShowLegendKeyNil sets the value for ShowLegendKey to be an explicit nil

### UnsetShowLegendKey
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowLegendKey()`

UnsetShowLegendKey ensures that no value is present for ShowLegendKey, not even an explicit nil
### GetShowPercentage

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowPercentage() bool`

GetShowPercentage returns the ShowPercentage field if non-nil, zero value otherwise.

### GetShowPercentageOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowPercentageOk() (*bool, bool)`

GetShowPercentageOk returns a tuple with the ShowPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPercentage

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowPercentage(v bool)`

SetShowPercentage sets ShowPercentage field to given value.

### HasShowPercentage

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowPercentage() bool`

HasShowPercentage returns a boolean if a field has been set.

### SetShowPercentageNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowPercentageNil(b bool)`

 SetShowPercentageNil sets the value for ShowPercentage to be an explicit nil

### UnsetShowPercentage
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowPercentage()`

UnsetShowPercentage ensures that no value is present for ShowPercentage, not even an explicit nil
### GetShowSeriesName

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowSeriesName() bool`

GetShowSeriesName returns the ShowSeriesName field if non-nil, zero value otherwise.

### GetShowSeriesNameOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowSeriesNameOk() (*bool, bool)`

GetShowSeriesNameOk returns a tuple with the ShowSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSeriesName

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowSeriesName(v bool)`

SetShowSeriesName sets ShowSeriesName field to given value.

### HasShowSeriesName

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowSeriesName() bool`

HasShowSeriesName returns a boolean if a field has been set.

### SetShowSeriesNameNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowSeriesNameNil(b bool)`

 SetShowSeriesNameNil sets the value for ShowSeriesName to be an explicit nil

### UnsetShowSeriesName
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowSeriesName()`

UnsetShowSeriesName ensures that no value is present for ShowSeriesName, not even an explicit nil
### GetShowValue

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowValue() bool`

GetShowValue returns the ShowValue field if non-nil, zero value otherwise.

### GetShowValueOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetShowValueOk() (*bool, bool)`

GetShowValueOk returns a tuple with the ShowValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowValue

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowValue(v bool)`

SetShowValue sets ShowValue field to given value.

### HasShowValue

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasShowValue() bool`

HasShowValue returns a boolean if a field has been set.

### SetShowValueNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetShowValueNil(b bool)`

 SetShowValueNil sets the value for ShowValue to be an explicit nil

### UnsetShowValue
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetShowValue()`

UnsetShowValue ensures that no value is present for ShowValue, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetFormat() AnyOfmicrosoftGraphWorkbookChartDataLabelFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartDataLabels) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartDataLabelFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetFormat(v AnyOfmicrosoftGraphWorkbookChartDataLabelFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartDataLabels) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartDataLabels) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartDataLabels) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


