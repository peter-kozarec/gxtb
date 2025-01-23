package gxtb

import (
	"encoding/json"
	"fmt"
)

type apiResponse struct {
	Status     bool            `json:"status"`
	ReturnData json.RawMessage `json:"returnData,omitempty"`
	ErrorCode  string          `json:"errorCode,omitempty"`
	ErrorDescr string          `json:"errorDescr,omitempty"`
}

func unmarshalResponse(data []byte, returnData interface{}) error {

	var response apiResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return fmt.Errorf("unable to unmarshal response: %w", err)
	}

	if !response.Status {
		return fmt.Errorf("%s - %s", response.ErrorCode, response.ErrorDescr)
	}

	if returnData != nil {
		if err := json.Unmarshal(response.ReturnData, returnData); err != nil {
			return fmt.Errorf("unable to unmarshal returnData:: %w", err)
		}
	}

	return nil
}
