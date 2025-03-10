// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package playintegrity provides access to the Google Play Integrity API.
//
// For product documentation, see: https://developer.android.com/google/play/integrity
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/playintegrity/v1"
//	...
//	ctx := context.Background()
//	playintegrityService, err := playintegrity.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	playintegrityService, err := playintegrity.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	playintegrityService, err := playintegrity.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package playintegrity // import "google.golang.org/api/playintegrity/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "playintegrity:v1"
const apiName = "playintegrity"
const apiVersion = "v1"
const basePath = "https://playintegrity.googleapis.com/"
const basePathTemplate = "https://playintegrity.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://playintegrity.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Private Service: https://www.googleapis.com/auth/playintegrity
	PlayintegrityScope = "https://www.googleapis.com/auth/playintegrity"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/playintegrity",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.EnableNewAuthLibrary())
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1 *V1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

// AccountActivity: (Restricted Access) Contains a signal helping apps
// differentiating between likely genuine and likely non-genuine user traffic.
type AccountActivity struct {
	// ActivityLevel: Required. Indicates the activity level of the account.
	//
	// Possible values:
	//   "ACTIVITY_LEVEL_UNSPECIFIED" - Activity level has not been set.
	//   "UNEVALUATED" - Account activity level is not evaluated.
	//   "UNUSUAL" - Unusual activity for at least one of the user accounts on the
	// device.
	//   "UNKNOWN" - Insufficient activity to verify the user account on the
	// device.
	//   "TYPICAL_BASIC" - Typical activity for the user account or accounts on the
	// device.
	//   "TYPICAL_STRONG" - Typical for the user account or accounts on the device,
	// with harder to replicate signals.
	ActivityLevel string `json:"activityLevel,omitempty"`
	// ForceSendFields is a list of field names (e.g. "ActivityLevel") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "ActivityLevel") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s AccountActivity) MarshalJSON() ([]byte, error) {
	type NoMethod AccountActivity
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// AccountDetails: Contains the account information such as the licensing
// status for the user in the scope.
type AccountDetails struct {
	// AccountActivity: (Restricted Access) Details about the account activity for
	// the user in the scope.
	AccountActivity *AccountActivity `json:"accountActivity,omitempty"`
	// AppLicensingVerdict: Required. Details about the licensing status of the
	// user for the app in the scope.
	//
	// Possible values:
	//   "UNKNOWN" - Play does not have sufficient information to evaluate
	// licensing details
	//   "LICENSED" - The app and certificate match the versions distributed by
	// Play.
	//   "UNLICENSED" - The certificate or package name does not match Google Play
	// records.
	//   "UNEVALUATED" - Licensing details were not evaluated since a necessary
	// requirement was missed. For example DeviceIntegrity did not meet the minimum
	// bar or the application was not a known Play version.
	AppLicensingVerdict string `json:"appLicensingVerdict,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AccountActivity") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AccountActivity") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s AccountDetails) MarshalJSON() ([]byte, error) {
	type NoMethod AccountDetails
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// AppAccessRiskVerdict: Contains signals about others apps on the device which
// could be used to access or control the requesting app.
type AppAccessRiskVerdict struct {
	// AppsDetected: List of detected app types signalled for App Access Risk.
	//
	// Possible values:
	//   "APPS_DETECTED_UNSPECIFIED" - Apps detected is unspecified.
	//   "KNOWN_INSTALLED" - One or more apps is installed by Google Play or
	// preloaded on the system partition by the device manufacturer.
	//   "KNOWN_CAPTURING" - One or more apps installed by Google Play or preloaded
	// on the device is running that could be used to read or capture the
	// requesting app, such as a screen recording app.
	//   "KNOWN_OVERLAYS" - One or more apps installed by Google Play or preloaded
	// on the device is running that could be used to display overlays over the
	// requesting app.
	//   "KNOWN_CONTROLLING" - One or more apps installed by Google Play or
	// preloaded on the device is running that could be used to control the device,
	// such as a remote support app.
	//   "UNKNOWN_INSTALLED" - One or more unknown apps is installed, that were not
	// installed by Google Play or preloaded on the system partition by the device
	// manufacturer.
	//   "UNKNOWN_CAPTURING" - One or more unknown apps, which were not installed
	// by Google Play or preloaded on the device, is running that could be used to
	// read or capture the requesting app, such as a screen recording app.
	//   "UNKNOWN_OVERLAYS" - One or more unknown apps, which were not installed by
	// Google Play or preloaded on the device, is running that could be used to
	// display overlays over the requesting app.
	//   "UNKNOWN_CONTROLLING" - One or more unknown apps, which were not installed
	// by Google Play or preloaded on the device, is running that could be used to
	// control the device, such as a remote support app.
	AppsDetected []string `json:"appsDetected,omitempty"`
	// OtherApps: Deprecated: this field will be removed, please use apps_detected
	// instead. App access risk verdict related to apps that are not installed by
	// Google Play, and are not preloaded on the system image by the device
	// manufacturer.
	//
	// Possible values:
	//   "UNKNOWN" - Risk type is unknown.
	//   "UNEVALUATED" - App access risk was not evaluated because a requirement
	// was missed, such as the device not being trusted enough.
	//   "NOT_INSTALLED" - No apps under this field are installed on the device.
	// This is only valid for the other apps field.
	//   "INSTALLED" - One or more apps under this field are installed on the
	// device.
	//   "CAPTURING" - Apps under this field are running that could be used to read
	// or capture inputs and outputs of the requesting app, such as screen
	// recording apps.
	//   "CONTROLLING" - Apps under this field are running that could be used to
	// control the device and inputs and outputs of the requesting app, such as
	// remote controlling apps.
	OtherApps string `json:"otherApps,omitempty"`
	// PlayOrSystemApps: Deprecated: this field will be removed, please use
	// apps_detected instead. App access risk verdict related to apps that are not
	// installed by the Google Play Store, and are not preloaded on the system
	// image by the device manufacturer.
	//
	// Possible values:
	//   "UNKNOWN" - Risk type is unknown.
	//   "UNEVALUATED" - App access risk was not evaluated because a requirement
	// was missed, such as the device not being trusted enough.
	//   "NOT_INSTALLED" - No apps under this field are installed on the device.
	// This is only valid for the other apps field.
	//   "INSTALLED" - One or more apps under this field are installed on the
	// device.
	//   "CAPTURING" - Apps under this field are running that could be used to read
	// or capture inputs and outputs of the requesting app, such as screen
	// recording apps.
	//   "CONTROLLING" - Apps under this field are running that could be used to
	// control the device and inputs and outputs of the requesting app, such as
	// remote controlling apps.
	PlayOrSystemApps string `json:"playOrSystemApps,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AppsDetected") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AppsDetected") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s AppAccessRiskVerdict) MarshalJSON() ([]byte, error) {
	type NoMethod AppAccessRiskVerdict
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// AppIntegrity: Contains the application integrity information.
type AppIntegrity struct {
	// AppRecognitionVerdict: Required. Details about the app recognition verdict
	//
	// Possible values:
	//   "UNKNOWN" - Play does not have sufficient information to evaluate app
	// integrity
	//   "PLAY_RECOGNIZED" - The app and certificate match the versions distributed
	// by Play.
	//   "UNRECOGNIZED_VERSION" - The certificate or package name does not match
	// Google Play records.
	//   "UNEVALUATED" - Application integrity was not evaluated since a necessary
	// requirement was missed. For example DeviceIntegrity did not meet the minimum
	// bar.
	AppRecognitionVerdict string `json:"appRecognitionVerdict,omitempty"`
	// CertificateSha256Digest: The SHA256 hash of the requesting app's signing
	// certificates (base64 web-safe encoded). Set iff app_recognition_verdict !=
	// UNEVALUATED.
	CertificateSha256Digest []string `json:"certificateSha256Digest,omitempty"`
	// PackageName: Package name of the application under attestation. Set iff
	// app_recognition_verdict != UNEVALUATED.
	PackageName string `json:"packageName,omitempty"`
	// VersionCode: Version code of the application. Set iff
	// app_recognition_verdict != UNEVALUATED.
	VersionCode int64 `json:"versionCode,omitempty,string"`
	// ForceSendFields is a list of field names (e.g. "AppRecognitionVerdict") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AppRecognitionVerdict") to
	// include in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s AppIntegrity) MarshalJSON() ([]byte, error) {
	type NoMethod AppIntegrity
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// DecodeIntegrityTokenRequest: Request to decode the integrity token.
type DecodeIntegrityTokenRequest struct {
	// IntegrityToken: Encoded integrity token.
	IntegrityToken string `json:"integrityToken,omitempty"`
	// ForceSendFields is a list of field names (e.g. "IntegrityToken") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "IntegrityToken") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s DecodeIntegrityTokenRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DecodeIntegrityTokenRequest
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// DecodeIntegrityTokenResponse: Response containing the decoded integrity
// payload.
type DecodeIntegrityTokenResponse struct {
	// TokenPayloadExternal: Plain token payload generated from the decoded
	// integrity token.
	TokenPayloadExternal *TokenPayloadExternal `json:"tokenPayloadExternal,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "TokenPayloadExternal") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "TokenPayloadExternal") to include
	// in API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s DecodeIntegrityTokenResponse) MarshalJSON() ([]byte, error) {
	type NoMethod DecodeIntegrityTokenResponse
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// DeviceIntegrity: Contains the device attestation information. Next tag: 4
type DeviceIntegrity struct {
	// DeviceRecognitionVerdict: Details about the integrity of the device the app
	// is running on.
	//
	// Possible values:
	//   "UNKNOWN" - Play does not have sufficient information to evaluate device
	// integrity
	//   "MEETS_BASIC_INTEGRITY" - App is running on a device that passes basic
	// system integrity checks, but may not meet Android platform compatibility
	// requirements and may not be approved to run Google Play services.
	//   "MEETS_DEVICE_INTEGRITY" - App is running on GMS Android device with
	// Google Play services.
	//   "MEETS_STRONG_INTEGRITY" - App is running on GMS Android device with
	// Google Play services and has a strong guarantee of system integrity such as
	// a hardware-backed keystore.
	//   "MEETS_VIRTUAL_INTEGRITY" - App is running on an Android emulator with
	// Google Play services which meets core Android compatibility requirements.
	DeviceRecognitionVerdict []string `json:"deviceRecognitionVerdict,omitempty"`
	// RecentDeviceActivity: Details about the device activity of the device the
	// app is running on.
	RecentDeviceActivity *RecentDeviceActivity `json:"recentDeviceActivity,omitempty"`
	// ForceSendFields is a list of field names (e.g. "DeviceRecognitionVerdict")
	// to unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "DeviceRecognitionVerdict") to
	// include in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s DeviceIntegrity) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceIntegrity
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// EnvironmentDetails: Contains information about the environment Play
// Integrity API runs in, e.g. Play Protect verdict.
type EnvironmentDetails struct {
	// AppAccessRiskVerdict: The evaluation of the App Access Risk verdicts.
	AppAccessRiskVerdict *AppAccessRiskVerdict `json:"appAccessRiskVerdict,omitempty"`
	// PlayProtectVerdict: The evaluation of Play Protect verdict.
	//
	// Possible values:
	//   "PLAY_PROTECT_VERDICT_UNSPECIFIED" - Play Protect verdict has not been
	// set.
	//   "UNEVALUATED" - Play Protect state was not evaluated. Device may not be
	// trusted.
	//   "NO_ISSUES" - Play Protect is on and no issues found.
	//   "NO_DATA" - Play Protect is on but no scan has been performed yet. The
	// device or Play Store app may have been reset.
	//   "MEDIUM_RISK" - Play Protect is on and warnings found.
	//   "HIGH_RISK" - Play Protect is on and high severity issues found.
	//   "POSSIBLE_RISK" - Play Protect is turned off. Turn on Play Protect.
	PlayProtectVerdict string `json:"playProtectVerdict,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AppAccessRiskVerdict") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AppAccessRiskVerdict") to include
	// in API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s EnvironmentDetails) MarshalJSON() ([]byte, error) {
	type NoMethod EnvironmentDetails
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// RecentDeviceActivity: Recent device activity can help developers identify
// devices that have exhibited hyperactive attestation activity, which could be
// a sign of an attack or token farming.
type RecentDeviceActivity struct {
	// DeviceActivityLevel: Required. Indicates the activity level of the device.
	//
	// Possible values:
	//   "DEVICE_ACTIVITY_LEVEL_UNSPECIFIED" - Device activity level has not been
	// set.
	//   "UNEVALUATED" - Device activity level has not been evaluated.
	//   "LEVEL_1" - Indicates the amount of used tokens. See the documentation for
	// details.
	//   "LEVEL_2" - Indicates the amount of used tokens. See the documentation for
	// details.
	//   "LEVEL_3" - Indicates the amount of used tokens. See the documentation for
	// details.
	//   "LEVEL_4" - Indicates the amount of used tokens. See the documentation for
	// details.
	DeviceActivityLevel string `json:"deviceActivityLevel,omitempty"`
	// ForceSendFields is a list of field names (e.g. "DeviceActivityLevel") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "DeviceActivityLevel") to include
	// in API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s RecentDeviceActivity) MarshalJSON() ([]byte, error) {
	type NoMethod RecentDeviceActivity
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// RequestDetails: Contains the integrity request information.
type RequestDetails struct {
	// Nonce: Nonce that was provided in the request (which is base64 web-safe
	// no-wrap).
	Nonce string `json:"nonce,omitempty"`
	// RequestHash: Request hash that was provided in the request.
	RequestHash string `json:"requestHash,omitempty"`
	// RequestPackageName: Required. Application package name this attestation was
	// requested for. Note: This field makes no guarantees or promises on the
	// caller integrity. For details on application integrity, check
	// application_integrity.
	RequestPackageName string `json:"requestPackageName,omitempty"`
	// TimestampMillis: Required. Timestamp, in milliseconds, of the integrity
	// application request.
	TimestampMillis int64 `json:"timestampMillis,omitempty,string"`
	// ForceSendFields is a list of field names (e.g. "Nonce") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Nonce") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s RequestDetails) MarshalJSON() ([]byte, error) {
	type NoMethod RequestDetails
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// TestingDetails: Contains additional information generated for testing
// responses.
type TestingDetails struct {
	// IsTestingResponse: Required. Indicates that the information contained in
	// this payload is a testing response that is statically overridden for a
	// tester.
	IsTestingResponse bool `json:"isTestingResponse,omitempty"`
	// ForceSendFields is a list of field names (e.g. "IsTestingResponse") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "IsTestingResponse") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s TestingDetails) MarshalJSON() ([]byte, error) {
	type NoMethod TestingDetails
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

// TokenPayloadExternal: Contains basic app information and integrity signals
// like device attestation and licensing details.
type TokenPayloadExternal struct {
	// AccountDetails: Required. Details about the Play Store account.
	AccountDetails *AccountDetails `json:"accountDetails,omitempty"`
	// AppIntegrity: Required. Details about the application integrity.
	AppIntegrity *AppIntegrity `json:"appIntegrity,omitempty"`
	// DeviceIntegrity: Required. Details about the device integrity.
	DeviceIntegrity *DeviceIntegrity `json:"deviceIntegrity,omitempty"`
	// EnvironmentDetails: Details of the environment Play Integrity API runs in.
	EnvironmentDetails *EnvironmentDetails `json:"environmentDetails,omitempty"`
	// RequestDetails: Required. Details about the integrity request.
	RequestDetails *RequestDetails `json:"requestDetails,omitempty"`
	// TestingDetails: Indicates that this payload is generated for testing
	// purposes and contains any additional data that is linked with testing
	// status.
	TestingDetails *TestingDetails `json:"testingDetails,omitempty"`
	// ForceSendFields is a list of field names (e.g. "AccountDetails") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "AccountDetails") to include in
	// API requests with the JSON null value. By default, fields with empty values
	// are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s TokenPayloadExternal) MarshalJSON() ([]byte, error) {
	type NoMethod TokenPayloadExternal
	return gensupport.MarshalJSON(NoMethod(s), s.ForceSendFields, s.NullFields)
}

type V1DecodeIntegrityTokenCall struct {
	s                           *Service
	packageName                 string
	decodeintegritytokenrequest *DecodeIntegrityTokenRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// DecodeIntegrityToken: Decodes the integrity token and returns the token
// payload.
//
//   - packageName: Package name of the app the attached integrity token belongs
//     to.
func (r *V1Service) DecodeIntegrityToken(packageName string, decodeintegritytokenrequest *DecodeIntegrityTokenRequest) *V1DecodeIntegrityTokenCall {
	c := &V1DecodeIntegrityTokenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageName = packageName
	c.decodeintegritytokenrequest = decodeintegritytokenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *V1DecodeIntegrityTokenCall) Fields(s ...googleapi.Field) *V1DecodeIntegrityTokenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *V1DecodeIntegrityTokenCall) Context(ctx context.Context) *V1DecodeIntegrityTokenCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *V1DecodeIntegrityTokenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1DecodeIntegrityTokenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.decodeintegritytokenrequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+packageName}:decodeIntegrityToken")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"packageName": c.packageName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "playintegrity.decodeIntegrityToken" call.
// Any non-2xx status code is an error. Response headers are in either
// *DecodeIntegrityTokenResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *V1DecodeIntegrityTokenCall) Do(opts ...googleapi.CallOption) (*DecodeIntegrityTokenResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &DecodeIntegrityTokenResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}
