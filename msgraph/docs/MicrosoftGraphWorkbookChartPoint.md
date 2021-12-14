# MicrosoftGraphWorkbookChartPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Returns the value of a chart point. Read-only. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartPointFormat**](anyOf&lt;microsoft.graph.workbookChartPointFormat&gt;.md) | Encapsulates the format properties chart point. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartPoint

`func NewMicrosoftGraphWorkbookChartPoint() *MicrosoftGraphWorkbookChartPoint`

NewMicrosoftGraphWorkbookChartPoint instantiates a new MicrosoftGraphWorkbookChartPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartPointWithDefaults

`func NewMicrosoftGraphWorkbookChartPointWithDefaults() *MicrosoftGraphWorkbookChartPoint`

NewMicrosoftGraphWorkbookChartPointWithDefaults instantiates a new MicrosoftGraphWorkbookChartPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartPoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartPoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *MicrosoftGraphWorkbookChartPoint) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetValueOk() (*AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphWorkbookChartPoint) SetValue(v AnyOfobject)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphWorkbookChartPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphWorkbookChartPoint) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphWorkbookChartPoint) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphWorkbookChartPoint) GetFormat() AnyOfmicrosoftGraphWorkbookChartPointFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartPointFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartPoint) SetFormat(v AnyOfmicrosoftGraphWorkbookChartPointFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartPoint) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartPoint) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartPoint) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


