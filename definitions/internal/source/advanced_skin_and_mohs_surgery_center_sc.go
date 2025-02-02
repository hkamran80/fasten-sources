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

// https://fhir-myrecord.cerner.com/r4/b438bc62-3485-4cac-8230-6d587898797b/metadata
func GetSourceAdvancedSkinAndMohsSurgeryCenterSc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/b438bc62-3485-4cac-8230-6d587898797b/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/b438bc62-3485-4cac-8230-6d587898797b/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/b438bc62-3485-4cac-8230-6d587898797b"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/b438bc62-3485-4cac-8230-6d587898797b"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAdvancedSkinAndMohsSurgeryCenterSc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Advanced Skin and Mohs Surgery Center, SC"
	sourceDef.SourceType = pkg.SourceTypeAdvancedSkinAndMohsSurgeryCenterSc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "http://www.asmsc.doctor"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
