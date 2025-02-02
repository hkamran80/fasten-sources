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

// https://fhir-myrecord.cerner.com/r4/087ff9c5-549d-4a6e-a079-607c5bcde7a6/metadata
func GetSourceClinicaLasCondes(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/087ff9c5-549d-4a6e-a079-607c5bcde7a6/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/087ff9c5-549d-4a6e-a079-607c5bcde7a6/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/087ff9c5-549d-4a6e-a079-607c5bcde7a6"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/087ff9c5-549d-4a6e-a079-607c5bcde7a6"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeClinicaLasCondes]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Clinica Las Condes"
	sourceDef.SourceType = pkg.SourceTypeClinicaLasCondes
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.clinicalascondes.cl/HOME"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
