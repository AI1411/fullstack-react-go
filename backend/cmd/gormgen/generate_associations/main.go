package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/domain/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler(
		"myuser",
		"mypassword",
		"localhost",
		"5432",
		"m_app",
		applogger.New(applogger.DefaultConfig()),
	)
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn)

	// Generate the code
	g.Execute()

	// 生成したModelにRelation情報を手動追加（これだけは手動対応が必要）
	allModels := []interface{}{
		// ユーザーモデル
		g.GenerateModel(
			model.TableNameUser,
			gen.FieldRelateModel(field.HasMany, "UserInterests", model.UserInterest{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "Prefecture", model.Prefecture{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "Education", model.Education{}, nil),
			gen.FieldRelateModel(field.HasMany, "Communities", model.Community{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"CreatorID"},
					"references": {"ID"},
				},
			}),
			gen.FieldRelateModel(field.HasMany, "CommunityMembers", model.CommunityMember{}, nil),
			gen.FieldRelateModel(field.HasMany, "Likes", model.Like{}, nil),
			gen.FieldRelateModel(field.HasMany, "Matches", model.Match{}, nil),
			gen.FieldRelateModel(field.HasMany, "Reports", model.Report{}, nil),
			gen.FieldRelateModel(field.HasMany, "Footprints", model.Footprint{}, nil),
			gen.FieldRelateModel(field.HasMany, "UserImages", model.UserImage{}, nil),
			gen.FieldRelateModel(field.HasMany, "Tweets", model.Tweet{}, nil),
			gen.FieldRelateModel(field.HasMany, "Blocklists", model.Blocklist{}, nil),
		),
		// ユーザー興味モデル
		g.GenerateModel(
			model.TableNameUserInterest,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "Interest", model.Interest{}, nil),
		),
		// 地域モデル
		g.GenerateModel(
			model.TableNameRegion,
			gen.FieldRelateModel(field.HasMany, "Prefectures", model.Prefecture{}, nil),
		),
		// 都道府県モデル
		g.GenerateModel(
			model.TableNamePrefecture,
			gen.FieldRelateModel(field.BelongsTo, "Region", model.Region{}, nil),
			gen.FieldRelateModel(field.HasMany, "Users", model.User{}, nil),
		),
		// カテゴリーモデル
		g.GenerateModel(
			model.TableNameCategory,
			gen.FieldRelateModel(field.HasMany, "Interests", model.Interest{}, nil),
		),
		// 興味モデル
		g.GenerateModel(
			model.TableNameInterest,
			gen.FieldRelateModel(field.BelongsTo, "Category", model.Category{}, nil),
			gen.FieldRelateModel(field.HasMany, "UserInterests", model.UserInterest{}, nil),
		),
		// コミュニティモデル
		g.GenerateModel(
			model.TableNameCommunity,
			gen.FieldRelateModel(field.Many2Many, "CommunityMembers", model.CommunityMember{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"CreatorID"},
					"references": {"ID"},
				},
			}),
		),
		// コミュニティメンバーモデル
		g.GenerateModel(
			model.TableNameCommunityMember,
			gen.FieldRelateModel(field.BelongsTo, "Community", model.Community{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
		),
		// 通知モデル
		g.GenerateModel(
			model.TableNameNotification,
			gen.FieldRelateModel(field.HasMany, "Users", model.User{}, nil),
		),
		// 学歴モデル
		g.GenerateModel(
			model.TableNameEducation,
			gen.FieldRelateModel(field.HasMany, "Users", model.User{}, nil),
		),
		// いいねモデル
		g.GenerateModel(
			model.TableNameLike,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "LikedUser", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"LikedUserID"},
					"references": {"ID"},
				},
			}),
		),
		// マッチモデル
		g.GenerateModel(
			model.TableNameMatch,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "MatchedUser", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"MatchedUserID"},
					"references": {"ID"},
				},
			}),
		),
		// 通報モデル
		g.GenerateModel(
			model.TableNameReport,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "ReportedUser", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"ReportedUserID"},
					"references": {"ID"},
				},
			}),
		),
		// フットプリントモデル
		g.GenerateModel(
			model.TableNameFootprint,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "VisitedUser", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"VisitedUserID"},
					"references": {"ID"},
				},
			}),
		),
		// ユーザー画像モデル
		g.GenerateModel(
			model.TableNameUserImage,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
		),
		// ツイートモデル
		g.GenerateModel(
			model.TableNameTweet,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
		),
		// ブロックリストモデル
		g.GenerateModel(
			model.TableNameBlocklist,
			gen.FieldRelateModel(field.BelongsTo, "User", model.User{}, nil),
			gen.FieldRelateModel(field.BelongsTo, "BlockedUser", model.User{}, &field.RelateConfig{
				GORMTag: map[string][]string{
					"foreignKey": {"BlockedUserID"},
					"references": {"ID"},
				},
			}),
		),
	}

	g.ApplyBasic(allModels...)

	// Generate the code
	g.Execute()
}
