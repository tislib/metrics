// Code generated by apbr generate. DO NOT EDIT.
// versions:
// 	apbr generate v1.2

//go:build !codeanalysis

package model

import "github.com/google/uuid"

type TestResource struct {
	Id          *uuid.UUID `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Version     int32      `json:"version,omitempty"`
}

func (s TestResource) GetId() *uuid.UUID {
	return s.Id
}
func (s TestResource) GetName() *string {
	return s.Name
}
func (s TestResource) GetDescription() *string {
	return s.Description
}
func (s TestResource) GetVersion() int32 {
	return s.Version
}
