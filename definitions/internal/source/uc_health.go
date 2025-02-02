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

// https://epic-soap.uchealth.com/FHIRProxy/api/FHIR/R4/metadata
func GetSourceUcHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epic-soap.uchealth.com/FHIRProxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epic-soap.uchealth.com/FHIRProxy/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epic-soap.uchealth.com/FHIRProxy/oauth2/register"

	sourceDef.Audience = "https://epic-soap.uchealth.com/FHIRProxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epic-soap.uchealth.com/FHIRProxy/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUcHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "UC Health"
	sourceDef.SourceType = pkg.SourceTypeUcHealth
	sourceDef.Category = []string{"261QP2300X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1013297340"}}
	sourceDef.PatientAccessUrl = "https://www.uchealth.com/en"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
