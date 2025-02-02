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

// https://fhir.ep0.hos.allscriptscloud.com/Fhir/metadata
func GetSourceSheppardPrattHealthSystem4(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.ep0.hos.allscriptscloud.com/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://fhir.ep0.hos.allscriptscloud.com/authorization/connect/token"

	sourceDef.Audience = "https://fhir.ep0.hos.allscriptscloud.com/Fhir"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.ep0.hos.allscriptscloud.com/Fhir"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSheppardPrattHealthSystem4]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Sheppard Pratt Health System"
	sourceDef.SourceType = pkg.SourceTypeSheppardPrattHealthSystem4
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
