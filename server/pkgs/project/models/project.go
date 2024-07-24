package project_models

import (
	authModel "server/pkgs/auth/models"
)

type Project struct {
	ID                        string         `json:"id"`
	Name                      string         `json:"name"`
	Project_manager           authModel.User `json:"project_manager"`
	Technical_project_manager authModel.User `json:"technical_project_manager"`
}
