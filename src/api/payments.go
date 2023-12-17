package api

import (
	"database/sql"
	"net/http"

	db "github.com/Chien179/NMCBookstoreBE/src/db/sqlc"
	"github.com/Chien179/NMCBookstoreBE/src/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) createPayment(ctx *gin.Context, PaymentID string, OrderID int64, ToAddress string, TotalShipping float64, SubTotal float64, Status string, email string) {
	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	shipping, err := server.createShipping(ctx, ToAddress, TotalShipping)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	order, err := server.store.GetOrder(ctx, OrderID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if Status != "failed" {
		arg := db.UpdateOrderParams{
			ID: order.ID,
			Status: sql.NullString{
				String: "paid",
				Valid:  true,
			},
		}

		_, err := server.store.UpdateOrder(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		user, err := server.store.GetUserByEmail(ctx, email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		userArg := db.UpdateUserParams{
			Rank: sql.NullInt32{
				Int32: user.Rank + int32(SubTotal),
				Valid: true,
			},
		}

		server.store.UpdateUser(ctx, userArg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	arg := db.CreatePaymentParams{
		ID:         PaymentID,
		Username:   authPayLoad.Username,
		OrderID:    order.ID,
		ShippingID: shipping.ID,
		Subtotal:   order.SubTotal,
		Status:     Status,
	}

	_, err = server.store.CreatePayment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
