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
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

const (
	// number of seconds after a daemon is treated as not alive
	expiresIn = 3
	// number of seconds that will be added to the expiresIn value to calculate if a daemon is alive or not
	delta = 2
)

type DaemonService struct {
	caller  PhotoServiceCaller
	daemons map[uuid.UUID]*model.Daemon
	mu      sync.Mutex
}

func NewDaemonService(c PhotoServiceCaller) DaemonService {
	return DaemonService{
		caller:  c,
		daemons: make(map[uuid.UUID]*model.Daemon),
	}
}

func (s *DaemonService) Register(newDaemon *model.Daemon) (uuid.UUID, uint32, error) {
	d := s.findInExistingDaemon(newDaemon)

	if d.New {
		fmt.Printf("%s Registering a new daemon %s (%s) located at %s:%s...",
			color.GreenString("!"),
			color.GreenString(d.Name),
			color.GreenString(d.Version),
			color.CyanString(d.Hostname), color.CyanString(strconv.Itoa(d.Port)))
	}

	if s.validateDaemonConnection(d) {
		s.activate(d)

		s.daemons[d.Id] = d

		_, _ = color.New(color.FgGreen, color.Bold).Println(" ✓ OK")

		return d.Id, expiresIn, nil
	} else {

		_, _ = color.New(color.FgRed, color.Bold).Println(" ✗ KO")

		return [16]byte{}, 0, errors.New("could not establish connection to your daemon")
	}
}

func (s *DaemonService) findInExistingDaemon(d *model.Daemon) *model.Daemon {
	for _, daemon := range s.daemons {
		if d.Name == daemon.Name &&
			d.Owner == daemon.Owner &&
			d.Hostname == daemon.Hostname {
			daemon.Alive = false
			return daemon
		}
	}

	return d
}

func (s *DaemonService) HeartBeat(daemon *model.Daemon) error {
	currentDaemon, exists := s.daemons[daemon.Id]

	if exists {
		s.activate(currentDaemon)
	} else {
		_, _, err := s.Register(daemon)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DaemonService) activate(d *model.Daemon) {
	if !d.Alive {
		fmt.Printf("%s Daemon %s reconnected ...",
			color.GreenString("!"),
			color.GreenString(d.Name))
	}
	s.mu.Lock()
	d.LastSeen = time.Now()
	d.Alive = true
	s.mu.Unlock()
}

func (s *DaemonService) validateDaemonConnection(d *model.Daemon) bool {
	_, err := s.caller.Exists(context.Background(), d, "0")
	if err != nil {
		return false
	}

	return true
}

func (s *DaemonService) Watch() {
	for {
		for _, d := range s.daemons {

			if d.LastSeen.Add(time.Duration(expiresIn+delta)*time.Second).Before(time.Now()) && d.Alive {
				s.mu.Lock()
				d.Alive = false
				s.mu.Unlock()

				fmt.Printf("%s Daemon %s unregistered.\n",
					color.YellowString("!"),
					color.YellowString(d.Name))
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (s *DaemonService) List(owner string) []*model.Daemon {
	daemons := make([]*model.Daemon, 0)

	for _, daemon := range s.daemons {
		if owner == daemon.Owner {
			daemons = append(daemons, daemon)
		}
	}

	return daemons
}

func (s *DaemonService) ById(id uuid.UUID) (*model.Daemon, error) {
	daemon, exists := s.daemons[id]

	if exists {
		if !daemon.Alive {
			return nil, status.Error(codes.NotFound, "daemon not active")
		}
		return daemon, nil
	} else {
		return nil, status.Error(codes.NotFound, "daemon not found")
	}
}

func (s *DaemonService) ExistsById(id uuid.UUID) bool {
	daemon, exists := s.daemons[id]

	if exists {
		if !daemon.Alive {
			return false
		}
		return true
	} else {
		return false
	}
}
