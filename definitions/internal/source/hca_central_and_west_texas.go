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

// https://stdavidsfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4/metadata
func GetSourceHcaCentralAndWestTexas(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://stdavidsfhirprd.app.medcity.net/fhir-proxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://stdavidsfhirprd.app.medcity.net/fhir-proxy/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://stdavidsfhirprd.app.medcity.net/fhir-proxy/oauth2/register"

	sourceDef.Audience = "https://stdavidsfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://stdavidsfhirprd.app.medcity.net/fhir-proxy/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHcaCentralAndWestTexas]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "HCA Central and West Texas"
	sourceDef.SourceType = pkg.SourceTypeHcaCentralAndWestTexas
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "hca-mountain.png"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
