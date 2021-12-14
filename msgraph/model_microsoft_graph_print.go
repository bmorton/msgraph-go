/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphPrint struct for MicrosoftGraphPrint
type MicrosoftGraphPrint struct {
	// Tenant-wide settings for the Universal Print service.
	Settings AnyOfmicrosoftGraphPrintSettings `json:"settings,omitempty"`
	// The list of available print connectors.
	Connectors *[]MicrosoftGraphPrintConnector `json:"connectors,omitempty"`
	// The list of print long running operations.
	Operations *[]MicrosoftGraphPrintOperation `json:"operations,omitempty"`
	// The list of printers registered in the tenant.
	Printers *[]MicrosoftGraphPrinter `json:"printers,omitempty"`
	// The list of available Universal Print service endpoints.
	Services *[]MicrosoftGraphPrintService `json:"services,omitempty"`
	// The list of printer shares registered in the tenant.
	Shares *[]MicrosoftGraphPrinterShare `json:"shares,omitempty"`
	// List of abstract definition for a task that can be triggered when various events occur within Universal Print.
	TaskDefinitions *[]MicrosoftGraphPrintTaskDefinition `json:"taskDefinitions,omitempty"`
}

// NewMicrosoftGraphPrint instantiates a new MicrosoftGraphPrint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrint() *MicrosoftGraphPrint {
	this := MicrosoftGraphPrint{}
	return &this
}

// NewMicrosoftGraphPrintWithDefaults instantiates a new MicrosoftGraphPrint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintWithDefaults() *MicrosoftGraphPrint {
	this := MicrosoftGraphPrint{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrint) GetSettings() AnyOfmicrosoftGraphPrintSettings {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintSettings
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrint) GetSettingsOk() (*AnyOfmicrosoftGraphPrintSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AnyOfmicrosoftGraphPrintSettings and assigns it to the Settings field.
func (o *MicrosoftGraphPrint) SetSettings(v AnyOfmicrosoftGraphPrintSettings) {
	o.Settings = v
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetConnectors() []MicrosoftGraphPrintConnector {
	if o == nil || o.Connectors == nil {
		var ret []MicrosoftGraphPrintConnector
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetConnectorsOk() (*[]MicrosoftGraphPrintConnector, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []MicrosoftGraphPrintConnector and assigns it to the Connectors field.
func (o *MicrosoftGraphPrint) SetConnectors(v []MicrosoftGraphPrintConnector) {
	o.Connectors = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetOperations() []MicrosoftGraphPrintOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphPrintOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetOperationsOk() (*[]MicrosoftGraphPrintOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphPrintOperation and assigns it to the Operations field.
func (o *MicrosoftGraphPrint) SetOperations(v []MicrosoftGraphPrintOperation) {
	o.Operations = &v
}

// GetPrinters returns the Printers field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetPrinters() []MicrosoftGraphPrinter {
	if o == nil || o.Printers == nil {
		var ret []MicrosoftGraphPrinter
		return ret
	}
	return *o.Printers
}

// GetPrintersOk returns a tuple with the Printers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetPrintersOk() (*[]MicrosoftGraphPrinter, bool) {
	if o == nil || o.Printers == nil {
		return nil, false
	}
	return o.Printers, true
}

// HasPrinters returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasPrinters() bool {
	if o != nil && o.Printers != nil {
		return true
	}

	return false
}

// SetPrinters gets a reference to the given []MicrosoftGraphPrinter and assigns it to the Printers field.
func (o *MicrosoftGraphPrint) SetPrinters(v []MicrosoftGraphPrinter) {
	o.Printers = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetServices() []MicrosoftGraphPrintService {
	if o == nil || o.Services == nil {
		var ret []MicrosoftGraphPrintService
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetServicesOk() (*[]MicrosoftGraphPrintService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []MicrosoftGraphPrintService and assigns it to the Services field.
func (o *MicrosoftGraphPrint) SetServices(v []MicrosoftGraphPrintService) {
	o.Services = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetShares() []MicrosoftGraphPrinterShare {
	if o == nil || o.Shares == nil {
		var ret []MicrosoftGraphPrinterShare
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetSharesOk() (*[]MicrosoftGraphPrinterShare, bool) {
	if o == nil || o.Shares == nil {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasShares() bool {
	if o != nil && o.Shares != nil {
		return true
	}

	return false
}

// SetShares gets a reference to the given []MicrosoftGraphPrinterShare and assigns it to the Shares field.
func (o *MicrosoftGraphPrint) SetShares(v []MicrosoftGraphPrinterShare) {
	o.Shares = &v
}

// GetTaskDefinitions returns the TaskDefinitions field value if set, zero value otherwise.
func (o *MicrosoftGraphPrint) GetTaskDefinitions() []MicrosoftGraphPrintTaskDefinition {
	if o == nil || o.TaskDefinitions == nil {
		var ret []MicrosoftGraphPrintTaskDefinition
		return ret
	}
	return *o.TaskDefinitions
}

// GetTaskDefinitionsOk returns a tuple with the TaskDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrint) GetTaskDefinitionsOk() (*[]MicrosoftGraphPrintTaskDefinition, bool) {
	if o == nil || o.TaskDefinitions == nil {
		return nil, false
	}
	return o.TaskDefinitions, true
}

// HasTaskDefinitions returns a boolean if a field has been set.
func (o *MicrosoftGraphPrint) HasTaskDefinitions() bool {
	if o != nil && o.TaskDefinitions != nil {
		return true
	}

	return false
}

// SetTaskDefinitions gets a reference to the given []MicrosoftGraphPrintTaskDefinition and assigns it to the TaskDefinitions field.
func (o *MicrosoftGraphPrint) SetTaskDefinitions(v []MicrosoftGraphPrintTaskDefinition) {
	o.TaskDefinitions = &v
}

func (o MicrosoftGraphPrint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Connectors != nil {
		toSerialize["connectors"] = o.Connectors
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Printers != nil {
		toSerialize["printers"] = o.Printers
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Shares != nil {
		toSerialize["shares"] = o.Shares
	}
	if o.TaskDefinitions != nil {
		toSerialize["taskDefinitions"] = o.TaskDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrint struct {
	value *MicrosoftGraphPrint
	isSet bool
}

func (v NullableMicrosoftGraphPrint) Get() *MicrosoftGraphPrint {
	return v.value
}

func (v *NullableMicrosoftGraphPrint) Set(val *MicrosoftGraphPrint) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrint) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrint(val *MicrosoftGraphPrint) *NullableMicrosoftGraphPrint {
	return &NullableMicrosoftGraphPrint{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


