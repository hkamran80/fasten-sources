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

// https://fhir4.healow.com/fhir/r4/AIHGCD/metadata
func GetSourcePerezOrthopedicspllcDbaSouthTexasOrthopedicSpecialistsVictoriaTx(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/AIHGCD"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/AIHGCD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/AIHGCD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePerezOrthopedicspllcDbaSouthTexasOrthopedicSpecialistsVictoriaTx]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Perez Orthopedicspllc Dba South Texas Orthopedic Specialists Victoria TX"
	sourceDef.SourceType = pkg.SourceTypePerezOrthopedicspllcDbaSouthTexasOrthopedicSpecialistsVictoriaTx
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
