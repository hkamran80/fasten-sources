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

// https://fhir-myrecord.cerner.com/r4/a1aa7562-925a-46a1-8ac2-dbe276eb4e56/metadata
func GetSourceGilaRiverHealthCareCorporation(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/a1aa7562-925a-46a1-8ac2-dbe276eb4e56/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/a1aa7562-925a-46a1-8ac2-dbe276eb4e56/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/a1aa7562-925a-46a1-8ac2-dbe276eb4e56"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/a1aa7562-925a-46a1-8ac2-dbe276eb4e56"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeGilaRiverHealthCareCorporation]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Gila River Health Care Corporation"
	sourceDef.SourceType = pkg.SourceTypeGilaRiverHealthCareCorporation
	sourceDef.Category = []string{"251B00000X", "251S00000X", "261QE0700X", "282NC0060X", "282NR1301X", "332800000X", "341600000X", "343900000X"}
	sourceDef.Aliases = []string{"RED TAIL HAWK BHS-CM", "RED TAIL HAWK BHS-OP", "RED TAIL HAWK HEALTH CENTER", "RED TAIL HAWK HEALTH CENTER PHARMACY"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1053523357", "1114430147", "1376538231", "1386633915", "1437664182", "1629579230", "1821599424", "1851441075", "1851600191", "1891776902"}}
	sourceDef.PatientAccessUrl = "http://grhc.org/"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
