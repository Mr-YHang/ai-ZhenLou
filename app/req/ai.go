package req

import "errors"

type TalkReq struct {
	UserID     int64  `json:"user_id"`
	DialogueID string `json:"dialogue_id"`
	Ask        string `json:"ask"`
}

func (TalkReq *TalkReq) Check() error {
	if TalkReq.UserID == 0 {
		return errors.New("用户信息异常，请您重新登录")
	}

	if len(TalkReq.Ask) == 0 {
		return errors.New("抱歉，我没能理解您的问题")
	}

	return nil
}
