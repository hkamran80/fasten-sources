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
func GetSourceEyePhysiciansOfSussexCounty(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEyePhysiciansOfSussexCounty]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Eye Physicians Of Sussex County"
	sourceDef.SourceType = pkg.SourceTypeEyePhysiciansOfSussexCounty
	sourceDef.Category = []string{"207W00000X", "261QA1903X"}
	sourceDef.Aliases = []string{"EYE PHYSICIANS OF SUSSEX COUNTY SURGICAL SUITE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1801993217", "1861429672"}}
	sourceDef.PatientAccessUrl = "https://www.eyephysiciansofsussexcounty.com/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
