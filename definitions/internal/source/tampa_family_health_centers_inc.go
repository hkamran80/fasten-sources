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

// https://api.platform.athenahealth.com/fhir/r4/metadata
func GetSourceTampaFamilyHealthCentersInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTampaFamilyHealthCentersInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "TAMPA FAMILY HEALTH CENTERS INC"
	sourceDef.SourceType = pkg.SourceTypeTampaFamilyHealthCentersInc
	sourceDef.Category = []string{"183500000X", "261QD0000X", "261QF0400X", "333600000X", "3336C0002X"}
	sourceDef.Aliases = []string{"TAMPA COMMUNITY HEALTH CTR.", "TAMPA COMMUNITY HEALTH CTRS", "TAMPA FAMILY HEALTH CENTERS", "TAMPA FAMILY HEALTH CENTERS, INC", "TFHC#23 - PHARMACY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1164754701", "1245608348", "1255675682", "1306988910", "1417325515", "1477604486", "1568503985", "1629340922", "1760881437", "1811281488"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}
