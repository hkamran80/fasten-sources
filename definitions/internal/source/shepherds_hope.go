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

// https://fhir-myrecord.cerner.com/r4/d60ebf1e-4f43-4ee3-aa8d-81eddc9a9b5d/metadata
func GetSourceShepherdsHope(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/d60ebf1e-4f43-4ee3-aa8d-81eddc9a9b5d/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/d60ebf1e-4f43-4ee3-aa8d-81eddc9a9b5d/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/d60ebf1e-4f43-4ee3-aa8d-81eddc9a9b5d"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/d60ebf1e-4f43-4ee3-aa8d-81eddc9a9b5d"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeShepherdsHope]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Shepherd's Hope"
	sourceDef.SourceType = pkg.SourceTypeShepherdsHope
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.shepherdshope.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
