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

// https://fhir-myrecord.cerner.com/r4/ab4ef341-80b4-40a7-999d-52b52314ec6e/metadata
func GetSourceSweenyCommunityHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ab4ef341-80b4-40a7-999d-52b52314ec6e/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ab4ef341-80b4-40a7-999d-52b52314ec6e/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ab4ef341-80b4-40a7-999d-52b52314ec6e"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ab4ef341-80b4-40a7-999d-52b52314ec6e"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSweenyCommunityHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Sweeny Community Hospital"
	sourceDef.SourceType = pkg.SourceTypeSweenyCommunityHospital
	sourceDef.Category = []string{"101YP2500X", "207LP2900X", "207P00000X", "207Q00000X", "207R00000X", "225100000X", "225200000X", "275N00000X", "282NC0060X", "3416L0300X", "363A00000X", "363L00000X"}
	sourceDef.Aliases = []string{"SWEENY COMMUNITY HOSPITAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1023011657"}}
	sourceDef.BrandLogo = "sweeny-community-hospital.jpeg"
	sourceDef.PatientAccessUrl = "https://sweenyhospital.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
