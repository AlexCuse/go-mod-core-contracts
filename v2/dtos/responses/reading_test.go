//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

func TestNewReadingCountResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedMessage := "unit test message"
	expectedCount := uint32(1000)
	actual := NewReadingCountResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedCount)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCount, actual.Count)
}

func TestNewReadingCountResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedCount := uint32(1000)
	actual := NewReadingCountResponseNoMessage(expectedRequestID, expectedStatusCode, expectedCount)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedCount, actual.Count)
}

func TestNewReadingResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedMessage := "unit test message"
	expectedReading := dtos.BaseReading{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewReadingResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedReading)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedReading, actual.Reading)
}

func TestNewReadingResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedReading := dtos.BaseReading{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewReadingResponseNoMessage(expectedRequestID, expectedStatusCode, expectedReading)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedReading, actual.Reading)
}
