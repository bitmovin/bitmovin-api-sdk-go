package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type CencDrmListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

func (q *CencDrmListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
