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

// MicrosoftGraphPrintJobConfiguration struct for MicrosoftGraphPrintJobConfiguration
type MicrosoftGraphPrintJobConfiguration struct {
	// Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
	Collate NullableBool `json:"collate,omitempty"`
	// The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
	ColorMode AnyOfmicrosoftGraphPrintColorMode `json:"colorMode,omitempty"`
	// The number of copies that should be printed. Read-only.
	Copies NullableInt32 `json:"copies,omitempty"`
	// The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
	Dpi NullableInt32 `json:"dpi,omitempty"`
	// The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
	DuplexMode AnyOfmicrosoftGraphPrintDuplexMode `json:"duplexMode,omitempty"`
	// The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
	FeedOrientation AnyOfmicrosoftGraphPrinterFeedOrientation `json:"feedOrientation,omitempty"`
	// Finishing processes to use when printing.
	Finishings *[]*AnyOfmicrosoftGraphPrintFinishing `json:"finishings,omitempty"`
	FitPdfToPage NullableBool `json:"fitPdfToPage,omitempty"`
	// The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
	InputBin NullableString `json:"inputBin,omitempty"`
	// The margin settings to use when printing.
	Margin AnyOfmicrosoftGraphPrintMargin `json:"margin,omitempty"`
	// The media size to use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values listed in the printerCapabilities topic.
	MediaSize NullableString `json:"mediaSize,omitempty"`
	// The default media (such as paper) type to print the document on.
	MediaType NullableString `json:"mediaType,omitempty"`
	// The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
	MultipageLayout AnyOfmicrosoftGraphPrintMultipageLayout `json:"multipageLayout,omitempty"`
	// The orientation setting the printer should use when printing the job. Valid values are described in the following table.
	Orientation AnyOfmicrosoftGraphPrintOrientation `json:"orientation,omitempty"`
	// The output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
	OutputBin NullableString `json:"outputBin,omitempty"`
	// The page ranges to print. Read-only.
	PageRanges *[]*AnyOfmicrosoftGraphIntegerRange `json:"pageRanges,omitempty"`
	// The number of document pages to print on each sheet.
	PagesPerSheet NullableInt32 `json:"pagesPerSheet,omitempty"`
	// The print quality to use when printing the job. Valid values are described in the table below. Read-only.
	Quality AnyOfmicrosoftGraphPrintQuality `json:"quality,omitempty"`
	// Specifies how the printer should scale the document data to fit the requested media. Valid values are described in the following table.
	Scaling AnyOfmicrosoftGraphPrintScaling `json:"scaling,omitempty"`
}

// NewMicrosoftGraphPrintJobConfiguration instantiates a new MicrosoftGraphPrintJobConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintJobConfiguration() *MicrosoftGraphPrintJobConfiguration {
	this := MicrosoftGraphPrintJobConfiguration{}
	return &this
}

// NewMicrosoftGraphPrintJobConfigurationWithDefaults instantiates a new MicrosoftGraphPrintJobConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintJobConfigurationWithDefaults() *MicrosoftGraphPrintJobConfiguration {
	this := MicrosoftGraphPrintJobConfiguration{}
	return &this
}

// GetCollate returns the Collate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetCollate() bool {
	if o == nil || o.Collate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Collate.Get()
}

// GetCollateOk returns a tuple with the Collate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetCollateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Collate.Get(), o.Collate.IsSet()
}

// HasCollate returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasCollate() bool {
	if o != nil && o.Collate.IsSet() {
		return true
	}

	return false
}

// SetCollate gets a reference to the given NullableBool and assigns it to the Collate field.
func (o *MicrosoftGraphPrintJobConfiguration) SetCollate(v bool) {
	o.Collate.Set(&v)
}
// SetCollateNil sets the value for Collate to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetCollateNil() {
	o.Collate.Set(nil)
}

// UnsetCollate ensures that no value is present for Collate, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetCollate() {
	o.Collate.Unset()
}

// GetColorMode returns the ColorMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetColorMode() AnyOfmicrosoftGraphPrintColorMode {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintColorMode
		return ret
	}
	return o.ColorMode
}

// GetColorModeOk returns a tuple with the ColorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetColorModeOk() (*AnyOfmicrosoftGraphPrintColorMode, bool) {
	if o == nil || o.ColorMode == nil {
		return nil, false
	}
	return &o.ColorMode, true
}

