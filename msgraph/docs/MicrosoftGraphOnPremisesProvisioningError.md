# MicrosoftGraphOnPremisesProvisioningError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** | Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property. | [optional] 
**OccurredDateTime** | Pointer to **NullableTime** | The date and time at which the error occurred. | [optional] 
**PropertyCausingError** | Pointer to **NullableString** | Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress | [optional] 
**Value** | Pointer to **NullableString** | Value of the property causing the error. | [optional] 

## Methods

### NewMicrosoftGraphOnPremisesProvisioningError

`func NewMicrosoftGraphOnPremisesProvisioningError() *MicrosoftGraphOnPremisesProvisioningError`

NewMicrosoftGraphOnPremisesProvisioningError instantiates a new MicrosoftGraphOnPremisesProvisioningError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnPremisesProvisioningErrorWithDefaults

`func NewMicrosoftGraphOnPremisesProvisioningErrorWithDefaults() *MicrosoftGraphOnPremisesProvisioningError`

NewMicrosoftGraphOnPremisesProvisioningErrorWithDefaults instantiates a new MicrosoftGraphOnPremisesProvisioningError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MicrosoftGraphOnPremisesProvisioningError) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MicrosoftGraphOnPremisesProvisioningError) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetOccurredDateTime

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetOccurredDateTime() time.Time`

GetOccurredDateTime returns the OccurredDateTime field if non-nil, zero value otherwise.

### GetOccurredDateTimeOk

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetOccurredDateTimeOk() (*time.Time, bool)`

GetOccurredDateTimeOk returns a tuple with the OccurredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredDateTime

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetOccurredDateTime(v time.Time)`

SetOccurredDateTime sets OccurredDateTime field to given value.

### HasOccurredDateTime

`func (o *MicrosoftGraphOnPremisesProvisioningError) HasOccurredDateTime() bool`

HasOccurredDateTime returns a boolean if a field has been set.

### SetOccurredDateTimeNil

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetOccurredDateTimeNil(b bool)`

 SetOccurredDateTimeNil sets the value for OccurredDateTime to be an explicit nil

### UnsetOccurredDateTime
`func (o *MicrosoftGraphOnPremisesProvisioningError) UnsetOccurredDateTime()`

UnsetOccurredDateTime ensures that no value is present for OccurredDateTime, not even an explicit nil
### GetPropertyCausingError

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetPropertyCausingError() string`

GetPropertyCausingError returns the PropertyCausingError field if non-nil, zero value otherwise.

### GetPropertyCausingErrorOk

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetPropertyCausingErrorOk() (*string, bool)`

GetPropertyCausingErrorOk returns a tuple with the PropertyCausingError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyCausingError

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetPropertyCausingError(v string)`

SetPropertyCausingError sets PropertyCausingError field to given value.

### HasPropertyCausingError

`func (o *MicrosoftGraphOnPremisesProvisioningError) HasPropertyCausingError() bool`

HasPropertyCausingError returns a boolean if a field has been set.

### SetPropertyCausingErrorNil

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetPropertyCausingErrorNil(b bool)`

 SetPropertyCausingErrorNil sets the value for PropertyCausingError to be an explicit nil

### UnsetPropertyCausingError
`func (o *MicrosoftGraphOnPremisesProvisioningError) UnsetPropertyCausingError()`

UnsetPropertyCausingError ensures that no value is present for PropertyCausingError, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphOnPremisesProvisioningError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphOnPremisesProvisioningError) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphOnPremisesProvisioningError) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphOnPremisesProvisioningError) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


