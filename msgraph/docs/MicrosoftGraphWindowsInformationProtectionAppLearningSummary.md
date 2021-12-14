# MicrosoftGraphWindowsInformationProtectionAppLearningSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ApplicationName** | Pointer to **NullableString** | Application Name | [optional] 
**ApplicationType** | Pointer to [**AnyOfmicrosoftGraphApplicationType**](anyOf&lt;microsoft.graph.applicationType&gt;.md) | Application Type. Possible values are: universal, desktop. | [optional] 
**DeviceCount** | Pointer to **int32** | Device Count | [optional] 

## Methods

### NewMicrosoftGraphWindowsInformationProtectionAppLearningSummary

`func NewMicrosoftGraphWindowsInformationProtectionAppLearningSummary() *MicrosoftGraphWindowsInformationProtectionAppLearningSummary`

NewMicrosoftGraphWindowsInformationProtectionAppLearningSummary instantiates a new MicrosoftGraphWindowsInformationProtectionAppLearningSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWindowsInformationProtectionAppLearningSummaryWithDefaults

`func NewMicrosoftGraphWindowsInformationProtectionAppLearningSummaryWithDefaults() *MicrosoftGraphWindowsInformationProtectionAppLearningSummary`

NewMicrosoftGraphWindowsInformationProtectionAppLearningSummaryWithDefaults instantiates a new MicrosoftGraphWindowsInformationProtectionAppLearningSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationType() AnyOfmicrosoftGraphApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationTypeOk() (*AnyOfmicrosoftGraphApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationType(v AnyOfmicrosoftGraphApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


