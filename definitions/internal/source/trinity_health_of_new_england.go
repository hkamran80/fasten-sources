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

// https://epicprdext.stfranciscare.org/FhirProxy/api/FHIR/R4/metadata
func GetSourceTrinityHealthOfNewEngland(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epicprdext.stfranciscare.org/FhirProxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicprdext.stfranciscare.org/FhirProxy/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epicprdext.stfranciscare.org/FhirProxy/oauth2/register"

	sourceDef.Audience = "https://epicprdext.stfranciscare.org/FhirProxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicprdext.stfranciscare.org/FhirProxy/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTrinityHealthOfNewEngland]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Trinity Health of New England"
	sourceDef.SourceType = pkg.SourceTypeTrinityHealthOfNewEngland
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
