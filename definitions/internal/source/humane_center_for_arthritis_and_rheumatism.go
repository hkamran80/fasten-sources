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

// https://fhir-myrecord.cerner.com/r4/ecb0768b-f462-488f-a03b-0c40f4d77fbb/metadata
func GetSourceHumaneCenterForArthritisAndRheumatism(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/ecb0768b-f462-488f-a03b-0c40f4d77fbb/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/ecb0768b-f462-488f-a03b-0c40f4d77fbb/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/ecb0768b-f462-488f-a03b-0c40f4d77fbb"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/ecb0768b-f462-488f-a03b-0c40f4d77fbb"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHumaneCenterForArthritisAndRheumatism]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Humane Center for Arthritis and Rheumatism"
	sourceDef.SourceType = pkg.SourceTypeHumaneCenterForArthritisAndRheumatism
	sourceDef.Category = []string{"207RR0500X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1982009551"}}
	sourceDef.PatientAccessUrl = "http://www.humanecenter.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
