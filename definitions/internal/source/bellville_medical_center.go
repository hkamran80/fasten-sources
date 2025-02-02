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

// https://fhir-myrecord.cerner.com/r4/16a6a2d0-9070-4585-9404-d87e500b89cb/metadata
func GetSourceBellvilleMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/16a6a2d0-9070-4585-9404-d87e500b89cb/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/16a6a2d0-9070-4585-9404-d87e500b89cb/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/16a6a2d0-9070-4585-9404-d87e500b89cb"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/16a6a2d0-9070-4585-9404-d87e500b89cb"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBellvilleMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Bellville Medical Center"
	sourceDef.SourceType = pkg.SourceTypeBellvilleMedicalCenter
	sourceDef.Category = []string{"207P00000X", "207Q00000X", "207R00000X", "208600000X", "208C00000X", "261QR1300X", "282N00000X"}
	sourceDef.Aliases = []string{"BELLVILLE INTERNAL AND FAMILY MEDICINE CLINIC", "BELLVILLE MEDICAL CENTER", "MID COAST MEDICAL CENTER-BELLVILLE", "MID COAST MEDICAL CLINIC-BELLVILLE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1457066151", "1477857332", "1902488240"}}
	sourceDef.PatientAccessUrl = "http://www.bellvillemc.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
