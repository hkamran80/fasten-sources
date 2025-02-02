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

// https://fhir-myrecord.cerner.com/r4/8b053a4e-d04f-420f-b772-45b3551bdda7/metadata
func GetSourceMulticareHealthSystem1(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/8b053a4e-d04f-420f-b772-45b3551bdda7/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/8b053a4e-d04f-420f-b772-45b3551bdda7/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/8b053a4e-d04f-420f-b772-45b3551bdda7"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/8b053a4e-d04f-420f-b772-45b3551bdda7"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMulticareHealthSystem1]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "MultiCare Health System"
	sourceDef.SourceType = pkg.SourceTypeMulticareHealthSystem1
	sourceDef.Category = []string{"104100000X", "207P00000X", "207PS0010X", "207Q00000X", "207QA0000X", "207QA0505X", "207QG0300X", "207QS0010X", "207RS0010X", "207SG0201X", "207V00000X", "207XX0005X", "207XX0801X", "2080S0010X", "2081S0010X", "2083S0010X", "2085B0100X", "2085D0003X", "2085H0002X", "2085N0700X", "2085N0904X", "2085P0229X", "2085R0001X", "2085R0204X", "2085R0205X", "2085U0001X", "213E00000X", "261QR0200X", "261QR0206X", "261QR0207X", "261QR0208X", "261QU0200X", "332B00000X", "332BP3500X", "333600000X", "3336H0001X", "3336I0012X", "3336S0011X", "363A00000X", "363L00000X"}
	sourceDef.Aliases = []string{"MULTICARE HOME INFUSION", "MULTICARE IMAGING - INLAND NW REGION", "OLYMPIC SPORTS & SPINE"}
	sourceDef.Identifiers = map[string][]string{"http://hl7.org/fhir/sid/us-npi": []string{"1154873826", "1235180589", "1346793874", "1407808413", "1548374408", "1639123003", "1679042188", "1700383213", "1700839537", "1710359344"}}
	sourceDef.BrandLogo = "multicare-health-system.png"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
