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

// https://fhir-myrecord.cerner.com/r4/d7c7003a-96ae-4246-aaaa-19b18f13056e/metadata
func GetSourceSheridanMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/d7c7003a-96ae-4246-aaaa-19b18f13056e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/d7c7003a-96ae-4246-aaaa-19b18f13056e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/d7c7003a-96ae-4246-aaaa-19b18f13056e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/d7c7003a-96ae-4246-aaaa-19b18f13056e"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSheridanMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Sheridan Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeSheridanMemorialHospital
	sourceDef.Category = []string{"207P00000X", "261QU0200X", "282N00000X"}
	sourceDef.Aliases = []string{"SHERIDAN MEMORIAL HOSPITAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1487707188"}}
	sourceDef.PatientAccessUrl = "https://www.sheridanhospital.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
