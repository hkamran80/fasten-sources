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

// https://fhir-myrecord.cerner.com/r4/56a037e8-3411-4ae3-badc-5241023ea3af/metadata
func GetSourceJamesRMilneDoPa(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/56a037e8-3411-4ae3-badc-5241023ea3af/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/56a037e8-3411-4ae3-badc-5241023ea3af/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/56a037e8-3411-4ae3-badc-5241023ea3af"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/56a037e8-3411-4ae3-badc-5241023ea3af"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeJamesRMilneDoPa]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "James R. Milne, DO, PA"
	sourceDef.SourceType = pkg.SourceTypeJamesRMilneDoPa
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "james-r-milne-do-pa.jpg"
	sourceDef.PatientAccessUrl = "https://www.yellowpages.com/oakland-park-fl/mip/dr-james-r-milne-do-496131229"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
