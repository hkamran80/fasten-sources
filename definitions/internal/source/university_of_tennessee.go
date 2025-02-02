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

// https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4/metadata
func GetSourceUniversityOfTennessee(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeUniversityOfTennessee]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "University Of Tennessee"
	sourceDef.SourceType = pkg.SourceTypeUniversityOfTennessee
	sourceDef.Category = []string{"1223P0106X", "1223P0221X", "1223X0400X", "156FX1700X", "207Q00000X", "225100000X", "229N00000X", "282N00000X", "363LF0000X", "390200000X"}
	sourceDef.Aliases = []string{"DURABLE MEDICAL EQUIPMENT", "UNIV. OF TN HEALTH SCIENCE CENTER COLLEGE OF NURSING"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1174221337", "1235666793", "1386735298", "1427420561", "1487723094", "1598767980", "1740577428", "1780686170", "1821164724", "1902278179"}}
	sourceDef.BrandLogo = "university-of-tennessee.svg"
	sourceDef.PatientAccessUrl = "https://www.utk.edu/"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
