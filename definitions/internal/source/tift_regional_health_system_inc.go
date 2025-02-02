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

// https://fhir-myrecord.cerner.com/r4/6e2dde13-d4ea-4659-900f-6a252ed00859/metadata
func GetSourceTiftRegionalHealthSystemInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/6e2dde13-d4ea-4659-900f-6a252ed00859/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/6e2dde13-d4ea-4659-900f-6a252ed00859/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/6e2dde13-d4ea-4659-900f-6a252ed00859"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/6e2dde13-d4ea-4659-900f-6a252ed00859"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTiftRegionalHealthSystemInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Tift Regional Health System, Inc."
	sourceDef.SourceType = pkg.SourceTypeTiftRegionalHealthSystemInc
	sourceDef.Category = []string{"174400000X", "251G00000X", "261QE0700X", "261QR1300X"}
	sourceDef.Aliases = []string{"COOK PRIMARY CARE", "HOSPICE OF TIFT AREA", "SOUTHWELL MEDICAL ADEL PRIMARY CARE", "TIFT REGIONAL DIALYSIS CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1831298660", "1871816918", "1902909427"}}
	sourceDef.BrandLogo = "tift-regional-health-system-inc.jpg"
	sourceDef.PatientAccessUrl = "https://mysouthwell.com"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
