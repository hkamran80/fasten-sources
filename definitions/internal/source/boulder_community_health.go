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

// https://prevprox.bch.org/FHIRproxyPRD/api/FHIR/R4/metadata
func GetSourceBoulderCommunityHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://prevprox.bch.org/FHIRproxyPRD/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://prevprox.bch.org/FHIRproxyPRD/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://prevprox.bch.org/FHIRproxyPRD/oauth2/register"

	sourceDef.Audience = "https://prevprox.bch.org/FHIRproxyPRD/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://prevprox.bch.org/FHIRproxyPRD/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBoulderCommunityHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Boulder Community Health"
	sourceDef.SourceType = pkg.SourceTypeBoulderCommunityHealth
	sourceDef.Category = []string{"2085R0202X", "261Q00000X", "261QM2500X", "261QP2300X", "282N00000X"}
	sourceDef.Aliases = []string{"BOULDER COMMUNITY HOSPITAL", "BOULDER HEART (BOULDER CAMPUS)", "BOULDER HEART LONGMONT", "BOULDER HEART OF ERIE", "BOULDER VALLEY PULMONOLOGY AT ERIE", "BOULDER VALLEY SURGICAL ASSOCIATES OF B.C.H.", "COMMUNITY MEDICAL ASSOC OF B.C.H.", "COMMUNITY MEDICAL ASSOC OF BCH", "COMMUNITY MEDICAL ASSOCIATES OF BCH", "FOOTHILLS NEONATAL PRACTICE", "GUNBARREL FAMILY MEDICINE", "INTERNAL MEDICINE ASSOCIATES LAFAYETTE", "SPRUCE STREET INTERNAL MEDICINE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1073160420", "1124674932", "1225687965", "1285233882", "1487803383", "1558965939", "1598313553", "1649879214", "1740837137", "1942856745"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
