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

// https://fhir-myrecord.cerner.com/r4/43a4de57-0683-4b0f-912a-33d5300b9c77/metadata
func GetSourceVirginiaLightnerMd(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/43a4de57-0683-4b0f-912a-33d5300b9c77/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/43a4de57-0683-4b0f-912a-33d5300b9c77/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/43a4de57-0683-4b0f-912a-33d5300b9c77"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/43a4de57-0683-4b0f-912a-33d5300b9c77"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeVirginiaLightnerMd]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Virginia Lightner, MD"
	sourceDef.SourceType = pkg.SourceTypeVirginiaLightnerMd
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "virginia-lightner-md.jpg"
	sourceDef.PatientAccessUrl = "https://www.dukehealth.org/find-doctors-physicians/virginia-lightner-md-phd"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
