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

// https://epicproxy.et1095.epichosted.com/FHIRProxy/api/FHIR/R4/metadata
func GetSourceKennedyKriegerInstitute(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://epicproxy.et1095.epichosted.com/FHIRProxy/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://epicproxy.et1095.epichosted.com/FHIRProxy/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://epicproxy.et1095.epichosted.com/FHIRProxy/oauth2/register"

	sourceDef.Audience = "https://epicproxy.et1095.epichosted.com/FHIRProxy/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://epicproxy.et1095.epichosted.com/FHIRProxy/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeKennedyKriegerInstitute]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "Kennedy Krieger Institute"
	sourceDef.SourceType = pkg.SourceTypeKennedyKriegerInstitute
	sourceDef.Category = []string{"103TC2200X", "261QD1600X", "261QH0700X", "281PC2000X", "282E00000X", "283Q00000X", "283XC2000X", "284300000X", "291U00000X"}
	sourceDef.Aliases = []string{"KENNEDY KRIEGER INSTITUTE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1053572958", "1619075926", "1649595141", "1699778415", "1710387550", "1730505363"}}
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
