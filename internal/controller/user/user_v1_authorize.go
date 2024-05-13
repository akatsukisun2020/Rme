package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"rme/api/user/v1"
)

func (c *ControllerV1) Authorize(ctx context.Context, req *v1.AuthorizeReq) (res *v1.AuthorizeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
