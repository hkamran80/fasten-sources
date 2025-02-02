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
func GetSourceCrossroadsHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCrossroadsHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Crossroads Health"
	sourceDef.SourceType = pkg.SourceTypeCrossroadsHealth
	sourceDef.Category = []string{"101YA0400X", "207Q00000X", "251S00000X", "261QM0850X", "261QM0855X", "3336C0003X"}
	sourceDef.Aliases = []string{"CROSSROADS LAKE COUNTY ADOLESCENT COUNSELING", "CROSSROADS LAKE COUNTY ADOLESCENT COUNSELING SERVICE", "CROSSROADS LAKE COUNTY ADOLESCENT COUNSELING SERVICES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1093307654", "1154680627", "1356303911", "1437733169", "1518226075"}}
	sourceDef.PatientAccessUrl = "https://crossroadshealth.org/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