// HasColorMode returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasColorMode() bool {
	if o != nil && o.ColorMode != nil {
		return true
	}

	return false
}

// SetColorMode gets a reference to the given AnyOfmicrosoftGraphPrintColorMode and assigns it to the ColorMode field.
func (o *MicrosoftGraphPrintJobConfiguration) SetColorMode(v AnyOfmicrosoftGraphPrintColorMode) {
	o.ColorMode = v
}

// GetCopies returns the Copies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetCopies() int32 {
	if o == nil || o.Copies.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Copies.Get()
}

// GetCopiesOk returns a tuple with the Copies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetCopiesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Copies.Get(), o.Copies.IsSet()
}

// HasCopies returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasCopies() bool {
	if o != nil && o.Copies.IsSet() {
		return true
	}

	return false
}

// SetCopies gets a reference to the given NullableInt32 and assigns it to the Copies field.
func (o *MicrosoftGraphPrintJobConfiguration) SetCopies(v int32) {
	o.Copies.Set(&v)
}
// SetCopiesNil sets the value for Copies to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetCopiesNil() {
	o.Copies.Set(nil)
}

// UnsetCopies ensures that no value is present for Copies, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetCopies() {
	o.Copies.Unset()
}

// GetDpi returns the Dpi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetDpi() int32 {
	if o == nil || o.Dpi.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Dpi.Get()
}

// GetDpiOk returns a tuple with the Dpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetDpiOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dpi.Get(), o.Dpi.IsSet()
}

// HasDpi returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasDpi() bool {
	if o != nil && o.Dpi.IsSet() {
		return true
	}

	return false
}

// SetDpi gets a reference to the given NullableInt32 and assigns it to the Dpi field.
func (o *MicrosoftGraphPrintJobConfiguration) SetDpi(v int32) {
	o.Dpi.Set(&v)
}
// SetDpiNil sets the value for Dpi to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetDpiNil() {
	o.Dpi.Set(nil)
}

// UnsetDpi ensures that no value is present for Dpi, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetDpi() {
	o.Dpi.Unset()
}

// GetDuplexMode returns the DuplexMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetDuplexMode() AnyOfmicrosoftGraphPrintDuplexMode {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintDuplexMode
		return ret
	}
	return o.DuplexMode
}

// GetDuplexModeOk returns a tuple with the DuplexMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetDuplexModeOk() (*AnyOfmicrosoftGraphPrintDuplexMode, bool) {
	if o == nil || o.DuplexMode == nil {
		return nil, false
	}
	return &o.DuplexMode, true
}

// HasDuplexMode returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasDuplexMode() bool {
	if o != nil && o.DuplexMode != nil {
		return true
	}

	return false
}

// SetDuplexMode gets a reference to the given AnyOfmicrosoftGraphPrintDuplexMode and assigns it to the DuplexMode field.
func (o *MicrosoftGraphPrintJobConfiguration) SetDuplexMode(v AnyOfmicrosoftGraphPrintDuplexMode) {
	o.DuplexMode = v
}

// GetFeedOrientation returns the FeedOrientation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetFeedOrientation() AnyOfmicrosoftGraphPrinterFeedOrientation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrinterFeedOrientation
		return ret
	}
	return o.FeedOrientation
}

// GetFeedOrientationOk returns a tuple with the FeedOrientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetFeedOrientationOk() (*AnyOfmicrosoftGraphPrinterFeedOrientation, bool) {
	if o == nil || o.FeedOrientation == nil {
		return nil, false
	}
	return &o.FeedOrientation, true
}

// HasFeedOrientation returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasFeedOrientation() bool {
	if o != nil && o.FeedOrientation != nil {
		return true
	}

	return false
}

// SetFeedOrientation gets a reference to the given AnyOfmicrosoftGraphPrinterFeedOrientation and assigns it to the FeedOrientation field.
func (o *MicrosoftGraphPrintJobConfiguration) SetFeedOrientation(v AnyOfmicrosoftGraphPrinterFeedOrientation) {
	o.FeedOrientation = v
}

// GetFinishings returns the Finishings field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintJobConfiguration) GetFinishings() []*AnyOfmicrosoftGraphPrintFinishing {
	if o == nil || o.Finishings == nil {
		var ret []*AnyOfmicrosoftGraphPrintFinishing
		return ret
	}
	return *o.Finishings
}

