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

// https://cookicfg.cookchildrens.org/CookFHIR/api/FHIR/R4/metadata
func GetSourceCookChildrensHealthCareSystem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://cookicfg.cookchildrens.org/CookFHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://cookicfg.cookchildrens.org/CookFHIR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://cookicfg.cookchildrens.org/CookFHIR/oauth2/register"

	sourceDef.Audience = "https://cookicfg.cookchildrens.org/CookFHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://cookicfg.cookchildrens.org/CookFHIR/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCookChildrensHealthCareSystem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Cook Children’s Health Care System"
	sourceDef.SourceType = pkg.SourceTypeCookChildrensHealthCareSystem
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
