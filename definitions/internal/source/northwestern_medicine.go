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

// https://nmepicproxy.nm.org/FHIR-PRD/api/FHIR/R4/metadata
func GetSourceNorthwesternMedicine(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://nmepicproxy.nm.org/FHIR-PRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://nmepicproxy.nm.org/FHIR-PRD/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://nmepicproxy.nm.org/FHIR-PRD/oauth2/register"

	sourceDef.Audience = "https://nmepicproxy.nm.org/FHIR-PRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://nmepicproxy.nm.org/FHIR-PRD/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNorthwesternMedicine]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Northwestern Medicine"
	sourceDef.SourceType = pkg.SourceTypeNorthwesternMedicine
	sourceDef.Category = []string{"261Q00000X", "261QM0801X", "282N00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1306417076", "1649668401", "1730541699"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
