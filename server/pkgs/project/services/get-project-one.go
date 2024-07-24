package projects

// imports
import (
	"context"
	"fmt"
	graphqlClient "server/clients/graphql"
	authModel "server/pkgs/auth/models"
	projectModel "server/pkgs/project/models"
)

func GetProject(id string) (projectModel.Project, error) {
	var query struct {
		Projects_by_pk struct {
			ID              string `json:"id"`
			Name            string `json:"name"`
			Project_manager struct {
				ID           string `json:"id"`
				First_name   string `json:"first_name"`
				Last_name    string `json:"last_name"`
				Email        string `json:"email"`
				Password     string `json:"password"`
				Phone_number string `json:"phone_number"`
			} `json:"project_manager"`
			Technical_project_manager struct {
				ID           string `json:"id"`
				First_name   string `json:"first_name"`
				Last_name    string `json:"last_name"`
				Email        string `json:"email"`
				Password     string `json:"password"`
				Phone_number string `json:"phone_number"`
			} `json:"technical_project_manager"`
		} `graphql:"projects_by_pk(id: $id)"`
	}

	type uuid string
	variables := map[string]interface{}{
		"id": uuid(id),
	}
	err := graphqlClient.SystemClient().Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err.Error(), "when querying project")
		return projectModel.Project{}, err
	}

	project := projectModel.Project{

		ID:   query.Projects_by_pk.ID,
		Name: query.Projects_by_pk.Name,
		Project_manager: authModel.User{
			ID:          query.Projects_by_pk.Project_manager.ID,
			FirstName:   query.Projects_by_pk.Project_manager.First_name,
			LastName:    query.Projects_by_pk.Project_manager.Last_name,
			Email:       query.Projects_by_pk.Project_manager.Email,
			PhoneNumber: query.Projects_by_pk.Project_manager.Phone_number,
		},
		Technical_project_manager: authModel.User{
			ID:          query.Projects_by_pk.Technical_project_manager.ID,
			FirstName:   query.Projects_by_pk.Technical_project_manager.First_name,
			LastName:    query.Projects_by_pk.Technical_project_manager.Last_name,
			Email:       query.Projects_by_pk.Technical_project_manager.Email,
			PhoneNumber: query.Projects_by_pk.Technical_project_manager.Phone_number,
		},
	}
	return project, nil
}
