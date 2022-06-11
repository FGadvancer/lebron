// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v1/home/banner",
				Handler: HomeBannerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/flashsale",
				Handler: FlashSaleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/recommend",
				Handler: RecommendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/category/list",
				Handler: CategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/cart/list",
				Handler: CartListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/product/comment",
				Handler: ProductCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/order/list",
				Handler: OrderListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/product/detail",
				Handler: ProductDetailHandler(serverCtx),
			},
		},
	)
}
