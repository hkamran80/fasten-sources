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

// https://fhir-myrecord.cerner.com/r4/a87a863d-8666-48d4-bbb9-1908bfb27a5f/metadata
func GetSourceNorthernArizonaHealthcare1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a87a863d-8666-48d4-bbb9-1908bfb27a5f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a87a863d-8666-48d4-bbb9-1908bfb27a5f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a87a863d-8666-48d4-bbb9-1908bfb27a5f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a87a863d-8666-48d4-bbb9-1908bfb27a5f"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNorthernArizonaHealthcare1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Northern Arizona Healthcare"
	sourceDef.SourceType = pkg.SourceTypeNorthernArizonaHealthcare1
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "northern-arizona-healthcare.svg"
	sourceDef.PatientAccessUrl = "https://www.nahealth.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
