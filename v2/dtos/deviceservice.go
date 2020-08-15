//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

// DeviceService represents a service that is responsible for proxying connectivity between a set of devices and the
// EdgeX Foundry core services.
// This object and its properties correspond to the DeviceService object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceService
type DeviceService struct {
	common.Versionable `json:",inline"`
	Id                 string   `json:"id,omitempty"`
	Name               string   `json:"name" validate:"required"`
	Created            int64    `json:"created,omitempty"`
	Modified           int64    `json:"modified,omitempty"`
	Description        string   `json:"description,omitempty"`
	LastConnected      int64    `json:"lastConnected,omitempty"`
	LastReported       int64    `json:"lastReported,omitempty"`
	OperatingState     string   `json:"operatingState" validate:"oneof='ENABLED' 'DISABLED'"`
	Labels             []string `json:"labels,omitempty"`
	BaseAddress        string   `json:"baseAddress" validate:"required,uri"`
	AdminState         string   `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
}

// UpdateDeviceService represents a service that is responsible for proxying connectivity between a set of devices and the
// EdgeX Foundry core services.
// This object and its properties correspond to the UpdateDeviceService object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/UpdateDeviceService
type UpdateDeviceService struct {
	Id             *string  `json:"id" validate:"required_without=Name"`
	Name           *string  `json:"name" validate:"required_without=Id"`
	BaseAddress    *string  `json:"baseAddress"`
	OperatingState *string  `json:"operatingState" validate:"omitempty,oneof='ENABLED' 'DISABLED'"`
	Labels         []string `json:"labels"`
	AdminState     *string  `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
}
