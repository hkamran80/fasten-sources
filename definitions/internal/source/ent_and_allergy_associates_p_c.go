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

// https://fhir-myrecord.cerner.com/r4/15dd2bda-9d00-4c3d-967c-f374acf8d6ab/metadata
func GetSourceEntAndAllergyAssociatesPC(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/15dd2bda-9d00-4c3d-967c-f374acf8d6ab/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/15dd2bda-9d00-4c3d-967c-f374acf8d6ab/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/15dd2bda-9d00-4c3d-967c-f374acf8d6ab"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/15dd2bda-9d00-4c3d-967c-f374acf8d6ab"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEntAndAllergyAssociatesPC]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "ENT & Allergy Associates, P.C."
	sourceDef.SourceType = pkg.SourceTypeEntAndAllergyAssociatesPC
	sourceDef.Category = []string{"207Y00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1598792244"}}
	sourceDef.PatientAccessUrl = "https://www.entalabama.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
