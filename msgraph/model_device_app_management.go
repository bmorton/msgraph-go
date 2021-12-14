/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// DeviceAppManagement Singleton entity that acts as a container for all device app management functionality.
type DeviceAppManagement struct {
	// Whether the account is enabled for syncing applications from the Microsoft Store for Business.
	IsEnabledForMicrosoftStoreForBusiness *bool `json:"isEnabledForMicrosoftStoreForBusiness,omitempty"`
	// The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
	MicrosoftStoreForBusinessLanguage NullableString `json:"microsoftStoreForBusinessLanguage,omitempty"`
	// The last time an application sync from the Microsoft Store for Business was completed.
	MicrosoftStoreForBusinessLastCompletedApplicationSyncTime *time.Time `json:"microsoftStoreForBusinessLastCompletedApplicationSyncTime,omitempty"`
	// The last time the apps from the Microsoft Store for Business were synced successfully for the account.
	MicrosoftStoreForBusinessLastSuccessfulSyncDateTime *time.Time `json:"microsoftStoreForBusinessLastSuccessfulSyncDateTime,omitempty"`
	// The Managed eBook.
	ManagedEBooks *[]MicrosoftGraphManagedEBook `json:"managedEBooks,omitempty"`
	// The mobile app categories.
	MobileAppCategories *[]MicrosoftGraphMobileAppCategory `json:"mobileAppCategories,omitempty"`
	// The Managed Device Mobile Application Configurations.
	MobileAppConfigurations *[]MicrosoftGraphManagedDeviceMobileAppConfiguration `json:"mobileAppConfigurations,omitempty"`
	// The mobile apps.
	MobileApps *[]MicrosoftGraphMobileApp `json:"mobileApps,omitempty"`
	// List of Vpp tokens for this organization.
	VppTokens *[]MicrosoftGraphVppToken `json:"vppTokens,omitempty"`
	// Android managed app policies.
	AndroidManagedAppProtections *[]MicrosoftGraphAndroidManagedAppProtection `json:"androidManagedAppProtections,omitempty"`
	// Default managed app policies.
	DefaultManagedAppProtections *[]MicrosoftGraphDefaultManagedAppProtection `json:"defaultManagedAppProtections,omitempty"`
	// iOS managed app policies.
	IosManagedAppProtections *[]MicrosoftGraphIosManagedAppProtection `json:"iosManagedAppProtections,omitempty"`
	// Managed app policies.
	ManagedAppPolicies *[]MicrosoftGraphManagedAppPolicy `json:"managedAppPolicies,omitempty"`
	// The managed app registrations.
	ManagedAppRegistrations *[]MicrosoftGraphManagedAppRegistration `json:"managedAppRegistrations,omitempty"`
	// The managed app statuses.
	ManagedAppStatuses *[]MicrosoftGraphManagedAppStatus `json:"managedAppStatuses,omitempty"`
	// Windows information protection for apps running on devices which are MDM enrolled.
	MdmWindowsInformationProtectionPolicies *[]MicrosoftGraphMdmWindowsInformationProtectionPolicy `json:"mdmWindowsInformationProtectionPolicies,omitempty"`
	// Targeted managed app configurations.
	TargetedManagedAppConfigurations *[]MicrosoftGraphTargetedManagedAppConfiguration `json:"targetedManagedAppConfigurations,omitempty"`
	// Windows information protection for apps running on devices which are not MDM enrolled.
	WindowsInformationProtectionPolicies *[]MicrosoftGraphWindowsInformationProtectionPolicy `json:"windowsInformationProtectionPolicies,omitempty"`
}

// NewDeviceAppManagement instantiates a new DeviceAppManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAppManagement() *DeviceAppManagement {
	this := DeviceAppManagement{}
	return &this
}

// NewDeviceAppManagementWithDefaults instantiates a new DeviceAppManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAppManagementWithDefaults() *DeviceAppManagement {
	this := DeviceAppManagement{}
	return &this
}

// GetIsEnabledForMicrosoftStoreForBusiness returns the IsEnabledForMicrosoftStoreForBusiness field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness() bool {
	if o == nil || o.IsEnabledForMicrosoftStoreForBusiness == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabledForMicrosoftStoreForBusiness
}

// GetIsEnabledForMicrosoftStoreForBusinessOk returns a tuple with the IsEnabledForMicrosoftStoreForBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusinessOk() (*bool, bool) {
	if o == nil || o.IsEnabledForMicrosoftStoreForBusiness == nil {
		return nil, false
	}
	return o.IsEnabledForMicrosoftStoreForBusiness, true
}

