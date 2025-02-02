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

// https://fhir-myrecord.cerner.com/r4/24c3ea72-6b18-4192-9262-677c412fb10a/metadata
func GetSourceVanDiestMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/24c3ea72-6b18-4192-9262-677c412fb10a/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/24c3ea72-6b18-4192-9262-677c412fb10a/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/24c3ea72-6b18-4192-9262-677c412fb10a"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/24c3ea72-6b18-4192-9262-677c412fb10a"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeVanDiestMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Van Diest Medical Center"
	sourceDef.SourceType = pkg.SourceTypeVanDiestMedicalCenter
	sourceDef.Category = []string{"3416L0300X"}
	sourceDef.Aliases = []string{"HAMILTON HOSPITAL AMBULANCE", "VAN DIEST MEDICAL CENTER", "VAN DIEST MEDICAL CENTER AMBULANCE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1083767321"}}
	sourceDef.PatientAccessUrl = "https://www.vandiestmc.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
