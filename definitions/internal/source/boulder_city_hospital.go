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

// https://fhir-myrecord.cerner.com/r4/fab57722-243a-4069-b32a-bcc74c7779f7/metadata
func GetSourceBoulderCityHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/fab57722-243a-4069-b32a-bcc74c7779f7/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/fab57722-243a-4069-b32a-bcc74c7779f7/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/fab57722-243a-4069-b32a-bcc74c7779f7"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/fab57722-243a-4069-b32a-bcc74c7779f7"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBoulderCityHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Boulder City Hospital"
	sourceDef.SourceType = pkg.SourceTypeBoulderCityHospital
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.bchcares.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
