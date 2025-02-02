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

// https://fhir-myrecord.cerner.com/r4/a9daee2a-e42f-4702-9c42-486fbda403f6/metadata
func GetSourceNeoshoMemorialRegionalMedicalCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a9daee2a-e42f-4702-9c42-486fbda403f6/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a9daee2a-e42f-4702-9c42-486fbda403f6/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a9daee2a-e42f-4702-9c42-486fbda403f6"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a9daee2a-e42f-4702-9c42-486fbda403f6"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeNeoshoMemorialRegionalMedicalCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Neosho Memorial Regional Medical Center"
	sourceDef.SourceType = pkg.SourceTypeNeoshoMemorialRegionalMedicalCenter
	sourceDef.Category = []string{"207X00000X", "208600000X", "251B00000X", "261QM2500X", "261QP2300X", "261QR1300X", "275N00000X", "282NC0060X", "332B00000X", "332BX2000X", "364SE0003X", "364SF0001X"}
	sourceDef.Aliases = []string{"HOME CARE PRODUCTS OF NEOSHO MEMORIAL", "NEW HORIZONS COUNSELING", "NMRMC ERIE FAMILY CARE CLINIC", "NMRMC FAMILY MEDICINE", "NMRMC ORTHOPEDIC CLINIC", "NMRMC SURGERY CLINIC", "NMRMC SWINGBED", "NMRMC WOMEN'S HEALTH CENTER"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1023251469", "1073566949", "1275179434", "1497121735", "1497869374", "1619261211", "1679902126", "1730557166", "1750485082", "1912014473"}}
	sourceDef.BrandLogo = "neosho-memorial-regional-medical-center.svg"
	sourceDef.PatientAccessUrl = "https://www.nmrmc.com/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
