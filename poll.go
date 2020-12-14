package linego

import (
	api "github.com/sakura-rip/linego/talkservice"
	"strconv"
	"strings"
)

func (cl *LineClient) fetchOps() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOps(
		cl.ctx,
		cl.SaveData.LastRevision,
		cl.SaveData.Count,
		cl.SaveData.GlobalRev,
		cl.SaveData.IndividualRev,
	)
	return res, err
}

func (cl *LineClient) fetchOperations() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOperations(
		cl.ctx,
		cl.SaveData.LastRevision,
		cl.SaveData.Count,
	)
	return res, err
}

func (cl *LineClient) setRevision(rev int64) {
	if cl.OperationValue.localRev > rev {
		cl.OperationValue.localRev = rev
	}
}
