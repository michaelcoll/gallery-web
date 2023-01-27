/*
 * Copyright (c) 2022 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package presentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhotoApiController_ExtractRangeHeader(t *testing.T) {
	// Given

	// When
	start, end, err := extractRangeHeader("photo=0-25")

	// Then
	if err != nil {
		assert.Fail(t, "should be valid", err)
	}
	assert.Equal(t, 0, start, "start value not valid")
	assert.Equal(t, 25, end, "end value not valid")
}

func TestPhotoApiController_ExtractRangeHeader_NoEnd(t *testing.T) {
	// Given

	// When
	start, end, err := extractRangeHeader("photo=25-")

	// Then
	if err != nil {
		assert.Fail(t, "should be valid", err)
	}
	assert.Equal(t, 25, start, "start value not valid")
	assert.Equal(t, 0, end, "end value not valid")
}

func TestPhotoApiController_ExtractRangeHeader_InvalidUnit(t *testing.T) {
	// Given

	// When
	_, _, err := extractRangeHeader("truc=0-25")

	// Then
	assert.Errorf(t, err, "unit invalid")
}

func TestPhotoApiController_ExtractRangeHeader_InvalidRange(t *testing.T) {
	// Given

	// When
	_, _, err := extractRangeHeader("truc=25-5")

	// Then
	assert.Errorf(t, err, "range invalid")
}

func TestPhotoApiController_ExtractRangeHeader_InvalidStart(t *testing.T) {
	// Given

	// When
	_, _, err := extractRangeHeader("truc=a-5")

	// Then
	assert.Errorf(t, err, "range invalid")
}

func TestPhotoApiController_ExtractRangeHeader_InvalidEnd(t *testing.T) {
	// Given

	// When
	_, _, err := extractRangeHeader("truc=42-A")

	// Then
	assert.Errorf(t, err, "range invalid")
}
