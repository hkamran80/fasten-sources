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

// https://epcapp.mhd.com/arr-prd-fhir/api/FHIR/R4/metadata
func GetSourceMethodistHealthSystem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epcapp.mhd.com/arr-prd-fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epcapp.mhd.com/arr-prd-fhir/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epcapp.mhd.com/arr-prd-fhir/oauth2/register"

	sourceDef.Audience = "https://epcapp.mhd.com/arr-prd-fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epcapp.mhd.com/arr-prd-fhir/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMethodistHealthSystem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Methodist Health System"
	sourceDef.SourceType = pkg.SourceTypeMethodistHealthSystem
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
