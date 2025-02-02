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

// https://fhir4.healow.com/fhir/r4/FEHIBA/metadata
func GetSourceFamilyMedicalAssociates1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEclinicalworks(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://oauthserver.eclinicalworks.com/oauth/oauth2/token"

	sourceDef.Issuer = "https://fhir4.healow.com/fhir/r4/FEHIBA"
	sourceDef.Audience = "https://fhir4.healow.com/fhir/r4/FEHIBA"

	sourceDef.ApiEndpointBaseUrl = "https://fhir4.healow.com/fhir/r4/FEHIBA"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFamilyMedicalAssociates1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEclinicalworks))

	sourceDef.Display = "Family Medical Associates"
	sourceDef.SourceType = pkg.SourceTypeFamilyMedicalAssociates1
	sourceDef.Category = []string{"207Q00000X", "207R00000X", "208000000X", "208D00000X", "261QP2300X"}
	sourceDef.Aliases = []string{"COMMUNITY MEDICAL ASSOC OF B.C.H.", "FAMILY MEDICAL ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1184730020", "1194734855", "1205871548", "1497302871"}}
	sourceDef.SecretKeyPrefix = "eclinicalworks"

	return sourceDef, err
}
