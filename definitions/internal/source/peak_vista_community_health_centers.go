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
func GetSourcePeakVistaCommunityHealthCenters(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePeakVistaCommunityHealthCenters]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Peak Vista Community Health Centers"
	sourceDef.SourceType = pkg.SourceTypePeakVistaCommunityHealthCenters
	sourceDef.Category = []string{"261QD0000X", "261QF0400X"}
	sourceDef.Aliases = []string{"CONVENIENT CARE CENTER AND PEDIATRIC HEALTH CENTER AT ACADEMY", "DENTAL HEALTH CENTER", "DOWNTOWN HEALTH CENTER", "FAMILY HEALTH CENTER AT MITCHELL HIGH SCHOOL", "HEALTH CENTER AT JET WING", "HEALTH CENTER AT UNION", "HEALTH CENTER AT WAHSATCH", "HEATH CENTER AT ACADEMY SUITE 3500", "LOGAN HEALTH CENTER AT MYRON STRATTON", "PEDIATRIC HEALTH CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1033412796", "1124321880", "1396048054", "1437622206", "1497787535", "1639714041", "1689977365", "1841593225", "1962740381", "1972245959"}}
	sourceDef.PatientAccessUrl = "https://www.peakvista.org/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
