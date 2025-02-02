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

// https://fhir-myrecord.cerner.com/r4/7fIEmT-aoS0xFmowQrxBz0hBdpd93tVq/metadata
func GetSourceDonaldMBirchMd(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/7fIEmT-aoS0xFmowQrxBz0hBdpd93tVq/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/7fIEmT-aoS0xFmowQrxBz0hBdpd93tVq/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/7fIEmT-aoS0xFmowQrxBz0hBdpd93tVq"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/7fIEmT-aoS0xFmowQrxBz0hBdpd93tVq"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDonaldMBirchMd]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Donald M. Birch, MD"
	sourceDef.SourceType = pkg.SourceTypeDonaldMBirchMd
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.BrandLogo = "donald-m-birch-md.jpg"
	sourceDef.PatientAccessUrl = "https://healthcare.ascension.org/doctors/1740270420/donald-m-birch-rochester-hills-mi"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
