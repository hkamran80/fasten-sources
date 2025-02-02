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

// https://fhir-myrecord.cerner.com/r4/f6f3d034-0e4d-4dff-9105-38992c3c2b63/metadata
func GetSourceTheVillagesTriCountyMedicalCenterIncDBAUfHealthTheVillagesHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/f6f3d034-0e4d-4dff-9105-38992c3c2b63/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/f6f3d034-0e4d-4dff-9105-38992c3c2b63/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/f6f3d034-0e4d-4dff-9105-38992c3c2b63"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/f6f3d034-0e4d-4dff-9105-38992c3c2b63"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTheVillagesTriCountyMedicalCenterIncDBAUfHealthTheVillagesHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "The Villages Tri-County Medical Center, Inc. d/b/a UF Health The Villages Hospital"
	sourceDef.SourceType = pkg.SourceTypeTheVillagesTriCountyMedicalCenterIncDBAUfHealthTheVillagesHospital
	sourceDef.Category = []string{}
	sourceDef.Aliases = []string{}
	sourceDef.PatientAccessUrl = "https://ufhealth.org/uf-health-villages-hospital"
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}
