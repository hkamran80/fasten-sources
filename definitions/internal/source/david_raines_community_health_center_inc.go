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
func GetSourceDavidRainesCommunityHealthCenterInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDavidRainesCommunityHealthCenterInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "David Raines Community Health Center Inc"
	sourceDef.SourceType = pkg.SourceTypeDavidRainesCommunityHealthCenterInc
	sourceDef.Category = []string{"261QF0400X", "333600000X", "3336C0002X", "3336C0003X", "3336M0003X"}
	sourceDef.Aliases = []string{"DAVID RAINES COMMUNITY HEALTH CENTER - MOBILE DENTAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1144705930", "1265438998", "1639489990", "1861415721"}}
	sourceDef.BrandLogo = "david-raines-community-health-center-inc.jpeg"
	sourceDef.PatientAccessUrl = "https://www.davidraineschc.org/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
