# VppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppleId** | Pointer to **NullableString** | The apple Id associated with the given Apple Volume Purchase Program Token. | [optional] 
**AutomaticallyUpdateApps** | Pointer to **bool** | Whether or not apps for the VPP token will be automatically updated. | [optional] 
**CountryOrRegion** | Pointer to **NullableString** | Whether or not apps for the VPP token will be automatically updated. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | The expiration date time of the Apple Volume Purchase Program Token. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last modification date time associated with the Apple Volume Purchase Program Token. | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token. | [optional] 
**LastSyncStatus** | Pointer to [**AnyOfmicrosoftGraphVppTokenSyncStatus**](anyOf&lt;microsoft.graph.vppTokenSyncStatus&gt;.md) | Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed. | [optional] 
**OrganizationName** | Pointer to **NullableString** | The organization associated with the Apple Volume Purchase Program Token | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphVppTokenState**](anyOf&lt;microsoft.graph.vppTokenState&gt;.md) | Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. | [optional] 
**Token** | Pointer to **NullableString** | The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program. | [optional] 
**VppTokenAccountType** | Pointer to [**AnyOfmicrosoftGraphVppTokenAccountType**](anyOf&lt;microsoft.graph.vppTokenAccountType&gt;.md) | The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education. | [optional] 

## Methods

### NewVppToken

`func NewVppToken() *VppToken`

NewVppToken instantiates a new VppToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVppTokenWithDefaults

`func NewVppTokenWithDefaults() *VppToken`

NewVppTokenWithDefaults instantiates a new VppToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppleId

`func (o *VppToken) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *VppToken) GetAppleIdOk() (*string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleId

`func (o *VppToken) SetAppleId(v string)`

SetAppleId sets AppleId field to given value.

### HasAppleId

`func (o *VppToken) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### SetAppleIdNil

`func (o *VppToken) SetAppleIdNil(b bool)`

 SetAppleIdNil sets the value for AppleId to be an explicit nil

### UnsetAppleId
`func (o *VppToken) UnsetAppleId()`

UnsetAppleId ensures that no value is present for AppleId, not even an explicit nil
### GetAutomaticallyUpdateApps

`func (o *VppToken) GetAutomaticallyUpdateApps() bool`

GetAutomaticallyUpdateApps returns the AutomaticallyUpdateApps field if non-nil, zero value otherwise.

### GetAutomaticallyUpdateAppsOk

`func (o *VppToken) GetAutomaticallyUpdateAppsOk() (*bool, bool)`

GetAutomaticallyUpdateAppsOk returns a tuple with the AutomaticallyUpdateApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyUpdateApps

`func (o *VppToken) SetAutomaticallyUpdateApps(v bool)`

SetAutomaticallyUpdateApps sets AutomaticallyUpdateApps field to given value.

### HasAutomaticallyUpdateApps

`func (o *VppToken) HasAutomaticallyUpdateApps() bool`

HasAutomaticallyUpdateApps returns a boolean if a field has been set.

### GetCountryOrRegion

`func (o *VppToken) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *VppToken) GetCountryOrRegionOk() (*string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOrRegion

`func (o *VppToken) SetCountryOrRegion(v string)`

SetCountryOrRegion sets CountryOrRegion field to given value.

### HasCountryOrRegion

`func (o *VppToken) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegionNil

`func (o *VppToken) SetCountryOrRegionNil(b bool)`

 SetCountryOrRegionNil sets the value for CountryOrRegion to be an explicit nil

### UnsetCountryOrRegion
`func (o *VppToken) UnsetCountryOrRegion()`

UnsetCountryOrRegion ensures that no value is present for CountryOrRegion, not even an explicit nil
### GetExpirationDateTime

`func (o *VppToken) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *VppToken) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *VppToken) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *VppToken) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *VppToken) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *VppToken) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *VppToken) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *VppToken) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetLastSyncDateTime

`func (o *VppToken) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *VppToken) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *VppToken) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *VppToken) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetLastSyncStatus

`func (o *VppToken) GetLastSyncStatus() AnyOfmicrosoftGraphVppTokenSyncStatus`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *VppToken) GetLastSyncStatusOk() (*AnyOfmicrosoftGraphVppTokenSyncStatus, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *VppToken) SetLastSyncStatus(v AnyOfmicrosoftGraphVppTokenSyncStatus)`

SetLastSyncStatus sets LastSyncStatus field to given value.

### HasLastSyncStatus

`func (o *VppToken) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### SetLastSyncStatusNil

`func (o *VppToken) SetLastSyncStatusNil(b bool)`

 SetLastSyncStatusNil sets the value for LastSyncStatus to be an explicit nil

### UnsetLastSyncStatus
`func (o *VppToken) UnsetLastSyncStatus()`

UnsetLastSyncStatus ensures that no value is present for LastSyncStatus, not even an explicit nil
### GetOrganizationName

`func (o *VppToken) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *VppToken) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *VppToken) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *VppToken) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *VppToken) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *VppToken) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetState

`func (o *VppToken) GetState() AnyOfmicrosoftGraphVppTokenState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VppToken) GetStateOk() (*AnyOfmicrosoftGraphVppTokenState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VppToken) SetState(v AnyOfmicrosoftGraphVppTokenState)`

SetState sets State field to given value.

### HasState

`func (o *VppToken) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *VppToken) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *VppToken) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetToken

`func (o *VppToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VppToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VppToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VppToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *VppToken) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *VppToken) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetVppTokenAccountType

`func (o *VppToken) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType`

GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.

### GetVppTokenAccountTypeOk

`func (o *VppToken) GetVppTokenAccountTypeOk() (*AnyOfmicrosoftGraphVppTokenAccountType, bool)`

GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTokenAccountType

`func (o *VppToken) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType)`

SetVppTokenAccountType sets VppTokenAccountType field to given value.

### HasVppTokenAccountType

`func (o *VppToken) HasVppTokenAccountType() bool`

HasVppTokenAccountType returns a boolean if a field has been set.

### SetVppTokenAccountTypeNil

`func (o *VppToken) SetVppTokenAccountTypeNil(b bool)`

 SetVppTokenAccountTypeNil sets the value for VppTokenAccountType to be an explicit nil

### UnsetVppTokenAccountType
`func (o *VppToken) UnsetVppTokenAccountType()`

UnsetVppTokenAccountType ensures that no value is present for VppTokenAccountType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


