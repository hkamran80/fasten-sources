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

// https://fhir-myrecord.cerner.com/r4/2229c7e9-33af-4844-8811-015c0ae6032b/metadata
func GetSourceManyRiversWholeHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/2229c7e9-33af-4844-8811-015c0ae6032b/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/2229c7e9-33af-4844-8811-015c0ae6032b/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/2229c7e9-33af-4844-8811-015c0ae6032b"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/2229c7e9-33af-4844-8811-015c0ae6032b"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeManyRiversWholeHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Many Rivers Whole Health"
	sourceDef.SourceType = pkg.SourceTypeManyRiversWholeHealth
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "many-rivers-whole-health.jpeg"
	sourceDef.PatientAccessUrl = "https://manyriverswholehealth.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
