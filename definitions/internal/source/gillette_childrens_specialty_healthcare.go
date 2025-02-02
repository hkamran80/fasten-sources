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

// https://fhir-myrecord.cerner.com/r4/1a27c25c-f98c-457f-a68f-46b79305bda6/metadata
func GetSourceGilletteChildrensSpecialtyHealthcare(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/1a27c25c-f98c-457f-a68f-46b79305bda6/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/1a27c25c-f98c-457f-a68f-46b79305bda6/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/1a27c25c-f98c-457f-a68f-46b79305bda6"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/1a27c25c-f98c-457f-a68f-46b79305bda6"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGilletteChildrensSpecialtyHealthcare]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Gillette Children's Specialty Healthcare"
	sourceDef.SourceType = pkg.SourceTypeGilletteChildrensSpecialtyHealthcare
	sourceDef.Category = []string{"261QM1300X"}
	sourceDef.Aliases = []string{}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1225119175"}}
	sourceDef.BrandLogo = "gillette-childrens-specialty-healthcare.jpg"
	sourceDef.PatientAccessUrl = "https://www.gillettechildrens.org/locations/st-paul-campus"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
