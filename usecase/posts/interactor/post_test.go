package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/ilmimris/learn-unittest/domain/model"
	"github.com/ilmimris/learn-unittest/domain/service"
	presenter "github.com/ilmimris/learn-unittest/shared/mock/presenter"
	repository "github.com/ilmimris/learn-unittest/shared/mock/repository"
)

func TestGetPosts(t *testing.T) {
	Convey("Test Get all posts", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoPost := repository.NewMockPostRepository(ctrl)
		repoAuthor := repository.NewMockAuthorRepository(ctrl)
		presenterPost := presenter.NewMockPostPresenter(ctrl)

		pi := NewPostInteractor(presenterPost, repoPost, repoAuthor)

		var posts []*model.Post
		posts = append(posts, &model.Post{
			ID:       1,
			Title:    "Title 1",
			Content:  "Content 1",
			AuthorID: int64(1),
		})
		posts = append(posts, &model.Post{
			ID:       2,
			Title:    "Title 2",
			Content:  "Content 2",
			AuthorID: int64(1),
		})

		author := &model.Author{
			ID:   int64(1),
			Name: "John Doe",
		}

		var resultsPosts []*service.Post
		resultsPosts = append(resultsPosts, &service.Post{
			ID:      1,
			Title:   "Title 1",
			Content: "Content 1",
			Author:  author.Name,
		})
		resultsPosts = append(resultsPosts, &service.Post{
			ID:      2,
			Title:   "Title 2",
			Content: "Content 2",
			Author:  author.Name,
		})

		Convey("Negative Scenario", func() {
			Convey("Should return error when fetch posts from db", func() {
				repoPost.EXPECT().FetchPosts(gomock.Any()).Return(nil, errors.New("error"))

				res, err := pi.GetPosts(context.Background())
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
			Convey("Should return error when fetch author from db for each post", func() {
				repoPost.EXPECT().FetchPosts(gomock.Any()).Return(posts, nil)
				repoAuthor.EXPECT().FetchAuthorByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

				res, err := pi.GetPosts(context.Background())
				So(err, ShouldNotBeNil)
				So(res, ShouldBeNil)
			})
		})
		Convey("Positive Scenario", func() {
			Convey("Should return all post", func() {
				repoPost.EXPECT().FetchPosts(gomock.Any()).Return(posts, nil)
				repoAuthor.EXPECT().FetchAuthorByID(gomock.Any(), gomock.Any()).Return(author, nil).AnyTimes()
				presenterPost.EXPECT().ResponsePosts(gomock.Any(), gomock.Any()).Return(resultsPosts, nil)

				res, err := pi.GetPosts(context.Background())
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
	})
}
