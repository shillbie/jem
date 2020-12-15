package linego

import (
	api "github.com/shillbie/jem/talkservice"
	"strconv"
	"strings"
)

func (cl *LineClient) FetchOps() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOps(
		cl.ctx,
		cl.SaveData.LastRevision,
		cl.SaveData.Count,
		cl.SaveData.GlobalRev,
		cl.SaveData.IndividualRev,
	)
	return res, err
}

func (cl *LineClient) FetchOperations() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOperations(
		cl.ctx,
		cl.SaveData.LastRevision,
		cl.SaveData.Count,
	)
	return res, err
}

func (cl *LineClient) SetRevision(rev int64) {
	if cl.SaveData.LastRevision < rev {
		cl.SaveData.LastRevision = rev
	}
}

func (cl *LineClient) SetIndividualRev(op *api.Operation) {
	if op.Param1 != "" {
		sps := strings.Split(op.Param1, "")
		if len(sps) != 0 {
			res, err := strconv.ParseInt(sps[0], 10, 64)
			if err == nil {
				cl.SaveData.IndividualRev = res
			}
		}
	}
}

func (cl *LineClient) SetGlobalRev(op *api.Operation) {
	if op.Param2 != "" {
		sps := strings.Split(op.Param2, "")
		if len(sps) != 0 {
			res, err := strconv.ParseInt(sps[0], 10, 64)
			if err == nil {
				cl.SaveData.GlobalRev = res
			}
		}
	}
}
