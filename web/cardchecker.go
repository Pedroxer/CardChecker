package web

import (
	"github.com/Pedroxer/CardChecker/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cardCheckerRequest struct {
	Card string `json:"card" binding:"required"`
}

func (serv *Server) validCheck(ctx *gin.Context) {
	var req cardCheckerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	valid := util.LuhnAlgorithm(req.Card)
	if valid {
		ctx.JSON(http.StatusOK, "Card is valid")
		return
	}
	ctx.JSON(http.StatusOK, "Card is not valid")
}
