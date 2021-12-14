# WindowsInformationProtectionAppLearningSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **NullableString** | Application Name | [optional] 
**ApplicationType** | Pointer to [**AnyOfmicrosoftGraphApplicationType**](anyOf&lt;microsoft.graph.applicationType&gt;.md) | Application Type. Possible values are: universal, desktop. | [optional] 
**DeviceCount** | Pointer to **int32** | Device Count | [optional] 

## Methods

### NewWindowsInformationProtectionAppLearningSummary

`func NewWindowsInformationProtectionAppLearningSummary() *WindowsInformationProtectionAppLearningSummary`

NewWindowsInformationProtectionAppLearningSummary instantiates a new WindowsInformationProtectionAppLearningSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsInformationProtectionAppLearningSummaryWithDefaults

`func NewWindowsInformationProtectionAppLearningSummaryWithDefaults() *WindowsInformationProtectionAppLearningSummary`

NewWindowsInformationProtectionAppLearningSummaryWithDefaults instantiates a new WindowsInformationProtectionAppLearningSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *WindowsInformationProtectionAppLearningSummary) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationType() AnyOfmicrosoftGraphApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationTypeOk() (*AnyOfmicrosoftGraphApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationType(v AnyOfmicrosoftGraphApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *WindowsInformationProtectionAppLearningSummary) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


