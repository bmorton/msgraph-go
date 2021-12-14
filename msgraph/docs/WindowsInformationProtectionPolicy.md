# WindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysWithoutContactBeforeUnenroll** | Pointer to **int32** | Offline interval before app data is wiped (days) | [optional] 
**MdmEnrollmentUrl** | Pointer to **NullableString** | Enrollment url for the MDM | [optional] 
**MinutesOfInactivityBeforeDeviceLock** | Pointer to **int32** | Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 &lt;&#x3D; X &lt;&#x3D; 999. | [optional] 
**NumberOfPastPinsRemembered** | Pointer to **int32** | Integer value that specifies the number of past PINs that can be associated to a user account that can&#39;t be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PasswordMaximumAttemptCount** | Pointer to **int32** | The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 &lt;&#x3D; X &lt;&#x3D; 16 for desktop and 0 &lt;&#x3D; X &lt;&#x3D; 999 for mobile devices. | [optional] 
**PinExpirationDays** | Pointer to **int32** | Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user&#39;s PIN will never expire. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PinLowercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow. | [optional] 
**PinMinimumLength** | Pointer to **int32** | Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest. | [optional] 
**PinSpecialCharacters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! &#39; # $ % &amp; &#39; ( )  + , - . / : ; &lt; &#x3D; &gt; ? @ [ / ] ^  &#x60; { | [optional] 
**PinUppercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow. | [optional] 
**RevokeOnMdmHandoffDisabled** | Pointer to **bool** | New property in RS2, pending documentation | [optional] 
**WindowsHelloForBusinessBlocked** | Pointer to **bool** | Boolean value that sets Windows Hello for Business as a method for signing into Windows. | [optional] 

## Methods

### NewWindowsInformationProtectionPolicy

`func NewWindowsInformationProtectionPolicy() *WindowsInformationProtectionPolicy`

NewWindowsInformationProtectionPolicy instantiates a new WindowsInformationProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsInformationProtectionPolicyWithDefaults

`func NewWindowsInformationProtectionPolicyWithDefaults() *WindowsInformationProtectionPolicy`

NewWindowsInformationProtectionPolicyWithDefaults instantiates a new WindowsInformationProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll() int32`

GetDaysWithoutContactBeforeUnenroll returns the DaysWithoutContactBeforeUnenroll field if non-nil, zero value otherwise.

### GetDaysWithoutContactBeforeUnenrollOk

`func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenrollOk() (*int32, bool)`

GetDaysWithoutContactBeforeUnenrollOk returns a tuple with the DaysWithoutContactBeforeUnenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(v int32)`

SetDaysWithoutContactBeforeUnenroll sets DaysWithoutContactBeforeUnenroll field to given value.

### HasDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) HasDaysWithoutContactBeforeUnenroll() bool`

HasDaysWithoutContactBeforeUnenroll returns a boolean if a field has been set.

### GetMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrl() string`

GetMdmEnrollmentUrl returns the MdmEnrollmentUrl field if non-nil, zero value otherwise.

### GetMdmEnrollmentUrlOk

`func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrlOk() (*string, bool)`

GetMdmEnrollmentUrlOk returns a tuple with the MdmEnrollmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(v string)`

SetMdmEnrollmentUrl sets MdmEnrollmentUrl field to given value.

### HasMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) HasMdmEnrollmentUrl() bool`

HasMdmEnrollmentUrl returns a boolean if a field has been set.

### SetMdmEnrollmentUrlNil

`func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrlNil(b bool)`

 SetMdmEnrollmentUrlNil sets the value for MdmEnrollmentUrl to be an explicit nil

### UnsetMdmEnrollmentUrl
`func (o *WindowsInformationProtectionPolicy) UnsetMdmEnrollmentUrl()`

UnsetMdmEnrollmentUrl ensures that no value is present for MdmEnrollmentUrl, not even an explicit nil
### GetMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock() int32`

GetMinutesOfInactivityBeforeDeviceLock returns the MinutesOfInactivityBeforeDeviceLock field if non-nil, zero value otherwise.

### GetMinutesOfInactivityBeforeDeviceLockOk

`func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLockOk() (*int32, bool)`

GetMinutesOfInactivityBeforeDeviceLockOk returns a tuple with the MinutesOfInactivityBeforeDeviceLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(v int32)`

SetMinutesOfInactivityBeforeDeviceLock sets MinutesOfInactivityBeforeDeviceLock field to given value.

### HasMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) HasMinutesOfInactivityBeforeDeviceLock() bool`

HasMinutesOfInactivityBeforeDeviceLock returns a boolean if a field has been set.

### GetNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered() int32`

GetNumberOfPastPinsRemembered returns the NumberOfPastPinsRemembered field if non-nil, zero value otherwise.

### GetNumberOfPastPinsRememberedOk

`func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRememberedOk() (*int32, bool)`

GetNumberOfPastPinsRememberedOk returns a tuple with the NumberOfPastPinsRemembered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(v int32)`

SetNumberOfPastPinsRemembered sets NumberOfPastPinsRemembered field to given value.

### HasNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) HasNumberOfPastPinsRemembered() bool`

HasNumberOfPastPinsRemembered returns a boolean if a field has been set.

### GetPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount() int32`

GetPasswordMaximumAttemptCount returns the PasswordMaximumAttemptCount field if non-nil, zero value otherwise.

### GetPasswordMaximumAttemptCountOk

`func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCountOk() (*int32, bool)`

