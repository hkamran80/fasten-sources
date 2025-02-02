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
func GetSourceCanyonlandsCommunityHealthCare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCanyonlandsCommunityHealthCare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Canyonlands Community Health Care"
	sourceDef.SourceType = pkg.SourceTypeCanyonlandsCommunityHealthCare
	sourceDef.Category = []string{"1223G0001X", "207Q00000X", "207R00000X", "207V00000X", "261QC1500X", "261QF0400X", "363AM0700X", "363AS0400X", "363L00000X"}
	sourceDef.Aliases = []string{"LAKE POWELL MEDICAL CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1134568660", "1144755182", "1154380392", "1417218876", "1568430676", "1609966860", "1619945904", "1669441564", "1780653741", "1861558595"}}
	sourceDef.PatientAccessUrl = "https://canyonlandschc.org/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
