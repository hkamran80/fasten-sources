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

// https://fhir-myrecord.cerner.com/r4/0d26bd04-9128-4f58-8e6d-ffda5cd00f89/metadata
func GetSourceAssociatesInAnesthesia(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/0d26bd04-9128-4f58-8e6d-ffda5cd00f89/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/0d26bd04-9128-4f58-8e6d-ffda5cd00f89/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/0d26bd04-9128-4f58-8e6d-ffda5cd00f89"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/0d26bd04-9128-4f58-8e6d-ffda5cd00f89"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAssociatesInAnesthesia]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Associates In Anesthesia"
	sourceDef.SourceType = pkg.SourceTypeAssociatesInAnesthesia
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "http://www.aiamd.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