// HasIsEnabledForMicrosoftStoreForBusiness returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasIsEnabledForMicrosoftStoreForBusiness() bool {
	if o != nil && o.IsEnabledForMicrosoftStoreForBusiness != nil {
		return true
	}

	return false
}

// SetIsEnabledForMicrosoftStoreForBusiness gets a reference to the given bool and assigns it to the IsEnabledForMicrosoftStoreForBusiness field.
func (o *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(v bool) {
	o.IsEnabledForMicrosoftStoreForBusiness = &v
}

// GetMicrosoftStoreForBusinessLanguage returns the MicrosoftStoreForBusinessLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage() string {
	if o == nil || o.MicrosoftStoreForBusinessLanguage.Get() == nil {
		var ret string
		return ret
	}
	return *o.MicrosoftStoreForBusinessLanguage.Get()
}

// GetMicrosoftStoreForBusinessLanguageOk returns a tuple with the MicrosoftStoreForBusinessLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MicrosoftStoreForBusinessLanguage.Get(), o.MicrosoftStoreForBusinessLanguage.IsSet()
}

// HasMicrosoftStoreForBusinessLanguage returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMicrosoftStoreForBusinessLanguage() bool {
	if o != nil && o.MicrosoftStoreForBusinessLanguage.IsSet() {
		return true
	}

	return false
}

// SetMicrosoftStoreForBusinessLanguage gets a reference to the given NullableString and assigns it to the MicrosoftStoreForBusinessLanguage field.
func (o *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(v string) {
	o.MicrosoftStoreForBusinessLanguage.Set(&v)
}
// SetMicrosoftStoreForBusinessLanguageNil sets the value for MicrosoftStoreForBusinessLanguage to be an explicit nil
func (o *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguageNil() {
	o.MicrosoftStoreForBusinessLanguage.Set(nil)
}

// UnsetMicrosoftStoreForBusinessLanguage ensures that no value is present for MicrosoftStoreForBusinessLanguage, not even an explicit nil
func (o *DeviceAppManagement) UnsetMicrosoftStoreForBusinessLanguage() {
	o.MicrosoftStoreForBusinessLanguage.Unset()
}

// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() time.Time {
	if o == nil || o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime == nil {
		var ret time.Time
		return ret
	}
	return *o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime
}

// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk returns a tuple with the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk() (*time.Time, bool) {
	if o == nil || o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime == nil {
		return nil, false
	}
	return o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime, true
}

// HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() bool {
	if o != nil && o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime != nil {
		return true
	}

	return false
}

// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets a reference to the given time.Time and assigns it to the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field.
func (o *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(v time.Time) {
	o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime = &v
}

// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() time.Time {
	if o == nil || o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime
}

// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk returns a tuple with the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk() (*time.Time, bool) {
	if o == nil || o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime == nil {
		return nil, false
	}
	return o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime, true
}

// HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() bool {
	if o != nil && o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime != nil {
		return true
	}

	return false
}

// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets a reference to the given time.Time and assigns it to the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field.
func (o *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(v time.Time) {
	o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime = &v
}

// GetManagedEBooks returns the ManagedEBooks field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetManagedEBooks() []MicrosoftGraphManagedEBook {
	if o == nil || o.ManagedEBooks == nil {
		var ret []MicrosoftGraphManagedEBook
		return ret
	}
	return *o.ManagedEBooks
}

// GetManagedEBooksOk returns a tuple with the ManagedEBooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetManagedEBooksOk() (*[]MicrosoftGraphManagedEBook, bool) {
	if o == nil || o.ManagedEBooks == nil {
		return nil, false
	}
	return o.ManagedEBooks, true
}

// HasManagedEBooks returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasManagedEBooks() bool {
	if o != nil && o.ManagedEBooks != nil {
		return true
	}

	return false
}

// SetManagedEBooks gets a reference to the given []MicrosoftGraphManagedEBook and assigns it to the ManagedEBooks field.
func (o *DeviceAppManagement) SetManagedEBooks(v []MicrosoftGraphManagedEBook) {
	o.ManagedEBooks = &v
}

// GetMobileAppCategories returns the MobileAppCategories field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMobileAppCategories() []MicrosoftGraphMobileAppCategory {
	if o == nil || o.MobileAppCategories == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret
	}
	return *o.MobileAppCategories
}

