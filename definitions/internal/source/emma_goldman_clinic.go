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

// https://fhir-myrecord.cerner.com/r4/52689f1e-9825-4b5d-8853-96de57ae62a9/metadata
func GetSourceEmmaGoldmanClinic(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/52689f1e-9825-4b5d-8853-96de57ae62a9/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/52689f1e-9825-4b5d-8853-96de57ae62a9/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/52689f1e-9825-4b5d-8853-96de57ae62a9"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/52689f1e-9825-4b5d-8853-96de57ae62a9"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeEmmaGoldmanClinic]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Emma Goldman Clinic"
	sourceDef.SourceType = pkg.SourceTypeEmmaGoldmanClinic
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://emmagoldman.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
