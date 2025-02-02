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

// https://fhir4.healow.com/fhir/r4/GBIEAD/metadata
func GetSourceBlueCoastCardiology(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/GBIEAD"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/GBIEAD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/GBIEAD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBlueCoastCardiology]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Blue Coast Cardiology"
	sourceDef.SourceType = pkg.SourceTypeBlueCoastCardiology
	sourceDef.Category = []string{"207RC0000X"}
	sourceDef.Aliases = []string{"BLUE COAST CARDIOLOGY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1649526518"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
