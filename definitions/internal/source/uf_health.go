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

// https://epicsoap.shands.ufl.edu/FHIR/api/FHIR/R4/metadata
func GetSourceUfHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epicsoap.shands.ufl.edu/FHIR/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicsoap.shands.ufl.edu/FHIR/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epicsoap.shands.ufl.edu/FHIR/oauth2/register"

	sourceDef.Audience = "https://epicsoap.shands.ufl.edu/FHIR/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicsoap.shands.ufl.edu/FHIR/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUfHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "UF Health"
	sourceDef.SourceType = pkg.SourceTypeUfHealth
	sourceDef.Category = []string{"282NC0060X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1033663752"}}
	sourceDef.PatientAccessUrl = "https://ufhealth.org/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
