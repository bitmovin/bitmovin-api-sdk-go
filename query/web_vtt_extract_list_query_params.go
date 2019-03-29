package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type WebVttExtractListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *WebVttExtractListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
