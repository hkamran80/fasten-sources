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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75426/metadata
func GetSourceNaeemAQaziMdPc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/43466265-f765-4059-8317-a4ec0113eadd/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/43466265-f765-4059-8317-a4ec0113eadd/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75426"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75426"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNaeemAQaziMdPc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Naeem A Qazi Md Pc"
	sourceDef.SourceType = pkg.SourceTypeNaeemAQaziMdPc
	sourceDef.Category = []string{"174400000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1467533323"}}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
