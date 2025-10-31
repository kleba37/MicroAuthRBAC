package MessageResponses

type AuthOperationResponse struct {
	Status  int    `json:"statusCode"`
	Access  bool   `json:"isAccess"`
	Message string `json:"message"`
}
