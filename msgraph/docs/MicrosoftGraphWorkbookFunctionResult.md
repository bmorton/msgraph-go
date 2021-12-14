# MicrosoftGraphWorkbookFunctionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphWorkbookFunctionResult

`func NewMicrosoftGraphWorkbookFunctionResult() *MicrosoftGraphWorkbookFunctionResult`

NewMicrosoftGraphWorkbookFunctionResult instantiates a new MicrosoftGraphWorkbookFunctionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookFunctionResultWithDefaults

`func NewMicrosoftGraphWorkbookFunctionResultWithDefaults() *MicrosoftGraphWorkbookFunctionResult`

NewMicrosoftGraphWorkbookFunctionResultWithDefaults instantiates a new MicrosoftGraphWorkbookFunctionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookFunctionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookFunctionResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookFunctionResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetError

`func (o *MicrosoftGraphWorkbookFunctionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphWorkbookFunctionResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphWorkbookFunctionResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphWorkbookFunctionResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphWorkbookFunctionResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphWorkbookFunctionResult) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetValueOk() (*AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphWorkbookFunctionResult) SetValue(v AnyOfobject)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphWorkbookFunctionResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphWorkbookFunctionResult) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphWorkbookFunctionResult) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


