# WorkbookApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationMode** | Pointer to **string** | Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual. | [optional] 

## Methods

### NewWorkbookApplication

`func NewWorkbookApplication() *WorkbookApplication`

NewWorkbookApplication instantiates a new WorkbookApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookApplicationWithDefaults

`func NewWorkbookApplicationWithDefaults() *WorkbookApplication`

NewWorkbookApplicationWithDefaults instantiates a new WorkbookApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationMode

`func (o *WorkbookApplication) GetCalculationMode() string`

GetCalculationMode returns the CalculationMode field if non-nil, zero value otherwise.

### GetCalculationModeOk

`func (o *WorkbookApplication) GetCalculationModeOk() (*string, bool)`

GetCalculationModeOk returns a tuple with the CalculationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMode

`func (o *WorkbookApplication) SetCalculationMode(v string)`

SetCalculationMode sets CalculationMode field to given value.

### HasCalculationMode

`func (o *WorkbookApplication) HasCalculationMode() bool`

HasCalculationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


