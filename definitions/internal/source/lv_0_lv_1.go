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

// https://scmprodweb.lv0.hos.allscriptscloud.com/R4/open-Prod/metadata
func GetSourceLv0Lv1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/098c56d7-4d28-424c-9772-a99a008c3f63/connect/authorize"
	sourceDef.TokenEndpoint = "https://global.open.allscripts.com/fhirroute/fmhpatientauth/fmhorgid/098c56d7-4d28-424c-9772-a99a008c3f63/connect/token"

	sourceDef.Audience = "https://scmprodweb.lv0.hos.allscriptscloud.com/R4/open-Prod"

	sourceDef.ApiEndpointBaseUrl = "https://scmprodweb.lv0.hos.allscriptscloud.com/R4/open-Prod"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeLv0Lv1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "LV0 - LV1"
	sourceDef.SourceType = pkg.SourceTypeLv0Lv1
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}
