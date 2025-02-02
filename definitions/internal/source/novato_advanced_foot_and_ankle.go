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

// https://fhir-myrecord.cerner.com/r4/58af3cfb-25bc-40a1-b39f-a39b55dc2d85/metadata
func GetSourceNovatoAdvancedFootAndAnkle(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/58af3cfb-25bc-40a1-b39f-a39b55dc2d85/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/58af3cfb-25bc-40a1-b39f-a39b55dc2d85/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/58af3cfb-25bc-40a1-b39f-a39b55dc2d85"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/58af3cfb-25bc-40a1-b39f-a39b55dc2d85"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNovatoAdvancedFootAndAnkle]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Novato Advanced Foot & Ankle"
	sourceDef.SourceType = pkg.SourceTypeNovatoAdvancedFootAndAnkle
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://novatoadvancedfootandankle.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
