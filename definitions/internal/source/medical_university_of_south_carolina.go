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

// https://fhirprod.musc.edu/fhirprod/api/FHIR/R4/metadata
func GetSourceMedicalUniversityOfSouthCarolina(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhirprod.musc.edu/fhirprod/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://fhirprod.musc.edu/fhirprod/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://fhirprod.musc.edu/fhirprod/oauth2/register"

	sourceDef.Audience = "https://fhirprod.musc.edu/fhirprod/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://fhirprod.musc.edu/fhirprod/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMedicalUniversityOfSouthCarolina]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Medical University of South Carolina"
	sourceDef.SourceType = pkg.SourceTypeMedicalUniversityOfSouthCarolina
	sourceDef.Category = []string{"261QM2500X", "261QP2000X", "282N00000X", "282NC0060X", "282NC2000X", "291U00000X"}
	sourceDef.Aliases = []string{"MUSC PATHOLOGY OUTREACH SERVICES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1194924589", "1275840423", "1285818104", "1295161420", "1437636057", "1508052283", "1538452313", "1568526796", "1659560670", "1780714907"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
