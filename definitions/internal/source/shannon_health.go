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

// https://epicrpx.shannonhealth.org/FHIR_ARR/api/FHIR/R4/metadata
func GetSourceShannonHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epicrpx.shannonhealth.org/FHIR_ARR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicrpx.shannonhealth.org/FHIR_ARR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epicrpx.shannonhealth.org/FHIR_ARR/oauth2/register"

	sourceDef.Audience = "https://epicrpx.shannonhealth.org/FHIR_ARR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicrpx.shannonhealth.org/FHIR_ARR/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeShannonHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Shannon Health"
	sourceDef.SourceType = pkg.SourceTypeShannonHealth
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
