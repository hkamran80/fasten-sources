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
func GetSourceSouthwellAmbulatoryInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSouthwellAmbulatoryInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "SOUTHWELL AMBULATORY INC"
	sourceDef.SourceType = pkg.SourceTypeSouthwellAmbulatoryInc
	sourceDef.Category = []string{"207Q00000X", "207R00000X", "207RG0100X", "207RI0200X", "207V00000X", "208000000X", "261QE0800X", "261QM1300X", "363LP2300X"}
	sourceDef.Aliases = []string{"SOUTHWELL GASTROENTEROLOGY", "SOUTHWELL INFECTIOUS DISEASE", "SOUTHWELL INTERNAL MEDICINE", "SOUTHWELL LOWNDES ENDOSCOPY CENTER", "SOUTHWELL OB/GYN", "SOUTHWELL PEDIATRICS VALDOSTA", "SOUTHWELL PRIMARY CARE", "SOUTHWELL VALDOSTA ENDOSCOPY CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1164069472", "1215574926", "1306466206", "1497330625", "1598304974", "1720625296", "1720693047", "1962035360", "1982296364"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}
