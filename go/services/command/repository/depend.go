package repository

import (
	"go.uber.org/fx"
)

var ReqDepend = fx.Options(
	fx.Provide(
		NewCategoryRepositorySqlBoiler,
		NewProductRepositorySQLBoiler,
	),
)
