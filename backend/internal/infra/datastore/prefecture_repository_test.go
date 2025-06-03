package datastore_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
	"github.com/AI1411/fullstack-react-go/internal/infra/db"
	"github.com/AI1411/fullstack-react-go/tests/testutils"
)

func TestPrefectureRepository_Find(t *testing.T) {
	tests := []struct {
		name    string
		want    []*model.Prefecture
		wantErr bool
		setup   func(t *testing.T, client db.Client)
		useMock bool
	}{
		{
			name: "Success",
			want: []*model.Prefecture{
				{
					ID:   13,
					Name: "東京都",
					Code: "23",
				},
				{
					ID:   27,
					Name: "大阪府",
					Code: "27",
				},
			},
			setup: func(t *testing.T, client db.Client) {
				client.Conn(context.Background()).Exec(
					"INSERT INTO prefectures (id, name, code) VALUES (13, '東京都', '23'), (27, '大阪府', '27')",
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)

			client := testutils.SetupTestDB(t)
			defer client.Close()

			repo := datastore.NewPrefectureRepository(ctx, client)

			testutils.TruncateAllTables(t, client)

			if tt.setup != nil {
				tt.setup(t, client)
			}

			got, err := repo.Find(ctx)
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)

			if tt.want != nil {
				if !cmp.Equal(tt.want, got) {
					t.Errorf("diff %s", cmp.Diff(tt.want, got))
				}
			}
		})
	}
}

func TestPrefectureRepository_FindByID(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    *model.Prefecture
		wantErr bool
		setup   func(t *testing.T, client db.Client)
	}{
		{
			name: "Success/isActive trueのみ取得",
			code: "23",
			want: &model.Prefecture{
				ID:   13,
				Name: "東京都",
				Code: "23",
				Municipalities: []model.Municipality{
					{
						ID:                    1,
						PrefectureCode:        "23",
						OrganizationCode:      "230001",
						PrefectureNameKanji:   "東京都",
						MunicipalityNameKanji: "千代田区",
						PrefectureNameKana:    "ﾄｳｷｮｳﾄ",
						MunicipalityNameKana:  "ﾁﾖﾀﾞｸ",
						IsActive:              true,
					},
					{
						ID:                    2,
						PrefectureCode:        "23",
						OrganizationCode:      "230002",
						PrefectureNameKanji:   "東京都",
						MunicipalityNameKanji: "港区",
						PrefectureNameKana:    "ﾄｳｷｮｳﾄ",
						MunicipalityNameKana:  "ﾐﾅﾄｸ",
						IsActive:              true,
					},
				},
			},
			setup: func(t *testing.T, client db.Client) {
				r := require.New(t)
				client.Conn(context.Background()).Exec(
					"INSERT INTO prefectures (id, name, code) VALUES (13, '東京都', '23')",
				)
				r.NoError(client.Conn(context.Background()).Exec(
					"INSERT INTO municipalities (prefecture_code, organization_code, prefecture_name_kanji, municipality_name_kanji, prefecture_name_kana, municipality_name_kana) VALUES (?, ?, ?, ?, ?, ?)",
					"23", "230001", "東京都", "千代田区", "ﾄｳｷｮｳﾄ", "ﾁﾖﾀﾞｸ").Error,
				)
				r.NoError(client.Conn(context.Background()).Exec(
					"INSERT INTO municipalities (prefecture_code, organization_code, prefecture_name_kanji, municipality_name_kanji, prefecture_name_kana, municipality_name_kana) VALUES (?, ?, ?, ?, ?, ?)",
					"23", "230002", "東京都", "港区", "ﾄｳｷｮｳﾄ", "ﾐﾅﾄｸ").Error,
				)
				r.NoError(client.Conn(context.Background()).Exec(
					"INSERT INTO municipalities (prefecture_code, organization_code, prefecture_name_kanji, municipality_name_kanji, prefecture_name_kana, municipality_name_kana, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)",
					"23", "230003", "東京都", "三鷹市", "ﾄｳｷｮｳﾄ", "ﾐﾀｶｼ", false).Error,
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			a := assert.New(t)

			client := testutils.SetupTestDB(t)
			defer client.Close()

			repo := datastore.NewPrefectureRepository(ctx, client)

			testutils.TruncateAllTables(t, client)

			if tt.setup != nil {
				tt.setup(t, client)
			}

			got, err := repo.FindByID(ctx, tt.code)
			if tt.wantErr {
				a.Error(err)
				return
			}
			a.NoError(err)

			if !cmp.Equal(tt.want, got) {
				t.Errorf("diff %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
