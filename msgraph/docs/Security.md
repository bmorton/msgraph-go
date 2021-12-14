# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]MicrosoftGraphAlert**](MicrosoftGraphAlert.md) | Read-only. Nullable. | [optional] 
**SecureScoreControlProfiles** | Pointer to [**[]MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md) |  | [optional] 
**SecureScores** | Pointer to [**[]MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md) |  | [optional] 

## Methods

### NewSecurity

`func NewSecurity() *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *Security) GetAlerts() []MicrosoftGraphAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Security) GetAlertsOk() (*[]MicrosoftGraphAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Security) SetAlerts(v []MicrosoftGraphAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *Security) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetSecureScoreControlProfiles

`func (o *Security) GetSecureScoreControlProfiles() []MicrosoftGraphSecureScoreControlProfile`

GetSecureScoreControlProfiles returns the SecureScoreControlProfiles field if non-nil, zero value otherwise.

### GetSecureScoreControlProfilesOk

`func (o *Security) GetSecureScoreControlProfilesOk() (*[]MicrosoftGraphSecureScoreControlProfile, bool)`

GetSecureScoreControlProfilesOk returns a tuple with the SecureScoreControlProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureScoreControlProfiles

`func (o *Security) SetSecureScoreControlProfiles(v []MicrosoftGraphSecureScoreControlProfile)`

SetSecureScoreControlProfiles sets SecureScoreControlProfiles field to given value.

### HasSecureScoreControlProfiles

`func (o *Security) HasSecureScoreControlProfiles() bool`

HasSecureScoreControlProfiles returns a boolean if a field has been set.

### GetSecureScores

`func (o *Security) GetSecureScores() []MicrosoftGraphSecureScore`

GetSecureScores returns the SecureScores field if non-nil, zero value otherwise.

### GetSecureScoresOk

`func (o *Security) GetSecureScoresOk() (*[]MicrosoftGraphSecureScore, bool)`

GetSecureScoresOk returns a tuple with the SecureScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureScores

`func (o *Security) SetSecureScores(v []MicrosoftGraphSecureScore)`

SetSecureScores sets SecureScores field to given value.

### HasSecureScores

`func (o *Security) HasSecureScores() bool`

HasSecureScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


