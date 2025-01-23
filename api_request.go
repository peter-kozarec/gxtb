package gxtb

import (
	"encoding/json"
	"fmt"
)

type apiRequest struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
}

func marshalRequest(command string, arguments interface{}) (data []byte, err error) {

	request := apiRequest{Command: command, Arguments: nil}

	if arguments != nil {
		request.Arguments, err = json.Marshal(arguments)

		if err != nil {
			return nil, fmt.Errorf("unable to marshal request arguments: %w", err)
		}
	}

	data, err = json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request: %w", err)
	}

	return data, err
}
