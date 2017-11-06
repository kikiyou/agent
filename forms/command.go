package forms

type CommandForm struct {
	Path    string `form:"path" json:"path"`
	Command string `form:"command" json:"command"`
	Html    string `form:"html" json:"html" `
	Token   string `form:"token" json:"token" binding:"required"`
	CommandID string `form:"commandID" json:"commandID" `
}
