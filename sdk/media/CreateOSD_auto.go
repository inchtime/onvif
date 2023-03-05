// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/inchtime/onvif"
	"github.com/inchtime/onvif/sdk"
	"github.com/inchtime/onvif/media"
)

// Call_CreateOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateOSDResponse.
func Call_CreateOSD(ctx context.Context, dev *onvif.Device, request media.CreateOSD) (media.CreateOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateOSDResponse media.CreateOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateOSDResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateOSD")
		return reply.Body.CreateOSDResponse, errors.Annotate(err, "reply")
	}
}
