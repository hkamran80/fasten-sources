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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10036320/metadata
func GetSourceLaureateMedicalGroup(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/51AFF0CC-3B27-4590-B895-A35400E5F189/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/51AFF0CC-3B27-4590-B895-A35400E5F189/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10036320"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10036320"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeLaureateMedicalGroup]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Laureate Medical Group"
	sourceDef.SourceType = pkg.SourceTypeLaureateMedicalGroup
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
