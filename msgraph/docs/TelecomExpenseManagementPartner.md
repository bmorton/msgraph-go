# TelecomExpenseManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAuthorized** | Pointer to **bool** | Whether the partner&#39;s AAD app has been authorized to access Intune. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name of the TEM partner. | [optional] 
**Enabled** | Pointer to **bool** | Whether Intune&#39;s connection to the TEM service is currently enabled or disabled. | [optional] 
**LastConnectionDateTime** | Pointer to **time.Time** | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 
**Url** | Pointer to **NullableString** | URL of the TEM partner&#39;s administrative control panel, where an administrator can configure their TEM service. | [optional] 

## Methods

### NewTelecomExpenseManagementPartner

`func NewTelecomExpenseManagementPartner() *TelecomExpenseManagementPartner`

NewTelecomExpenseManagementPartner instantiates a new TelecomExpenseManagementPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelecomExpenseManagementPartnerWithDefaults

`func NewTelecomExpenseManagementPartnerWithDefaults() *TelecomExpenseManagementPartner`

NewTelecomExpenseManagementPartnerWithDefaults instantiates a new TelecomExpenseManagementPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAuthorized

`func (o *TelecomExpenseManagementPartner) GetAppAuthorized() bool`

GetAppAuthorized returns the AppAuthorized field if non-nil, zero value otherwise.

### GetAppAuthorizedOk

`func (o *TelecomExpenseManagementPartner) GetAppAuthorizedOk() (*bool, bool)`

GetAppAuthorizedOk returns a tuple with the AppAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAuthorized

`func (o *TelecomExpenseManagementPartner) SetAppAuthorized(v bool)`

SetAppAuthorized sets AppAuthorized field to given value.

### HasAppAuthorized

`func (o *TelecomExpenseManagementPartner) HasAppAuthorized() bool`

HasAppAuthorized returns a boolean if a field has been set.

### GetDisplayName

`func (o *TelecomExpenseManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TelecomExpenseManagementPartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TelecomExpenseManagementPartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TelecomExpenseManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TelecomExpenseManagementPartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TelecomExpenseManagementPartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnabled

`func (o *TelecomExpenseManagementPartner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TelecomExpenseManagementPartner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TelecomExpenseManagementPartner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TelecomExpenseManagementPartner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTimeOk() (*time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime sets LastConnectionDateTime field to given value.

### HasLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### GetUrl

`func (o *TelecomExpenseManagementPartner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TelecomExpenseManagementPartner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TelecomExpenseManagementPartner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TelecomExpenseManagementPartner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TelecomExpenseManagementPartner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TelecomExpenseManagementPartner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


