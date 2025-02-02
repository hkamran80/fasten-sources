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

// https://fhir-myrecord.cerner.com/r4/EmqFrr057wTBNwNrvJBrzi2LHJQc9ijd/metadata
func GetSourceDareFootCareLlc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/EmqFrr057wTBNwNrvJBrzi2LHJQc9ijd/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/EmqFrr057wTBNwNrvJBrzi2LHJQc9ijd/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/EmqFrr057wTBNwNrvJBrzi2LHJQc9ijd"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/EmqFrr057wTBNwNrvJBrzi2LHJQc9ijd"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDareFootCareLlc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "DARE Foot Care, LLC"
	sourceDef.SourceType = pkg.SourceTypeDareFootCareLlc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://npino.com/podiatrist/1629519517-dare-foot-care-llc/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
