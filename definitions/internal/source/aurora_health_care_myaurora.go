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

// https://EpicFHIR.aurora.org/FHIR/MYAURORA/api/FHIR/R4/.well-known/smart-configuration
// https://EpicFHIR.aurora.org/FHIR/MYAURORA/api/FHIR/R4/metadata
func GetSourceAuroraHealthCareMyaurora(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://EpicFHIR.aurora.org/FHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://EpicFHIR.aurora.org/FHIR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://EpicFHIR.aurora.org/FHIR/oauth2/register"

	sourceDef.Audience = "https://EpicFHIR.aurora.org/FHIR/MYAURORA/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://EpicFHIR.aurora.org/FHIR/MYAURORA/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAuroraHealthCareMyaurora]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Aurora Health Care - myAurora"
	sourceDef.SourceType = pkg.SourceTypeAuroraHealthCareMyaurora
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
