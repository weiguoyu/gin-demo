// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package resource_controller

import (
	"gin-demo/pkg/logger"
	"gin-demo/pkg/models/dbschema"
	"gin-demo/pkg/models/project"
	mysql "gin-demo/pkg/util/db"
)

func GetProjects(inputs project.APIGetProjectsInput) (total uint32, projects []project.Project, err error) {
	db := mysql.GetMysqlInstance()

	projects = make([]project.Project, 0)
	total, err = db.New().
		Select(project.ProjectColumns...).
		From(dbschema.TableProject).
		Where(mysql.BuildFilterConditions(inputs, dbschema.TableProject)).
		CountAndLoad(&projects)
	if err != nil {
		logger.Errorf("get projects error, %v", err)
		return
	}
	return
}
