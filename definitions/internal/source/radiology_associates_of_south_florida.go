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

// https://fhir-myrecord.cerner.com/r4/f9aed9f4-5c6e-4527-8cf9-3a241a404968/metadata
func GetSourceRadiologyAssociatesOfSouthFlorida(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f9aed9f4-5c6e-4527-8cf9-3a241a404968/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f9aed9f4-5c6e-4527-8cf9-3a241a404968/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f9aed9f4-5c6e-4527-8cf9-3a241a404968"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f9aed9f4-5c6e-4527-8cf9-3a241a404968"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRadiologyAssociatesOfSouthFlorida]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Radiology Associates of South Florida"
	sourceDef.SourceType = pkg.SourceTypeRadiologyAssociatesOfSouthFlorida
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.superdoctors.com/florida/hospital-clinic/Radiology-Associates-of-South-Florida/98aea428-dd19-4f25-8bc5-3deff457f7e2.html"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
