package response

import "github.com/wangrui19970405/wu-shi-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
