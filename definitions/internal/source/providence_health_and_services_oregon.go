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

// https://fhir-myrecord.cerner.com/r4/57956409-cd51-43ce-a57e-5c0de3d73983/metadata
func GetSourceProvidenceHealthAndServicesOregon(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/57956409-cd51-43ce-a57e-5c0de3d73983/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/57956409-cd51-43ce-a57e-5c0de3d73983/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/57956409-cd51-43ce-a57e-5c0de3d73983"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/57956409-cd51-43ce-a57e-5c0de3d73983"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeProvidenceHealthAndServicesOregon]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Providence Health & Services - Oregon"
	sourceDef.SourceType = pkg.SourceTypeProvidenceHealthAndServicesOregon
	sourceDef.Category = []string{"204E00000X", "207RC0000X", "207RH0002X", "207VG0400X", "261Q00000X", "261QX0100X", "363L00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083059778", "1114255247", "1235259532", "1245416965", "1407109937", "1447561626", "1639593387", "1790025104", "1801107081", "1861825820"}}
	sourceDef.PatientAccessUrl = "https://www.providence.org/locations/or/community-health-division"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
