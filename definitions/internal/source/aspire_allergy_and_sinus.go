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
func GetSourceAspireAllergyAndSinus(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAspireAllergyAndSinus]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "Aspire Allergy & Sinus"
	sourceDef.SourceType = pkg.SourceTypeAspireAllergyAndSinus
	sourceDef.Category = []string{"207KA0200X"}
	sourceDef.Aliases = []string{"ASPIRE ALLERGY & SINUS"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1073120945", "1497362362", "1669089538", "1881201754"}}
	sourceDef.PatientAccessUrl = "https://worker.mturk.com/projects/3CTCX9NXCJJWWANBW47QJI84M2SJLA/tasks/3M47JKRKDAYOO8QA0RUNUQWDEU8685?assignment_id=3S4AW7T80PF7AC491UFDRHUU7CML4A&from_queue=true"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