// GetMobileAppCategoriesOk returns a tuple with the MobileAppCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMobileAppCategoriesOk() (*[]MicrosoftGraphMobileAppCategory, bool) {
	if o == nil || o.MobileAppCategories == nil {
		return nil, false
	}
	return o.MobileAppCategories, true
}

// HasMobileAppCategories returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMobileAppCategories() bool {
	if o != nil && o.MobileAppCategories != nil {
		return true
	}

	return false
}

// SetMobileAppCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the MobileAppCategories field.
func (o *DeviceAppManagement) SetMobileAppCategories(v []MicrosoftGraphMobileAppCategory) {
	o.MobileAppCategories = &v
}

// GetMobileAppConfigurations returns the MobileAppConfigurations field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMobileAppConfigurations() []MicrosoftGraphManagedDeviceMobileAppConfiguration {
	if o == nil || o.MobileAppConfigurations == nil {
		var ret []MicrosoftGraphManagedDeviceMobileAppConfiguration
		return ret
	}
	return *o.MobileAppConfigurations
}

// GetMobileAppConfigurationsOk returns a tuple with the MobileAppConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMobileAppConfigurationsOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfiguration, bool) {
	if o == nil || o.MobileAppConfigurations == nil {
		return nil, false
	}
	return o.MobileAppConfigurations, true
}

// HasMobileAppConfigurations returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMobileAppConfigurations() bool {
	if o != nil && o.MobileAppConfigurations != nil {
		return true
	}

	return false
}

// SetMobileAppConfigurations gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfiguration and assigns it to the MobileAppConfigurations field.
func (o *DeviceAppManagement) SetMobileAppConfigurations(v []MicrosoftGraphManagedDeviceMobileAppConfiguration) {
	o.MobileAppConfigurations = &v
}

// GetMobileApps returns the MobileApps field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMobileApps() []MicrosoftGraphMobileApp {
	if o == nil || o.MobileApps == nil {
		var ret []MicrosoftGraphMobileApp
		return ret
	}
	return *o.MobileApps
}

// GetMobileAppsOk returns a tuple with the MobileApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMobileAppsOk() (*[]MicrosoftGraphMobileApp, bool) {
	if o == nil || o.MobileApps == nil {
		return nil, false
	}
	return o.MobileApps, true
}

// HasMobileApps returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMobileApps() bool {
	if o != nil && o.MobileApps != nil {
		return true
	}

	return false
}

// SetMobileApps gets a reference to the given []MicrosoftGraphMobileApp and assigns it to the MobileApps field.
func (o *DeviceAppManagement) SetMobileApps(v []MicrosoftGraphMobileApp) {
	o.MobileApps = &v
}

// GetVppTokens returns the VppTokens field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetVppTokens() []MicrosoftGraphVppToken {
	if o == nil || o.VppTokens == nil {
		var ret []MicrosoftGraphVppToken
		return ret
	}
	return *o.VppTokens
}

// GetVppTokensOk returns a tuple with the VppTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetVppTokensOk() (*[]MicrosoftGraphVppToken, bool) {
	if o == nil || o.VppTokens == nil {
		return nil, false
	}
	return o.VppTokens, true
}

// HasVppTokens returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasVppTokens() bool {
	if o != nil && o.VppTokens != nil {
		return true
	}

	return false
}

// SetVppTokens gets a reference to the given []MicrosoftGraphVppToken and assigns it to the VppTokens field.
func (o *DeviceAppManagement) SetVppTokens(v []MicrosoftGraphVppToken) {
	o.VppTokens = &v
}

// GetAndroidManagedAppProtections returns the AndroidManagedAppProtections field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetAndroidManagedAppProtections() []MicrosoftGraphAndroidManagedAppProtection {
	if o == nil || o.AndroidManagedAppProtections == nil {
		var ret []MicrosoftGraphAndroidManagedAppProtection
		return ret
	}
	return *o.AndroidManagedAppProtections
}

// GetAndroidManagedAppProtectionsOk returns a tuple with the AndroidManagedAppProtections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetAndroidManagedAppProtectionsOk() (*[]MicrosoftGraphAndroidManagedAppProtection, bool) {
	if o == nil || o.AndroidManagedAppProtections == nil {
		return nil, false
	}
	return o.AndroidManagedAppProtections, true
}

// HasAndroidManagedAppProtections returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasAndroidManagedAppProtections() bool {
	if o != nil && o.AndroidManagedAppProtections != nil {
		return true
	}

	return false
}

