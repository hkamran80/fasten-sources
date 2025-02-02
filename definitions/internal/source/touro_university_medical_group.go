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

// https://fhir-myrecord.cerner.com/r4/e0d732aa-1d5c-42bf-b672-8fabe3feb437/metadata
func GetSourceTouroUniversityMedicalGroup(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/e0d732aa-1d5c-42bf-b672-8fabe3feb437/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/e0d732aa-1d5c-42bf-b672-8fabe3feb437/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/e0d732aa-1d5c-42bf-b672-8fabe3feb437"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/e0d732aa-1d5c-42bf-b672-8fabe3feb437"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTouroUniversityMedicalGroup]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Touro University Medical Group"
	sourceDef.SourceType = pkg.SourceTypeTouroUniversityMedicalGroup
	sourceDef.Category = []string{"207R00000X", "2084N0400X", "2084N0600X", "2084P0800X", "2084P0804X", "261QM1300X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1215573613"}}
	sourceDef.PatientAccessUrl = "https://www.mytumg.org/contact/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
