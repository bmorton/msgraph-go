# MicrosoftGraphUserSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadUserId** | Pointer to **NullableString** | AAD User object identifier (GUID) - represents the physical/multi-account user entity. | [optional] 
**AccountName** | Pointer to **NullableString** | Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName). | [optional] 
**DomainName** | Pointer to **NullableString** | NetBIOS/Active Directory domain of user account (that is, domain/account format). | [optional] 
**EmailRole** | Pointer to [**AnyOfmicrosoftGraphEmailRole**](anyOf&lt;microsoft.graph.emailRole&gt;.md) | For email-related alerts - user account&#39;s email &#39;role&#39;. Possible values are: unknown, sender, recipient. | [optional] 
**IsVpn** | Pointer to **NullableBool** | Indicates whether the user logged on through a VPN. | [optional] 
**LogonDateTime** | Pointer to **NullableTime** | Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**LogonId** | Pointer to **NullableString** | User sign-in ID. | [optional] 
**LogonIp** | Pointer to **NullableString** | IP Address the sign-in request originated from. | [optional] 
**LogonLocation** | Pointer to **NullableString** | Location (by IP address mapping) associated with a user sign-in event by this user. | [optional] 
**LogonType** | Pointer to [**AnyOfmicrosoftGraphLogonType**](anyOf&lt;microsoft.graph.logonType&gt;.md) | Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service. | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **NullableString** | Active Directory (on-premises) Security Identifier (SID) of the user. | [optional] 
**RiskScore** | Pointer to **NullableString** | Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage. | [optional] 
**UserAccountType** | Pointer to [**AnyOfmicrosoftGraphUserAccountSecurityType**](anyOf&lt;microsoft.graph.userAccountSecurityType&gt;.md) | User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | User sign-in name - internet format: (user account name)@(user account DNS domain name). | [optional] 

## Methods

### NewMicrosoftGraphUserSecurityState

`func NewMicrosoftGraphUserSecurityState() *MicrosoftGraphUserSecurityState`

NewMicrosoftGraphUserSecurityState instantiates a new MicrosoftGraphUserSecurityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserSecurityStateWithDefaults

`func NewMicrosoftGraphUserSecurityStateWithDefaults() *MicrosoftGraphUserSecurityState`

NewMicrosoftGraphUserSecurityStateWithDefaults instantiates a new MicrosoftGraphUserSecurityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadUserId

`func (o *MicrosoftGraphUserSecurityState) GetAadUserId() string`

GetAadUserId returns the AadUserId field if non-nil, zero value otherwise.

### GetAadUserIdOk

`func (o *MicrosoftGraphUserSecurityState) GetAadUserIdOk() (*string, bool)`

GetAadUserIdOk returns a tuple with the AadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadUserId

`func (o *MicrosoftGraphUserSecurityState) SetAadUserId(v string)`

SetAadUserId sets AadUserId field to given value.

### HasAadUserId

`func (o *MicrosoftGraphUserSecurityState) HasAadUserId() bool`

HasAadUserId returns a boolean if a field has been set.

### SetAadUserIdNil

`func (o *MicrosoftGraphUserSecurityState) SetAadUserIdNil(b bool)`

 SetAadUserIdNil sets the value for AadUserId to be an explicit nil

### UnsetAadUserId
`func (o *MicrosoftGraphUserSecurityState) UnsetAadUserId()`

UnsetAadUserId ensures that no value is present for AadUserId, not even an explicit nil
### GetAccountName

`func (o *MicrosoftGraphUserSecurityState) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MicrosoftGraphUserSecurityState) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *MicrosoftGraphUserSecurityState) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *MicrosoftGraphUserSecurityState) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *MicrosoftGraphUserSecurityState) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *MicrosoftGraphUserSecurityState) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetDomainName

`func (o *MicrosoftGraphUserSecurityState) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *MicrosoftGraphUserSecurityState) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *MicrosoftGraphUserSecurityState) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *MicrosoftGraphUserSecurityState) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *MicrosoftGraphUserSecurityState) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *MicrosoftGraphUserSecurityState) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetEmailRole

`func (o *MicrosoftGraphUserSecurityState) GetEmailRole() AnyOfmicrosoftGraphEmailRole`

GetEmailRole returns the EmailRole field if non-nil, zero value otherwise.

### GetEmailRoleOk

