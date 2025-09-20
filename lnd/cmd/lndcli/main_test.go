////////////////////////////////////////////////////////////////////////////////
//	lndcli/lnd_client_test.go  -  May-4-2022  -  aldebap
//
//	unit tests for pld client
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"testing"
)

type testScenario struct {
	name            string
	command         string
	arguments       []string
	expectedPayload string
	expectedError   string
}

func TestFormatRequestPayload(t *testing.T) {

	pldServer = "http://localhost:8080"

	var testCases []testScenario = []testScenario{
		//	test error handling
		{
			name:          "invalid command path error handling",
			command:       "meta/debug_level",
			arguments:     []string{},
			expectedError: `invalid pld command: meta/debug_level`,
		},
		{
			name:          "invalid command CLI argument error handling",
			command:       "meta/debuglevel",
			arguments:     []string{"--show", "--level_spec=debug", "--amount=100000"},
			expectedError: `invalid command argument: --amount=100000`,
		},
		{
			name:    "missing square bracket before array elements",
			command: "util/seed/changepassphrase",
			arguments: []string{"--current_seed_passphrase_bin=cGFzc3dvcmQ=", `--current_seed="plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", ` +
				`"example", "around"]`, "--new_seed_passphrase=password"},
			expectedError: `error parsing arguments: array argument must be delimitted by square brackets: "plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", "example", ` +
				`"around"]`,
		},
		{
			name:    "missing square bracket after array elements",
			command: "util/seed/changepassphrase",
			arguments: []string{"--current_seed_passphrase_bin=cGFzc3dvcmQ=", `--current_seed=["plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", ` +
				`"example", "around"`, "--new_seed_passphrase=password"},
			expectedError: `error parsing arguments: array argument must be delimitted by square brackets: ["plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", "example", ` +
				`"around"`,
		},
		{
			name:    "missing open double quotes before array element",
			command: "util/seed/changepassphrase",
			arguments: []string{"--current_seed_passphrase_bin=cGFzc3dvcmQ=", `--current_seed=[plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", ` +
				`"example", "around"]`, "--new_seed_passphrase=password"},
			expectedError: `error parsing arguments: array element must be delimitted by double quotes: plastic"`,
		},
		{
			name:    "missing close double quotes after array element",
			command: "util/seed/changepassphrase",
			arguments: []string{"--current_seed_passphrase_bin=cGFzc3dvcmQ=", `--current_seed=["plastic, "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", ` +
				`"example", "around"]`, "--new_seed_passphrase=password"},
			expectedError: `error parsing arguments: array element must be delimitted by double quotes: "plastic`,
		},
		//	test commands of "meta" group
		{
			name:            "debuglevel CLI options",
			command:         "meta/debuglevel",
			arguments:       []string{"--show", "--level_spec=debug"},
			expectedPayload: `{ "show": true, "level_spec": "debug" }`,
		},
		{
			name:            "getinfo CLI options",
			command:         "meta/getinfo",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "version CLI options",
			command:         "meta/version",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		//	test commands of "util/seed" group
		{
			name:    "changepassphrase CLI options",
			command: "util/seed/changepassphrase",
			arguments: []string{"--current_seed_passphrase_bin=cGFzc3dvcmQ=", `--current_seed=["plastic", "hollow", ` +
				`"mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", ` +
				`"example", "around"]`, "--new_seed_passphrase=password"},
			expectedPayload: `{ "current_seed_passphrase_bin": "cGFzc3dvcmQ=", ` +
				`"current_seed": [ "plastic", "hollow", "mansion", "keep", "into", "cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", "example", "around" ], ` +
				`"new_seed_passphrase": "password" }`,
		},
		{
			name:            "genseed CLI options",
			command:         "util/seed/create",
			arguments:       []string{"--seed_passphrase_bin=cGFzc3dvcmQ="},
			expectedPayload: `{ "seed_passphrase_bin": "cGFzc3dvcmQ=" }`,
		},
		//	test commands of "Lightning/Channel" group
		{
			name:    "openchannel CLI options",
			command: "lightning/channel/open",
			arguments: []string{"--node_pubkey=a0a1a2a3a4a5a6a7a8a9aa=", "--local_funding_amount=100000", "--push_sat=500000", "--target_conf=50",
				"--sat_per_byte=250", "--private", "--min_htlc_msat=2048", "--remote_csv_delay=256", "--min_confs=5", "--close_address=C70534dd",
				"--funding_shim.chan_point_shim.amt=2050750", "--funding_shim.chan_point_shim.chan_point.funding_txid_bytes=f0f1f2f3f4f5f6f7f8f9ff",
				"--funding_shim.chan_point_shim.chan_point.output_index=282748", "--funding_shim.chan_point_shim.local_key.raw_key_bytes=1011121314151617181910",
				"--funding_shim.chan_point_shim.local_key.key_loc.key_family=25", "--funding_shim.chan_point_shim.local_key.key_loc.key_index=70",
				"--funding_shim.chan_point_shim.remote_key=e0e2e3e4e5e6e7e8e9ee", "--funding_shim.chan_point_shim.pending_chan_id=632d129",
				"--funding_shim.chan_point_shim.thaw_height=16274648", "--funding_shim.psbt_shim.pending_chan_id=374649828379",
				"--funding_shim.psbt_shim.base_psbt=3754629387292", "--funding_shim.psbt_shim.no_publish",
				"--remote_max_value_in_flight_msat=100500", "--remote_max_htlcs=287", "--max_local_csv=209",
			},
			expectedPayload: `{ "node_pubkey": "a0a1a2a3a4a5a6a7a8a9aa=", "local_funding_amount": 100000, "push_sat": 500000, "target_conf": 50, ` +
				`"sat_per_byte": 250, "private": true, "min_htlc_msat": 2048, "remote_csv_delay": 256, "min_confs": 5, "close_address": "C70534dd", ` +
				`"funding_shim": { "chan_point_shim": { "amt": 2050750, "chan_point": { "funding_txid_bytes": "f0f1f2f3f4f5f6f7f8f9ff", "output_index": 282748 }, ` +
				`"local_key": { "raw_key_bytes": "1011121314151617181910", "key_loc": { "key_family": 25, "key_index": 70 } }, "remote_key": "e0e2e3e4e5e6e7e8e9ee", ` +
				`"pending_chan_id": "632d129", "thaw_height": 16274648 }, "psbt_shim": { "pending_chan_id": "374649828379", "base_psbt": "3754629387292", ` +
				`"no_publish": true } }, "remote_max_value_in_flight_msat": 100500, "remote_max_htlcs": 287, "max_local_csv": 209 }`,
		},
		{
			name:            "closechannel CLI options",
			command:         "lightning/channel/close",
			arguments:       []string{"--channel_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa", "--force", "--delivery_address=d0d1d2d3d4d5d6d7d8d9dd"},
			expectedPayload: `{ "channel_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa" }, "force": true, "delivery_address": "d0d1d2d3d4d5d6d7d8d9dd" }`,
		},
		{
			name:            "abandonchannel CLI options",
			command:         "lightning/channel/abandon",
			arguments:       []string{"--channel_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa", "--channel_point.output_index=282748", "--pending_funding_shim_only"},
			expectedPayload: `{ "channel_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa", "output_index": 282748 }, "pending_funding_shim_only": true }`,
		},
		{
			name:            "channelbalance CLI options",
			command:         "lightning/channel/balance",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "pendingchannels CLI options",
			command:         "lightning/channel/pending",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "listchannels CLI options",
			command:         "lightning/channel",
			arguments:       []string{"--active_only", "--public_only", "--peer=a0a1a2a3a4a5a6a7a8a9aa="},
			expectedPayload: `{ "active_only": true, "public_only": true, "peer": "a0a1a2a3a4a5a6a7a8a9aa=" }`,
		},
		{
			name:            "closedchannels CLI options",
			command:         "lightning/channel/closed",
			arguments:       []string{"--local_force", "--abandoned"},
			expectedPayload: `{ "local_force": true, "abandoned": true }`,
		},
		{
			name:            "getnetworkinfo CLI options",
			command:         "lightning/channel/networkinfo",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "feereport CLI options",
			command:         "lightning/channel/feereport",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "updatechanpolicy CLI options",
			command:         "lightning/channel/policy",
			arguments:       []string{"--global", "--chan_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa", "--chan_point.output_index=282748", "--base_fee_msat=10", "--fee_rate=100"},
			expectedPayload: `{ "global": true, "chan_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa", "output_index": 282748 }, "base_fee_msat": 10, "fee_rate": 100 }`,
		},
		//	test commands of "Lightning/Channel/Backup" group
		{
			name:            "exportchanbackup CLI options",
			command:         "lightning/channel/backup/export",
			arguments:       []string{"--chan_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa", "--chan_point.output_index=282748"},
			expectedPayload: `{ "chan_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa", "output_index": 282748 } }`,
		},
		{
			name:    "restorechanbackup CLI options",
			command: "lightning/channel/backup/restore",
			arguments: []string{"--chan_backups.chan_backups.chan_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa",
				"--chan_backups.chan_backups.chan_point.output_index=282748", "--chan_backups.chan_backups.chan_backup=RW5jcnlwdGVkIENoYW4gQmFja3Vw"},
			expectedPayload: `{ "chan_backups": { "chan_backups": [ "chan_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa", "output_index": 282748 }, ` +
				`"chan_backup": "RW5jcnlwdGVkIENoYW4gQmFja3Vw" ] } }`,
		},
		{
			name:    "verifychanbackup CLI options",
			command: "lightning/channel/backup/verify",
			arguments: []string{"--single_chan_backups.chan_backups.chan_point.funding_txid_str=a0a1a2a3a4a5a6a7a8a9aa",
				"--single_chan_backups.chan_backups.chan_point.output_index=282748", "--single_chan_backups.chan_backups.chan_backup=RW5jcnlwdGVkIENoYW4gQmFja3Vw",
				"--multi_chan_backup.chan_points.funding_txid_str=b0b1b2b3b4b5b6b7b8b9bb"},
			expectedPayload: `{ "single_chan_backups": { "chan_backups": [ "chan_point": { "funding_txid_str": "a0a1a2a3a4a5a6a7a8a9aa", "output_index": 282748 }, ` +
				`"chan_backup": "RW5jcnlwdGVkIENoYW4gQmFja3Vw" ] }, "multi_chan_backup": { "chan_points": [ "funding_txid_str": "b0b1b2b3b4b5b6b7b8b9bb" ] } }`,
		},
		//	test commands of "Lightning/Graph" group
		{
			name:            "describegraph CLI options",
			command:         "lightning/graph",
			arguments:       []string{"--include_unannounced"},
			expectedPayload: `{ "include_unannounced": true }`,
		},
		{
			name:            "getnodemetrics CLI options",
			command:         "lightning/graph/nodemetrics",
			arguments:       []string{"--types.UNKNOWN", "--types.BETWEENNESS_CENTRALITY"},
			expectedPayload: `{ "types": [ "UNKNOWN", "BETWEENNESS_CENTRALITY" ] }`,
		},
		{
			name:            "getchaninfo CLI options",
			command:         "lightning/graph/channel",
			arguments:       []string{"--chan_id=123456"},
			expectedPayload: `{ "chan_id": 123456 }`,
		},
		{
			name:            "getnodeinfo CLI options",
			command:         "lightning/graph/nodeinfo",
			arguments:       []string{"--pub_key=a0a1a2a3a4a5a6a7a8a9aa", "--include_channels"},
			expectedPayload: `{ "pub_key": "a0a1a2a3a4a5a6a7a8a9aa", "include_channels": true }`,
		},
		//	test commands of "Lightning/Invoice" group
		{
			name:            "addinvoice CLI options",
			command:         "lightning/invoice/create",
			arguments:       []string{"--memo=xpto", "--r_preimage=0123456789abcdef0123456789abcdef", "--r_hash=00112233445566778899", "--value=10", "--expiry=3600"},
			expectedPayload: `{ "memo": "xpto", "r_preimage": "0123456789abcdef0123456789abcdef", "r_hash": "00112233445566778899", "value": 10, "expiry": 3600 }`,
		},
		{
			name:            "lookupinvoice CLI options",
			command:         "lightning/invoice/lookup",
			arguments:       []string{"--r_hash=00112233445566778899"},
			expectedPayload: `{ "r_hash": "00112233445566778899" }`,
		},
		{
			name:            "listinvoices CLI options",
			command:         "lightning/invoice",
			arguments:       []string{"--index_offset=1", "--num_max_invoices=10", "--reversed"},
			expectedPayload: `{ "index_offset": 1, "num_max_invoices": 10, "reversed": true }`,
		},
		{
			name:            "decodepayreq CLI options",
			command:         "lightning/invoice/decodepayreq",
			arguments:       []string{"--pay_req=123456"},
			expectedPayload: `{ "pay_req": "123456" }`,
		},
		//	test commands of "Lightning/Payment" group
		{
			name:            "sendpayment CLI options",
			command:         "lightning/payment/send",
			arguments:       []string{"--dest=010203040506070809", "--payment_hash=02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e", "--amt=100000"},
			expectedPayload: `{ "dest": "010203040506070809", "amt": 100000, "payment_hash": "02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e" }`,
		},
		/*
			//	TODO: help for this command is missing !
			{
				name:            "payinvoice CLI options",
				command:         "/lightning/payment/payinvoice",
				arguments:       []string{""},
				expectedPayload: `{  }`,
			},
		*/
		{
			name:            "sendtoroute CLI options",
			command:         "lightning/payment/sendtoroute",
			arguments:       []string{"--payment_hash=02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e", "--route.hops.amt_to_forward=100000"},
			expectedPayload: `{ "payment_hash": "02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e", "route": { "hops": [ "amt_to_forward": 100000 ] } }`,
		},
		{
			name:            "listpayments CLI options",
			command:         "lightning/payment",
			arguments:       []string{"--include_incomplete", "--max_payments=50"},
			expectedPayload: `{ "include_incomplete": true, "max_payments": 50 }`,
		},
		/*
			{
				name:            "trackpayment CLI options",
				command:         "/lightning/payment/track",
				arguments:       []string{""},
				expectedPayload: `{  }`,
			},
		*/
		//	TODO: discuss with Caleb how to deal wth this type of fields (route_hints.hop_hints)
		{
			name:    "queryroutes CLI options",
			command: "lightning/payment/queryroutes",
			arguments: []string{"--pub_key=a0a1a2a3a4a5a6a7a8a9aa", "--amt=100000", "--route_hints.hop_hints.node_id=123456", "--dest_features.INITIAL_ROUING_SYNC",
				"--dest_features.STATIC_REMOTE_KEY_REQ"},
			expectedPayload: `{ "pub_key": "a0a1a2a3a4a5a6a7a8a9aa", "amt": 100000, "route_hints": [ "hop_hints": [ "node_id": "123456" ] ], ` +
				`"dest_features": [ "INITIAL_ROUING_SYNC", "STATIC_REMOTE_KEY_REQ" ] }`,
		},
		{
			name:            "fwdinghistory CLI options",
			command:         "lightning/payment/fwdinghistory",
			arguments:       []string{"--start_time=100", "--num_max_events=25"},
			expectedPayload: `{ "start_time": 100, "num_max_events": 25 }`,
		},
		{
			name:            "querymc CLI options",
			command:         "lightning/payment/querymc",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "queryprob CLI options",
			command:         "lightning/payment/queryprob",
			arguments:       []string{"--from_node=123456", "--to_node=789012"},
			expectedPayload: `{ "from_node": "123456", "to_node": "789012" }`,
		},
		{
			name:            "resetmc CLI options",
			command:         "lightning/payment/resetmc",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "buildroute CLI options",
			command:         "lightning/payment/buildroute",
			arguments:       []string{"--amt_msat=100000", "--outgoing_chan_id=0", "--hop_pubkeys=a0a1a2a3a4a5a6a7a8a9aa"},
			expectedPayload: `{ "amt_msat": 100000, "outgoing_chan_id": 0, "hop_pubkeys": "a0a1a2a3a4a5a6a7a8a9aa" }`,
		},
		//	test commands of "Lightning/Peer" group
		{
			name:            "connect CLI options",
			command:         "lightning/peer/connect",
			arguments:       []string{"--addr.pubkey=a0a1a2a3a4a5a6a7a8a9aa", "--perm"},
			expectedPayload: `{ "addr": { "pubkey": "a0a1a2a3a4a5a6a7a8a9aa" }, "perm": true }`,
		},
		{
			name:            "disconnect CLI options",
			command:         "lightning/peer/disconnect",
			arguments:       []string{"--pub_key=a0a1a2a3a4a5a6a7a8a9aa"},
			expectedPayload: `{ "pub_key": "a0a1a2a3a4a5a6a7a8a9aa" }`,
		},
		{
			name:            "listpeers CLI options",
			command:         "lightning/peer",
			arguments:       []string{},
			expectedPayload: ``,
		},
		//	test commands of "neutrino" group
		{
			name:            "bcasttransaction CLI options",
			command:         "neutrino/bcasttransaction",
			arguments:       []string{"--tx=f0f1f2f3f4f5f6f7f8f9ff"},
			expectedPayload: `{ "tx": "f0f1f2f3f4f5f6f7f8f9ff" }`,
		},
		{
			name:      "estimatefee CLI options",
			command:   "neutrino/estimatefee",
			arguments: []string{"--AddrToAmount.key=2625478276", "--AddrToAmount.value=8827484919"},
			//	TODO: how to parse the help to correctly build this payload ?
			expectedPayload: `{ "AddrToAmount": [ "key": "2625478276", "value": 8827484919 ] }`,
		},
		//	test commands of "wallet" group
		{
			name:            "walletbalance CLI options",
			command:         "wallet/balance",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "changePassphrase CLI options",
			command:         "wallet/changepassphrase",
			arguments:       []string{"--current_passphrase=p4sswd", "--new_passphrase=n3wP4sswd"},
			expectedPayload: `{ "current_passphrase": "p4sswd", "new_passphrase": "n3wP4sswd" }`,
		},
		{
			name:            "checkPassphrase CLI options",
			command:         "wallet/checkpassphrase",
			arguments:       []string{"--wallet_passphrase=p4sswd"},
			expectedPayload: `{ "wallet_passphrase": "p4sswd" }`,
		},
		{
			name:    "create_wallet CLI options",
			command: "wallet/create",
			arguments: []string{"--wallet_passphrase=p4sswd", `--wallet_seed=["plastic","hollow",` +
				`"mansion","keep","into","cloth","awesome","salmon","reopen","inner","replace","dice","life",` +
				`"example","around"]`, "--seed_passphrase_bin=s33dP4sswd"},
			expectedPayload: `{ "wallet_passphrase": "p4sswd", "wallet_seed": [ "plastic", "hollow", "mansion", "keep", "into", ` +
				`"cloth", "awesome", "salmon", "reopen", "inner", "replace", "dice", "life", "example", "around" ], ` +
				`"seed_passphrase_bin": "s33dP4sswd" }`,
		},
		{
			name:            "getsecret CLI options",
			command:         "wallet/getsecret",
			arguments:       []string{"--name=Isaac Assimov"},
			expectedPayload: `{ "name": "Isaac Assimov" }`,
		},
		{
			name:            "getwalletseed CLI options",
			command:         "wallet/seed",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "unlock_wallet CLI options",
			command:         "wallet/unlock",
			arguments:       []string{"--wallet_passphrase=p4sswd"},
			expectedPayload: `{ "wallet_passphrase": "p4sswd" }`,
		},
		//	test commands of "wallet/address" group
		{
			name:            "getaddressbalances CLI options",
			command:         "wallet/address/balances",
			arguments:       []string{"--minconf=10", "--showzerobalance"},
			expectedPayload: `{ "minconf": 10, "showzerobalance": true }`,
		},
		{
			name:            "newaddress CLI options",
			command:         "wallet/address/create",
			arguments:       []string{"--legacy"},
			expectedPayload: `{ "legacy": true }`,
		},
		{
			name:            "dumpprivkey CLI options",
			command:         "wallet/address/dumpprivkey",
			arguments:       []string{"--address=pkt1q85n69mzthdxlwutn6dr6f7kwyd9nv8ulasdaqz"},
			expectedPayload: `{ "address": "pkt1q85n69mzthdxlwutn6dr6f7kwyd9nv8ulasdaqz" }`,
		},
		{
			name:            "importprivkey CLI options",
			command:         "wallet/address/import",
			arguments:       []string{"--private_key=d0d1d2d3d4d5d6d7d8d9dd", "--rescan"},
			expectedPayload: `{ "private_key": "d0d1d2d3d4d5d6d7d8d9dd", "rescan": true }`,
		},
		{
			name:            "signmessage CLI options",
			command:         "wallet/address/signmessage",
			arguments:       []string{"--msg=pldctl unit tests", "--key_loc.key_family=1"},
			expectedPayload: `{ "msg": "pldctl unit tests", "key_loc": { "key_family": 1 } }`,
		},
		//	test commands of "wallet/networkstewardvote" group
		{
			name:            "getnetworkstewardvote CLI options",
			command:         "wallet/networkstewardvote",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "setnetworkstewardvote CLI options",
			command:         "wallet/networkstewardvote/set",
			arguments:       []string{"--vote_against=0", "--vote_for=1"},
			expectedPayload: `{ "vote_against": "0", "vote_for": "1" }`,
		},
		//	test commands of "wallet/transaction" group
		{
			name:            "gettransaction CLI options",
			command:         "wallet/transaction",
			arguments:       []string{"--txid=e0e1e2e3e4e5e6e7e8e9ee"},
			expectedPayload: `{ "txid": "e0e1e2e3e4e5e6e7e8e9ee" }`,
		},
		{
			name:            "createtransaction CLI options",
			command:         "wallet/transaction/create",
			arguments:       []string{"--to_address=2734648278367", "--amount=100000"},
			expectedPayload: `{ "to_address": "2734648278367", "amount": 100000 }`,
		},
		{
			name:            "query CLI options",
			command:         "wallet/transaction/query",
			arguments:       []string{"--start_height=1000000", "--end_height=1500000"},
			expectedPayload: `{ "start_height": 1000000, "end_height": 1500000 }`,
		},
		{
			name:            "sendcoins CLI options",
			command:         "wallet/transaction/sendcoins",
			arguments:       []string{"--addr=2734648278367", "--send_all"},
			expectedPayload: `{ "addr": "2734648278367", "send_all": true }`,
		},
		{
			name:            "sendfrom CLI options",
			command:         "wallet/transaction/sendfrom",
			arguments:       []string{"--to_address=2734648278367", "--amount=100000", `--from_address=[ "8274745638725", "475647384723" ]`},
			expectedPayload: `{ "to_address": "2734648278367", "amount": 100000, "from_address": [ "8274745638725", "475647384723" ] }`,
		},
		{
			//	TODO: again, how to parse the help to correctly build this payload ?
			name:            "sendmany CLI options",
			command:         "wallet/transaction/sendmany",
			arguments:       []string{"--AddrToAmount.key=2734648278367", "--AddrToAmount.value=100000", "--label=Payment for mobile phone bill"},
			expectedPayload: `{ "AddrToAmount": [ "key": "2734648278367", "value": 100000 ], "label": "Payment for mobile phone bill" }`,
		},
		//	test commands of "wallet/unspent" group
		{
			name:            "listunspent CLI options",
			command:         "wallet/unspent",
			arguments:       []string{"--min_confs=10", "--max_confs=50"},
			expectedPayload: `{ "min_confs": 10, "max_confs": 50 }`,
		},
		{
			name:            "resync CLI options",
			command:         "wallet/unspent/resync",
			arguments:       []string{"--from_height=1200000"},
			expectedPayload: `{ "from_height": 1200000 }`,
		},
		{
			name:            "stopresync CLI options",
			command:         "wallet/unspent/stopresync",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		//	test commands of "wallet/unspent/lock" group
		{
			name:            "listlockunspent CLI options",
			command:         "wallet/unspent/lock",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "lockunspent CLI options",
			command:         "wallet/unspent/lock/create",
			arguments:       []string{"--unlock", "--transactions.txid=e0e1e2e3e4e5e6e7e8e9ee"},
			expectedPayload: `{ "unlock": true, "transactions": [ "txid": "e0e1e2e3e4e5e6e7e8e9ee" ] }`,
		},
		//	test commands of "wtclient/tower" group
		{
			name:            "listtowers CLI options",
			command:         "wtclient/tower",
			arguments:       []string{},
			expectedPayload: ``,
		},
		{
			name:            "createwatchtower CLI options",
			command:         "wtclient/tower/create",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "gettowerinfo CLI options",
			command:         "wtclient/tower/getinfo",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "gettowerpolicy CLI options",
			command:         "wtclient/tower/policy",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "removewatchtower CLI options",
			command:         "wtclient/tower/remove",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
		{
			name:            "gettowerstats CLI options",
			command:         "wtclient/tower/stats",
			arguments:       []string{},
			expectedPayload: `{  }`,
		},
	}

	for _, testCase := range testCases {

		t.Logf(">>> Test the conversion of all CLI arguments into REST JSon request payload")

		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.expectedPayload
			got, err := formatRequestPayload(testCase.command, testCase.arguments)
			if err != nil {
				if len(testCase.expectedError) == 0 {
					t.Errorf("Unexpected error formatting the payload: %s", err)
					return
				}
				want = testCase.expectedError
				got = err.Error()
			}

			if want != got {
				t.Errorf("Error formatting the payload: got '%s', want '%s'", got, want)
			}
		})
	}
}
