# WorkbookChartDataLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkbookChartDataLabels

`func NewWorkbookChartDataLabels() *WorkbookChartDataLabels`

NewWorkbookChartDataLabels instantiates a new WorkbookChartDataLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartDataLabelsWithDefaults

`func NewWorkbookChartDataLabelsWithDefaults() *WorkbookChartDataLabels`

NewWorkbookChartDataLabelsWithDefaults instantiates a new WorkbookChartDataLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *WorkbookChartDataLabels) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkbookChartDataLabels) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WorkbookChartDataLabels) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WorkbookChartDataLabels) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *WorkbookChartDataLabels) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *WorkbookChartDataLabels) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetSeparator

`func (o *WorkbookChartDataLabels) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *WorkbookChartDataLabels) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *WorkbookChartDataLabels) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *WorkbookChartDataLabels) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *WorkbookChartDataLabels) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *WorkbookChartDataLabels) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetShowBubbleSize

`func (o *WorkbookChartDataLabels) GetShowBubbleSize() bool`

GetShowBubbleSize returns the ShowBubbleSize field if non-nil, zero value otherwise.

### GetShowBubbleSizeOk

`func (o *WorkbookChartDataLabels) GetShowBubbleSizeOk() (*bool, bool)`

GetShowBubbleSizeOk returns a tuple with the ShowBubbleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBubbleSize

`func (o *WorkbookChartDataLabels) SetShowBubbleSize(v bool)`

SetShowBubbleSize sets ShowBubbleSize field to given value.

### HasShowBubbleSize

`func (o *WorkbookChartDataLabels) HasShowBubbleSize() bool`

HasShowBubbleSize returns a boolean if a field has been set.

### SetShowBubbleSizeNil

`func (o *WorkbookChartDataLabels) SetShowBubbleSizeNil(b bool)`

 SetShowBubbleSizeNil sets the value for ShowBubbleSize to be an explicit nil

### UnsetShowBubbleSize
`func (o *WorkbookChartDataLabels) UnsetShowBubbleSize()`

UnsetShowBubbleSize ensures that no value is present for ShowBubbleSize, not even an explicit nil
### GetShowCategoryName

`func (o *WorkbookChartDataLabels) GetShowCategoryName() bool`

GetShowCategoryName returns the ShowCategoryName field if non-nil, zero value otherwise.

### GetShowCategoryNameOk

`func (o *WorkbookChartDataLabels) GetShowCategoryNameOk() (*bool, bool)`

GetShowCategoryNameOk returns a tuple with the ShowCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCategoryName

`func (o *WorkbookChartDataLabels) SetShowCategoryName(v bool)`

SetShowCategoryName sets ShowCategoryName field to given value.

### HasShowCategoryName

`func (o *WorkbookChartDataLabels) HasShowCategoryName() bool`

HasShowCategoryName returns a boolean if a field has been set.

### SetShowCategoryNameNil

`func (o *WorkbookChartDataLabels) SetShowCategoryNameNil(b bool)`

 SetShowCategoryNameNil sets the value for ShowCategoryName to be an explicit nil

### UnsetShowCategoryName
`func (o *WorkbookChartDataLabels) UnsetShowCategoryName()`

UnsetShowCategoryName ensures that no value is present for ShowCategoryName, not even an explicit nil
### GetShowLegendKey

`func (o *WorkbookChartDataLabels) GetShowLegendKey() bool`

GetShowLegendKey returns the ShowLegendKey field if non-nil, zero value otherwise.

### GetShowLegendKeyOk

`func (o *WorkbookChartDataLabels) GetShowLegendKeyOk() (*bool, bool)`

GetShowLegendKeyOk returns a tuple with the ShowLegendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegendKey

`func (o *WorkbookChartDataLabels) SetShowLegendKey(v bool)`

SetShowLegendKey sets ShowLegendKey field to given value.

### HasShowLegendKey

`func (o *WorkbookChartDataLabels) HasShowLegendKey() bool`

HasShowLegendKey returns a boolean if a field has been set.

### SetShowLegendKeyNil

`func (o *WorkbookChartDataLabels) SetShowLegendKeyNil(b bool)`

 SetShowLegendKeyNil sets the value for ShowLegendKey to be an explicit nil

### UnsetShowLegendKey
`func (o *WorkbookChartDataLabels) UnsetShowLegendKey()`

UnsetShowLegendKey ensures that no value is present for ShowLegendKey, not even an explicit nil
### GetShowPercentage

`func (o *WorkbookChartDataLabels) GetShowPercentage() bool`

GetShowPercentage returns the ShowPercentage field if non-nil, zero value otherwise.

### GetShowPercentageOk

`func (o *WorkbookChartDataLabels) GetShowPercentageOk() (*bool, bool)`

GetShowPercentageOk returns a tuple with the ShowPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPercentage

`func (o *WorkbookChartDataLabels) SetShowPercentage(v bool)`

SetShowPercentage sets ShowPercentage field to given value.

### HasShowPercentage

`func (o *WorkbookChartDataLabels) HasShowPercentage() bool`

HasShowPercentage returns a boolean if a field has been set.

### SetShowPercentageNil

`func (o *WorkbookChartDataLabels) SetShowPercentageNil(b bool)`

 SetShowPercentageNil sets the value for ShowPercentage to be an explicit nil

### UnsetShowPercentage
`func (o *WorkbookChartDataLabels) UnsetShowPercentage()`

UnsetShowPercentage ensures that no value is present for ShowPercentage, not even an explicit nil
### GetShowSeriesName

`func (o *WorkbookChartDataLabels) GetShowSeriesName() bool`

GetShowSeriesName returns the ShowSeriesName field if non-nil, zero value otherwise.

### GetShowSeriesNameOk

`func (o *WorkbookChartDataLabels) GetShowSeriesNameOk() (*bool, bool)`

GetShowSeriesNameOk returns a tuple with the ShowSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSeriesName

`func (o *WorkbookChartDataLabels) SetShowSeriesName(v bool)`

SetShowSeriesName sets ShowSeriesName field to given value.

### HasShowSeriesName

`func (o *WorkbookChartDataLabels) HasShowSeriesName() bool`

HasShowSeriesName returns a boolean if a field has been set.

### SetShowSeriesNameNil

`func (o *WorkbookChartDataLabels) SetShowSeriesNameNil(b bool)`

 SetShowSeriesNameNil sets the value for ShowSeriesName to be an explicit nil

### UnsetShowSeriesName
`func (o *WorkbookChartDataLabels) UnsetShowSeriesName()`

UnsetShowSeriesName ensures that no value is present for ShowSeriesName, not even an explicit nil
### GetShowValue

`func (o *WorkbookChartDataLabels) GetShowValue() bool`

GetShowValue returns the ShowValue field if non-nil, zero value otherwise.

### GetShowValueOk

`func (o *WorkbookChartDataLabels) GetShowValueOk() (*bool, bool)`

GetShowValueOk returns a tuple with the ShowValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowValue

`func (o *WorkbookChartDataLabels) SetShowValue(v bool)`

SetShowValue sets ShowValue field to given value.

### HasShowValue

`func (o *WorkbookChartDataLabels) HasShowValue() bool`

HasShowValue returns a boolean if a field has been set.

### SetShowValueNil

`func (o *WorkbookChartDataLabels) SetShowValueNil(b bool)`

 SetShowValueNil sets the value for ShowValue to be an explicit nil

### UnsetShowValue
`func (o *WorkbookChartDataLabels) UnsetShowValue()`

UnsetShowValue ensures that no value is present for ShowValue, not even an explicit nil
### GetFormat

`func (o *WorkbookChartDataLabels) GetFormat() AnyOfmicrosoftGraphWorkbookChartDataLabelFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartDataLabels) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartDataLabelFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChartDataLabels) SetFormat(v AnyOfmicrosoftGraphWorkbookChartDataLabelFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChartDataLabels) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChartDataLabels) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChartDataLabels) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


