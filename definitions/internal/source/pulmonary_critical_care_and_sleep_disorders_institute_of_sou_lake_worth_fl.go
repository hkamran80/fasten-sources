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

// https://fhir4.healow.com/fhir/r4/IEEEBD/metadata
func GetSourcePulmonaryCriticalCareAndSleepDisordersInstituteOfSouLakeWorthFl(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/IEEEBD"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/IEEEBD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/IEEEBD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePulmonaryCriticalCareAndSleepDisordersInstituteOfSouLakeWorthFl]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Pulmonary Critical Care and Sleep Disorders Institute of Sou LAKE WORTH FL"
	sourceDef.SourceType = pkg.SourceTypePulmonaryCriticalCareAndSleepDisordersInstituteOfSouLakeWorthFl
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
