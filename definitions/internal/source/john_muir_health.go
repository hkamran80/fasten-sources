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

// https://fhir.johnmuirhealth.com/fhir-prd/api/FHIR/R4/metadata
func GetSourceJohnMuirHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.johnmuirhealth.com/fhir-prd/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://fhir.johnmuirhealth.com/fhir-prd/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://fhir.johnmuirhealth.com/fhir-prd/oauth2/register"

	sourceDef.Audience = "https://fhir.johnmuirhealth.com/fhir-prd/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.johnmuirhealth.com/fhir-prd/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeJohnMuirHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "John Muir Health"
	sourceDef.SourceType = pkg.SourceTypeJohnMuirHealth
	sourceDef.Category = []string{"204E00000X", "207RC0000X", "207RG0100X", "207RP1001X", "207W00000X", "207X00000X", "207Y00000X", "2084N0400X", "208600000X", "2086S0122X", "2086S0129X", "208800000X", "273Y00000X", "282N00000X", "291U00000X", "333600000X", "3336C0003X", "363A00000X", "363L00000X"}
	sourceDef.Aliases = []string{"JOHN MUIR HEALTH, WALNUT CREEK MEDICAL CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1164799391", "1174890305", "1275753493", "1336416569", "1447269303", "1619049491", "1699099937", "1740215219", "1760602726", "1801163035"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
