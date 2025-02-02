// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package platform

import (
	"context"
	base "github.com/fastenhealth/fasten-sources/clients/internal/base"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

type sourceClientCerner struct {
	models.SourceClient
}

/*
https://groups.google.com/g/cerner-fhir-developers
http://fhir.cerner.com/millennium/r4/#authorization
*/
// https://fhir-ehr.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/.well-known/smart-configuration
// https://fhir-myrecord.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/metadata
// https://docs.google.com/document/d/10RnVyF1etl_17pyCyK96tyhUWRbrTyEcqpwzW-Z-Ybs/edit
func GetSourceClientCerner(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, error) {
	baseClient, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	if err != nil {
		return nil, err
	}
	// API requires the following headers for every request
	baseClient.Headers["Accept"] = "application/fhir+json"

	return sourceClientCerner{baseClient}, err
}

// Operation-PatientEverything is not supported - https://build.fhir.org/operation-patient-everything.html
// Manually processing individual resources
func (c sourceClientCerner) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	supportedResources := append(c.GetUsCoreResources(), []string{"Account", "Appointment", "Consent", "FamilyMemberHistory", "InsurancePlan", "MedicationRequest", "NutritionOrder", "Person", "Provenance", "Questionnaire", "QuestionnaireResponse", "RelatedPerson", "Schedule", "ServiceRequest", "Slot"}...)
	return c.SyncAllByResourceName(db, supportedResources)
}
