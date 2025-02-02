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

// https://fhir-myrecord.cerner.com/r4/4186e40c-ae1e-4c84-b5ea-1707df8e59de/metadata
func GetSourceColonAndRectalSurgeryAssociates(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/4186e40c-ae1e-4c84-b5ea-1707df8e59de/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/4186e40c-ae1e-4c84-b5ea-1707df8e59de/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/4186e40c-ae1e-4c84-b5ea-1707df8e59de"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/4186e40c-ae1e-4c84-b5ea-1707df8e59de"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeColonAndRectalSurgeryAssociates]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Colon & Rectal Surgery Associates"
	sourceDef.SourceType = pkg.SourceTypeColonAndRectalSurgeryAssociates
	sourceDef.Category = []string{"208C00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1184689507"}}
	sourceDef.PatientAccessUrl = "https://colonrectal.net/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
