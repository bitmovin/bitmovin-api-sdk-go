package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type PrimeTimeDrmListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

func (q *PrimeTimeDrmListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
