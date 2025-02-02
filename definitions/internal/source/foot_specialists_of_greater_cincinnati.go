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

// https://fhir-myrecord.cerner.com/r4/bb0f3815-61cd-44f3-85f8-eeb2e99ef22a/metadata
func GetSourceFootSpecialistsOfGreaterCincinnati(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/bb0f3815-61cd-44f3-85f8-eeb2e99ef22a/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/bb0f3815-61cd-44f3-85f8-eeb2e99ef22a/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/bb0f3815-61cd-44f3-85f8-eeb2e99ef22a"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/bb0f3815-61cd-44f3-85f8-eeb2e99ef22a"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFootSpecialistsOfGreaterCincinnati]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Foot Specialists of Greater Cincinnati"
	sourceDef.SourceType = pkg.SourceTypeFootSpecialistsOfGreaterCincinnati
	sourceDef.Category = []string{"213ES0131X", "335E00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1184782690", "1235152448", "1275691784", "1467511162", "1548545395", "1730247248"}}
	sourceDef.PatientAccessUrl = "http://www.nkyfootdoc.com/contact/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
