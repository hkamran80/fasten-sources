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

// https://fhir-myrecord.cerner.com/r4/81eb3df5-7690-4c07-9f5d-6ffda1c2ed5c/metadata
func GetSourceMchcpMissouriConsolidatedHealthCarePlan(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/81eb3df5-7690-4c07-9f5d-6ffda1c2ed5c/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/81eb3df5-7690-4c07-9f5d-6ffda1c2ed5c/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/81eb3df5-7690-4c07-9f5d-6ffda1c2ed5c"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/81eb3df5-7690-4c07-9f5d-6ffda1c2ed5c"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMchcpMissouriConsolidatedHealthCarePlan]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "MCHCP Missouri Consolidated Health Care Plan"
	sourceDef.SourceType = pkg.SourceTypeMchcpMissouriConsolidatedHealthCarePlan
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://missourigreendoctor.com/?gclid=Cj0KCQiAx6ugBhCcARIsAGNmMbiveXecAfXJ4luH0eK_7OtC0CXoklndAnxjPQxnHG-VkUVmLl_8KgUaAjymEALw_wcB"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
