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

// https://fhir-myrecord.cerner.com/r4/022992dd-25dd-480b-9df5-6959b1bac6bc/metadata
func GetSourceDrAmyMorris(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/022992dd-25dd-480b-9df5-6959b1bac6bc/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/022992dd-25dd-480b-9df5-6959b1bac6bc/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/022992dd-25dd-480b-9df5-6959b1bac6bc"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/022992dd-25dd-480b-9df5-6959b1bac6bc"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDrAmyMorris]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Dr. Amy Morris"
	sourceDef.SourceType = pkg.SourceTypeDrAmyMorris
	sourceDef.Category = []string{"207N00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1316914583"}}
	sourceDef.PatientAccessUrl = "https://patientportal.advancedmd.com/126242/account/logon"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
