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

// https://fhir4.healow.com/fhir/r4/HBCCAA/metadata
func GetSourceOrthopaedicInstituteOfOhio(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/HBCCAA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/HBCCAA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/HBCCAA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeOrthopaedicInstituteOfOhio]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Orthopaedic Institute of Ohio"
	sourceDef.SourceType = pkg.SourceTypeOrthopaedicInstituteOfOhio
	sourceDef.Category = []string{"207X00000X", "208100000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1073783916", "1083870026", "1114048444", "1205956299", "1205957537", "1497909568", "1952555021"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