// GetFinishingsOk returns a tuple with the Finishings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintJobConfiguration) GetFinishingsOk() (*[]*AnyOfmicrosoftGraphPrintFinishing, bool) {
	if o == nil || o.Finishings == nil {
		return nil, false
	}
	return o.Finishings, true
}

// HasFinishings returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasFinishings() bool {
	if o != nil && o.Finishings != nil {
		return true
	}

	return false
}

// SetFinishings gets a reference to the given []*AnyOfmicrosoftGraphPrintFinishing and assigns it to the Finishings field.
func (o *MicrosoftGraphPrintJobConfiguration) SetFinishings(v []*AnyOfmicrosoftGraphPrintFinishing) {
	o.Finishings = &v
}

// GetFitPdfToPage returns the FitPdfToPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetFitPdfToPage() bool {
	if o == nil || o.FitPdfToPage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FitPdfToPage.Get()
}

// GetFitPdfToPageOk returns a tuple with the FitPdfToPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetFitPdfToPageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FitPdfToPage.Get(), o.FitPdfToPage.IsSet()
}

// HasFitPdfToPage returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasFitPdfToPage() bool {
	if o != nil && o.FitPdfToPage.IsSet() {
		return true
	}

	return false
}

// SetFitPdfToPage gets a reference to the given NullableBool and assigns it to the FitPdfToPage field.
func (o *MicrosoftGraphPrintJobConfiguration) SetFitPdfToPage(v bool) {
	o.FitPdfToPage.Set(&v)
}
// SetFitPdfToPageNil sets the value for FitPdfToPage to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetFitPdfToPageNil() {
	o.FitPdfToPage.Set(nil)
}

// UnsetFitPdfToPage ensures that no value is present for FitPdfToPage, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetFitPdfToPage() {
	o.FitPdfToPage.Unset()
}

// GetInputBin returns the InputBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetInputBin() string {
	if o == nil || o.InputBin.Get() == nil {
		var ret string
		return ret
	}
	return *o.InputBin.Get()
}

// GetInputBinOk returns a tuple with the InputBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetInputBinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InputBin.Get(), o.InputBin.IsSet()
}

// HasInputBin returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasInputBin() bool {
	if o != nil && o.InputBin.IsSet() {
		return true
	}

	return false
}

// SetInputBin gets a reference to the given NullableString and assigns it to the InputBin field.
func (o *MicrosoftGraphPrintJobConfiguration) SetInputBin(v string) {
	o.InputBin.Set(&v)
}
// SetInputBinNil sets the value for InputBin to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetInputBinNil() {
	o.InputBin.Set(nil)
}

// UnsetInputBin ensures that no value is present for InputBin, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetInputBin() {
	o.InputBin.Unset()
}

// GetMargin returns the Margin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetMargin() AnyOfmicrosoftGraphPrintMargin {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintMargin
		return ret
	}
	return o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetMarginOk() (*AnyOfmicrosoftGraphPrintMargin, bool) {
	if o == nil || o.Margin == nil {
		return nil, false
	}
	return &o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasMargin() bool {
	if o != nil && o.Margin != nil {
		return true
	}

	return false
}

// SetMargin gets a reference to the given AnyOfmicrosoftGraphPrintMargin and assigns it to the Margin field.
func (o *MicrosoftGraphPrintJobConfiguration) SetMargin(v AnyOfmicrosoftGraphPrintMargin) {
	o.Margin = v
}

// GetMediaSize returns the MediaSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetMediaSize() string {
	if o == nil || o.MediaSize.Get() == nil {
		var ret string
		return ret
	}
	return *o.MediaSize.Get()
}

// GetMediaSizeOk returns a tuple with the MediaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetMediaSizeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MediaSize.Get(), o.MediaSize.IsSet()
}

// HasMediaSize returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasMediaSize() bool {
	if o != nil && o.MediaSize.IsSet() {
		return true
	}

	return false
}

// SetMediaSize gets a reference to the given NullableString and assigns it to the MediaSize field.
func (o *MicrosoftGraphPrintJobConfiguration) SetMediaSize(v string) {
	o.MediaSize.Set(&v)
}
// SetMediaSizeNil sets the value for MediaSize to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetMediaSizeNil() {
	o.MediaSize.Set(nil)
}

