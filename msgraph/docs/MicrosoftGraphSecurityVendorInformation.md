# MicrosoftGraphSecurityVendorInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **NullableString** | Specific provider (product/service - not vendor company); for example, WindowsDefenderATP. | [optional] 
**ProviderVersion** | Pointer to **NullableString** | Version of the provider or subprovider, if it exists, that generated the alert. Required | [optional] 
**SubProvider** | Pointer to **NullableString** | Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen. | [optional] 
**Vendor** | Pointer to **NullableString** | Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required | [optional] 

## Methods

### NewMicrosoftGraphSecurityVendorInformation

`func NewMicrosoftGraphSecurityVendorInformation() *MicrosoftGraphSecurityVendorInformation`

NewMicrosoftGraphSecurityVendorInformation instantiates a new MicrosoftGraphSecurityVendorInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSecurityVendorInformationWithDefaults

`func NewMicrosoftGraphSecurityVendorInformationWithDefaults() *MicrosoftGraphSecurityVendorInformation`

NewMicrosoftGraphSecurityVendorInformationWithDefaults instantiates a new MicrosoftGraphSecurityVendorInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *MicrosoftGraphSecurityVendorInformation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MicrosoftGraphSecurityVendorInformation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MicrosoftGraphSecurityVendorInformation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MicrosoftGraphSecurityVendorInformation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *MicrosoftGraphSecurityVendorInformation) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *MicrosoftGraphSecurityVendorInformation) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetProviderVersion

`func (o *MicrosoftGraphSecurityVendorInformation) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *MicrosoftGraphSecurityVendorInformation) GetProviderVersionOk() (*string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *MicrosoftGraphSecurityVendorInformation) SetProviderVersion(v string)`

SetProviderVersion sets ProviderVersion field to given value.

### HasProviderVersion

`func (o *MicrosoftGraphSecurityVendorInformation) HasProviderVersion() bool`

HasProviderVersion returns a boolean if a field has been set.

### SetProviderVersionNil

`func (o *MicrosoftGraphSecurityVendorInformation) SetProviderVersionNil(b bool)`

 SetProviderVersionNil sets the value for ProviderVersion to be an explicit nil

### UnsetProviderVersion
`func (o *MicrosoftGraphSecurityVendorInformation) UnsetProviderVersion()`

UnsetProviderVersion ensures that no value is present for ProviderVersion, not even an explicit nil
### GetSubProvider

`func (o *MicrosoftGraphSecurityVendorInformation) GetSubProvider() string`

GetSubProvider returns the SubProvider field if non-nil, zero value otherwise.

### GetSubProviderOk

`func (o *MicrosoftGraphSecurityVendorInformation) GetSubProviderOk() (*string, bool)`

GetSubProviderOk returns a tuple with the SubProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProvider

`func (o *MicrosoftGraphSecurityVendorInformation) SetSubProvider(v string)`

SetSubProvider sets SubProvider field to given value.

### HasSubProvider

`func (o *MicrosoftGraphSecurityVendorInformation) HasSubProvider() bool`

HasSubProvider returns a boolean if a field has been set.

### SetSubProviderNil

`func (o *MicrosoftGraphSecurityVendorInformation) SetSubProviderNil(b bool)`

 SetSubProviderNil sets the value for SubProvider to be an explicit nil

### UnsetSubProvider
`func (o *MicrosoftGraphSecurityVendorInformation) UnsetSubProvider()`

UnsetSubProvider ensures that no value is present for SubProvider, not even an explicit nil
### GetVendor

`func (o *MicrosoftGraphSecurityVendorInformation) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MicrosoftGraphSecurityVendorInformation) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MicrosoftGraphSecurityVendorInformation) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MicrosoftGraphSecurityVendorInformation) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *MicrosoftGraphSecurityVendorInformation) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *MicrosoftGraphSecurityVendorInformation) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


