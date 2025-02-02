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

// https://fhir-myrecord.cerner.com/r4/2c834eee-5231-4dd1-ba8d-28931b99c0cc/metadata
func GetSourceTheFootCareGroupPc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/2c834eee-5231-4dd1-ba8d-28931b99c0cc/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/2c834eee-5231-4dd1-ba8d-28931b99c0cc/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/2c834eee-5231-4dd1-ba8d-28931b99c0cc"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/2c834eee-5231-4dd1-ba8d-28931b99c0cc"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTheFootCareGroupPc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "The Foot Care Group, PC."
	sourceDef.SourceType = pkg.SourceTypeTheFootCareGroupPc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.yourfootdoctor.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
