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

// https://fhir-myrecord.cerner.com/r4/86251441-7ffd-4c0e-91b6-1c16f5c92270/metadata
func GetSourceDrPadmaGuptaMdPa(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/86251441-7ffd-4c0e-91b6-1c16f5c92270/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/86251441-7ffd-4c0e-91b6-1c16f5c92270/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/86251441-7ffd-4c0e-91b6-1c16f5c92270"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/86251441-7ffd-4c0e-91b6-1c16f5c92270"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDrPadmaGuptaMdPa]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Dr. Padma Gupta, MD, PA"
	sourceDef.SourceType = pkg.SourceTypeDrPadmaGuptaMdPa
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://doctor.webmd.com/doctor/padma-gupta-90b725c4-8315-49f5-a3eb-10e54f2b743d-overview"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
