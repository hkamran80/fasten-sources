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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10076808/metadata
func GetSourceBeauregardPhysicianGroupLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/patientauthv2/2975cfee-3bcf-4701-9ba4-c3344a567b88/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/patientauthv2/2975cfee-3bcf-4701-9ba4-c3344a567b88/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10076808"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/10076808"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBeauregardPhysicianGroupLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Beauregard Physician Group Llc"
	sourceDef.SourceType = pkg.SourceTypeBeauregardPhysicianGroupLlc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
