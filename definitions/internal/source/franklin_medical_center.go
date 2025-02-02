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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75374/metadata
func GetSourceFranklinMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/e44f1017-ebb7-4e5d-8936-a3bf009f3cf6/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/e44f1017-ebb7-4e5d-8936-a3bf009f3cf6/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75374"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/75374"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFranklinMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Franklin Medical Center"
	sourceDef.SourceType = pkg.SourceTypeFranklinMedicalCenter
	sourceDef.Category = []string{"207Q00000X", "261QP2300X", "261QR1300X", "275N00000X", "282NR1301X"}
	sourceDef.Aliases = []string{"CROWVILLE HEALTH CENTER", "FRANKLIN MEDICAL CENTER", "GILBERT HEALTH CENTER", "NEWELLTON RURAL HEALTH CLINIC", "WINNSBORO RURAL HEALTH CLINIC"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1053492579", "1205884723", "1356399539", "1467575670", "1497068951", "1669002002", "1720296775", "1831144054", "1841389251", "1861573388", "1902851124"}}
	sourceDef.PatientAccessUrl = "https://www.fmc-cares.com/"
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
