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

// https://fhir-myrecord.cerner.com/r4/635f51c4-b647-4331-af69-de60fc5578b8/metadata
func GetSourceWinchesterFootAndAnkleAssociatesPllc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/635f51c4-b647-4331-af69-de60fc5578b8/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/635f51c4-b647-4331-af69-de60fc5578b8/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/635f51c4-b647-4331-af69-de60fc5578b8"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/635f51c4-b647-4331-af69-de60fc5578b8"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeWinchesterFootAndAnkleAssociatesPllc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Winchester Foot and Ankle Associates, PLLC"
	sourceDef.SourceType = pkg.SourceTypeWinchesterFootAndAnkleAssociatesPllc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "http://www.winchesterfootandankle.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
