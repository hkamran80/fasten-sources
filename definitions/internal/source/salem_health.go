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

// https://prd.salemhealth.org/fhir/api/FHIR/R4/metadata
func GetSourceSalemHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://prd.salemhealth.org/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://prd.salemhealth.org/fhir/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://prd.salemhealth.org/fhir/oauth2/register"

	sourceDef.Audience = "https://prd.salemhealth.org/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://prd.salemhealth.org/fhir/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSalemHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Salem Health"
	sourceDef.SourceType = pkg.SourceTypeSalemHealth
	sourceDef.Category = []string{"207Q00000X", "251E00000X", "261QP2300X", "261QR0200X", "273R00000X", "273Y00000X", "282N00000X", "291U00000X", "3336C0003X"}
	sourceDef.Aliases = []string{"RIVER ROAD PHARMACY", "SALEM HOSPITAL", "SSALEM HOSPITAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1013505288", "1164099248", "1265431829", "1568059707", "1578112066", "1689128308", "1952952079"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
