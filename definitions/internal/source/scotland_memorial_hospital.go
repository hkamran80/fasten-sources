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

// https://fhir-myrecord.cerner.com/r4/bf2bc40a-17cb-4353-8de2-68c75c849bc8/metadata
func GetSourceScotlandMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/bf2bc40a-17cb-4353-8de2-68c75c849bc8/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/bf2bc40a-17cb-4353-8de2-68c75c849bc8/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/bf2bc40a-17cb-4353-8de2-68c75c849bc8"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/bf2bc40a-17cb-4353-8de2-68c75c849bc8"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeScotlandMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Scotland Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeScotlandMemorialHospital
	sourceDef.Category = []string{"207Q00000X", "261QR1300X", "332B00000X"}
	sourceDef.Aliases = []string{"SCOTLAND FAMILY PRACTICE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1265426191", "1700937711", "1932301769"}}
	sourceDef.BrandLogo = "scotland-memorial-hospital.jpg"
	sourceDef.PatientAccessUrl = "https://www.facebook.com/scotlandhealthcare/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
