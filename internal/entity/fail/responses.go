package fail

import "github.com/mephistolie/chefbook-backend-common/responses/fail"

var (
	GrpcNameLength = fail.CreateGrpcClient(fail.TypeInvalidBody, "maximum first & last name lengths is 64 symbols")

	GrpcDescriptionLength = fail.CreateGrpcClient(fail.TypeInvalidBody, "maximum description length is 150 symbols")
)
