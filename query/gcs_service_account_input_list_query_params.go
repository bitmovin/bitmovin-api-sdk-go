package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type GcsServiceAccountInputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *GcsServiceAccountInputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
