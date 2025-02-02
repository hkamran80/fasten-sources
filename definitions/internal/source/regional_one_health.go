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

// https://fhir-myrecord.cerner.com/r4/ab796338-ac9c-40c5-863d-ce75caaa302c/metadata
func GetSourceRegionalOneHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ab796338-ac9c-40c5-863d-ce75caaa302c/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ab796338-ac9c-40c5-863d-ce75caaa302c/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ab796338-ac9c-40c5-863d-ce75caaa302c"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ab796338-ac9c-40c5-863d-ce75caaa302c"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRegionalOneHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Regional One Health"
	sourceDef.SourceType = pkg.SourceTypeRegionalOneHealth
	sourceDef.Category = []string{"282N00000X", "333600000X", "3336C0002X", "3336C0003X", "3336S0011X"}
	sourceDef.Aliases = []string{"REGIONAL ONE HEALTH"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1144213117", "1710959614"}}
	sourceDef.PatientAccessUrl = "https://www.regionalonehealth.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
