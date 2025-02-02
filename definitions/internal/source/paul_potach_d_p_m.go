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

// https://fhir-myrecord.cerner.com/r4/25323118-f910-447e-97a3-efa84a07950e/metadata
func GetSourcePaulPotachDPM(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/25323118-f910-447e-97a3-efa84a07950e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/25323118-f910-447e-97a3-efa84a07950e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/25323118-f910-447e-97a3-efa84a07950e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/25323118-f910-447e-97a3-efa84a07950e"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePaulPotachDPM]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Paul Potach D.P.M"
	sourceDef.SourceType = pkg.SourceTypePaulPotachDPM
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.wheelingfootdoctor.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
