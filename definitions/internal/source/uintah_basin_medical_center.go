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

// https://fhir-myrecord.cerner.com/r4/f11b499f-f009-416a-ba91-79d8b37ce81a/metadata
func GetSourceUintahBasinMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f11b499f-f009-416a-ba91-79d8b37ce81a/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f11b499f-f009-416a-ba91-79d8b37ce81a/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f11b499f-f009-416a-ba91-79d8b37ce81a"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f11b499f-f009-416a-ba91-79d8b37ce81a"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUintahBasinMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Uintah Basin Medical Center"
	sourceDef.SourceType = pkg.SourceTypeUintahBasinMedicalCenter
	sourceDef.Category = []string{"251E00000X", "251G00000X", "261Q00000X", "261QE0700X", "261QM1300X", "261QR1300X", "282N00000X", "332B00000X", "341600000X"}
	sourceDef.Aliases = []string{"DME", "UBMC CLINIC", "UBMC CLINIC DUCHESNE", "UINTAH BASIN HOME HEALTH", "UINTAH BASIN MEDICAL CENTER DIALYSIS", "UINTAH BASIN MEDICAL CENTER DIALYSIS VERNAL", "UINTAH BASIN MEDICAL CENTER EMS", "UINTAH BASIN MEDICAL CENTER THE CLINIC VERNAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1235736059", "1265613541", "1508822347", "1598335903", "1629034376", "1679193338", "1760180566", "1851920151", "1871556217", "1922280007"}}
	sourceDef.BrandLogo = "uintah-basin-medical-center.jpeg"
	sourceDef.PatientAccessUrl = "https://ubh.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
