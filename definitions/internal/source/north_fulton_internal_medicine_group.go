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

// https://fhir4.eclinicalworks.com/fhir/r4/HHDCAA/metadata
func GetSourceNorthFultonInternalMedicineGroup(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Audience = "https://fhir4.eclinicalworks.com/fhir/r4/HHDCAA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.eclinicalworks.com/fhir/r4/HHDCAA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNorthFultonInternalMedicineGroup]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "North Fulton Internal Medicine Group"
	sourceDef.SourceType = pkg.SourceTypeNorthFultonInternalMedicineGroup
	sourceDef.Category = []string{"261QP2300X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1396708087"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
