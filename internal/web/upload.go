package web

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *Server) getUploadHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "multimedia.upload.tmpl", &uiPage{})
}

func (s *Server) postUploadHandler(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		s.l.Errorf("Upload torrent file failed: %s", err)
		displayError(ctx, http.StatusBadRequest, "Upload file failed")
		return
	}
	f, err := file.Open()
	if err != nil {
		s.l.Errorf("Upload torrent file failed: %s", err)
		displayError(ctx, http.StatusBadRequest, "Upload file failed")
		return
	}

	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		s.l.Errorf("Upload torrent file failed: %s", err)
		displayError(ctx, http.StatusBadRequest, "Upload file failed")
		return
	}

	if err = s.TorrentService.Add(buf); err != nil {
		s.l.Errorf("Add torrent failed %s", err)
		displayError(ctx, http.StatusInternalServerError, "Add torrent failed")
		return
	}
	displayOK(ctx, "Файл загружен", "/torrents")
}