// SetAndroidManagedAppProtections gets a reference to the given []MicrosoftGraphAndroidManagedAppProtection and assigns it to the AndroidManagedAppProtections field.
func (o *DeviceAppManagement) SetAndroidManagedAppProtections(v []MicrosoftGraphAndroidManagedAppProtection) {
	o.AndroidManagedAppProtections = &v
}

// GetDefaultManagedAppProtections returns the DefaultManagedAppProtections field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetDefaultManagedAppProtections() []MicrosoftGraphDefaultManagedAppProtection {
	if o == nil || o.DefaultManagedAppProtections == nil {
		var ret []MicrosoftGraphDefaultManagedAppProtection
		return ret
	}
	return *o.DefaultManagedAppProtections
}

// GetDefaultManagedAppProtectionsOk returns a tuple with the DefaultManagedAppProtections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetDefaultManagedAppProtectionsOk() (*[]MicrosoftGraphDefaultManagedAppProtection, bool) {
	if o == nil || o.DefaultManagedAppProtections == nil {
		return nil, false
	}
	return o.DefaultManagedAppProtections, true
}

// HasDefaultManagedAppProtections returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasDefaultManagedAppProtections() bool {
	if o != nil && o.DefaultManagedAppProtections != nil {
		return true
	}

	return false
}

// SetDefaultManagedAppProtections gets a reference to the given []MicrosoftGraphDefaultManagedAppProtection and assigns it to the DefaultManagedAppProtections field.
func (o *DeviceAppManagement) SetDefaultManagedAppProtections(v []MicrosoftGraphDefaultManagedAppProtection) {
	o.DefaultManagedAppProtections = &v
}

// GetIosManagedAppProtections returns the IosManagedAppProtections field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetIosManagedAppProtections() []MicrosoftGraphIosManagedAppProtection {
	if o == nil || o.IosManagedAppProtections == nil {
		var ret []MicrosoftGraphIosManagedAppProtection
		return ret
	}
	return *o.IosManagedAppProtections
}

// GetIosManagedAppProtectionsOk returns a tuple with the IosManagedAppProtections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetIosManagedAppProtectionsOk() (*[]MicrosoftGraphIosManagedAppProtection, bool) {
	if o == nil || o.IosManagedAppProtections == nil {
		return nil, false
	}
	return o.IosManagedAppProtections, true
}

// HasIosManagedAppProtections returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasIosManagedAppProtections() bool {
	if o != nil && o.IosManagedAppProtections != nil {
		return true
	}

	return false
}

// SetIosManagedAppProtections gets a reference to the given []MicrosoftGraphIosManagedAppProtection and assigns it to the IosManagedAppProtections field.
func (o *DeviceAppManagement) SetIosManagedAppProtections(v []MicrosoftGraphIosManagedAppProtection) {
	o.IosManagedAppProtections = &v
}

// GetManagedAppPolicies returns the ManagedAppPolicies field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetManagedAppPolicies() []MicrosoftGraphManagedAppPolicy {
	if o == nil || o.ManagedAppPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret
	}
	return *o.ManagedAppPolicies
}

// GetManagedAppPoliciesOk returns a tuple with the ManagedAppPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetManagedAppPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool) {
	if o == nil || o.ManagedAppPolicies == nil {
		return nil, false
	}
	return o.ManagedAppPolicies, true
}

// HasManagedAppPolicies returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasManagedAppPolicies() bool {
	if o != nil && o.ManagedAppPolicies != nil {
		return true
	}

	return false
}

// SetManagedAppPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the ManagedAppPolicies field.
func (o *DeviceAppManagement) SetManagedAppPolicies(v []MicrosoftGraphManagedAppPolicy) {
	o.ManagedAppPolicies = &v
}

// GetManagedAppRegistrations returns the ManagedAppRegistrations field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration {
	if o == nil || o.ManagedAppRegistrations == nil {
		var ret []MicrosoftGraphManagedAppRegistration
		return ret
	}
	return *o.ManagedAppRegistrations
}

// GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetManagedAppRegistrationsOk() (*[]MicrosoftGraphManagedAppRegistration, bool) {
	if o == nil || o.ManagedAppRegistrations == nil {
		return nil, false
	}
	return o.ManagedAppRegistrations, true
}

// HasManagedAppRegistrations returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasManagedAppRegistrations() bool {
	if o != nil && o.ManagedAppRegistrations != nil {
		return true
	}

	return false
}

