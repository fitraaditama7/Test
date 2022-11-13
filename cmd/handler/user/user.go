package userhandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test/cmd/constant"
	"test/cmd/handler"
	"test/cmd/mapper"
	"test/cmd/model"
	"test/cmd/service"
	"test/pkg/responses"
	"test/pkg/router"
	"test/pkg/utils"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) handler.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	ctx, cancel := context.WithTimeout(ctx, constant.ContextTimeout)
	defer cancel()

	var body model.RegisterUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	err := body.Validate()
	if err != nil {
		log.Println(err)
		responses.Error(w, err)
		return
	}

	response, err := h.userService.Register(ctx, body)
	if err != nil {
		log.Println(err)
		responses.Error(w, err)
		return
	}

	responses.Success(w, response)
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	ctx, cancel := context.WithTimeout(ctx, constant.ContextTimeout)
	defer cancel()
	var queryParam = r.URL.Query()
	var offset = 0
	var limit = 10
	var orderBy = "name"
	var sortBy = "desc"
	var query = r.URL.Query().Get("q")
	var err error

	if queryParam.Get("of") != "" {
		offset, err = strconv.Atoi(queryParam.Get("of"))
		if err != nil {
			log.Println(err)
			responses.Error(w, constant.InvalidQueryOffsetAndLimitTypeData)
			return
		}
	}

	if queryParam.Get("lt") != "" {
		limit, err = strconv.Atoi(queryParam.Get("lt"))
		if err != nil {
			log.Println(err)
			responses.Error(w, constant.InvalidQueryOffsetAndLimitTypeData)
			return
		}
	}

	if utils.Contains(constant.QueryOrderBy, queryParam.Get("ob")) {
		orderBy = queryParam.Get("ob")
	}

	if utils.Contains(constant.QuerySortBy, queryParam.Get("sb")) {
		sortBy = queryParam.Get("sb")
	}

	commonParam := mapper.ToCommonParam(offset, limit, query, orderBy, sortBy)

	response, err := h.userService.List(ctx, commonParam)
	if err != nil {
		log.Println(err)
		responses.Error(w, err)
		return
	}
	responses.Success(w, response)
}

func (h *userHandler) Detail(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	ctx, cancel := context.WithTimeout(ctx, constant.ContextTimeout)
	defer cancel()

	var userID = router.Param(ctx, "user_id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Println(err)
		responses.Error(w, constant.InvalidUserID)
		return
	}

	user, err := h.userService.Detail(ctx, id)
	if err != nil {
		log.Println(err)
		responses.Error(w, err)
		return
	}
	responses.Success(w, user)
}
