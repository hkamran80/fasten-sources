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

// https://fhir-myrecord.cerner.com/r4/51f96a70-43c0-40c0-931a-19c1f8bc3766/metadata
func GetSourceMichaelFEsberDpmPc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/51f96a70-43c0-40c0-931a-19c1f8bc3766/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/51f96a70-43c0-40c0-931a-19c1f8bc3766/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/51f96a70-43c0-40c0-931a-19c1f8bc3766"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/51f96a70-43c0-40c0-931a-19c1f8bc3766"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMichaelFEsberDpmPc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Michael F. Esber, DPM PC"
	sourceDef.SourceType = pkg.SourceTypeMichaelFEsberDpmPc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "michael-f-esber-dpm-pc.svg"
	sourceDef.PatientAccessUrl = "https://doctors.bannerhealth.com/provider/Michael+Esber/447216"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