// SetManagedAppRegistrations gets a reference to the given []MicrosoftGraphManagedAppRegistration and assigns it to the ManagedAppRegistrations field.
func (o *DeviceAppManagement) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration) {
	o.ManagedAppRegistrations = &v
}

// GetManagedAppStatuses returns the ManagedAppStatuses field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetManagedAppStatuses() []MicrosoftGraphManagedAppStatus {
	if o == nil || o.ManagedAppStatuses == nil {
		var ret []MicrosoftGraphManagedAppStatus
		return ret
	}
	return *o.ManagedAppStatuses
}

// GetManagedAppStatusesOk returns a tuple with the ManagedAppStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetManagedAppStatusesOk() (*[]MicrosoftGraphManagedAppStatus, bool) {
	if o == nil || o.ManagedAppStatuses == nil {
		return nil, false
	}
	return o.ManagedAppStatuses, true
}

// HasManagedAppStatuses returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasManagedAppStatuses() bool {
	if o != nil && o.ManagedAppStatuses != nil {
		return true
	}

	return false
}

// SetManagedAppStatuses gets a reference to the given []MicrosoftGraphManagedAppStatus and assigns it to the ManagedAppStatuses field.
func (o *DeviceAppManagement) SetManagedAppStatuses(v []MicrosoftGraphManagedAppStatus) {
	o.ManagedAppStatuses = &v
}

// GetMdmWindowsInformationProtectionPolicies returns the MdmWindowsInformationProtectionPolicies field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies() []MicrosoftGraphMdmWindowsInformationProtectionPolicy {
	if o == nil || o.MdmWindowsInformationProtectionPolicies == nil {
		var ret []MicrosoftGraphMdmWindowsInformationProtectionPolicy
		return ret
	}
	return *o.MdmWindowsInformationProtectionPolicies
}

// GetMdmWindowsInformationProtectionPoliciesOk returns a tuple with the MdmWindowsInformationProtectionPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetMdmWindowsInformationProtectionPoliciesOk() (*[]MicrosoftGraphMdmWindowsInformationProtectionPolicy, bool) {
	if o == nil || o.MdmWindowsInformationProtectionPolicies == nil {
		return nil, false
	}
	return o.MdmWindowsInformationProtectionPolicies, true
}

// HasMdmWindowsInformationProtectionPolicies returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasMdmWindowsInformationProtectionPolicies() bool {
	if o != nil && o.MdmWindowsInformationProtectionPolicies != nil {
		return true
	}

	return false
}

// SetMdmWindowsInformationProtectionPolicies gets a reference to the given []MicrosoftGraphMdmWindowsInformationProtectionPolicy and assigns it to the MdmWindowsInformationProtectionPolicies field.
func (o *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(v []MicrosoftGraphMdmWindowsInformationProtectionPolicy) {
	o.MdmWindowsInformationProtectionPolicies = &v
}

// GetTargetedManagedAppConfigurations returns the TargetedManagedAppConfigurations field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetTargetedManagedAppConfigurations() []MicrosoftGraphTargetedManagedAppConfiguration {
	if o == nil || o.TargetedManagedAppConfigurations == nil {
		var ret []MicrosoftGraphTargetedManagedAppConfiguration
		return ret
	}
	return *o.TargetedManagedAppConfigurations
}

// GetTargetedManagedAppConfigurationsOk returns a tuple with the TargetedManagedAppConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetTargetedManagedAppConfigurationsOk() (*[]MicrosoftGraphTargetedManagedAppConfiguration, bool) {
	if o == nil || o.TargetedManagedAppConfigurations == nil {
		return nil, false
	}
	return o.TargetedManagedAppConfigurations, true
}

// HasTargetedManagedAppConfigurations returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasTargetedManagedAppConfigurations() bool {
	if o != nil && o.TargetedManagedAppConfigurations != nil {
		return true
	}

	return false
}

// SetTargetedManagedAppConfigurations gets a reference to the given []MicrosoftGraphTargetedManagedAppConfiguration and assigns it to the TargetedManagedAppConfigurations field.
func (o *DeviceAppManagement) SetTargetedManagedAppConfigurations(v []MicrosoftGraphTargetedManagedAppConfiguration) {
	o.TargetedManagedAppConfigurations = &v
}

// GetWindowsInformationProtectionPolicies returns the WindowsInformationProtectionPolicies field value if set, zero value otherwise.
func (o *DeviceAppManagement) GetWindowsInformationProtectionPolicies() []MicrosoftGraphWindowsInformationProtectionPolicy {
	if o == nil || o.WindowsInformationProtectionPolicies == nil {
		var ret []MicrosoftGraphWindowsInformationProtectionPolicy
		return ret
	}
	return *o.WindowsInformationProtectionPolicies
}

