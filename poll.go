package linego

import api "github.com/sakura-rip/linego/talkservice"

func (cl *LineClient) fetchOps() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOps(
		cl.ctx,
		cl.OperationValue.localRev,
		cl.OperationValue.count,
		cl.OperationValue.globalRev,
		cl.OperationValue.individualRev,
	)
	return res, err
}

func (cl *LineClient) fetchOperations() ([]*api.Operation, error) {
	res, err := cl.Poll.FetchOperations(
		cl.ctx,
		cl.OperationValue.localRev,
		cl.OperationValue.count,
	)
	return res, err
}

func (cl *LineClient) setRevision(rev int64) {
	if cl.OperationValue.localRev > rev {
		cl.OperationValue.localRev = rev
	}
}
