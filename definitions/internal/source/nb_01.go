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

// https://fhir.nb0.hos.allscriptscloud.com/open/metadata
func GetSourceNb01(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/44e67798-95eb-4d45-aabf-a22a062a0e48/connect/authorize"
	sourceDef.TokenEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/44e67798-95eb-4d45-aabf-a22a062a0e48/connect/token"

	sourceDef.Audience = "https://fhir.nb0.hos.allscriptscloud.com/open"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nb0.hos.allscriptscloud.com/open"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNb01]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "NB0"
	sourceDef.SourceType = pkg.SourceTypeNb01
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
