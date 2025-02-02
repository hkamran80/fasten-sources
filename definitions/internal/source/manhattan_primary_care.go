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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067932/.well-known/smart-configuration
// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067932/metadata
func GetSourceManhattanPrimaryCare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/5a50bf13-90ed-47a8-9cf8-abcf012040cf/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/5a50bf13-90ed-47a8-9cf8-abcf012040cf/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067932"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067932"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeManhattanPrimaryCare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Manhattan Primary Care"
	sourceDef.SourceType = pkg.SourceTypeManhattanPrimaryCare
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
