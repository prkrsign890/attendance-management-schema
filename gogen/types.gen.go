// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package Openapi

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for EmployeeGender.
const (
	EmployeeGenderFEMALE EmployeeGender = "FEMALE"

	EmployeeGenderMALE EmployeeGender = "MALE"

	EmployeeGenderNEITHER EmployeeGender = "NEITHER"

	EmployeeGenderNOTANSWER EmployeeGender = "NOT_ANSWER"
)

// Defines values for EmployeePosition.
const (
	EmployeePositionFULLTIME EmployeePosition = "FULL_TIME"

	EmployeePositionPARTTIME EmployeePosition = "PART_TIME"
)

// Employee defines model for Employee.
type Employee struct {

	// 誕生日
	// ISO 8601のYYYY-MM-DD形式
	Birthday openapi_types.Date `json:"birthday"`

	// 名前
	FirstName string `json:"firstName"`

	// 性別
	Gender EmployeeGender `json:"gender"`

	// 苗字
	LastName string `json:"lastName"`

	// 役割
	Position EmployeePosition `json:"position"`
}

// 性別
type EmployeeGender string

// 役割
type EmployeePosition string

// Employees defines model for Employees.
type Employees []Employee

// PostEmployeeRequest defines model for PostEmployeeRequest.
type PostEmployeeRequest Employee

// PostApiEmployeeJSONBody defines parameters for PostApiEmployee.
type PostApiEmployeeJSONBody PostEmployeeRequest

// PostApiEmployeeJSONRequestBody defines body for PostApiEmployee for application/json ContentType.
type PostApiEmployeeJSONRequestBody PostApiEmployeeJSONBody

