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

// https://fhir4.healow.com/fhir/r4/DCCHBD/metadata
func GetSourceAdvancedPainManagement1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/DCCHBD"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/DCCHBD"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/DCCHBD"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAdvancedPainManagement1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Advanced Pain Management"
	sourceDef.SourceType = pkg.SourceTypeAdvancedPainManagement1
	sourceDef.Category = []string{"174400000X", "261Q00000X", "261QA1903X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003062282", "1043612450", "1225284409", "1477982163", "1588715510", "1639325814", "1699768713", "1831345016", "1841446028", "1972759470"}}
	sourceDef.BrandLogo = "advanced-pain-management.png"
	sourceDef.PatientAccessUrl = "https://apmhealth.com/"
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
