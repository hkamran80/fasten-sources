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

// https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4/metadata
func GetSourceCommunicareHealthCenters(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCommunicareHealthCenters]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Communicare Health Centers"
	sourceDef.SourceType = pkg.SourceTypeCommunicareHealthCenters
	sourceDef.Category = []string{"251S00000X", "261QF0400X", "261QM0855X", "261QM1300X"}
	sourceDef.Aliases = []string{"COMMUNICARE FAMILY WELLNESS CENTER - UNITY BUILDING", "COMMUNICARE HEALTH CENTERS", "HANSEN FAMILY HEALTH CENTERS", "VIDA FAMILY HEALTH CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003228750", "1063067502", "1073006607", "1104207208", "1114059474", "1124246947", "1215561808", "1275284309", "1356569172", "1366091191", "1386882231", "1427357458", "1447478268", "1528608890", "1558469064", "1669115630", "1710087515", "1750701454", "1881330579", "1952690463"}}
	sourceDef.PatientAccessUrl = "https://communicaresa.org/wp-content/uploads/2021/07/CC_Logo_White_2021-e1626108594478.png"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
