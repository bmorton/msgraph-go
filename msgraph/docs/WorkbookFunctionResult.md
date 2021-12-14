# WorkbookFunctionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### NewWorkbookFunctionResult

`func NewWorkbookFunctionResult() *WorkbookFunctionResult`

NewWorkbookFunctionResult instantiates a new WorkbookFunctionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookFunctionResultWithDefaults

`func NewWorkbookFunctionResultWithDefaults() *WorkbookFunctionResult`

NewWorkbookFunctionResultWithDefaults instantiates a new WorkbookFunctionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WorkbookFunctionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkbookFunctionResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkbookFunctionResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WorkbookFunctionResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WorkbookFunctionResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WorkbookFunctionResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetValue

`func (o *WorkbookFunctionResult) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkbookFunctionResult) GetValueOk() (*AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkbookFunctionResult) SetValue(v AnyOfobject)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkbookFunctionResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkbookFunctionResult) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkbookFunctionResult) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


