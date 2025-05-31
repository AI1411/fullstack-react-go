package main

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	applogger "github.com/AI1411/fullstack-react-go/internal/infra/logger"
)

func main() {
	ctx := context.Background()
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/domain/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler(
		db.DefaultDatabaseConfig(),
		applogger.New(applogger.DefaultConfig()),
	)
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn(ctx))

	// Generate the code
	g.Execute()

	// 生成したModelにRelation情報を手動追加（これだけは手動対応が必要）
	allModels := []interface{}{
		// ユーザーモデル
		g.GenerateModel(
			model.TableNameRegion,
			gen.FieldRelateModel(field.HasMany, "Prefectures", model.Prefecture{}, nil),
		),
		g.GenerateModel(
			model.TableNamePrefecture,
			gen.FieldRelateModel(field.HasMany, "Disasters", model.Disaster{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "Region", model.Region{}, nil),
		),
		g.GenerateModel(
			model.TableNameDisaster,
			gen.FieldRelateModel(field.BelongsTo, "Prefecture", model.Prefecture{}, nil),
			gen.FieldRelateModel(field.HasMany, "Timelines", model.Timeline{}, nil),
			gen.FieldRelateModel(field.HasMany, "DisasterDocuments", model.DisasterDocument{}, nil),
		),

		g.GenerateModel(model.TableNameUser),

		g.GenerateModel(
			model.TableNameTimeline,
			gen.FieldRelateModel(field.BelongsTo, "Disaster", model.Disaster{}, nil),
		),

		g.GenerateModel(
			model.TableNameDisasterDocument,
			gen.FieldRelateModel(field.BelongsTo, "Disaster", model.Disaster{}, nil),
		),

		g.GenerateModel(
			model.TableNameSupportApplication,
		),
		g.GenerateModel(
			model.TableNameDamageLevel,
		),
		g.GenerateModel(
			model.TableNameAssessment,
		),
		g.GenerateModel(
			model.TableNameAssessmentComment,
		),
		g.GenerateModel(
			model.TableNameAssessmentItem,
		),
		g.GenerateModel(
			model.TableNameFacilityType,
		),
		g.GenerateModel(
			model.TableNameGisDatum,
		),
		g.GenerateModel(
			model.TableNameNotification,
		),
		g.GenerateModel(
			model.TableNameRole,
		),
		g.GenerateModel(
			model.TableNameUserOrganization,
		),
	}

	g.ApplyBasic(allModels...)

	// Generate the code
	g.Execute()
}
