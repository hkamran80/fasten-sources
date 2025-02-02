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

// https://fhir-myrecord.cerner.com/r4/e35355ac-dfca-48b9-9dd9-0b7c408c2a66/metadata
func GetSourceLakeRegionHealthcareCorporation(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/e35355ac-dfca-48b9-9dd9-0b7c408c2a66/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/e35355ac-dfca-48b9-9dd9-0b7c408c2a66/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/e35355ac-dfca-48b9-9dd9-0b7c408c2a66"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/e35355ac-dfca-48b9-9dd9-0b7c408c2a66"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeLakeRegionHealthcareCorporation]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Lake Region Healthcare Corporation"
	sourceDef.SourceType = pkg.SourceTypeLakeRegionHealthcareCorporation
	sourceDef.Category = []string{"111N00000X", "273R00000X", "273Y00000X", "282N00000X", "314000000X", "332900000X", "332B00000X", "332H00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1093713372", "1194862557", "1265148480", "1922471804"}}
	sourceDef.BrandLogo = "lake-region-healthcare-corporation.jpg"
	sourceDef.PatientAccessUrl = "https://www.lrhc.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
