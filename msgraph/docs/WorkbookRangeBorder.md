# WorkbookRangeBorder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **NullableString** | HTML color code representing the color of the border line, of the form #RRGGBB (e.g. &#39;FFA500&#39;) or as a named HTML color (e.g. &#39;orange&#39;). | [optional] 
**SideIndex** | Pointer to **NullableString** | Constant value that indicates the specific side of the border. The possible values are: EdgeTop, EdgeBottom, EdgeLeft, EdgeRight, InsideVertical, InsideHorizontal, DiagonalDown, DiagonalUp. Read-only. | [optional] 
**Style** | Pointer to **NullableString** | One of the constants of line style specifying the line style for the border. The possible values are: None, Continuous, Dash, DashDot, DashDotDot, Dot, Double, SlantDashDot. | [optional] 
**Weight** | Pointer to **NullableString** | Specifies the weight of the border around a range. The possible values are: Hairline, Thin, Medium, Thick. | [optional] 

## Methods

### NewWorkbookRangeBorder

`func NewWorkbookRangeBorder() *WorkbookRangeBorder`

NewWorkbookRangeBorder instantiates a new WorkbookRangeBorder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookRangeBorderWithDefaults

`func NewWorkbookRangeBorderWithDefaults() *WorkbookRangeBorder`

NewWorkbookRangeBorderWithDefaults instantiates a new WorkbookRangeBorder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *WorkbookRangeBorder) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkbookRangeBorder) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkbookRangeBorder) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WorkbookRangeBorder) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *WorkbookRangeBorder) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *WorkbookRangeBorder) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetSideIndex

`func (o *WorkbookRangeBorder) GetSideIndex() string`

GetSideIndex returns the SideIndex field if non-nil, zero value otherwise.

### GetSideIndexOk

`func (o *WorkbookRangeBorder) GetSideIndexOk() (*string, bool)`

GetSideIndexOk returns a tuple with the SideIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideIndex

`func (o *WorkbookRangeBorder) SetSideIndex(v string)`

SetSideIndex sets SideIndex field to given value.

### HasSideIndex

`func (o *WorkbookRangeBorder) HasSideIndex() bool`

HasSideIndex returns a boolean if a field has been set.

### SetSideIndexNil

`func (o *WorkbookRangeBorder) SetSideIndexNil(b bool)`

 SetSideIndexNil sets the value for SideIndex to be an explicit nil

### UnsetSideIndex
`func (o *WorkbookRangeBorder) UnsetSideIndex()`

UnsetSideIndex ensures that no value is present for SideIndex, not even an explicit nil
### GetStyle

`func (o *WorkbookRangeBorder) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *WorkbookRangeBorder) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *WorkbookRangeBorder) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *WorkbookRangeBorder) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *WorkbookRangeBorder) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *WorkbookRangeBorder) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetWeight

`func (o *WorkbookRangeBorder) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WorkbookRangeBorder) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WorkbookRangeBorder) SetWeight(v string)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WorkbookRangeBorder) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *WorkbookRangeBorder) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *WorkbookRangeBorder) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