// GetWindowsInformationProtectionPoliciesOk returns a tuple with the WindowsInformationProtectionPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAppManagement) GetWindowsInformationProtectionPoliciesOk() (*[]MicrosoftGraphWindowsInformationProtectionPolicy, bool) {
	if o == nil || o.WindowsInformationProtectionPolicies == nil {
		return nil, false
	}
	return o.WindowsInformationProtectionPolicies, true
}

// HasWindowsInformationProtectionPolicies returns a boolean if a field has been set.
func (o *DeviceAppManagement) HasWindowsInformationProtectionPolicies() bool {
	if o != nil && o.WindowsInformationProtectionPolicies != nil {
		return true
	}

	return false
}

// SetWindowsInformationProtectionPolicies gets a reference to the given []MicrosoftGraphWindowsInformationProtectionPolicy and assigns it to the WindowsInformationProtectionPolicies field.
func (o *DeviceAppManagement) SetWindowsInformationProtectionPolicies(v []MicrosoftGraphWindowsInformationProtectionPolicy) {
	o.WindowsInformationProtectionPolicies = &v
}

func (o DeviceAppManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabledForMicrosoftStoreForBusiness != nil {
		toSerialize["isEnabledForMicrosoftStoreForBusiness"] = o.IsEnabledForMicrosoftStoreForBusiness
	}
	if o.MicrosoftStoreForBusinessLanguage.IsSet() {
		toSerialize["microsoftStoreForBusinessLanguage"] = o.MicrosoftStoreForBusinessLanguage.Get()
	}
	if o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime != nil {
		toSerialize["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = o.MicrosoftStoreForBusinessLastCompletedApplicationSyncTime
	}
	if o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime != nil {
		toSerialize["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = o.MicrosoftStoreForBusinessLastSuccessfulSyncDateTime
	}
	if o.ManagedEBooks != nil {
		toSerialize["managedEBooks"] = o.ManagedEBooks
	}
	if o.MobileAppCategories != nil {
		toSerialize["mobileAppCategories"] = o.MobileAppCategories
	}
	if o.MobileAppConfigurations != nil {
		toSerialize["mobileAppConfigurations"] = o.MobileAppConfigurations
	}
	if o.MobileApps != nil {
		toSerialize["mobileApps"] = o.MobileApps
	}
	if o.VppTokens != nil {
		toSerialize["vppTokens"] = o.VppTokens
	}
	if o.AndroidManagedAppProtections != nil {
		toSerialize["androidManagedAppProtections"] = o.AndroidManagedAppProtections
	}
	if o.DefaultManagedAppProtections != nil {
		toSerialize["defaultManagedAppProtections"] = o.DefaultManagedAppProtections
	}
	if o.IosManagedAppProtections != nil {
		toSerialize["iosManagedAppProtections"] = o.IosManagedAppProtections
	}
	if o.ManagedAppPolicies != nil {
		toSerialize["managedAppPolicies"] = o.ManagedAppPolicies
	}
	if o.ManagedAppRegistrations != nil {
		toSerialize["managedAppRegistrations"] = o.ManagedAppRegistrations
	}
	if o.ManagedAppStatuses != nil {
		toSerialize["managedAppStatuses"] = o.ManagedAppStatuses
	}
	if o.MdmWindowsInformationProtectionPolicies != nil {
		toSerialize["mdmWindowsInformationProtectionPolicies"] = o.MdmWindowsInformationProtectionPolicies
	}
	if o.TargetedManagedAppConfigurations != nil {
		toSerialize["targetedManagedAppConfigurations"] = o.TargetedManagedAppConfigurations
	}
	if o.WindowsInformationProtectionPolicies != nil {
		toSerialize["windowsInformationProtectionPolicies"] = o.WindowsInformationProtectionPolicies
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAppManagement struct {
	value *DeviceAppManagement
	isSet bool
}

func (v NullableDeviceAppManagement) Get() *DeviceAppManagement {
	return v.value
}

func (v *NullableDeviceAppManagement) Set(val *DeviceAppManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAppManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAppManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAppManagement(val *DeviceAppManagement) *NullableDeviceAppManagement {
	return &NullableDeviceAppManagement{value: val, isSet: true}
}

func (v NullableDeviceAppManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAppManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


