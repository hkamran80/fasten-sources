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
func GetSourceGoodmanCampbellBrainAndSpine(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGoodmanCampbellBrainAndSpine]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Goodman Campbell Brain And Spine"
	sourceDef.SourceType = pkg.SourceTypeGoodmanCampbellBrainAndSpine
	sourceDef.Category = []string{"163W00000X", "207LP2900X", "207T00000X", "208100000X", "2085R0204X", "363A00000X", "363L00000X"}
	sourceDef.Aliases = []string{"GOODMAN CAMPBELL BRAIN AND SPINE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1275580318"}}
	sourceDef.BrandLogo = "goodman-campbell-brain-and-spine.svg"
	sourceDef.PatientAccessUrl = "https://www.goodmancampbell.com/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
