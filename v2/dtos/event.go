//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

// Event represents a single measurable event read from a device
// This object and its properties correspond to the Event object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/Event
type Event struct {
	common.Versionable `json:",inline"`
	ID                 string        `json:"id"`
	Pushed             int64         `json:"pushed,omitempty"`
	DeviceName         string        `json:"deviceName"`
	Created            int64         `json:"created"`
	Modified           int64         `json:"modified,omitempty"`
	Origin             int64         `json:"origin"`
	Readings           []BaseReading `json:"readings"`
}