// UnsetMediaSize ensures that no value is present for MediaSize, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetMediaSize() {
	o.MediaSize.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetMediaType() string {
	if o == nil || o.MediaType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetMediaTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasMediaType() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableString and assigns it to the MediaType field.
func (o *MicrosoftGraphPrintJobConfiguration) SetMediaType(v string) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetMediaType() {
	o.MediaType.Unset()
}

// GetMultipageLayout returns the MultipageLayout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetMultipageLayout() AnyOfmicrosoftGraphPrintMultipageLayout {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintMultipageLayout
		return ret
	}
	return o.MultipageLayout
}

// GetMultipageLayoutOk returns a tuple with the MultipageLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetMultipageLayoutOk() (*AnyOfmicrosoftGraphPrintMultipageLayout, bool) {
	if o == nil || o.MultipageLayout == nil {
		return nil, false
	}
	return &o.MultipageLayout, true
}

// HasMultipageLayout returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasMultipageLayout() bool {
	if o != nil && o.MultipageLayout != nil {
		return true
	}

	return false
}

// SetMultipageLayout gets a reference to the given AnyOfmicrosoftGraphPrintMultipageLayout and assigns it to the MultipageLayout field.
func (o *MicrosoftGraphPrintJobConfiguration) SetMultipageLayout(v AnyOfmicrosoftGraphPrintMultipageLayout) {
	o.MultipageLayout = v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetOrientation() AnyOfmicrosoftGraphPrintOrientation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintOrientation
		return ret
	}
	return o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetOrientationOk() (*AnyOfmicrosoftGraphPrintOrientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return &o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given AnyOfmicrosoftGraphPrintOrientation and assigns it to the Orientation field.
func (o *MicrosoftGraphPrintJobConfiguration) SetOrientation(v AnyOfmicrosoftGraphPrintOrientation) {
	o.Orientation = v
}

// GetOutputBin returns the OutputBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetOutputBin() string {
	if o == nil || o.OutputBin.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutputBin.Get()
}

// GetOutputBinOk returns a tuple with the OutputBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetOutputBinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OutputBin.Get(), o.OutputBin.IsSet()
}

// HasOutputBin returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasOutputBin() bool {
	if o != nil && o.OutputBin.IsSet() {
		return true
	}

	return false
}

// SetOutputBin gets a reference to the given NullableString and assigns it to the OutputBin field.
func (o *MicrosoftGraphPrintJobConfiguration) SetOutputBin(v string) {
	o.OutputBin.Set(&v)
}
// SetOutputBinNil sets the value for OutputBin to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetOutputBinNil() {
	o.OutputBin.Set(nil)
}

// UnsetOutputBin ensures that no value is present for OutputBin, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetOutputBin() {
	o.OutputBin.Unset()
}

// GetPageRanges returns the PageRanges field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintJobConfiguration) GetPageRanges() []*AnyOfmicrosoftGraphIntegerRange {
	if o == nil || o.PageRanges == nil {
		var ret []*AnyOfmicrosoftGraphIntegerRange
		return ret
	}
	return *o.PageRanges
}

// GetPageRangesOk returns a tuple with the PageRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintJobConfiguration) GetPageRangesOk() (*[]*AnyOfmicrosoftGraphIntegerRange, bool) {
	if o == nil || o.PageRanges == nil {
		return nil, false
	}
	return o.PageRanges, true
}

// HasPageRanges returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasPageRanges() bool {
	if o != nil && o.PageRanges != nil {
		return true
	}

	return false
}

// SetPageRanges gets a reference to the given []*AnyOfmicrosoftGraphIntegerRange and assigns it to the PageRanges field.
func (o *MicrosoftGraphPrintJobConfiguration) SetPageRanges(v []*AnyOfmicrosoftGraphIntegerRange) {
	o.PageRanges = &v
}

// GetPagesPerSheet returns the PagesPerSheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetPagesPerSheet() int32 {
	if o == nil || o.PagesPerSheet.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PagesPerSheet.Get()
}

// GetPagesPerSheetOk returns a tuple with the PagesPerSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetPagesPerSheetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PagesPerSheet.Get(), o.PagesPerSheet.IsSet()
}

// HasPagesPerSheet returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasPagesPerSheet() bool {
	if o != nil && o.PagesPerSheet.IsSet() {
		return true
	}

	return false
}

// SetPagesPerSheet gets a reference to the given NullableInt32 and assigns it to the PagesPerSheet field.
func (o *MicrosoftGraphPrintJobConfiguration) SetPagesPerSheet(v int32) {
	o.PagesPerSheet.Set(&v)
}
// SetPagesPerSheetNil sets the value for PagesPerSheet to be an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) SetPagesPerSheetNil() {
	o.PagesPerSheet.Set(nil)
}

