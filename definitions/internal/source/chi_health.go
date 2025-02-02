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

// https://rp.chihealth.com/fhir/api/FHIR/R4/metadata
func GetSourceChiHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://rp.chihealth.com/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://rp.chihealth.com/fhir/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://rp.chihealth.com/fhir/oauth2/register"

	sourceDef.Audience = "https://rp.chihealth.com/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://rp.chihealth.com/fhir/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeChiHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "CHI Health"
	sourceDef.SourceType = pkg.SourceTypeChiHealth
	sourceDef.Category = []string{"101Y00000X", "103T00000X", "1041C0700X", "133V00000X", "163W00000X", "207R00000X", "2084P0800X", "208M00000X", "3336C0003X", "363L00000X", "364S00000X"}
	sourceDef.Aliases = []string{"CHI HEALTH"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1033288923", "1598423220"}}
	sourceDef.BrandLogo = "chi-health.jpg"
	sourceDef.PatientAccessUrl = "https://www.chihealth.com/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
