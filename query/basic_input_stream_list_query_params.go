package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type BasicInputStreamListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *BasicInputStreamListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
