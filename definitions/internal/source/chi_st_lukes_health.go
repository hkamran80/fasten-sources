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

// https://rpsouth.catholichealth.net/fhir/api/FHIR/R4/metadata
func GetSourceChiStLukesHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://rpsouth.catholichealth.net/fhir/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://rpsouth.catholichealth.net/fhir/oauth2/token"
	sourceDef.RegistrationEndpoint = "https://rpsouth.catholichealth.net/fhir/oauth2/register"

	sourceDef.Audience = "https://rpsouth.catholichealth.net/fhir/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://rpsouth.catholichealth.net/fhir/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeChiStLukesHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "CHI St. Luke's Health"
	sourceDef.SourceType = pkg.SourceTypeChiStLukesHealth
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://www.stlukeshealth.org/"
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}
