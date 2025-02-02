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

// https://surescripts.mdmercy.com/fhir-prd/api/FHIR/R4/metadata
func GetSourceMercyHealthServicesMd(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://interconnect.mdmercy.com/fhir-prd/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://interconnect.mdmercy.com/fhir-prd/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://interconnect.mdmercy.com/fhir-prd/oauth2/register"

	sourceDef.Audience = "https://surescripts.mdmercy.com/fhir-prd/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://surescripts.mdmercy.com/fhir-prd/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMercyHealthServicesMd]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Mercy Health Services (MD)"
	sourceDef.SourceType = pkg.SourceTypeMercyHealthServicesMd
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
