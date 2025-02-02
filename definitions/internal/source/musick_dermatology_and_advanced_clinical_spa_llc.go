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

// https://fhir-myrecord.cerner.com/r4/GRzkF6yCy_Pe-6F3SJzcXkCEiy_u6zZe/metadata
func GetSourceMusickDermatologyAndAdvancedClinicalSpaLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/GRzkF6yCy_Pe-6F3SJzcXkCEiy_u6zZe/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/GRzkF6yCy_Pe-6F3SJzcXkCEiy_u6zZe/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/GRzkF6yCy_Pe-6F3SJzcXkCEiy_u6zZe"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/GRzkF6yCy_Pe-6F3SJzcXkCEiy_u6zZe"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMusickDermatologyAndAdvancedClinicalSpaLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Musick Dermatology & Advanced Clinical Spa, LLC"
	sourceDef.SourceType = pkg.SourceTypeMusickDermatologyAndAdvancedClinicalSpaLlc
	sourceDef.Category = []string{"207N00000X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1922260587"}}
	sourceDef.PatientAccessUrl = "https://www.musickdermatology.com"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
