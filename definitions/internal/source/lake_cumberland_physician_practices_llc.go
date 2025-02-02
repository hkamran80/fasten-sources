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

// https://api.platform.athenahealth.com/fhir/r4/metadata
func GetSourceLakeCumberlandPhysicianPracticesLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAthena(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/authorize"
	sourceDef.TokenEndpoint = "https://api.platform.athenahealth.com/oauth2/v1/token"

	sourceDef.Audience = "https://api.platform.athenahealth.com/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://api.platform.athenahealth.com/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeLakeCumberlandPhysicianPracticesLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAthena))

	sourceDef.Display = "Lake Cumberland Physician Practices LLC"
	sourceDef.SourceType = pkg.SourceTypeLakeCumberlandPhysicianPracticesLlc
	sourceDef.Category = []string{"207Q00000X", "207R00000X", "207RG0100X", "207T00000X", "207X00000X", "207XX0005X", "2084N0400X", "2084P0800X", "2084S0012X", "208600000X", "332B00000X", "363A00000X", "363AM0700X", "363AS0400X", "363L00000X", "363LA2100X", "363LA2200X", "363LF0000X"}
	sourceDef.Aliases = []string{"GASTROENTEROLOGY OF LAKE CUMBERLAND", "LAKE CUMBERLAND SURGERY SPECIALISTS"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1104261403", "1174091383", "1386999027", "1437440054", "1689115172", "1700198884", "1811396195", "1982999843"}}
	sourceDef.SecretKeyPrefix = "athena"

	return sourceDef, err
}
