package controller

type App struct {
	Posts interface{ PostController }
}