`func (o *MicrosoftGraphUserSecurityState) GetEmailRoleOk() (*AnyOfmicrosoftGraphEmailRole, bool)`

GetEmailRoleOk returns a tuple with the EmailRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRole

`func (o *MicrosoftGraphUserSecurityState) SetEmailRole(v AnyOfmicrosoftGraphEmailRole)`

SetEmailRole sets EmailRole field to given value.

### HasEmailRole

`func (o *MicrosoftGraphUserSecurityState) HasEmailRole() bool`

HasEmailRole returns a boolean if a field has been set.

### SetEmailRoleNil

`func (o *MicrosoftGraphUserSecurityState) SetEmailRoleNil(b bool)`

 SetEmailRoleNil sets the value for EmailRole to be an explicit nil

### UnsetEmailRole
`func (o *MicrosoftGraphUserSecurityState) UnsetEmailRole()`

UnsetEmailRole ensures that no value is present for EmailRole, not even an explicit nil
### GetIsVpn

`func (o *MicrosoftGraphUserSecurityState) GetIsVpn() bool`

GetIsVpn returns the IsVpn field if non-nil, zero value otherwise.

### GetIsVpnOk

`func (o *MicrosoftGraphUserSecurityState) GetIsVpnOk() (*bool, bool)`

GetIsVpnOk returns a tuple with the IsVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpn

`func (o *MicrosoftGraphUserSecurityState) SetIsVpn(v bool)`

SetIsVpn sets IsVpn field to given value.

### HasIsVpn

`func (o *MicrosoftGraphUserSecurityState) HasIsVpn() bool`

HasIsVpn returns a boolean if a field has been set.

### SetIsVpnNil

`func (o *MicrosoftGraphUserSecurityState) SetIsVpnNil(b bool)`

 SetIsVpnNil sets the value for IsVpn to be an explicit nil

### UnsetIsVpn
`func (o *MicrosoftGraphUserSecurityState) UnsetIsVpn()`

UnsetIsVpn ensures that no value is present for IsVpn, not even an explicit nil
### GetLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) GetLogonDateTime() time.Time`

GetLogonDateTime returns the LogonDateTime field if non-nil, zero value otherwise.

### GetLogonDateTimeOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonDateTimeOk() (*time.Time, bool)`

GetLogonDateTimeOk returns a tuple with the LogonDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) SetLogonDateTime(v time.Time)`

SetLogonDateTime sets LogonDateTime field to given value.

### HasLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) HasLogonDateTime() bool`

HasLogonDateTime returns a boolean if a field has been set.

### SetLogonDateTimeNil

`func (o *MicrosoftGraphUserSecurityState) SetLogonDateTimeNil(b bool)`

 SetLogonDateTimeNil sets the value for LogonDateTime to be an explicit nil

### UnsetLogonDateTime
`func (o *MicrosoftGraphUserSecurityState) UnsetLogonDateTime()`

UnsetLogonDateTime ensures that no value is present for LogonDateTime, not even an explicit nil
### GetLogonId

