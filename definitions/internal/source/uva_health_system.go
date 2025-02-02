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

// https://hscsesoap.hscs.virginia.edu/FHIRProxy/api/FHIR/R4/metadata
func GetSourceUvaHealthSystem(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://hscsesoap.hscs.virginia.edu/oauth2/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://hscsesoap.hscs.virginia.edu/oauth2/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://hscsesoap.hscs.virginia.edu/oauth2/oauth2/register"

	sourceDef.Audience = "https://hscsesoap.hscs.virginia.edu/FHIRProxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://hscsesoap.hscs.virginia.edu/FHIRProxy/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUvaHealthSystem]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "UVA Health System"
	sourceDef.SourceType = pkg.SourceTypeUvaHealthSystem
	sourceDef.Category = []string{"282N00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1689669814"}}
	sourceDef.PatientAccessUrl = "https://uvahealth.com/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
