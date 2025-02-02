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

// https://fhir-myrecord.cerner.com/r4/5f5b322c-2217-4325-ad8e-025b10008aca/metadata
func GetSourceMunisingMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/5f5b322c-2217-4325-ad8e-025b10008aca/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/5f5b322c-2217-4325-ad8e-025b10008aca/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/5f5b322c-2217-4325-ad8e-025b10008aca"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/5f5b322c-2217-4325-ad8e-025b10008aca"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMunisingMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Munising Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeMunisingMemorialHospital
	sourceDef.Category = []string{"207Q00000X", "207X00000X", "261QR1300X", "275N00000X"}
	sourceDef.Aliases = []string{"MUNISING MEMORIAL HOSPITAL"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1043562762", "1750894648"}}
	sourceDef.PatientAccessUrl = "https://munisingmemorial.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
