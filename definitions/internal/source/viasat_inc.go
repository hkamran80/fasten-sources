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

// https://fhir-myrecord.cerner.com/r4/86fe2a9d-7304-48e6-91d1-0feeb3fee8ac/metadata
func GetSourceViasatInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/86fe2a9d-7304-48e6-91d1-0feeb3fee8ac/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/86fe2a9d-7304-48e6-91d1-0feeb3fee8ac/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/86fe2a9d-7304-48e6-91d1-0feeb3fee8ac"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/86fe2a9d-7304-48e6-91d1-0feeb3fee8ac"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeViasatInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Viasat, Inc."
	sourceDef.SourceType = pkg.SourceTypeViasatInc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.facebook.com/Viasat/events/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
