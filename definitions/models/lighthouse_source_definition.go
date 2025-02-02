package models

import "github.com/fastenhealth/fasten-sources/pkg"

// LighthouseSource
//https://vteapif1.aetna.com/fhirdemo/v1/fhirserver_auth/oauth2/authorize?
//response_type=code&
//client_id=5c47935b-29a7-4346-a01b-649a11d94dc5&
//scope=launch%2Fpatient%20patient%2FObservation.rs%20patient%2FPatient.rs%20offline_access&
//redirect_uri=https://sandbox.fastenhealth.com/api/callback/aetna&
//aud=https://vteapif1.aetna.com/fhirdemo&
//code_challenge=XXXXXX&
//code_challenge_method=XXXXX
//
// Similar in functionality to https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html#example-request
// /apis/fhir/.well-known/smart-configuration
type LighthouseSourceDefinition struct {

	// Smart-On-FHIR configuration
	// https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html#example-request

	//oauth endpoints
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	IntrospectionEndpoint string `json:"introspection_endpoint"`
	RegistrationEndpoint  string `json:"registration_endpoint"` //optional - required when Dynamic Client Registration mode is set

	Scopes                        []string `json:"scopes_supported"`
	Issuer                        string   `json:"issuer"`
	GrantTypesSupported           []string `json:"grant_types_supported"`
	ResponseType                  []string `json:"response_types_supported"`
	ResponseModesSupported        []string `json:"response_modes_supported"`
	Audience                      string   `json:"aud"`                              //optional - required for some providers
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"` // If populated: PKCE is supported (can be used with Confidential true or false)

	//Fasten custom configuration
	UserInfoEndpoint   string `json:"userinfo_endpoint"`     //optional - supported by some providers, not others.
	ApiEndpointBaseUrl string `json:"api_endpoint_base_url"` //api endpoint we'll communicate with after authentication
	ClientId           string `json:"client_id"`
	RedirectUri        string `json:"redirect_uri"` //lighthouse url the provider will redirect to (registered with App)

	Confidential                  bool   `json:"confidential"`                     //if enabled, requires client_secret to authenticate with provider (PKCE)
	DynamicClientRegistrationMode string `json:"dynamic_client_registration_mode"` //if enabled, will dynamically register client with provider (https://oauth.net/2/dynamic-client-registration/)
	CORSRelayRequired             bool   `json:"cors_relay_required"`              //if true, requires CORS proxy/relay, as provider does not return proper response to CORS preflight
	SecretKeyPrefix               string `json:"-"`                                //the secret key prefix to use, if empty (default) will use the sourceType value

	//Display information
	PlatformType pkg.SourceType      `json:"platform_type"`
	Display      string              `json:"display"` //display name of the source
	SourceType   pkg.SourceType      `json:"source_type"`
	Category     []string            `json:"category"`
	Aliases      []string            `json:"aliases"`
	Identifiers  map[string][]string `json:"identifiers"`

	Hidden                   bool   `json:"hidden"`
	BrandLogo                string `json:"brand_logo,omitempty"` //if set, will override the default logo filename (sourceType.png). Use for non-png extensions
	PatientAccessUrl         string `json:"patient_access_url,omitempty"`
	PatientAccessDescription string `json:"patient_access_description,omitempty"`
}
