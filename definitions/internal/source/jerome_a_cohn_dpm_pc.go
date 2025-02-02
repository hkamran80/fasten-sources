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

// https://fhir-myrecord.cerner.com/r4/11c92883-3ba8-4054-8136-467108df0876/metadata
func GetSourceJeromeACohnDpmPc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/11c92883-3ba8-4054-8136-467108df0876/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/11c92883-3ba8-4054-8136-467108df0876/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/11c92883-3ba8-4054-8136-467108df0876"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/11c92883-3ba8-4054-8136-467108df0876"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeJeromeACohnDpmPc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Jerome A. Cohn, DPM PC"
	sourceDef.SourceType = pkg.SourceTypeJeromeACohnDpmPc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.healthgrades.com/physician/dr-jerome-cohn-2vpr2"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
