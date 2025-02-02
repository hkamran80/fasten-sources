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

// https://fhir-myrecord.cerner.com/r4/9cfe5b83-b5f6-41dc-99dc-2bd4b3e16e68/metadata
func GetSourceAtiSmRichburgHealthCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/9cfe5b83-b5f6-41dc-99dc-2bd4b3e16e68/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/9cfe5b83-b5f6-41dc-99dc-2bd4b3e16e68/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/9cfe5b83-b5f6-41dc-99dc-2bd4b3e16e68"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/9cfe5b83-b5f6-41dc-99dc-2bd4b3e16e68"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeAtiSmRichburgHealthCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "ATI SM Richburg Health Center"
	sourceDef.SourceType = pkg.SourceTypeAtiSmRichburgHealthCenter
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "ati-sm-occupational-health-center.png"
	sourceDef.PatientAccessUrl = "https://worker.mturk.com/projects/3CTCX9NXCJJWWANBW47QJI84M2SJLA/tasks/3JYPJ2TAZV55DTI2N1Z2AI4GW3BFP1?assignment_id=352YTHGRO9A5OGLJA90FZP1NQVF4HO&from_queue=true"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
