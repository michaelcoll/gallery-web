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

package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

func TestDaemonService_HeartBeat(t *testing.T) {
	// Given
	daemon := givenDaemon("testDaemon1", false)
	service := NewDaemonService(nil)
	service.daemons[daemon.Id] = daemon

	// When
	_ = service.HeartBeat(daemon)

	// Then
	assert.Equal(t, true, daemon.Alive, "should be alive")

}

func TestDaemonService_List(t *testing.T) {
	// Given
	daemon0 := givenDaemon("testDaemon0", false)
	daemon1 := givenDaemon("testDaemon1", false)
	daemon2 := givenDaemon("testDaemon2", false)
	service := NewDaemonService(nil)
	service.daemons[daemon0.Id] = daemon0
	service.daemons[daemon1.Id] = daemon1
	service.daemons[daemon2.Id] = daemon2

	// When
	list := service.List("me")

	// Then
	assert.Contains(t, list, daemon0, "should contain testDaemon0")
	assert.Contains(t, list, daemon1, "should contain testDaemon1")
	assert.Contains(t, list, daemon2, "should contain testDaemon2")
}

func TestDaemonService_ById(t *testing.T) {
	// Given
	daemon0 := givenDaemon("testDaemon0", true)
	daemon1 := givenDaemon("testDaemon1", false)
	service := NewDaemonService(nil)
	service.daemons[daemon0.Id] = daemon0
	service.daemons[daemon1.Id] = daemon1

	// When
	daemon, err := service.ById(daemon0.Id)
	if err != nil {
		assert.Fail(t, "should get a daemon (%v)", err)
	}

	// Then
	assert.Equal(t, daemon0, daemon, "should be testDaemon0")
}

func TestDaemonService_ById_Inactive(t *testing.T) {
	// Given
	daemon0 := givenDaemon("testDaemon0", true)
	daemon1 := givenDaemon("testDaemon1", false)
	service := NewDaemonService(nil)
	service.daemons[daemon0.Id] = daemon0
	service.daemons[daemon1.Id] = daemon1

	// When
	_, err := service.ById(daemon1.Id)

	// Then
	assert.Errorf(t, err, "daemon not active")
}

func givenDaemon(name string, alive bool) *model.Daemon {
	return &model.Daemon{
		Id:    uuid.New(),
		Name:  name,
		Alive: alive,
		Owner: "me",
	}
}
