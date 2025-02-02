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

// https://fhir-myrecord.cerner.com/r4/3e12b969-e9cf-4559-9980-aa408e074b16/metadata
func GetSourceOsfSaintClareMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/3e12b969-e9cf-4559-9980-aa408e074b16/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/3e12b969-e9cf-4559-9980-aa408e074b16/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/3e12b969-e9cf-4559-9980-aa408e074b16"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/3e12b969-e9cf-4559-9980-aa408e074b16"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeOsfSaintClareMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "OSF Saint Clare Medical Center"
	sourceDef.SourceType = pkg.SourceTypeOsfSaintClareMedicalCenter
	sourceDef.Category = []string{"261Q00000X", "275N00000X", "282NC0060X"}
	sourceDef.Aliases = []string{"OSF SAINT CLARE MEDICAL CENTER", "OSF SAINT CLARE MEDICAL CENTER SWING BED"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1093308439", "1447843891", "1790363505"}}
	sourceDef.PatientAccessUrl = "https://www.osfhealthcare.org/saint-clare/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
