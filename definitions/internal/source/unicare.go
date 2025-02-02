// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://patient360c.unicare.com/P360Member/api/fhir-r4/metadata
// https://patient360c.unicare.com/P360Member/fhir/documentation?prefix=fhir-r4
func GetSourceUnicare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := GetSourceAnthem(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://patient360c.unicare.com/P360Member/identityserver/connect/authorize"
	sourceDef.TokenEndpoint = "https://patient360c.unicare.com/P360Member/identityserver/connect/token"

	sourceDef.Audience = "https://patient360c.unicare.com/P360Member/api/fhir-r4"

	sourceDef.ApiEndpointBaseUrl = "https://patient360c.unicare.com/P360Member/api/fhir-r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUnicare]; clientIdOk {
		sourceDef.ClientId = clientId
	}

	sourceDef.Display = "Unicare"
	sourceDef.SourceType = pkg.SourceTypeUnicare
	sourceDef.Category = []string{"332B00000X", "335E00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1770802845"}}
	sourceDef.SecretKeyPrefix = "anthem"

	return sourceDef, err
}
