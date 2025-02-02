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

// https://fhir-myrecord.cerner.com/r4/3cdd3fbc-bd2e-4e4a-872a-b66a3e5748d2/metadata
func GetSourceFayetteCountyMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/3cdd3fbc-bd2e-4e4a-872a-b66a3e5748d2/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/3cdd3fbc-bd2e-4e4a-872a-b66a3e5748d2/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/3cdd3fbc-bd2e-4e4a-872a-b66a3e5748d2"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/3cdd3fbc-bd2e-4e4a-872a-b66a3e5748d2"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFayetteCountyMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Fayette County Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeFayetteCountyMemorialHospital
	sourceDef.Category = []string{"207Q00000X", "207R00000X", "208600000X", "341600000X"}
	sourceDef.Aliases = []string{"FAYETTE COUNTY MEMORIAL HOSPITAL MEDICAL AND SURGICAL ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1043380579", "1518119791"}}
	sourceDef.PatientAccessUrl = "https://www.adena.org/locations/detail.dT/adena-fayette-medical-center"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
