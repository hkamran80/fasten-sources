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

// https://fhir-myrecord.cerner.com/r4/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb/metadata
func GetSourceMadeliaCommunityHospitalAndClinicDBAMadeliaHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMadeliaCommunityHospitalAndClinicDBAMadeliaHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Madelia Community Hospital and Clinic (D/B/A Madelia Health)"
	sourceDef.SourceType = pkg.SourceTypeMadeliaCommunityHospitalAndClinicDBAMadeliaHealth
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "madelia-community-hospital-and-clinic-d-b-a-madelia-health.jpg"
	sourceDef.PatientAccessUrl = "https://www.madeliahealth.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
