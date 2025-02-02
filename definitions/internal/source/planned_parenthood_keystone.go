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
func GetSourcePlannedParenthoodKeystone(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePlannedParenthoodKeystone]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Planned Parenthood Keystone"
	sourceDef.SourceType = pkg.SourceTypePlannedParenthoodKeystone
	sourceDef.Category = []string{"207VG0400X", "251K00000X", "261Q00000X", "261QA0005X", "261QF0050X"}
	sourceDef.Aliases = []string{"BENSALEM MEDICAL CENTER", "HARRISBURG MEDICAL CENTER", "LANCASTER MEDICAL CENTER", "PLANNED PARENTHOOD KEYSTONE", "READING MEDICAL CENTER", "YORK MEDICAL CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083866263", "1093975823", "1487655981", "1568639029", "1659531804", "1699814103", "1851464317", "1881184760"}}
	sourceDef.BrandLogo = "planned-parenthood-keystone.jpg"
	sourceDef.PatientAccessUrl = "https://www.plannedparenthood.org/planned-parenthood-keystone?utm_campaign=keystone-affiliate&utm_medium=organic&utm_source=local-listing"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
