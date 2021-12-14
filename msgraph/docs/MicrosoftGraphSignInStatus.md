# MicrosoftGraphSignInStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **NullableString** | Provides additional details on the sign-in activity | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | Provides the 5-6 digit error code that&#39;s generated during a sign-in failure. Check out the list of error codes and messages. | [optional] 
**FailureReason** | Pointer to **NullableString** | Provides the error message or the reason for failure for the corresponding sign-in activity. Check out the list of error codes and messages. | [optional] 

## Methods

### NewMicrosoftGraphSignInStatus

`func NewMicrosoftGraphSignInStatus() *MicrosoftGraphSignInStatus`

NewMicrosoftGraphSignInStatus instantiates a new MicrosoftGraphSignInStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSignInStatusWithDefaults

`func NewMicrosoftGraphSignInStatusWithDefaults() *MicrosoftGraphSignInStatus`

NewMicrosoftGraphSignInStatusWithDefaults instantiates a new MicrosoftGraphSignInStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) GetAdditionalDetails() string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MicrosoftGraphSignInStatus) GetAdditionalDetailsOk() (*string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) SetAdditionalDetails(v string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetailsNil

`func (o *MicrosoftGraphSignInStatus) SetAdditionalDetailsNil(b bool)`

 SetAdditionalDetailsNil sets the value for AdditionalDetails to be an explicit nil

### UnsetAdditionalDetails
`func (o *MicrosoftGraphSignInStatus) UnsetAdditionalDetails()`

UnsetAdditionalDetails ensures that no value is present for AdditionalDetails, not even an explicit nil
### GetErrorCode

`func (o *MicrosoftGraphSignInStatus) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphSignInStatus) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MicrosoftGraphSignInStatus) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MicrosoftGraphSignInStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *MicrosoftGraphSignInStatus) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *MicrosoftGraphSignInStatus) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetFailureReason

`func (o *MicrosoftGraphSignInStatus) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MicrosoftGraphSignInStatus) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MicrosoftGraphSignInStatus) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MicrosoftGraphSignInStatus) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *MicrosoftGraphSignInStatus) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *MicrosoftGraphSignInStatus) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


