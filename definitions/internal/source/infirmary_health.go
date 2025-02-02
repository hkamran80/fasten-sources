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

// https://ssproxyprod.infirmaryhealth.org/epicFHIR/api/FHIR/R4/metadata
func GetSourceInfirmaryHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://ssproxyprod.infirmaryhealth.org/epicFHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://ssproxyprod.infirmaryhealth.org/epicFHIR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://ssproxyprod.infirmaryhealth.org/epicFHIR/oauth2/register"

	sourceDef.Audience = "https://ssproxyprod.infirmaryhealth.org/epicFHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://ssproxyprod.infirmaryhealth.org/epicFHIR/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeInfirmaryHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Infirmary Health"
	sourceDef.SourceType = pkg.SourceTypeInfirmaryHealth
	sourceDef.Category = []string{"282E00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1548710197"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
