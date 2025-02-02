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

// https://fhir-myrecord.cerner.com/r4/c0c316ba-6c8e-4f8d-898b-bb21c58aa05f/metadata
func GetSourceSmithCountyMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/c0c316ba-6c8e-4f8d-898b-bb21c58aa05f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/c0c316ba-6c8e-4f8d-898b-bb21c58aa05f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/c0c316ba-6c8e-4f8d-898b-bb21c58aa05f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/c0c316ba-6c8e-4f8d-898b-bb21c58aa05f"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSmithCountyMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Smith County Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeSmithCountyMemorialHospital
	sourceDef.Category = []string{"275N00000X", "282NC0060X"}
	sourceDef.Aliases = []string{"SMITH COUNTY MEMORIAL HOSPITAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1700985116", "1760498141", "1851334734"}}
	sourceDef.PatientAccessUrl = "https://www.scmhks.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