`func (o *MicrosoftGraphUserSecurityState) GetLogonId() string`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonIdOk() (*string, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonId

`func (o *MicrosoftGraphUserSecurityState) SetLogonId(v string)`

SetLogonId sets LogonId field to given value.

### HasLogonId

`func (o *MicrosoftGraphUserSecurityState) HasLogonId() bool`

HasLogonId returns a boolean if a field has been set.

### SetLogonIdNil

`func (o *MicrosoftGraphUserSecurityState) SetLogonIdNil(b bool)`

 SetLogonIdNil sets the value for LogonId to be an explicit nil

### UnsetLogonId
`func (o *MicrosoftGraphUserSecurityState) UnsetLogonId()`

UnsetLogonId ensures that no value is present for LogonId, not even an explicit nil
### GetLogonIp

`func (o *MicrosoftGraphUserSecurityState) GetLogonIp() string`

GetLogonIp returns the LogonIp field if non-nil, zero value otherwise.

### GetLogonIpOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonIpOk() (*string, bool)`

GetLogonIpOk returns a tuple with the LogonIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonIp

`func (o *MicrosoftGraphUserSecurityState) SetLogonIp(v string)`

SetLogonIp sets LogonIp field to given value.

### HasLogonIp

`func (o *MicrosoftGraphUserSecurityState) HasLogonIp() bool`

HasLogonIp returns a boolean if a field has been set.

### SetLogonIpNil

`func (o *MicrosoftGraphUserSecurityState) SetLogonIpNil(b bool)`

 SetLogonIpNil sets the value for LogonIp to be an explicit nil

### UnsetLogonIp
`func (o *MicrosoftGraphUserSecurityState) UnsetLogonIp()`

UnsetLogonIp ensures that no value is present for LogonIp, not even an explicit nil
### GetLogonLocation

`func (o *MicrosoftGraphUserSecurityState) GetLogonLocation() string`

GetLogonLocation returns the LogonLocation field if non-nil, zero value otherwise.

### GetLogonLocationOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonLocationOk() (*string, bool)`

GetLogonLocationOk returns a tuple with the LogonLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonLocation

`func (o *MicrosoftGraphUserSecurityState) SetLogonLocation(v string)`

SetLogonLocation sets LogonLocation field to given value.

### HasLogonLocation

`func (o *MicrosoftGraphUserSecurityState) HasLogonLocation() bool`

HasLogonLocation returns a boolean if a field has been set.

### SetLogonLocationNil

`func (o *MicrosoftGraphUserSecurityState) SetLogonLocationNil(b bool)`

 SetLogonLocationNil sets the value for LogonLocation to be an explicit nil

### UnsetLogonLocation
`func (o *MicrosoftGraphUserSecurityState) UnsetLogonLocation()`

UnsetLogonLocation ensures that no value is present for LogonLocation, not even an explicit nil
### GetLogonType

`func (o *MicrosoftGraphUserSecurityState) GetLogonType() AnyOfmicrosoftGraphLogonType`

GetLogonType returns the LogonType field if non-nil, zero value otherwise.

### GetLogonTypeOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonTypeOk() (*AnyOfmicrosoftGraphLogonType, bool)`

GetLogonTypeOk returns a tuple with the LogonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonType

`func (o *MicrosoftGraphUserSecurityState) SetLogonType(v AnyOfmicrosoftGraphLogonType)`

SetLogonType sets LogonType field to given value.

### HasLogonType

`func (o *MicrosoftGraphUserSecurityState) HasLogonType() bool`

HasLogonType returns a boolean if a field has been set.

### SetLogonTypeNil

`func (o *MicrosoftGraphUserSecurityState) SetLogonTypeNil(b bool)`

 SetLogonTypeNil sets the value for LogonType to be an explicit nil

### UnsetLogonType
`func (o *MicrosoftGraphUserSecurityState) UnsetLogonType()`

UnsetLogonType ensures that no value is present for LogonType, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphUserSecurityState) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *MicrosoftGraphUserSecurityState) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *MicrosoftGraphUserSecurityState) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetRiskScore

`func (o *MicrosoftGraphUserSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphUserSecurityState) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *MicrosoftGraphUserSecurityState) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *MicrosoftGraphUserSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScoreNil

`func (o *MicrosoftGraphUserSecurityState) SetRiskScoreNil(b bool)`

 SetRiskScoreNil sets the value for RiskScore to be an explicit nil

### UnsetRiskScore
`func (o *MicrosoftGraphUserSecurityState) UnsetRiskScore()`

UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil
### GetUserAccountType

`func (o *MicrosoftGraphUserSecurityState) GetUserAccountType() AnyOfmicrosoftGraphUserAccountSecurityType`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *MicrosoftGraphUserSecurityState) GetUserAccountTypeOk() (*AnyOfmicrosoftGraphUserAccountSecurityType, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountType

`func (o *MicrosoftGraphUserSecurityState) SetUserAccountType(v AnyOfmicrosoftGraphUserAccountSecurityType)`

SetUserAccountType sets UserAccountType field to given value.

### HasUserAccountType

`func (o *MicrosoftGraphUserSecurityState) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### SetUserAccountTypeNil

`func (o *MicrosoftGraphUserSecurityState) SetUserAccountTypeNil(b bool)`

 SetUserAccountTypeNil sets the value for UserAccountType to be an explicit nil

### UnsetUserAccountType
`func (o *MicrosoftGraphUserSecurityState) UnsetUserAccountType()`

UnsetUserAccountType ensures that no value is present for UserAccountType, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphUserSecurityState) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphUserSecurityState) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphUserSecurityState) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


