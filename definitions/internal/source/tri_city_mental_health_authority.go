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

// https://fhir-myrecord.cerner.com/r4/a64573bd-48d1-453d-bc0b-03ac15ad915e/metadata
func GetSourceTriCityMentalHealthAuthority(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a64573bd-48d1-453d-bc0b-03ac15ad915e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a64573bd-48d1-453d-bc0b-03ac15ad915e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a64573bd-48d1-453d-bc0b-03ac15ad915e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a64573bd-48d1-453d-bc0b-03ac15ad915e"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTriCityMentalHealthAuthority]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Tri-City Mental Health Authority"
	sourceDef.SourceType = pkg.SourceTypeTriCityMentalHealthAuthority
	sourceDef.Category = []string{"251S00000X", "261QM0801X"}
	sourceDef.Aliases = []string{"TRI CITY MENTAL HEALTH AUTHORITY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1013540707", "1093366650", "1518287028", "1568095131", "1598398166", "1962053611", "1992356646"}}
	sourceDef.PatientAccessUrl = "https://www.facebook.com/TriCityMHS/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
