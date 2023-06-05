package mallard

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func checkBypass(ctx *fiber.Ctx, opts *Opts) error {
	apiKey := ctx.GetReqHeaders()[opts.bypassHeader]
	if apiKey != opts.apiKey {
		return errors.New("invalid or missing api key")
	}

	return nil
}
