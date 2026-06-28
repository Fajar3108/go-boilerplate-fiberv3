package helpers

type APIResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Meta       any    `json:"meta,omitempty"`
	Details    any    `json:"details,omitempty"`
	ErrorCode  string `json:"error_code,omitempty"`
	Trace      any    `json:"trace,omitempty"`
}

func NewAPIResponse(statusCode int, message string) *APIResponse {
	return &APIResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (ar *APIResponse) Success(data, meta, details any) *APIResponse {
	ar.Data = data
	ar.Meta = meta
	ar.Details = details

	return ar
}

func (ar *APIResponse) Error(errorCode string, details, trace any) *APIResponse {
	ar.ErrorCode = errorCode
	ar.Details = details
	ar.Trace = trace

	return ar
}
