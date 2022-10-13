package request

import (
	"fmt"
	"github.com/peepoclown/golang-yandex-metrika/internal/constants"
	"strconv"
	"strings"
	"time"
)

type StatsRequest struct {
	CounterId          int64
	DateStart          time.Time
	DateFinish         time.Time
	GroupBy            string
	Metrics            []string
	Dimensions         []string
	IsIncludeUndefined bool
	IsProposedAccuracy bool
	Accuracy           int64
	Rows               []string
	Filters            string
	path               string
}

func (req *StatsRequest) GetPath() string {
	return constants.AppMetrikaApiStatsPath
}

func (req *StatsRequest) GetRequestParams() map[string]string {
	params := make(map[string]string)

	params["id"] = strconv.FormatInt(req.CounterId, 10)
	params["date1"] = req.DateStart.Format(constants.ApiRequestDateTimeFormat)
	params["date2"] = req.DateFinish.Format(constants.ApiRequestDateTimeFormat)
	params["group"] = req.GroupBy
	params["metrics"] = strings.Join(req.Metrics, ",")
	params["dimensions"] = strings.Join(req.Dimensions, ",")
	params["include_undefined"] = strconv.FormatBool(req.IsIncludeUndefined)
	params["proposedAccuracy"] = strconv.FormatBool(req.IsIncludeUndefined)
	params["accuracy"] = strconv.FormatInt(req.Accuracy, 10)

	if req.Rows != nil {
		rows := make([]string, 0, cap(req.Rows))
		for _, row := range req.Rows {
			rows = append(rows, fmt.Sprintf("[\"%s\"]", row))
		}
		params["row_ids"] = fmt.Sprintf("[%s]", strings.Join(rows, ","))
	}

	if req.Filters != "" {
		params["filters"] = req.Filters
	}

	return params
}
