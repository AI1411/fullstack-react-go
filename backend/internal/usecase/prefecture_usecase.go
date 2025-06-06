//go:generate mockgen -source=prefecture_usecase.go -destination=../../tests/mock/usecase/prefecture_usecase.mock.go
package usecase

import (
	"context"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	domain "github.com/AI1411/fullstack-react-go/internal/domain/repository"
)

type PrefectureUseCase interface {
	ListPrefectures(ctx context.Context) ([]*model.Prefecture, error)
	GetPrefectureByID(ctx context.Context, code string) (*model.Prefecture, error)
}

type prefectureUseCase struct {
	prefectureRepository domain.PrefectureRepository
}

func NewPrefectureUseCase(
	prefectureRepository domain.PrefectureRepository,
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

func (u *prefectureUseCase) GetPrefectureByID(ctx context.Context, code string) (*model.Prefecture, error) {
	prefecture, err := u.prefectureRepository.FindByID(ctx, code)
	if err != nil {
		return nil, err
	}

	return prefecture, nil
}
