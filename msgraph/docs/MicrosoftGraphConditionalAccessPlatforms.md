# MicrosoftGraphConditionalAccessPlatforms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludePlatforms** | Pointer to [**[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform**](AnyOfmicrosoftGraphConditionalAccessDevicePlatform.md) | Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue. | [optional] 
**IncludePlatforms** | Pointer to [**[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform**](AnyOfmicrosoftGraphConditionalAccessDevicePlatform.md) | Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessPlatforms

`func NewMicrosoftGraphConditionalAccessPlatforms() *MicrosoftGraphConditionalAccessPlatforms`

NewMicrosoftGraphConditionalAccessPlatforms instantiates a new MicrosoftGraphConditionalAccessPlatforms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessPlatformsWithDefaults

`func NewMicrosoftGraphConditionalAccessPlatformsWithDefaults() *MicrosoftGraphConditionalAccessPlatforms`

NewMicrosoftGraphConditionalAccessPlatformsWithDefaults instantiates a new MicrosoftGraphConditionalAccessPlatforms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) GetExcludePlatforms() []AnyOfmicrosoftGraphConditionalAccessDevicePlatform`

GetExcludePlatforms returns the ExcludePlatforms field if non-nil, zero value otherwise.

### GetExcludePlatformsOk

`func (o *MicrosoftGraphConditionalAccessPlatforms) GetExcludePlatformsOk() (*[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform, bool)`

GetExcludePlatformsOk returns a tuple with the ExcludePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) SetExcludePlatforms(v []AnyOfmicrosoftGraphConditionalAccessDevicePlatform)`

SetExcludePlatforms sets ExcludePlatforms field to given value.

### HasExcludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) HasExcludePlatforms() bool`

HasExcludePlatforms returns a boolean if a field has been set.

### GetIncludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) GetIncludePlatforms() []AnyOfmicrosoftGraphConditionalAccessDevicePlatform`

GetIncludePlatforms returns the IncludePlatforms field if non-nil, zero value otherwise.

### GetIncludePlatformsOk

`func (o *MicrosoftGraphConditionalAccessPlatforms) GetIncludePlatformsOk() (*[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform, bool)`

GetIncludePlatformsOk returns a tuple with the IncludePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) SetIncludePlatforms(v []AnyOfmicrosoftGraphConditionalAccessDevicePlatform)`

SetIncludePlatforms sets IncludePlatforms field to given value.

### HasIncludePlatforms

`func (o *MicrosoftGraphConditionalAccessPlatforms) HasIncludePlatforms() bool`

HasIncludePlatforms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


