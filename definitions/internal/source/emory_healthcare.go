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

// https://epicrp-prd.eushc.org/OAUTH2-PRD/api/FHIR/R4/metadata
func GetSourceEmoryHealthcare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epicrp-prd.eushc.org/OAUTH2-PRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicrp-prd.eushc.org/OAUTH2-PRD/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epicrp-prd.eushc.org/OAUTH2-PRD/oauth2/register"

	sourceDef.Audience = "https://epicrp-prd.eushc.org/OAUTH2-PRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicrp-prd.eushc.org/OAUTH2-PRD/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEmoryHealthcare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Emory Healthcare"
	sourceDef.SourceType = pkg.SourceTypeEmoryHealthcare
	sourceDef.Category = []string{"1041C0700X", "261QM2500X", "282N00000X", "282NC0060X", "313M00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1154643336", "1235339227", "1336345982", "1417157405", "1578834008", "1598964082", "1659674992", "1740689496", "1801834908", "1972792885"}}
	sourceDef.PatientAccessUrl = "https://www.emoryhealthcare.org/index.html"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
