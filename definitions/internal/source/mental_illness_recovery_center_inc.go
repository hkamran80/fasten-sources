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

// https://fhir-myrecord.cerner.com/r4/7878c7ba-d1a3-4bd2-9b45-7234d7142c12/metadata
func GetSourceMentalIllnessRecoveryCenterInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/7878c7ba-d1a3-4bd2-9b45-7234d7142c12/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/7878c7ba-d1a3-4bd2-9b45-7234d7142c12/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/7878c7ba-d1a3-4bd2-9b45-7234d7142c12"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/7878c7ba-d1a3-4bd2-9b45-7234d7142c12"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMentalIllnessRecoveryCenterInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Mental Illness Recovery Center Inc."
	sourceDef.SourceType = pkg.SourceTypeMentalIllnessRecoveryCenterInc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "mental-illness-recovery-center-inc.svg"
	sourceDef.PatientAccessUrl = "https://www.mirci.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
