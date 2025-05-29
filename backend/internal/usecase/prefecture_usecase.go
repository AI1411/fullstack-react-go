package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type PrefectureUseCase interface {
	ListPrefectures(ctx context.Context) ([]*model.Prefecture, error)
	GetPrefectureByID(ctx context.Context, id int32) (*model.Prefecture, error)
}

type prefectureUseCase struct {
	prefectureRepository datastore.PrefectureRepository
}

func NewPrefectureUseCase(
	prefectureRepository datastore.PrefectureRepository,
) PrefectureUseCase {
	return &prefectureUseCase{
		prefectureRepository: prefectureRepository,
	}
}

func (u *prefectureUseCase) ListPrefectures(ctx context.Context) ([]*model.Prefecture, error) {
	prefectures, err := u.prefectureRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return prefectures, nil
}

func (u *prefectureUseCase) GetPrefectureByID(ctx context.Context, id int32) (*model.Prefecture, error) {
	prefecture, err := u.prefectureRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return prefecture, nil
}
