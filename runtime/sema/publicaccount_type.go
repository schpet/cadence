/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import (
	"github.com/onflow/cadence/runtime/common"
)

const PublicAccountTypeName = "PublicAccount"
const PublicAccountAddressField = "address"
const PublicAccountBalanceField = "balance"
const PublicAccountAvailableBalanceField = "availableBalance"
const PublicAccountStorageUsedField = "storageUsed"
const PublicAccountStorageCapacityField = "storageCapacity"
const PublicAccountGetCapabilityField = "getCapability"
const PublicAccountGetTargetLinkField = "getLinkTarget"
const PublicAccountKeysField = "keys"

// PublicAccountType represents the publicly accessible portion of an account.
//
var PublicAccountType = func() *CompositeType {

	publicAccountType := &CompositeType{
		Identifier:         PublicAccountTypeName,
		Kind:               common.CompositeKindStructure,
		hasComputedMembers: true,
		importable:         false,

		nestedTypes: func() *StringTypeOrderedMap {
			nestedTypes := NewStringTypeOrderedMap()
			nestedTypes.Set(AccountKeysTypeName, PublicAccountKeysType)
			return nestedTypes
		}(),
	}

	var members = []*Member{
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountAddressField,
			&AddressType{},
			accountTypeAddressFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountBalanceField,
			UFix64Type,
			accountTypeAccountBalanceFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountAvailableBalanceField,
			UFix64Type,
			accountTypeAccountAvailableBalanceFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountStorageUsedField,
			UInt64Type,
			accountTypeStorageUsedFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountStorageCapacityField,
			UInt64Type,
			accountTypeStorageCapacityFieldDocString,
		),
		NewPublicFunctionMember(
			publicAccountType,
			PublicAccountGetCapabilityField,
			publicAccountTypeGetCapabilityFunctionType,
			publicAccountTypeGetLinkTargetFunctionDocString,
		),
		NewPublicFunctionMember(
			publicAccountType,
			PublicAccountGetTargetLinkField,
			accountTypeGetLinkTargetFunctionType,
			accountTypeGetLinkTargetFunctionDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			PublicAccountKeysField,
			PublicAccountKeysType,
			accountTypeKeysFieldDocString,
		),
	}

	publicAccountType.Members = GetMembersAsMap(members)
	publicAccountType.Fields = getFieldNames(members)
	return publicAccountType
}()

// PublicAccountKeysType represents the keys associated with a public account.
var PublicAccountKeysType = func() *CompositeType {

	accountKeys := &CompositeType{
		Identifier: AccountKeysTypeName,
		Kind:       common.CompositeKindStructure,
		importable: false,
	}

	var members = []*Member{
		NewPublicFunctionMember(
			accountKeys,
			AccountKeysGetFunctionName,
			accountKeysTypeGetFunctionType,
			accountKeysTypeGetFunctionDocString,
		),
	}

	accountKeys.Members = GetMembersAsMap(members)
	accountKeys.Fields = getFieldNames(members)
	return accountKeys
}()

func init() {
	// Set the container type after initializing the AccountKeysTypes, to avoid initializing loop.
	PublicAccountKeysType.SetContainerType(PublicAccountType)
}

const publicAccountTypeGetLinkTargetFunctionDocString = `
Returns the capability at the given public path, or nil if it does not exist
`
