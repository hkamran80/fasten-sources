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

// https://fhir-myrecord.cerner.com/r4/940a6958-277a-4baa-b4ad-64568a7458eb/metadata
func GetSourceRichardsMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/940a6958-277a-4baa-b4ad-64568a7458eb/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/940a6958-277a-4baa-b4ad-64568a7458eb/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/940a6958-277a-4baa-b4ad-64568a7458eb"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/940a6958-277a-4baa-b4ad-64568a7458eb"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeRichardsMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Richards Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeRichardsMemorialHospital
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.ahd.com/free_profile.php?hcfa_id=a790b3b0d010d499a0255cc681573129&ek=318a20d2a6134fb7eaa7731713c31f24"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
