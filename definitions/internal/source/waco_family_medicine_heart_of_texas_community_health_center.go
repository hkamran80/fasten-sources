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

// https://haiku.wacofhc.org/FHIR/api/FHIR/R4/metadata
func GetSourceWacoFamilyMedicineHeartOfTexasCommunityHealthCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://haiku.wacofhc.org/FHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://haiku.wacofhc.org/FHIR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://haiku.wacofhc.org/FHIR/oauth2/register"

	sourceDef.Audience = "https://haiku.wacofhc.org/FHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://haiku.wacofhc.org/FHIR/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWacoFamilyMedicineHeartOfTexasCommunityHealthCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Waco Family Medicine (Heart of Texas Community Health Center)"
	sourceDef.SourceType = pkg.SourceTypeWacoFamilyMedicineHeartOfTexasCommunityHealthCenter
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
