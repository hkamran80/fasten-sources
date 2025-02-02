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

// https://fhir-myrecord.cerner.com/r4/ntTIZxOIDS2tB2PUqr33FYQo0UWd2W9R/metadata
func GetSourceWrayCommunityDistrictHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ntTIZxOIDS2tB2PUqr33FYQo0UWd2W9R/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ntTIZxOIDS2tB2PUqr33FYQo0UWd2W9R/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ntTIZxOIDS2tB2PUqr33FYQo0UWd2W9R"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ntTIZxOIDS2tB2PUqr33FYQo0UWd2W9R"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWrayCommunityDistrictHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Wray Community District Hospital"
	sourceDef.SourceType = pkg.SourceTypeWrayCommunityDistrictHospital
	sourceDef.Category = []string{"208D00000X", "225100000X", "275N00000X", "282NC0060X"}
	sourceDef.Aliases = []string{"THE WRAY CLINIC", "WRAY COMMUNITY DISTRICT HOSPITAL", "WRAY PHYSICAL THERAPY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083640239", "1477587202", "1831312925", "1912486820"}}
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
