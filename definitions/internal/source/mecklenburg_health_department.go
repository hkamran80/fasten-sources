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

// https://fhir-myrecord.cerner.com/r4/5213ec22-3df5-4b49-ab92-487589c502a5/metadata
func GetSourceMecklenburgHealthDepartment(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/5213ec22-3df5-4b49-ab92-487589c502a5/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/5213ec22-3df5-4b49-ab92-487589c502a5/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/5213ec22-3df5-4b49-ab92-487589c502a5"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/5213ec22-3df5-4b49-ab92-487589c502a5"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMecklenburgHealthDepartment]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Mecklenburg Health Department"
	sourceDef.SourceType = pkg.SourceTypeMecklenburgHealthDepartment
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://m.facebook.com/profile.php?id=147448061979552"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
