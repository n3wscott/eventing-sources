/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"log"

	"github.com/knative/eventing-sources/pkg/kncloudevents"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"

	"github.com/gorilla/websocket"
)

var (
	sink   string
	source string

	// CloudEvents specific parameters
	eventType   string
	eventSource string
)

func init() {
	flag.StringVar(&sink, "sink", "", "the host url to send messages to")
	flag.StringVar(&source, "source", "", "the url to get messages from")
	flag.StringVar(&eventType, "eventType", "websocket-event", "the event-type (CloudEvents)")
	flag.StringVar(&eventSource, "eventSource", "", "the event-source (CloudEvents)")
}

func main() {
	flag.Parse()

	// "source" flag must not be empty for operation.
	if source == "" {
		log.Fatal("A valid source url must be defined.")
	}

	// The event's source defaults to the URL of where it was taken from.
	if eventSource == "" {
		eventSource = source
	}

	ce, err := kncloudevents.NewDefaultClient(sink)
	if err != nil {
		log.Fatalf("Failed to create a http cloudevent client: %s", err.Error())
	}

	ws, _, err := websocket.DefaultDialer.Dial(source, nil)
	if err != nil {
		log.Fatal("error connecting:", err)
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			// TODO(markusthoemmes): Handle failures and reconnect
			log.Println("error while reading message:", err)
			return
		}

		event := cloudevents.Event{
			Context: cloudevents.EventContextV02{
				Type:   eventType,
				Source: *types.ParseURLRef(eventSource),
			}.AsV02(),
			Data: message,
		}
		if _, err := ce.Send(context.TODO(), event); err != nil {
			log.Printf("sending event to channel failed: %v", err)
		}
	}
}
