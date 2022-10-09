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

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

const expiresIn = 2

type DaemonService struct {
	c       PhotoServiceCaller
	daemons map[uuid.UUID]*model.Daemon
	mu      sync.Mutex
}

func NewDaemonService(c PhotoServiceCaller) DaemonService {
	return DaemonService{
		c:       c,
		daemons: make(map[uuid.UUID]*model.Daemon),
	}
}

func (s *DaemonService) Register(d *model.Daemon) (uuid.UUID, int32, error) {
	fmt.Printf("%s Registering a new daemon %s (%s) located at %s:%s...",
		color.GreenString("❗"),
		color.GreenString(d.Name),
		color.GreenString(d.Version),
		color.CyanString(d.Hostname), color.CyanString(strconv.Itoa(d.Port)))

	if s.validateDaemonConnection(d) {
		s.activate(d)

		s.daemons[d.Id] = d

		fmt.Println(color.GreenString(" ✅ OK"))

		return d.Id, expiresIn, nil
	} else {

		fmt.Println(color.RedString(" ❌ KO"))

		return [16]byte{}, 0, errors.New("could not establish connection to your daemon")
	}
}

func (s *DaemonService) HeartBeat(id uuid.UUID) error {
	daemon, exists := s.daemons[id]

	if exists {
		s.activate(daemon)
	} else {
		fmt.Printf("%s Daemon %s not found...\n", color.RedString("❌"), color.RedString(id.String()))
		return errors.New("daemon not found")
	}

	return nil
}

func (s *DaemonService) activate(d *model.Daemon) {
	if !d.Alive {
		fmt.Printf("%s Daemon %s reconnected.\n",
			color.GreenString("❗"),
			color.GreenString(d.Name))
	}
	s.mu.Lock()
	d.NextSee = time.Now().Add(time.Duration(expiresIn) * time.Second)
	d.Alive = true
	s.mu.Unlock()
}

func (s *DaemonService) validateDaemonConnection(d *model.Daemon) bool {
	_, err := s.c.Exists(context.Background(), *d, "0")
	if err != nil {
		return false
	}

	return true
}

func (s *DaemonService) Watch() {
	for {
		for _, d := range s.daemons {

			if d.NextSee.Before(time.Now()) && d.Alive {
				s.mu.Lock()
				d.Alive = false
				s.mu.Unlock()

				fmt.Printf("%s Daemon %s unregistered.\n",
					color.YellowString("❗"),
					color.YellowString(d.Name))
			}
		}

		time.Sleep(1 * time.Second)
	}
}
