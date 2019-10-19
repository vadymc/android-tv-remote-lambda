package main

import (
	alexa "github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	startSqs()
	lambda.Start(Handler)
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return dispatch(request), nil
}

func dispatch(request alexa.Request) alexa.Response {
	switch request.Body.Type {
	case "LaunchRequest":
		return DispatchLaunchRequest(request)
	case "IntentRequest":
		return DispatchIntentRequest(request)
	default:
		return alexa.NewSimpleResponse("Unknown request", "unknown request")
	}
}

func DispatchLaunchRequest(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Launch request", "opened session")
}
func DispatchIntentRequest(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "VolumeUp":
		response = handleVolumeUp(request)
	case "VolumeDown":
		response = handleVolumeDown(request)
	case "SetVolume":
		response = handleSetVolume(request)
	case "Mute":
		response = handleCommand("MUTE")
	case "Power":
		response = handleCommand("POWER")
	default:
		response = alexa.NewSimpleResponse("Unknown request", "cannot handle intent")
	}

	return response
}

func handleVolumeUp(request alexa.Request) alexa.Response {
	if val, exists := request.Body.Intent.Slots["level"]; exists {
		sendToSqs("VOLUME_UP " + val.Value)
	} else {
		sendToSqs("VOLUME_UP")
	}

	return alexa.NewSimpleResponse("Volume Up", "done")
}

func handleVolumeDown(request alexa.Request) alexa.Response {
	if val, exists := request.Body.Intent.Slots["level"]; exists {
		sendToSqs("VOLUME_DOWN " + val.Value)
	} else {
		sendToSqs("VOLUME_DOWN")
	}
	return alexa.NewSimpleResponse("Volume Down", "done")
}

func handleSetVolume(request alexa.Request) alexa.Response {
	if val, exists := request.Body.Intent.Slots["level"]; exists {
		sendToSqs("SET_VOLUME " + val.Value)
	}
	return alexa.NewSimpleResponse("Set Volume", "done")
}

func handleCommand(command string) alexa.Response {
	sendToSqs(command)
	return alexa.NewSimpleResponse(command, "done")
}
