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

// https://khnarr.ketthealth.com/FHIR-PROD/api/FHIR/R4/metadata
func GetSourceKetteringHealthNetwork(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://KHNARR.KETTHEALTH.COM/FHIR-PROD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://KHNARR.KETTHEALTH.COM/FHIR-PROD/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://KHNARR.KETTHEALTH.COM/FHIR-PROD/oauth2/register"

	sourceDef.Audience = "https://khnarr.ketthealth.com/FHIR-PROD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://khnarr.ketthealth.com/FHIR-PROD/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeKetteringHealthNetwork]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Kettering Health Network"
	sourceDef.SourceType = pkg.SourceTypeKetteringHealthNetwork
	sourceDef.Category = []string{"282N00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1356669840"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
