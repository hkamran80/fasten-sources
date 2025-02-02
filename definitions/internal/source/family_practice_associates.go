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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/57963/metadata
func GetSourceFamilyPracticeAssociates(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/63151d27-19a2-4386-a353-a33b00fe9fc4/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/63151d27-19a2-4386-a353-a33b00fe9fc4/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/57963"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/57963"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeFamilyPracticeAssociates]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Family Practice Associates"
	sourceDef.SourceType = pkg.SourceTypeFamilyPracticeAssociates
	sourceDef.Category = []string{"207Q00000X", "208D00000X", "261QP2300X", "302F00000X", "363A00000X", "363L00000X"}
	sourceDef.Aliases = []string{"DEACONESS IL CLINIC ELEVEN S", "FAMILY HEALTH CENTER", "FAMILY HEALTH CONVENIENT CARE CLINIC", "FAMILY PRACTICE ASSOCIATES"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1104014463", "1366499931", "1447329370", "1568513273", "1568576619", "1679952154", "1700982295", "1730623943", "1962128256", "1982685293"}}
	sourceDef.PatientAccessUrl = "https://www.ourfpa.com/"
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