// UnsetPagesPerSheet ensures that no value is present for PagesPerSheet, not even an explicit nil
func (o *MicrosoftGraphPrintJobConfiguration) UnsetPagesPerSheet() {
	o.PagesPerSheet.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetQuality() AnyOfmicrosoftGraphPrintQuality {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintQuality
		return ret
	}
	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetQualityOk() (*AnyOfmicrosoftGraphPrintQuality, bool) {
	if o == nil || o.Quality == nil {
		return nil, false
	}
	return &o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasQuality() bool {
	if o != nil && o.Quality != nil {
		return true
	}

	return false
}

// SetQuality gets a reference to the given AnyOfmicrosoftGraphPrintQuality and assigns it to the Quality field.
func (o *MicrosoftGraphPrintJobConfiguration) SetQuality(v AnyOfmicrosoftGraphPrintQuality) {
	o.Quality = v
}

// GetScaling returns the Scaling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobConfiguration) GetScaling() AnyOfmicrosoftGraphPrintScaling {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintScaling
		return ret
	}
	return o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobConfiguration) GetScalingOk() (*AnyOfmicrosoftGraphPrintScaling, bool) {
	if o == nil || o.Scaling == nil {
		return nil, false
	}
	return &o.Scaling, true
}

// HasScaling returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobConfiguration) HasScaling() bool {
	if o != nil && o.Scaling != nil {
		return true
	}

	return false
}

// SetScaling gets a reference to the given AnyOfmicrosoftGraphPrintScaling and assigns it to the Scaling field.
func (o *MicrosoftGraphPrintJobConfiguration) SetScaling(v AnyOfmicrosoftGraphPrintScaling) {
	o.Scaling = v
}

func (o MicrosoftGraphPrintJobConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collate.IsSet() {
		toSerialize["collate"] = o.Collate.Get()
	}
	if o.ColorMode != nil {
		toSerialize["colorMode"] = o.ColorMode
	}
	if o.Copies.IsSet() {
		toSerialize["copies"] = o.Copies.Get()
	}
	if o.Dpi.IsSet() {
		toSerialize["dpi"] = o.Dpi.Get()
	}
	if o.DuplexMode != nil {
		toSerialize["duplexMode"] = o.DuplexMode
	}
	if o.FeedOrientation != nil {
		toSerialize["feedOrientation"] = o.FeedOrientation
	}
	if o.Finishings != nil {
		toSerialize["finishings"] = o.Finishings
	}
	if o.FitPdfToPage.IsSet() {
		toSerialize["fitPdfToPage"] = o.FitPdfToPage.Get()
	}
	if o.InputBin.IsSet() {
		toSerialize["inputBin"] = o.InputBin.Get()
	}
	if o.Margin != nil {
		toSerialize["margin"] = o.Margin
	}
	if o.MediaSize.IsSet() {
		toSerialize["mediaSize"] = o.MediaSize.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["mediaType"] = o.MediaType.Get()
	}
	if o.MultipageLayout != nil {
		toSerialize["multipageLayout"] = o.MultipageLayout
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.OutputBin.IsSet() {
		toSerialize["outputBin"] = o.OutputBin.Get()
	}
	if o.PageRanges != nil {
		toSerialize["pageRanges"] = o.PageRanges
	}
	if o.PagesPerSheet.IsSet() {
		toSerialize["pagesPerSheet"] = o.PagesPerSheet.Get()
	}
	if o.Quality != nil {
		toSerialize["quality"] = o.Quality
	}
	if o.Scaling != nil {
		toSerialize["scaling"] = o.Scaling
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintJobConfiguration struct {
	value *MicrosoftGraphPrintJobConfiguration
	isSet bool
}

func (v NullableMicrosoftGraphPrintJobConfiguration) Get() *MicrosoftGraphPrintJobConfiguration {
	return v.value
}

func (v *NullableMicrosoftGraphPrintJobConfiguration) Set(val *MicrosoftGraphPrintJobConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintJobConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintJobConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintJobConfiguration(val *MicrosoftGraphPrintJobConfiguration) *NullableMicrosoftGraphPrintJobConfiguration {
	return &NullableMicrosoftGraphPrintJobConfiguration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintJobConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintJobConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


