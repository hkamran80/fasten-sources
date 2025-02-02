// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46/metadata
func GetSourceTriCountyMentalHealthServices(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/b39fa050-f9bd-471d-b705-2049f60e1d46/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/b39fa050-f9bd-471d-b705-2049f60e1d46/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/b39fa050-f9bd-471d-b705-2049f60e1d46"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTriCountyMentalHealthServices]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Tri-County Mental Health Services"
	sourceDef.SourceType = pkg.SourceTypeTriCountyMentalHealthServices
	sourceDef.Category = []string{"101YM0800X", "251B00000X", "251C00000X", "261QM0850X", "320800000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1013243211", "1215979158", "1386970580", "1669708863", "1689921629"}}
	sourceDef.BrandLogo = "tri-county-mental-health-services.jpg"
	sourceDef.PatientAccessUrl = "https://www.tcmhs.org/contact"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
