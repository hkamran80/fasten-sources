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

// https://fhir-myrecord.cerner.com/r4/527bda78-aebb-4618-87d0-e8d3b484f2ea/metadata
func GetSourceIntegrisHealth2(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/527bda78-aebb-4618-87d0-e8d3b484f2ea/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/527bda78-aebb-4618-87d0-e8d3b484f2ea/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/527bda78-aebb-4618-87d0-e8d3b484f2ea"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/527bda78-aebb-4618-87d0-e8d3b484f2ea"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeIntegrisHealth2]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Integris Health"
	sourceDef.SourceType = pkg.SourceTypeIntegrisHealth2
	sourceDef.Category = []string{"261QR0401X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083065452"}}
	sourceDef.BrandLogo = "integris-health.png"
	sourceDef.PatientAccessUrl = "https://integrisok.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
