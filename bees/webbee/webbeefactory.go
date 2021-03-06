/*
 *    Copyright (C) 2014 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package webbee

import (
	"github.com/muesli/beehive/bees"
)

type WebBeeFactory struct {
	bees.BeeFactory
}

// Interface impl

func (factory *WebBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := WebBee{
		Bee:  bees.NewBee(name, factory.Name(), description),
		addr: options.GetValue("addr").(string),
		path: options.GetValue("path").(string),
	}

	return &bee
}

func (factory *WebBeeFactory) Name() string {
	return "webbee"
}

func (factory *WebBeeFactory) Description() string {
	return "A RESTful HTTP module for beehive"
}

func (factory *WebBeeFactory) Image() string {
	return factory.Name() + ".png"
}

func (factory *WebBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		bees.BeeOptionDescriptor{
			Name:        "addr",
			Description: "Which addr to listen on, eg: 0.0.0.0:12345",
			Type:        "string",
			Mandatory:   true,
		},
		bees.BeeOptionDescriptor{
			Name:        "path",
			Description: "Which path to expect GET/POST requests on, eg: /foobar",
			Type:        "string",
			Mandatory:   true,
		},
	}
	return opts
}

func (factory *WebBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		bees.EventDescriptor{
			Namespace:   factory.Name(),
			Name:        "post",
			Description: "A POST call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				bees.PlaceholderDescriptor{
					Name:        "json",
					Description: "JSON map received from caller",
					Type:        "map",
				},
				bees.PlaceholderDescriptor{
					Name:        "ip",
					Description: "IP of the caller",
					Type:        "string",
				},
			},
		},
		bees.EventDescriptor{
			Namespace:   factory.Name(),
			Name:        "get",
			Description: "A GET call was received by the HTTP server",
			Options: []bees.PlaceholderDescriptor{
				bees.PlaceholderDescriptor{
					Name:        "query_params",
					Description: "Map of query parameters received from caller",
					Type:        "map",
				},
				bees.PlaceholderDescriptor{
					Name:        "ip",
					Description: "IP of the caller",
					Type:        "string",
				},
			},
		},
	}
	return events
}

func (factory *WebBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		bees.ActionDescriptor{
			Namespace:   factory.Name(),
			Name:        "post",
			Description: "Does a POST request",
			Options: []bees.PlaceholderDescriptor{
				bees.PlaceholderDescriptor{
					Name:        "json",
					Description: "Data to send",
					Type:        "string",
				},
				bees.PlaceholderDescriptor{
					Name:        "url",
					Description: "Where to connect to",
					Type:        "string",
				},
			},
		},
	}
	return actions
}

func init() {
	f := WebBeeFactory{}
	bees.RegisterFactory(&f)
}
