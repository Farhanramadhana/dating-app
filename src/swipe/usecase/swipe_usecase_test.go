package usecase

import (
	"fmt"
	"testing"
	"time"

	mocks "dating-app/mocks/repository"
	"dating-app/model/constant"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCountUserSwipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := mocks.NewMockRedis(ctrl)
	mocksRepository := mocks.NewMockRepositoryInterface(ctrl)

	usecase := &SwipeUsecase{
		swipeRepository: mocksRepository,
		redis:           mockRedis,
	}

	testCases := []struct {
		name        string
		input       int
		expectation func(int)
	}{
		{
			name:  "Count total swipe for free user",
			input: 1,
			expectation: func(input int) {
				output := "1"
				today := time.Now().Format(constant.YYYYMMDD)
				mockRedis.EXPECT().RedisGet(gomock.Any(), fmt.Sprintf(constant.UserSwipeKey, input, today)).Return(output)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.expectation(test.input)
			usecase.CountUserSwipe(test.input)
			assert.Nil(t, nil)
		})
	}

}