GetPasswordMaximumAttemptCountOk returns a tuple with the PasswordMaximumAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(v int32)`

SetPasswordMaximumAttemptCount sets PasswordMaximumAttemptCount field to given value.

### HasPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) HasPasswordMaximumAttemptCount() bool`

HasPasswordMaximumAttemptCount returns a boolean if a field has been set.

### GetPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) GetPinExpirationDays() int32`

GetPinExpirationDays returns the PinExpirationDays field if non-nil, zero value otherwise.

### GetPinExpirationDaysOk

`func (o *WindowsInformationProtectionPolicy) GetPinExpirationDaysOk() (*int32, bool)`

GetPinExpirationDaysOk returns a tuple with the PinExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) SetPinExpirationDays(v int32)`

SetPinExpirationDays sets PinExpirationDays field to given value.

### HasPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) HasPinExpirationDays() bool`

HasPinExpirationDays returns a boolean if a field has been set.

### GetPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinLowercaseLetters returns the PinLowercaseLetters field if non-nil, zero value otherwise.

### GetPinLowercaseLettersOk

`func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLettersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinLowercaseLettersOk returns a tuple with the PinLowercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) SetPinLowercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinLowercaseLetters sets PinLowercaseLetters field to given value.

### HasPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) HasPinLowercaseLetters() bool`

HasPinLowercaseLetters returns a boolean if a field has been set.

### SetPinLowercaseLettersNil

`func (o *WindowsInformationProtectionPolicy) SetPinLowercaseLettersNil(b bool)`

 SetPinLowercaseLettersNil sets the value for PinLowercaseLetters to be an explicit nil

### UnsetPinLowercaseLetters
`func (o *WindowsInformationProtectionPolicy) UnsetPinLowercaseLetters()`

UnsetPinLowercaseLetters ensures that no value is present for PinLowercaseLetters, not even an explicit nil
### GetPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) GetPinMinimumLength() int32`

GetPinMinimumLength returns the PinMinimumLength field if non-nil, zero value otherwise.

### GetPinMinimumLengthOk

`func (o *WindowsInformationProtectionPolicy) GetPinMinimumLengthOk() (*int32, bool)`

GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) SetPinMinimumLength(v int32)`

SetPinMinimumLength sets PinMinimumLength field to given value.

### HasPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) HasPinMinimumLength() bool`

HasPinMinimumLength returns a boolean if a field has been set.

### GetPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharacters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinSpecialCharacters returns the PinSpecialCharacters field if non-nil, zero value otherwise.

### GetPinSpecialCharactersOk

`func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharactersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinSpecialCharactersOk returns a tuple with the PinSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) SetPinSpecialCharacters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinSpecialCharacters sets PinSpecialCharacters field to given value.

### HasPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) HasPinSpecialCharacters() bool`

HasPinSpecialCharacters returns a boolean if a field has been set.

### SetPinSpecialCharactersNil

`func (o *WindowsInformationProtectionPolicy) SetPinSpecialCharactersNil(b bool)`

 SetPinSpecialCharactersNil sets the value for PinSpecialCharacters to be an explicit nil

### UnsetPinSpecialCharacters
`func (o *WindowsInformationProtectionPolicy) UnsetPinSpecialCharacters()`

UnsetPinSpecialCharacters ensures that no value is present for PinSpecialCharacters, not even an explicit nil
### GetPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinUppercaseLetters returns the PinUppercaseLetters field if non-nil, zero value otherwise.

### GetPinUppercaseLettersOk

`func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLettersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinUppercaseLettersOk returns a tuple with the PinUppercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) SetPinUppercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinUppercaseLetters sets PinUppercaseLetters field to given value.

### HasPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) HasPinUppercaseLetters() bool`

HasPinUppercaseLetters returns a boolean if a field has been set.

### SetPinUppercaseLettersNil

`func (o *WindowsInformationProtectionPolicy) SetPinUppercaseLettersNil(b bool)`

 SetPinUppercaseLettersNil sets the value for PinUppercaseLetters to be an explicit nil

### UnsetPinUppercaseLetters
`func (o *WindowsInformationProtectionPolicy) UnsetPinUppercaseLetters()`

UnsetPinUppercaseLetters ensures that no value is present for PinUppercaseLetters, not even an explicit nil
### GetRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled() bool`

GetRevokeOnMdmHandoffDisabled returns the RevokeOnMdmHandoffDisabled field if non-nil, zero value otherwise.

### GetRevokeOnMdmHandoffDisabledOk

`func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabledOk() (*bool, bool)`

GetRevokeOnMdmHandoffDisabledOk returns a tuple with the RevokeOnMdmHandoffDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(v bool)`

SetRevokeOnMdmHandoffDisabled sets RevokeOnMdmHandoffDisabled field to given value.

### HasRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) HasRevokeOnMdmHandoffDisabled() bool`

HasRevokeOnMdmHandoffDisabled returns a boolean if a field has been set.

### GetWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked() bool`

GetWindowsHelloForBusinessBlocked returns the WindowsHelloForBusinessBlocked field if non-nil, zero value otherwise.

### GetWindowsHelloForBusinessBlockedOk

`func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlockedOk() (*bool, bool)`

GetWindowsHelloForBusinessBlockedOk returns a tuple with the WindowsHelloForBusinessBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(v bool)`

SetWindowsHelloForBusinessBlocked sets WindowsHelloForBusinessBlocked field to given value.

### HasWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) HasWindowsHelloForBusinessBlocked() bool`

HasWindowsHelloForBusinessBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


