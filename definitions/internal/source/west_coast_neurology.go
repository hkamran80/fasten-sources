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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067926/metadata
func GetSourceWestCoastNeurology(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/07fd0afe-1cb4-4906-8ed2-a52b00b842fe/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/07fd0afe-1cb4-4906-8ed2-a52b00b842fe/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067926"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10067926"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWestCoastNeurology]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "West Coast Neurology"
	sourceDef.SourceType = pkg.SourceTypeWestCoastNeurology
	sourceDef.Category = []string{"174400000X"}
	sourceDef.Aliases = []string{"HARISH J PATEL MD", "WEST COAST NEUROLOGY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1063632800", "1821566845"}}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
