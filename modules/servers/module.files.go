package servers

import (
	"github.com/jamee5e/jame-shop-tutorial/modules/files/filesHandlers"
	"github.com/jamee5e/jame-shop-tutorial/modules/files/filesUsecases"
)

type IFilesModule interface {
	Init()
	Usecase() filesUsecases.IFilesUsecase
	Handler() filesHandlers.IFilesHandler
}

type filesModule struct {
	*moduleFactory
	usecase filesUsecases.IFilesUsecase
	handler filesHandlers.IFilesHandler
}

func (m *moduleFactory) FilesModule() IFilesModule {
	usecase := filesUsecases.FilesUsecase(m.server.cfg)
	handler := filesHandlers.FilesHandler(m.server.cfg, usecase)

	return &filesModule{
		moduleFactory: m,
		usecase:       usecase,
		handler:       handler,
	}
}

func (f *filesModule) Init() {
	router := f.router.Group("/files")
	router.Post("/upload", f.mid.JwtAuth(), f.mid.Authorize(2), f.handler.UploadFiles)
	router.Patch("/delete", f.mid.JwtAuth(), f.mid.Authorize(2), f.handler.DeleteFile)
}

func (f *filesModule) Usecase() filesUsecases.IFilesUsecase { return f.usecase }
func (f *filesModule) Handler() filesHandlers.IFilesHandler { return f.handler }
