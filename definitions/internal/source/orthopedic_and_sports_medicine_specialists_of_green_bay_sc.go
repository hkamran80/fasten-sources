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

// https://fhir-myrecord.cerner.com/r4/0cda6ed7-9e18-4112-ab11-d5a45a40c8d5/metadata
func GetSourceOrthopedicAndSportsMedicineSpecialistsOfGreenBaySc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/0cda6ed7-9e18-4112-ab11-d5a45a40c8d5/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/0cda6ed7-9e18-4112-ab11-d5a45a40c8d5/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/0cda6ed7-9e18-4112-ab11-d5a45a40c8d5"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/0cda6ed7-9e18-4112-ab11-d5a45a40c8d5"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeOrthopedicAndSportsMedicineSpecialistsOfGreenBaySc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Orthopedic and Sports Medicine Specialists of Green Bay, SC"
	sourceDef.SourceType = pkg.SourceTypeOrthopedicAndSportsMedicineSpecialistsOfGreenBaySc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://osmsgb.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
