package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @ID RetrieveFile
// @Summary Get content of a file
// @Tags File
// @Accept json
// @Produce octet-stream
// @Param id path int true "File ID"
// @Param Range header string false "HTTP Range Header"
// @Success 200 {file} file
// @Success 206 {file} file
// @Failure 500 {object} api.HTTPError
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /file/{id}/retrieve [get]
func (s *Server) retrieveFile(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseUint(c.ParamValues()[0], 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, HTTPError{Err: "failed to parse path parameter as number"})
	}
	data, name, modTime, err := s.fileHandler.RetrieveFileHandler(ctx, s.db.WithContext(ctx), s.retriever, id)
	if err != nil {
		return httpResponseFromError(c, err)
	}
	c.Response().Header().Add("Content-Type", "application/octet-stream")
	http.ServeContent(c.Response(), c.Request(), name, modTime, data)
	return data.Close()
}
