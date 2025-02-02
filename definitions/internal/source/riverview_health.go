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

// https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/api/FHIR/R4/metadata
func GetSourceRiverviewHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/oauth2/register"

	sourceDef.Audience = "https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epic-fhir.mercy.net/PRDFHIRSTL/RVH/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRiverviewHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Riverview Health"
	sourceDef.SourceType = pkg.SourceTypeRiverviewHealth
	sourceDef.Category = []string{"261QR1300X", "273Y00000X", "282N00000X", "333600000X", "3336C0003X", "3336I0012X", "3336L0003X", "363LF0000X"}
	sourceDef.Aliases = []string{"RIVERVIEW HEALTH"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083763650", "1225275928", "1255688370", "1316945066", "1386309185"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
