package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/Chien179/NMCBookstoreBE/db/sqlc"
	"github.com/Chien179/NMCBookstoreBE/token"
	"github.com/gin-gonic/gin"
)

type createOrderRequest struct {
	Carts []db.Cart `json:"carts",binding:"required"`
}

// @Summary      Create order
// @Description  Use this API to create order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        Request body createOrderRequest  true  "Create order"
// @Success      200 {object} db.Order
// @failure	 	 400
// @failure		 500
// @Router       /users/orders [post]
func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayLoad == nil {
		err := errors.New("Login first")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	order, err := server.store.CreateOrder(ctx, authPayLoad.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, cart := range req.Carts {
		arg := db.CreateTransactionParams{
			OrdersID: order.ID,
			BooksID:  cart.BooksID,
			Amount:   cart.Amount,
			Total:    cart.Total,
		}
		_, err := server.store.CreateTransaction(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

	}

	ctx.JSON(http.StatusOK, order)
}

type deleteOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// @Summary      Delete order
// @Description  Use this API to delete order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        ID path int  true  "Delete order"
// @Success      200
// @failure	 	 400
// @failure	 	 401
// @failure	 	 404
// @failure		 500
// @Router       /users/orders/delete/{id} [delete]
func (server *Server) deleteOrder(ctx *gin.Context) {
	var req deleteOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := server.store.GetOrder(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if order.Username != authPayLoad.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = server.store.DeleteOrder(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Order deleted successfully")
}

type listOrderRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type listOrderResponse struct {
	ID    int64     `uri:"id"`
	Books []db.Book `json:"books"`
}

// @Summary      List order
// @Description  Use this API to list order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        Query query listOrderRequest  true  "List order"
// @Success      200 {object} []listOrderResponse
// @failure	 	 400
// @failure	 	 404
// @failure		 500
// @Router       /users/orders [get]
func (server *Server) listOrderPaid(ctx *gin.Context) {
	var req listOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListOdersByUserNameParams{
		Username: authPayLoad.Username,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	orders, err := server.store.ListOdersByUserName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := []listOrderResponse{}

	for _, order := range orders {
		if order.Status == "paid" {
			transactions, err := server.store.ListTransactionsByOrderID(ctx, order.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			books, err := server.listBookByTransactions(ctx, transactions)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			rsp = append(rsp, listOrderResponse{
				ID:    order.ID,
				Books: books,
			})
		}
	}

	ctx.JSON(http.StatusOK, rsp)
}
