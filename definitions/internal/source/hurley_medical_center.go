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

// https://fhir.hurleymc.com/fhir/api/FHIR/R4/metadata
func GetSourceHurleyMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.hurleymc.com/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://fhir.hurleymc.com/fhir/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://fhir.hurleymc.com/fhir/oauth2/register"

	sourceDef.Audience = "https://fhir.hurleymc.com/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.hurleymc.com/fhir/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHurleyMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Hurley Medical Center"
	sourceDef.SourceType = pkg.SourceTypeHurleyMedicalCenter
	sourceDef.Category = []string{"1041C0700X", "207Q00000X", "207R00000X", "207V00000X", "2085R0202X", "208G00000X", "261QM0850X", "282N00000X", "282NW0100X"}
	sourceDef.Aliases = []string{"LAPEER FAMILY CARE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1003271792", "1093207425", "1093915373", "1316383706", "1710296348", "1730420217", "1760730469", "1821178567", "1841545209", "1952709834"}}
	sourceDef.PatientAccessUrl = "https://www.hurleymc.com/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
