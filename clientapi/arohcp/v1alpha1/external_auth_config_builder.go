/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// ExternalAuthConfigBuilder contains the data and logic needed to build 'external_auth_config' objects.
//
// Represents an external authentication configuration
type ExternalAuthConfigBuilder struct {
	bitmap_       uint32
	id            string
	href          string
	externalAuths *ExternalAuthListBuilder
	state         ExternalAuthConfigState
	enabled       bool
}

// NewExternalAuthConfig creates a new builder of 'external_auth_config' objects.
func NewExternalAuthConfig() *ExternalAuthConfigBuilder {
	return &ExternalAuthConfigBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *ExternalAuthConfigBuilder) Link(value bool) *ExternalAuthConfigBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *ExternalAuthConfigBuilder) ID(value string) *ExternalAuthConfigBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *ExternalAuthConfigBuilder) HREF(value string) *ExternalAuthConfigBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *ExternalAuthConfigBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// Enabled sets the value of the 'enabled' attribute to the given value.
func (b *ExternalAuthConfigBuilder) Enabled(value bool) *ExternalAuthConfigBuilder {
	b.enabled = value
	b.bitmap_ |= 8
	return b
}

// ExternalAuths sets the value of the 'external_auths' attribute to the given values.
func (b *ExternalAuthConfigBuilder) ExternalAuths(value *ExternalAuthListBuilder) *ExternalAuthConfigBuilder {
	b.externalAuths = value
	b.bitmap_ |= 16
	return b
}

// State sets the value of the 'state' attribute to the given value.
//
// Representation of the possible values for the state field of an external authentication configuration
func (b *ExternalAuthConfigBuilder) State(value ExternalAuthConfigState) *ExternalAuthConfigBuilder {
	b.state = value
	b.bitmap_ |= 32
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ExternalAuthConfigBuilder) Copy(object *ExternalAuthConfig) *ExternalAuthConfigBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.enabled = object.enabled
	if object.externalAuths != nil {
		b.externalAuths = NewExternalAuthList().Copy(object.externalAuths)
	} else {
		b.externalAuths = nil
	}
	b.state = object.state
	return b
}

// Build creates a 'external_auth_config' object using the configuration stored in the builder.
func (b *ExternalAuthConfigBuilder) Build() (object *ExternalAuthConfig, err error) {
	object = new(ExternalAuthConfig)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.enabled = b.enabled
	if b.externalAuths != nil {
		object.externalAuths, err = b.externalAuths.Build()
		if err != nil {
			return
		}
	}
	object.state = b.state
	return
}
