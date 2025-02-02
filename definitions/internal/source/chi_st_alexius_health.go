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

// https://rp.chihealth.com/fhir/FHIRSTA/api/FHIR/R4/metadata
func GetSourceChiStAlexiusHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://rp.chihealth.com/fhir/FHIRSTA/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://rp.chihealth.com/fhir/FHIRSTA/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://rp.chihealth.com/fhir/FHIRSTA/oauth2/register"

	sourceDef.Audience = "https://rp.chihealth.com/fhir/FHIRSTA/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://rp.chihealth.com/fhir/FHIRSTA/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeChiStAlexiusHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "CHI St. Alexius Health"
	sourceDef.SourceType = pkg.SourceTypeChiStAlexiusHealth
	sourceDef.Category = []string{"133V00000X", "261QA1903X", "261QR1300X", "282N00000X", "367500000X"}
	sourceDef.Aliases = []string{"CHI ST. ALEXIUS HEALTH", "CHI ST. ALEXIUS HEALTH TURTLE LAKE", "CHI ST. ALEXIUS HEALTH TURTLE LAKE CLINIC"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1306832654", "1336708536", "1629121074"}}
	sourceDef.PatientAccessUrl = "https://www.chistalexiushealth.org/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
