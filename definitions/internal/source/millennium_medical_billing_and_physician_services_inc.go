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

// https://fhir-myrecord.cerner.com/r4/04131558-4c98-43ea-8da0-5fb393cbdb29/metadata
func GetSourceMillenniumMedicalBillingAndPhysicianServicesInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/04131558-4c98-43ea-8da0-5fb393cbdb29/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/04131558-4c98-43ea-8da0-5fb393cbdb29/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/04131558-4c98-43ea-8da0-5fb393cbdb29"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/04131558-4c98-43ea-8da0-5fb393cbdb29"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMillenniumMedicalBillingAndPhysicianServicesInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Millennium Medical Billing & Physician Services, Inc."
	sourceDef.SourceType = pkg.SourceTypeMillenniumMedicalBillingAndPhysicianServicesInc
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://millenniummedicalbilling.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
