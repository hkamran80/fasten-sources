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

// https://fhir-myrecord.cerner.com/r4/823d67d6-97c5-40a2-b084-56d06cb82099/metadata
func GetSourceMadisonCountyMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/823d67d6-97c5-40a2-b084-56d06cb82099/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/823d67d6-97c5-40a2-b084-56d06cb82099/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/823d67d6-97c5-40a2-b084-56d06cb82099"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/823d67d6-97c5-40a2-b084-56d06cb82099"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMadisonCountyMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Madison County Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeMadisonCountyMemorialHospital
	sourceDef.Category = []string{"163W00000X", "2085R0202X", "251G00000X", "261QE0002X", "261QP2300X", "261QR1300X", "275N00000X", "282NC0060X", "363L00000X"}
	sourceDef.Aliases = []string{"HEALTH TRUST PHYSICIANS CLINIC"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1093929390", "1194880112", "1386709129", "1457520744", "1497810097", "1639314941", "1710042973", "1831104611", "1962567990"}}
	sourceDef.PatientAccessUrl = "https://www.madisonhealth.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
