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

// https://fhir-myrecord.cerner.com/r4/kSfiRTuNDYx_39U4kLLwpT-2wyiW22lf/metadata
func GetSourceAdvancedFootAndAnkleClinicLlp(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/kSfiRTuNDYx_39U4kLLwpT-2wyiW22lf/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/kSfiRTuNDYx_39U4kLLwpT-2wyiW22lf/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/kSfiRTuNDYx_39U4kLLwpT-2wyiW22lf"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/kSfiRTuNDYx_39U4kLLwpT-2wyiW22lf"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAdvancedFootAndAnkleClinicLlp]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Advanced Foot & Ankle Clinic, LLP"
	sourceDef.SourceType = pkg.SourceTypeAdvancedFootAndAnkleClinicLlp
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "advanced-foot-and-ankle-clinic-llp.webp"
	sourceDef.PatientAccessUrl = "https://www.advancedfootandankleclinic.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
