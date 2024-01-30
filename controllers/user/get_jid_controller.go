package user

import (
	"github.com/Mauricio-Carrion/LeadMeBackend/models/user"
)

func GetJidController(userUuid string) (string, error) {
		data, err := user.GetJidModel(userUuid)	

		if err != nil {
			return "", err
		}
		
		return data.Jid, nil
}	