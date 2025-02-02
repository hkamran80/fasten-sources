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

// https://fhir-myrecord.cerner.com/r4/4211ccc0-df3c-4100-a077-f6b59ef3ddba/metadata
func GetSourceCityOfChandler(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/4211ccc0-df3c-4100-a077-f6b59ef3ddba/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/4211ccc0-df3c-4100-a077-f6b59ef3ddba/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/4211ccc0-df3c-4100-a077-f6b59ef3ddba"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/4211ccc0-df3c-4100-a077-f6b59ef3ddba"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeCityOfChandler]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "City of Chandler"
	sourceDef.SourceType = pkg.SourceTypeCityOfChandler
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "city-of-chandler.svg"
	sourceDef.PatientAccessUrl = "https://www.chandleraz.gov/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
