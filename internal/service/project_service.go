package service

import (
	"Outsourcing/cmd/gorm/model"
	"Outsourcing/cmd/gorm/query"
	"Outsourcing/internal/repository/dao"
	"Outsourcing/model/dto"
	"Outsourcing/pkg/config"
	"Outsourcing/pkg/util"
	"fmt"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

var ProjectService = &projectService{}

type projectService struct{}

// SendTask
//
//	@Description:	项目发布
//	@receiver		userService
//	@parameter		projectParam
//	@return			result 0代表失败 1代表成功
//	@return			err 失败返回err 成功返回nil
func (projectService) SendTask(projectParam *dto.ProjectPublishParam, user *model.User) (result int, err error) {

	zap.S().Info("接受到的参数为", projectParam)

	q := query.Use(dao.DB)

	err = q.Transaction(func(tx *query.Query) error {
		//操作project file
		project := &model.Project{
			ID:                  projectParam.ProjectId,
			ProjectPublisherID:  user.ID,
			ProjectName:         projectParam.ProjectName,
			ProjectLogo:         projectParam.ProjectLogo,
			ProjectBudget:       projectParam.ProjectBudget.InexactFloat64(),
			ProjectShortIntro:   projectParam.ProjectShortIntro,
			ProjectIntroduction: projectParam.ProjectIntroduction,
			ProjectDuration:     projectParam.ProjectDuration,
			WorkExplain:         projectParam.WorkExplain,
			StartTime:           projectParam.StartTime,
			EndTime:             projectParam.EndTime,
			CreateTime:          time.Now(),
			UpdateTime:          time.Now(),
			ScheduleID:          1,
			IsCompanyPublish:    int32(projectParam.IsCompanyPublish),
			WorkType:            int32(projectParam.WorkType),
			DeleteTag:           0,
			Approved:            0,
			Approver:            0,
		}
		if err := tx.Project.Create(project); err != nil {
			zap.Error(err)
			return err
		}

		//操作project position type relation--position type
		//创建切片插入project position type relation
		var positionTypeRelations []*model.ProjectPositionTypeRelation

		//构建切片 数据插入切片
		for _, v := range projectParam.EngineerTypeIdList {
			member := model.ProjectPositionTypeRelation{
				ID:             0,
				ProjectID:      project.ID,
				PositionTypeID: 0,
			}
			member.PositionTypeID = v
			positionTypeRelations = append(positionTypeRelations, &member)
		}
		//批量插入
		if err = tx.ProjectPositionTypeRelation.Create(positionTypeRelations...); err != nil {
			zap.Error(err)
			return err
		}

		//操作project schedule relation--schedule
		if err = tx.ProjectScheduleRelation.Create(&model.ProjectScheduleRelation{
			ID:         0,
			ProjectID:  project.ID,
			ScheduleID: 1,
			CreateTime: time.Now(),
		}); err != nil {
			zap.Error(err)
			return err
		}

		//操作project type relation--project type
		var projectTypeRelations []*model.ProjectTypeRelation
		//构建切片 数据插入切片
		for _, v := range projectParam.ProjectTypeIdList {
			member := model.ProjectTypeRelation{
				ID:            0,
				ProjectID:     project.ID,
				ProjectTypeID: 0,
			}
			member.ProjectTypeID = v
			projectTypeRelations = append(projectTypeRelations, &member)
		}
		//批量插入
		if err = tx.ProjectTypeRelation.Create(projectTypeRelations...); err != nil {
			zap.Error(err)
			return err
		}

		//project files--file type--
		var projectFile []*model.ProjectFile
		for _, file := range projectParam.FileList {
			//改写文件的信息
			fileName := fmt.Sprintf(
				"%s-%s%s",
				user.Username, //放用户名
				time.Now().Format(config.DigitDateTimeLayout),
				filepath.Ext(file.Filename),
			)

			// 文件保存路径
			Path := config.UploadsPath + fileName
			// 保存文件
			err = util.SaveUploadedFile(file, Path)

			member := model.ProjectFile{
				ID:         0,
				ProjectID:  project.ID,
				UserID:     user.ID,
				FileTypeID: 0,
				FileName:   fileName,
				FileURL:    Path,
				CreateTime: time.Now(),
			}
			ext := filepath.Ext(file.Filename)
			if ext == "png" || ext == "JPG" {
				member.FileTypeID = 2
			} else {
				member.FileTypeID = 1
			}
			projectFile = append(projectFile, &member)
		}
		if err = tx.ProjectFile.Create(projectFile...); err != nil {
			zap.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}
	return 1, nil

}

// QueryTask
//
//	@Description: 查询符合条件的任务
//	@receiver projectService
//	@param projectParam
//	@return find
func (projectService) QueryTask(projectParam *dto.ProjectQueryParam) (find []*model.Project) {
	zap.S().Info("开始查询项目 条件为：", projectParam)

	p := query.Project
	u := query.User
	s := query.Schedule
	pt := query.ProjectType
	ptr := query.ProjectTypeRelation
	pptr := query.ProjectPositionTypeRelation
	pt1 := query.PositionType
	where := p.Select(p.ID, p.ProjectName, p.ProjectPublisherID, u.Username.As("project_publisher_name"),
		u.Phone.As("phone"), p.ProjectBudget, p.ProjectLogo, p.ProjectShortIntro, p.ProjectIntroduction,
		p.ProjectDuration, p.ProjectDuration, p.WorkExplain, p.CreateTime, p.UpdateTime, p.WorkType,
		p.ScheduleID.As("project_schedule_id"), s.ScheduleName.As("project_schedule_name")).
		LeftJoin(s, p.ScheduleID.EqCol(s.ID)).
		LeftJoin(u, p.ProjectPublisherID.EqCol(u.ID)).
		LeftJoin(ptr, ptr.ProjectID.EqCol(p.ID)).
		LeftJoin(pptr, pptr.ProjectID.EqCol(p.ID)).
		LeftJoin(pt1, pt1.ID.EqCol(pptr.PositionTypeID)).
		Where(p.DeleteTag.Eq(0), p.ScheduleID.Lt(4), p.Approved.Eq(1))

	zap.S().Info(where.Find())

	if projectParam.ID != 0 {
		where.Where(p.ID.Eq(projectParam.ID))
	}

	if projectParam.ProjectPublisherID != 0 {
		where.Where(p.ProjectPublisherID.Eq(projectParam.ProjectPublisherID))
	}
	if projectParam.MinProjectBudget != 0 {
		where.Where(p.ProjectBudget.Gte(float64(projectParam.MinProjectBudget)))
	}
	if projectParam.MaxProjectBudget != 0 {
		where.Where(p.ProjectBudget.Lte(float64(projectParam.MaxProjectBudget)))
	}
	if projectParam.ProjectDuration != "" {
		where.Where(p.ProjectDuration.Eq(projectParam.ProjectDuration))
	}
	if projectParam.CreateTimeBegin.After(time.Time{}) {
		where.Where(p.CreateTime.Gte(projectParam.CreateTimeBegin))
	}
	if projectParam.CreateTimeEnd.After(time.Time{}) {
		where.Where(p.CreateTime.Lte(projectParam.CreateTimeEnd))
	}
	if projectParam.UpdateTimeBegin.After(time.Time{}) {
		where.Where(p.UpdateTime.Gte(projectParam.UpdateTimeBegin))
	}
	if projectParam.UpdateTimeEnd.After(time.Time{}) {
		where.Where(p.UpdateTime.Lte(projectParam.UpdateTimeEnd))
	}
	if projectParam.ProjectScheduleID != 0 {
		where.Where(p.ScheduleID.Eq(projectParam.ProjectScheduleID))
	}
	if projectParam.WorkType != 0 {
		where.Where(p.WorkType.Eq(int32(projectParam.WorkType)))
	}
	if projectParam.ProjectPositionTypeList != nil && len(projectParam.ProjectPositionTypeList) > 0 {
		where.Where(pt1.ID.In(projectParam.ProjectPositionTypeList...))
	}
	if projectParam.ProjectTypeList != nil && len(projectParam.ProjectTypeList) > 0 {
		where.Where(pt.ID.In(projectParam.ProjectTypeList...))
	}

	order := where.Group(p.ID).Order()

	if projectParam.OrderType == 0 && projectParam.OrderWay == 0 {
		order.Order(p.ProjectBudget)
	} else if projectParam.OrderType == 0 && projectParam.OrderWay == 1 {
		order.Order(p.ProjectBudget.Desc())
	} else if projectParam.OrderType == 1 && projectParam.OrderWay == 0 {
		order.Order(p.CreateTime)
	} else if projectParam.OrderType == 1 && projectParam.OrderWay == 1 {
		order.Order(p.CreateTime.Desc())
	}

	find, err := order.Find()

	if err != nil {
		zap.S().Error(err)
	}

	return find

}
