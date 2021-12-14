# MicrosoftGraphSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Alerts** | Pointer to [**[]MicrosoftGraphAlert**](MicrosoftGraphAlert.md) | Read-only. Nullable. | [optional] 
**SecureScoreControlProfiles** | Pointer to [**[]MicrosoftGraphSecureScoreControlProfile**](MicrosoftGraphSecureScoreControlProfile.md) |  | [optional] 
**SecureScores** | Pointer to [**[]MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md) |  | [optional] 

## Methods

### NewMicrosoftGraphSecurity

`func NewMicrosoftGraphSecurity() *MicrosoftGraphSecurity`

NewMicrosoftGraphSecurity instantiates a new MicrosoftGraphSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSecurityWithDefaults

`func NewMicrosoftGraphSecurityWithDefaults() *MicrosoftGraphSecurity`

NewMicrosoftGraphSecurityWithDefaults instantiates a new MicrosoftGraphSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSecurity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSecurity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSecurity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSecurity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlerts

`func (o *MicrosoftGraphSecurity) GetAlerts() []MicrosoftGraphAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *MicrosoftGraphSecurity) GetAlertsOk() (*[]MicrosoftGraphAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *MicrosoftGraphSecurity) SetAlerts(v []MicrosoftGraphAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *MicrosoftGraphSecurity) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetSecureScoreControlProfiles

`func (o *MicrosoftGraphSecurity) GetSecureScoreControlProfiles() []MicrosoftGraphSecureScoreControlProfile`

GetSecureScoreControlProfiles returns the SecureScoreControlProfiles field if non-nil, zero value otherwise.

### GetSecureScoreControlProfilesOk

`func (o *MicrosoftGraphSecurity) GetSecureScoreControlProfilesOk() (*[]MicrosoftGraphSecureScoreControlProfile, bool)`

GetSecureScoreControlProfilesOk returns a tuple with the SecureScoreControlProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureScoreControlProfiles

`func (o *MicrosoftGraphSecurity) SetSecureScoreControlProfiles(v []MicrosoftGraphSecureScoreControlProfile)`

SetSecureScoreControlProfiles sets SecureScoreControlProfiles field to given value.

### HasSecureScoreControlProfiles

`func (o *MicrosoftGraphSecurity) HasSecureScoreControlProfiles() bool`

HasSecureScoreControlProfiles returns a boolean if a field has been set.

### GetSecureScores

`func (o *MicrosoftGraphSecurity) GetSecureScores() []MicrosoftGraphSecureScore`

GetSecureScores returns the SecureScores field if non-nil, zero value otherwise.

### GetSecureScoresOk

`func (o *MicrosoftGraphSecurity) GetSecureScoresOk() (*[]MicrosoftGraphSecureScore, bool)`

GetSecureScoresOk returns a tuple with the SecureScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureScores

`func (o *MicrosoftGraphSecurity) SetSecureScores(v []MicrosoftGraphSecureScore)`

SetSecureScores sets SecureScores field to given value.

### HasSecureScores

`func (o *MicrosoftGraphSecurity) HasSecureScores() bool`

HasSecureScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


