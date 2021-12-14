# MicrosoftGraphCloudAppSecuritySessionControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the session control is enabled. | [optional] 
**CloudAppSecurityType** | Pointer to [**AnyOfmicrosoftGraphCloudAppSecuritySessionControlType**](anyOf&lt;microsoft.graph.cloudAppSecuritySessionControlType&gt;.md) | Possible values are: mcasConfigured, monitorOnly, blockDownloads, unknownFutureValue. For more information, see Deploy Conditional Access App Control for featured apps. | [optional] 

## Methods

### NewMicrosoftGraphCloudAppSecuritySessionControl

`func NewMicrosoftGraphCloudAppSecuritySessionControl() *MicrosoftGraphCloudAppSecuritySessionControl`

NewMicrosoftGraphCloudAppSecuritySessionControl instantiates a new MicrosoftGraphCloudAppSecuritySessionControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCloudAppSecuritySessionControlWithDefaults

`func NewMicrosoftGraphCloudAppSecuritySessionControlWithDefaults() *MicrosoftGraphCloudAppSecuritySessionControl`

NewMicrosoftGraphCloudAppSecuritySessionControlWithDefaults instantiates a new MicrosoftGraphCloudAppSecuritySessionControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MicrosoftGraphCloudAppSecuritySessionControl) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetCloudAppSecurityType

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetCloudAppSecurityType() AnyOfmicrosoftGraphCloudAppSecuritySessionControlType`

GetCloudAppSecurityType returns the CloudAppSecurityType field if non-nil, zero value otherwise.

### GetCloudAppSecurityTypeOk

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetCloudAppSecurityTypeOk() (*AnyOfmicrosoftGraphCloudAppSecuritySessionControlType, bool)`

GetCloudAppSecurityTypeOk returns a tuple with the CloudAppSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAppSecurityType

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetCloudAppSecurityType(v AnyOfmicrosoftGraphCloudAppSecuritySessionControlType)`

SetCloudAppSecurityType sets CloudAppSecurityType field to given value.

### HasCloudAppSecurityType

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) HasCloudAppSecurityType() bool`

HasCloudAppSecurityType returns a boolean if a field has been set.

### SetCloudAppSecurityTypeNil

`func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetCloudAppSecurityTypeNil(b bool)`

 SetCloudAppSecurityTypeNil sets the value for CloudAppSecurityType to be an explicit nil

### UnsetCloudAppSecurityType
`func (o *MicrosoftGraphCloudAppSecuritySessionControl) UnsetCloudAppSecurityType()`

UnsetCloudAppSecurityType ensures that no value is present for CloudAppSecurityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


