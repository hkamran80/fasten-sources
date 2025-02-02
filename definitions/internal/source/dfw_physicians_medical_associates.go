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

// https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4/metadata
func GetSourceDfwPhysiciansMedicalAssociates(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceNextgen(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/authorize"
	sourceDef.TokenEndpoint = "https://fhir.nextgen.com/nge/prod/patient-oauth/token"

	sourceDef.Audience = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.nextgen.com/nge/prod/fhir-api-r4/fhir/r4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDfwPhysiciansMedicalAssociates]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeNextgen))

	sourceDef.Display = "DFW Physicians Medical Associates"
	sourceDef.SourceType = pkg.SourceTypeDfwPhysiciansMedicalAssociates
	sourceDef.Category = []string{"103TC2200X", "207Q00000X", "207R00000X", "207RC0000X", "207RI0011X", "207RP1001X", "2080C0008X", "2084A0401X", "2084P0804X", "2084P0805X", "363L00000X", "363LP2300X"}
	sourceDef.Aliases = []string{"DFWPMA", "DFWPMA FCVC", "FRISCO CARDIAC AND VASCULAR CARE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1376014076"}}
	sourceDef.BrandLogo = "dfw-physicians-medical-associates.jpg"
	sourceDef.PatientAccessUrl = "https://pay.instamed.com/Form/PaymentPortal/Default?id=DFWPHYS"
	sourceDef.SecretKeyPrefix = "nextgen"

	return sourceDef, err
}
