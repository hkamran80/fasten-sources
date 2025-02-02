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

// https://fhir-myrecord.cerner.com/r4/06e9d567-9657-447f-97fb-d967654387e7/metadata
func GetSourceSaundersMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/06e9d567-9657-447f-97fb-d967654387e7/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/06e9d567-9657-447f-97fb-d967654387e7/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/06e9d567-9657-447f-97fb-d967654387e7"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/06e9d567-9657-447f-97fb-d967654387e7"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSaundersMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Saunders Medical Center"
	sourceDef.SourceType = pkg.SourceTypeSaundersMedicalCenter
	sourceDef.Category = []string{"261QP2300X", "261QR1300X", "275N00000X", "282NC0060X", "313M00000X", "314000000X", "333600000X", "3336C0003X", "3336L0003X"}
	sourceDef.Aliases = []string{"SAUNDERS MEDICAL CENTER LTC", "SAUNDERS MEDICAL CENTER LTC PHARMACY", "SAUNDERS MEDICAL CENTER LTC SKILLED"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1013364678", "1326416850", "1417337635", "1467442244", "1609866409", "1780642660", "1831188705"}}
	sourceDef.PatientAccessUrl = "https://saundersmedicalcenter.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
