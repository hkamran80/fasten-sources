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

// https://fhir-myrecord.cerner.com/r4/82c1b592-2342-43bd-a292-07e36f400517/metadata
func GetSourceMemorialHospitalOfSweetwaterCounty(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/82c1b592-2342-43bd-a292-07e36f400517/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/82c1b592-2342-43bd-a292-07e36f400517/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/82c1b592-2342-43bd-a292-07e36f400517"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/82c1b592-2342-43bd-a292-07e36f400517"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMemorialHospitalOfSweetwaterCounty]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Memorial Hospital of Sweetwater County"
	sourceDef.SourceType = pkg.SourceTypeMemorialHospitalOfSweetwaterCounty
	sourceDef.Category = []string{"207R00000X", "282N00000X"}
	sourceDef.Aliases = []string{"COUNTY OF SWEETWATER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1063840775", "1437312436", "1558361949"}}
	sourceDef.PatientAccessUrl = "https://www.sweetwatermemorial.com/?utm_source=GMB&utm_medium=organic&utm_campaign=rocksprings"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
