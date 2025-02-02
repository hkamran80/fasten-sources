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

// https://psacesoap.whhs.com/interconnect-fhir-prd/api/FHIR/R4/metadata
func GetSourceWashingtonHospitalHealthcareSystem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://psacesoap.whhs.com/interconnect-fhir-prd/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://psacesoap.whhs.com/interconnect-fhir-prd/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://psacesoap.whhs.com/interconnect-fhir-prd/oauth2/register"

	sourceDef.Audience = "https://psacesoap.whhs.com/interconnect-fhir-prd/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://psacesoap.whhs.com/interconnect-fhir-prd/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWashingtonHospitalHealthcareSystem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Washington Hospital Healthcare System"
	sourceDef.SourceType = pkg.SourceTypeWashingtonHospitalHealthcareSystem
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
