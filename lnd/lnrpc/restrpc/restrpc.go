package restrpc

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkt-cash/pktd/btcjson"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/connmgr/banmgr"
	"github.com/pkt-cash/pktd/lnd/chainreg"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/restrpc/help"
	"github.com/pkt-cash/pktd/lnd/lnrpc/routerrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/verrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/wtclientrpc"
	"github.com/pkt-cash/pktd/lnd/walletunlocker"
	"github.com/pkt-cash/pktd/neutrino"
	"github.com/pkt-cash/pktd/pktlog/log"
	"github.com/pkt-cash/pktd/pktwallet/waddrmgr"
	"github.com/pkt-cash/pktd/pktwallet/wallet"
)

type RpcFunc struct {
	command string
	req     proto.Message
	res     proto.Message
	f       func(c *RpcContext, m proto.Message) (proto.Message, er.R)
}

var rpcFunctions []RpcFunc = []RpcFunc{
	//	>>> lightning/channel subCategory commands

	//	service openchannel  -  URI /lightning/channel/open
	{
		command: help.CommandOpenChannel,
		req:     (*lnrpc.OpenChannelRequest)(nil),
		res:     (*lnrpc.ChannelPoint)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			openChannelReq, ok := m.(*lnrpc.OpenChannelRequest)
			if !ok {
				return nil, er.New("Argument is not a OpenChannelRequest")
			}

			//	open a channel
			cc, errr := c.withRpcServer()
			if cc != nil {
				var openChannelResp *lnrpc.ChannelPoint

				openChannelResp, err := cc.OpenChannelSync(context.TODO(), openChannelReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return openChannelResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service closechannel  -  URI /lightning/channel/close
	{
		command: help.CommandCloseChannel,
		req:     (*lnrpc.CloseChannelRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			closeChannelReq, ok := m.(*lnrpc.CloseChannelRequest)
			if !ok {
				return nil, er.New("Argument is not a CloseChannelRequest")
			}

			//	close a channel
			cc, errr := c.withRpcServer()
			if cc != nil {
				err := cc.CloseChannel(closeChannelReq, nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service abandonchannel  -  URI /lightning/channel/abandon
	{
		command: help.CommandAbandonChannel,
		req:     (*lnrpc.AbandonChannelRequest)(nil),
		res:     (*lnrpc.AbandonChannelResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			abandonChannelReq, ok := m.(*lnrpc.AbandonChannelRequest)
			if !ok {
				return nil, er.New("Argument is not a AbandonChannelRequest")
			}

			//	abandon a channel
			cc, errr := c.withRpcServer()
			if cc != nil {
				var abandonChannelResp *lnrpc.AbandonChannelResponse

				abandonChannelResp, err := cc.AbandonChannel(context.TODO(), abandonChannelReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return abandonChannelResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service channelbalance  -  URI /lightning/channel/balance
	{
		command: help.CommandChannelBalance,
		req:     nil,
		res:     (*lnrpc.ChannelBalanceResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the channel balance
			cc, errr := c.withRpcServer()
			if cc != nil {
				var channelBalanceResp *lnrpc.ChannelBalanceResponse

				channelBalanceResp, err := cc.ChannelBalance(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return channelBalanceResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service pendingchannels  -  URI /lightning/channel/pending
	{
		command: help.CommandPendingChannels,
		req:     nil,
		res:     (*lnrpc.PendingChannelsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get pending channels info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var pendingChannelsResp *lnrpc.PendingChannelsResponse

				pendingChannelsResp, err := cc.PendingChannels(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return pendingChannelsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service listchannels  -  URI /lightning/channel
	{
		command: help.CommandListChannels,
		req:     (*lnrpc.ListChannelsRequest)(nil),
		res:     (*lnrpc.ListChannelsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			listChannelsReq, ok := m.(*lnrpc.ListChannelsRequest)
			if !ok {
				return nil, er.New("Argument is not a ListChannelsRequest")
			}

			//	get a list of channels
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listChannelsResp *lnrpc.ListChannelsResponse

				listChannelsResp, err := cc.ListChannels(context.TODO(), listChannelsReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listChannelsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service closedchannels  -  URI /lightning/channel/closed
	{
		command: help.CommandClosedChannels,
		req:     (*lnrpc.ClosedChannelsRequest)(nil),
		res:     (*lnrpc.ClosedChannelsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			closeChannelReq, ok := m.(*lnrpc.ClosedChannelsRequest)
			if !ok {
				return nil, er.New("Argument is not a ClosedChannelRequest")
			}

			//	get a list of all closed channels
			cc, errr := c.withRpcServer()
			if cc != nil {
				var closedChannelsResp *lnrpc.ClosedChannelsResponse

				closedChannelsResp, err := cc.ClosedChannels(context.TODO(), closeChannelReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return closedChannelsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service getnetworkinfo  -  URI /lightning/channel/networkinfo
	{
		command: help.CommandGetNetworkInfo,
		req:     nil,
		res:     (*lnrpc.NetworkInfo)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get network info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var networkInfoResp *lnrpc.NetworkInfo

				networkInfoResp, err := cc.GetNetworkInfo(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return networkInfoResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service feereport  -  URI /lightning/channel/feereport
	{
		command: help.CommandFeeReport,
		req:     nil,
		res:     (*lnrpc.FeeReportResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get fee report
			cc, errr := c.withRpcServer()
			if cc != nil {
				var feeReportResp *lnrpc.FeeReportResponse

				feeReportResp, err := cc.FeeReport(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return feeReportResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service updatechanpolicy  -  URI /lightning/channel/policy
	{
		command: help.CommandUpdateChanPolicy,
		req:     (*lnrpc.PolicyUpdateRequest)(nil),
		res:     (*lnrpc.PolicyUpdateResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			policyUpdateReq, ok := m.(*lnrpc.PolicyUpdateRequest)
			if !ok {
				return nil, er.New("Argument is not a ClosedChannelRequest")
			}

			//	invoke Lightning update chan policy command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var policyUpdateResp *lnrpc.PolicyUpdateResponse

				policyUpdateResp, err := cc.UpdateChannelPolicy(context.TODO(), policyUpdateReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return policyUpdateResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> lightning/channel/backup subCategory commands

	//	service exportchanbackup  -  URI /lightning/channel/backup/export
	{
		command: help.CommandExportChanBackup,
		req:     (*lnrpc.ExportChannelBackupRequest)(nil),
		res:     (*lnrpc.ChannelBackup)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			exportChannelBackupReq, ok := m.(*lnrpc.ExportChannelBackupRequest)
			if !ok {
				return nil, er.New("Argument is not a ExportChannelBackupRequest")
			}

			//	invoke Lightning export chan backup command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var channelBackupResp *lnrpc.ChannelBackup

				channelBackupResp, err := cc.ExportChannelBackup(context.TODO(), exportChannelBackupReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return channelBackupResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service verifychanbackup  -  URI /lightning/channel/backup/verify
	{
		command: help.CommandVerifyChanBackup,
		req:     (*lnrpc.ChanBackupSnapshot)(nil),
		res:     (*lnrpc.VerifyChanBackupResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			chanBackupSnapshotReq, ok := m.(*lnrpc.ChanBackupSnapshot)
			if !ok {
				return nil, er.New("Argument is not a ChanBackupSnapshot")
			}

			//	invoke Lightning verify chan backup command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var verifyChanBackupResp *lnrpc.VerifyChanBackupResponse

				verifyChanBackupResp, err := cc.VerifyChanBackup(context.TODO(), chanBackupSnapshotReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return verifyChanBackupResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service restorechanbackup  -  URI /lightning/channel/backup/restore
	{
		command: help.CommandRestoreChanBackup,
		req:     (*lnrpc.RestoreChanBackupRequest)(nil),
		res:     (*lnrpc.RestoreBackupResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			restoreChanBackupReq, ok := m.(*lnrpc.RestoreChanBackupRequest)
			if !ok {
				return nil, er.New("Argument is not a RestoreChanBackupRequest")
			}

			//	invoke Lightning restore chan backup command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var restoreBackupResp *lnrpc.RestoreBackupResponse

				restoreBackupResp, err := cc.RestoreChannelBackups(context.TODO(), restoreChanBackupReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return restoreBackupResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> lightning/graph subCategory commands

	//	service describegraph  -  URI /lightning/graph
	{
		command: help.CommandDescribeGraph,
		req:     (*lnrpc.ChannelGraphRequest)(nil),
		res:     (*lnrpc.ChannelGraph)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			channelGraphReq, ok := m.(*lnrpc.ChannelGraphRequest)
			if !ok {
				return nil, er.New("Argument is not a ChannelGraphRequest")
			}

			//	get graph description info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var channelGraphResp *lnrpc.ChannelGraph

				channelGraphResp, err := cc.DescribeGraph(context.TODO(), channelGraphReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return channelGraphResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service getnodemetrics  -  URI /lightning/graph/nodemetrics
	{
		command: help.CommandGetNodeMetrics,
		req:     (*lnrpc.NodeMetricsRequest)(nil),
		res:     (*lnrpc.NodeMetricsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			nodeMetricsReq, ok := m.(*lnrpc.NodeMetricsRequest)
			if !ok {
				return nil, er.New("Argument is not a NodeMetricsRequest")
			}

			//	get node metrics info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var nodeMetricsResp *lnrpc.NodeMetricsResponse

				nodeMetricsResp, err := cc.GetNodeMetrics(context.TODO(), nodeMetricsReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return nodeMetricsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service getchaninfo  -  URI /lightning/graph/channel
	{
		command: help.CommandGetChanInfo,
		req:     (*lnrpc.ChanInfoRequest)(nil),
		res:     (*lnrpc.ChannelEdge)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			chanInfoReq, ok := m.(*lnrpc.ChanInfoRequest)
			if !ok {
				return nil, er.New("Argument is not a ChanInfoRequest")
			}

			//	get chan info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var channelEdgeResp *lnrpc.ChannelEdge

				channelEdgeResp, err := cc.GetChanInfo(context.TODO(), chanInfoReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return channelEdgeResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service getnodeinfo  -  URI /lightning/graph/nodeinfo
	{
		command: help.CommandGetNodeInfo,
		req:     (*lnrpc.NodeInfoRequest)(nil),
		res:     (*lnrpc.NodeInfo)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			nodeInfoReq, ok := m.(*lnrpc.NodeInfoRequest)
			if !ok {
				return nil, er.New("Argument is not a NodeInfoRequest")
			}

			//	get node info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var nodeInfoResp *lnrpc.NodeInfo

				nodeInfoResp, err := cc.GetNodeInfo(context.TODO(), nodeInfoReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return nodeInfoResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> lightning/invoice subCategory commands

	//	service addinvoice  -  URI /lightning/invoice/create
	{
		command: help.CommandAddInvoice,
		req:     (*lnrpc.Invoice)(nil),
		res:     (*lnrpc.AddInvoiceResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			invoiceReq, ok := m.(*lnrpc.Invoice)
			if !ok {
				return nil, er.New("Argument is not a Invoice")
			}

			//	add an invoice
			cc, errr := c.withRpcServer()
			if cc != nil {
				var addInvoiceResp *lnrpc.AddInvoiceResponse

				addInvoiceResp, err := cc.AddInvoice(context.TODO(), invoiceReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return addInvoiceResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service lookupinvoice  -  URI /lightning/invoice/lookup
	{
		command: help.CommandLookupInvoice,
		req:     (*lnrpc.PaymentHash)(nil),
		res:     (*lnrpc.Invoice)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			paymentHashReq, ok := m.(*lnrpc.PaymentHash)
			if !ok {
				return nil, er.New("Argument is not a PaymentHash")
			}

			//	lookup an invoice
			cc, errr := c.withRpcServer()
			if cc != nil {
				var InvoiceResp *lnrpc.Invoice

				InvoiceResp, err := cc.LookupInvoice(context.TODO(), paymentHashReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return InvoiceResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service listinvoices  -  URI /lightning/invoice
	{
		command: help.CommandListInvoices,
		req:     (*lnrpc.ListInvoiceRequest)(nil),
		res:     (*lnrpc.ListInvoiceResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			listInvoiceReq, ok := m.(*lnrpc.ListInvoiceRequest)
			if !ok {
				return nil, er.New("Argument is not a ListInvoiceRequest")
			}

			//	list all invoices
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listInvoiceResp *lnrpc.ListInvoiceResponse

				listInvoiceResp, err := cc.ListInvoices(context.TODO(), listInvoiceReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listInvoiceResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service decodepayreq  -  URI /lightning/invoice/decodepayreq
	{
		command: help.CommandDecodePayreq,
		req:     (*lnrpc.PayReqString)(nil),
		res:     (*lnrpc.PayReq)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			payReqStringReq, ok := m.(*lnrpc.PayReqString)
			if !ok {
				return nil, er.New("Argument is not a PayReqString")
			}

			//	decode payment request
			cc, errr := c.withRpcServer()
			if cc != nil {
				var payReqResp *lnrpc.PayReq

				payReqResp, err := cc.DecodePayReq(context.TODO(), payReqStringReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return payReqResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> lightning/payment subCategory command

	//	service sendpayment  -  URI /lightning/payment/send
	{
		command: help.CommandSendPayment,
		req:     (*lnrpc.SendRequest)(nil),
		res:     (*lnrpc.SendResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			sendReq, ok := m.(*lnrpc.SendRequest)
			if !ok {
				return nil, er.New("Argument is not a SendRequest")
			}

			//	invoke Lightning send payment command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var sendResp *lnrpc.SendResponse

				sendResp, err := cc.SendPaymentSync(context.TODO(), sendReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return sendResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service payinvoice  -  URI /lightning/payment/payinvoice
	{
		command: help.CommandPayInvoice,
		req:     (*routerrpc.SendPaymentRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			sendPaymentReq, ok := m.(*routerrpc.SendPaymentRequest)
			if !ok {
				return nil, er.New("Argument is not a SendPaymentRequest")
			}

			//	invoke Lightning send payment command
			cc, errr := c.withRouterServer()
			if cc != nil {
				err := cc.SendPaymentV2(sendPaymentReq, nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service sendtoroute  -  URI /lightning/payment/sendtoroute
	{
		command: help.CommandSendToRoute,
		req:     (*lnrpc.SendToRouteRequest)(nil),
		res:     (*lnrpc.SendResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			sendToRouteReq, ok := m.(*lnrpc.SendToRouteRequest)
			if !ok {
				return nil, er.New("Argument is not a SendToRouteRequest")
			}

			//	invoke Lightning send to route command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var sendResp *lnrpc.SendResponse

				sendResp, err := cc.SendToRouteSync(context.TODO(), sendToRouteReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return sendResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service listpayments  -  URI /lightning/payment
	{
		command: help.CommandListPayments,
		req:     (*lnrpc.ListPaymentsRequest)(nil),
		res:     (*lnrpc.ListPaymentsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			listPaymentsReq, ok := m.(*lnrpc.ListPaymentsRequest)
			if !ok {
				return nil, er.New("Argument is not a ListPaymentsRequest")
			}

			//	invoke Lightning list payments command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listPaymentsResp *lnrpc.ListPaymentsResponse

				listPaymentsResp, err := cc.ListPayments(context.TODO(), listPaymentsReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listPaymentsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	TODO: service trackpayment  -  URI /lightning/payment/???
	{
		command: help.CommandTrackPayment,
		req:     (*routerrpc.TrackPaymentRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			trackPaymentReq, ok := m.(*routerrpc.TrackPaymentRequest)
			if !ok {
				return nil, er.New("Argument is not a TrackPaymentRequest")
			}

			//	invoke Lightning send payment command
			cc, errr := c.withRouterServer()
			if cc != nil {
				err := cc.TrackPaymentV2(trackPaymentReq, nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service queryroutes  -  URI /lightning/payment/queryroutes
	{
		command: help.CommandQueryRoutes,
		req:     (*lnrpc.QueryRoutesRequest)(nil),
		res:     (*lnrpc.QueryRoutesResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			queryRoutesReq, ok := m.(*lnrpc.QueryRoutesRequest)
			if !ok {
				return nil, er.New("Argument is not a QueryRoutesRequest")
			}

			//	invoke Lightning query routes command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var queryRoutesResp *lnrpc.QueryRoutesResponse

				queryRoutesResp, err := cc.QueryRoutes(context.TODO(), queryRoutesReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return queryRoutesResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service fwdinghistory  -  URI /lightning/payment/fwdinghistory
	{
		command: help.CommandFwdingHistory,
		req:     (*lnrpc.ForwardingHistoryRequest)(nil),
		res:     (*lnrpc.ForwardingHistoryResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			forwardingHistoryReq, ok := m.(*lnrpc.ForwardingHistoryRequest)
			if !ok {
				return nil, er.New("Argument is not a ForwardingHistoryRequest")
			}

			//	invoke Lightning forwarding history command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var forwardingHistoryResp *lnrpc.ForwardingHistoryResponse

				forwardingHistoryResp, err := cc.ForwardingHistory(context.TODO(), forwardingHistoryReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return forwardingHistoryResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service querymc  -  URI /lightning/payment/querymc
	{
		command: help.CommandQueryMc,
		req:     nil,
		res:     (*routerrpc.QueryMissionControlResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	query the mission control status
			cc, errr := c.withRouterServer()
			if cc != nil {
				var queryMissionControlResp *routerrpc.QueryMissionControlResponse

				queryMissionControlResp, err := cc.QueryMissionControl(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return queryMissionControlResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service queryprob  -  URI /lightning/payment/queryprob
	{
		command: help.CommandQueryProb,
		req:     (*routerrpc.QueryProbabilityRequest)(nil),
		res:     (*routerrpc.QueryProbabilityResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			queryProbabilityReq, ok := m.(*routerrpc.QueryProbabilityRequest)
			if !ok {
				return nil, er.New("Argument is not a QueryProbabilityRequest")
			}

			//	invoke the probability service
			cc, errr := c.withRouterServer()
			if cc != nil {
				var queryProbabilityResp *routerrpc.QueryProbabilityResponse

				queryProbabilityResp, err := cc.QueryProbability(context.TODO(), queryProbabilityReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return queryProbabilityResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service resetmc  -  URI /lightning/payment/resetmc
	{
		command: help.CommandResetMc,
		req:     nil,
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	invoke reset mission controle service
			cc, errr := c.withRouterServer()
			if cc != nil {
				_, err := cc.ResetMissionControl(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service buildroute  -  URI /lightning/payment/buildroute
	{
		command: help.CommandBuildRoute,
		req:     (*routerrpc.BuildRouteRequest)(nil),
		res:     (*routerrpc.BuildRouteResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			buildRouteReq, ok := m.(*routerrpc.BuildRouteRequest)
			if !ok {
				return nil, er.New("Argument is not a BuildRouteRequest")
			}

			//	invoke reset mission controle service
			cc, errr := c.withRouterServer()
			if cc != nil {
				var buildRouteResp *routerrpc.BuildRouteResponse

				buildRouteResp, err := cc.BuildRoute(context.TODO(), buildRouteReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return buildRouteResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> lightning/peer subCategory command

	//	service connect  -  URI /lightning/peer/connect
	{
		command: help.CommandConnectPeer,
		req:     (*lnrpc.ConnectPeerRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			connectPeerReq, ok := m.(*lnrpc.ConnectPeerRequest)
			if !ok {
				return nil, er.New("Argument is not a ConnectPeerRequest")
			}

			//	invoke Lightning connect peer command
			cc, errr := c.withRpcServer()
			if cc != nil {
				_, err := cc.ConnectPeer(context.TODO(), connectPeerReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service disconnect  -  URI /lightning/peer/disconnect
	{
		command: help.CommandDisconnectPeer,
		req:     (*lnrpc.DisconnectPeerRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			disconnectPeerReq, ok := m.(*lnrpc.DisconnectPeerRequest)
			if !ok {
				return nil, er.New("Argument is not a DisconnectPeerRequest")
			}

			//	invoke Lightning disconnect peer command
			cc, errr := c.withRpcServer()
			if cc != nil {
				_, err := cc.DisconnectPeer(context.TODO(), disconnectPeerReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service listpeers  -  URI /lightning/peer
	{
		command: help.CommandListPeers,
		req:     nil,
		res:     (*lnrpc.ListPeersResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			listPeersReq := &lnrpc.ListPeersRequest{
				LatestError: true,
			}

			//	invoke Lightning list peers command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listPeersResp *lnrpc.ListPeersResponse

				listPeersResp, err := cc.ListPeers(context.TODO(), listPeersReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listPeersResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> meta category command

	//	service debug level  -  URI /meta/debuglevel
	{
		command: help.CommandDebugLevel,
		req:     (*lnrpc.DebugLevelRequest)(nil),
		res:     (*lnrpc.DebugLevelResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			debugLevelReq, ok := m.(*lnrpc.DebugLevelRequest)
			if !ok {
				return nil, er.New("Argument is not a DebugLevelRequest")
			}

			//	set Lightning debug level
			cc, errr := c.withRpcServer()
			if cc != nil {
				var debugLevelResp *lnrpc.DebugLevelResponse

				debugLevelResp, err := cc.DebugLevel(context.TODO(), debugLevelReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return debugLevelResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	MetaService get info  -  URI /meta/getinfo
	{
		command: help.CommandGetInfo,
		req:     nil,
		res:     (*lnrpc.GetInfo2Response)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			var ni lnrpc.NeutrinoInfo
			if n, _ := c.withNeutrino(); n != nil {
				neutrinoPeers := n.Peers()
				for i := range neutrinoPeers {
					var peerDesc lnrpc.PeerDesc
					neutrinoPeer := neutrinoPeers[i]

					peerDesc.BytesReceived = neutrinoPeer.BytesReceived()
					peerDesc.BytesSent = neutrinoPeer.BytesSent()
					peerDesc.LastRecv = neutrinoPeer.LastRecv().String()
					peerDesc.LastSend = neutrinoPeer.LastSend().String()
					peerDesc.Connected = neutrinoPeer.Connected()
					peerDesc.Addr = neutrinoPeer.Addr()
					peerDesc.Inbound = neutrinoPeer.Inbound()
					na := neutrinoPeer.NA()
					if na != nil {
						peerDesc.Na = na.IP.String() + ":" + strconv.Itoa(int(na.Port))
					}
					peerDesc.Id = neutrinoPeer.ID()
					peerDesc.UserAgent = neutrinoPeer.UserAgent()
					peerDesc.Services = neutrinoPeer.Services().String()
					peerDesc.VersionKnown = neutrinoPeer.VersionKnown()
					peerDesc.AdvertisedProtoVer = neutrinoPeer.Describe().AdvertisedProtoVer
					peerDesc.ProtocolVersion = neutrinoPeer.ProtocolVersion()
					peerDesc.SendHeadersPreferred = neutrinoPeer.Describe().SendHeadersPreferred
					peerDesc.VerAckReceived = neutrinoPeer.VerAckReceived()
					peerDesc.WitnessEnabled = neutrinoPeer.Describe().WitnessEnabled
					peerDesc.WireEncoding = strconv.Itoa(int(neutrinoPeer.Describe().WireEncoding))
					peerDesc.TimeOffset = neutrinoPeer.TimeOffset()
					peerDesc.TimeConnected = neutrinoPeer.Describe().TimeConnected.String()
					peerDesc.StartingHeight = neutrinoPeer.StartingHeight()
					peerDesc.LastBlock = neutrinoPeer.LastBlock()
					if neutrinoPeer.LastAnnouncedBlock() != nil {
						peerDesc.LastAnnouncedBlock = neutrinoPeer.LastAnnouncedBlock().CloneBytes()
					}
					peerDesc.LastPingNonce = neutrinoPeer.LastPingNonce()
					peerDesc.LastPingTime = neutrinoPeer.LastPingTime().String()
					peerDesc.LastPingMicros = neutrinoPeer.LastPingMicros()

					ni.Peers = append(ni.Peers, &peerDesc)
				}
				n.BanMgr().ForEachIp(func(bi banmgr.BanInfo) er.R {
					ban := lnrpc.NeutrinoBan{}
					ban.Addr = bi.Addr
					ban.Reason = bi.Reason
					ban.EndTime = bi.BanExpiresTime.String()
					ban.BanScore = bi.BanScore

					ni.Bans = append(ni.Bans, &ban)
					return nil
				})

				neutrionoQueries := n.GetActiveQueries()
				for i := range neutrionoQueries {
					nq := lnrpc.NeutrinoQuery{}
					query := neutrionoQueries[i]
					if query.Peer != nil {
						nq.Peer = query.Peer.String()
					} else {
						nq.Peer = "<nil>"
					}
					nq.Command = query.Command
					nq.ReqNum = query.ReqNum
					nq.CreateTime = query.CreateTime
					nq.LastRequestTime = query.LastRequestTime
					nq.LastResponseTime = query.LastResponseTime

					ni.Queries = append(ni.Queries, &nq)
				}

				bb, err := n.BestBlock()
				if err != nil {
					return nil, err
				}
				ni.BlockHash = bb.Hash.String()
				ni.Height = bb.Height
				ni.BlockTimestamp = bb.Timestamp.String()
				ni.IsSyncing = !n.IsCurrent()
			}

			var walletInfo *lnrpc.WalletInfo
			if w, _ := c.withWallet(); w != nil {
				mgrStamp := w.Manager.SyncedTo()
				walletStats := &lnrpc.WalletStats{}
				w.ReadStats(func(ws *btcjson.WalletStats) {
					walletStats.MaintenanceInProgress = ws.MaintenanceInProgress
					walletStats.MaintenanceName = ws.MaintenanceName
					walletStats.MaintenanceCycles = int32(ws.MaintenanceCycles)
					walletStats.MaintenanceLastBlockVisited = int32(ws.MaintenanceLastBlockVisited)
					walletStats.Syncing = ws.Syncing
					if ws.SyncStarted != nil {
						walletStats.SyncStarted = ws.SyncStarted.String()
					}
					walletStats.SyncRemainingSeconds = ws.SyncRemainingSeconds
					walletStats.SyncCurrentBlock = ws.SyncCurrentBlock
					walletStats.SyncFrom = ws.SyncFrom
					walletStats.SyncTo = ws.SyncTo
					walletStats.BirthdayBlock = ws.BirthdayBlock
				})
				walletInfo = &lnrpc.WalletInfo{
					CurrentBlockHash:      mgrStamp.Hash.String(),
					CurrentHeight:         mgrStamp.Height,
					CurrentBlockTimestamp: mgrStamp.Timestamp.String(),
					WalletVersion:         int32(waddrmgr.LatestMgrVersion),
					WalletStats:           walletStats,
				}
			}

			// Get Lightning info
			var lightning *lnrpc.GetInfoResponse
			if cc, _ := c.withRpcServer(); cc != nil {
				if l, err := cc.GetInfo(context.TODO(), nil); err != nil {
					return nil, er.E(err)
				} else {
					lightning = l
				}
			}

			return &lnrpc.GetInfo2Response{
				Neutrino:  &ni,
				Wallet:    walletInfo,
				Lightning: lightning,
			}, nil
		},
	},
	//	service to stop the pld daemon  -  URI /meta/stop
	{
		command: help.CommandStop,
		req:     nil,
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	invoke Lightning stop daemon command
			cc, errr := c.withRpcServer()
			if cc != nil {
				_, err := cc.StopDaemon(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service daemon version  -  URI /meta/version
	{
		command: help.CommandVersion,
		req:     nil,
		res:     (*verrpc.Version)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get daemon version
			cc, errr := c.withVerRPCServer()
			if cc != nil {
				var versionResp *verrpc.Version

				versionResp, err := cc.GetVersion(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return versionResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service force pld crash  -  URI /meta/crash
	{
		command: help.CommandCrash,
		req:     nil,
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	force a crash by defer a nil pointer
			meta, errr := c.withMetaServer()
			if meta != nil {
				meta.ForceCrash(context.TODO(), nil)

				return &RestEmptyResponse{}, nil
			} else {
				return nil, errr
			}
		},
	},

	//	>>> wallet category commands

	//	Wallet balance  -  URI /wallet/changepassphrase
	//	requires unlocked wallet -> access to rpcServer
	{
		command: help.CommandWalletBalance,
		req:     nil,
		res:     (*lnrpc.WalletBalanceResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.WalletBalance(context.TODO(), nil); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	MetaService change wallet password  -  URI /wallet/changepassphrase
	{
		command: help.CommandChangePassphrase,
		req:     (*lnrpc.ChangePasswordRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.ChangePasswordRequest)
			if !ok {
				return nil, er.New("Argument is not a ChangePasswordRequest")
			}
			meta, err := c.withMetaServer()
			if err != nil {
				return nil, err
			}
			_, errr := meta.ChangePassword(context.TODO(), req)
			if errr != nil {
				return nil, er.E(errr)
			}
			return &RestEmptyResponse{}, nil
		},
	},
	//	MetaService check wallet password  -  URI /wallet/checkpassphrase
	{
		command: help.CommandCheckPassphrase,
		req:     (*lnrpc.CheckPasswordRequest)(nil),
		res:     (*lnrpc.CheckPasswordResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			checkPasswordReq, ok := m.(*lnrpc.CheckPasswordRequest)
			if !ok {
				return nil, er.New("Argument is not a CheckPasswordRequest")
			}

			//	check the passphrase
			var checkPasswordResp *lnrpc.CheckPasswordResponse

			meta, errr := c.withMetaServer()
			if errr != nil {
				return nil, errr
			}

			checkPasswordResp, err := meta.CheckPassword(context.TODO(), checkPasswordReq)
			if err != nil {
				return nil, er.E(err)
			}

			return checkPasswordResp, nil
		},
	},
	//	WalletUnlocker: Wallet create  -  URI /wallet/create
	//	Will try to create/restore wallet
	{
		command: help.CommandCreateWallet,
		req:     (*lnrpc.InitWalletRequest)(nil), // Use init wallet structure to create
		res:     (*RestEmptyResponse)(nil),

		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			initWalletReq, ok := m.(*lnrpc.InitWalletRequest)
			if !ok {
				return nil, er.New("Argument is not a InitWalletRequest")
			}

			//	init wallet
			cc, errr := c.withUnlocker()
			if errr != nil {
				return nil, errr
			}

			_, err := cc.InitWallet(context.TODO(), initWalletReq)
			if err != nil {
				return nil, er.E(err)
			}
			return &RestEmptyResponse{}, nil
		},
	},
	//	service getsecret  -  URI /wallet/getsecret
	{
		command: help.CommandGetSecret,
		req:     (*lnrpc.GetSecretRequest)(nil),
		res:     (*lnrpc.GetSecretResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			getSecretReq, ok := m.(*lnrpc.GetSecretRequest)
			if !ok {
				return nil, er.New("Argument is not a GetSecretRequest")
			}

			//	invoke wallet get secret command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var getSecretResp *lnrpc.GetSecretResponse

				getSecretResp, err := cc.GetSecret(context.TODO(), getSecretReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return getSecretResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	GetWalletSeed  -  URI /wallet/seed
	{
		command: help.CommandGetWalletSeed,
		req:     nil,
		res:     (*lnrpc.GetWalletSeedResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.GetWalletSeed(context.TODO(), nil); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	WalletUnlocker: Wallet unlock  -  URI /wallet/unlock
	//	Will try to unlock the wallet with the password(s) provided
	{
		command: help.CommandUnlockWallet,
		req:     (*lnrpc.UnlockWalletRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.UnlockWalletRequest)
			if !ok {
				return nil, er.New("Argument is not a UnlockWalletRequest")
			}
			if u, err := c.withUnlocker(); err != nil {
				return nil, err
			} else if _, err := u.UnlockWallet0(context.TODO(), req); err != nil {
				return nil, err
			}
			return &RestEmptyResponse{}, nil
		},
	},

	//	>>> wallet/networkstewardvote subCategory command

	//	service getnetworkstewardvote  -  URI /wallet/networkstewardvote
	{
		command: help.CommandGetNetworkStewardVote,
		req:     nil,
		res:     (*lnrpc.GetNetworkStewardVoteResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get network steward vote
			cc, errr := c.withRpcServer()
			if cc != nil {
				var getNetworkStewardVoteResp *lnrpc.GetNetworkStewardVoteResponse

				getNetworkStewardVoteResp, err := cc.GetNetworkStewardVote(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return getNetworkStewardVoteResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service setnetworkstewardvote  -  URI /wallet/networkstewardvote/set
	{
		command: help.CommandSetNetworkStewardVote,
		req:     (*lnrpc.SetNetworkStewardVoteRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			setNetworkStewardVoteReq, ok := m.(*lnrpc.SetNetworkStewardVoteRequest)
			if !ok {
				return nil, er.New("Argument is not a SetNetworkStewardVoteRequest")
			}

			//	set network steward vote
			cc, errr := c.withRpcServer()
			if cc != nil {
				_, err := cc.SetNetworkStewardVote(context.TODO(), setNetworkStewardVoteReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> wallet/transaction subCategory command

	//	GetTransaction  -  URI /wallet/transaction
	{
		command: help.CommandGetTransaction,
		req:     (*lnrpc.GetTransactionRequest)(nil),
		res:     (*lnrpc.GetTransactionResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.GetTransactionRequest)
			if !ok {
				return nil, er.New("Argument is not a GetTransactionRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.GetTransaction(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	service createtransaction  -  URI /wallet/transaction/create
	{
		command: help.CommandCreateTransaction,
		req:     (*lnrpc.CreateTransactionRequest)(nil),
		res:     (*lnrpc.CreateTransactionResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			createTransactionReq, ok := m.(*lnrpc.CreateTransactionRequest)
			if !ok {
				return nil, er.New("Argument is not a CreateTransactionRequest")
			}

			//	invoke wallet create transaction command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var createTransactionResp *lnrpc.CreateTransactionResponse

				createTransactionResp, err := cc.CreateTransaction(context.TODO(), createTransactionReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return createTransactionResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	Wallet transactions  -  URI /wallet/transaction/query
	//	requires unlocked wallet -> access to rpcServer
	{
		command: help.CommandQueryTransactions,
		req:     (*lnrpc.GetTransactionsRequest)(nil),
		res:     (*lnrpc.TransactionDetails)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.GetTransactionsRequest)
			if !ok {
				return nil, er.New("Argument is not a GetTransactionsRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.GetTransactions(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	service sendcoins  -  URI /wallet/transaction/sendcoins
	{
		command: help.CommandSendCoins,
		req:     (*lnrpc.SendCoinsRequest)(nil),
		res:     (*lnrpc.SendCoinsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			sendCoinsReq, ok := m.(*lnrpc.SendCoinsRequest)
			if !ok {
				return nil, er.New("Argument is not a SendCoinsRequest")
			}

			//	send coins to one addresses
			cc, errr := c.withRpcServer()
			if cc != nil {
				var sendCoinsResp *lnrpc.SendCoinsResponse

				sendCoinsResp, err := cc.SendCoins(context.TODO(), sendCoinsReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return sendCoinsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	Sendfrom  -  URI /wallet/transaction/sendfrom
	{
		command: help.CommandSendFrom,
		req:     (*lnrpc.SendFromRequest)(nil),
		res:     (*lnrpc.SendFromResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.SendFromRequest)
			if !ok {
				return nil, er.New("Argument is not a SendFromRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.SendFrom(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	service sendmany  -  URI /wallet/transaction/sendmany
	{
		command: help.CommandSendMany,
		req:     (*lnrpc.SendManyRequest)(nil),
		res:     (*lnrpc.SendManyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			sendManyReq, ok := m.(*lnrpc.SendManyRequest)
			if !ok {
				return nil, er.New("Argument is not a SendManyRequest")
			}

			//	send coins to many addresses
			cc, errr := c.withRpcServer()
			if cc != nil {
				var sendManyResp *lnrpc.SendManyResponse

				sendManyResp, err := cc.SendMany(context.TODO(), sendManyReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return sendManyResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> wallet/unspent subCategory command

	//	service listunspent  -  URI /wallet/unspent
	{
		command: help.CommandListUnspent,
		req:     (*lnrpc.ListUnspentRequest)(nil),
		res:     (*lnrpc.ListUnspentResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			listUnspentReq, ok := m.(*lnrpc.ListUnspentRequest)
			if !ok {
				return nil, er.New("Argument is not a ListUnspentRequest")
			}

			//	get a list of available utxos
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listUnspentResp *lnrpc.ListUnspentResponse

				listUnspentResp, err := cc.ListUnspent(context.TODO(), listUnspentReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listUnspentResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	Resync  -  URI /wallet/unspent/resync
	{
		command: help.CommandResync,
		req:     (*lnrpc.ReSyncChainRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.ReSyncChainRequest)
			if !ok {
				return nil, er.New("Argument is not a ReSyncChainRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if _, err := server.ReSync(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	StopResync  -  URI /wallet/unspent/stopresync
	{
		command: help.CommandStopResync,
		req:     nil,
		res:     (*lnrpc.StopReSyncResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.StopReSync(context.TODO(), nil); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},

	//	>>> wallet/unspent/lock subCategory command

	//	service listlockunspent  -  URI /wallet/unspent/lock
	{
		command: help.CommandListLockUnspent,
		req:     nil,
		res:     (*lnrpc.ListLockUnspentResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	invoke wallet list lock unspent command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var listLockUnspentResp *lnrpc.ListLockUnspentResponse

				listLockUnspentResp, err := cc.ListLockUnspent(context.TODO(), nil)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listLockUnspentResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service lockunspent  -  URI /wallet/unspent/lock/create
	{
		command: help.CommandLockUnspent,
		req:     (*lnrpc.LockUnspentRequest)(nil),
		res:     (*lnrpc.LockUnspentResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			lockUnspentReq, ok := m.(*lnrpc.LockUnspentRequest)
			if !ok {
				return nil, er.New("Argument is not a LockUnspentRequest")
			}

			//	invoke wallet lock unspent command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var lockUnspentResp *lnrpc.LockUnspentResponse

				lockUnspentResp, err := cc.LockUnspent(context.TODO(), lockUnspentReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return lockUnspentResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> wallet/address subCategory command

	//	GetAddressBalances  -  URI /wallet/address/balances
	{
		command: help.CommandGetAddressBalances,
		req:     (*lnrpc.GetAddressBalancesRequest)(nil),
		res:     (*lnrpc.GetAddressBalancesResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.GetAddressBalancesRequest)
			if !ok {
				return nil, er.New("Argument is not a GetAddressBalancesRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.GetAddressBalances(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	New wallet address  -  URI /wallet/address/create
	//	requires unlocked wallet -> access to rpcServer
	{
		command: help.CommandNewAddress,
		req:     (*lnrpc.GetNewAddressRequest)(nil),
		res:     (*lnrpc.GetNewAddressResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {
			req, ok := m.(*lnrpc.GetNewAddressRequest)
			if !ok {
				return nil, er.New("Argument is not a GetNewAddressRequest")
			}
			if server, err := c.withRpcServer(); server != nil {
				if l, err := server.GetNewAddress(context.TODO(), req); err != nil {
					return nil, er.E(err)
				} else {
					return l, nil
				}
			} else {
				return nil, err
			}
		},
	},
	//	service dumpprivkey  -  URI /wallet/address/dumpprivkey
	{
		command: help.CommandDumpPrivkey,
		req:     (*lnrpc.DumpPrivKeyRequest)(nil),
		res:     (*lnrpc.DumpPrivKeyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			dumpPrivKeyReq, ok := m.(*lnrpc.DumpPrivKeyRequest)
			if !ok {
				return nil, er.New("Argument is not a DumpPrivKeyRequest")
			}

			//	invoke wallet dump private key command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var dumpPrivKeyResp *lnrpc.DumpPrivKeyResponse

				dumpPrivKeyResp, err := cc.DumpPrivKey(context.TODO(), dumpPrivKeyReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return dumpPrivKeyResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service importprivkey  -  URI /wallet/address/import
	{
		command: help.CommandImportPrivkey,
		req:     (*lnrpc.ImportPrivKeyRequest)(nil),
		res:     (*lnrpc.ImportPrivKeyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			importPrivKeyReq, ok := m.(*lnrpc.ImportPrivKeyRequest)
			if !ok {
				return nil, er.New("Argument is not a ImportPrivKeyRequest")
			}

			//	invoke wallet import private key command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var importPrivKeyResp *lnrpc.ImportPrivKeyResponse

				importPrivKeyResp, err := cc.ImportPrivKey(context.TODO(), importPrivKeyReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return importPrivKeyResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service signmessage  -  URI /wallet/address/signmessage
	{
		command: help.CommandSignMessage,
		req:     (*lnrpc.SignMessageRequest)(nil),
		res:     (*lnrpc.SignMessageResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			signMessageReq, ok := m.(*lnrpc.SignMessageRequest)
			if !ok {
				return nil, er.New("Argument is not a SignMessageRequest")
			}

			//	invoke wallet sign message command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var signMessageResp *lnrpc.SignMessageResponse

				signMessageResp, err := cc.SignMessage(context.TODO(), signMessageReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return signMessageResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	Decode transaction service  -  URI /wallet/transaction/decode
	{
		command: help.CommandDecodeRawTransaction,
		req:     (*lnrpc.DecodeRawTransactionRequest)(nil),
		res:     (*lnrpc.DecodeRawTransactionResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			decodeReq, ok := m.(*lnrpc.DecodeRawTransactionRequest)
			if !ok {
				return nil, er.New("Argument is not a DecodeRawTransactionRequest")
			}

			//	generate a new seed
			cc, errr := c.withRpcServer()
			if cc != nil {
				var decodeResp *lnrpc.DecodeRawTransactionResponse

				decodeResp, err := cc.DecodeRawTransaction(context.TODO(), decodeReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return decodeResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	>>> neutrino category command

	//	service bcasttransaction  -  URI /neutrino/bcasttransaction
	{
		command: help.CommandBcastTransaction,
		req:     (*lnrpc.BcastTransactionRequest)(nil),
		res:     (*lnrpc.BcastTransactionResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			bcastTransactionReq, ok := m.(*lnrpc.BcastTransactionRequest)
			if !ok {
				return nil, er.New("Argument is not a BcastTransactionRequest")
			}

			//	invoke Lightning broadcast transaction in chain command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var bcastTransactionResp *lnrpc.BcastTransactionResponse

				bcastTransactionResp, err := cc.BcastTransaction(context.TODO(), bcastTransactionReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return bcastTransactionResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service estimatefee  -  URI /neutrino/estimatefee
	{
		command: help.CommandEstimateFee,
		req:     (*lnrpc.EstimateFeeRequest)(nil),
		res:     (*lnrpc.EstimateFeeResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			estimateFeeReq, ok := m.(*lnrpc.EstimateFeeRequest)
			if !ok {
				return nil, er.New("Argument is not a EstimateFeeRequest")
			}

			//	get estimate fee info
			cc, errr := c.withRpcServer()
			if cc != nil {
				var estimateFeeResp *lnrpc.EstimateFeeResponse

				estimateFeeResp, err := cc.EstimateFee(context.TODO(), estimateFeeReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return estimateFeeResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> util/seed subCategory command

	//	Change Passphrase service  -  URI /util/seed/changepassphrase
	{
		command: help.CommandChangeSeedPassphrase,
		req:     (*lnrpc.ChangeSeedPassphraseRequest)(nil),
		res:     (*lnrpc.ChangeSeedPassphraseResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			changeSeedPassphraseReq, ok := m.(*lnrpc.ChangeSeedPassphraseRequest)
			if !ok {
				return nil, er.New("Argument is not a ChangeSeedPassphraseRequest")
			}

			//	invoke Lightning change seed passphrase command
			cc, errr := c.withRpcServer()
			if cc != nil {
				var changeSeedPassphraseResp *lnrpc.ChangeSeedPassphraseResponse

				changeSeedPassphraseResp, err := cc.ChangeSeedPassphrase(context.TODO(), changeSeedPassphraseReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return changeSeedPassphraseResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	GenSeed service  -  URI /util/seed/create
	{
		command: help.CommandGenSeed,
		req:     (*lnrpc.GenSeedRequest)(nil),
		res:     (*lnrpc.GenSeedResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			genSeedReq, ok := m.(*lnrpc.GenSeedRequest)
			if !ok {
				return nil, er.New("Argument is not a GenSeedRequest")
			}

			//	generate a new seed
			cc, errr := c.withUnlocker()
			if cc != nil {
				var genSeedResp *lnrpc.GenSeedResponse

				genSeedResp, err := cc.GenSeed(context.TODO(), genSeedReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return genSeedResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},

	//	>>> wtclient/tower subCategory command

	//	service CreateWatchTower  -  URI /wtclient/tower/create
	{
		command: help.CommandCreateWatchTower,
		req:     (*wtclientrpc.AddTowerRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			addTowerReq, ok := m.(*wtclientrpc.AddTowerRequest)
			if !ok {
				return nil, er.New("Argument is not a AddTowerRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				_, err := cc.AddTower(context.TODO(), addTowerReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service RemoveWatchTower  -  URI /wtclient/tower/remove
	{
		command: help.CommandRemoveTower,
		req:     (*wtclientrpc.RemoveTowerRequest)(nil),
		res:     (*RestEmptyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			removeTowerReq, ok := m.(*wtclientrpc.RemoveTowerRequest)
			if !ok {
				return nil, er.New("Argument is not a RemoveTowerRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				_, err := cc.RemoveTower(context.TODO(), removeTowerReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return &RestEmptyResponse{}, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service ListTowers  -  URI /wtclient/tower
	{
		command: help.CommandListTowers,
		req:     (*wtclientrpc.ListTowersRequest)(nil),
		res:     (*wtclientrpc.ListTowersResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			listTowersReq, ok := m.(*wtclientrpc.ListTowersRequest)
			if !ok {
				return nil, er.New("Argument is not a ListTowersRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				var listTowersResp *wtclientrpc.ListTowersResponse

				listTowersResp, err := cc.ListTowers(context.TODO(), listTowersReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return listTowersResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service GetTowerInfo  -  URI /wtclient/tower/getinfo
	{
		command: help.CommandGetTowerInfo,
		req:     (*wtclientrpc.GetTowerInfoRequest)(nil),
		res:     (*wtclientrpc.Tower)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			getTowerInfoReq, ok := m.(*wtclientrpc.GetTowerInfoRequest)
			if !ok {
				return nil, er.New("Argument is not a GetTowerInfoRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				var towerResp *wtclientrpc.Tower

				towerResp, err := cc.GetTowerInfo(context.TODO(), getTowerInfoReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return towerResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service GetTowerStats  -  URI /wtclient/tower/stats
	{
		command: help.CommandGetTowerStats,
		req:     nil,
		res:     (*wtclientrpc.StatsResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			statsReq, ok := m.(*wtclientrpc.StatsRequest)
			if !ok {
				return nil, er.New("Argument is not a StatsRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				var statsResp *wtclientrpc.StatsResponse

				statsResp, err := cc.Stats(context.TODO(), statsReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return statsResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
	//	service GetTowerPolicy  -  URI /wtclient/tower/policy
	{
		command: help.CommandGetTowerPolicy,
		req:     (*wtclientrpc.PolicyRequest)(nil),
		res:     (*wtclientrpc.PolicyResponse)(nil),
		f: func(c *RpcContext, m proto.Message) (proto.Message, er.R) {

			//	get the request payload
			policyReq, ok := m.(*wtclientrpc.PolicyRequest)
			if !ok {
				return nil, er.New("Argument is not a PolicyRequest")
			}

			//	invoke wallet get transactions command
			cc, errr := c.withWatchTowerClient()

			if cc != nil {
				var policyResp *wtclientrpc.PolicyResponse

				policyResp, err := cc.Policy(context.TODO(), policyReq)
				if err != nil {
					return nil, er.E(err)
				} else {
					return policyResp, nil
				}
			} else {
				return nil, errr
			}
		},
	},
}

type RpcContext struct {
	MaybeCC               *chainreg.ChainControl
	MaybeNeutrino         *neutrino.ChainService
	MaybeWallet           *wallet.Wallet
	MaybeRpcServer        lnrpc.LightningServer
	MaybeWalletUnlocker   *walletunlocker.UnlockerService
	MaybeMetaService      lnrpc.MetaServiceServer
	MaybeVerRPCServer     verrpc.VersionerServer
	MaybeRouterServer     routerrpc.RouterServer
	MaybeWatchTowerClient wtclientrpc.WatchtowerClientClient
}

func with(thing interface{}, name string) er.R {
	if thing == nil {
		return er.Errorf("Could not call function because [%s] is not yet ready", name)
	}
	return nil
}
func (c *RpcContext) withCC() (*chainreg.ChainControl, er.R) {
	return c.MaybeCC, with(c.MaybeCC, "ChainController")
}
func (c *RpcContext) withNeutrino() (*neutrino.ChainService, er.R) {
	return c.MaybeNeutrino, with(c.MaybeNeutrino, "Neutrino")
}
func (c *RpcContext) withWallet() (*wallet.Wallet, er.R) {
	return c.MaybeWallet, with(c.MaybeWallet, "Wallet")
}
func (c *RpcContext) withRpcServer() (lnrpc.LightningServer, er.R) {
	return c.MaybeRpcServer, with(c.MaybeRpcServer, "LightningServer")
}
func (c *RpcContext) withUnlocker() (*walletunlocker.UnlockerService, er.R) {
	return c.MaybeWalletUnlocker, with(c.MaybeWalletUnlocker, "UnlockerService")
}
func (c *RpcContext) withMetaServer() (lnrpc.MetaServiceServer, er.R) {
	return c.MaybeMetaService, with(c.MaybeMetaService, "MetaServiceService")
}
func (c *RpcContext) withVerRPCServer() (verrpc.VersionerServer, er.R) {
	return c.MaybeVerRPCServer, with(c.MaybeVerRPCServer, "VersionerService")
}
func (c *RpcContext) withRouterServer() (routerrpc.RouterServer, er.R) {
	return c.MaybeRouterServer, with(c.MaybeRouterServer, "RouterServer")
}

func (c *RpcContext) withWatchTowerClient() (wtclientrpc.WatchtowerClientClient, er.R) {
	return c.MaybeWatchTowerClient, with(c.MaybeWatchTowerClient, "WatchTowerClient")
}

type SimpleHandler struct {
	rf RpcFunc
	c  *RpcContext
}

func unmarshal1(r *http.Request, m proto.Message, isJson bool) er.R {
	if b, err := io.ReadAll(r.Body); err != nil {
		return er.E(err)
	} else if isJson {
		// Use jsoniter for unmarshaling because it is far more forgiving
		if err := jsoniter.Unmarshal(b, m); err != nil {
			return er.E(err)
		}
	} else if err := proto.Unmarshal(b, m); err != nil {
		return er.E(err)
	}
	return nil
}

func unmarshal(r *http.Request, m proto.Message, isJson bool) er.R {
	if isJson {
		if err := jsonpb.Unmarshal(r.Body, m); err != nil {
			return er.E(err)
		}
	} else {
		if b, err := io.ReadAll(r.Body); err != nil {
			return er.E(err)
		} else if err := proto.Unmarshal(b, m); err != nil {
			return er.E(err)
		}
	}
	return nil
}
func marshal(w http.ResponseWriter, m proto.Message, isJson bool) er.R {
	if m == nil {
		return nil
	}
	if isJson {
		marshaler := jsonpb.Marshaler{
			OrigName:     false,
			EnumsAsInts:  false,
			EmitDefaults: true,
			Indent:       "\t",
		}
		if s, err := marshaler.MarshalToString(m); err != nil {
			return er.E(err)
		} else if _, err := io.WriteString(w, s); err != nil {
			return er.E(err)
		}
	} else {
		if b, err := proto.Marshal(m); err != nil {
			return er.E(err)
		} else if _, err := w.Write(b); err != nil {
			return er.E(err)
		}
	}
	return nil
}

func (s *SimpleHandler) ServeHttpOrErr(w http.ResponseWriter, r *http.Request, isJson bool) er.R {
	var commandInfo help.CommandInfo

	for _, commandInfo = range help.CommandInfoData {
		if commandInfo.Command == s.rf.command {
			break
		}
	}

	//	check if the URI is for command help
	if r.RequestURI == help.URI_prefix+helpURI_prefix+commandInfo.Path {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return er.New("405 - Request should be a GET because the help endpoint requires no input")
		}
		if commandInfo.HelpInfo == nil {
			w.WriteHeader(http.StatusNotFound)
			return er.New("404 - help not available for command: " + commandInfo.Path)
		}

		err := marshalHelp(w, commandInfo.HelpInfo())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		return nil
	}

	//	command URI handler
	var req proto.Message

	if s.rf.req != nil {
		if r.Method != "POST" {
			//	if it's not a POST, check if the endpoint allows GET
			if !commandInfo.AllowGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return er.New("405 - Request should be a POST because the endpoint requires input")
			}

			//	not a POST and not a GET so, it's defenitely not allowed
			if r.Method != "GET" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return er.New("405 - Method not allowed: " + r.Method)
			}
		}
		req1 := reflect.New(reflect.TypeOf(s.rf.req).Elem())
		if r, ok := req1.Interface().(proto.Message); !ok {
			panic("elem is not a proto.Message")
		} else {
			req = r
		}
		//	check the method again, because if it's a GET, there's no request payload
		if r.Method == "POST" {
			if err := unmarshal(r, req, isJson); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return err
			}
		}
	}
	if res, err := s.rf.f(s.c, req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	} else if err := marshal(w, res, isJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func (s *SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct := strings.ToLower(r.Header.Get("Content-Type"))
	isJson := strings.Contains(ct, "application/json")
	if !isJson && !strings.Contains(ct, "application/protobuf") {
		if r.Method == "GET" {
			isJson = true
		} else {
			w.Header().Set("Connection", "close")
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "415 - Invalid content type, must be json or protobuf", http.StatusUnsupportedMediaType)
			return
		}
	}
	if err := s.ServeHttpOrErr(w, r, isJson); err != nil {
		if err = marshal(w, &lnrpc.RestError{
			Message: err.Message(),
			Stack:   err.Stack(),
		}, isJson); err != nil {
			log.Errorf("Error replying to request for [%s] from [%s] - error sending error, giving up: [%s]",
				r.RequestURI, r.RemoteAddr, err)
		}
	}
}

func RestHandlers(c *RpcContext) *mux.Router {
	r := mux.NewRouter()
	for _, rf := range rpcFunctions {
		var commandInfo help.CommandInfo
		var commandFound bool

		for _, commandInfo = range help.CommandInfoData {
			if commandInfo.Command == rf.command {
				commandFound = true
				break
			}
		}
		//	if all commands in s.rf (rpcFunctions) slice have an entry in help.CommandInfoData
		//		this panic condition will never be reached
		if !commandFound {
			panic("[panic] Command info about command " + rf.command + "not found")
		}

		r.Handle(help.URI_prefix+commandInfo.Path, &SimpleHandler{c: c, rf: rf})
		r.Handle(help.URI_prefix+helpURI_prefix+commandInfo.Path, &SimpleHandler{c: c, rf: rf})
	}

	//	add a handler for endpoint not found (404)
	r.NotFoundHandler = http.HandlerFunc(func(httpResponse http.ResponseWriter, r *http.Request) {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "404 - invalid endpoint: for help on all endpoints go to /api/v1 URI", http.StatusNotFound)
	})

	//	add a handler for websocket endpoint
	r.Handle(help.URI_prefix+"/meta/websocket", http.HandlerFunc(func(httpResponse http.ResponseWriter, httpRequest *http.Request) {
		webSocketHandler(c, httpResponse, httpRequest)
	}))

	return r
}
