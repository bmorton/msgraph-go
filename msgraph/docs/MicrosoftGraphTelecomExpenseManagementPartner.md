# MicrosoftGraphTelecomExpenseManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppAuthorized** | Pointer to **bool** | Whether the partner&#39;s AAD app has been authorized to access Intune. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name of the TEM partner. | [optional] 
**Enabled** | Pointer to **bool** | Whether Intune&#39;s connection to the TEM service is currently enabled or disabled. | [optional] 
**LastConnectionDateTime** | Pointer to **time.Time** | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 
**Url** | Pointer to **NullableString** | URL of the TEM partner&#39;s administrative control panel, where an administrator can configure their TEM service. | [optional] 

## Methods

### NewMicrosoftGraphTelecomExpenseManagementPartner

`func NewMicrosoftGraphTelecomExpenseManagementPartner() *MicrosoftGraphTelecomExpenseManagementPartner`

NewMicrosoftGraphTelecomExpenseManagementPartner instantiates a new MicrosoftGraphTelecomExpenseManagementPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTelecomExpenseManagementPartnerWithDefaults

`func NewMicrosoftGraphTelecomExpenseManagementPartnerWithDefaults() *MicrosoftGraphTelecomExpenseManagementPartner`

NewMicrosoftGraphTelecomExpenseManagementPartnerWithDefaults instantiates a new MicrosoftGraphTelecomExpenseManagementPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetAppAuthorized() bool`

GetAppAuthorized returns the AppAuthorized field if non-nil, zero value otherwise.

### GetAppAuthorizedOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetAppAuthorizedOk() (*bool, bool)`

GetAppAuthorizedOk returns a tuple with the AppAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetAppAuthorized(v bool)`

SetAppAuthorized sets AppAuthorized field to given value.

### HasAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasAppAuthorized() bool`

HasAppAuthorized returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTelecomExpenseManagementPartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetLastConnectionDateTimeOk() (*time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime sets LastConnectionDateTime field to given value.

### HasLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### GetUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MicrosoftGraphTelecomExpenseManagementPartner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


