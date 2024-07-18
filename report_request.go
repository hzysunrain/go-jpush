package jpush

type ReportStatusRequest struct {
	MsgId           int      `json:"msg_id"`
	RegistrationIds []string `json:"registration_ids"`
	Date            string   `json:"date,omitempty"`
}
