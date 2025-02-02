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

// https://fhir-myrecord.cerner.com/r4/d13cd954-ca68-4a08-a95d-1ae54e7487f0/metadata
func GetSourcePami2(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/d13cd954-ca68-4a08-a95d-1ae54e7487f0/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/d13cd954-ca68-4a08-a95d-1ae54e7487f0/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/d13cd954-ca68-4a08-a95d-1ae54e7487f0"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/d13cd954-ca68-4a08-a95d-1ae54e7487f0"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePami2]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "PAMI"
	sourceDef.SourceType = pkg.SourceTypePami2
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "pami.jpg"
	sourceDef.PatientAccessUrl = "http://www.epami.net/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
