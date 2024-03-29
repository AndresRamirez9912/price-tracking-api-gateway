package models

type GetUserRequest struct {
	AccessToken string `json:"AccessToken"`
}

type generalResponse struct {
	Success      bool        `json:"success"`
	Response     interface{} `json:"response"`
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
}

type GetUserResponse struct {
	Response struct {
		MFAOptions          interface{} `json:"MFAOptions"`
		PreferredMfaSetting interface{} `json:"PreferredMfaSetting"`
		UserAttributes      []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"UserAttributes"`
		UserMFASettingList interface{} `json:"UserMFASettingList"`
		Username           string      `json:"Username"`
	} `json:"response"`
	generalResponse
}

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
}
