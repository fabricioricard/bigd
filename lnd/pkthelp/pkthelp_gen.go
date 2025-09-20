package pkthelp
func mkdouble() Type {
    return Type{
        Name: "float64",
    }
}
func mkfloat() Type {
    return Type{
        Name: "float32",
    }
}
func mkint32() Type {
    return Type{
        Name: "int32",
    }
}
func mkint64() Type {
    return Type{
        Name: "int64",
    }
}
func mkuint32() Type {
    return Type{
        Name: "uint32",
    }
}
func mkuint64() Type {
    return Type{
        Name: "uint64",
    }
}
func mksint32() Type {
    return Type{
        Name: "int32",
    }
}
func mksint64() Type {
    return Type{
        Name: "int64",
    }
}
func mkfixed32() Type {
    return Type{
        Name: "uint32",
    }
}
func mkfixed64() Type {
    return Type{
        Name: "uint64",
    }
}
func mksfixed32() Type {
    return Type{
        Name: "int32",
    }
}
func mksfixed64() Type {
    return Type{
        Name: "int64",
    }
}
func mkbool() Type {
    return Type{
        Name: "bool",
    }
}
func mkstring() Type {
    return Type{
        Name: "string",
    }
}
func mkbytes() Type {
    return Type{
        Name: "[]byte",
    }
}
func mklnrpc_AddressType() Type {
    return Type{
        Name: "lnrpc_AddressType",
        Description: []string{
            "`AddressType` has to be one of:",
            "",
            "- `p2wkh`: Pay to witness key hash (`WITNESS_PUBKEY_HASH` = 0)",
            "- `np2wkh`: Pay to nested witness key hash (`NESTED_PUBKEY_HASH` = 1)",
        },
        Fields: []Field{
            {
                Name: "WITNESS_PUBKEY_HASH",
                Type: EnumVarientType,
            },
            {
                Name: "NESTED_PUBKEY_HASH",
                Type: EnumVarientType,
            },
            {
                Name: "UNUSED_WITNESS_PUBKEY_HASH",
                Type: EnumVarientType,
            },
            {
                Name: "UNUSED_NESTED_PUBKEY_HASH",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_ChannelCloseSummary_ClosureType() Type {
    return Type{
        Name: "lnrpc_ChannelCloseSummary_ClosureType",
        Fields: []Field{
            {
                Name: "COOPERATIVE_CLOSE",
                Type: EnumVarientType,
            },
            {
                Name: "LOCAL_FORCE_CLOSE",
                Type: EnumVarientType,
            },
            {
                Name: "REMOTE_FORCE_CLOSE",
                Type: EnumVarientType,
            },
            {
                Name: "BREACH_CLOSE",
                Type: EnumVarientType,
            },
            {
                Name: "FUNDING_CANCELED",
                Type: EnumVarientType,
            },
            {
                Name: "ABANDONED",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_ChannelEventUpdate_UpdateType() Type {
    return Type{
        Name: "lnrpc_ChannelEventUpdate_UpdateType",
        Fields: []Field{
            {
                Name: "OPEN_CHANNEL",
                Type: EnumVarientType,
            },
            {
                Name: "CLOSED_CHANNEL",
                Type: EnumVarientType,
            },
            {
                Name: "ACTIVE_CHANNEL",
                Type: EnumVarientType,
            },
            {
                Name: "INACTIVE_CHANNEL",
                Type: EnumVarientType,
            },
            {
                Name: "PENDING_OPEN_CHANNEL",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_CommitmentType() Type {
    return Type{
        Name: "lnrpc_CommitmentType",
        Fields: []Field{
            {
                Name: "LEGACY",
                Description: []string{
                    "A channel using the legacy commitment format having tweaked to_remote",
                    "keys.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "STATIC_REMOTE_KEY",
                Description: []string{
                    "A channel that uses the modern commitment format where the key in the",
                    "output of the remote party does not change each state. This makes back",
                    "up and recovery easier as when the channel is closed, the funds go",
                    "directly to that key.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "ANCHORS",
                Description: []string{
                    "A channel that uses a commitment format that has anchor outputs on the",
                    "commitments, allowing fee bumping after a force close transaction has",
                    "been broadcast.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "UNKNOWN_COMMITMENT_TYPE",
                Description: []string{
                    "Returned when the commitment type isn't known or unavailable.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_Failure_FailureCode() Type {
    return Type{
        Name: "lnrpc_Failure_FailureCode",
        Fields: []Field{
            {
                Name: "RESERVED",
                Description: []string{
                    "The numbers assigned in this enumeration match the failure codes as",
                    "defined in BOLT #4. Because protobuf 3 requires enums to start with 0,",
                    "a RESERVED value is added.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "INCORRECT_OR_UNKNOWN_PAYMENT_DETAILS",
                Type: EnumVarientType,
            },
            {
                Name: "INCORRECT_PAYMENT_AMOUNT",
                Type: EnumVarientType,
            },
            {
                Name: "FINAL_INCORRECT_CLTV_EXPIRY",
                Type: EnumVarientType,
            },
            {
                Name: "FINAL_INCORRECT_HTLC_AMOUNT",
                Type: EnumVarientType,
            },
            {
                Name: "FINAL_EXPIRY_TOO_SOON",
                Type: EnumVarientType,
            },
            {
                Name: "INVALID_REALM",
                Type: EnumVarientType,
            },
            {
                Name: "EXPIRY_TOO_SOON",
                Type: EnumVarientType,
            },
            {
                Name: "INVALID_ONION_VERSION",
                Type: EnumVarientType,
            },
            {
                Name: "INVALID_ONION_HMAC",
                Type: EnumVarientType,
            },
            {
                Name: "INVALID_ONION_KEY",
                Type: EnumVarientType,
            },
            {
                Name: "AMOUNT_BELOW_MINIMUM",
                Type: EnumVarientType,
            },
            {
                Name: "FEE_INSUFFICIENT",
                Type: EnumVarientType,
            },
            {
                Name: "INCORRECT_CLTV_EXPIRY",
                Type: EnumVarientType,
            },
            {
                Name: "CHANNEL_DISABLED",
                Type: EnumVarientType,
            },
            {
                Name: "TEMPORARY_CHANNEL_FAILURE",
                Type: EnumVarientType,
            },
            {
                Name: "REQUIRED_NODE_FEATURE_MISSING",
                Type: EnumVarientType,
            },
            {
                Name: "REQUIRED_CHANNEL_FEATURE_MISSING",
                Type: EnumVarientType,
            },
            {
                Name: "UNKNOWN_NEXT_PEER",
                Type: EnumVarientType,
            },
            {
                Name: "TEMPORARY_NODE_FAILURE",
                Type: EnumVarientType,
            },
            {
                Name: "PERMANENT_NODE_FAILURE",
                Type: EnumVarientType,
            },
            {
                Name: "PERMANENT_CHANNEL_FAILURE",
                Type: EnumVarientType,
            },
            {
                Name: "EXPIRY_TOO_FAR",
                Type: EnumVarientType,
            },
            {
                Name: "MPP_TIMEOUT",
                Type: EnumVarientType,
            },
            {
                Name: "INTERNAL_FAILURE",
                Description: []string{
                    "An internal error occurred.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "UNKNOWN_FAILURE",
                Description: []string{
                    "The error source is known, but the failure itself couldn't be decoded.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "UNREADABLE_FAILURE",
                Description: []string{
                    "An unreadable failure result is returned if the received failure message",
                    "cannot be decrypted. In that case the error source is unknown.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_FeatureBit() Type {
    return Type{
        Name: "lnrpc_FeatureBit",
        Fields: []Field{
            {
                Name: "DATALOSS_PROTECT_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "DATALOSS_PROTECT_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "INITIAL_ROUING_SYNC",
                Type: EnumVarientType,
            },
            {
                Name: "UPFRONT_SHUTDOWN_SCRIPT_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "UPFRONT_SHUTDOWN_SCRIPT_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "GOSSIP_QUERIES_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "GOSSIP_QUERIES_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "TLV_ONION_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "TLV_ONION_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "EXT_GOSSIP_QUERIES_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "EXT_GOSSIP_QUERIES_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "STATIC_REMOTE_KEY_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "STATIC_REMOTE_KEY_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "PAYMENT_ADDR_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "PAYMENT_ADDR_OPT",
                Type: EnumVarientType,
            },
            {
                Name: "MPP_REQ",
                Type: EnumVarientType,
            },
            {
                Name: "MPP_OPT",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_HTLCAttempt_HTLCStatus() Type {
    return Type{
        Name: "lnrpc_HTLCAttempt_HTLCStatus",
        Fields: []Field{
            {
                Name: "IN_FLIGHT",
                Type: EnumVarientType,
            },
            {
                Name: "SUCCEEDED",
                Type: EnumVarientType,
            },
            {
                Name: "FAILED",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_Initiator() Type {
    return Type{
        Name: "lnrpc_Initiator",
        Fields: []Field{
            {
                Name: "INITIATOR_UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "INITIATOR_LOCAL",
                Type: EnumVarientType,
            },
            {
                Name: "INITIATOR_REMOTE",
                Type: EnumVarientType,
            },
            {
                Name: "INITIATOR_BOTH",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_Invoice_InvoiceState() Type {
    return Type{
        Name: "lnrpc_Invoice_InvoiceState",
        Fields: []Field{
            {
                Name: "OPEN",
                Type: EnumVarientType,
            },
            {
                Name: "SETTLED",
                Type: EnumVarientType,
            },
            {
                Name: "CANCELED",
                Type: EnumVarientType,
            },
            {
                Name: "ACCEPTED",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_InvoiceHTLCState() Type {
    return Type{
        Name: "lnrpc_InvoiceHTLCState",
        Fields: []Field{
            {
                Name: "ACCEPTED",
                Type: EnumVarientType,
            },
            {
                Name: "SETTLED",
                Type: EnumVarientType,
            },
            {
                Name: "CANCELED",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_NodeMetricType() Type {
    return Type{
        Name: "lnrpc_NodeMetricType",
        Fields: []Field{
            {
                Name: "UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "BETWEENNESS_CENTRALITY",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_Payment_PaymentStatus() Type {
    return Type{
        Name: "lnrpc_Payment_PaymentStatus",
        Fields: []Field{
            {
                Name: "UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "IN_FLIGHT",
                Type: EnumVarientType,
            },
            {
                Name: "SUCCEEDED",
                Type: EnumVarientType,
            },
            {
                Name: "FAILED",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_PaymentFailureReason() Type {
    return Type{
        Name: "lnrpc_PaymentFailureReason",
        Fields: []Field{
            {
                Name: "FAILURE_REASON_NONE",
                Description: []string{
                    "Payment isn't failed (yet).",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILURE_REASON_TIMEOUT",
                Description: []string{
                    "There are more routes to try, but the payment timeout was exceeded.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILURE_REASON_NO_ROUTE",
                Description: []string{
                    "All possible routes were tried and failed permanently. Or were no",
                    "routes to the destination at all.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILURE_REASON_ERROR",
                Description: []string{
                    "A non-recoverable error has occured.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILURE_REASON_INCORRECT_PAYMENT_DETAILS",
                Description: []string{
                    "Payment details incorrect (unknown hash, invalid amt or",
                    "invalid final cltv delta)",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILURE_REASON_INSUFFICIENT_BALANCE",
                Description: []string{
                    "Insufficient local balance.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_Peer_SyncType() Type {
    return Type{
        Name: "lnrpc_Peer_SyncType",
        Fields: []Field{
            {
                Name: "UNKNOWN_SYNC",
                Description: []string{
                    "Denotes that we cannot determine the peer's current sync type.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "ACTIVE_SYNC",
                Description: []string{
                    "Denotes that we are actively receiving new graph updates from the peer.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "PASSIVE_SYNC",
                Description: []string{
                    "Denotes that we are not receiving new graph updates from the peer.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_PeerEvent_EventType() Type {
    return Type{
        Name: "lnrpc_PeerEvent_EventType",
        Fields: []Field{
            {
                Name: "PEER_ONLINE",
                Type: EnumVarientType,
            },
            {
                Name: "PEER_OFFLINE",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_ForceClosedChannel_AnchorState() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_ForceClosedChannel_AnchorState",
        Fields: []Field{
            {
                Name: "LIMBO",
                Type: EnumVarientType,
            },
            {
                Name: "RECOVERED",
                Type: EnumVarientType,
            },
            {
                Name: "LOST",
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_ResolutionOutcome() Type {
    return Type{
        Name: "lnrpc_ResolutionOutcome",
        Fields: []Field{
            {
                Name: "OUTCOME_UNKNOWN",
                Description: []string{
                    "Outcome unknown.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "CLAIMED",
                Description: []string{
                    "An output was claimed on chain.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "UNCLAIMED",
                Description: []string{
                    "An output was left unclaimed on chain.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "ABANDONED",
                Description: []string{
                    "ResolverOutcomeAbandoned indicates that an output that we did not",
                    "claim on chain, for example an anchor that we did not sweep and a",
                    "third party claimed on chain, or a htlc that we could not decode",
                    "so left unclaimed.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FIRST_STAGE",
                Description: []string{
                    "If we force closed our channel, our htlcs need to be claimed in two",
                    "stages. This outcome represents the broadcast of a timeout or success",
                    "transaction for this two stage htlc claim.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "TIMEOUT",
                Description: []string{
                    "A htlc was timed out on chain.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mklnrpc_ResolutionType() Type {
    return Type{
        Name: "lnrpc_ResolutionType",
        Fields: []Field{
            {
                Name: "TYPE_UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "ANCHOR",
                Description: []string{
                    "We resolved an anchor output.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "INCOMING_HTLC",
                Description: []string{
                    "We are resolving an incoming htlc on chain. This if this htlc is",
                    "claimed, we swept the incoming htlc with the preimage. If it is timed",
                    "out, our peer swept the timeout path.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "OUTGOING_HTLC",
                Description: []string{
                    "We are resolving an outgoing htlc on chain. If this htlc is claimed,",
                    "the remote party swept the htlc with the preimage. If it is timed out,",
                    "we swept it with the timeout path.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "COMMIT",
                Description: []string{
                    "We force closed and need to sweep our time locked commitment output.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mkwalletrpc_WitnessType() Type {
    return Type{
        Name: "walletrpc_WitnessType",
        Fields: []Field{
            {
                Name: "UNKNOWN_WITNESS",
                Type: EnumVarientType,
            },
            {
                Name: "COMMITMENT_TIME_LOCK",
                Description: []string{
                    "A witness that allows us to spend the output of a commitment transaction",
                    "after a relative lock-time lockout.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "COMMITMENT_NO_DELAY",
                Description: []string{
                    "A witness that allows us to spend a settled no-delay output immediately on a",
                    "counterparty's commitment transaction.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "COMMITMENT_REVOKE",
                Description: []string{
                    "A witness that allows us to sweep the settled output of a malicious",
                    "counterparty's who broadcasts a revoked commitment transaction.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_OFFERED_REVOKE",
                Description: []string{
                    "A witness that allows us to sweep an HTLC which we offered to the remote",
                    "party in the case that they broadcast a revoked commitment state.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_ACCEPTED_REVOKE",
                Description: []string{
                    "A witness that allows us to sweep an HTLC output sent to us in the case that",
                    "the remote party broadcasts a revoked commitment state.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_OFFERED_TIMEOUT_SECOND_LEVEL",
                Description: []string{
                    "A witness that allows us to sweep an HTLC output that we extended to a",
                    "party, but was never fulfilled.  This HTLC output isn't directly on the",
                    "commitment transaction, but is the result of a confirmed second-level HTLC",
                    "transaction. As a result, we can only spend this after a CSV delay.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_ACCEPTED_SUCCESS_SECOND_LEVEL",
                Description: []string{
                    "A witness that allows us to sweep an HTLC output that was offered to us, and",
                    "for which we have a payment preimage. This HTLC output isn't directly on our",
                    "commitment transaction, but is the result of confirmed second-level HTLC",
                    "transaction. As a result, we can only spend this after a CSV delay.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_OFFERED_REMOTE_TIMEOUT",
                Description: []string{
                    "A witness that allows us to sweep an HTLC that we offered to the remote",
                    "party which lies in the commitment transaction of the remote party. We can",
                    "spend this output after the absolute CLTV timeout of the HTLC as passed.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_ACCEPTED_REMOTE_SUCCESS",
                Description: []string{
                    "A witness that allows us to sweep an HTLC that was offered to us by the",
                    "remote party. We use this witness in the case that the remote party goes to",
                    "chain, and we know the pre-image to the HTLC. We can sweep this without any",
                    "additional timeout.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_SECOND_LEVEL_REVOKE",
                Description: []string{
                    "A witness that allows us to sweep an HTLC from the remote party's commitment",
                    "transaction in the case that the broadcast a revoked commitment, but then",
                    "also immediately attempt to go to the second level to claim the HTLC.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "WITNESS_KEY_HASH",
                Description: []string{
                    "A witness type that allows us to spend a regular p2wkh output that's sent to",
                    "an output which is under complete control of the backing wallet.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "NESTED_WITNESS_KEY_HASH",
                Description: []string{
                    "A witness type that allows us to sweep an output that sends to a nested P2SH",
                    "script that pays to a key solely under our control.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "COMMITMENT_ANCHOR",
                Description: []string{
                    "A witness type that allows us to spend our anchor on the commitment",
                    "transaction.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mkrouterrpc_FailureDetail() Type {
    return Type{
        Name: "routerrpc_FailureDetail",
        Fields: []Field{
            {
                Name: "UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "NO_DETAIL",
                Type: EnumVarientType,
            },
            {
                Name: "ONION_DECODE",
                Type: EnumVarientType,
            },
            {
                Name: "LINK_NOT_ELIGIBLE",
                Type: EnumVarientType,
            },
            {
                Name: "ON_CHAIN_TIMEOUT",
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_EXCEEDS_MAX",
                Type: EnumVarientType,
            },
            {
                Name: "INSUFFICIENT_BALANCE",
                Type: EnumVarientType,
            },
            {
                Name: "INCOMPLETE_FORWARD",
                Type: EnumVarientType,
            },
            {
                Name: "HTLC_ADD_FAILED",
                Type: EnumVarientType,
            },
            {
                Name: "FORWARDS_DISABLED",
                Type: EnumVarientType,
            },
            {
                Name: "INVOICE_CANCELED",
                Type: EnumVarientType,
            },
            {
                Name: "INVOICE_UNDERPAID",
                Type: EnumVarientType,
            },
            {
                Name: "INVOICE_EXPIRY_TOO_SOON",
                Type: EnumVarientType,
            },
            {
                Name: "INVOICE_NOT_OPEN",
                Type: EnumVarientType,
            },
            {
                Name: "MPP_INVOICE_TIMEOUT",
                Type: EnumVarientType,
            },
            {
                Name: "ADDRESS_MISMATCH",
                Type: EnumVarientType,
            },
            {
                Name: "SET_TOTAL_MISMATCH",
                Type: EnumVarientType,
            },
            {
                Name: "SET_TOTAL_TOO_LOW",
                Type: EnumVarientType,
            },
            {
                Name: "SET_OVERPAID",
                Type: EnumVarientType,
            },
            {
                Name: "UNKNOWN_INVOICE",
                Type: EnumVarientType,
            },
            {
                Name: "INVALID_KEYSEND",
                Type: EnumVarientType,
            },
            {
                Name: "MPP_IN_PROGRESS",
                Type: EnumVarientType,
            },
            {
                Name: "CIRCULAR_ROUTE",
                Type: EnumVarientType,
            },
        },
    }
}
func mkrouterrpc_HtlcEvent_EventType() Type {
    return Type{
        Name: "routerrpc_HtlcEvent_EventType",
        Fields: []Field{
            {
                Name: "UNKNOWN",
                Type: EnumVarientType,
            },
            {
                Name: "SEND",
                Type: EnumVarientType,
            },
            {
                Name: "RECEIVE",
                Type: EnumVarientType,
            },
            {
                Name: "FORWARD",
                Type: EnumVarientType,
            },
        },
    }
}
func mkrouterrpc_PaymentState() Type {
    return Type{
        Name: "routerrpc_PaymentState",
        Fields: []Field{
            {
                Name: "IN_FLIGHT",
                Description: []string{
                    "Payment is still in flight.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "SUCCEEDED",
                Description: []string{
                    "Payment completed successfully.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILED_TIMEOUT",
                Description: []string{
                    "There are more routes to try, but the payment timeout was exceeded.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILED_NO_ROUTE",
                Description: []string{
                    "All possible routes were tried and failed permanently. Or were no",
                    "routes to the destination at all.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILED_ERROR",
                Description: []string{
                    "A non-recoverable error has occured.",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILED_INCORRECT_PAYMENT_DETAILS",
                Description: []string{
                    "Payment details incorrect (unknown hash, invalid amt or",
                    "invalid final cltv delta)",
                },
                Type: EnumVarientType,
            },
            {
                Name: "FAILED_INSUFFICIENT_BALANCE",
                Description: []string{
                    "Insufficient local balance.",
                },
                Type: EnumVarientType,
            },
        },
    }
}
func mkrouterrpc_ResolveHoldForwardAction() Type {
    return Type{
        Name: "routerrpc_ResolveHoldForwardAction",
        Fields: []Field{
            {
                Name: "SETTLE",
                Type: EnumVarientType,
            },
            {
                Name: "FAIL",
                Type: EnumVarientType,
            },
            {
                Name: "RESUME",
                Type: EnumVarientType,
            },
        },
    }
}
func mkwatchtowerrpc_GetInfoRequest() Type {
    return Type{
        Name: "watchtowerrpc_GetInfoRequest",
    }
}
func mkwatchtowerrpc_GetInfoResponse() Type {
    return Type{
        Name: "watchtowerrpc_GetInfoResponse",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The public key of the watchtower.",
                },
                Type: mkbytes(),
            },
            {
                Name: "listeners",
                Description: []string{
                    "The listening addresses of the watchtower.",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "uris",
                Description: []string{
                    "The URIs of the watchtower.",
                },
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnclipb_VersionResponse() Type {
    return Type{
        Name: "lnclipb_VersionResponse",
        Fields: []Field{
            {
                Name: "pldctl",
                Description: []string{
                    "The version information for pldctl.",
                },
                Type: mkverrpc_Version(),
            },
            {
                Name: "pld",
                Description: []string{
                    "The version information for pld.",
                },
                Type: mkverrpc_Version(),
            },
        },
    }
}
func mklnrpc_NeutrinoBan() Type {
    return Type{
        Name: "lnrpc_NeutrinoBan",
        Fields: []Field{
            {
                Name: "addr",
                Type: mkstring(),
            },
            {
                Name: "reason",
                Type: mkstring(),
            },
            {
                Name: "end_time",
                Type: mkstring(),
            },
            {
                Name: "ban_score",
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_NeutrinoInfo() Type {
    return Type{
        Name: "lnrpc_NeutrinoInfo",
        Fields: []Field{
            {
                Name: "peers",
                Repeated: true,
                Type: mklnrpc_PeerDesc(),
            },
            {
                Name: "bans",
                Repeated: true,
                Type: mklnrpc_NeutrinoBan(),
            },
            {
                Name: "queries",
                Repeated: true,
                Type: mklnrpc_NeutrinoQuery(),
            },
            {
                Name: "block_hash",
                Type: mkstring(),
            },
            {
                Name: "height",
                Type: mkint32(),
            },
            {
                Name: "block_timestamp",
                Type: mkstring(),
            },
            {
                Name: "is_syncing",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_NeutrinoQuery() Type {
    return Type{
        Name: "lnrpc_NeutrinoQuery",
        Fields: []Field{
            {
                Name: "peer",
                Type: mkstring(),
            },
            {
                Name: "command",
                Type: mkstring(),
            },
            {
                Name: "req_num",
                Type: mkuint32(),
            },
            {
                Name: "create_time",
                Type: mkuint32(),
            },
            {
                Name: "last_request_time",
                Type: mkuint32(),
            },
            {
                Name: "last_response_time",
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_PeerDesc() Type {
    return Type{
        Name: "lnrpc_PeerDesc",
        Fields: []Field{
            {
                Name: "bytes_received",
                Type: mkuint64(),
            },
            {
                Name: "bytes_sent",
                Type: mkuint64(),
            },
            {
                Name: "last_recv",
                Type: mkstring(),
            },
            {
                Name: "last_send",
                Type: mkstring(),
            },
            {
                Name: "connected",
                Type: mkbool(),
            },
            {
                Name: "addr",
                Type: mkstring(),
            },
            {
                Name: "inbound",
                Type: mkbool(),
            },
            {
                Name: "na",
                Description: []string{
                    "netaddress address:port",
                },
                Type: mkstring(),
            },
            {
                Name: "id",
                Type: mkint32(),
            },
            {
                Name: "user_agent",
                Type: mkstring(),
            },
            {
                Name: "services",
                Type: mkstring(),
            },
            {
                Name: "version_known",
                Type: mkbool(),
            },
            {
                Name: "advertised_proto_ver",
                Type: mkuint32(),
            },
            {
                Name: "protocol_version",
                Type: mkuint32(),
            },
            {
                Name: "send_headers_preferred",
                Type: mkbool(),
            },
            {
                Name: "ver_ack_received",
                Type: mkbool(),
            },
            {
                Name: "witness_enabled",
                Type: mkbool(),
            },
            {
                Name: "wire_encoding",
                Type: mkstring(),
            },
            {
                Name: "time_offset",
                Type: mkint64(),
            },
            {
                Name: "time_connected",
                Type: mkstring(),
            },
            {
                Name: "starting_height",
                Type: mkint32(),
            },
            {
                Name: "last_block",
                Type: mkint32(),
            },
            {
                Name: "last_announced_block",
                Type: mkbytes(),
            },
            {
                Name: "last_ping_nonce",
                Type: mkuint64(),
            },
            {
                Name: "last_ping_time",
                Type: mkstring(),
            },
            {
                Name: "last_ping_micros",
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_WalletInfo() Type {
    return Type{
        Name: "lnrpc_WalletInfo",
        Fields: []Field{
            {
                Name: "current_block_hash",
                Type: mkstring(),
            },
            {
                Name: "current_height",
                Type: mkint32(),
            },
            {
                Name: "current_block_timestamp",
                Type: mkstring(),
            },
            {
                Name: "wallet_version",
                Type: mkint32(),
            },
            {
                Name: "wallet_stats",
                Type: mklnrpc_WalletStats(),
            },
        },
    }
}
func mklnrpc_WalletStats() Type {
    return Type{
        Name: "lnrpc_WalletStats",
        Fields: []Field{
            {
                Name: "maintenance_in_progress",
                Type: mkbool(),
            },
            {
                Name: "maintenance_name",
                Type: mkstring(),
            },
            {
                Name: "maintenance_cycles",
                Type: mkint32(),
            },
            {
                Name: "maintenance_last_block_visited",
                Type: mkint32(),
            },
            {
                Name: "time_of_last_maintenance",
                Type: mkstring(),
            },
            {
                Name: "syncing",
                Type: mkbool(),
            },
            {
                Name: "sync_started",
                Type: mkstring(),
            },
            {
                Name: "sync_remaining_seconds",
                Type: mkint64(),
            },
            {
                Name: "sync_current_block",
                Type: mkint32(),
            },
            {
                Name: "sync_from",
                Type: mkint32(),
            },
            {
                Name: "sync_to",
                Type: mkint32(),
            },
            {
                Name: "birthday_block",
                Type: mkint32(),
            },
        },
    }
}
func mkverrpc_Version() Type {
    return Type{
        Name: "verrpc_Version",
        Fields: []Field{
            {
                Name: "commit",
                Description: []string{
                    "A verbose description of the daemon's commit.",
                },
                Type: mkstring(),
            },
            {
                Name: "commit_hash",
                Description: []string{
                    "The SHA1 commit hash that the daemon is compiled with.",
                },
                Type: mkstring(),
            },
            {
                Name: "version",
                Description: []string{
                    "The semantic version.",
                },
                Type: mkstring(),
            },
            {
                Name: "app_major",
                Description: []string{
                    "The major application version.",
                },
                Type: mkuint32(),
            },
            {
                Name: "app_minor",
                Description: []string{
                    "The minor application version.",
                },
                Type: mkuint32(),
            },
            {
                Name: "app_patch",
                Description: []string{
                    "The application patch number.",
                },
                Type: mkuint32(),
            },
            {
                Name: "app_pre_release",
                Description: []string{
                    "The application pre-release modifier, possibly empty.",
                },
                Type: mkstring(),
            },
            {
                Name: "build_tags",
                Description: []string{
                    "The list of build tags that were supplied during compilation.",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "go_version",
                Description: []string{
                    "The version of go that compiled the executable.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkverrpc_VersionRequest() Type {
    return Type{
        Name: "verrpc_VersionRequest",
    }
}
func mkchainrpc_BlockEpoch() Type {
    return Type{
        Name: "chainrpc_BlockEpoch",
        Fields: []Field{
            {
                Name: "hash",
                Description: []string{
                    "The hash of the block.",
                },
                Type: mkbytes(),
            },
            {
                Name: "height",
                Description: []string{
                    "The height of the block.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkchainrpc_ConfDetails() Type {
    return Type{
        Name: "chainrpc_ConfDetails",
        Fields: []Field{
            {
                Name: "raw_tx",
                Description: []string{
                    "The raw bytes of the confirmed transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "block_hash",
                Description: []string{
                    "The hash of the block in which the confirmed transaction was included in.",
                },
                Type: mkbytes(),
            },
            {
                Name: "block_height",
                Description: []string{
                    "The height of the block in which the confirmed transaction was included",
                    "in.",
                },
                Type: mkuint32(),
            },
            {
                Name: "tx_index",
                Description: []string{
                    "The index of the confirmed transaction within the transaction.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkchainrpc_ConfEvent() Type {
    return Type{
        Name: "chainrpc_ConfEvent",
        Fields: []Field{
            {
                Name: "conf",
                Description: []string{
                    "An event that includes the confirmation details of the request",
                    "(txid/ouput script).",
                },
                Type: mkchainrpc_ConfDetails(),
            },
            {
                Name: "reorg",
                Description: []string{
                    "An event send when the transaction of the request is reorged out of the",
                    "chain.",
                },
                Type: mkchainrpc_Reorg(),
            },
        },
    }
}
func mkchainrpc_ConfRequest() Type {
    return Type{
        Name: "chainrpc_ConfRequest",
        Fields: []Field{
            {
                Name: "txid",
                Description: []string{
                    "The transaction hash for which we should request a confirmation notification",
                    "for. If set to a hash of all zeros, then the confirmation notification will",
                    "be requested for the script instead.",
                },
                Type: mkbytes(),
            },
            {
                Name: "script",
                Description: []string{
                    "An output script within a transaction with the hash above which will be used",
                    "by light clients to match block filters. If the transaction hash is set to a",
                    "hash of all zeros, then a confirmation notification will be requested for",
                    "this script instead.",
                },
                Type: mkbytes(),
            },
            {
                Name: "num_confs",
                Description: []string{
                    "The number of desired confirmations the transaction/output script should",
                    "reach before dispatching a confirmation notification.",
                },
                Type: mkuint32(),
            },
            {
                Name: "height_hint",
                Description: []string{
                    "The earliest height in the chain for which the transaction/output script",
                    "could have been included in a block. This should in most cases be set to the",
                    "broadcast height of the transaction/output script.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkchainrpc_Outpoint() Type {
    return Type{
        Name: "chainrpc_Outpoint",
        Fields: []Field{
            {
                Name: "hash",
                Description: []string{
                    "The hash of the transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "index",
                Description: []string{
                    "The index of the output within the transaction.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkchainrpc_Reorg() Type {
    return Type{
        Name: "chainrpc_Reorg",
        Description: []string{
            "TODO(wilmer): need to know how the client will use this first.",
        },
    }
}
func mkchainrpc_SpendDetails() Type {
    return Type{
        Name: "chainrpc_SpendDetails",
        Fields: []Field{
            {
                Name: "spending_outpoint",
                Description: []string{
                    "The outpoint was that spent.",
                },
                Type: mkchainrpc_Outpoint(),
            },
            {
                Name: "raw_spending_tx",
                Description: []string{
                    "The raw bytes of the spending transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "spending_tx_hash",
                Description: []string{
                    "The hash of the spending transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "spending_input_index",
                Description: []string{
                    "The input of the spending transaction that fulfilled the spend request.",
                },
                Type: mkuint32(),
            },
            {
                Name: "spending_height",
                Description: []string{
                    "The height at which the spending transaction was included in a block.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkchainrpc_SpendEvent() Type {
    return Type{
        Name: "chainrpc_SpendEvent",
        Fields: []Field{
            {
                Name: "spend",
                Description: []string{
                    "An event that includes the details of the spending transaction of the",
                    "request (outpoint/output script).",
                },
                Type: mkchainrpc_SpendDetails(),
            },
            {
                Name: "reorg",
                Description: []string{
                    "An event sent when the spending transaction of the request was",
                    "reorged out of the chain.",
                },
                Type: mkchainrpc_Reorg(),
            },
        },
    }
}
func mkchainrpc_SpendRequest() Type {
    return Type{
        Name: "chainrpc_SpendRequest",
        Fields: []Field{
            {
                Name: "outpoint",
                Description: []string{
                    "The outpoint for which we should request a spend notification for. If set to",
                    "a zero outpoint, then the spend notification will be requested for the",
                    "script instead.",
                },
                Type: mkchainrpc_Outpoint(),
            },
            {
                Name: "script",
                Description: []string{
                    "The output script for the outpoint above. This will be used by light clients",
                    "to match block filters. If the outpoint is set to a zero outpoint, then a",
                    "spend notification will be requested for this script instead.",
                },
                Type: mkbytes(),
            },
            {
                Name: "height_hint",
                Description: []string{
                    "The earliest height in the chain for which the outpoint/output script could",
                    "have been spent. This should in most cases be set to the broadcast height of",
                    "the outpoint/output script.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkwtclientrpc_AddTowerRequest() Type {
    return Type{
        Name: "wtclientrpc_AddTowerRequest",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The identifying public key of the watchtower to add.",
                },
                Type: mkbytes(),
            },
            {
                Name: "address",
                Description: []string{
                    "A network address the watchtower is reachable over.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkwtclientrpc_AddTowerResponse() Type {
    return Type{
        Name: "wtclientrpc_AddTowerResponse",
    }
}
func mkwtclientrpc_GetTowerInfoRequest() Type {
    return Type{
        Name: "wtclientrpc_GetTowerInfoRequest",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The identifying public key of the watchtower to retrieve information for.",
                },
                Type: mkbytes(),
            },
            {
                Name: "include_sessions",
                Description: []string{
                    "Whether we should include sessions with the watchtower in the response.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwtclientrpc_ListTowersRequest() Type {
    return Type{
        Name: "wtclientrpc_ListTowersRequest",
        Fields: []Field{
            {
                Name: "include_sessions",
                Description: []string{
                    "Whether we should include sessions with the watchtower in the response.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwtclientrpc_ListTowersResponse() Type {
    return Type{
        Name: "wtclientrpc_ListTowersResponse",
        Fields: []Field{
            {
                Name: "towers",
                Description: []string{
                    "The list of watchtowers available for new backups.",
                },
                Repeated: true,
                Type: mkwtclientrpc_Tower(),
            },
        },
    }
}
func mkwtclientrpc_PolicyRequest() Type {
    return Type{
        Name: "wtclientrpc_PolicyRequest",
    }
}
func mkwtclientrpc_PolicyResponse() Type {
    return Type{
        Name: "wtclientrpc_PolicyResponse",
        Fields: []Field{
            {
                Name: "max_updates",
                Description: []string{
                    "The maximum number of updates each session we negotiate with watchtowers",
                    "should allow.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sweep_sat_per_byte",
                Description: []string{
                    "The fee rate, in satoshis per vbyte, that will be used by watchtowers for",
                    "justice transactions in response to channel breaches.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkwtclientrpc_RemoveTowerRequest() Type {
    return Type{
        Name: "wtclientrpc_RemoveTowerRequest",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The identifying public key of the watchtower to remove.",
                },
                Type: mkbytes(),
            },
            {
                Name: "address",
                Description: []string{
                    "If set, then the record for this address will be removed, indicating that is",
                    "is stale. Otherwise, the watchtower will no longer be used for future",
                    "session negotiations and backups.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkwtclientrpc_RemoveTowerResponse() Type {
    return Type{
        Name: "wtclientrpc_RemoveTowerResponse",
    }
}
func mkwtclientrpc_StatsRequest() Type {
    return Type{
        Name: "wtclientrpc_StatsRequest",
    }
}
func mkwtclientrpc_StatsResponse() Type {
    return Type{
        Name: "wtclientrpc_StatsResponse",
        Fields: []Field{
            {
                Name: "num_backups",
                Description: []string{
                    "The total number of backups made to all active and exhausted watchtower",
                    "sessions.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_pending_backups",
                Description: []string{
                    "The total number of backups that are pending to be acknowledged by all",
                    "active and exhausted watchtower sessions.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_failed_backups",
                Description: []string{
                    "The total number of backups that all active and exhausted watchtower",
                    "sessions have failed to acknowledge.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_sessions_acquired",
                Description: []string{
                    "The total number of new sessions made to watchtowers.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_sessions_exhausted",
                Description: []string{
                    "The total number of watchtower sessions that have been exhausted.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkwtclientrpc_Tower() Type {
    return Type{
        Name: "wtclientrpc_Tower",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The identifying public key of the watchtower.",
                },
                Type: mkbytes(),
            },
            {
                Name: "addresses",
                Description: []string{
                    "The list of addresses the watchtower is reachable over.",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "active_session_candidate",
                Description: []string{
                    "Whether the watchtower is currently a candidate for new sessions.",
                },
                Type: mkbool(),
            },
            {
                Name: "num_sessions",
                Description: []string{
                    "The number of sessions that have been negotiated with the watchtower.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sessions",
                Description: []string{
                    "The list of sessions that have been negotiated with the watchtower.",
                },
                Repeated: true,
                Type: mkwtclientrpc_TowerSession(),
            },
        },
    }
}
func mkwtclientrpc_TowerSession() Type {
    return Type{
        Name: "wtclientrpc_TowerSession",
        Fields: []Field{
            {
                Name: "num_backups",
                Description: []string{
                    "The total number of successful backups that have been made to the",
                    "watchtower session.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_pending_backups",
                Description: []string{
                    "The total number of backups in the session that are currently pending to be",
                    "acknowledged by the watchtower.",
                },
                Type: mkuint32(),
            },
            {
                Name: "max_backups",
                Description: []string{
                    "The maximum number of backups allowed by the watchtower session.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sweep_sat_per_byte",
                Description: []string{
                    "The fee rate, in satoshis per vbyte, that will be used by the watchtower for",
                    "the justice transaction in the event of a channel breach.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_AbandonChannelRequest() Type {
    return Type{
        Name: "lnrpc_AbandonChannelRequest",
        Fields: []Field{
            {
                Name: "channel_point",
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "pending_funding_shim_only",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_AbandonChannelResponse() Type {
    return Type{
        Name: "lnrpc_AbandonChannelResponse",
    }
}
func mklnrpc_AddInvoiceResponse() Type {
    return Type{
        Name: "lnrpc_AddInvoiceResponse",
        Fields: []Field{
            {
                Name: "r_hash",
                Type: mkbytes(),
            },
            {
                Name: "payment_request",
                Description: []string{
                    "A bare-bones invoice for a payment within the Lightning Network. With the",
                    "details of the invoice, the sender has all the data necessary to send a",
                    "payment to the recipient. Represented as bech32 encoding.",
                },
                Type: mkstring(),
            },
            {
                Name: "add_index",
                Description: []string{
                    "The \"add\" index of this invoice. Each newly created invoice will increment",
                    "this index making it monotonically increasing. Callers to the",
                    "SubscribeInvoices call can use this to instantly get notified of all added",
                    "invoices with an add_index greater than this one.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_Amount() Type {
    return Type{
        Name: "lnrpc_Amount",
        Fields: []Field{
            {
                Name: "sat",
                Description: []string{
                    "Value denominated in satoshis.",
                },
                Type: mkuint64(),
            },
            {
                Name: "msat",
                Description: []string{
                    "Value denominated in milli-satoshis.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_BcastTransactionRequest() Type {
    return Type{
        Name: "lnrpc_BcastTransactionRequest",
        Fields: []Field{
            {
                Name: "tx",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_BcastTransactionResponse() Type {
    return Type{
        Name: "lnrpc_BcastTransactionResponse",
        Fields: []Field{
            {
                Name: "txn_hash",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Chain() Type {
    return Type{
        Name: "lnrpc_Chain",
        Fields: []Field{
            {
                Name: "chain",
                Description: []string{
                    "The blockchain the node is on (eg bitcoin, litecoin)",
                },
                Type: mkstring(),
            },
            {
                Name: "network",
                Description: []string{
                    "The network the node is on (eg regtest, testnet, mainnet)",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_ChanBackupExportRequest() Type {
    return Type{
        Name: "lnrpc_ChanBackupExportRequest",
    }
}
func mklnrpc_ChanBackupSnapshot() Type {
    return Type{
        Name: "lnrpc_ChanBackupSnapshot",
        Fields: []Field{
            {
                Name: "single_chan_backups",
                Description: []string{
                    "The set of new channels that have been added since the last channel backup",
                    "snapshot was requested.",
                },
                Type: mklnrpc_ChannelBackups(),
            },
            {
                Name: "multi_chan_backup",
                Description: []string{
                    "A multi-channel backup that covers all open channels currently known to",
                    "lnd.",
                },
                Type: mklnrpc_MultiChanBackup(),
            },
        },
    }
}
func mklnrpc_ChanInfoRequest() Type {
    return Type{
        Name: "lnrpc_ChanInfoRequest",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_ChanPointShim() Type {
    return Type{
        Name: "lnrpc_ChanPointShim",
        Fields: []Field{
            {
                Name: "amt",
                Description: []string{
                    "The size of the pre-crafted output to be used as the channel point for this",
                    "channel funding.",
                },
                Type: mkint64(),
            },
            {
                Name: "chan_point",
                Description: []string{
                    "The target channel point to refrence in created commitment transactions.",
                },
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "local_key",
                Description: []string{
                    "Our local key to use when creating the multi-sig output.",
                },
                Type: mklnrpc_KeyDescriptor(),
            },
            {
                Name: "remote_key",
                Description: []string{
                    "The key of the remote party to use when creating the multi-sig output.",
                },
                Type: mkbytes(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "If non-zero, then this will be used as the pending channel ID on the wire",
                    "protocol to initate the funding request. This is an optional field, and",
                    "should only be set if the responder is already expecting a specific pending",
                    "channel ID.",
                },
                Type: mkbytes(),
            },
            {
                Name: "thaw_height",
                Description: []string{
                    "This uint32 indicates if this channel is to be considered 'frozen'. A frozen",
                    "channel does not allow a cooperative channel close by the initiator. The",
                    "thaw_height is the height that this restriction stops applying to the",
                    "channel. The height can be interpreted in two ways: as a relative height if",
                    "the value is less than 500,000, or as an absolute height otherwise.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ChangeSeedPassphraseRequest() Type {
    return Type{
        Name: "lnrpc_ChangeSeedPassphraseRequest",
        Fields: []Field{
            {
                Name: "current_seed_passphrase",
                Description: []string{
                    "current_seed_passphrase is the optional user specified passphrase that",
                    "encrypts the current seed.",
                },
                Type: mkstring(),
            },
            {
                Name: "current_seed_passphrase_bin",
                Description: []string{
                    "current_seed_passphrase_bin overrides current_seed_passphrase if specified,",
                    "for binary representation of the current seed passphrase. If using JSON then",
                    "this field must be base64 encoded.",
                },
                Type: mkbytes(),
            },
            {
                Name: "current_seed",
                Description: []string{
                    "current_seed is the seed whose passphrase is going to be changed",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "new_seed_passphrase",
                Description: []string{
                    "new_seed_passphrase is the optional user specified passphrase that will be used",
                    "to encrypt the seed.",
                },
                Type: mkstring(),
            },
            {
                Name: "new_seed_passphrase_bin",
                Description: []string{
                    "new_seed_passphrase_bin overrides new_seed_passphrase if specified, for binary",
                    "representation of the passphrase. If using JSON then this field must be base64",
                    "encoded.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_ChangeSeedPassphraseResponse() Type {
    return Type{
        Name: "lnrpc_ChangeSeedPassphraseResponse",
        Fields: []Field{
            {
                Name: "seed",
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Channel() Type {
    return Type{
        Name: "lnrpc_Channel",
        Fields: []Field{
            {
                Name: "active",
                Description: []string{
                    "Whether this channel is active or not",
                },
                Type: mkbool(),
            },
            {
                Name: "remote_pubkey",
                Description: []string{
                    "The identity pubkey of the remote node",
                },
                Type: mkbytes(),
            },
            {
                Name: "channel_point",
                Description: []string{
                    "The outpoint (txid:index) of the funding transaction. With this value, Bob",
                    "will be able to generate a signature for Alice's version of the commitment",
                    "transaction.",
                },
                Type: mkstring(),
            },
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "capacity",
                Description: []string{
                    "The total amount of funds held in this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "local_balance",
                Description: []string{
                    "This node's current balance in this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "remote_balance",
                Description: []string{
                    "The counterparty's current balance in this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "commit_fee",
                Description: []string{
                    "The amount calculated to be paid in fees for the current set of commitment",
                    "transactions. The fee amount is persisted with the channel in order to",
                    "allow the fee amount to be removed and recalculated with each channel state",
                    "update, including updates that happen after a system restart.",
                },
                Type: mkint64(),
            },
            {
                Name: "commit_weight",
                Description: []string{
                    "The weight of the commitment transaction",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_per_kw",
                Description: []string{
                    "The required number of satoshis per kilo-weight that the requester will pay",
                    "at all times, for both the funding transaction and commitment transaction.",
                    "This value can later be updated once the channel is open.",
                },
                Type: mkint64(),
            },
            {
                Name: "unsettled_balance",
                Description: []string{
                    "The unsettled balance in this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "total_satoshis_sent",
                Description: []string{
                    "The total number of satoshis we've sent within this channel.",
                },
                Type: mkint64(),
            },
            {
                Name: "total_satoshis_received",
                Description: []string{
                    "The total number of satoshis we've received within this channel.",
                },
                Type: mkint64(),
            },
            {
                Name: "num_updates",
                Description: []string{
                    "The total number of updates conducted within this channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "pending_htlcs",
                Description: []string{
                    "The list of active, uncleared HTLCs currently pending within the channel.",
                },
                Repeated: true,
                Type: mklnrpc_HTLC(),
            },
            {
                Name: "csv_delay",
                Description: []string{
                    "Deprecated. The CSV delay expressed in relative blocks. If the channel is",
                    "force closed, we will need to wait for this many blocks before we can regain",
                    "our funds.",
                },
                Type: mkuint32(),
            },
            {
                Name: "private",
                Description: []string{
                    "Whether this channel is advertised to the network or not.",
                },
                Type: mkbool(),
            },
            {
                Name: "initiator",
                Description: []string{
                    "True if we were the ones that created the channel.",
                },
                Type: mkbool(),
            },
            {
                Name: "chan_status_flags",
                Description: []string{
                    "A set of flags showing the current state of the channel.",
                },
                Type: mkstring(),
            },
            {
                Name: "local_chan_reserve_sat",
                Description: []string{
                    "Deprecated. The minimum satoshis this node is required to reserve in its",
                    "balance.",
                },
                Type: mkint64(),
            },
            {
                Name: "remote_chan_reserve_sat",
                Description: []string{
                    "Deprecated. The minimum satoshis the other node is required to reserve in",
                    "its balance.",
                },
                Type: mkint64(),
            },
            {
                Name: "static_remote_key",
                Description: []string{
                    "Deprecated. Use commitment_type.",
                },
                Type: mkbool(),
            },
            {
                Name: "commitment_type",
                Description: []string{
                    "The commitment type used by this channel.",
                },
                Type: mklnrpc_CommitmentType(),
            },
            {
                Name: "lifetime",
                Description: []string{
                    "The number of seconds that the channel has been monitored by the channel",
                    "scoring system. Scores are currently not persisted, so this value may be",
                    "less than the lifetime of the channel [EXPERIMENTAL].",
                },
                Type: mkint64(),
            },
            {
                Name: "uptime",
                Description: []string{
                    "The number of seconds that the remote peer has been observed as being online",
                    "by the channel scoring system over the lifetime of the channel",
                    "[EXPERIMENTAL].",
                },
                Type: mkint64(),
            },
            {
                Name: "close_address",
                Description: []string{
                    "Close address is the address that we will enforce payout to on cooperative",
                    "close if the channel was opened utilizing option upfront shutdown. This",
                    "value can be set on channel open by setting close_address in an open channel",
                    "request. If this value is not set, you can still choose a payout address by",
                    "cooperatively closing with the delivery_address field set.",
                },
                Type: mkstring(),
            },
            {
                Name: "push_amount_sat",
                Description: []string{
                    "The amount that the initiator of the channel optionally pushed to the remote",
                    "party on channel open. This amount will be zero if the channel initiator did",
                    "not push any funds to the remote peer. If the initiator field is true, we",
                    "pushed this amount to our peer, if it is false, the remote peer pushed this",
                    "amount to us.",
                },
                Type: mkuint64(),
            },
            {
                Name: "thaw_height",
                Description: []string{
                    "This uint32 indicates if this channel is to be considered 'frozen'. A",
                    "frozen channel doest not allow a cooperative channel close by the",
                    "initiator. The thaw_height is the height that this restriction stops",
                    "applying to the channel. This field is optional, not setting it or using a",
                    "value of zero will mean the channel has no additional restrictions. The",
                    "height can be interpreted in two ways: as a relative height if the value is",
                    "less than 500,000, or as an absolute height otherwise.",
                },
                Type: mkuint32(),
            },
            {
                Name: "local_constraints",
                Description: []string{
                    "List constraints for the local node.",
                },
                Type: mklnrpc_ChannelConstraints(),
            },
            {
                Name: "remote_constraints",
                Description: []string{
                    "List constraints for the remote node.",
                },
                Type: mklnrpc_ChannelConstraints(),
            },
        },
    }
}
func mklnrpc_ChannelAcceptRequest() Type {
    return Type{
        Name: "lnrpc_ChannelAcceptRequest",
        Fields: []Field{
            {
                Name: "node_pubkey",
                Description: []string{
                    "The pubkey of the node that wishes to open an inbound channel.",
                },
                Type: mkbytes(),
            },
            {
                Name: "chain_hash",
                Description: []string{
                    "The hash of the genesis block that the proposed channel resides in.",
                },
                Type: mkbytes(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel id.",
                },
                Type: mkbytes(),
            },
            {
                Name: "funding_amt",
                Description: []string{
                    "The funding amount in satoshis that initiator wishes to use in the",
                    "channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "push_amt",
                Description: []string{
                    "The push amount of the proposed channel in millisatoshis.",
                },
                Type: mkuint64(),
            },
            {
                Name: "dust_limit",
                Description: []string{
                    "The dust limit of the initiator's commitment tx.",
                },
                Type: mkuint64(),
            },
            {
                Name: "max_value_in_flight",
                Description: []string{
                    "The maximum amount of coins in millisatoshis that can be pending in this",
                    "channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "channel_reserve",
                Description: []string{
                    "The minimum amount of satoshis the initiator requires us to have at all",
                    "times.",
                },
                Type: mkuint64(),
            },
            {
                Name: "min_htlc",
                Description: []string{
                    "The smallest HTLC in millisatoshis that the initiator will accept.",
                },
                Type: mkuint64(),
            },
            {
                Name: "fee_per_kw",
                Description: []string{
                    "The initial fee rate that the initiator suggests for both commitment",
                    "transactions.",
                },
                Type: mkuint64(),
            },
            {
                Name: "csv_delay",
                Description: []string{
                    "The number of blocks to use for the relative time lock in the pay-to-self",
                    "output of both commitment transactions.",
                },
                Type: mkuint32(),
            },
            {
                Name: "max_accepted_htlcs",
                Description: []string{
                    "The total number of incoming HTLC's that the initiator will accept.",
                },
                Type: mkuint32(),
            },
            {
                Name: "channel_flags",
                Description: []string{
                    "A bit-field which the initiator uses to specify proposed channel",
                    "behavior.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ChannelAcceptResponse() Type {
    return Type{
        Name: "lnrpc_ChannelAcceptResponse",
        Fields: []Field{
            {
                Name: "accept",
                Description: []string{
                    "Whether or not the client accepts the channel.",
                },
                Type: mkbool(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel id to which this response applies.",
                },
                Type: mkbytes(),
            },
            {
                Name: "error",
                Description: []string{
                    "An optional error to send the initiating party to indicate why the channel",
                    "was rejected. This field *should not* contain sensitive information, it will",
                    "be sent to the initiating party. This field should only be set if accept is",
                    "false, the channel will be rejected if an error is set with accept=true",
                    "because the meaning of this response is ambiguous. Limited to 500",
                    "characters.",
                },
                Type: mkstring(),
            },
            {
                Name: "upfront_shutdown",
                Description: []string{
                    "The upfront shutdown address to use if the initiating peer supports option",
                    "upfront shutdown script (see ListPeers for the features supported). Note",
                    "that the channel open will fail if this value is set for a peer that does",
                    "not support this feature bit.",
                },
                Type: mkstring(),
            },
            {
                Name: "csv_delay",
                Description: []string{
                    "The csv delay (in blocks) that we require for the remote party.",
                },
                Type: mkuint32(),
            },
            {
                Name: "reserve_sat",
                Description: []string{
                    "The reserve amount in satoshis that we require the remote peer to adhere to.",
                    "We require that the remote peer always have some reserve amount allocated to",
                    "them so that there is always a disincentive to broadcast old state (if they",
                    "hold 0 sats on their side of the channel, there is nothing to lose).",
                },
                Type: mkuint64(),
            },
            {
                Name: "in_flight_max_msat",
                Description: []string{
                    "The maximum amount of funds in millisatoshis that we allow the remote peer",
                    "to have in outstanding htlcs.",
                },
                Type: mkuint64(),
            },
            {
                Name: "max_htlc_count",
                Description: []string{
                    "The maximum number of htlcs that the remote peer can offer us.",
                },
                Type: mkuint32(),
            },
            {
                Name: "min_htlc_in",
                Description: []string{
                    "The minimum value in millisatoshis for incoming htlcs on the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "min_accept_depth",
                Description: []string{
                    "The number of confirmations we require before we consider the channel open.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ChannelBackup() Type {
    return Type{
        Name: "lnrpc_ChannelBackup",
        Fields: []Field{
            {
                Name: "chan_point",
                Description: []string{
                    "Identifies the channel that this backup belongs to.",
                },
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "chan_backup",
                Description: []string{
                    "Is an encrypted single-chan backup. this can be passed to",
                    "RestoreChannelBackups, or the WalletUnlocker Init and Unlock methods in",
                    "order to trigger the recovery protocol.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_ChannelBackupSubscription() Type {
    return Type{
        Name: "lnrpc_ChannelBackupSubscription",
    }
}
func mklnrpc_ChannelBackups() Type {
    return Type{
        Name: "lnrpc_ChannelBackups",
        Fields: []Field{
            {
                Name: "chan_backups",
                Description: []string{
                    "A set of single-chan static channel backups.",
                },
                Repeated: true,
                Type: mklnrpc_ChannelBackup(),
            },
        },
    }
}
func mklnrpc_ChannelBalanceRequest() Type {
    return Type{
        Name: "lnrpc_ChannelBalanceRequest",
    }
}
func mklnrpc_ChannelBalanceResponse() Type {
    return Type{
        Name: "lnrpc_ChannelBalanceResponse",
        Fields: []Field{
            {
                Name: "balance",
                Description: []string{
                    "Deprecated. Sum of channels balances denominated in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "pending_open_balance",
                Description: []string{
                    "Deprecated. Sum of channels pending balances denominated in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "local_balance",
                Description: []string{
                    "Sum of channels local balances.",
                },
                Type: mklnrpc_Amount(),
            },
            {
                Name: "remote_balance",
                Description: []string{
                    "Sum of channels remote balances.",
                },
                Type: mklnrpc_Amount(),
            },
            {
                Name: "unsettled_local_balance",
                Description: []string{
                    "Sum of channels local unsettled balances.",
                },
                Type: mklnrpc_Amount(),
            },
            {
                Name: "unsettled_remote_balance",
                Description: []string{
                    "Sum of channels remote unsettled balances.",
                },
                Type: mklnrpc_Amount(),
            },
            {
                Name: "pending_open_local_balance",
                Description: []string{
                    "Sum of channels pending local balances.",
                },
                Type: mklnrpc_Amount(),
            },
            {
                Name: "pending_open_remote_balance",
                Description: []string{
                    "Sum of channels pending remote balances.",
                },
                Type: mklnrpc_Amount(),
            },
        },
    }
}
func mklnrpc_ChannelCloseSummary() Type {
    return Type{
        Name: "lnrpc_ChannelCloseSummary",
        Fields: []Field{
            {
                Name: "channel_point",
                Description: []string{
                    "The outpoint (txid:index) of the funding transaction.",
                },
                Type: mkstring(),
            },
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chain_hash",
                Description: []string{
                    "The hash of the genesis block that this channel resides within.",
                },
                Type: mkstring(),
            },
            {
                Name: "closing_tx_hash",
                Description: []string{
                    "The txid of the transaction which ultimately closed this channel.",
                },
                Type: mkstring(),
            },
            {
                Name: "remote_pubkey",
                Description: []string{
                    "Public key of the remote peer that we formerly had a channel with.",
                },
                Type: mkbytes(),
            },
            {
                Name: "capacity",
                Description: []string{
                    "Total capacity of the channel.",
                },
                Type: mkint64(),
            },
            {
                Name: "close_height",
                Description: []string{
                    "Height at which the funding transaction was spent.",
                },
                Type: mkuint32(),
            },
            {
                Name: "settled_balance",
                Description: []string{
                    "Settled balance at the time of channel closure",
                },
                Type: mkint64(),
            },
            {
                Name: "time_locked_balance",
                Description: []string{
                    "The sum of all the time-locked outputs at the time of channel closure",
                },
                Type: mkint64(),
            },
            {
                Name: "close_type",
                Description: []string{
                    "Details on how the channel was closed.",
                },
                Type: mklnrpc_ChannelCloseSummary_ClosureType(),
            },
            {
                Name: "open_initiator",
                Description: []string{
                    "Open initiator is the party that initiated opening the channel. Note that",
                    "this value may be unknown if the channel was closed before we migrated to",
                    "store open channel information after close.",
                },
                Type: mklnrpc_Initiator(),
            },
            {
                Name: "close_initiator",
                Description: []string{
                    "Close initiator indicates which party initiated the close. This value will",
                    "be unknown for channels that were cooperatively closed before we started",
                    "tracking cooperative close initiators. Note that this indicates which party",
                    "initiated a close, and it is possible for both to initiate cooperative or",
                    "force closes, although only one party's close will be confirmed on chain.",
                },
                Type: mklnrpc_Initiator(),
            },
            {
                Name: "resolutions",
                Repeated: true,
                Type: mklnrpc_Resolution(),
            },
        },
    }
}
func mklnrpc_ChannelCloseUpdate() Type {
    return Type{
        Name: "lnrpc_ChannelCloseUpdate",
        Fields: []Field{
            {
                Name: "closing_txid",
                Type: mkbytes(),
            },
            {
                Name: "success",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ChannelConstraints() Type {
    return Type{
        Name: "lnrpc_ChannelConstraints",
        Fields: []Field{
            {
                Name: "csv_delay",
                Description: []string{
                    "The CSV delay expressed in relative blocks. If the channel is force closed,",
                    "we will need to wait for this many blocks before we can regain our funds.",
                },
                Type: mkuint32(),
            },
            {
                Name: "chan_reserve_sat",
                Description: []string{
                    "The minimum satoshis this node is required to reserve in its balance.",
                },
                Type: mkuint64(),
            },
            {
                Name: "dust_limit_sat",
                Description: []string{
                    "The dust limit (in satoshis) of the initiator's commitment tx.",
                },
                Type: mkuint64(),
            },
            {
                Name: "max_pending_amt_msat",
                Description: []string{
                    "The maximum amount of coins in millisatoshis that can be pending in this",
                    "channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "min_htlc_msat",
                Description: []string{
                    "The smallest HTLC in millisatoshis that the initiator will accept.",
                },
                Type: mkuint64(),
            },
            {
                Name: "max_accepted_htlcs",
                Description: []string{
                    "The total number of incoming HTLC's that the initiator will accept.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ChannelEdge() Type {
    return Type{
        Name: "lnrpc_ChannelEdge",
        Description: []string{
            "A fully authenticated channel along with all its unique attributes.",
            "Once an authenticated channel announcement has been processed on the network,",
            "then an instance of ChannelEdgeInfo encapsulating the channels attributes is",
            "stored. The other portions relevant to routing policy of a channel are stored",
            "within a ChannelEdgePolicy for each direction of the channel.",
        },
        Fields: []Field{
            {
                Name: "channel_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chan_point",
                Type: mkstring(),
            },
            {
                Name: "last_update",
                Type: mkuint32(),
            },
            {
                Name: "node1_pub",
                Type: mkbytes(),
            },
            {
                Name: "node2_pub",
                Type: mkbytes(),
            },
            {
                Name: "capacity",
                Type: mkint64(),
            },
            {
                Name: "node1_policy",
                Type: mklnrpc_RoutingPolicy(),
            },
            {
                Name: "node2_policy",
                Type: mklnrpc_RoutingPolicy(),
            },
        },
    }
}
func mklnrpc_ChannelEdgeUpdate() Type {
    return Type{
        Name: "lnrpc_ChannelEdgeUpdate",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chan_point",
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "capacity",
                Type: mkint64(),
            },
            {
                Name: "routing_policy",
                Type: mklnrpc_RoutingPolicy(),
            },
            {
                Name: "advertising_node",
                Type: mkstring(),
            },
            {
                Name: "connecting_node",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_ChannelEventSubscription() Type {
    return Type{
        Name: "lnrpc_ChannelEventSubscription",
    }
}
func mklnrpc_ChannelEventUpdate() Type {
    return Type{
        Name: "lnrpc_ChannelEventUpdate",
        Fields: []Field{
            {
                Name: "open_channel",
                Type: mklnrpc_Channel(),
            },
            {
                Name: "closed_channel",
                Type: mklnrpc_ChannelCloseSummary(),
            },
            {
                Name: "active_channel",
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "inactive_channel",
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "pending_open_channel",
                Type: mklnrpc_PendingUpdate(),
            },
            {
                Name: "type",
                Type: mklnrpc_ChannelEventUpdate_UpdateType(),
            },
        },
    }
}
func mklnrpc_ChannelFeeReport() Type {
    return Type{
        Name: "lnrpc_ChannelFeeReport",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The short channel id that this fee report belongs to.",
                },
                Type: mkuint64(),
            },
            {
                Name: "channel_point",
                Description: []string{
                    "The channel that this fee report belongs to.",
                },
                Type: mkstring(),
            },
            {
                Name: "base_fee_msat",
                Description: []string{
                    "The base fee charged regardless of the number of milli-satoshis sent.",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_per_mil",
                Description: []string{
                    "The amount charged per milli-satoshis transferred expressed in",
                    "millionths of a satoshi.",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_rate",
                Description: []string{
                    "The effective fee rate in milli-satoshis. Computed by dividing the",
                    "fee_per_mil value by 1 million.",
                },
                Type: mkdouble(),
            },
        },
    }
}
func mklnrpc_ChannelGraph() Type {
    return Type{
        Name: "lnrpc_ChannelGraph",
        Description: []string{
            "Returns a new instance of the directed channel graph.",
        },
        Fields: []Field{
            {
                Name: "nodes",
                Description: []string{
                    "The list of `LightningNode`s in this channel graph",
                },
                Repeated: true,
                Type: mklnrpc_LightningNode(),
            },
            {
                Name: "edges",
                Description: []string{
                    "The list of `ChannelEdge`s in this channel graph",
                },
                Repeated: true,
                Type: mklnrpc_ChannelEdge(),
            },
        },
    }
}
func mklnrpc_ChannelGraphRequest() Type {
    return Type{
        Name: "lnrpc_ChannelGraphRequest",
        Fields: []Field{
            {
                Name: "include_unannounced",
                Description: []string{
                    "Whether unannounced channels are included in the response or not. If set,",
                    "unannounced channels are included. Unannounced channels are both private",
                    "channels, and public channels that are not yet announced to the network.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ChannelOpenUpdate() Type {
    return Type{
        Name: "lnrpc_ChannelOpenUpdate",
        Fields: []Field{
            {
                Name: "channel_point",
                Type: mklnrpc_ChannelPoint(),
            },
        },
    }
}
func mklnrpc_ChannelPoint() Type {
    return Type{
        Name: "lnrpc_ChannelPoint",
        Fields: []Field{
            {
                Name: "funding_txid_bytes",
                Description: []string{
                    "Txid of the funding transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "funding_txid_str",
                Description: []string{
                    "Hex-encoded string representing the byte-reversed hash of the funding",
                    "transaction.",
                },
                Type: mkstring(),
            },
            {
                Name: "output_index",
                Description: []string{
                    "The index of the output of the funding transaction",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ChannelUpdate() Type {
    return Type{
        Name: "lnrpc_ChannelUpdate",
        Fields: []Field{
            {
                Name: "signature",
                Description: []string{
                    "The signature that validates the announced data and proves the ownership",
                    "of node id.",
                },
                Type: mkbytes(),
            },
            {
                Name: "chain_hash",
                Description: []string{
                    "The target chain that this channel was opened within. This value",
                    "should be the genesis hash of the target chain. Along with the short",
                    "channel ID, this uniquely identifies the channel globally in a",
                    "blockchain.",
                },
                Type: mkbytes(),
            },
            {
                Name: "chan_id",
                Description: []string{
                    "The unique description of the funding transaction.",
                },
                Type: mkuint64(),
            },
            {
                Name: "timestamp",
                Description: []string{
                    "A timestamp that allows ordering in the case of multiple announcements.",
                    "We should ignore the message if timestamp is not greater than the",
                    "last-received.",
                },
                Type: mkuint32(),
            },
            {
                Name: "message_flags",
                Description: []string{
                    "The bitfield that describes whether optional fields are present in this",
                    "update. Currently, the least-significant bit must be set to 1 if the",
                    "optional field MaxHtlc is present.",
                },
                Type: mkuint32(),
            },
            {
                Name: "channel_flags",
                Description: []string{
                    "The bitfield that describes additional meta-data concerning how the",
                    "update is to be interpreted. Currently, the least-significant bit must be",
                    "set to 0 if the creating node corresponds to the first node in the",
                    "previously sent channel announcement and 1 otherwise. If the second bit",
                    "is set, then the channel is set to be disabled.",
                },
                Type: mkuint32(),
            },
            {
                Name: "time_lock_delta",
                Description: []string{
                    "The minimum number of blocks this node requires to be added to the expiry",
                    "of HTLCs. This is a security parameter determined by the node operator.",
                    "This value represents the required gap between the time locks of the",
                    "incoming and outgoing HTLC's set to this node.",
                },
                Type: mkuint32(),
            },
            {
                Name: "htlc_minimum_msat",
                Description: []string{
                    "The minimum HTLC value which will be accepted.",
                },
                Type: mkuint64(),
            },
            {
                Name: "base_fee",
                Description: []string{
                    "The base fee that must be used for incoming HTLC's to this particular",
                    "channel. This value will be tacked onto the required for a payment",
                    "independent of the size of the payment.",
                },
                Type: mkuint32(),
            },
            {
                Name: "fee_rate",
                Description: []string{
                    "The fee rate that will be charged per millionth of a satoshi.",
                },
                Type: mkuint32(),
            },
            {
                Name: "htlc_maximum_msat",
                Description: []string{
                    "The maximum HTLC value which will be accepted.",
                },
                Type: mkuint64(),
            },
            {
                Name: "extra_opaque_data",
                Description: []string{
                    "The set of data that was appended to this message, some of which we may",
                    "not actually know how to iterate or parse. By holding onto this data, we",
                    "ensure that we're able to properly validate the set of signatures that",
                    "cover these new fields, and ensure we're able to make upgrades to the",
                    "network in a forwards compatible manner.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_CloseChannelRequest() Type {
    return Type{
        Name: "lnrpc_CloseChannelRequest",
        Fields: []Field{
            {
                Name: "channel_point",
                Description: []string{
                    "The outpoint (txid:index) of the funding transaction. With this value, Bob",
                    "will be able to generate a signature for Alice's version of the commitment",
                    "transaction.",
                },
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "force",
                Description: []string{
                    "If true, then the channel will be closed forcibly. This means the",
                    "current commitment transaction will be signed and broadcast.",
                },
                Type: mkbool(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that the closure transaction should be",
                    "confirmed by.",
                },
                Type: mkint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "A manual fee rate set in sat/byte that should be used when crafting the",
                    "closure transaction.",
                },
                Type: mkint64(),
            },
            {
                Name: "delivery_address",
                Description: []string{
                    "An optional address to send funds to in the case of a cooperative close.",
                    "If the channel was opened with an upfront shutdown script and this field",
                    "is set, the request to close will fail because the channel must pay out",
                    "to the upfront shutdown addresss.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_CloseStatusUpdate() Type {
    return Type{
        Name: "lnrpc_CloseStatusUpdate",
        Fields: []Field{
            {
                Name: "close_pending",
                Type: mklnrpc_PendingUpdate(),
            },
            {
                Name: "chan_close",
                Type: mklnrpc_ChannelCloseUpdate(),
            },
        },
    }
}
func mklnrpc_ClosedChannelUpdate() Type {
    return Type{
        Name: "lnrpc_ClosedChannelUpdate",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "capacity",
                Type: mkint64(),
            },
            {
                Name: "closed_height",
                Type: mkuint32(),
            },
            {
                Name: "chan_point",
                Type: mklnrpc_ChannelPoint(),
            },
        },
    }
}
func mklnrpc_ClosedChannelsRequest() Type {
    return Type{
        Name: "lnrpc_ClosedChannelsRequest",
        Fields: []Field{
            {
                Name: "cooperative",
                Type: mkbool(),
            },
            {
                Name: "local_force",
                Type: mkbool(),
            },
            {
                Name: "remote_force",
                Type: mkbool(),
            },
            {
                Name: "breach",
                Type: mkbool(),
            },
            {
                Name: "funding_canceled",
                Type: mkbool(),
            },
            {
                Name: "abandoned",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ClosedChannelsResponse() Type {
    return Type{
        Name: "lnrpc_ClosedChannelsResponse",
        Fields: []Field{
            {
                Name: "channels",
                Repeated: true,
                Type: mklnrpc_ChannelCloseSummary(),
            },
        },
    }
}
func mklnrpc_ConfirmationUpdate() Type {
    return Type{
        Name: "lnrpc_ConfirmationUpdate",
        Fields: []Field{
            {
                Name: "block_sha",
                Type: mkbytes(),
            },
            {
                Name: "block_height",
                Type: mkint32(),
            },
            {
                Name: "num_confs_left",
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ConnectPeerRequest() Type {
    return Type{
        Name: "lnrpc_ConnectPeerRequest",
        Fields: []Field{
            {
                Name: "addr",
                Description: []string{
                    "Lightning address of the peer, in the format `<pubkey>@host`",
                },
                Type: mklnrpc_LightningAddress(),
            },
            {
                Name: "perm",
                Description: []string{
                    "If set, the daemon will attempt to persistently connect to the target",
                    "peer. Otherwise, the call will be synchronous.",
                },
                Type: mkbool(),
            },
            {
                Name: "timeout",
                Description: []string{
                    "The connection timeout value (in seconds) for this request. It won't affect",
                    "other requests.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_ConnectPeerResponse() Type {
    return Type{
        Name: "lnrpc_ConnectPeerResponse",
    }
}
func mklnrpc_CreateTransactionRequest() Type {
    return Type{
        Name: "lnrpc_CreateTransactionRequest",
        Fields: []Field{
            {
                Name: "to_address",
                Description: []string{
                    "Address which we will be paying to",
                },
                Type: mkstring(),
            },
            {
                Name: "amount",
                Description: []string{
                    "Number of PKT to send",
                },
                Type: mkdouble(),
            },
            {
                Name: "from_address",
                Description: []string{
                    "Addresses which can be selected for sourcing funds from",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "electrum_format",
                Description: []string{
                    "Output an electrum format transaction",
                },
                Type: mkbool(),
            },
            {
                Name: "change_address",
                Type: mkstring(),
            },
            {
                Name: "input_min_height",
                Type: mkint32(),
            },
            {
                Name: "min_conf",
                Type: mkint32(),
            },
            {
                Name: "vote",
                Type: mkbool(),
            },
            {
                Name: "max_inputs",
                Type: mkint32(),
            },
            {
                Name: "autolock",
                Type: mkstring(),
            },
            {
                Name: "sign",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_CreateTransactionResponse() Type {
    return Type{
        Name: "lnrpc_CreateTransactionResponse",
        Fields: []Field{
            {
                Name: "transaction",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_DebugLevelRequest() Type {
    return Type{
        Name: "lnrpc_DebugLevelRequest",
        Fields: []Field{
            {
                Name: "show",
                Type: mkbool(),
            },
            {
                Name: "level_spec",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_DebugLevelResponse() Type {
    return Type{
        Name: "lnrpc_DebugLevelResponse",
        Fields: []Field{
            {
                Name: "sub_systems",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_DecodeRawTransactionRequest() Type {
    return Type{
        Name: "lnrpc_DecodeRawTransactionRequest",
        Fields: []Field{
            {
                Name: "hex_tx",
                Type: mkstring(),
            },
            {
                Name: "vin_extra",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_DecodeRawTransactionResponse() Type {
    return Type{
        Name: "lnrpc_DecodeRawTransactionResponse",
        Fields: []Field{
            {
                Name: "txid",
                Type: mkstring(),
            },
            {
                Name: "version",
                Type: mkint32(),
            },
            {
                Name: "locktime",
                Type: mkuint32(),
            },
            {
                Name: "sfee",
                Type: mkstring(),
            },
            {
                Name: "size",
                Type: mkint32(),
            },
            {
                Name: "vsize",
                Type: mkint32(),
            },
            {
                Name: "vin",
                Repeated: true,
                Type: mklnrpc_VinPrevOut(),
            },
            {
                Name: "vout",
                Repeated: true,
                Type: mklnrpc_Vout(),
            },
        },
    }
}
func mklnrpc_DeleteAllPaymentsRequest() Type {
    return Type{
        Name: "lnrpc_DeleteAllPaymentsRequest",
    }
}
func mklnrpc_DeleteAllPaymentsResponse() Type {
    return Type{
        Name: "lnrpc_DeleteAllPaymentsResponse",
    }
}
func mklnrpc_DisconnectPeerRequest() Type {
    return Type{
        Name: "lnrpc_DisconnectPeerRequest",
        Fields: []Field{
            {
                Name: "pub_key",
                Description: []string{
                    "The pubkey of the node to disconnect from",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_DisconnectPeerResponse() Type {
    return Type{
        Name: "lnrpc_DisconnectPeerResponse",
    }
}
func mklnrpc_DumpPrivKeyRequest() Type {
    return Type{
        Name: "lnrpc_DumpPrivKeyRequest",
        Fields: []Field{
            {
                Name: "address",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_DumpPrivKeyResponse() Type {
    return Type{
        Name: "lnrpc_DumpPrivKeyResponse",
        Fields: []Field{
            {
                Name: "private_key",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_EdgeLocator() Type {
    return Type{
        Name: "lnrpc_EdgeLocator",
        Fields: []Field{
            {
                Name: "channel_id",
                Description: []string{
                    "The short channel id of this edge.",
                },
                Type: mkuint64(),
            },
            {
                Name: "direction_reverse",
                Description: []string{
                    "The direction of this edge. If direction_reverse is false, the direction",
                    "of this edge is from the channel endpoint with the lexicographically smaller",
                    "pub key to the endpoint with the larger pub key. If direction_reverse is",
                    "is true, the edge goes the other way.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_EstimateFeeRequest() Type {
    return Type{
        Name: "lnrpc_EstimateFeeRequest",
        Fields: []Field{
            {
                Name: "AddrToAmount",
                Description: []string{
                    "The map from addresses to amounts for the transaction.",
                },
                Repeated: true,
                Type: mklnrpc_EstimateFeeRequest_AddrToAmountEntry(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that this transaction should be confirmed",
                    "by.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_EstimateFeeRequest_AddrToAmountEntry() Type {
    return Type{
        Name: "lnrpc_EstimateFeeRequest_AddrToAmountEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_EstimateFeeResponse() Type {
    return Type{
        Name: "lnrpc_EstimateFeeResponse",
        Fields: []Field{
            {
                Name: "fee_sat",
                Description: []string{
                    "The total fee in satoshis.",
                },
                Type: mkint64(),
            },
            {
                Name: "feerate_sat_per_byte",
                Description: []string{
                    "The fee rate in satoshi/byte.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_ExportChannelBackupRequest() Type {
    return Type{
        Name: "lnrpc_ExportChannelBackupRequest",
        Fields: []Field{
            {
                Name: "chan_point",
                Description: []string{
                    "The target channel point to obtain a back up for.",
                },
                Type: mklnrpc_ChannelPoint(),
            },
        },
    }
}
func mklnrpc_Failure() Type {
    return Type{
        Name: "lnrpc_Failure",
        Fields: []Field{
            {
                Name: "code",
                Description: []string{
                    "Failure code as defined in the Lightning spec",
                },
                Type: mklnrpc_Failure_FailureCode(),
            },
            {
                Name: "channel_update",
                Description: []string{
                    "An optional channel update message.",
                },
                Type: mklnrpc_ChannelUpdate(),
            },
            {
                Name: "htlc_msat",
                Description: []string{
                    "A failure type-dependent htlc value.",
                },
                Type: mkuint64(),
            },
            {
                Name: "onion_sha_256",
                Description: []string{
                    "The sha256 sum of the onion payload.",
                },
                Type: mkbytes(),
            },
            {
                Name: "cltv_expiry",
                Description: []string{
                    "A failure type-dependent cltv expiry value.",
                },
                Type: mkuint32(),
            },
            {
                Name: "flags",
                Description: []string{
                    "A failure type-dependent flags value.",
                },
                Type: mkuint32(),
            },
            {
                Name: "failure_source_index",
                Description: []string{
                    "The position in the path of the intermediate or final node that generated",
                    "the failure message. Position zero is the sender node.",
                },
                Type: mkuint32(),
            },
            {
                Name: "height",
                Description: []string{
                    "A failure type-dependent block height.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_Feature() Type {
    return Type{
        Name: "lnrpc_Feature",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "is_required",
                Type: mkbool(),
            },
            {
                Name: "is_known",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_FeeLimit() Type {
    return Type{
        Name: "lnrpc_FeeLimit",
        Fields: []Field{
            {
                Name: "fixed",
                Description: []string{
                    "The fee limit expressed as a fixed amount of satoshis.",
                    "",
                    "The fields fixed and fixed_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "fixed_msat",
                Description: []string{
                    "The fee limit expressed as a fixed amount of millisatoshis.",
                    "",
                    "The fields fixed and fixed_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "percent",
                Description: []string{
                    "The fee limit expressed as a percentage of the payment amount.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_FeeReportRequest() Type {
    return Type{
        Name: "lnrpc_FeeReportRequest",
    }
}
func mklnrpc_FeeReportResponse() Type {
    return Type{
        Name: "lnrpc_FeeReportResponse",
        Fields: []Field{
            {
                Name: "channel_fees",
                Description: []string{
                    "An array of channel fee reports which describes the current fee schedule",
                    "for each channel.",
                },
                Repeated: true,
                Type: mklnrpc_ChannelFeeReport(),
            },
            {
                Name: "day_fee_sum",
                Description: []string{
                    "The total amount of fee revenue (in satoshis) the switch has collected",
                    "over the past 24 hrs.",
                },
                Type: mkuint64(),
            },
            {
                Name: "week_fee_sum",
                Description: []string{
                    "The total amount of fee revenue (in satoshis) the switch has collected",
                    "over the past 1 week.",
                },
                Type: mkuint64(),
            },
            {
                Name: "month_fee_sum",
                Description: []string{
                    "The total amount of fee revenue (in satoshis) the switch has collected",
                    "over the past 1 month.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_FloatMetric() Type {
    return Type{
        Name: "lnrpc_FloatMetric",
        Fields: []Field{
            {
                Name: "value",
                Description: []string{
                    "Arbitrary float value.",
                },
                Type: mkdouble(),
            },
            {
                Name: "normalized_value",
                Description: []string{
                    "The value normalized to [0,1] or [-1,1].",
                },
                Type: mkdouble(),
            },
        },
    }
}
func mklnrpc_ForwardingEvent() Type {
    return Type{
        Name: "lnrpc_ForwardingEvent",
        Fields: []Field{
            {
                Name: "timestamp",
                Description: []string{
                    "Timestamp is the time (unix epoch offset) that this circuit was",
                    "completed.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chan_id_in",
                Description: []string{
                    "The incoming channel ID that carried the HTLC that created the circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chan_id_out",
                Description: []string{
                    "The outgoing channel ID that carried the preimage that completed the",
                    "circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_in",
                Description: []string{
                    "The total amount (in satoshis) of the incoming HTLC that created half",
                    "the circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_out",
                Description: []string{
                    "The total amount (in satoshis) of the outgoing HTLC that created the",
                    "second half of the circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "fee",
                Description: []string{
                    "The total fee (in satoshis) that this payment circuit carried.",
                },
                Type: mkuint64(),
            },
            {
                Name: "fee_msat",
                Description: []string{
                    "The total fee (in milli-satoshis) that this payment circuit carried.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_in_msat",
                Description: []string{
                    "The total amount (in milli-satoshis) of the incoming HTLC that created",
                    "half the circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_out_msat",
                Description: []string{
                    "The total amount (in milli-satoshis) of the outgoing HTLC that created",
                    "the second half of the circuit.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_ForwardingHistoryRequest() Type {
    return Type{
        Name: "lnrpc_ForwardingHistoryRequest",
        Fields: []Field{
            {
                Name: "start_time",
                Description: []string{
                    "Start time is the starting point of the forwarding history request. All",
                    "records beyond this point will be included, respecting the end time, and",
                    "the index offset.",
                },
                Type: mkuint64(),
            },
            {
                Name: "end_time",
                Description: []string{
                    "End time is the end point of the forwarding history request. The",
                    "response will carry at most 50k records between the start time and the",
                    "end time. The index offset can be used to implement pagination.",
                },
                Type: mkuint64(),
            },
            {
                Name: "index_offset",
                Description: []string{
                    "Index offset is the offset in the time series to start at. As each",
                    "response can only contain 50k records, callers can use this to skip",
                    "around within a packed time series.",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_max_events",
                Description: []string{
                    "The max number of events to return in the response to this query.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ForwardingHistoryResponse() Type {
    return Type{
        Name: "lnrpc_ForwardingHistoryResponse",
        Fields: []Field{
            {
                Name: "forwarding_events",
                Description: []string{
                    "A list of forwarding events from the time slice of the time series",
                    "specified in the request.",
                },
                Repeated: true,
                Type: mklnrpc_ForwardingEvent(),
            },
            {
                Name: "last_offset_index",
                Description: []string{
                    "The index of the last time in the set of returned forwarding events. Can",
                    "be used to seek further, pagination style.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_FundingPsbtFinalize() Type {
    return Type{
        Name: "lnrpc_FundingPsbtFinalize",
        Fields: []Field{
            {
                Name: "signed_psbt",
                Description: []string{
                    "The funded PSBT that contains all witness data to send the exact channel",
                    "capacity amount to the PK script returned in the open channel message in a",
                    "previous step. Cannot be set at the same time as final_raw_tx.",
                },
                Type: mkbytes(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel ID of the channel to get the PSBT for.",
                },
                Type: mkbytes(),
            },
            {
                Name: "final_raw_tx",
                Description: []string{
                    "As an alternative to the signed PSBT with all witness data, the final raw",
                    "wire format transaction can also be specified directly. Cannot be set at the",
                    "same time as signed_psbt.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_FundingPsbtVerify() Type {
    return Type{
        Name: "lnrpc_FundingPsbtVerify",
        Fields: []Field{
            {
                Name: "funded_psbt",
                Description: []string{
                    "The funded but not yet signed PSBT that sends the exact channel capacity",
                    "amount to the PK script returned in the open channel message in a previous",
                    "step.",
                },
                Type: mkbytes(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel ID of the channel to get the PSBT for.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_FundingShim() Type {
    return Type{
        Name: "lnrpc_FundingShim",
        Fields: []Field{
            {
                Name: "chan_point_shim",
                Description: []string{
                    "A channel shim where the channel point was fully constructed outside",
                    "of lnd's wallet and the transaction might already be published.",
                },
                Type: mklnrpc_ChanPointShim(),
            },
            {
                Name: "psbt_shim",
                Description: []string{
                    "A channel shim that uses a PSBT to fund and sign the channel funding",
                    "transaction.",
                },
                Type: mklnrpc_PsbtShim(),
            },
        },
    }
}
func mklnrpc_FundingShimCancel() Type {
    return Type{
        Name: "lnrpc_FundingShimCancel",
        Fields: []Field{
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel ID of the channel to cancel the funding shim for.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_FundingStateStepResp() Type {
    return Type{
        Name: "lnrpc_FundingStateStepResp",
    }
}
func mklnrpc_FundingTransitionMsg() Type {
    return Type{
        Name: "lnrpc_FundingTransitionMsg",
        Fields: []Field{
            {
                Name: "shim_register",
                Description: []string{
                    "The funding shim to register. This should be used before any",
                    "channel funding has began by the remote party, as it is intended as a",
                    "preparatory step for the full channel funding.",
                },
                Type: mklnrpc_FundingShim(),
            },
            {
                Name: "shim_cancel",
                Description: []string{
                    "Used to cancel an existing registered funding shim.",
                },
                Type: mklnrpc_FundingShimCancel(),
            },
            {
                Name: "psbt_verify",
                Description: []string{
                    "Used to continue a funding flow that was initiated to be executed",
                    "through a PSBT. This step verifies that the PSBT contains the correct",
                    "outputs to fund the channel.",
                },
                Type: mklnrpc_FundingPsbtVerify(),
            },
            {
                Name: "psbt_finalize",
                Description: []string{
                    "Used to continue a funding flow that was initiated to be executed",
                    "through a PSBT. This step finalizes the funded and signed PSBT, finishes",
                    "negotiation with the peer and finally publishes the resulting funding",
                    "transaction.",
                },
                Type: mklnrpc_FundingPsbtFinalize(),
            },
        },
    }
}
func mklnrpc_GetAddressBalancesRequest() Type {
    return Type{
        Name: "lnrpc_GetAddressBalancesRequest",
        Fields: []Field{
            {
                Name: "minconf",
                Description: []string{
                    "Minimum number of confirmations for coins to be considered received",
                },
                Type: mkint32(),
            },
            {
                Name: "showzerobalance",
                Description: []string{
                    "If true then addresses which have been created but carry zero balance will be included",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_GetAddressBalancesResponse() Type {
    return Type{
        Name: "lnrpc_GetAddressBalancesResponse",
        Fields: []Field{
            {
                Name: "addrs",
                Repeated: true,
                Type: mklnrpc_GetAddressBalancesResponseAddr(),
            },
        },
    }
}
func mklnrpc_GetAddressBalancesResponseAddr() Type {
    return Type{
        Name: "lnrpc_GetAddressBalancesResponseAddr",
        Fields: []Field{
            {
                Name: "address",
                Description: []string{
                    "The address which has this balance",
                },
                Type: mkstring(),
            },
            {
                Name: "total",
                Description: []string{
                    "Total balance in coins",
                },
                Type: mkdouble(),
            },
            {
                Name: "stotal",
                Description: []string{
                    "Total balance (atomic units)",
                },
                Type: mkint64(),
            },
            {
                Name: "spendable",
                Description: []string{
                    "Balance which is currently spendable (coins)",
                },
                Type: mkdouble(),
            },
            {
                Name: "sspendable",
                Description: []string{
                    "Balance which is currently spendable (atomic units)",
                },
                Type: mkint64(),
            },
            {
                Name: "immaturereward",
                Description: []string{
                    "Mined coins which have not yet matured (coins)",
                },
                Type: mkdouble(),
            },
            {
                Name: "simmaturereward",
                Description: []string{
                    "Mined coins which have not yet matured (atomic units)",
                },
                Type: mkint64(),
            },
            {
                Name: "unconfirmed",
                Description: []string{
                    "Unconfirmed balance in coins",
                },
                Type: mkdouble(),
            },
            {
                Name: "sunconfirmed",
                Description: []string{
                    "Unconfirmed balance in atomic units",
                },
                Type: mkint64(),
            },
            {
                Name: "outputcount",
                Description: []string{
                    "The number of transaction outputs which make up the balance",
                },
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_GetInfoRequest() Type {
    return Type{
        Name: "lnrpc_GetInfoRequest",
    }
}
func mklnrpc_GetInfoResponse() Type {
    return Type{
        Name: "lnrpc_GetInfoResponse",
        Fields: []Field{
            {
                Name: "version",
                Description: []string{
                    "The version of the LND software that the node is running.",
                },
                Type: mkstring(),
            },
            {
                Name: "commit_hash",
                Description: []string{
                    "The SHA1 commit hash that the daemon is compiled with.",
                },
                Type: mkstring(),
            },
            {
                Name: "identity_pubkey",
                Description: []string{
                    "The identity pubkey of the current node.",
                },
                Type: mkbytes(),
            },
            {
                Name: "alias",
                Description: []string{
                    "If applicable, the alias of the current node, e.g. \"bob\"",
                },
                Type: mkstring(),
            },
            {
                Name: "color",
                Description: []string{
                    "The color of the current node in hex code format",
                },
                Type: mkstring(),
            },
            {
                Name: "num_pending_channels",
                Description: []string{
                    "Number of pending channels",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_active_channels",
                Description: []string{
                    "Number of active channels",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_inactive_channels",
                Description: []string{
                    "Number of inactive channels",
                },
                Type: mkuint32(),
            },
            {
                Name: "num_peers",
                Description: []string{
                    "Number of peers",
                },
                Type: mkuint32(),
            },
            {
                Name: "block_height",
                Description: []string{
                    "The node's current view of the height of the best block",
                },
                Type: mkuint32(),
            },
            {
                Name: "block_hash",
                Description: []string{
                    "The node's current view of the hash of the best block",
                },
                Type: mkstring(),
            },
            {
                Name: "best_header_timestamp",
                Description: []string{
                    "Timestamp of the block best known to the wallet",
                },
                Type: mkint64(),
            },
            {
                Name: "synced_to_chain",
                Description: []string{
                    "Whether the wallet's view is synced to the main chain",
                },
                Type: mkbool(),
            },
            {
                Name: "synced_to_graph",
                Description: []string{
                    "Whether we consider ourselves synced with the public channel graph.",
                },
                Type: mkbool(),
            },
            {
                Name: "testnet",
                Description: []string{
                    "Whether the current node is connected to testnet. This field is",
                    "deprecated and the network field should be used instead",
                },
                Type: mkbool(),
            },
            {
                Name: "chains",
                Description: []string{
                    "A list of active chains the node is connected to",
                },
                Repeated: true,
                Type: mklnrpc_Chain(),
            },
            {
                Name: "uris",
                Description: []string{
                    "The URIs of the current node.",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "features",
                Description: []string{
                    "Features that our node has advertised in our init message, node",
                    "announcements and invoices.",
                },
                Repeated: true,
                Type: mklnrpc_GetInfoResponse_FeaturesEntry(),
            },
        },
    }
}
func mklnrpc_GetInfoResponse_FeaturesEntry() Type {
    return Type{
        Name: "lnrpc_GetInfoResponse_FeaturesEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint32(),
            },
            {
                Name: "value",
                Type: mklnrpc_Feature(),
            },
        },
    }
}
func mklnrpc_GetNetworkStewardVoteRequest() Type {
    return Type{
        Name: "lnrpc_GetNetworkStewardVoteRequest",
    }
}
func mklnrpc_GetNetworkStewardVoteResponse() Type {
    return Type{
        Name: "lnrpc_GetNetworkStewardVoteResponse",
        Fields: []Field{
            {
                Name: "vote_against",
                Type: mkstring(),
            },
            {
                Name: "vote_for",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_GetNewAddressRequest() Type {
    return Type{
        Name: "lnrpc_GetNewAddressRequest",
        Fields: []Field{
            {
                Name: "legacy",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_GetNewAddressResponse() Type {
    return Type{
        Name: "lnrpc_GetNewAddressResponse",
        Fields: []Field{
            {
                Name: "address",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_GetRecoveryInfoRequest() Type {
    return Type{
        Name: "lnrpc_GetRecoveryInfoRequest",
    }
}
func mklnrpc_GetRecoveryInfoResponse() Type {
    return Type{
        Name: "lnrpc_GetRecoveryInfoResponse",
        Fields: []Field{
            {
                Name: "recovery_mode",
                Description: []string{
                    "Whether the wallet is in recovery mode",
                },
                Type: mkbool(),
            },
            {
                Name: "recovery_finished",
                Description: []string{
                    "Whether the wallet recovery progress is finished",
                },
                Type: mkbool(),
            },
            {
                Name: "progress",
                Description: []string{
                    "The recovery progress, ranging from 0 to 1.",
                },
                Type: mkdouble(),
            },
        },
    }
}
func mklnrpc_GetSecretRequest() Type {
    return Type{
        Name: "lnrpc_GetSecretRequest",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_GetSecretResponse() Type {
    return Type{
        Name: "lnrpc_GetSecretResponse",
        Fields: []Field{
            {
                Name: "secret",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_GetTransactionDetailsResult() Type {
    return Type{
        Name: "lnrpc_GetTransactionDetailsResult",
        Fields: []Field{
            {
                Name: "address",
                Type: mkstring(),
            },
            {
                Name: "amount",
                Type: mkdouble(),
            },
            {
                Name: "category",
                Type: mkstring(),
            },
            {
                Name: "vout",
                Type: mkuint32(),
            },
            {
                Name: "amount_units",
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_GetTransactionRequest() Type {
    return Type{
        Name: "lnrpc_GetTransactionRequest",
        Fields: []Field{
            {
                Name: "txid",
                Type: mkstring(),
            },
            {
                Name: "includewatchonly",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_GetTransactionResponse() Type {
    return Type{
        Name: "lnrpc_GetTransactionResponse",
        Fields: []Field{
            {
                Name: "transaction",
                Type: mklnrpc_TransactionResult(),
            },
        },
    }
}
func mklnrpc_GetTransactionsRequest() Type {
    return Type{
        Name: "lnrpc_GetTransactionsRequest",
        Fields: []Field{
            {
                Name: "start_height",
                Description: []string{
                    "The height from which to list transactions, inclusive.",
                },
                Type: mkint32(),
            },
            {
                Name: "end_height",
                Description: []string{
                    "The height until which to list transactions, inclusive. To include",
                    "unconfirmed transactions, this value should be set to -1, which will",
                    "return transactions from start_height until the current chain tip and",
                    "unconfirmed transactions. If no end_height is provided, the call will",
                    "default to this option.",
                },
                Type: mkint32(),
            },
            {
                Name: "txns_limit",
                Type: mkint32(),
            },
            {
                Name: "txns_skip",
                Type: mkint32(),
            },
            {
                Name: "coinbase",
                Type: mkint32(),
            },
            {
                Name: "reversed",
                Description: []string{
                    "If set, the payments returned will result from seeking backwards from the",
                    "specified index offset. This can be used to paginate backwards. The order",
                    "of the returned payments is always oldest first (ascending index order).",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_GetWalletSeedRequest() Type {
    return Type{
        Name: "lnrpc_GetWalletSeedRequest",
    }
}
func mklnrpc_GetWalletSeedResponse() Type {
    return Type{
        Name: "lnrpc_GetWalletSeedResponse",
        Fields: []Field{
            {
                Name: "seed",
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_GraphTopologySubscription() Type {
    return Type{
        Name: "lnrpc_GraphTopologySubscription",
    }
}
func mklnrpc_GraphTopologyUpdate() Type {
    return Type{
        Name: "lnrpc_GraphTopologyUpdate",
        Fields: []Field{
            {
                Name: "node_updates",
                Repeated: true,
                Type: mklnrpc_NodeUpdate(),
            },
            {
                Name: "channel_updates",
                Repeated: true,
                Type: mklnrpc_ChannelEdgeUpdate(),
            },
            {
                Name: "closed_chans",
                Repeated: true,
                Type: mklnrpc_ClosedChannelUpdate(),
            },
        },
    }
}
func mklnrpc_HTLC() Type {
    return Type{
        Name: "lnrpc_HTLC",
        Fields: []Field{
            {
                Name: "incoming",
                Type: mkbool(),
            },
            {
                Name: "amount",
                Type: mkint64(),
            },
            {
                Name: "hash_lock",
                Type: mkbytes(),
            },
            {
                Name: "expiration_height",
                Type: mkuint32(),
            },
            {
                Name: "htlc_index",
                Description: []string{
                    "Index identifying the htlc on the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "forwarding_channel",
                Description: []string{
                    "If this HTLC is involved in a forwarding operation, this field indicates",
                    "the forwarding channel. For an outgoing htlc, it is the incoming channel.",
                    "For an incoming htlc, it is the outgoing channel. When the htlc",
                    "originates from this node or this node is the final destination,",
                    "forwarding_channel will be zero. The forwarding channel will also be zero",
                    "for htlcs that need to be forwarded but don't have a forwarding decision",
                    "persisted yet.",
                },
                Type: mkuint64(),
            },
            {
                Name: "forwarding_htlc_index",
                Description: []string{
                    "Index identifying the htlc on the forwarding channel.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_HTLCAttempt() Type {
    return Type{
        Name: "lnrpc_HTLCAttempt",
        Fields: []Field{
            {
                Name: "status",
                Description: []string{
                    "The status of the HTLC.",
                },
                Type: mklnrpc_HTLCAttempt_HTLCStatus(),
            },
            {
                Name: "route",
                Description: []string{
                    "The route taken by this HTLC.",
                },
                Type: mklnrpc_Route(),
            },
            {
                Name: "attempt_time_ns",
                Description: []string{
                    "The time in UNIX nanoseconds at which this HTLC was sent.",
                },
                Type: mkint64(),
            },
            {
                Name: "resolve_time_ns",
                Description: []string{
                    "The time in UNIX nanoseconds at which this HTLC was settled or failed.",
                    "This value will not be set if the HTLC is still IN_FLIGHT.",
                },
                Type: mkint64(),
            },
            {
                Name: "failure",
                Description: []string{
                    "Detailed htlc failure info.",
                },
                Type: mklnrpc_Failure(),
            },
            {
                Name: "preimage",
                Description: []string{
                    "The preimage that was used to settle the HTLC.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_Hop() Type {
    return Type{
        Name: "lnrpc_Hop",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The unique channel ID for the channel. The first 3 bytes are the block",
                    "height, the next 3 the index within the block, and the last 2 bytes are the",
                    "output index for the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "chan_capacity",
                Type: mkint64(),
            },
            {
                Name: "amt_to_forward",
                Type: mkint64(),
            },
            {
                Name: "fee",
                Type: mkint64(),
            },
            {
                Name: "expiry",
                Type: mkuint32(),
            },
            {
                Name: "amt_to_forward_msat",
                Type: mkint64(),
            },
            {
                Name: "fee_msat",
                Type: mkint64(),
            },
            {
                Name: "pub_key",
                Description: []string{
                    "An optional public key of the hop. If the public key is given, the payment",
                    "can be executed without relying on a copy of the channel graph.",
                },
                Type: mkstring(),
            },
            {
                Name: "tlv_payload",
                Description: []string{
                    "If set to true, then this hop will be encoded using the new variable length",
                    "TLV format. Note that if any custom tlv_records below are specified, then",
                    "this field MUST be set to true for them to be encoded properly.",
                },
                Type: mkbool(),
            },
            {
                Name: "mpp_record",
                Description: []string{
                    "An optional TLV record that signals the use of an MPP payment. If present,",
                    "the receiver will enforce that that the same mpp_record is included in the",
                    "final hop payload of all non-zero payments in the HTLC set. If empty, a",
                    "regular single-shot payment is or was attempted.",
                },
                Type: mklnrpc_MPPRecord(),
            },
            {
                Name: "custom_records",
                Description: []string{
                    "An optional set of key-value TLV records. This is useful within the context",
                    "of the SendToRoute call as it allows callers to specify arbitrary K-V pairs",
                    "to drop off at each hop within the onion.",
                },
                Repeated: true,
                Type: mklnrpc_Hop_CustomRecordsEntry(),
            },
        },
    }
}
func mklnrpc_Hop_CustomRecordsEntry() Type {
    return Type{
        Name: "lnrpc_Hop_CustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_HopHint() Type {
    return Type{
        Name: "lnrpc_HopHint",
        Fields: []Field{
            {
                Name: "node_id",
                Description: []string{
                    "The public key of the node at the start of the channel.",
                },
                Type: mkbytes(),
            },
            {
                Name: "chan_id",
                Description: []string{
                    "The unique identifier of the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "fee_base_msat",
                Description: []string{
                    "The base fee of the channel denominated in millisatoshis.",
                },
                Type: mkuint32(),
            },
            {
                Name: "fee_proportional_millionths",
                Description: []string{
                    "The fee rate of the channel for sending one satoshi across it denominated in",
                    "millionths of a satoshi.",
                },
                Type: mkuint32(),
            },
            {
                Name: "cltv_expiry_delta",
                Description: []string{
                    "The time-lock delta of the channel.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ImportPrivKeyRequest() Type {
    return Type{
        Name: "lnrpc_ImportPrivKeyRequest",
        Fields: []Field{
            {
                Name: "private_key",
                Type: mkstring(),
            },
            {
                Name: "rescan",
                Type: mkbool(),
            },
            {
                Name: "legacy",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ImportPrivKeyResponse() Type {
    return Type{
        Name: "lnrpc_ImportPrivKeyResponse",
        Fields: []Field{
            {
                Name: "address",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Invoice() Type {
    return Type{
        Name: "lnrpc_Invoice",
        Fields: []Field{
            {
                Name: "memo",
                Description: []string{
                    "An optional memo to attach along with the invoice. Used for record keeping",
                    "purposes for the invoice's creator, and will also be set in the description",
                    "field of the encoded payment request if the description_hash field is not",
                    "being used.",
                },
                Type: mkstring(),
            },
            {
                Name: "r_preimage",
                Description: []string{
                    "The 32 byte preimage which will allow settling an incoming HTLC.",
                },
                Type: mkbytes(),
            },
            {
                Name: "r_hash",
                Description: []string{
                    "The hash of the preimage.",
                },
                Type: mkbytes(),
            },
            {
                Name: "value",
                Description: []string{
                    "The value of this invoice in satoshis",
                    "",
                    "The fields value and value_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "value_msat",
                Description: []string{
                    "The value of this invoice in millisatoshis",
                    "",
                    "The fields value and value_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "settled",
                Description: []string{
                    "Whether this invoice has been fulfilled",
                },
                Type: mkbool(),
            },
            {
                Name: "creation_date",
                Description: []string{
                    "When this invoice was created",
                },
                Type: mkint64(),
            },
            {
                Name: "settle_date",
                Description: []string{
                    "When this invoice was settled",
                },
                Type: mkint64(),
            },
            {
                Name: "payment_request",
                Description: []string{
                    "A bare-bones invoice for a payment within the Lightning Network. With the",
                    "details of the invoice, the sender has all the data necessary to send a",
                    "payment to the recipient.",
                },
                Type: mkstring(),
            },
            {
                Name: "description_hash",
                Description: []string{
                    "Hash (SHA-256) of a description of the payment. Used if the description of",
                    "payment (memo) is too long to naturally fit within the description field",
                    "of an encoded payment request.",
                },
                Type: mkbytes(),
            },
            {
                Name: "expiry",
                Description: []string{
                    "Payment request expiry time in seconds. Default is 3600 (1 hour).",
                },
                Type: mkint64(),
            },
            {
                Name: "fallback_addr",
                Description: []string{
                    "Fallback on-chain address.",
                },
                Type: mkstring(),
            },
            {
                Name: "cltv_expiry",
                Description: []string{
                    "Delta to use for the time-lock of the CLTV extended to the final hop.",
                },
                Type: mkuint64(),
            },
            {
                Name: "route_hints",
                Description: []string{
                    "Route hints that can each be individually used to assist in reaching the",
                    "invoice's destination.",
                },
                Repeated: true,
                Type: mklnrpc_RouteHint(),
            },
            {
                Name: "private",
                Description: []string{
                    "Whether this invoice should include routing hints for private channels.",
                },
                Type: mkbool(),
            },
            {
                Name: "add_index",
                Description: []string{
                    "The \"add\" index of this invoice. Each newly created invoice will increment",
                    "this index making it monotonically increasing. Callers to the",
                    "SubscribeInvoices call can use this to instantly get notified of all added",
                    "invoices with an add_index greater than this one.",
                },
                Type: mkuint64(),
            },
            {
                Name: "settle_index",
                Description: []string{
                    "The \"settle\" index of this invoice. Each newly settled invoice will",
                    "increment this index making it monotonically increasing. Callers to the",
                    "SubscribeInvoices call can use this to instantly get notified of all",
                    "settled invoices with an settle_index greater than this one.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_paid_sat",
                Description: []string{
                    "The amount that was accepted for this invoice, in satoshis. This will ONLY",
                    "be set if this invoice has been settled. We provide this field as if the",
                    "invoice was created with a zero value, then we need to record what amount",
                    "was ultimately accepted. Additionally, it's possible that the sender paid",
                    "MORE that was specified in the original invoice. So we'll record that here",
                    "as well.",
                },
                Type: mkint64(),
            },
            {
                Name: "amt_paid_msat",
                Description: []string{
                    "The amount that was accepted for this invoice, in millisatoshis. This will",
                    "ONLY be set if this invoice has been settled. We provide this field as if",
                    "the invoice was created with a zero value, then we need to record what",
                    "amount was ultimately accepted. Additionally, it's possible that the sender",
                    "paid MORE that was specified in the original invoice. So we'll record that",
                    "here as well.",
                },
                Type: mkint64(),
            },
            {
                Name: "state",
                Description: []string{
                    "The state the invoice is in.",
                },
                Type: mklnrpc_Invoice_InvoiceState(),
            },
            {
                Name: "htlcs",
                Description: []string{
                    "List of HTLCs paying to this invoice [EXPERIMENTAL].",
                },
                Repeated: true,
                Type: mklnrpc_InvoiceHTLC(),
            },
            {
                Name: "features",
                Description: []string{
                    "List of features advertised on the invoice.",
                },
                Repeated: true,
                Type: mklnrpc_Invoice_FeaturesEntry(),
            },
            {
                Name: "is_keysend",
                Description: []string{
                    "Indicates if this invoice was a spontaneous payment that arrived via keysend",
                    "[EXPERIMENTAL].",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_Invoice_FeaturesEntry() Type {
    return Type{
        Name: "lnrpc_Invoice_FeaturesEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint32(),
            },
            {
                Name: "value",
                Type: mklnrpc_Feature(),
            },
        },
    }
}
func mklnrpc_InvoiceHTLC() Type {
    return Type{
        Name: "lnrpc_InvoiceHTLC",
        Description: []string{
            "Details of an HTLC that paid to an invoice",
        },
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "Short channel id over which the htlc was received.",
                },
                Type: mkuint64(),
            },
            {
                Name: "htlc_index",
                Description: []string{
                    "Index identifying the htlc on the channel.",
                },
                Type: mkuint64(),
            },
            {
                Name: "amt_msat",
                Description: []string{
                    "The amount of the htlc in msat.",
                },
                Type: mkuint64(),
            },
            {
                Name: "accept_height",
                Description: []string{
                    "Block height at which this htlc was accepted.",
                },
                Type: mkint32(),
            },
            {
                Name: "accept_time",
                Description: []string{
                    "Time at which this htlc was accepted.",
                },
                Type: mkint64(),
            },
            {
                Name: "resolve_time",
                Description: []string{
                    "Time at which this htlc was settled or canceled.",
                },
                Type: mkint64(),
            },
            {
                Name: "expiry_height",
                Description: []string{
                    "Block height at which this htlc expires.",
                },
                Type: mkint32(),
            },
            {
                Name: "state",
                Description: []string{
                    "Current state the htlc is in.",
                },
                Type: mklnrpc_InvoiceHTLCState(),
            },
            {
                Name: "custom_records",
                Description: []string{
                    "Custom tlv records.",
                },
                Repeated: true,
                Type: mklnrpc_InvoiceHTLC_CustomRecordsEntry(),
            },
            {
                Name: "mpp_total_amt_msat",
                Description: []string{
                    "The total amount of the mpp payment in msat.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_InvoiceHTLC_CustomRecordsEntry() Type {
    return Type{
        Name: "lnrpc_InvoiceHTLC_CustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_InvoiceSubscription() Type {
    return Type{
        Name: "lnrpc_InvoiceSubscription",
        Fields: []Field{
            {
                Name: "add_index",
                Description: []string{
                    "If specified (non-zero), then we'll first start by sending out",
                    "notifications for all added indexes with an add_index greater than this",
                    "value. This allows callers to catch up on any events they missed while they",
                    "weren't connected to the streaming RPC.",
                },
                Type: mkuint64(),
            },
            {
                Name: "settle_index",
                Description: []string{
                    "If specified (non-zero), then we'll first start by sending out",
                    "notifications for all settled indexes with an settle_index greater than",
                    "this value. This allows callers to catch up on any events they missed while",
                    "they weren't connected to the streaming RPC.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_KeyDescriptor() Type {
    return Type{
        Name: "lnrpc_KeyDescriptor",
        Fields: []Field{
            {
                Name: "raw_key_bytes",
                Description: []string{
                    "The raw bytes of the key being identified.",
                },
                Type: mkbytes(),
            },
            {
                Name: "key_loc",
                Description: []string{
                    "The key locator that identifies which key to use for signing.",
                },
                Type: mklnrpc_KeyLocator(),
            },
        },
    }
}
func mklnrpc_KeyLocator() Type {
    return Type{
        Name: "lnrpc_KeyLocator",
        Fields: []Field{
            {
                Name: "key_family",
                Description: []string{
                    "The family of key being identified.",
                },
                Type: mkint32(),
            },
            {
                Name: "key_index",
                Description: []string{
                    "The precise index of the key being identified.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_LightningAddress() Type {
    return Type{
        Name: "lnrpc_LightningAddress",
        Fields: []Field{
            {
                Name: "pubkey",
                Description: []string{
                    "The identity pubkey of the Lightning node",
                },
                Type: mkstring(),
            },
            {
                Name: "host",
                Description: []string{
                    "The network location of the lightning node, e.g. `69.69.69.69:1337` or",
                    "`localhost:10011`",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_LightningNode() Type {
    return Type{
        Name: "lnrpc_LightningNode",
        Description: []string{
            "An individual vertex/node within the channel graph. A node is",
            "connected to other nodes by one or more channel edges emanating from it. As the",
            "graph is directed, a node will also have an incoming edge attached to it for",
            "each outgoing edge.",
        },
        Fields: []Field{
            {
                Name: "last_update",
                Type: mkuint32(),
            },
            {
                Name: "pub_key",
                Description: []string{
                    "The public key of the node",
                },
                Type: mkbytes(),
            },
            {
                Name: "alias",
                Type: mkstring(),
            },
            {
                Name: "addresses",
                Repeated: true,
                Type: mklnrpc_NodeAddress(),
            },
            {
                Name: "color",
                Type: mkstring(),
            },
            {
                Name: "features",
                Repeated: true,
                Type: mklnrpc_LightningNode_FeaturesEntry(),
            },
        },
    }
}
func mklnrpc_LightningNode_FeaturesEntry() Type {
    return Type{
        Name: "lnrpc_LightningNode_FeaturesEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint32(),
            },
            {
                Name: "value",
                Type: mklnrpc_Feature(),
            },
        },
    }
}
func mklnrpc_ListChannelsRequest() Type {
    return Type{
        Name: "lnrpc_ListChannelsRequest",
        Fields: []Field{
            {
                Name: "active_only",
                Type: mkbool(),
            },
            {
                Name: "inactive_only",
                Type: mkbool(),
            },
            {
                Name: "public_only",
                Type: mkbool(),
            },
            {
                Name: "private_only",
                Type: mkbool(),
            },
            {
                Name: "peer",
                Description: []string{
                    "Filters the response for channels with a target peer's pubkey. If peer is",
                    "empty, all channels will be returned.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_ListChannelsResponse() Type {
    return Type{
        Name: "lnrpc_ListChannelsResponse",
        Fields: []Field{
            {
                Name: "channels",
                Description: []string{
                    "The list of active channels",
                },
                Repeated: true,
                Type: mklnrpc_Channel(),
            },
        },
    }
}
func mklnrpc_ListInvoiceRequest() Type {
    return Type{
        Name: "lnrpc_ListInvoiceRequest",
        Fields: []Field{
            {
                Name: "pending_only",
                Description: []string{
                    "If set, only invoices that are not settled and not canceled will be returned",
                    "in the response.",
                },
                Type: mkbool(),
            },
            {
                Name: "index_offset",
                Description: []string{
                    "The index of an invoice that will be used as either the start or end of a",
                    "query to determine which invoices should be returned in the response.",
                },
                Type: mkuint64(),
            },
            {
                Name: "num_max_invoices",
                Description: []string{
                    "The max number of invoices to return in the response to this query.",
                },
                Type: mkuint64(),
            },
            {
                Name: "reversed",
                Description: []string{
                    "If set, the invoices returned will result from seeking backwards from the",
                    "specified index offset. This can be used to paginate backwards.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ListInvoiceResponse() Type {
    return Type{
        Name: "lnrpc_ListInvoiceResponse",
        Fields: []Field{
            {
                Name: "invoices",
                Description: []string{
                    "A list of invoices from the time slice of the time series specified in the",
                    "request.",
                },
                Repeated: true,
                Type: mklnrpc_Invoice(),
            },
            {
                Name: "last_index_offset",
                Description: []string{
                    "The index of the last item in the set of returned invoices. This can be used",
                    "to seek further, pagination style.",
                },
                Type: mkuint64(),
            },
            {
                Name: "first_index_offset",
                Description: []string{
                    "The index of the last item in the set of returned invoices. This can be used",
                    "to seek backwards, pagination style.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_ListLockUnspentRequest() Type {
    return Type{
        Name: "lnrpc_ListLockUnspentRequest",
    }
}
func mklnrpc_ListLockUnspentResponse() Type {
    return Type{
        Name: "lnrpc_ListLockUnspentResponse",
        Fields: []Field{
            {
                Name: "locked_unspent",
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_ListPaymentsRequest() Type {
    return Type{
        Name: "lnrpc_ListPaymentsRequest",
        Fields: []Field{
            {
                Name: "include_incomplete",
                Description: []string{
                    "If true, then return payments that have not yet fully completed. This means",
                    "that pending payments, as well as failed payments will show up if this",
                    "field is set to true. This flag doesn't change the meaning of the indices,",
                    "which are tied to individual payments.",
                },
                Type: mkbool(),
            },
            {
                Name: "index_offset",
                Description: []string{
                    "The index of a payment that will be used as either the start or end of a",
                    "query to determine which payments should be returned in the response. The",
                    "index_offset is exclusive. In the case of a zero index_offset, the query",
                    "will start with the oldest payment when paginating forwards, or will end",
                    "with the most recent payment when paginating backwards.",
                },
                Type: mkuint64(),
            },
            {
                Name: "max_payments",
                Description: []string{
                    "The maximal number of payments returned in the response to this query.",
                },
                Type: mkuint64(),
            },
            {
                Name: "reversed",
                Description: []string{
                    "If set, the payments returned will result from seeking backwards from the",
                    "specified index offset. This can be used to paginate backwards. The order",
                    "of the returned payments is always oldest first (ascending index order).",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ListPaymentsResponse() Type {
    return Type{
        Name: "lnrpc_ListPaymentsResponse",
        Fields: []Field{
            {
                Name: "payments",
                Description: []string{
                    "The list of payments",
                },
                Repeated: true,
                Type: mklnrpc_Payment(),
            },
            {
                Name: "first_index_offset",
                Description: []string{
                    "The index of the first item in the set of returned payments. This can be",
                    "used as the index_offset to continue seeking backwards in the next request.",
                },
                Type: mkuint64(),
            },
            {
                Name: "last_index_offset",
                Description: []string{
                    "The index of the last item in the set of returned payments. This can be used",
                    "as the index_offset to continue seeking forwards in the next request.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_ListPeersRequest() Type {
    return Type{
        Name: "lnrpc_ListPeersRequest",
        Fields: []Field{
            {
                Name: "latest_error",
                Description: []string{
                    "If true, only the last error that our peer sent us will be returned with",
                    "the peer's information, rather than the full set of historic errors we have",
                    "stored.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ListPeersResponse() Type {
    return Type{
        Name: "lnrpc_ListPeersResponse",
        Fields: []Field{
            {
                Name: "peers",
                Description: []string{
                    "The list of currently connected peers",
                },
                Repeated: true,
                Type: mklnrpc_Peer(),
            },
        },
    }
}
func mklnrpc_ListUnspentRequest() Type {
    return Type{
        Name: "lnrpc_ListUnspentRequest",
        Fields: []Field{
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations to be included.",
                },
                Type: mkint32(),
            },
            {
                Name: "max_confs",
                Description: []string{
                    "The maximum number of confirmations to be included.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_ListUnspentResponse() Type {
    return Type{
        Name: "lnrpc_ListUnspentResponse",
        Fields: []Field{
            {
                Name: "utxos",
                Description: []string{
                    "A list of utxos",
                },
                Repeated: true,
                Type: mklnrpc_Utxo(),
            },
        },
    }
}
func mklnrpc_LockUnspentRequest() Type {
    return Type{
        Name: "lnrpc_LockUnspentRequest",
        Fields: []Field{
            {
                Name: "unlock",
                Type: mkbool(),
            },
            {
                Name: "transactions",
                Repeated: true,
                Type: mklnrpc_LockUnspentTransaction(),
            },
            {
                Name: "lockname",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_LockUnspentResponse() Type {
    return Type{
        Name: "lnrpc_LockUnspentResponse",
        Fields: []Field{
            {
                Name: "result",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_LockUnspentTransaction() Type {
    return Type{
        Name: "lnrpc_LockUnspentTransaction",
        Fields: []Field{
            {
                Name: "txid",
                Type: mkstring(),
            },
            {
                Name: "vout",
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_MPPRecord() Type {
    return Type{
        Name: "lnrpc_MPPRecord",
        Fields: []Field{
            {
                Name: "payment_addr",
                Description: []string{
                    "A unique, random identifier used to authenticate the sender as the intended",
                    "payer of a multi-path payment. The payment_addr must be the same for all",
                    "subpayments, and match the payment_addr provided in the receiver's invoice.",
                    "The same payment_addr must be used on all subpayments.",
                },
                Type: mkbytes(),
            },
            {
                Name: "total_amt_msat",
                Description: []string{
                    "The total amount in milli-satoshis being sent as part of a larger multi-path",
                    "payment. The caller is responsible for ensuring subpayments to the same node",
                    "and payment_hash sum exactly to total_amt_msat. The same",
                    "total_amt_msat must be used on all subpayments.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_MultiChanBackup() Type {
    return Type{
        Name: "lnrpc_MultiChanBackup",
        Fields: []Field{
            {
                Name: "chan_points",
                Description: []string{
                    "Is the set of all channels that are included in this multi-channel backup.",
                },
                Repeated: true,
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "multi_chan_backup",
                Description: []string{
                    "A single encrypted blob containing all the static channel backups of the",
                    "channel listed above. This can be stored as a single file or blob, and",
                    "safely be replaced with any prior/future versions.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_NetworkInfo() Type {
    return Type{
        Name: "lnrpc_NetworkInfo",
        Fields: []Field{
            {
                Name: "graph_diameter",
                Type: mkuint32(),
            },
            {
                Name: "avg_out_degree",
                Type: mkdouble(),
            },
            {
                Name: "max_out_degree",
                Type: mkuint32(),
            },
            {
                Name: "num_nodes",
                Type: mkuint32(),
            },
            {
                Name: "num_channels",
                Type: mkuint32(),
            },
            {
                Name: "total_network_capacity",
                Type: mkint64(),
            },
            {
                Name: "avg_channel_size",
                Type: mkdouble(),
            },
            {
                Name: "min_channel_size",
                Type: mkint64(),
            },
            {
                Name: "max_channel_size",
                Type: mkint64(),
            },
            {
                Name: "median_channel_size_sat",
                Type: mkint64(),
            },
            {
                Name: "num_zombie_chans",
                Description: []string{
                    "The number of edges marked as zombies.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_NetworkInfoRequest() Type {
    return Type{
        Name: "lnrpc_NetworkInfoRequest",
    }
}
func mklnrpc_NewAddressRequest() Type {
    return Type{
        Name: "lnrpc_NewAddressRequest",
        Fields: []Field{
            {
                Name: "type",
                Description: []string{
                    "The address type",
                },
                Type: mklnrpc_AddressType(),
            },
        },
    }
}
func mklnrpc_NewAddressResponse() Type {
    return Type{
        Name: "lnrpc_NewAddressResponse",
        Fields: []Field{
            {
                Name: "address",
                Description: []string{
                    "The newly generated wallet address",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_NodeAddress() Type {
    return Type{
        Name: "lnrpc_NodeAddress",
        Fields: []Field{
            {
                Name: "network",
                Type: mkstring(),
            },
            {
                Name: "addr",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_NodeInfo() Type {
    return Type{
        Name: "lnrpc_NodeInfo",
        Fields: []Field{
            {
                Name: "node",
                Description: []string{
                    "An individual vertex/node within the channel graph. A node is",
                    "connected to other nodes by one or more channel edges emanating from it. As",
                    "the graph is directed, a node will also have an incoming edge attached to",
                    "it for each outgoing edge.",
                },
                Type: mklnrpc_LightningNode(),
            },
            {
                Name: "num_channels",
                Description: []string{
                    "The total number of channels for the node.",
                },
                Type: mkuint32(),
            },
            {
                Name: "total_capacity",
                Description: []string{
                    "The sum of all channels capacity for the node, denominated in satoshis.",
                },
                Type: mkint64(),
            },
            {
                Name: "channels",
                Description: []string{
                    "A list of all public channels for the node.",
                },
                Repeated: true,
                Type: mklnrpc_ChannelEdge(),
            },
        },
    }
}
func mklnrpc_NodeInfoRequest() Type {
    return Type{
        Name: "lnrpc_NodeInfoRequest",
        Fields: []Field{
            {
                Name: "pub_key",
                Description: []string{
                    "The 33-byte compressed public of the target node",
                },
                Type: mkbytes(),
            },
            {
                Name: "include_channels",
                Description: []string{
                    "If true, will include all known channels associated with the node.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_NodeMetricsRequest() Type {
    return Type{
        Name: "lnrpc_NodeMetricsRequest",
        Fields: []Field{
            {
                Name: "types",
                Description: []string{
                    "The requested node metrics.",
                },
                Repeated: true,
                Type: mklnrpc_NodeMetricType(),
            },
        },
    }
}
func mklnrpc_NodeMetricsResponse() Type {
    return Type{
        Name: "lnrpc_NodeMetricsResponse",
        Fields: []Field{
            {
                Name: "betweenness_centrality",
                Description: []string{
                    "Betweenness centrality is the sum of the ratio of shortest paths that pass",
                    "through the node for each pair of nodes in the graph (not counting paths",
                    "starting or ending at this node).",
                    "Map of node pubkey to betweenness centrality of the node. Normalized",
                    "values are in the [0,1] closed interval.",
                },
                Repeated: true,
                Type: mklnrpc_NodeMetricsResponse_BetweennessCentralityEntry(),
            },
        },
    }
}
func mklnrpc_NodeMetricsResponse_BetweennessCentralityEntry() Type {
    return Type{
        Name: "lnrpc_NodeMetricsResponse_BetweennessCentralityEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mklnrpc_FloatMetric(),
            },
        },
    }
}
func mklnrpc_NodePair() Type {
    return Type{
        Name: "lnrpc_NodePair",
        Fields: []Field{
            {
                Name: "from",
                Description: []string{
                    "The sending node of the pair.",
                },
                Type: mkbytes(),
            },
            {
                Name: "to",
                Description: []string{
                    "The receiving node of the pair.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_NodeUpdate() Type {
    return Type{
        Name: "lnrpc_NodeUpdate",
        Fields: []Field{
            {
                Name: "addresses",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "identity_key",
                Type: mkstring(),
            },
            {
                Name: "global_features",
                Type: mkbytes(),
            },
            {
                Name: "alias",
                Type: mkstring(),
            },
            {
                Name: "color",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Op() Type {
    return Type{
        Name: "lnrpc_Op",
        Fields: []Field{
            {
                Name: "entity",
                Type: mkstring(),
            },
            {
                Name: "actions",
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_OpenChannelRequest() Type {
    return Type{
        Name: "lnrpc_OpenChannelRequest",
        Fields: []Field{
            {
                Name: "node_pubkey",
                Description: []string{
                    "The pubkey of the node to open a channel with.",
                },
                Type: mkbytes(),
            },
            {
                Name: "local_funding_amount",
                Description: []string{
                    "The number of pktoshis the wallet should commit to the channel",
                },
                Type: mkint64(),
            },
            {
                Name: "push_sat",
                Description: []string{
                    "The number of pktoshis to push to the remote side as part of the initial",
                    "commitment state",
                },
                Type: mkint64(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that the funding transaction should be",
                    "confirmed by.",
                },
                Type: mkint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "A manual fee rate set in sat/byte that should be used when crafting the",
                    "funding transaction.",
                },
                Type: mkint64(),
            },
            {
                Name: "private",
                Description: []string{
                    "Whether this channel should be private, not announced to the greater",
                    "network.",
                },
                Type: mkbool(),
            },
            {
                Name: "min_htlc_msat",
                Description: []string{
                    "The minimum value in millisatoshi we will require for incoming HTLCs on",
                    "the channel.",
                },
                Type: mkint64(),
            },
            {
                Name: "remote_csv_delay",
                Description: []string{
                    "The delay we require on the remote's commitment transaction. If this is",
                    "not set, it will be scaled automatically with the channel size.",
                },
                Type: mkuint32(),
            },
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations each one of your outputs used for",
                    "the funding transaction must satisfy.",
                },
                Type: mkint32(),
            },
            {
                Name: "spend_unconfirmed",
                Description: []string{
                    "Whether unconfirmed outputs should be used as inputs for the funding",
                    "transaction.",
                },
                Type: mkbool(),
            },
            {
                Name: "close_address",
                Description: []string{
                    "Close address is an optional address which specifies the address to which",
                    "funds should be paid out to upon cooperative close. This field may only be",
                    "set if the peer supports the option upfront feature bit (call listpeers",
                    "to check). The remote peer will only accept cooperative closes to this",
                    "address if it is set.",
                    "",
                    "Note: If this value is set on channel creation, you will *not* be able to",
                    "cooperatively close out to a different address.",
                },
                Type: mkstring(),
            },
            {
                Name: "funding_shim",
                Description: []string{
                    "Funding shims are an optional argument that allow the caller to intercept",
                    "certain funding functionality. For example, a shim can be provided to use a",
                    "particular key for the commitment key (ideally cold) rather than use one",
                    "that is generated by the wallet as normal, or signal that signing will be",
                    "carried out in an interactive manner (PSBT based).",
                },
                Type: mklnrpc_FundingShim(),
            },
            {
                Name: "remote_max_value_in_flight_msat",
                Description: []string{
                    "The maximum amount of coins in millisatoshi that can be pending within",
                    "the channel. It only applies to the remote party.",
                },
                Type: mkuint64(),
            },
            {
                Name: "remote_max_htlcs",
                Description: []string{
                    "The maximum number of concurrent HTLCs we will allow the remote party to add",
                    "to the commitment transaction.",
                },
                Type: mkuint32(),
            },
            {
                Name: "max_local_csv",
                Description: []string{
                    "Max local csv is the maximum csv delay we will allow for our own commitment",
                    "transaction.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_OpenStatusUpdate() Type {
    return Type{
        Name: "lnrpc_OpenStatusUpdate",
        Fields: []Field{
            {
                Name: "chan_pending",
                Description: []string{
                    "Signals that the channel is now fully negotiated and the funding",
                    "transaction published.",
                },
                Type: mklnrpc_PendingUpdate(),
            },
            {
                Name: "chan_open",
                Description: []string{
                    "Signals that the channel's funding transaction has now reached the",
                    "required number of confirmations on chain and can be used.",
                },
                Type: mklnrpc_ChannelOpenUpdate(),
            },
            {
                Name: "psbt_fund",
                Description: []string{
                    "Signals that the funding process has been suspended and the construction",
                    "of a PSBT that funds the channel PK script is now required.",
                },
                Type: mklnrpc_ReadyForPsbtFunding(),
            },
            {
                Name: "pending_chan_id",
                Description: []string{
                    "The pending channel ID of the created channel. This value may be used to",
                    "further the funding flow manually via the FundingStateStep method.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_OutPoint() Type {
    return Type{
        Name: "lnrpc_OutPoint",
        Fields: []Field{
            {
                Name: "txid_bytes",
                Description: []string{
                    "Raw bytes representing the transaction id.",
                },
                Type: mkbytes(),
            },
            {
                Name: "txid_str",
                Description: []string{
                    "Reversed, hex-encoded string representing the transaction id.",
                },
                Type: mkstring(),
            },
            {
                Name: "output_index",
                Description: []string{
                    "The index of the output on the transaction.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_PayReq() Type {
    return Type{
        Name: "lnrpc_PayReq",
        Fields: []Field{
            {
                Name: "destination",
                Type: mkbytes(),
            },
            {
                Name: "payment_hash",
                Type: mkbytes(),
            },
            {
                Name: "num_satoshis",
                Type: mkint64(),
            },
            {
                Name: "timestamp",
                Type: mkint64(),
            },
            {
                Name: "expiry",
                Type: mkint64(),
            },
            {
                Name: "description",
                Type: mkstring(),
            },
            {
                Name: "description_hash",
                Type: mkbytes(),
            },
            {
                Name: "fallback_addr",
                Type: mkstring(),
            },
            {
                Name: "cltv_expiry",
                Type: mkint64(),
            },
            {
                Name: "route_hints",
                Repeated: true,
                Type: mklnrpc_RouteHint(),
            },
            {
                Name: "payment_addr",
                Type: mkbytes(),
            },
            {
                Name: "num_msat",
                Type: mkint64(),
            },
            {
                Name: "features",
                Repeated: true,
                Type: mklnrpc_PayReq_FeaturesEntry(),
            },
        },
    }
}
func mklnrpc_PayReq_FeaturesEntry() Type {
    return Type{
        Name: "lnrpc_PayReq_FeaturesEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint32(),
            },
            {
                Name: "value",
                Type: mklnrpc_Feature(),
            },
        },
    }
}
func mklnrpc_PayReqString() Type {
    return Type{
        Name: "lnrpc_PayReqString",
        Fields: []Field{
            {
                Name: "pay_req",
                Description: []string{
                    "The payment request string to be decoded",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Payment() Type {
    return Type{
        Name: "lnrpc_Payment",
        Fields: []Field{
            {
                Name: "payment_hash",
                Description: []string{
                    "The payment hash",
                },
                Type: mkstring(),
            },
            {
                Name: "value",
                Description: []string{
                    "Deprecated, use value_sat or value_msat.",
                },
                Type: mkint64(),
            },
            {
                Name: "creation_date",
                Description: []string{
                    "Deprecated, use creation_time_ns",
                },
                Type: mkint64(),
            },
            {
                Name: "fee",
                Description: []string{
                    "Deprecated, use fee_sat or fee_msat.",
                },
                Type: mkint64(),
            },
            {
                Name: "payment_preimage",
                Description: []string{
                    "The payment preimage",
                },
                Type: mkbytes(),
            },
            {
                Name: "value_sat",
                Description: []string{
                    "The value of the payment in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "value_msat",
                Description: []string{
                    "The value of the payment in milli-satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "payment_request",
                Description: []string{
                    "The optional payment request being fulfilled.",
                },
                Type: mkstring(),
            },
            {
                Name: "status",
                Description: []string{
                    "The status of the payment.",
                },
                Type: mklnrpc_Payment_PaymentStatus(),
            },
            {
                Name: "fee_sat",
                Description: []string{
                    "The fee paid for this payment in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_msat",
                Description: []string{
                    "The fee paid for this payment in milli-satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "creation_time_ns",
                Description: []string{
                    "The time in UNIX nanoseconds at which the payment was created.",
                },
                Type: mkint64(),
            },
            {
                Name: "htlcs",
                Description: []string{
                    "The HTLCs made in attempt to settle the payment.",
                },
                Repeated: true,
                Type: mklnrpc_HTLCAttempt(),
            },
            {
                Name: "payment_index",
                Description: []string{
                    "The creation index of this payment. Each payment can be uniquely identified",
                    "by this index, which may not strictly increment by 1 for payments made in",
                    "older versions of lnd.",
                },
                Type: mkuint64(),
            },
            {
                Name: "failure_reason",
                Type: mklnrpc_PaymentFailureReason(),
            },
        },
    }
}
func mklnrpc_PaymentHash() Type {
    return Type{
        Name: "lnrpc_PaymentHash",
        Fields: []Field{
            {
                Name: "r_hash",
                Description: []string{
                    "The payment hash of the invoice to be looked up.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_Peer() Type {
    return Type{
        Name: "lnrpc_Peer",
        Fields: []Field{
            {
                Name: "pub_key",
                Description: []string{
                    "The identity pubkey of the peer",
                },
                Type: mkbytes(),
            },
            {
                Name: "address",
                Description: []string{
                    "Network address of the peer; eg `127.0.0.1:10011`",
                },
                Type: mkstring(),
            },
            {
                Name: "bytes_sent",
                Description: []string{
                    "Bytes of data transmitted to this peer",
                },
                Type: mkuint64(),
            },
            {
                Name: "bytes_recv",
                Description: []string{
                    "Bytes of data transmitted from this peer",
                },
                Type: mkuint64(),
            },
            {
                Name: "sat_sent",
                Description: []string{
                    "Satoshis sent to this peer",
                },
                Type: mkint64(),
            },
            {
                Name: "sat_recv",
                Description: []string{
                    "Satoshis received from this peer",
                },
                Type: mkint64(),
            },
            {
                Name: "inbound",
                Description: []string{
                    "A channel is inbound if the counterparty initiated the channel",
                },
                Type: mkbool(),
            },
            {
                Name: "ping_time",
                Description: []string{
                    "Ping time to this peer",
                },
                Type: mkint64(),
            },
            {
                Name: "sync_type",
                Description: []string{
                    "The type of sync we are currently performing with this peer.",
                },
                Type: mklnrpc_Peer_SyncType(),
            },
            {
                Name: "features",
                Description: []string{
                    "Features advertised by the remote peer in their init message.",
                },
                Repeated: true,
                Type: mklnrpc_Peer_FeaturesEntry(),
            },
            {
                Name: "errors",
                Description: []string{
                    "The latest errors received from our peer with timestamps, limited to the 10",
                    "most recent errors. These errors are tracked across peer connections, but",
                    "are not persisted across lnd restarts. Note that these errors are only",
                    "stored for peers that we have channels open with, to prevent peers from",
                    "spamming us with errors at no cost.",
                },
                Repeated: true,
                Type: mklnrpc_TimestampedError(),
            },
            {
                Name: "flap_count",
                Description: []string{
                    "The number of times we have recorded this peer going offline or coming",
                    "online, recorded across restarts. Note that this value is decreased over",
                    "time if the peer has not recently flapped, so that we can forgive peers",
                    "with historically high flap counts.",
                },
                Type: mkint32(),
            },
            {
                Name: "last_flap_ns",
                Description: []string{
                    "The timestamp of the last flap we observed for this peer. If this value is",
                    "zero, we have not observed any flaps for this peer.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_Peer_FeaturesEntry() Type {
    return Type{
        Name: "lnrpc_Peer_FeaturesEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint32(),
            },
            {
                Name: "value",
                Type: mklnrpc_Feature(),
            },
        },
    }
}
func mklnrpc_PeerEvent() Type {
    return Type{
        Name: "lnrpc_PeerEvent",
        Fields: []Field{
            {
                Name: "pub_key",
                Description: []string{
                    "The identity pubkey of the peer.",
                },
                Type: mkbytes(),
            },
            {
                Name: "type",
                Type: mklnrpc_PeerEvent_EventType(),
            },
        },
    }
}
func mklnrpc_PeerEventSubscription() Type {
    return Type{
        Name: "lnrpc_PeerEventSubscription",
    }
}
func mklnrpc_PendingChannelsRequest() Type {
    return Type{
        Name: "lnrpc_PendingChannelsRequest",
    }
}
func mklnrpc_PendingChannelsResponse() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse",
        Fields: []Field{
            {
                Name: "total_limbo_balance",
                Description: []string{
                    "The balance in satoshis encumbered in pending channels",
                },
                Type: mkint64(),
            },
            {
                Name: "pending_open_channels",
                Description: []string{
                    "Channels pending opening",
                },
                Repeated: true,
                Type: mklnrpc_PendingChannelsResponse_PendingOpenChannel(),
            },
            {
                Name: "pending_closing_channels",
                Description: []string{
                    "Deprecated: Channels pending closing previously contained cooperatively",
                    "closed channels with a single confirmation. These channels are now",
                    "considered closed from the time we see them on chain.",
                },
                Repeated: true,
                Type: mklnrpc_PendingChannelsResponse_ClosedChannel(),
            },
            {
                Name: "pending_force_closing_channels",
                Description: []string{
                    "Channels pending force closing",
                },
                Repeated: true,
                Type: mklnrpc_PendingChannelsResponse_ForceClosedChannel(),
            },
            {
                Name: "waiting_close_channels",
                Description: []string{
                    "Channels waiting for closing tx to confirm",
                },
                Repeated: true,
                Type: mklnrpc_PendingChannelsResponse_WaitingCloseChannel(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_ClosedChannel() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_ClosedChannel",
        Fields: []Field{
            {
                Name: "channel",
                Description: []string{
                    "The pending channel to be closed",
                },
                Type: mklnrpc_PendingChannelsResponse_PendingChannel(),
            },
            {
                Name: "closing_txid",
                Description: []string{
                    "The transaction id of the closing transaction",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_Commitments() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_Commitments",
        Fields: []Field{
            {
                Name: "local_txid",
                Description: []string{
                    "Hash of the local version of the commitment tx.",
                },
                Type: mkstring(),
            },
            {
                Name: "remote_txid",
                Description: []string{
                    "Hash of the remote version of the commitment tx.",
                },
                Type: mkstring(),
            },
            {
                Name: "remote_pending_txid",
                Description: []string{
                    "Hash of the remote pending version of the commitment tx.",
                },
                Type: mkstring(),
            },
            {
                Name: "local_commit_fee_sat",
                Description: []string{
                    "The amount in satoshis calculated to be paid in fees for the local",
                    "commitment.",
                },
                Type: mkuint64(),
            },
            {
                Name: "remote_commit_fee_sat",
                Description: []string{
                    "The amount in satoshis calculated to be paid in fees for the remote",
                    "commitment.",
                },
                Type: mkuint64(),
            },
            {
                Name: "remote_pending_commit_fee_sat",
                Description: []string{
                    "The amount in satoshis calculated to be paid in fees for the remote",
                    "pending commitment.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_ForceClosedChannel() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_ForceClosedChannel",
        Fields: []Field{
            {
                Name: "channel",
                Description: []string{
                    "The pending channel to be force closed",
                },
                Type: mklnrpc_PendingChannelsResponse_PendingChannel(),
            },
            {
                Name: "closing_txid",
                Description: []string{
                    "The transaction id of the closing transaction",
                },
                Type: mkstring(),
            },
            {
                Name: "limbo_balance",
                Description: []string{
                    "The balance in satoshis encumbered in this pending channel",
                },
                Type: mkint64(),
            },
            {
                Name: "maturity_height",
                Description: []string{
                    "The height at which funds can be swept into the wallet",
                },
                Type: mkuint32(),
            },
            {
                Name: "blocks_til_maturity",
                Description: []string{
                    "Remaining # of blocks until the commitment output can be swept.",
                    "Negative values indicate how many blocks have passed since becoming",
                    "mature.",
                },
                Type: mkint32(),
            },
            {
                Name: "recovered_balance",
                Description: []string{
                    "The total value of funds successfully recovered from this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "pending_htlcs",
                Repeated: true,
                Type: mklnrpc_PendingHTLC(),
            },
            {
                Name: "anchor",
                Type: mklnrpc_PendingChannelsResponse_ForceClosedChannel_AnchorState(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_PendingChannel() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_PendingChannel",
        Fields: []Field{
            {
                Name: "remote_node_pub",
                Type: mkbytes(),
            },
            {
                Name: "channel_point",
                Type: mkstring(),
            },
            {
                Name: "capacity",
                Type: mkint64(),
            },
            {
                Name: "local_balance",
                Type: mkint64(),
            },
            {
                Name: "remote_balance",
                Type: mkint64(),
            },
            {
                Name: "local_chan_reserve_sat",
                Description: []string{
                    "The minimum satoshis this node is required to reserve in its",
                    "balance.",
                },
                Type: mkint64(),
            },
            {
                Name: "remote_chan_reserve_sat",
                Description: []string{
                    "The minimum satoshis the other node is required to reserve in its",
                    "balance.",
                },
                Type: mkint64(),
            },
            {
                Name: "initiator",
                Description: []string{
                    "The party that initiated opening the channel.",
                },
                Type: mklnrpc_Initiator(),
            },
            {
                Name: "commitment_type",
                Description: []string{
                    "The commitment type used by this channel.",
                },
                Type: mklnrpc_CommitmentType(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_PendingOpenChannel() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_PendingOpenChannel",
        Fields: []Field{
            {
                Name: "channel",
                Description: []string{
                    "The pending channel",
                },
                Type: mklnrpc_PendingChannelsResponse_PendingChannel(),
            },
            {
                Name: "confirmation_height",
                Description: []string{
                    "The height at which this channel will be confirmed",
                },
                Type: mkuint32(),
            },
            {
                Name: "commit_fee",
                Description: []string{
                    "The amount calculated to be paid in fees for the current set of",
                    "commitment transactions. The fee amount is persisted with the channel",
                    "in order to allow the fee amount to be removed and recalculated with",
                    "each channel state update, including updates that happen after a system",
                    "restart.",
                },
                Type: mkint64(),
            },
            {
                Name: "commit_weight",
                Description: []string{
                    "The weight of the commitment transaction",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_per_kw",
                Description: []string{
                    "The required number of satoshis per kilo-weight that the requester will",
                    "pay at all times, for both the funding transaction and commitment",
                    "transaction. This value can later be updated once the channel is open.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_PendingChannelsResponse_WaitingCloseChannel() Type {
    return Type{
        Name: "lnrpc_PendingChannelsResponse_WaitingCloseChannel",
        Fields: []Field{
            {
                Name: "channel",
                Description: []string{
                    "The pending channel waiting for closing tx to confirm",
                },
                Type: mklnrpc_PendingChannelsResponse_PendingChannel(),
            },
            {
                Name: "limbo_balance",
                Description: []string{
                    "The balance in satoshis encumbered in this channel",
                },
                Type: mkint64(),
            },
            {
                Name: "commitments",
                Description: []string{
                    "A list of valid commitment transactions. Any of these can confirm at",
                    "this point.",
                },
                Type: mklnrpc_PendingChannelsResponse_Commitments(),
            },
        },
    }
}
func mklnrpc_PendingHTLC() Type {
    return Type{
        Name: "lnrpc_PendingHTLC",
        Fields: []Field{
            {
                Name: "incoming",
                Description: []string{
                    "The direction within the channel that the htlc was sent",
                },
                Type: mkbool(),
            },
            {
                Name: "amount",
                Description: []string{
                    "The total value of the htlc",
                },
                Type: mkint64(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The final output to be swept back to the user's wallet",
                },
                Type: mkstring(),
            },
            {
                Name: "maturity_height",
                Description: []string{
                    "The next block height at which we can spend the current stage",
                },
                Type: mkuint32(),
            },
            {
                Name: "blocks_til_maturity",
                Description: []string{
                    "The number of blocks remaining until the current stage can be swept.",
                    "Negative values indicate how many blocks have passed since becoming",
                    "mature.",
                },
                Type: mkint32(),
            },
            {
                Name: "stage",
                Description: []string{
                    "Indicates whether the htlc is in its first or second stage of recovery",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_PendingUpdate() Type {
    return Type{
        Name: "lnrpc_PendingUpdate",
        Fields: []Field{
            {
                Name: "txid",
                Type: mkbytes(),
            },
            {
                Name: "output_index",
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_PolicyUpdateRequest() Type {
    return Type{
        Name: "lnrpc_PolicyUpdateRequest",
        Fields: []Field{
            {
                Name: "global",
                Description: []string{
                    "If set, then this update applies to all currently active channels.",
                },
                Type: mkbool(),
            },
            {
                Name: "chan_point",
                Description: []string{
                    "If set, this update will target a specific channel.",
                },
                Type: mklnrpc_ChannelPoint(),
            },
            {
                Name: "base_fee_msat",
                Description: []string{
                    "The base fee charged regardless of the number of milli-satoshis sent.",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_rate",
                Description: []string{
                    "The effective fee rate in milli-satoshis. The precision of this value",
                    "goes up to 6 decimal places, so 1e-6.",
                },
                Type: mkdouble(),
            },
            {
                Name: "time_lock_delta",
                Description: []string{
                    "The required timelock delta for HTLCs forwarded over the channel.",
                },
                Type: mkuint32(),
            },
            {
                Name: "max_htlc_msat",
                Description: []string{
                    "If set, the maximum HTLC size in milli-satoshis. If unset, the maximum",
                    "HTLC will be unchanged.",
                },
                Type: mkuint64(),
            },
            {
                Name: "min_htlc_msat",
                Description: []string{
                    "The minimum HTLC size in milli-satoshis. Only applied if",
                    "min_htlc_msat_specified is true.",
                },
                Type: mkuint64(),
            },
            {
                Name: "min_htlc_msat_specified",
                Description: []string{
                    "If true, min_htlc_msat is applied.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_PolicyUpdateResponse() Type {
    return Type{
        Name: "lnrpc_PolicyUpdateResponse",
    }
}
func mklnrpc_PrevOut() Type {
    return Type{
        Name: "lnrpc_PrevOut",
        Fields: []Field{
            {
                Name: "address",
                Type: mkstring(),
            },
            {
                Name: "value_coins",
                Type: mkdouble(),
            },
            {
                Name: "svalue",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_PsbtShim() Type {
    return Type{
        Name: "lnrpc_PsbtShim",
        Fields: []Field{
            {
                Name: "pending_chan_id",
                Description: []string{
                    "A unique identifier of 32 random bytes that will be used as the pending",
                    "channel ID to identify the PSBT state machine when interacting with it and",
                    "on the wire protocol to initiate the funding request.",
                },
                Type: mkbytes(),
            },
            {
                Name: "base_psbt",
                Description: []string{
                    "An optional base PSBT the new channel output will be added to. If this is",
                    "non-empty, it must be a binary serialized PSBT.",
                },
                Type: mkbytes(),
            },
            {
                Name: "no_publish",
                Description: []string{
                    "If a channel should be part of a batch (multiple channel openings in one",
                    "transaction), it can be dangerous if the whole batch transaction is",
                    "published too early before all channel opening negotiations are completed.",
                    "This flag prevents this particular channel from broadcasting the transaction",
                    "after the negotiation with the remote peer. In a batch of channel openings",
                    "this flag should be set to true for every channel but the very last.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_QueryRoutesRequest() Type {
    return Type{
        Name: "lnrpc_QueryRoutesRequest",
        Fields: []Field{
            {
                Name: "pub_key",
                Description: []string{
                    "The 33-byte public key for the payment destination",
                },
                Type: mkbytes(),
            },
            {
                Name: "amt",
                Description: []string{
                    "The amount to send expressed in satoshis.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "amt_msat",
                Description: []string{
                    "The amount to send expressed in millisatoshis.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "final_cltv_delta",
                Description: []string{
                    "An optional CLTV delta from the current height that should be used for the",
                    "timelock of the final hop. Note that unlike SendPayment, QueryRoutes does",
                    "not add any additional block padding on top of final_ctlv_delta. This",
                    "padding of a few blocks needs to be added manually or otherwise failures may",
                    "happen when a block comes in while the payment is in flight.",
                },
                Type: mkint32(),
            },
            {
                Name: "fee_limit",
                Description: []string{
                    "The maximum number of satoshis that will be paid as a fee of the payment.",
                    "This value can be represented either as a percentage of the amount being",
                    "sent, or as a fixed amount of the maximum fee the user is willing the pay to",
                    "send the payment.",
                },
                Type: mklnrpc_FeeLimit(),
            },
            {
                Name: "ignored_nodes",
                Description: []string{
                    "A list of nodes to ignore during path finding.",
                },
                Repeated: true,
                Type: mkbytes(),
            },
            {
                Name: "ignored_edges",
                Description: []string{
                    "Deprecated. A list of edges to ignore during path finding.",
                },
                Repeated: true,
                Type: mklnrpc_EdgeLocator(),
            },
            {
                Name: "source_pub_key",
                Description: []string{
                    "The source node where the request route should originated from. If empty,",
                    "self is assumed.",
                },
                Type: mkbytes(),
            },
            {
                Name: "use_mission_control",
                Description: []string{
                    "If set to true, edge probabilities from mission control will be used to get",
                    "the optimal route.",
                },
                Type: mkbool(),
            },
            {
                Name: "ignored_pairs",
                Description: []string{
                    "A list of directed node pairs that will be ignored during path finding.",
                },
                Repeated: true,
                Type: mklnrpc_NodePair(),
            },
            {
                Name: "cltv_limit",
                Description: []string{
                    "An optional maximum total time lock for the route. If the source is empty or",
                    "ourselves, this should not exceed lnd's `--max-cltv-expiry` setting. If",
                    "zero, then the value of `--max-cltv-expiry` is used as the limit.",
                },
                Type: mkuint32(),
            },
            {
                Name: "dest_custom_records",
                Description: []string{
                    "An optional field that can be used to pass an arbitrary set of TLV records",
                    "to a peer which understands the new records. This can be used to pass",
                    "application specific data during the payment attempt. If the destination",
                    "does not support the specified recrods, and error will be returned.",
                    "Record types are required to be in the custom range >= 65536.",
                },
                Repeated: true,
                Type: mklnrpc_QueryRoutesRequest_DestCustomRecordsEntry(),
            },
            {
                Name: "outgoing_chan_id",
                Description: []string{
                    "The channel id of the channel that must be taken to the first hop. If zero,",
                    "any channel may be used.",
                },
                Type: mkuint64(),
            },
            {
                Name: "last_hop_pubkey",
                Description: []string{
                    "The pubkey of the last hop of the route. If empty, any hop may be used.",
                },
                Type: mkbytes(),
            },
            {
                Name: "route_hints",
                Description: []string{
                    "Optional route hints to reach the destination through private channels.",
                },
                Repeated: true,
                Type: mklnrpc_RouteHint(),
            },
            {
                Name: "dest_features",
                Description: []string{
                    "Features assumed to be supported by the final node. All transitive feature",
                    "dependencies must also be set properly. For a given feature bit pair, either",
                    "optional or remote may be set, but not both. If this field is nil or empty,",
                    "the router will try to load destination features from the graph as a",
                    "fallback.",
                },
                Repeated: true,
                Type: mklnrpc_FeatureBit(),
            },
        },
    }
}
func mklnrpc_QueryRoutesRequest_DestCustomRecordsEntry() Type {
    return Type{
        Name: "lnrpc_QueryRoutesRequest_DestCustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_QueryRoutesResponse() Type {
    return Type{
        Name: "lnrpc_QueryRoutesResponse",
        Fields: []Field{
            {
                Name: "routes",
                Description: []string{
                    "The route that results from the path finding operation. This is still a",
                    "repeated field to retain backwards compatibility.",
                },
                Repeated: true,
                Type: mklnrpc_Route(),
            },
            {
                Name: "success_prob",
                Description: []string{
                    "The success probability of the returned route based on the current mission",
                    "control state. [EXPERIMENTAL]",
                },
                Type: mkdouble(),
            },
        },
    }
}
func mklnrpc_ReSyncChainRequest() Type {
    return Type{
        Name: "lnrpc_ReSyncChainRequest",
        Fields: []Field{
            {
                Name: "from_height",
                Type: mkint32(),
            },
            {
                Name: "to_height",
                Type: mkint32(),
            },
            {
                Name: "addresses",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "drop_db",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_ReSyncChainResponse() Type {
    return Type{
        Name: "lnrpc_ReSyncChainResponse",
    }
}
func mklnrpc_ReadyForPsbtFunding() Type {
    return Type{
        Name: "lnrpc_ReadyForPsbtFunding",
        Fields: []Field{
            {
                Name: "funding_address",
                Description: []string{
                    "The P2WSH address of the channel funding multisig address that the below",
                    "specified amount in satoshis needs to be sent to.",
                },
                Type: mkstring(),
            },
            {
                Name: "funding_amount",
                Description: []string{
                    "The exact amount in satoshis that needs to be sent to the above address to",
                    "fund the pending channel.",
                },
                Type: mkint64(),
            },
            {
                Name: "psbt",
                Description: []string{
                    "A raw PSBT that contains the pending channel output. If a base PSBT was",
                    "provided in the PsbtShim, this is the base PSBT with one additional output.",
                    "If no base PSBT was specified, this is an otherwise empty PSBT with exactly",
                    "one output.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_Resolution() Type {
    return Type{
        Name: "lnrpc_Resolution",
        Fields: []Field{
            {
                Name: "resolution_type",
                Description: []string{
                    "The type of output we are resolving.",
                },
                Type: mklnrpc_ResolutionType(),
            },
            {
                Name: "outcome",
                Description: []string{
                    "The outcome of our on chain action that resolved the outpoint.",
                },
                Type: mklnrpc_ResolutionOutcome(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The outpoint that was spent by the resolution.",
                },
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "amount_sat",
                Description: []string{
                    "The amount that was claimed by the resolution.",
                },
                Type: mkuint64(),
            },
            {
                Name: "sweep_txid",
                Description: []string{
                    "The hex-encoded transaction ID of the sweep transaction that spent the",
                    "output.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_RestError() Type {
    return Type{
        Name: "lnrpc_RestError",
        Fields: []Field{
            {
                Name: "message",
                Type: mkstring(),
            },
            {
                Name: "stack",
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_RestoreBackupResponse() Type {
    return Type{
        Name: "lnrpc_RestoreBackupResponse",
    }
}
func mklnrpc_RestoreChanBackupRequest() Type {
    return Type{
        Name: "lnrpc_RestoreChanBackupRequest",
        Fields: []Field{
            {
                Name: "chan_backups",
                Description: []string{
                    "The channels to restore as a list of channel/backup pairs.",
                },
                Type: mklnrpc_ChannelBackups(),
            },
            {
                Name: "multi_chan_backup",
                Description: []string{
                    "The channels to restore in the packed multi backup format.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_Route() Type {
    return Type{
        Name: "lnrpc_Route",
        Description: []string{
            "A path through the channel graph which runs over one or more channels in",
            "succession. This struct carries all the information required to craft the",
            "Sphinx onion packet, and send the payment along the first hop in the path. A",
            "route is only selected as valid if all the channels have sufficient capacity to",
            "carry the initial payment amount after fees are accounted for.",
        },
        Fields: []Field{
            {
                Name: "total_time_lock",
                Description: []string{
                    "The cumulative (final) time lock across the entire route. This is the CLTV",
                    "value that should be extended to the first hop in the route. All other hops",
                    "will decrement the time-lock as advertised, leaving enough time for all",
                    "hops to wait for or present the payment preimage to complete the payment.",
                },
                Type: mkuint32(),
            },
            {
                Name: "total_fees",
                Description: []string{
                    "The sum of the fees paid at each hop within the final route. In the case",
                    "of a one-hop payment, this value will be zero as we don't need to pay a fee",
                    "to ourselves.",
                },
                Type: mkint64(),
            },
            {
                Name: "total_amt",
                Description: []string{
                    "The total amount of funds required to complete a payment over this route.",
                    "This value includes the cumulative fees at each hop. As a result, the HTLC",
                    "extended to the first-hop in the route will need to have at least this many",
                    "satoshis, otherwise the route will fail at an intermediate node due to an",
                    "insufficient amount of fees.",
                },
                Type: mkint64(),
            },
            {
                Name: "hops",
                Description: []string{
                    "Contains details concerning the specific forwarding details at each hop.",
                },
                Repeated: true,
                Type: mklnrpc_Hop(),
            },
            {
                Name: "total_fees_msat",
                Description: []string{
                    "The total fees in millisatoshis.",
                },
                Type: mkint64(),
            },
            {
                Name: "total_amt_msat",
                Description: []string{
                    "The total amount in millisatoshis.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_RouteHint() Type {
    return Type{
        Name: "lnrpc_RouteHint",
        Fields: []Field{
            {
                Name: "hop_hints",
                Description: []string{
                    "A list of hop hints that when chained together can assist in reaching a",
                    "specific destination.",
                },
                Repeated: true,
                Type: mklnrpc_HopHint(),
            },
        },
    }
}
func mklnrpc_RoutingPolicy() Type {
    return Type{
        Name: "lnrpc_RoutingPolicy",
        Fields: []Field{
            {
                Name: "time_lock_delta",
                Type: mkuint32(),
            },
            {
                Name: "min_htlc",
                Type: mkint64(),
            },
            {
                Name: "fee_base_msat",
                Type: mkint64(),
            },
            {
                Name: "fee_rate_milli_msat",
                Type: mkint64(),
            },
            {
                Name: "disabled",
                Type: mkbool(),
            },
            {
                Name: "max_htlc_msat",
                Type: mkuint64(),
            },
            {
                Name: "last_update",
                Type: mkuint32(),
            },
        },
    }
}
func mklnrpc_ScriptSig() Type {
    return Type{
        Name: "lnrpc_ScriptSig",
        Fields: []Field{
            {
                Name: "asm",
                Type: mkstring(),
            },
            {
                Name: "hex",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SendCoinsRequest() Type {
    return Type{
        Name: "lnrpc_SendCoinsRequest",
        Fields: []Field{
            {
                Name: "addr",
                Description: []string{
                    "The address to send coins to",
                },
                Type: mkstring(),
            },
            {
                Name: "amount",
                Description: []string{
                    "The amount in satoshis to send",
                },
                Type: mkint64(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that this transaction should be confirmed",
                    "by.",
                },
                Type: mkint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "A manual fee rate set in sat/byte that should be used when crafting the",
                    "transaction.",
                },
                Type: mkint64(),
            },
            {
                Name: "send_all",
                Description: []string{
                    "If set, then the amount field will be ignored, and lnd will attempt to",
                    "send all the coins under control of the internal wallet to the specified",
                    "address.",
                },
                Type: mkbool(),
            },
            {
                Name: "label",
                Description: []string{
                    "An optional label for the transaction, limited to 500 characters.",
                },
                Type: mkstring(),
            },
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations each one of your outputs used for",
                    "the transaction must satisfy.",
                },
                Type: mkint32(),
            },
            {
                Name: "spend_unconfirmed",
                Description: []string{
                    "Whether unconfirmed outputs should be used as inputs for the transaction.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_SendCoinsResponse() Type {
    return Type{
        Name: "lnrpc_SendCoinsResponse",
        Fields: []Field{
            {
                Name: "txid",
                Description: []string{
                    "The transaction ID of the transaction",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SendFromRequest() Type {
    return Type{
        Name: "lnrpc_SendFromRequest",
        Fields: []Field{
            {
                Name: "to_address",
                Type: mkstring(),
            },
            {
                Name: "amount",
                Type: mkdouble(),
            },
            {
                Name: "from_address",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "min_conf",
                Type: mkint32(),
            },
            {
                Name: "max_inputs",
                Type: mkint32(),
            },
            {
                Name: "min_height",
                Type: mkint32(),
            },
        },
    }
}
func mklnrpc_SendFromResponse() Type {
    return Type{
        Name: "lnrpc_SendFromResponse",
        Fields: []Field{
            {
                Name: "tx_hash",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SendManyRequest() Type {
    return Type{
        Name: "lnrpc_SendManyRequest",
        Fields: []Field{
            {
                Name: "AddrToAmount",
                Description: []string{
                    "The map from addresses to amounts",
                },
                Repeated: true,
                Type: mklnrpc_SendManyRequest_AddrToAmountEntry(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that this transaction should be confirmed",
                    "by.",
                },
                Type: mkint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "A manual fee rate set in sat/byte that should be used when crafting the",
                    "transaction.",
                },
                Type: mkint64(),
            },
            {
                Name: "label",
                Description: []string{
                    "An optional label for the transaction, limited to 500 characters.",
                },
                Type: mkstring(),
            },
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations each one of your outputs used for",
                    "the transaction must satisfy.",
                },
                Type: mkint32(),
            },
            {
                Name: "spend_unconfirmed",
                Description: []string{
                    "Whether unconfirmed outputs should be used as inputs for the transaction.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_SendManyRequest_AddrToAmountEntry() Type {
    return Type{
        Name: "lnrpc_SendManyRequest_AddrToAmountEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_SendManyResponse() Type {
    return Type{
        Name: "lnrpc_SendManyResponse",
        Fields: []Field{
            {
                Name: "txid",
                Description: []string{
                    "The id of the transaction",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SendRequest() Type {
    return Type{
        Name: "lnrpc_SendRequest",
        Fields: []Field{
            {
                Name: "dest",
                Description: []string{
                    "The identity pubkey of the payment recipient.",
                },
                Type: mkbytes(),
            },
            {
                Name: "amt",
                Description: []string{
                    "The amount to send expressed in satoshis.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "amt_msat",
                Description: []string{
                    "The amount to send expressed in millisatoshis.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "payment_hash",
                Description: []string{
                    "The hash to use within the payment's HTLC.",
                },
                Type: mkbytes(),
            },
            {
                Name: "payment_request",
                Description: []string{
                    "A bare-bones invoice for a payment within the Lightning Network. With the",
                    "details of the invoice, the sender has all the data necessary to send a",
                    "payment to the recipient.",
                },
                Type: mkstring(),
            },
            {
                Name: "final_cltv_delta",
                Description: []string{
                    "The CLTV delta from the current height that should be used to set the",
                    "timelock for the final hop.",
                },
                Type: mkint32(),
            },
            {
                Name: "fee_limit",
                Description: []string{
                    "The maximum number of satoshis that will be paid as a fee of the payment.",
                    "This value can be represented either as a percentage of the amount being",
                    "sent, or as a fixed amount of the maximum fee the user is willing the pay to",
                    "send the payment.",
                },
                Type: mklnrpc_FeeLimit(),
            },
            {
                Name: "outgoing_chan_id",
                Description: []string{
                    "The channel id of the channel that must be taken to the first hop. If zero,",
                    "any channel may be used.",
                },
                Type: mkuint64(),
            },
            {
                Name: "last_hop_pubkey",
                Description: []string{
                    "The pubkey of the last hop of the route. If empty, any hop may be used.",
                },
                Type: mkbytes(),
            },
            {
                Name: "cltv_limit",
                Description: []string{
                    "An optional maximum total time lock for the route. This should not exceed",
                    "lnd's `--max-cltv-expiry` setting. If zero, then the value of",
                    "`--max-cltv-expiry` is enforced.",
                },
                Type: mkuint32(),
            },
            {
                Name: "dest_custom_records",
                Description: []string{
                    "An optional field that can be used to pass an arbitrary set of TLV records",
                    "to a peer which understands the new records. This can be used to pass",
                    "application specific data during the payment attempt. Record types are",
                    "required to be in the custom range >= 65536.",
                },
                Repeated: true,
                Type: mklnrpc_SendRequest_DestCustomRecordsEntry(),
            },
            {
                Name: "allow_self_payment",
                Description: []string{
                    "If set, circular payments to self are permitted.",
                },
                Type: mkbool(),
            },
            {
                Name: "dest_features",
                Description: []string{
                    "Features assumed to be supported by the final node. All transitive feature",
                    "dependencies must also be set properly. For a given feature bit pair, either",
                    "optional or remote may be set, but not both. If this field is nil or empty,",
                    "the router will try to load destination features from the graph as a",
                    "fallback.",
                },
                Repeated: true,
                Type: mklnrpc_FeatureBit(),
            },
        },
    }
}
func mklnrpc_SendRequest_DestCustomRecordsEntry() Type {
    return Type{
        Name: "lnrpc_SendRequest_DestCustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_SendResponse() Type {
    return Type{
        Name: "lnrpc_SendResponse",
        Fields: []Field{
            {
                Name: "payment_error",
                Type: mkstring(),
            },
            {
                Name: "payment_preimage",
                Type: mkbytes(),
            },
            {
                Name: "payment_route",
                Type: mklnrpc_Route(),
            },
            {
                Name: "payment_hash",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_SendToRouteRequest() Type {
    return Type{
        Name: "lnrpc_SendToRouteRequest",
        Fields: []Field{
            {
                Name: "payment_hash",
                Description: []string{
                    "The payment hash to use for the HTLC.",
                },
                Type: mkbytes(),
            },
            {
                Name: "route",
                Description: []string{
                    "Route that should be used to attempt to complete the payment.",
                },
                Type: mklnrpc_Route(),
            },
        },
    }
}
func mklnrpc_SetNetworkStewardVoteRequest() Type {
    return Type{
        Name: "lnrpc_SetNetworkStewardVoteRequest",
        Fields: []Field{
            {
                Name: "vote_against",
                Type: mkstring(),
            },
            {
                Name: "vote_for",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SetNetworkStewardVoteResponse() Type {
    return Type{
        Name: "lnrpc_SetNetworkStewardVoteResponse",
    }
}
func mklnrpc_SignMessageRequest() Type {
    return Type{
        Name: "lnrpc_SignMessageRequest",
        Fields: []Field{
            {
                Name: "msg",
                Description: []string{
                    "The message to be signed.",
                },
                Type: mkstring(),
            },
            {
                Name: "msg_bin",
                Description: []string{
                    "If specified then will override msg, binary form.",
                },
                Type: mkbytes(),
            },
            {
                Name: "address",
                Description: []string{
                    "Address to select for signing with.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_SignMessageResponse() Type {
    return Type{
        Name: "lnrpc_SignMessageResponse",
        Fields: []Field{
            {
                Name: "signature",
                Description: []string{
                    "The signature for the given message",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_StopReSyncRequest() Type {
    return Type{
        Name: "lnrpc_StopReSyncRequest",
    }
}
func mklnrpc_StopReSyncResponse() Type {
    return Type{
        Name: "lnrpc_StopReSyncResponse",
        Fields: []Field{
            {
                Name: "value",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_StopRequest() Type {
    return Type{
        Name: "lnrpc_StopRequest",
    }
}
func mklnrpc_StopResponse() Type {
    return Type{
        Name: "lnrpc_StopResponse",
    }
}
func mklnrpc_TimestampedError() Type {
    return Type{
        Name: "lnrpc_TimestampedError",
        Fields: []Field{
            {
                Name: "timestamp",
                Description: []string{
                    "The unix timestamp in seconds when the error occurred.",
                },
                Type: mkuint64(),
            },
            {
                Name: "error",
                Description: []string{
                    "The string representation of the error sent by our peer.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Transaction() Type {
    return Type{
        Name: "lnrpc_Transaction",
        Fields: []Field{
            {
                Name: "tx_hash",
                Description: []string{
                    "The transaction hash",
                },
                Type: mkstring(),
            },
            {
                Name: "amount",
                Description: []string{
                    "The transaction amount, denominated in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "num_confirmations",
                Description: []string{
                    "The number of confirmations",
                },
                Type: mkint32(),
            },
            {
                Name: "block_hash",
                Description: []string{
                    "The hash of the block this transaction was included in",
                },
                Type: mkstring(),
            },
            {
                Name: "block_height",
                Description: []string{
                    "The height of the block this transaction was included in",
                },
                Type: mkint32(),
            },
            {
                Name: "time_stamp",
                Description: []string{
                    "Timestamp of this transaction",
                },
                Type: mkint64(),
            },
            {
                Name: "total_fees",
                Description: []string{
                    "Fees paid for this transaction",
                },
                Type: mkint64(),
            },
            {
                Name: "dest_addresses",
                Description: []string{
                    "Addresses that received funds for this transaction",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "raw_tx_hex",
                Description: []string{
                    "The raw transaction hex.",
                },
                Type: mkbytes(),
            },
            {
                Name: "label",
                Description: []string{
                    "A label that was optionally set on transaction broadcast.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_TransactionDetails() Type {
    return Type{
        Name: "lnrpc_TransactionDetails",
        Fields: []Field{
            {
                Name: "transactions",
                Description: []string{
                    "The list of transactions relevant to the wallet.",
                },
                Repeated: true,
                Type: mklnrpc_Transaction(),
            },
        },
    }
}
func mklnrpc_TransactionResult() Type {
    return Type{
        Name: "lnrpc_TransactionResult",
        Fields: []Field{
            {
                Name: "amount",
                Type: mkdouble(),
            },
            {
                Name: "amount_units",
                Type: mkuint64(),
            },
            {
                Name: "fee",
                Type: mkdouble(),
            },
            {
                Name: "fee_units",
                Type: mkuint64(),
            },
            {
                Name: "confirmations",
                Type: mkint64(),
            },
            {
                Name: "block_hash",
                Type: mkstring(),
            },
            {
                Name: "block_index",
                Type: mkint64(),
            },
            {
                Name: "block_time",
                Type: mkint64(),
            },
            {
                Name: "txid",
                Type: mkstring(),
            },
            {
                Name: "wallet_conflicts",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "time",
                Type: mkint64(),
            },
            {
                Name: "time_received",
                Type: mkint64(),
            },
            {
                Name: "details",
                Repeated: true,
                Type: mklnrpc_GetTransactionDetailsResult(),
            },
            {
                Name: "raw",
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_Utxo() Type {
    return Type{
        Name: "lnrpc_Utxo",
        Fields: []Field{
            {
                Name: "address_type",
                Description: []string{
                    "The type of address",
                },
                Type: mklnrpc_AddressType(),
            },
            {
                Name: "address",
                Description: []string{
                    "The address",
                },
                Type: mkstring(),
            },
            {
                Name: "amount_sat",
                Description: []string{
                    "The value of the unspent coin in satoshis",
                },
                Type: mkint64(),
            },
            {
                Name: "pk_script",
                Description: []string{
                    "The pkscript in hex",
                },
                Type: mkstring(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The outpoint in format txid:n",
                },
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "confirmations",
                Description: []string{
                    "The number of confirmations for the Utxo",
                },
                Type: mkint64(),
            },
        },
    }
}
func mklnrpc_VerifyChanBackupResponse() Type {
    return Type{
        Name: "lnrpc_VerifyChanBackupResponse",
    }
}
func mklnrpc_VinPrevOut() Type {
    return Type{
        Name: "lnrpc_VinPrevOut",
        Fields: []Field{
            {
                Name: "coinbase",
                Type: mkstring(),
            },
            {
                Name: "txid",
                Type: mkstring(),
            },
            {
                Name: "vout",
                Type: mkuint32(),
            },
            {
                Name: "script_sig",
                Type: mklnrpc_ScriptSig(),
            },
            {
                Name: "sequence",
                Type: mkuint32(),
            },
            {
                Name: "witness",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "prev_out",
                Type: mklnrpc_PrevOut(),
            },
        },
    }
}
func mklnrpc_Vote() Type {
    return Type{
        Name: "lnrpc_Vote",
        Fields: []Field{
            {
                Name: "for",
                Type: mkstring(),
            },
            {
                Name: "against",
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_Vout() Type {
    return Type{
        Name: "lnrpc_Vout",
        Fields: []Field{
            {
                Name: "value_coins",
                Type: mkdouble(),
            },
            {
                Name: "svalue",
                Type: mkstring(),
            },
            {
                Name: "n",
                Type: mkuint32(),
            },
            {
                Name: "address",
                Type: mkstring(),
            },
            {
                Name: "vote",
                Type: mklnrpc_Vote(),
            },
        },
    }
}
func mklnrpc_WalletBalanceRequest() Type {
    return Type{
        Name: "lnrpc_WalletBalanceRequest",
    }
}
func mklnrpc_WalletBalanceResponse() Type {
    return Type{
        Name: "lnrpc_WalletBalanceResponse",
        Fields: []Field{
            {
                Name: "total_balance",
                Description: []string{
                    "The balance of the wallet",
                },
                Type: mkint64(),
            },
            {
                Name: "confirmed_balance",
                Description: []string{
                    "The confirmed balance of a wallet(with >= 1 confirmations)",
                },
                Type: mkint64(),
            },
            {
                Name: "unconfirmed_balance",
                Description: []string{
                    "The unconfirmed balance of a wallet(with 0 confirmations)",
                },
                Type: mkint64(),
            },
        },
    }
}
func mksignrpc_InputScript() Type {
    return Type{
        Name: "signrpc_InputScript",
        Fields: []Field{
            {
                Name: "witness",
                Description: []string{
                    "The serializes witness stack for the specified input.",
                },
                Repeated: true,
                Type: mkbytes(),
            },
            {
                Name: "sig_script",
                Description: []string{
                    "The optional sig script for the specified witness that will only be set if",
                    "the input specified is a nested p2sh witness program.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_InputScriptResp() Type {
    return Type{
        Name: "signrpc_InputScriptResp",
        Fields: []Field{
            {
                Name: "input_scripts",
                Description: []string{
                    "The set of fully valid input scripts requested.",
                },
                Repeated: true,
                Type: mksignrpc_InputScript(),
            },
        },
    }
}
func mksignrpc_KeyDescriptor() Type {
    return Type{
        Name: "signrpc_KeyDescriptor",
        Fields: []Field{
            {
                Name: "raw_key_bytes",
                Description: []string{
                    "The raw bytes of the key being identified. Either this or the KeyLocator",
                    "must be specified.",
                },
                Type: mkbytes(),
            },
            {
                Name: "key_loc",
                Description: []string{
                    "The key locator that identifies which key to use for signing. Either this",
                    "or the raw bytes of the target key must be specified.",
                },
                Type: mksignrpc_KeyLocator(),
            },
        },
    }
}
func mksignrpc_KeyLocator() Type {
    return Type{
        Name: "signrpc_KeyLocator",
        Fields: []Field{
            {
                Name: "key_family",
                Description: []string{
                    "The family of key being identified.",
                },
                Type: mkint32(),
            },
            {
                Name: "key_index",
                Description: []string{
                    "The precise index of the key being identified.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mksignrpc_SharedKeyRequest() Type {
    return Type{
        Name: "signrpc_SharedKeyRequest",
        Fields: []Field{
            {
                Name: "ephemeral_pubkey",
                Description: []string{
                    "The ephemeral public key to use for the DH key derivation.",
                },
                Type: mkbytes(),
            },
            {
                Name: "key_loc",
                Description: []string{
                    "Deprecated. The optional key locator of the local key that should be used.",
                    "If this parameter is not set then the node's identity private key will be",
                    "used.",
                },
                Type: mksignrpc_KeyLocator(),
            },
            {
                Name: "key_desc",
                Description: []string{
                    "A key descriptor describes the key used for performing ECDH. Either a key",
                    "locator or a raw public key is expected, if neither is supplied, defaults to",
                    "the node's identity private key.",
                },
                Type: mksignrpc_KeyDescriptor(),
            },
        },
    }
}
func mksignrpc_SharedKeyResponse() Type {
    return Type{
        Name: "signrpc_SharedKeyResponse",
        Fields: []Field{
            {
                Name: "shared_key",
                Description: []string{
                    "The shared public key, hashed with sha256.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_SignDescriptor() Type {
    return Type{
        Name: "signrpc_SignDescriptor",
        Fields: []Field{
            {
                Name: "key_desc",
                Description: []string{
                    "A descriptor that precisely describes *which* key to use for signing. This",
                    "may provide the raw public key directly, or require the Signer to re-derive",
                    "the key according to the populated derivation path.",
                    "",
                    "Note that if the key descriptor was obtained through walletrpc.DeriveKey,",
                    "then the key locator MUST always be provided, since the derived keys are not",
                    "persisted unlike with DeriveNextKey.",
                },
                Type: mksignrpc_KeyDescriptor(),
            },
            {
                Name: "single_tweak",
                Description: []string{
                    "A scalar value that will be added to the private key corresponding to the",
                    "above public key to obtain the private key to be used to sign this input.",
                    "This value is typically derived via the following computation:",
                    "",
                    "derivedKey = privkey + sha256(perCommitmentPoint || pubKey) mod N",
                },
                Type: mkbytes(),
            },
            {
                Name: "double_tweak",
                Description: []string{
                    "A private key that will be used in combination with its corresponding",
                    "private key to derive the private key that is to be used to sign the target",
                    "input. Within the Lightning protocol, this value is typically the",
                    "commitment secret from a previously revoked commitment transaction. This",
                    "value is in combination with two hash values, and the original private key",
                    "to derive the private key to be used when signing.",
                    "",
                    "k = (privKey*sha256(pubKey || tweakPub) +",
                    "tweakPriv*sha256(tweakPub || pubKey)) mod N",
                },
                Type: mkbytes(),
            },
            {
                Name: "witness_script",
                Description: []string{
                    "The full script required to properly redeem the output.  This field will",
                    "only be populated if a p2wsh or a p2sh output is being signed.",
                },
                Type: mkbytes(),
            },
            {
                Name: "output",
                Description: []string{
                    "A description of the output being spent. The value and script MUST be",
                    "provided.",
                },
                Type: mksignrpc_TxOut(),
            },
            {
                Name: "sighash",
                Description: []string{
                    "The target sighash type that should be used when generating the final",
                    "sighash, and signature.",
                },
                Type: mkuint32(),
            },
            {
                Name: "input_index",
                Description: []string{
                    "The target input within the transaction that should be signed.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mksignrpc_SignMessageReq() Type {
    return Type{
        Name: "signrpc_SignMessageReq",
        Fields: []Field{
            {
                Name: "msg",
                Description: []string{
                    "The message to be signed.",
                },
                Type: mkbytes(),
            },
            {
                Name: "key_loc",
                Description: []string{
                    "The key locator that identifies which key to use for signing.",
                },
                Type: mksignrpc_KeyLocator(),
            },
        },
    }
}
func mksignrpc_SignMessageResp() Type {
    return Type{
        Name: "signrpc_SignMessageResp",
        Fields: []Field{
            {
                Name: "signature",
                Description: []string{
                    "The signature for the given message in the fixed-size LN wire format.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_SignReq() Type {
    return Type{
        Name: "signrpc_SignReq",
        Fields: []Field{
            {
                Name: "raw_tx_bytes",
                Description: []string{
                    "The raw bytes of the transaction to be signed.",
                },
                Type: mkbytes(),
            },
            {
                Name: "sign_descs",
                Description: []string{
                    "A set of sign descriptors, for each input to be signed.",
                },
                Repeated: true,
                Type: mksignrpc_SignDescriptor(),
            },
        },
    }
}
func mksignrpc_SignResp() Type {
    return Type{
        Name: "signrpc_SignResp",
        Fields: []Field{
            {
                Name: "raw_sigs",
                Description: []string{
                    "A set of signatures realized in a fixed 64-byte format ordered in ascending",
                    "input order.",
                },
                Repeated: true,
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_TxOut() Type {
    return Type{
        Name: "signrpc_TxOut",
        Fields: []Field{
            {
                Name: "value",
                Description: []string{
                    "The value of the output being spent.",
                },
                Type: mkint64(),
            },
            {
                Name: "pk_script",
                Description: []string{
                    "The script of the output being spent.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_VerifyMessageReq() Type {
    return Type{
        Name: "signrpc_VerifyMessageReq",
        Fields: []Field{
            {
                Name: "msg",
                Description: []string{
                    "The message over which the signature is to be verified.",
                },
                Type: mkbytes(),
            },
            {
                Name: "signature",
                Description: []string{
                    "The fixed-size LN wire encoded signature to be verified over the given",
                    "message.",
                },
                Type: mkbytes(),
            },
            {
                Name: "pubkey",
                Description: []string{
                    "The public key the signature has to be valid for.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mksignrpc_VerifyMessageResp() Type {
    return Type{
        Name: "signrpc_VerifyMessageResp",
        Fields: []Field{
            {
                Name: "valid",
                Description: []string{
                    "Whether the signature was valid over the given message.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkautopilotrpc_ModifyStatusRequest() Type {
    return Type{
        Name: "autopilotrpc_ModifyStatusRequest",
        Fields: []Field{
            {
                Name: "enable",
                Description: []string{
                    "Whether the autopilot agent should be enabled or not.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkautopilotrpc_ModifyStatusResponse() Type {
    return Type{
        Name: "autopilotrpc_ModifyStatusResponse",
    }
}
func mkautopilotrpc_QueryScoresRequest() Type {
    return Type{
        Name: "autopilotrpc_QueryScoresRequest",
        Fields: []Field{
            {
                Name: "pubkeys",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "ignore_local_state",
                Description: []string{
                    "If set, we will ignore the local channel state when calculating scores.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkautopilotrpc_QueryScoresResponse() Type {
    return Type{
        Name: "autopilotrpc_QueryScoresResponse",
        Fields: []Field{
            {
                Name: "results",
                Repeated: true,
                Type: mkautopilotrpc_QueryScoresResponse_HeuristicResult(),
            },
        },
    }
}
func mkautopilotrpc_QueryScoresResponse_HeuristicResult() Type {
    return Type{
        Name: "autopilotrpc_QueryScoresResponse_HeuristicResult",
        Fields: []Field{
            {
                Name: "heuristic",
                Type: mkstring(),
            },
            {
                Name: "scores",
                Repeated: true,
                Type: mkautopilotrpc_QueryScoresResponse_HeuristicResult_ScoresEntry(),
            },
        },
    }
}
func mkautopilotrpc_QueryScoresResponse_HeuristicResult_ScoresEntry() Type {
    return Type{
        Name: "autopilotrpc_QueryScoresResponse_HeuristicResult_ScoresEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mkdouble(),
            },
        },
    }
}
func mkautopilotrpc_SetScoresRequest() Type {
    return Type{
        Name: "autopilotrpc_SetScoresRequest",
        Fields: []Field{
            {
                Name: "heuristic",
                Description: []string{
                    "The name of the heuristic to provide scores to.",
                },
                Type: mkstring(),
            },
            {
                Name: "scores",
                Description: []string{
                    "A map from hex-encoded public keys to scores. Scores must be in the range",
                    "[0.0, 1.0].",
                },
                Repeated: true,
                Type: mkautopilotrpc_SetScoresRequest_ScoresEntry(),
            },
        },
    }
}
func mkautopilotrpc_SetScoresRequest_ScoresEntry() Type {
    return Type{
        Name: "autopilotrpc_SetScoresRequest_ScoresEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mkdouble(),
            },
        },
    }
}
func mkautopilotrpc_SetScoresResponse() Type {
    return Type{
        Name: "autopilotrpc_SetScoresResponse",
    }
}
func mkautopilotrpc_StatusRequest() Type {
    return Type{
        Name: "autopilotrpc_StatusRequest",
    }
}
func mkautopilotrpc_StatusResponse() Type {
    return Type{
        Name: "autopilotrpc_StatusResponse",
        Fields: []Field{
            {
                Name: "active",
                Description: []string{
                    "Indicates whether the autopilot is active or not.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkrestrpc_RestEmptyResponse() Type {
    return Type{
        Name: "restrpc_RestEmptyResponse",
    }
}
func mkrestrpc_WebSocketError() Type {
    return Type{
        Name: "restrpc_WebSocketError",
        Description: []string{
            "WebSocket request and response JSon messages",
        },
        Fields: []Field{
            {
                Name: "http_code",
                Type: mkuint32(),
            },
            {
                Name: "message",
                Type: mkstring(),
            },
        },
    }
}
func mkrestrpc_WebSocketProtobufRequest() Type {
    return Type{
        Name: "restrpc_WebSocketProtobufRequest",
        Description: []string{
            "WebSocket request and response protobuf messages",
        },
        Fields: []Field{
            {
                Name: "endpoint",
                Description: []string{
                    "The rest endpoint to send the request to",
                },
                Type: mkstring(),
            },
            {
                Name: "request_id",
                Description: []string{
                    "An arbitrary string which will be reflected back in the response",
                },
                Type: mkstring(),
            },
            {
                Name: "has_more",
                Type: mkbool(),
            },
            {
                Name: "payload",
                Description: []string{
                    "The data to post to the REST endpoint, if any.",
                    "Make sure this is the correct data structure based on the endpoint you are posting to.",
                },
                Type: mkgoogle_protobuf_Any(),
            },
        },
    }
}
func mkrestrpc_WebSocketProtobufResponse() Type {
    return Type{
        Name: "restrpc_WebSocketProtobufResponse",
        Fields: []Field{
            {
                Name: "request_id",
                Type: mkstring(),
            },
            {
                Name: "has_more",
                Type: mkbool(),
            },
            {
                Name: "ok",
                Type: mkgoogle_protobuf_Any(),
            },
            {
                Name: "error",
                Type: mkrestrpc_WebSocketError(),
            },
        },
    }
}
func mkwalletrpc_AddrRequest() Type {
    return Type{
        Name: "walletrpc_AddrRequest",
        Description: []string{
            "No fields, as we always give out a p2wkh address.",
        },
    }
}
func mkwalletrpc_AddrResponse() Type {
    return Type{
        Name: "walletrpc_AddrResponse",
        Fields: []Field{
            {
                Name: "addr",
                Description: []string{
                    "The address encoded using a bech32 format.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkwalletrpc_BumpFeeRequest() Type {
    return Type{
        Name: "walletrpc_BumpFeeRequest",
        Fields: []Field{
            {
                Name: "outpoint",
                Description: []string{
                    "The input we're attempting to bump the fee of.",
                },
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that the input should be spent within.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "The fee rate, expressed in sat/byte, that should be used to spend the input",
                    "with.",
                },
                Type: mkuint32(),
            },
            {
                Name: "force",
                Description: []string{
                    "Whether this input must be force-swept. This means that it is swept even",
                    "if it has a negative yield.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwalletrpc_BumpFeeResponse() Type {
    return Type{
        Name: "walletrpc_BumpFeeResponse",
    }
}
func mkwalletrpc_EstimateFeeRequest() Type {
    return Type{
        Name: "walletrpc_EstimateFeeRequest",
        Fields: []Field{
            {
                Name: "conf_target",
                Description: []string{
                    "The number of confirmations to shoot for when estimating the fee.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mkwalletrpc_EstimateFeeResponse() Type {
    return Type{
        Name: "walletrpc_EstimateFeeResponse",
        Fields: []Field{
            {
                Name: "sat_per_kw",
                Description: []string{
                    "The amount of satoshis per kw that should be used in order to reach the",
                    "confirmation target in the request.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mkwalletrpc_FinalizePsbtRequest() Type {
    return Type{
        Name: "walletrpc_FinalizePsbtRequest",
        Fields: []Field{
            {
                Name: "funded_psbt",
                Description: []string{
                    "A PSBT that should be signed and finalized. The PSBT must contain all",
                    "required inputs, outputs, UTXO data and partial signatures of all other",
                    "signers.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkwalletrpc_FinalizePsbtResponse() Type {
    return Type{
        Name: "walletrpc_FinalizePsbtResponse",
        Fields: []Field{
            {
                Name: "signed_psbt",
                Description: []string{
                    "The fully signed and finalized transaction in PSBT format.",
                },
                Type: mkbytes(),
            },
            {
                Name: "raw_final_tx",
                Description: []string{
                    "The fully signed and finalized transaction in the raw wire format.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkwalletrpc_FundPsbtRequest() Type {
    return Type{
        Name: "walletrpc_FundPsbtRequest",
        Fields: []Field{
            {
                Name: "psbt",
                Description: []string{
                    "Use an existing PSBT packet as the template for the funded PSBT.",
                    "",
                    "The packet must contain at least one non-dust output. If one or more",
                    "inputs are specified, no coin selection is performed. In that case every",
                    "input must be an UTXO known to the wallet that has not been locked",
                    "before. The sum of all inputs must be sufficiently greater than the sum",
                    "of all outputs to pay a miner fee with the specified fee rate. A change",
                    "output is added to the PSBT if necessary.",
                },
                Type: mkbytes(),
            },
            {
                Name: "raw",
                Description: []string{
                    "Use the outputs and optional inputs from this raw template.",
                },
                Type: mkwalletrpc_TxTemplate(),
            },
            {
                Name: "target_conf",
                Description: []string{
                    "The target number of blocks that the transaction should be confirmed in.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sat_per_vbyte",
                Description: []string{
                    "The fee rate, expressed in sat/vbyte, that should be used to spend the",
                    "input with.",
                },
                Type: mkuint32(),
            },
        },
    }
}
func mkwalletrpc_FundPsbtResponse() Type {
    return Type{
        Name: "walletrpc_FundPsbtResponse",
        Fields: []Field{
            {
                Name: "funded_psbt",
                Description: []string{
                    "The funded but not yet signed PSBT packet.",
                },
                Type: mkbytes(),
            },
            {
                Name: "change_output_index",
                Description: []string{
                    "The index of the added change output or -1 if no change was left over.",
                },
                Type: mkint32(),
            },
            {
                Name: "locked_utxos",
                Description: []string{
                    "The list of lock leases that were acquired for the inputs in the funded PSBT",
                    "packet.",
                },
                Repeated: true,
                Type: mkwalletrpc_UtxoLease(),
            },
        },
    }
}
func mkwalletrpc_KeyReq() Type {
    return Type{
        Name: "walletrpc_KeyReq",
        Fields: []Field{
            {
                Name: "key_finger_print",
                Description: []string{
                    "Is the key finger print of the root pubkey that this request is targeting.",
                    "This allows the WalletKit to possibly serve out keys for multiple HD chains",
                    "via public derivation.",
                },
                Type: mkint32(),
            },
            {
                Name: "key_family",
                Description: []string{
                    "The target key family to derive a key from. In other contexts, this is",
                    "known as the \"account\".",
                },
                Type: mkint32(),
            },
        },
    }
}
func mkwalletrpc_LabelTransactionRequest() Type {
    return Type{
        Name: "walletrpc_LabelTransactionRequest",
        Fields: []Field{
            {
                Name: "txid",
                Description: []string{
                    "The txid of the transaction to label.",
                },
                Type: mkbytes(),
            },
            {
                Name: "label",
                Description: []string{
                    "The label to add to the transaction, limited to 500 characters.",
                },
                Type: mkstring(),
            },
            {
                Name: "overwrite",
                Description: []string{
                    "Whether to overwrite the existing label, if it is present.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwalletrpc_LabelTransactionResponse() Type {
    return Type{
        Name: "walletrpc_LabelTransactionResponse",
    }
}
func mkwalletrpc_LeaseOutputRequest() Type {
    return Type{
        Name: "walletrpc_LeaseOutputRequest",
        Fields: []Field{
            {
                Name: "id",
                Description: []string{
                    "An ID of 32 random bytes that must be unique for each distinct application",
                    "using this RPC which will be used to bound the output lease to.",
                },
                Type: mkbytes(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The identifying outpoint of the output being leased.",
                },
                Type: mklnrpc_OutPoint(),
            },
        },
    }
}
func mkwalletrpc_LeaseOutputResponse() Type {
    return Type{
        Name: "walletrpc_LeaseOutputResponse",
        Fields: []Field{
            {
                Name: "expiration",
                Description: []string{
                    "The absolute expiration of the output lease represented as a unix timestamp.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mkwalletrpc_ListSweepsRequest() Type {
    return Type{
        Name: "walletrpc_ListSweepsRequest",
        Fields: []Field{
            {
                Name: "verbose",
                Description: []string{
                    "Retrieve the full sweep transaction details. If false, only the sweep txids",
                    "will be returned. Note that some sweeps that LND publishes will have been",
                    "replaced-by-fee, so will not be included in this output.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwalletrpc_ListSweepsResponse() Type {
    return Type{
        Name: "walletrpc_ListSweepsResponse",
        Fields: []Field{
            {
                Name: "transaction_details",
                Type: mklnrpc_TransactionDetails(),
            },
            {
                Name: "transaction_ids",
                Type: mkwalletrpc_ListSweepsResponse_TransactionIDs(),
            },
        },
    }
}
func mkwalletrpc_ListSweepsResponse_TransactionIDs() Type {
    return Type{
        Name: "walletrpc_ListSweepsResponse_TransactionIDs",
        Fields: []Field{
            {
                Name: "transaction_ids",
                Description: []string{
                    "Reversed, hex-encoded string representing the transaction ids of the",
                    "sweeps that our node has broadcast. Note that these transactions may",
                    "not have confirmed yet, we record sweeps on broadcast, not confirmation.",
                },
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mkwalletrpc_ListUnspentRequest() Type {
    return Type{
        Name: "walletrpc_ListUnspentRequest",
        Fields: []Field{
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations to be included.",
                },
                Type: mkint32(),
            },
            {
                Name: "max_confs",
                Description: []string{
                    "The maximum number of confirmations to be included.",
                },
                Type: mkint32(),
            },
        },
    }
}
func mkwalletrpc_ListUnspentResponse() Type {
    return Type{
        Name: "walletrpc_ListUnspentResponse",
        Fields: []Field{
            {
                Name: "utxos",
                Description: []string{
                    "A list of utxos satisfying the specified number of confirmations.",
                },
                Repeated: true,
                Type: mklnrpc_Utxo(),
            },
        },
    }
}
func mkwalletrpc_PendingSweep() Type {
    return Type{
        Name: "walletrpc_PendingSweep",
        Fields: []Field{
            {
                Name: "outpoint",
                Description: []string{
                    "The outpoint of the output we're attempting to sweep.",
                },
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "witness_type",
                Description: []string{
                    "The witness type of the output we're attempting to sweep.",
                },
                Type: mkwalletrpc_WitnessType(),
            },
            {
                Name: "amount_sat",
                Description: []string{
                    "The value of the output we're attempting to sweep.",
                },
                Type: mkuint32(),
            },
            {
                Name: "sat_per_byte",
                Description: []string{
                    "The fee rate we'll use to sweep the output. The fee rate is only determined",
                    "once a sweeping transaction for the output is created, so it's possible for",
                    "this to be 0 before this.",
                },
                Type: mkuint32(),
            },
            {
                Name: "broadcast_attempts",
                Description: []string{
                    "The number of broadcast attempts we've made to sweep the output.",
                },
                Type: mkuint32(),
            },
            {
                Name: "next_broadcast_height",
                Description: []string{
                    "The next height of the chain at which we'll attempt to broadcast the",
                    "sweep transaction of the output.",
                },
                Type: mkuint32(),
            },
            {
                Name: "requested_conf_target",
                Description: []string{
                    "The requested confirmation target for this output.",
                },
                Type: mkuint32(),
            },
            {
                Name: "requested_sat_per_byte",
                Description: []string{
                    "The requested fee rate, expressed in sat/byte, for this output.",
                },
                Type: mkuint32(),
            },
            {
                Name: "force",
                Description: []string{
                    "Whether this input must be force-swept. This means that it is swept even",
                    "if it has a negative yield.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwalletrpc_PendingSweepsRequest() Type {
    return Type{
        Name: "walletrpc_PendingSweepsRequest",
    }
}
func mkwalletrpc_PendingSweepsResponse() Type {
    return Type{
        Name: "walletrpc_PendingSweepsResponse",
        Fields: []Field{
            {
                Name: "pending_sweeps",
                Description: []string{
                    "The set of outputs currently being swept by lnd's central batching engine.",
                },
                Repeated: true,
                Type: mkwalletrpc_PendingSweep(),
            },
        },
    }
}
func mkwalletrpc_PublishResponse() Type {
    return Type{
        Name: "walletrpc_PublishResponse",
        Fields: []Field{
            {
                Name: "publish_error",
                Description: []string{
                    "If blank, then no error occurred and the transaction was successfully",
                    "published. If not the empty string, then a string representation of the",
                    "broadcast error.",
                    "",
                    "TODO(roasbeef): map to a proper enum type",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkwalletrpc_ReleaseOutputRequest() Type {
    return Type{
        Name: "walletrpc_ReleaseOutputRequest",
        Fields: []Field{
            {
                Name: "id",
                Description: []string{
                    "The unique ID that was used to lock the output.",
                },
                Type: mkbytes(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The identifying outpoint of the output being released.",
                },
                Type: mklnrpc_OutPoint(),
            },
        },
    }
}
func mkwalletrpc_ReleaseOutputResponse() Type {
    return Type{
        Name: "walletrpc_ReleaseOutputResponse",
    }
}
func mkwalletrpc_SendOutputsRequest() Type {
    return Type{
        Name: "walletrpc_SendOutputsRequest",
        Fields: []Field{
            {
                Name: "sat_per_kw",
                Description: []string{
                    "The number of satoshis per kilo weight that should be used when crafting",
                    "this transaction.",
                },
                Type: mkint64(),
            },
            {
                Name: "outputs",
                Description: []string{
                    "A slice of the outputs that should be created in the transaction produced.",
                },
                Repeated: true,
                Type: mksignrpc_TxOut(),
            },
            {
                Name: "label",
                Description: []string{
                    "An optional label for the transaction, limited to 500 characters.",
                },
                Type: mkstring(),
            },
            {
                Name: "min_confs",
                Description: []string{
                    "The minimum number of confirmations each one of your outputs used for",
                    "the transaction must satisfy.",
                },
                Type: mkint32(),
            },
            {
                Name: "spend_unconfirmed",
                Description: []string{
                    "Whether unconfirmed outputs should be used as inputs for the transaction.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkwalletrpc_SendOutputsResponse() Type {
    return Type{
        Name: "walletrpc_SendOutputsResponse",
        Fields: []Field{
            {
                Name: "raw_tx",
                Description: []string{
                    "The serialized transaction sent out on the network.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkwalletrpc_Transaction() Type {
    return Type{
        Name: "walletrpc_Transaction",
        Fields: []Field{
            {
                Name: "tx_hex",
                Description: []string{
                    "The raw serialized transaction.",
                },
                Type: mkbytes(),
            },
            {
                Name: "label",
                Description: []string{
                    "An optional label to save with the transaction. Limited to 500 characters.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkwalletrpc_TxTemplate() Type {
    return Type{
        Name: "walletrpc_TxTemplate",
        Fields: []Field{
            {
                Name: "inputs",
                Description: []string{
                    "An optional list of inputs to use. Every input must be an UTXO known to the",
                    "wallet that has not been locked before. The sum of all inputs must be",
                    "sufficiently greater than the sum of all outputs to pay a miner fee with the",
                    "fee rate specified in the parent message.",
                    "",
                    "If no inputs are specified, coin selection will be performed instead and",
                    "inputs of sufficient value will be added to the resulting PSBT.",
                },
                Repeated: true,
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "outputs",
                Description: []string{
                    "A map of all addresses and the amounts to send to in the funded PSBT.",
                },
                Repeated: true,
                Type: mkwalletrpc_TxTemplate_OutputsEntry(),
            },
        },
    }
}
func mkwalletrpc_TxTemplate_OutputsEntry() Type {
    return Type{
        Name: "walletrpc_TxTemplate_OutputsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkstring(),
            },
            {
                Name: "value",
                Type: mkuint64(),
            },
        },
    }
}
func mkwalletrpc_UtxoLease() Type {
    return Type{
        Name: "walletrpc_UtxoLease",
        Fields: []Field{
            {
                Name: "id",
                Description: []string{
                    "A 32 byte random ID that identifies the lease.",
                },
                Type: mkbytes(),
            },
            {
                Name: "outpoint",
                Description: []string{
                    "The identifying outpoint of the output being leased.",
                },
                Type: mklnrpc_OutPoint(),
            },
            {
                Name: "expiration",
                Description: []string{
                    "The absolute expiration of the output lease represented as a unix timestamp.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mkrouterrpc_BuildRouteRequest() Type {
    return Type{
        Name: "routerrpc_BuildRouteRequest",
        Fields: []Field{
            {
                Name: "amt_msat",
                Description: []string{
                    "The amount to send expressed in msat. If set to zero, the minimum routable",
                    "amount is used.",
                },
                Type: mkint64(),
            },
            {
                Name: "final_cltv_delta",
                Description: []string{
                    "CLTV delta from the current height that should be used for the timelock",
                    "of the final hop",
                },
                Type: mkint32(),
            },
            {
                Name: "outgoing_chan_id",
                Description: []string{
                    "The channel id of the channel that must be taken to the first hop. If zero,",
                    "any channel may be used.",
                },
                Type: mkuint64(),
            },
            {
                Name: "hop_pubkeys",
                Description: []string{
                    "A list of hops that defines the route. This does not include the source hop",
                    "pubkey.",
                },
                Repeated: true,
                Type: mkbytes(),
            },
        },
    }
}
func mkrouterrpc_BuildRouteResponse() Type {
    return Type{
        Name: "routerrpc_BuildRouteResponse",
        Fields: []Field{
            {
                Name: "route",
                Description: []string{
                    "Fully specified route that can be used to execute the payment.",
                },
                Type: mklnrpc_Route(),
            },
        },
    }
}
func mkrouterrpc_CircuitKey() Type {
    return Type{
        Name: "routerrpc_CircuitKey",
        Fields: []Field{
            {
                Name: "chan_id",
                Description: []string{
                    "The id of the channel that the is part of this circuit.",
                },
                Type: mkuint64(),
            },
            {
                Name: "htlc_id",
                Description: []string{
                    "The index of the incoming htlc in the incoming channel.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mkrouterrpc_ForwardEvent() Type {
    return Type{
        Name: "routerrpc_ForwardEvent",
        Fields: []Field{
            {
                Name: "info",
                Description: []string{
                    "Info contains details about the htlc that was forwarded.",
                },
                Type: mkrouterrpc_HtlcInfo(),
            },
        },
    }
}
func mkrouterrpc_ForwardFailEvent() Type {
    return Type{
        Name: "routerrpc_ForwardFailEvent",
    }
}
func mkrouterrpc_ForwardHtlcInterceptRequest() Type {
    return Type{
        Name: "routerrpc_ForwardHtlcInterceptRequest",
        Fields: []Field{
            {
                Name: "incoming_circuit_key",
                Description: []string{
                    "The key of this forwarded htlc. It defines the incoming channel id and",
                    "the index in this channel.",
                },
                Type: mkrouterrpc_CircuitKey(),
            },
            {
                Name: "incoming_amount_msat",
                Description: []string{
                    "The incoming htlc amount.",
                },
                Type: mkuint64(),
            },
            {
                Name: "incoming_expiry",
                Description: []string{
                    "The incoming htlc expiry.",
                },
                Type: mkuint32(),
            },
            {
                Name: "payment_hash",
                Description: []string{
                    "The htlc payment hash. This value is not guaranteed to be unique per",
                    "request.",
                },
                Type: mkbytes(),
            },
            {
                Name: "outgoing_requested_chan_id",
                Description: []string{
                    "The requested outgoing channel id for this forwarded htlc. Because of",
                    "non-strict forwarding, this isn't necessarily the channel over which the",
                    "packet will be forwarded eventually. A different channel to the same peer",
                    "may be selected as well.",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_amount_msat",
                Description: []string{
                    "The outgoing htlc amount.",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_expiry",
                Description: []string{
                    "The outgoing htlc expiry.",
                },
                Type: mkuint32(),
            },
            {
                Name: "custom_records",
                Description: []string{
                    "Any custom records that were present in the payload.",
                },
                Repeated: true,
                Type: mkrouterrpc_ForwardHtlcInterceptRequest_CustomRecordsEntry(),
            },
            {
                Name: "onion_blob",
                Description: []string{
                    "The onion blob for the next hop",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkrouterrpc_ForwardHtlcInterceptRequest_CustomRecordsEntry() Type {
    return Type{
        Name: "routerrpc_ForwardHtlcInterceptRequest_CustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mkrouterrpc_ForwardHtlcInterceptResponse() Type {
    return Type{
        Name: "routerrpc_ForwardHtlcInterceptResponse",
        Description: []string{
            "ForwardHtlcInterceptResponse enables the caller to resolve a previously hold",
            "forward. The caller can choose either to:",
            "- `Resume`: Execute the default behavior (usually forward).",
            "- `Reject`: Fail the htlc backwards.",
            "- `Settle`: Settle this htlc with a given preimage.",
        },
        Fields: []Field{
            {
                Name: "incoming_circuit_key",
                Description: []string{
                    "The key of this forwarded htlc. It defines the incoming channel id and",
                    "the index in this channel.",
                },
                Type: mkrouterrpc_CircuitKey(),
            },
            {
                Name: "action",
                Description: []string{
                    "The resolve action for this intercepted htlc.",
                },
                Type: mkrouterrpc_ResolveHoldForwardAction(),
            },
            {
                Name: "preimage",
                Description: []string{
                    "The preimage in case the resolve action is Settle.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkrouterrpc_HtlcEvent() Type {
    return Type{
        Name: "routerrpc_HtlcEvent",
        Description: []string{
            "HtlcEvent contains the htlc event that was processed. These are served on a",
            "best-effort basis; events are not persisted, delivery is not guaranteed",
            "(in the event of a crash in the switch, forward events may be lost) and",
            "some events may be replayed upon restart. Events consumed from this package",
            "should be de-duplicated by the htlc's unique combination of incoming and",
            "outgoing channel id and htlc id. [EXPERIMENTAL]",
        },
        Fields: []Field{
            {
                Name: "incoming_channel_id",
                Description: []string{
                    "The short channel id that the incoming htlc arrived at our node on. This",
                    "value is zero for sends.",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_channel_id",
                Description: []string{
                    "The short channel id that the outgoing htlc left our node on. This value",
                    "is zero for receives.",
                },
                Type: mkuint64(),
            },
            {
                Name: "incoming_htlc_id",
                Description: []string{
                    "Incoming id is the index of the incoming htlc in the incoming channel.",
                    "This value is zero for sends.",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_htlc_id",
                Description: []string{
                    "Outgoing id is the index of the outgoing htlc in the outgoing channel.",
                    "This value is zero for receives.",
                },
                Type: mkuint64(),
            },
            {
                Name: "timestamp_ns",
                Description: []string{
                    "The time in unix nanoseconds that the event occurred.",
                },
                Type: mkuint64(),
            },
            {
                Name: "event_type",
                Description: []string{
                    "The event type indicates whether the htlc was part of a send, receive or",
                    "forward.",
                },
                Type: mkrouterrpc_HtlcEvent_EventType(),
            },
            {
                Name: "forward_event",
                Type: mkrouterrpc_ForwardEvent(),
            },
            {
                Name: "forward_fail_event",
                Type: mkrouterrpc_ForwardFailEvent(),
            },
            {
                Name: "settle_event",
                Type: mkrouterrpc_SettleEvent(),
            },
            {
                Name: "link_fail_event",
                Type: mkrouterrpc_LinkFailEvent(),
            },
        },
    }
}
func mkrouterrpc_HtlcInfo() Type {
    return Type{
        Name: "routerrpc_HtlcInfo",
        Fields: []Field{
            {
                Name: "incoming_timelock",
                Description: []string{
                    "The timelock on the incoming htlc.",
                },
                Type: mkuint32(),
            },
            {
                Name: "outgoing_timelock",
                Description: []string{
                    "The timelock on the outgoing htlc.",
                },
                Type: mkuint32(),
            },
            {
                Name: "incoming_amt_msat",
                Description: []string{
                    "The amount of the incoming htlc.",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_amt_msat",
                Description: []string{
                    "The amount of the outgoing htlc.",
                },
                Type: mkuint64(),
            },
        },
    }
}
func mkrouterrpc_LinkFailEvent() Type {
    return Type{
        Name: "routerrpc_LinkFailEvent",
        Fields: []Field{
            {
                Name: "info",
                Description: []string{
                    "Info contains details about the htlc that we failed.",
                },
                Type: mkrouterrpc_HtlcInfo(),
            },
            {
                Name: "wire_failure",
                Description: []string{
                    "FailureCode is the BOLT error code for the failure.",
                },
                Type: mklnrpc_Failure_FailureCode(),
            },
            {
                Name: "failure_detail",
                Description: []string{
                    "FailureDetail provides additional information about the reason for the",
                    "failure. This detail enriches the information provided by the wire message",
                    "and may be 'no detail' if the wire message requires no additional metadata.",
                },
                Type: mkrouterrpc_FailureDetail(),
            },
            {
                Name: "failure_string",
                Description: []string{
                    "A string representation of the link failure.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkrouterrpc_PairData() Type {
    return Type{
        Name: "routerrpc_PairData",
        Fields: []Field{
            {
                Name: "fail_time",
                Description: []string{
                    "Time of last failure.",
                },
                Type: mkint64(),
            },
            {
                Name: "fail_amt_sat",
                Description: []string{
                    "Lowest amount that failed to forward rounded to whole sats. This may be",
                    "set to zero if the failure is independent of amount.",
                },
                Type: mkint64(),
            },
            {
                Name: "fail_amt_msat",
                Description: []string{
                    "Lowest amount that failed to forward in millisats. This may be",
                    "set to zero if the failure is independent of amount.",
                },
                Type: mkint64(),
            },
            {
                Name: "success_time",
                Description: []string{
                    "Time of last success.",
                },
                Type: mkint64(),
            },
            {
                Name: "success_amt_sat",
                Description: []string{
                    "Highest amount that we could successfully forward rounded to whole sats.",
                },
                Type: mkint64(),
            },
            {
                Name: "success_amt_msat",
                Description: []string{
                    "Highest amount that we could successfully forward in millisats.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mkrouterrpc_PairHistory() Type {
    return Type{
        Name: "routerrpc_PairHistory",
        Description: []string{
            "PairHistory contains the mission control state for a particular node pair.",
        },
        Fields: []Field{
            {
                Name: "node_from",
                Description: []string{
                    "The source node pubkey of the pair.",
                },
                Type: mkbytes(),
            },
            {
                Name: "node_to",
                Description: []string{
                    "The destination node pubkey of the pair.",
                },
                Type: mkbytes(),
            },
            {
                Name: "history",
                Type: mkrouterrpc_PairData(),
            },
        },
    }
}
func mkrouterrpc_PaymentStatus() Type {
    return Type{
        Name: "routerrpc_PaymentStatus",
        Fields: []Field{
            {
                Name: "state",
                Description: []string{
                    "Current state the payment is in.",
                },
                Type: mkrouterrpc_PaymentState(),
            },
            {
                Name: "preimage",
                Description: []string{
                    "The pre-image of the payment when state is SUCCEEDED.",
                },
                Type: mkbytes(),
            },
            {
                Name: "htlcs",
                Description: []string{
                    "The HTLCs made in attempt to settle the payment [EXPERIMENTAL].",
                },
                Repeated: true,
                Type: mklnrpc_HTLCAttempt(),
            },
        },
    }
}
func mkrouterrpc_QueryMissionControlRequest() Type {
    return Type{
        Name: "routerrpc_QueryMissionControlRequest",
    }
}
func mkrouterrpc_QueryMissionControlResponse() Type {
    return Type{
        Name: "routerrpc_QueryMissionControlResponse",
        Description: []string{
            "QueryMissionControlResponse contains mission control state.",
        },
        Fields: []Field{
            {
                Name: "pairs",
                Description: []string{
                    "Node pair-level mission control state.",
                },
                Repeated: true,
                Type: mkrouterrpc_PairHistory(),
            },
        },
    }
}
func mkrouterrpc_QueryProbabilityRequest() Type {
    return Type{
        Name: "routerrpc_QueryProbabilityRequest",
        Fields: []Field{
            {
                Name: "from_node",
                Description: []string{
                    "The source node pubkey of the pair.",
                },
                Type: mkbytes(),
            },
            {
                Name: "to_node",
                Description: []string{
                    "The destination node pubkey of the pair.",
                },
                Type: mkbytes(),
            },
            {
                Name: "amt_msat",
                Description: []string{
                    "The amount for which to calculate a probability.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mkrouterrpc_QueryProbabilityResponse() Type {
    return Type{
        Name: "routerrpc_QueryProbabilityResponse",
        Fields: []Field{
            {
                Name: "probability",
                Description: []string{
                    "The success probability for the requested pair.",
                },
                Type: mkdouble(),
            },
            {
                Name: "history",
                Description: []string{
                    "The historical data for the requested pair.",
                },
                Type: mkrouterrpc_PairData(),
            },
        },
    }
}
func mkrouterrpc_ResetMissionControlRequest() Type {
    return Type{
        Name: "routerrpc_ResetMissionControlRequest",
    }
}
func mkrouterrpc_ResetMissionControlResponse() Type {
    return Type{
        Name: "routerrpc_ResetMissionControlResponse",
    }
}
func mkrouterrpc_RouteFeeRequest() Type {
    return Type{
        Name: "routerrpc_RouteFeeRequest",
        Fields: []Field{
            {
                Name: "dest",
                Description: []string{
                    "The destination once wishes to obtain a routing fee quote to.",
                },
                Type: mkbytes(),
            },
            {
                Name: "amt_sat",
                Description: []string{
                    "The amount one wishes to send to the target destination.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mkrouterrpc_RouteFeeResponse() Type {
    return Type{
        Name: "routerrpc_RouteFeeResponse",
        Fields: []Field{
            {
                Name: "routing_fee_msat",
                Description: []string{
                    "A lower bound of the estimated fee to the target destination within the",
                    "network, expressed in milli-satoshis.",
                },
                Type: mkint64(),
            },
            {
                Name: "time_lock_delay",
                Description: []string{
                    "An estimate of the worst case time delay that can occur. Note that callers",
                    "will still need to factor in the final CLTV delta of the last hop into this",
                    "value.",
                },
                Type: mkint64(),
            },
        },
    }
}
func mkrouterrpc_SendPaymentRequest() Type {
    return Type{
        Name: "routerrpc_SendPaymentRequest",
        Fields: []Field{
            {
                Name: "dest",
                Description: []string{
                    "The identity pubkey of the payment recipient",
                },
                Type: mkbytes(),
            },
            {
                Name: "amt",
                Description: []string{
                    "Number of satoshis to send.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "amt_msat",
                Description: []string{
                    "Number of millisatoshis to send.",
                    "",
                    "The fields amt and amt_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "payment_hash",
                Description: []string{
                    "The hash to use within the payment's HTLC",
                },
                Type: mkbytes(),
            },
            {
                Name: "final_cltv_delta",
                Description: []string{
                    "The CLTV delta from the current height that should be used to set the",
                    "timelock for the final hop.",
                },
                Type: mkint32(),
            },
            {
                Name: "payment_request",
                Description: []string{
                    "A bare-bones invoice for a payment within the Lightning Network.  With the",
                    "details of the invoice, the sender has all the data necessary to send a",
                    "payment to the recipient. The amount in the payment request may be zero. In",
                    "that case it is required to set the amt field as well. If no payment request",
                    "is specified, the following fields are required: dest, amt and payment_hash.",
                },
                Type: mkstring(),
            },
            {
                Name: "timeout_seconds",
                Description: []string{
                    "An upper limit on the amount of time we should spend when attempting to",
                    "fulfill the payment. This is expressed in seconds. If we cannot make a",
                    "successful payment within this time frame, an error will be returned.",
                    "This field must be non-zero.",
                },
                Type: mkint32(),
            },
            {
                Name: "fee_limit_sat",
                Description: []string{
                    "The maximum number of satoshis that will be paid as a fee of the payment.",
                    "If this field is left to the default value of 0, only zero-fee routes will",
                    "be considered. This usually means single hop routes connecting directly to",
                    "the destination. To send the payment without a fee limit, use max int here.",
                    "",
                    "The fields fee_limit_sat and fee_limit_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "fee_limit_msat",
                Description: []string{
                    "The maximum number of millisatoshis that will be paid as a fee of the",
                    "payment. If this field is left to the default value of 0, only zero-fee",
                    "routes will be considered. This usually means single hop routes connecting",
                    "directly to the destination. To send the payment without a fee limit, use",
                    "max int here.",
                    "",
                    "The fields fee_limit_sat and fee_limit_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "outgoing_chan_id",
                Description: []string{
                    "Deprecated, use outgoing_chan_ids. The channel id of the channel that must",
                    "be taken to the first hop. If zero, any channel may be used (unless",
                    "outgoing_chan_ids are set).",
                },
                Type: mkuint64(),
            },
            {
                Name: "outgoing_chan_ids",
                Description: []string{
                    "The channel ids of the channels are allowed for the first hop. If empty,",
                    "any channel may be used.",
                },
                Repeated: true,
                Type: mkuint64(),
            },
            {
                Name: "last_hop_pubkey",
                Description: []string{
                    "The pubkey of the last hop of the route. If empty, any hop may be used.",
                },
                Type: mkbytes(),
            },
            {
                Name: "cltv_limit",
                Description: []string{
                    "An optional maximum total time lock for the route. This should not exceed",
                    "lnd's `--max-cltv-expiry` setting. If zero, then the value of",
                    "`--max-cltv-expiry` is enforced.",
                },
                Type: mkint32(),
            },
            {
                Name: "route_hints",
                Description: []string{
                    "Optional route hints to reach the destination through private channels.",
                },
                Repeated: true,
                Type: mklnrpc_RouteHint(),
            },
            {
                Name: "dest_custom_records",
                Description: []string{
                    "An optional field that can be used to pass an arbitrary set of TLV records",
                    "to a peer which understands the new records. This can be used to pass",
                    "application specific data during the payment attempt. Record types are",
                    "required to be in the custom range >= 65536. When using REST, the values",
                    "must be encoded as base64.",
                },
                Repeated: true,
                Type: mkrouterrpc_SendPaymentRequest_DestCustomRecordsEntry(),
            },
            {
                Name: "allow_self_payment",
                Description: []string{
                    "If set, circular payments to self are permitted.",
                },
                Type: mkbool(),
            },
            {
                Name: "dest_features",
                Description: []string{
                    "Features assumed to be supported by the final node. All transitive feature",
                    "dependencies must also be set properly. For a given feature bit pair, either",
                    "optional or remote may be set, but not both. If this field is nil or empty,",
                    "the router will try to load destination features from the graph as a",
                    "fallback.",
                },
                Repeated: true,
                Type: mklnrpc_FeatureBit(),
            },
            {
                Name: "max_parts",
                Description: []string{
                    "The maximum number of partial payments that may be use to complete the full",
                    "amount.",
                },
                Type: mkuint32(),
            },
            {
                Name: "no_inflight_updates",
                Description: []string{
                    "If set, only the final payment update is streamed back. Intermediate updates",
                    "that show which htlcs are still in flight are suppressed.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkrouterrpc_SendPaymentRequest_DestCustomRecordsEntry() Type {
    return Type{
        Name: "routerrpc_SendPaymentRequest_DestCustomRecordsEntry",
        Fields: []Field{
            {
                Name: "key",
                Type: mkuint64(),
            },
            {
                Name: "value",
                Type: mkbytes(),
            },
        },
    }
}
func mkrouterrpc_SendToRouteRequest() Type {
    return Type{
        Name: "routerrpc_SendToRouteRequest",
        Fields: []Field{
            {
                Name: "payment_hash",
                Description: []string{
                    "The payment hash to use for the HTLC.",
                },
                Type: mkbytes(),
            },
            {
                Name: "route",
                Description: []string{
                    "Route that should be used to attempt to complete the payment.",
                },
                Type: mklnrpc_Route(),
            },
        },
    }
}
func mkrouterrpc_SendToRouteResponse() Type {
    return Type{
        Name: "routerrpc_SendToRouteResponse",
        Fields: []Field{
            {
                Name: "preimage",
                Description: []string{
                    "The preimage obtained by making the payment.",
                },
                Type: mkbytes(),
            },
            {
                Name: "failure",
                Description: []string{
                    "The failure message in case the payment failed.",
                },
                Type: mklnrpc_Failure(),
            },
        },
    }
}
func mkrouterrpc_SettleEvent() Type {
    return Type{
        Name: "routerrpc_SettleEvent",
    }
}
func mkrouterrpc_SubscribeHtlcEventsRequest() Type {
    return Type{
        Name: "routerrpc_SubscribeHtlcEventsRequest",
    }
}
func mkrouterrpc_TrackPaymentRequest() Type {
    return Type{
        Name: "routerrpc_TrackPaymentRequest",
        Fields: []Field{
            {
                Name: "payment_hash",
                Description: []string{
                    "The hash of the payment to look up.",
                },
                Type: mkbytes(),
            },
            {
                Name: "no_inflight_updates",
                Description: []string{
                    "If set, only the final payment update is streamed back. Intermediate updates",
                    "that show which htlcs are still in flight are suppressed.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkhelp_Field() Type {
    return Type{
        Name: "help_Field",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "description",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "repeated",
                Type: mkbool(),
            },
            {
                Name: "type",
                Type: mkhelp_Type(),
            },
        },
    }
}
func mkhelp_RestCommandCategory() Type {
    return Type{
        Name: "help_RestCommandCategory",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "description",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "endpoints",
                Repeated: true,
                Type: mkhelp_RestEndpoint(),
            },
            {
                Name: "subcategory",
                Repeated: true,
                Type: mkhelp_RestCommandCategory(),
            },
        },
    }
}
func mkhelp_RestEndpoint() Type {
    return Type{
        Name: "help_RestEndpoint",
        Fields: []Field{
            {
                Name: "URI",
                Type: mkstring(),
            },
            {
                Name: "short_description",
                Type: mkstring(),
            },
        },
    }
}
func mkhelp_RestHelpResponse() Type {
    return Type{
        Name: "help_RestHelpResponse",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "service",
                Type: mkstring(),
            },
            {
                Name: "description",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "request",
                Type: mkhelp_Type(),
            },
            {
                Name: "response",
                Type: mkhelp_Type(),
            },
        },
    }
}
func mkhelp_RestMasterHelpResponse() Type {
    return Type{
        Name: "help_RestMasterHelpResponse",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "description",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "category",
                Repeated: true,
                Type: mkhelp_RestCommandCategory(),
            },
        },
    }
}
func mkhelp_Type() Type {
    return Type{
        Name: "help_Type",
        Fields: []Field{
            {
                Name: "name",
                Type: mkstring(),
            },
            {
                Name: "description",
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "fields",
                Repeated: true,
                Type: mkhelp_Field(),
            },
        },
    }
}
func mkgoogle_protobuf_Any() Type {
    return Type{
        Name: "google_protobuf_Any",
        Description: []string{
            "`Any` contains an arbitrary serialized protocol buffer message along with a",
            "URL that describes the type of the serialized message.",
            "",
            "Protobuf library provides support to pack/unpack Any values in the form",
            "of utility functions or additional generated methods of the Any type.",
            "",
            "Example 1: Pack and unpack a message in C++.",
            "",
            "    Foo foo = ...;",
            "    Any any;",
            "    any.PackFrom(foo);",
            "    ...",
            "    if (any.UnpackTo(&foo)) {",
            "      ...",
            "    }",
            "",
            "Example 2: Pack and unpack a message in Java.",
            "",
            "    Foo foo = ...;",
            "    Any any = Any.pack(foo);",
            "    ...",
            "    if (any.is(Foo.class)) {",
            "      foo = any.unpack(Foo.class);",
            "    }",
            "",
            " Example 3: Pack and unpack a message in Python.",
            "",
            "    foo = Foo(...)",
            "    any = Any()",
            "    any.Pack(foo)",
            "    ...",
            "    if any.Is(Foo.DESCRIPTOR):",
            "      any.Unpack(foo)",
            "      ...",
            "",
            " Example 4: Pack and unpack a message in Go",
            "",
            "     foo := &pb.Foo{...}",
            "     any, err := ptypes.MarshalAny(foo)",
            "     ...",
            "     foo := &pb.Foo{}",
            "     if err := ptypes.UnmarshalAny(any, foo); err != nil {",
            "       ...",
            "     }",
            "",
            "The pack methods provided by protobuf library will by default use",
            "'type.googleapis.com/full.type.name' as the type URL and the unpack",
            "methods only use the fully qualified type name after the last '/'",
            "in the type URL, for example \"foo.bar.com/x/y.z\" will yield type",
            "name \"y.z\".",
            "",
            "",
            "JSON",
            "====",
            "The JSON representation of an `Any` value uses the regular",
            "representation of the deserialized, embedded message, with an",
            "additional field `@type` which contains the type URL. Example:",
            "",
            "    package google.profile;",
            "    message Person {",
            "      string first_name = 1;",
            "      string last_name = 2;",
            "    }",
            "",
            "    {",
            "      \"@type\": \"type.googleapis.com/google.profile.Person\",",
            "      \"firstName\": <string>,",
            "      \"lastName\": <string>",
            "    }",
            "",
            "If the embedded message type is well-known and has a custom JSON",
            "representation, that representation will be embedded adding a field",
            "`value` which holds the custom JSON in addition to the `@type`",
            "field. Example (for message [google.protobuf.Duration][]):",
            "",
            "    {",
            "      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",",
            "      \"value\": \"1.212s\"",
            "    }",
        },
        Fields: []Field{
            {
                Name: "type_url",
                Description: []string{
                    "A URL/resource name that uniquely identifies the type of the serialized",
                    "protocol buffer message. This string must contain at least",
                    "one \"/\" character. The last segment of the URL's path must represent",
                    "the fully qualified name of the type (as in",
                    "`path/google.protobuf.Duration`). The name should be in a canonical form",
                    "(e.g., leading \".\" is not accepted).",
                    "",
                    "In practice, teams usually precompile into the binary all types that they",
                    "expect it to use in the context of Any. However, for URLs which use the",
                    "scheme `http`, `https`, or no scheme, one can optionally set up a type",
                    "server that maps type URLs to message definitions as follows:",
                    "",
                    "* If no scheme is provided, `https` is assumed.",
                    "* An HTTP GET on the URL must yield a [google.protobuf.Type][]",
                    "  value in binary format, or produce an error.",
                    "* Applications are allowed to cache lookup results based on the",
                    "  URL, or have them precompiled into a binary to avoid any",
                    "  lookup. Therefore, binary compatibility needs to be preserved",
                    "  on changes to types. (Use versioned type names to manage",
                    "  breaking changes.)",
                    "",
                    "Note: this functionality is not currently available in the official",
                    "protobuf release, and it is not used for type URLs beginning with",
                    "type.googleapis.com.",
                    "",
                    "Schemes other than `http`, `https` (or the empty scheme) might be",
                    "used with implementation specific semantics.",
                },
                Type: mkstring(),
            },
            {
                Name: "value",
                Description: []string{
                    "Must be a valid serialized protocol buffer of the above specified type.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkinvoicesrpc_AddHoldInvoiceRequest() Type {
    return Type{
        Name: "invoicesrpc_AddHoldInvoiceRequest",
        Fields: []Field{
            {
                Name: "memo",
                Description: []string{
                    "An optional memo to attach along with the invoice. Used for record keeping",
                    "purposes for the invoice's creator, and will also be set in the description",
                    "field of the encoded payment request if the description_hash field is not",
                    "being used.",
                },
                Type: mkstring(),
            },
            {
                Name: "hash",
                Description: []string{
                    "The hash of the preimage",
                },
                Type: mkbytes(),
            },
            {
                Name: "value",
                Description: []string{
                    "The value of this invoice in satoshis",
                    "",
                    "The fields value and value_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "value_msat",
                Description: []string{
                    "The value of this invoice in millisatoshis",
                    "",
                    "The fields value and value_msat are mutually exclusive.",
                },
                Type: mkint64(),
            },
            {
                Name: "description_hash",
                Description: []string{
                    "Hash (SHA-256) of a description of the payment. Used if the description of",
                    "payment (memo) is too long to naturally fit within the description field",
                    "of an encoded payment request.",
                },
                Type: mkbytes(),
            },
            {
                Name: "expiry",
                Description: []string{
                    "Payment request expiry time in seconds. Default is 3600 (1 hour).",
                },
                Type: mkint64(),
            },
            {
                Name: "fallback_addr",
                Description: []string{
                    "Fallback on-chain address.",
                },
                Type: mkstring(),
            },
            {
                Name: "cltv_expiry",
                Description: []string{
                    "Delta to use for the time-lock of the CLTV extended to the final hop.",
                },
                Type: mkuint64(),
            },
            {
                Name: "route_hints",
                Description: []string{
                    "Route hints that can each be individually used to assist in reaching the",
                    "invoice's destination.",
                },
                Repeated: true,
                Type: mklnrpc_RouteHint(),
            },
            {
                Name: "private",
                Description: []string{
                    "Whether this invoice should include routing hints for private channels.",
                },
                Type: mkbool(),
            },
        },
    }
}
func mkinvoicesrpc_AddHoldInvoiceResp() Type {
    return Type{
        Name: "invoicesrpc_AddHoldInvoiceResp",
        Fields: []Field{
            {
                Name: "payment_request",
                Description: []string{
                    "A bare-bones invoice for a payment within the Lightning Network.  With the",
                    "details of the invoice, the sender has all the data necessary to send a",
                    "payment to the recipient.",
                },
                Type: mkstring(),
            },
        },
    }
}
func mkinvoicesrpc_CancelInvoiceMsg() Type {
    return Type{
        Name: "invoicesrpc_CancelInvoiceMsg",
        Fields: []Field{
            {
                Name: "payment_hash",
                Description: []string{
                    "Hash corresponding to the (hold) invoice to cancel.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkinvoicesrpc_CancelInvoiceResp() Type {
    return Type{
        Name: "invoicesrpc_CancelInvoiceResp",
    }
}
func mkinvoicesrpc_SettleInvoiceMsg() Type {
    return Type{
        Name: "invoicesrpc_SettleInvoiceMsg",
        Fields: []Field{
            {
                Name: "preimage",
                Description: []string{
                    "Externally discovered pre-image that should be used to settle the hold",
                    "invoice.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mkinvoicesrpc_SettleInvoiceResp() Type {
    return Type{
        Name: "invoicesrpc_SettleInvoiceResp",
    }
}
func mkinvoicesrpc_SubscribeSingleInvoiceRequest() Type {
    return Type{
        Name: "invoicesrpc_SubscribeSingleInvoiceRequest",
        Fields: []Field{
            {
                Name: "r_hash",
                Description: []string{
                    "Hash corresponding to the (hold) invoice to subscribe to.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_ChangePasswordRequest() Type {
    return Type{
        Name: "lnrpc_ChangePasswordRequest",
        Fields: []Field{
            {
                Name: "current_passphrase",
                Description: []string{
                    "current_password should be the current valid passphrase used to unlock the daemon.",
                },
                Type: mkstring(),
            },
            {
                Name: "current_password_bin",
                Description: []string{
                    "Binary form of current_passphrase, if specified will override current_passphrase.",
                    "When using JSON, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "new_passphrase",
                Description: []string{
                    "new_passphrase should be the new passphrase that will be needed to unlock the",
                    "daemon.",
                },
                Type: mkstring(),
            },
            {
                Name: "new_passphrase_bin",
                Description: []string{
                    "Binary form of new_passphrase, if specified will override new_passphrase.",
                    "When using JSON, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "wallet_name",
                Description: []string{
                    "wallet_name is optional, if specified will override default wallet.db",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_ChangePasswordResponse() Type {
    return Type{
        Name: "lnrpc_ChangePasswordResponse",
    }
}
func mklnrpc_CheckPasswordRequest() Type {
    return Type{
        Name: "lnrpc_CheckPasswordRequest",
        Fields: []Field{
            {
                Name: "wallet_passphrase",
                Description: []string{
                    "current_password should be the current valid passphrase used to unlock the daemon.",
                },
                Type: mkstring(),
            },
            {
                Name: "wallet_password_bin",
                Description: []string{
                    "Binary form of current_passphrase, if specified will override current_passphrase.",
                    "When using JSON, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "wallet_name",
                Description: []string{
                    "wallet_name is optional, if specified will override default wallet.db",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_CheckPasswordResponse() Type {
    return Type{
        Name: "lnrpc_CheckPasswordResponse",
        Fields: []Field{
            {
                Name: "valid_passphrase",
                Type: mkbool(),
            },
        },
    }
}
func mklnrpc_CrashRequest() Type {
    return Type{
        Name: "lnrpc_CrashRequest",
    }
}
func mklnrpc_CrashResponse() Type {
    return Type{
        Name: "lnrpc_CrashResponse",
    }
}
func mklnrpc_GetInfo2Request() Type {
    return Type{
        Name: "lnrpc_GetInfo2Request",
    }
}
func mklnrpc_GetInfo2Response() Type {
    return Type{
        Name: "lnrpc_GetInfo2Response",
        Fields: []Field{
            {
                Name: "neutrino",
                Type: mklnrpc_NeutrinoInfo(),
            },
            {
                Name: "wallet",
                Type: mklnrpc_WalletInfo(),
            },
            {
                Name: "lightning",
                Type: mklnrpc_GetInfoResponse(),
            },
        },
    }
}
func mklnrpc_GenSeedRequest() Type {
    return Type{
        Name: "lnrpc_GenSeedRequest",
        Fields: []Field{
            {
                Name: "seed_passphrase",
                Description: []string{
                    "seed_passphrase is the optional user specified passphrase that will be used",
                    "to encrypt the generated seed.",
                },
                Type: mkstring(),
            },
            {
                Name: "seed_passphrase_bin",
                Description: []string{
                    "seed_passphrase_bin overrides seed_passphrase if specified, for binary",
                    "representation of the passphrase. If using JSON then this field must be base64",
                    "encoded.",
                },
                Type: mkbytes(),
            },
            {
                Name: "seed_entropy",
                Description: []string{
                    "seed_entropy is an optional 16-bytes generated via CSPRNG. If not",
                    "specified, then a fresh set of randomness will be used to create the seed.",
                    "When using REST, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
        },
    }
}
func mklnrpc_GenSeedResponse() Type {
    return Type{
        Name: "lnrpc_GenSeedResponse",
        Fields: []Field{
            {
                Name: "seed",
                Description: []string{
                    "seed is a 15-word mnemonic that encodes a secret root seed used to generate",
                    "all private keys of the wallet.",
                },
                Repeated: true,
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_InitWalletRequest() Type {
    return Type{
        Name: "lnrpc_InitWalletRequest",
        Fields: []Field{
            {
                Name: "wallet_passphrase",
                Description: []string{
                    "wallet_passphrase is the passphrase that should be used to encrypt the",
                    "wallet. This MUST be at least 8 chars in length. After creation, this",
                    "password is required to unlock the daemon.",
                },
                Type: mkstring(),
            },
            {
                Name: "wallet_passphrase_bin",
                Description: []string{
                    "If specified, will override wallet_passphrase, but is expressed in binary.",
                    "When using REST, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "wallet_seed",
                Description: []string{
                    "wallet_seed is a 15-word wallet seed. This may have been generated by the",
                    "GenSeed method, or be an existing seed.",
                },
                Repeated: true,
                Type: mkstring(),
            },
            {
                Name: "seed_passphrase",
                Description: []string{
                    "seed_passphrase is an optional user provided passphrase that will be used",
                    "to encrypt the generated seed.",
                },
                Type: mkstring(),
            },
            {
                Name: "seed_passphrase_bin",
                Description: []string{
                    "If specified, will override seed_passphrase, but is expressed in binary.",
                    "When using REST, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "recovery_window",
                Description: []string{
                    "recovery_window is an optional argument specifying the address lookahead",
                    "when restoring a wallet seed. The recovery window applies to each",
                    "individual branch of the BIP44 derivation paths. Supplying a recovery",
                    "window of zero indicates that no addresses should be recovered, such after",
                    "the first initialization of the wallet.",
                },
                Type: mkint32(),
            },
            {
                Name: "channel_backups",
                Description: []string{
                    "channel_backups is an optional argument that allows clients to recover the",
                    "settled funds within a set of channels. This should be populated if the",
                    "user was unable to close out all channels and sweep funds before partial or",
                    "total data loss occurred. If specified, then after on-chain recovery of",
                    "funds, lnd begin to carry out the data loss recovery protocol in order to",
                    "recover the funds in each channel from a remote force closed transaction.",
                },
                Type: mklnrpc_ChanBackupSnapshot(),
            },
            {
                Name: "wallet_name",
                Description: []string{
                    "wallet_name is an optional argument that allows to define the ",
                    "wallet filename other than the default wallet.db",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_InitWalletResponse() Type {
    return Type{
        Name: "lnrpc_InitWalletResponse",
    }
}
func mklnrpc_UnlockWalletRequest() Type {
    return Type{
        Name: "lnrpc_UnlockWalletRequest",
        Fields: []Field{
            {
                Name: "wallet_passphrase",
                Description: []string{
                    "wallet_passphrase should be the current valid private passphrase for the daemon. This",
                    "will be required to decrypt on-disk material that the daemon requires to",
                    "function properly.",
                },
                Type: mkstring(),
            },
            {
                Name: "wallet_passphrase_bin",
                Description: []string{
                    "If specified, will override wallet_passphrase, but is expressed in binary.",
                    "When using REST, this field must be encoded as base64.",
                },
                Type: mkbytes(),
            },
            {
                Name: "recovery_window",
                Description: []string{
                    "recovery_window is an optional argument specifying the address lookahead",
                    "when restoring a wallet seed. The recovery window applies to each",
                    "individual branch of the BIP44 derivation paths. Supplying a recovery",
                    "window of zero indicates that no addresses should be recovered, such after",
                    "the first initialization of the wallet.",
                },
                Type: mkint32(),
            },
            {
                Name: "channel_backups",
                Description: []string{
                    "channel_backups is an optional argument that allows clients to recover the",
                    "settled funds within a set of channels. This should be populated if the",
                    "user was unable to close out all channels and sweep funds before partial or",
                    "total data loss occurred. If specified, then after on-chain recovery of",
                    "funds, lnd begin to carry out the data loss recovery protocol in order to",
                    "recover the funds in each channel from a remote force closed transaction.",
                },
                Type: mklnrpc_ChanBackupSnapshot(),
            },
            {
                Name: "wallet_name",
                Description: []string{
                    "wallet_name is optional when the user wants to load a specified wallet other ",
                    "than the default wallet.db",
                },
                Type: mkstring(),
            },
        },
    }
}
func mklnrpc_UnlockWalletResponse() Type {
    return Type{
        Name: "lnrpc_UnlockWalletResponse",
    }
}
func Watchtower_GetInfo() Method {
    return Method{
        Name: "GetInfo",
        Service: "Watchtower",
        Description: []string{
            "lncli: tower info",
            "GetInfo returns general information concerning the companion watchtower",
            "including its public key and URIs where the server is currently",
            "listening for clients.",
        },
        Req: mkwatchtowerrpc_GetInfoRequest(),
        Res: mkwatchtowerrpc_GetInfoResponse(),
    }
}
func Versioner_GetVersion() Method {
    return Method{
        Name: "GetVersion",
        Service: "Versioner",
        Category: "Meta",
        ShortDescription: "Display pldctl and pld version info",
        Description: []string{
            "GetVersion returns the current version and build information of the running",
            "daemon.",
        },
        Req: mkverrpc_VersionRequest(),
        Res: mkverrpc_Version(),
    }
}
func ChainNotifier_RegisterConfirmationsNtfn() Method {
    return Method{
        Name: "RegisterConfirmationsNtfn",
        Service: "ChainNotifier",
        Description: []string{
            "RegisterConfirmationsNtfn is a synchronous response-streaming RPC that",
            "registers an intent for a client to be notified once a confirmation request",
            "has reached its required number of confirmations on-chain.",
            "A client can specify whether the confirmation request should be for a",
            "particular transaction by its hash or for an output script by specifying a",
            "zero hash.",
        },
        Req: mkchainrpc_ConfRequest(),
        Res: mkchainrpc_ConfEvent(),
    }
}
func ChainNotifier_RegisterSpendNtfn() Method {
    return Method{
        Name: "RegisterSpendNtfn",
        Service: "ChainNotifier",
        Description: []string{
            "RegisterSpendNtfn is a synchronous response-streaming RPC that registers an",
            "intent for a client to be notification once a spend request has been spent",
            "by a transaction that has confirmed on-chain.",
            "A client can specify whether the spend request should be for a particular",
            "outpoint  or for an output script by specifying a zero outpoint.",
        },
        Req: mkchainrpc_SpendRequest(),
        Res: mkchainrpc_SpendEvent(),
    }
}
func ChainNotifier_RegisterBlockEpochNtfn() Method {
    return Method{
        Name: "RegisterBlockEpochNtfn",
        Service: "ChainNotifier",
        Description: []string{
            "RegisterBlockEpochNtfn is a synchronous response-streaming RPC that",
            "registers an intent for a client to be notified of blocks in the chain. The",
            "stream will return a hash and height tuple of a block for each new/stale",
            "block in the chain. It is the client's responsibility to determine whether",
            "the tuple returned is for a new or stale block in the chain.",
            "A client can also request a historical backlog of blocks from a particular",
            "point. This allows clients to be idempotent by ensuring that they do not",
            "missing processing a single block within the chain.",
        },
        Req: mkchainrpc_BlockEpoch(),
        Res: mkchainrpc_BlockEpoch(),
    }
}
func WatchtowerClient_AddTower() Method {
    return Method{
        Name: "AddTower",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Register a watchtower to use for future sessions/backups",
        Description: []string{
            "AddTower adds a new watchtower reachable at the given address and",
            "considers it for new sessions. If the watchtower already exists, then",
            "any new addresses included will be considered when dialing it for",
            "session negotiations and backups.",
        },
        Req: mkwtclientrpc_AddTowerRequest(),
        Res: mkwtclientrpc_AddTowerResponse(),
    }
}
func WatchtowerClient_RemoveTower() Method {
    return Method{
        Name: "RemoveTower",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Remove a watchtower to prevent its use for future sessions/backups",
        Description: []string{
            "RemoveTower removes a watchtower from being considered for future session",
            "negotiations and from being used for any subsequent backups until it's added",
            "again. If an address is provided, then this RPC only serves as a way of",
            "removing the address from the watchtower instead.",
        },
        Req: mkwtclientrpc_RemoveTowerRequest(),
        Res: mkwtclientrpc_RemoveTowerResponse(),
    }
}
func WatchtowerClient_ListTowers() Method {
    return Method{
        Name: "ListTowers",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Display information about all registered watchtowers",
        Description: []string{
            "ListTowers returns the list of watchtowers registered with the client.",
        },
        Req: mkwtclientrpc_ListTowersRequest(),
        Res: mkwtclientrpc_ListTowersResponse(),
    }
}
func WatchtowerClient_GetTowerInfo() Method {
    return Method{
        Name: "GetTowerInfo",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Display information about a specific registered watchtower",
        Description: []string{
            "GetTowerInfo retrieves information for a registered watchtower.",
        },
        Req: mkwtclientrpc_GetTowerInfoRequest(),
        Res: mkwtclientrpc_Tower(),
    }
}
func WatchtowerClient_Stats() Method {
    return Method{
        Name: "Stats",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Display the session stats of the watchtower client",
        Description: []string{
            "Stats returns the in-memory statistics of the client since startup.",
        },
        Req: mkwtclientrpc_StatsRequest(),
        Res: mkwtclientrpc_StatsResponse(),
    }
}
func WatchtowerClient_Policy() Method {
    return Method{
        Name: "Policy",
        Service: "WatchtowerClient",
        Category: "Watchtower",
        ShortDescription: "Display the active watchtower client policy configuration",
        Description: []string{
            "Policy returns the active watchtower client policy configuration.",
        },
        Req: mkwtclientrpc_PolicyRequest(),
        Res: mkwtclientrpc_PolicyResponse(),
    }
}
func Lightning_WalletBalance() Method {
    return Method{
        Name: "WalletBalance",
        Service: "Lightning",
        Category: "Wallet",
        ShortDescription: "Compute and display the wallet's current balance",
        Description: []string{
            "WalletBalance returns total unspent outputs(confirmed and unconfirmed), all",
            "confirmed unspent outputs and all unconfirmed unspent outputs under control",
            "of the wallet.",
        },
        Req: mklnrpc_WalletBalanceRequest(),
        Res: mklnrpc_WalletBalanceResponse(),
    }
}
func Lightning_GetAddressBalances() Method {
    return Method{
        Name: "GetAddressBalances",
        Service: "Lightning",
        Category: "Address",
        ShortDescription: "Compute and display balances for each address in the wallet",
        Description: []string{
            "GetAddressBalances returns the balance for each of the addresses in the wallet.",
        },
        Req: mklnrpc_GetAddressBalancesRequest(),
        Res: mklnrpc_GetAddressBalancesResponse(),
    }
}
func Lightning_ChannelBalance() Method {
    return Method{
        Name: "ChannelBalance",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Returns the sum of the total available channel balance across all open channels",
        Description: []string{
            "ChannelBalance returns a report on the total funds across all open channels,",
            "categorized in local/remote, pending local/remote and unsettled local/remote",
            "balances.",
        },
        Req: mklnrpc_ChannelBalanceRequest(),
        Res: mklnrpc_ChannelBalanceResponse(),
    }
}
func Lightning_GetTransactions() Method {
    return Method{
        Name: "GetTransactions",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "List transactions from the wallet",
        Description: []string{
            "GetTransactions returns a list describing all the known transactions",
            "relevant to the wallet.",
        },
        Req: mklnrpc_GetTransactionsRequest(),
        Res: mklnrpc_TransactionDetails(),
    }
}
func Lightning_EstimateFee() Method {
    return Method{
        Name: "EstimateFee",
        Service: "Lightning",
        Category: "Neutrino",
        ShortDescription: "Get fee estimates for sending bitcoin on-chain to multiple addresses",
        Description: []string{
            "EstimateFee asks the chain backend to estimate the fee rate and total fees",
            "for a transaction that pays to multiple specified outputs.",
            "When using REST, the `AddrToAmount` map type can be set by appending",
            "`&AddrToAmount[<address>]=<amount_to_send>` to the URL. Unfortunately this",
            "map type doesn't appear in the REST API documentation because of a bug in",
            "the grpc-gateway library.",
        },
        Req: mklnrpc_EstimateFeeRequest(),
        Res: mklnrpc_EstimateFeeResponse(),
    }
}
func Lightning_SendCoins() Method {
    return Method{
        Name: "SendCoins",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "Send bitcoin on-chain to an address",
        Description: []string{
            "SendCoins executes a request to send coins to a particular address. Unlike",
            "SendMany, this RPC call only allows creating a single output at a time. If",
            "neither target_conf, or sat_per_byte are set, then the internal wallet will",
            "consult its fee model to determine a fee for the default confirmation",
            "target.",
        },
        Req: mklnrpc_SendCoinsRequest(),
        Res: mklnrpc_SendCoinsResponse(),
    }
}
func Lightning_ListUnspent() Method {
    return Method{
        Name: "ListUnspent",
        Service: "Lightning",
        Description: []string{
            "lncli: `listunspent`",
            "Deprecated, use walletrpc.ListUnspent instead.",
            "ListUnspent returns a list of all utxos spendable by the wallet with a",
            "number of confirmations between the specified minimum and maximum.",
        },
        Req: mklnrpc_ListUnspentRequest(),
        Res: mklnrpc_ListUnspentResponse(),
    }
}
func Lightning_SubscribeTransactions() Method {
    return Method{
        Name: "SubscribeTransactions",
        Service: "Lightning",
        Description: []string{
            "SubscribeTransactions creates a uni-directional stream from the server to",
            "the client in which any newly discovered transactions relevant to the",
            "wallet are sent over.",
        },
        Req: mklnrpc_GetTransactionsRequest(),
        Res: mklnrpc_Transaction(),
    }
}
func Lightning_SendMany() Method {
    return Method{
        Name: "SendMany",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "Send bitcoin on-chain to multiple addresses",
        Description: []string{
            "SendMany handles a request for a transaction that creates multiple specified",
            "outputs in parallel. If neither target_conf, or sat_per_byte are set, then",
            "the internal wallet will consult its fee model to determine a fee for the",
            "default confirmation target.",
        },
        Req: mklnrpc_SendManyRequest(),
        Res: mklnrpc_SendManyResponse(),
    }
}
func Lightning_NewAddress() Method {
    return Method{
        Name: "NewAddress",
        Service: "Lightning",
        Description: []string{
            "NewAddress creates a new address under control of the local wallet.",
        },
        Req: mklnrpc_NewAddressRequest(),
        Res: mklnrpc_NewAddressResponse(),
    }
}
func Lightning_SignMessage() Method {
    return Method{
        Name: "SignMessage",
        Service: "Lightning",
        Category: "Address",
        ShortDescription: "Signs a message using the private key of a payment address",
        Description: []string{
            "SignMessage signs a message with this node's private key. The returned",
            "signature string is `zbase32` encoded and pubkey recoverable, meaning that",
            "only the message digest and signature are needed for verification.",
        },
        Req: mklnrpc_SignMessageRequest(),
        Res: mklnrpc_SignMessageResponse(),
    }
}
func Lightning_ConnectPeer() Method {
    return Method{
        Name: "ConnectPeer",
        Service: "Lightning",
        Category: "Peer",
        ShortDescription: "Connect to a remote pld peer",
        Description: []string{
            "ConnectPeer attempts to establish a connection to a remote peer. This is at",
            "the networking level, and is used for communication between nodes. This is",
            "distinct from establishing a channel with a peer.",
        },
        Req: mklnrpc_ConnectPeerRequest(),
        Res: mklnrpc_ConnectPeerResponse(),
    }
}
func Lightning_DisconnectPeer() Method {
    return Method{
        Name: "DisconnectPeer",
        Service: "Lightning",
        Category: "Peer",
        ShortDescription: "Disconnect a remote pld peer identified by public key",
        Description: []string{
            "DisconnectPeer attempts to disconnect one peer from another identified by a",
            "given pubKey. In the case that we currently have a pending or active channel",
            "with the target peer, then this action will be not be allowed.",
        },
        Req: mklnrpc_DisconnectPeerRequest(),
        Res: mklnrpc_DisconnectPeerResponse(),
    }
}
func Lightning_ListPeers() Method {
    return Method{
        Name: "ListPeers",
        Service: "Lightning",
        Category: "Peer",
        ShortDescription: "List all active, currently connected peers",
        Description: []string{
            "ListPeers returns a verbose listing of all currently active peers.",
        },
        Req: mklnrpc_ListPeersRequest(),
        Res: mklnrpc_ListPeersResponse(),
    }
}
func Lightning_SubscribePeerEvents() Method {
    return Method{
        Name: "SubscribePeerEvents",
        Service: "Lightning",
        Description: []string{
            "SubscribePeerEvents creates a uni-directional stream from the server to",
            "the client in which any events relevant to the state of peers are sent",
            "over. Events include peers going online and offline.",
        },
        Req: mklnrpc_PeerEventSubscription(),
        Res: mklnrpc_PeerEvent(),
    }
}
func Lightning_GetInfo() Method {
    return Method{
        Name: "GetInfo",
        Service: "Lightning",
        Description: []string{
            "GetInfo returns general information concerning the lightning node including",
            "it's identity pubkey, alias, the chains it is connected to, and information",
            "concerning the number of open+pending channels.",
        },
        Req: mklnrpc_GetInfoRequest(),
        Res: mklnrpc_GetInfoResponse(),
    }
}
func Lightning_GetRecoveryInfo() Method {
    return Method{
        Name: "GetRecoveryInfo",
        Service: "Lightning",
        Description: []string{
            "lncli: `getrecoveryinfo`",
            "GetRecoveryInfo returns information concerning the recovery mode including",
            "whether it's in a recovery mode, whether the recovery is finished, and the",
            "progress made so far.",
        },
        Req: mklnrpc_GetRecoveryInfoRequest(),
        Res: mklnrpc_GetRecoveryInfoResponse(),
    }
}
func Lightning_PendingChannels() Method {
    return Method{
        Name: "PendingChannels",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Display information pertaining to pending channels",
        Description: []string{
            "PendingChannels returns a list of all the channels that are currently",
            "considered \"pending\". A channel is pending if it has finished the funding",
            "workflow and is waiting for confirmations for the funding txn, or is in the",
            "process of closure, either initiated cooperatively or non-cooperatively.",
        },
        Req: mklnrpc_PendingChannelsRequest(),
        Res: mklnrpc_PendingChannelsResponse(),
    }
}
func Lightning_ListChannels() Method {
    return Method{
        Name: "ListChannels",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "List all open channels",
        Description: []string{
            "ListChannels returns a description of all the open channels that this node",
            "is a participant in.",
        },
        Req: mklnrpc_ListChannelsRequest(),
        Res: mklnrpc_ListChannelsResponse(),
    }
}
func Lightning_SubscribeChannelEvents() Method {
    return Method{
        Name: "SubscribeChannelEvents",
        Service: "Lightning",
        Description: []string{
            "SubscribeChannelEvents creates a uni-directional stream from the server to",
            "the client in which any updates relevant to the state of the channels are",
            "sent over. Events include new active channels, inactive channels, and closed",
            "channels.",
        },
        Req: mklnrpc_ChannelEventSubscription(),
        Res: mklnrpc_ChannelEventUpdate(),
    }
}
func Lightning_ClosedChannels() Method {
    return Method{
        Name: "ClosedChannels",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "List all closed channels",
        Description: []string{
            "ClosedChannels returns a description of all the closed channels that",
            "this node was a participant in.",
        },
        Req: mklnrpc_ClosedChannelsRequest(),
        Res: mklnrpc_ClosedChannelsResponse(),
    }
}
func Lightning_OpenChannelSync() Method {
    return Method{
        Name: "OpenChannelSync",
        Service: "Lightning",
        Description: []string{
            "OpenChannelSync is a synchronous version of the OpenChannel RPC call. This",
            "call is meant to be consumed by clients to the REST proxy. As with all",
            "other sync calls, all byte slices are intended to be populated as hex",
            "encoded strings.",
        },
        Req: mklnrpc_OpenChannelRequest(),
        Res: mklnrpc_ChannelPoint(),
    }
}
func Lightning_OpenChannel() Method {
    return Method{
        Name: "OpenChannel",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Open a channel to a node or an existing peer",
        Description: []string{
            "OpenChannel attempts to open a singly funded channel specified in the",
            "request to a remote peer. Users are able to specify a target number of",
            "blocks that the funding transaction should be confirmed in, or a manual fee",
            "rate to us for the funding transaction. If neither are specified, then a",
            "lax block confirmation target is used. Each OpenStatusUpdate will return",
            "the pending channel ID of the in-progress channel. Depending on the",
            "arguments specified in the OpenChannelRequest, this pending channel ID can",
            "then be used to manually progress the channel funding flow.",
        },
        Req: mklnrpc_OpenChannelRequest(),
        Res: mklnrpc_OpenStatusUpdate(),
    }
}
func Lightning_FundingStateStep() Method {
    return Method{
        Name: "FundingStateStep",
        Service: "Lightning",
        Description: []string{
            "FundingStateStep is an advanced funding related call that allows the caller",
            "to either execute some preparatory steps for a funding workflow, or",
            "manually progress a funding workflow. The primary way a funding flow is",
            "identified is via its pending channel ID. As an example, this method can be",
            "used to specify that we're expecting a funding flow for a particular",
            "pending channel ID, for which we need to use specific parameters.",
            "Alternatively, this can be used to interactively drive PSBT signing for",
            "funding for partially complete funding transactions.",
        },
        Req: mklnrpc_FundingTransitionMsg(),
        Res: mklnrpc_FundingStateStepResp(),
    }
}
func Lightning_ChannelAcceptor() Method {
    return Method{
        Name: "ChannelAcceptor",
        Service: "Lightning",
        Description: []string{
            "ChannelAcceptor dispatches a bi-directional streaming RPC in which",
            "OpenChannel requests are sent to the client and the client responds with",
            "a boolean that tells LND whether or not to accept the channel. This allows",
            "node operators to specify their own criteria for accepting inbound channels",
            "through a single persistent connection.",
        },
        Req: mklnrpc_ChannelAcceptResponse(),
        Res: mklnrpc_ChannelAcceptRequest(),
    }
}
func Lightning_CloseChannel() Method {
    return Method{
        Name: "CloseChannel",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Close an existing channel",
        Description: []string{
            "CloseChannel attempts to close an active channel identified by its channel",
            "outpoint (ChannelPoint). The actions of this method can additionally be",
            "augmented to attempt a force close after a timeout period in the case of an",
            "inactive peer. If a non-force close (cooperative closure) is requested,",
            "then the user can specify either a target number of blocks until the",
            "closure transaction is confirmed, or a manual fee rate. If neither are",
            "specified, then a default lax, block confirmation target is used.",
        },
        Req: mklnrpc_CloseChannelRequest(),
        Res: mklnrpc_CloseStatusUpdate(),
    }
}
func Lightning_AbandonChannel() Method {
    return Method{
        Name: "AbandonChannel",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Abandons an existing channel",
        Description: []string{
            "AbandonChannel removes all channel state from the database except for a",
            "close summary. This method can be used to get rid of permanently unusable",
            "channels due to bugs fixed in newer versions of lnd. This method can also be",
            "used to remove externally funded channels where the funding transaction was",
            "never broadcast. Only available for non-externally funded channels in dev",
            "build.",
        },
        Req: mklnrpc_AbandonChannelRequest(),
        Res: mklnrpc_AbandonChannelResponse(),
    }
}
func Lightning_SendPayment() Method {
    return Method{
        Name: "SendPayment",
        Service: "Lightning",
        Description: []string{
            "Deprecated, use routerrpc.SendPaymentV2. SendPayment dispatches a",
            "bi-directional streaming RPC for sending payments through the Lightning",
            "Network. A single RPC invocation creates a persistent bi-directional",
            "stream allowing clients to rapidly send payments through the Lightning",
            "Network with a single persistent connection.",
        },
        Req: mklnrpc_SendRequest(),
        Res: mklnrpc_SendResponse(),
    }
}
func Lightning_SendPaymentSync() Method {
    return Method{
        Name: "SendPaymentSync",
        Service: "Lightning",
        Description: []string{
            "SendPaymentSync is the synchronous non-streaming version of SendPayment.",
            "This RPC is intended to be consumed by clients of the REST proxy.",
            "Additionally, this RPC expects the destination's public key and the payment",
            "hash (if any) to be encoded as hex strings.",
        },
        Req: mklnrpc_SendRequest(),
        Res: mklnrpc_SendResponse(),
    }
}
func Lightning_SendToRoute() Method {
    return Method{
        Name: "SendToRoute",
        Service: "Lightning",
        Description: []string{
            "lncli: `sendtoroute`",
            "Deprecated, use routerrpc.SendToRouteV2. SendToRoute is a bi-directional",
            "streaming RPC for sending payment through the Lightning Network. This",
            "method differs from SendPayment in that it allows users to specify a full",
            "route manually. This can be used for things like rebalancing, and atomic",
            "swaps.",
        },
        Req: mklnrpc_SendToRouteRequest(),
        Res: mklnrpc_SendResponse(),
    }
}
func Lightning_SendToRouteSync() Method {
    return Method{
        Name: "SendToRouteSync",
        Service: "Lightning",
        Description: []string{
            "SendToRouteSync is a synchronous version of SendToRoute. It Will block",
            "until the payment either fails or succeeds.",
        },
        Req: mklnrpc_SendToRouteRequest(),
        Res: mklnrpc_SendResponse(),
    }
}
func Lightning_AddInvoice() Method {
    return Method{
        Name: "AddInvoice",
        Service: "Lightning",
        Category: "Invoice",
        ShortDescription: "Add a new invoice",
        Description: []string{
            "AddInvoice attempts to add a new invoice to the invoice database. Any",
            "duplicated invoices are rejected, therefore all invoices *must* have a",
            "unique payment preimage.",
        },
        Req: mklnrpc_Invoice(),
        Res: mklnrpc_AddInvoiceResponse(),
    }
}
func Lightning_ListInvoices() Method {
    return Method{
        Name: "ListInvoices",
        Service: "Lightning",
        Category: "Invoice",
        ShortDescription: "List all invoices currently stored within the database. Any active debug invoices are ignored",
        Description: []string{
            "ListInvoices returns a list of all the invoices currently stored within the",
            "database. Any active debug invoices are ignored. It has full support for",
            "paginated responses, allowing users to query for specific invoices through",
            "their add_index. This can be done by using either the first_index_offset or",
            "last_index_offset fields included in the response as the index_offset of the",
            "next request. By default, the first 100 invoices created will be returned.",
            "Backwards pagination is also supported through the Reversed flag.",
        },
        Req: mklnrpc_ListInvoiceRequest(),
        Res: mklnrpc_ListInvoiceResponse(),
    }
}
func Lightning_LookupInvoice() Method {
    return Method{
        Name: "LookupInvoice",
        Service: "Lightning",
        Category: "Invoice",
        ShortDescription: "Lookup an existing invoice by its payment hash",
        Description: []string{
            "LookupInvoice attempts to look up an invoice according to its payment hash.",
            "The passed payment hash *must* be exactly 32 bytes, if not, an error is",
            "returned.",
        },
        Req: mklnrpc_PaymentHash(),
        Res: mklnrpc_Invoice(),
    }
}
func Lightning_SubscribeInvoices() Method {
    return Method{
        Name: "SubscribeInvoices",
        Service: "Lightning",
        Description: []string{
            "SubscribeInvoices returns a uni-directional stream (server -> client) for",
            "notifying the client of newly added/settled invoices. The caller can",
            "optionally specify the add_index and/or the settle_index. If the add_index",
            "is specified, then we'll first start by sending add invoice events for all",
            "invoices with an add_index greater than the specified value. If the",
            "settle_index is specified, the next, we'll send out all settle events for",
            "invoices with a settle_index greater than the specified value. One or both",
            "of these fields can be set. If no fields are set, then we'll only send out",
            "the latest add/settle events.",
        },
        Req: mklnrpc_InvoiceSubscription(),
        Res: mklnrpc_Invoice(),
    }
}
func Lightning_DecodePayReq() Method {
    return Method{
        Name: "DecodePayReq",
        Service: "Lightning",
        Category: "Invoice",
        ShortDescription: "Decode a payment request",
        Description: []string{
            "DecodePayReq takes an encoded payment request string and attempts to decode",
            "it, returning a full description of the conditions encoded within the",
            "payment request.",
        },
        Req: mklnrpc_PayReqString(),
        Res: mklnrpc_PayReq(),
    }
}
func Lightning_ListPayments() Method {
    return Method{
        Name: "ListPayments",
        Service: "Lightning",
        Category: "Payment",
        ShortDescription: "List all outgoing payments",
        Description: []string{
            "ListPayments returns a list of all outgoing payments.",
        },
        Req: mklnrpc_ListPaymentsRequest(),
        Res: mklnrpc_ListPaymentsResponse(),
    }
}
func Lightning_DeleteAllPayments() Method {
    return Method{
        Name: "DeleteAllPayments",
        Service: "Lightning",
        Description: []string{
            "DeleteAllPayments deletes all outgoing payments from DB.",
        },
        Req: mklnrpc_DeleteAllPaymentsRequest(),
        Res: mklnrpc_DeleteAllPaymentsResponse(),
    }
}
func Lightning_DescribeGraph() Method {
    return Method{
        Name: "DescribeGraph",
        Service: "Lightning",
        Category: "Graph",
        ShortDescription: "Describe the network graph",
        Description: []string{
            "DescribeGraph returns a description of the latest graph state from the",
            "point of view of the node. The graph information is partitioned into two",
            "components: all the nodes/vertexes, and all the edges that connect the",
            "vertexes themselves. As this is a directed graph, the edges also contain",
            "the node directional specific routing policy which includes: the time lock",
            "delta, fee information, etc.",
        },
        Req: mklnrpc_ChannelGraphRequest(),
        Res: mklnrpc_ChannelGraph(),
    }
}
func Lightning_GetNodeMetrics() Method {
    return Method{
        Name: "GetNodeMetrics",
        Service: "Lightning",
        Category: "Graph",
        ShortDescription: "Get node metrics",
        Description: []string{
            "GetNodeMetrics returns node metrics calculated from the graph. Currently",
            "the only supported metric is betweenness centrality of individual nodes.",
        },
        Req: mklnrpc_NodeMetricsRequest(),
        Res: mklnrpc_NodeMetricsResponse(),
    }
}
func Lightning_GetChanInfo() Method {
    return Method{
        Name: "GetChanInfo",
        Service: "Lightning",
        Category: "Graph",
        ShortDescription: "Get the state of a channel",
        Description: []string{
            "GetChanInfo returns the latest authenticated network announcement for the",
            "given channel identified by its channel ID: an 8-byte integer which",
            "uniquely identifies the location of transaction's funding output within the",
            "blockchain.",
        },
        Req: mklnrpc_ChanInfoRequest(),
        Res: mklnrpc_ChannelEdge(),
    }
}
func Lightning_GetNodeInfo() Method {
    return Method{
        Name: "GetNodeInfo",
        Service: "Lightning",
        Category: "Graph",
        ShortDescription: "Get information on a specific node",
        Description: []string{
            "GetNodeInfo returns the latest advertised, aggregated, and authenticated",
            "channel information for the specified node identified by its public key.",
        },
        Req: mklnrpc_NodeInfoRequest(),
        Res: mklnrpc_NodeInfo(),
    }
}
func Lightning_QueryRoutes() Method {
    return Method{
        Name: "QueryRoutes",
        Service: "Lightning",
        Category: "Payment",
        ShortDescription: "Query a route to a destination",
        Description: []string{
            "QueryRoutes attempts to query the daemon's Channel Router for a possible",
            "route to a target destination capable of carrying a specific amount of",
            "satoshis. The returned route contains the full details required to craft and",
            "send an HTLC, also including the necessary information that should be",
            "present within the Sphinx packet encapsulated within the HTLC.",
            "When using REST, the `dest_custom_records` map type can be set by appending",
            "`&dest_custom_records[<record_number>]=<record_data_base64_url_encoded>`",
            "to the URL. Unfortunately this map type doesn't appear in the REST API",
            "documentation because of a bug in the grpc-gateway library.",
        },
        Req: mklnrpc_QueryRoutesRequest(),
        Res: mklnrpc_QueryRoutesResponse(),
    }
}
func Lightning_GetNetworkInfo() Method {
    return Method{
        Name: "GetNetworkInfo",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Get statistical information about the current state of the network",
        Description: []string{
            "GetNetworkInfo returns some basic stats about the known channel graph from",
            "the point of view of the node.",
        },
        Req: mklnrpc_NetworkInfoRequest(),
        Res: mklnrpc_NetworkInfo(),
    }
}
func Lightning_StopDaemon() Method {
    return Method{
        Name: "StopDaemon",
        Service: "Lightning",
        Category: "Meta",
        ShortDescription: "Stop and shutdown the daemon",
        Description: []string{
            "StopDaemon will send a shutdown request to the interrupt handler, triggering",
            "a graceful shutdown of the daemon.",
        },
        Req: mklnrpc_StopRequest(),
        Res: mklnrpc_StopResponse(),
    }
}
func Lightning_SubscribeChannelGraph() Method {
    return Method{
        Name: "SubscribeChannelGraph",
        Service: "Lightning",
        Description: []string{
            "SubscribeChannelGraph launches a streaming RPC that allows the caller to",
            "receive notifications upon any changes to the channel graph topology from",
            "the point of view of the responding node. Events notified include: new",
            "nodes coming online, nodes updating their authenticated attributes, new",
            "channels being advertised, updates in the routing policy for a directional",
            "channel edge, and when channels are closed on-chain.",
        },
        Req: mklnrpc_GraphTopologySubscription(),
        Res: mklnrpc_GraphTopologyUpdate(),
    }
}
func Lightning_DebugLevel() Method {
    return Method{
        Name: "DebugLevel",
        Service: "Lightning",
        Category: "Meta",
        ShortDescription: "Set the debug level",
        Description: []string{
            "DebugLevel allows a caller to programmatically set the logging verbosity of",
            "lnd. The logging can be targeted according to a coarse daemon-wide logging",
            "level, or in a granular fashion to specify the logging for a target",
            "sub-system.",
        },
        Req: mklnrpc_DebugLevelRequest(),
        Res: mklnrpc_DebugLevelResponse(),
    }
}
func Lightning_FeeReport() Method {
    return Method{
        Name: "FeeReport",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Display the current fee policies of all active channels",
        Description: []string{
            "FeeReport allows the caller to obtain a report detailing the current fee",
            "schedule enforced by the node globally for each channel.",
        },
        Req: mklnrpc_FeeReportRequest(),
        Res: mklnrpc_FeeReportResponse(),
    }
}
func Lightning_UpdateChannelPolicy() Method {
    return Method{
        Name: "UpdateChannelPolicy",
        Service: "Lightning",
        Category: "Channel",
        ShortDescription: "Update the channel policy for all channels, or a single channel",
        Description: []string{
            "UpdateChannelPolicy allows the caller to update the fee schedule and",
            "channel policies for all channels globally, or a particular channel.",
        },
        Req: mklnrpc_PolicyUpdateRequest(),
        Res: mklnrpc_PolicyUpdateResponse(),
    }
}
func Lightning_ForwardingHistory() Method {
    return Method{
        Name: "ForwardingHistory",
        Service: "Lightning",
        Category: "Payment",
        ShortDescription: "Query the history of all forwarded HTLCs",
        Description: []string{
            "ForwardingHistory allows the caller to query the htlcswitch for a record of",
            "all HTLCs forwarded within the target time range, and integer offset",
            "within that time range. If no time-range is specified, then the first chunk",
            "of the past 24 hrs of forwarding history are returned.",
            "A list of forwarding events are returned. The size of each forwarding event",
            "is 40 bytes, and the max message size able to be returned in gRPC is 4 MiB.",
            "As a result each message can only contain 50k entries. Each response has",
            "the index offset of the last entry. The index offset can be provided to the",
            "request to allow the caller to skip a series of records.",
        },
        Req: mklnrpc_ForwardingHistoryRequest(),
        Res: mklnrpc_ForwardingHistoryResponse(),
    }
}
func Lightning_ExportChannelBackup() Method {
    return Method{
        Name: "ExportChannelBackup",
        Service: "Lightning",
        Category: "Backup",
        ShortDescription: "Obtain a static channel back up for a selected channels, or all known channels",
        Description: []string{
            "ExportChannelBackup attempts to return an encrypted static channel backup",
            "for the target channel identified by it channel point. The backup is",
            "encrypted with a key generated from the aezeed seed of the user. The",
            "returned backup can either be restored using the RestoreChannelBackup",
            "method once lnd is running, or via the InitWallet and UnlockWallet methods",
            "from the WalletUnlocker service.",
        },
        Req: mklnrpc_ExportChannelBackupRequest(),
        Res: mklnrpc_ChannelBackup(),
    }
}
func Lightning_ExportAllChannelBackups() Method {
    return Method{
        Name: "ExportAllChannelBackups",
        Service: "Lightning",
        Description: []string{
            "ExportAllChannelBackups returns static channel backups for all existing",
            "channels known to lnd. A set of regular singular static channel backups for",
            "each channel are returned. Additionally, a multi-channel backup is returned",
            "as well, which contains a single encrypted blob containing the backups of",
            "each channel.",
        },
        Req: mklnrpc_ChanBackupExportRequest(),
        Res: mklnrpc_ChanBackupSnapshot(),
    }
}
func Lightning_VerifyChanBackup() Method {
    return Method{
        Name: "VerifyChanBackup",
        Service: "Lightning",
        Category: "Backup",
        ShortDescription: "Verify an existing channel backup",
        Description: []string{
            "VerifyChanBackup allows a caller to verify the integrity of a channel backup",
            "snapshot. This method will accept either a packed Single or a packed Multi.",
            "Specifying both will result in an error.",
        },
        Req: mklnrpc_ChanBackupSnapshot(),
        Res: mklnrpc_VerifyChanBackupResponse(),
    }
}
func Lightning_RestoreChannelBackups() Method {
    return Method{
        Name: "RestoreChannelBackups",
        Service: "Lightning",
        Category: "Backup",
        ShortDescription: "Restore an existing single or multi-channel static channel backup",
        Description: []string{
            "RestoreChannelBackups accepts a set of singular channel backups, or a",
            "single encrypted multi-chan backup and attempts to recover any funds",
            "remaining within the channel. If we are able to unpack the backup, then the",
            "new channel will be shown under listchannels, as well as pending channels.",
        },
        Req: mklnrpc_RestoreChanBackupRequest(),
        Res: mklnrpc_RestoreBackupResponse(),
    }
}
func Lightning_SubscribeChannelBackups() Method {
    return Method{
        Name: "SubscribeChannelBackups",
        Service: "Lightning",
        Description: []string{
            "SubscribeChannelBackups allows a client to sub-subscribe to the most up to",
            "date information concerning the state of all channel backups. Each time a",
            "new channel is added, we return the new set of channels, along with a",
            "multi-chan backup containing the backup info for all channels. Each time a",
            "channel is closed, we send a new update, which contains new new chan back",
            "ups, but the updated set of encrypted multi-chan backups with the closed",
            "channel(s) removed.",
        },
        Req: mklnrpc_ChannelBackupSubscription(),
        Res: mklnrpc_ChanBackupSnapshot(),
    }
}
func Lightning_ReSync() Method {
    return Method{
        Name: "ReSync",
        Service: "Lightning",
        Category: "Unspent",
        ShortDescription: "Scan over the chain to find any transactions which may not have been recorded in the wallet's database",
        Description: []string{
            "Scan over the chain to find any transactions which may not have been recorded in the wallet's database",
        },
        Req: mklnrpc_ReSyncChainRequest(),
        Res: mklnrpc_ReSyncChainResponse(),
    }
}
func Lightning_StopReSync() Method {
    return Method{
        Name: "StopReSync",
        Service: "Lightning",
        Category: "Unspent",
        ShortDescription: "Stop a re-synchronization job before it's completion",
        Description: []string{
            "Stop a re-synchronization job before it's completion",
        },
        Req: mklnrpc_StopReSyncRequest(),
        Res: mklnrpc_StopReSyncResponse(),
    }
}
func Lightning_GetWalletSeed() Method {
    return Method{
        Name: "GetWalletSeed",
        Service: "Lightning",
        Category: "Wallet",
        ShortDescription: "Get the wallet seed words for this wallet",
        Description: []string{
            "Get the wallet seed words for this wallet",
        },
        Req: mklnrpc_GetWalletSeedRequest(),
        Res: mklnrpc_GetWalletSeedResponse(),
    }
}
func Lightning_ChangeSeedPassphrase() Method {
    return Method{
        Name: "ChangeSeedPassphrase",
        Service: "Lightning",
        Category: "Seed",
        ShortDescription: "Alter the passphrase which is used to encrypt a wallet seed",
        Description: []string{
            "Change seed's passphrase",
        },
        Req: mklnrpc_ChangeSeedPassphraseRequest(),
        Res: mklnrpc_ChangeSeedPassphraseResponse(),
    }
}
func Lightning_GetSecret() Method {
    return Method{
        Name: "GetSecret",
        Service: "Lightning",
        Category: "Wallet",
        ShortDescription: "Get a secret seed",
        Description: []string{
            "Get a secret seed which is generated using the wallet's private key, this can be used as a password for another application",
        },
        Req: mklnrpc_GetSecretRequest(),
        Res: mklnrpc_GetSecretResponse(),
    }
}
func Lightning_ImportPrivKey() Method {
    return Method{
        Name: "ImportPrivKey",
        Service: "Lightning",
        Category: "Address",
        ShortDescription: "Imports a WIF-encoded private key to the 'imported' account",
        Description: []string{
            "Imports a WIF-encoded private key to the 'imported' account.",
        },
        Req: mklnrpc_ImportPrivKeyRequest(),
        Res: mklnrpc_ImportPrivKeyResponse(),
    }
}
func Lightning_DumpPrivKey() Method {
    return Method{
        Name: "DumpPrivKey",
        Service: "Lightning",
        Category: "Address",
        ShortDescription: "Returns the private key in WIF encoding that controls some wallet address",
        Description: []string{
            "Returns the private key in WIF encoding that controls some wallet address.",
        },
        Req: mklnrpc_DumpPrivKeyRequest(),
        Res: mklnrpc_DumpPrivKeyResponse(),
    }
}
func Lightning_ListLockUnspent() Method {
    return Method{
        Name: "ListLockUnspent",
        Service: "Lightning",
        Category: "Lock",
        ShortDescription: "Returns a JSON array of outpoints marked as locked (with lockunspent) for this wallet session",
        Description: []string{
            "Returns a JSON array of outpoints marked as locked (with lockunspent) for this wallet session.",
        },
        Req: mklnrpc_ListLockUnspentRequest(),
        Res: mklnrpc_ListLockUnspentResponse(),
    }
}
func Lightning_LockUnspent() Method {
    return Method{
        Name: "LockUnspent",
        Service: "Lightning",
        Category: "Lock",
        ShortDescription: "Locks or unlocks an unspent output",
        Description: []string{
            "Locks or unlocks an unspent output",
        },
        Req: mklnrpc_LockUnspentRequest(),
        Res: mklnrpc_LockUnspentResponse(),
    }
}
func Lightning_CreateTransaction() Method {
    return Method{
        Name: "CreateTransaction",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "Create a transaction but do not send it to the chain",
        Description: []string{
            "Create a transaction but po not send it to the chain",
        },
        Req: mklnrpc_CreateTransactionRequest(),
        Res: mklnrpc_CreateTransactionResponse(),
    }
}
func Lightning_GetNewAddress() Method {
    return Method{
        Name: "GetNewAddress",
        Service: "Lightning",
        Category: "Address",
        ShortDescription: "Generates a new address",
        Description: []string{
            "Generates and returns a new payment address",
        },
        Req: mklnrpc_GetNewAddressRequest(),
        Res: mklnrpc_GetNewAddressResponse(),
    }
}
func Lightning_GetTransaction() Method {
    return Method{
        Name: "GetTransaction",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "Returns a JSON object with details regarding a transaction relevant to this wallet",
        Description: []string{
            "Returns a JSON object with details regarding a transaction relevant to this wallet.",
        },
        Req: mklnrpc_GetTransactionRequest(),
        Res: mklnrpc_GetTransactionResponse(),
    }
}
func Lightning_SetNetworkStewardVote() Method {
    return Method{
        Name: "SetNetworkStewardVote",
        Service: "Lightning",
        Category: "Network Steward Vote",
        ShortDescription: "Configure the wallet to vote for a network steward when making payments (note: payments to segwit addresses cannot vote)",
        Description: []string{
            "Configure the wallet to vote for a network steward when making payments (note: payments to segwit addresses cannot vote)",
        },
        Req: mklnrpc_SetNetworkStewardVoteRequest(),
        Res: mklnrpc_SetNetworkStewardVoteResponse(),
    }
}
func Lightning_GetNetworkStewardVote() Method {
    return Method{
        Name: "GetNetworkStewardVote",
        Service: "Lightning",
        Category: "Network Steward Vote",
        ShortDescription: "Find out how the wallet is currently configured to vote in a network steward election",
        Description: []string{
            "Find out how the wallet is currently configured to vote in a network steward election",
        },
        Req: mklnrpc_GetNetworkStewardVoteRequest(),
        Res: mklnrpc_GetNetworkStewardVoteResponse(),
    }
}
func Lightning_BcastTransaction() Method {
    return Method{
        Name: "BcastTransaction",
        Service: "Lightning",
        Category: "Neutrino",
        ShortDescription: "Broadcast a transaction onchain",
        Description: []string{
            "Broadcast a transaction",
        },
        Req: mklnrpc_BcastTransactionRequest(),
        Res: mklnrpc_BcastTransactionResponse(),
    }
}
func Lightning_SendFrom() Method {
    return Method{
        Name: "SendFrom",
        Service: "Lightning",
        Category: "Transaction",
        ShortDescription: "Authors, signs, and sends a transaction that outputs some amount to a payment address",
        Description: []string{
            "SendFrom authors, signs, and sends a transaction that outputs some amount to a payment address.",
        },
        Req: mklnrpc_SendFromRequest(),
        Res: mklnrpc_SendFromResponse(),
    }
}
func Lightning_DecodeRawTransaction() Method {
    return Method{
        Name: "DecodeRawTransaction",
        Service: "Lightning",
        Category: "Util",
        ShortDescription: "Returns a JSON object representing the provided serialized, hex-encoded transaction.",
        Description: []string{
            "DecodeRawTransaction returns a JSON object representing the provided serialized, hex-encoded transaction.",
        },
        Req: mklnrpc_DecodeRawTransactionRequest(),
        Res: mklnrpc_DecodeRawTransactionResponse(),
    }
}
func Signer_SignOutputRaw() Method {
    return Method{
        Name: "SignOutputRaw",
        Service: "Signer",
        Description: []string{
            "SignOutputRaw is a method that can be used to generated a signature for a",
            "set of inputs/outputs to a transaction. Each request specifies details",
            "concerning how the outputs should be signed, which keys they should be",
            "signed with, and also any optional tweaks. The return value is a fixed",
            "64-byte signature (the same format as we use on the wire in Lightning).",
            "If we are  unable to sign using the specified keys, then an error will be",
            "returned.",
        },
        Req: mksignrpc_SignReq(),
        Res: mksignrpc_SignResp(),
    }
}
func Signer_ComputeInputScript() Method {
    return Method{
        Name: "ComputeInputScript",
        Service: "Signer",
        Description: []string{
            "ComputeInputScript generates a complete InputIndex for the passed",
            "transaction with the signature as defined within the passed SignDescriptor.",
            "This method should be capable of generating the proper input script for",
            "both regular p2wkh output and p2wkh outputs nested within a regular p2sh",
            "output.",
            "Note that when using this method to sign inputs belonging to the wallet,",
            "the only items of the SignDescriptor that need to be populated are pkScript",
            "in the TxOut field, the value in that same field, and finally the input",
            "index.",
        },
        Req: mksignrpc_SignReq(),
        Res: mksignrpc_InputScriptResp(),
    }
}
func Signer_SignMessage() Method {
    return Method{
        Name: "SignMessage",
        Service: "Signer",
        Description: []string{
            "SignMessage signs a message with the key specified in the key locator. The",
            "returned signature is fixed-size LN wire format encoded.",
            "The main difference to SignMessage in the main RPC is that a specific key is",
            "used to sign the message instead of the node identity private key.",
        },
        Req: mksignrpc_SignMessageReq(),
        Res: mksignrpc_SignMessageResp(),
    }
}
func Signer_VerifyMessage() Method {
    return Method{
        Name: "VerifyMessage",
        Service: "Signer",
        Description: []string{
            "VerifyMessage verifies a signature over a message using the public key",
            "provided. The signature must be fixed-size LN wire format encoded.",
            "The main difference to VerifyMessage in the main RPC is that the public key",
            "used to sign the message does not have to be a node known to the network.",
        },
        Req: mksignrpc_VerifyMessageReq(),
        Res: mksignrpc_VerifyMessageResp(),
    }
}
func Signer_DeriveSharedKey() Method {
    return Method{
        Name: "DeriveSharedKey",
        Service: "Signer",
        Description: []string{
            "DeriveSharedKey returns a shared secret key by performing Diffie-Hellman key",
            "derivation between the ephemeral public key in the request and the node's",
            "key specified in the key_desc parameter. Either a key locator or a raw",
            "public key is expected in the key_desc, if neither is supplied, defaults to",
            "the node's identity private key:",
            "P_shared = privKeyNode * ephemeralPubkey",
            "The resulting shared public key is serialized in the compressed format and",
            "hashed with sha256, resulting in the final key length of 256bit.",
        },
        Req: mksignrpc_SharedKeyRequest(),
        Res: mksignrpc_SharedKeyResponse(),
    }
}
func Autopilot_Status() Method {
    return Method{
        Name: "Status",
        Service: "Autopilot",
        Description: []string{
            "Status returns whether the daemon's autopilot agent is active.",
        },
        Req: mkautopilotrpc_StatusRequest(),
        Res: mkautopilotrpc_StatusResponse(),
    }
}
func Autopilot_ModifyStatus() Method {
    return Method{
        Name: "ModifyStatus",
        Service: "Autopilot",
        Description: []string{
            "ModifyStatus is used to modify the status of the autopilot agent, like",
            "enabling or disabling it.",
        },
        Req: mkautopilotrpc_ModifyStatusRequest(),
        Res: mkautopilotrpc_ModifyStatusResponse(),
    }
}
func Autopilot_QueryScores() Method {
    return Method{
        Name: "QueryScores",
        Service: "Autopilot",
        Description: []string{
            "QueryScores queries all available autopilot heuristics, in addition to any",
            "active combination of these heruristics, for the scores they would give to",
            "the given nodes.",
        },
        Req: mkautopilotrpc_QueryScoresRequest(),
        Res: mkautopilotrpc_QueryScoresResponse(),
    }
}
func Autopilot_SetScores() Method {
    return Method{
        Name: "SetScores",
        Service: "Autopilot",
        Description: []string{
            "SetScores attempts to set the scores used by the running autopilot agent,",
            "if the external scoring heuristic is enabled.",
        },
        Req: mkautopilotrpc_SetScoresRequest(),
        Res: mkautopilotrpc_SetScoresResponse(),
    }
}
func WalletKit_ListUnspent() Method {
    return Method{
        Name: "ListUnspent",
        Service: "WalletKit",
        Category: "Unspent",
        ShortDescription: "List utxos available for spending",
        Description: []string{
            "ListUnspent returns a list of all utxos spendable by the wallet with a",
            "number of confirmations between the specified minimum and maximum.",
        },
        Req: mkwalletrpc_ListUnspentRequest(),
        Res: mkwalletrpc_ListUnspentResponse(),
    }
}
func WalletKit_LeaseOutput() Method {
    return Method{
        Name: "LeaseOutput",
        Service: "WalletKit",
        Description: []string{
            "LeaseOutput locks an output to the given ID, preventing it from being",
            "available for any future coin selection attempts. The absolute time of the",
            "lock's expiration is returned. The expiration of the lock can be extended by",
            "successive invocations of this RPC. Outputs can be unlocked before their",
            "expiration through `ReleaseOutput`.",
        },
        Req: mkwalletrpc_LeaseOutputRequest(),
        Res: mkwalletrpc_LeaseOutputResponse(),
    }
}
func WalletKit_ReleaseOutput() Method {
    return Method{
        Name: "ReleaseOutput",
        Service: "WalletKit",
        Description: []string{
            "ReleaseOutput unlocks an output, allowing it to be available for coin",
            "selection if it remains unspent. The ID should match the one used to",
            "originally lock the output.",
        },
        Req: mkwalletrpc_ReleaseOutputRequest(),
        Res: mkwalletrpc_ReleaseOutputResponse(),
    }
}
func WalletKit_DeriveNextKey() Method {
    return Method{
        Name: "DeriveNextKey",
        Service: "WalletKit",
        Description: []string{
            "DeriveNextKey attempts to derive the *next* key within the key family",
            "(account in BIP43) specified. This method should return the next external",
            "child within this branch.",
        },
        Req: mkwalletrpc_KeyReq(),
        Res: mksignrpc_KeyDescriptor(),
    }
}
func WalletKit_DeriveKey() Method {
    return Method{
        Name: "DeriveKey",
        Service: "WalletKit",
        Description: []string{
            "DeriveKey attempts to derive an arbitrary key specified by the passed",
            "KeyLocator.",
        },
        Req: mksignrpc_KeyLocator(),
        Res: mksignrpc_KeyDescriptor(),
    }
}
func WalletKit_NextAddr() Method {
    return Method{
        Name: "NextAddr",
        Service: "WalletKit",
        Description: []string{
            "NextAddr returns the next unused address within the wallet.",
        },
        Req: mkwalletrpc_AddrRequest(),
        Res: mkwalletrpc_AddrResponse(),
    }
}
func WalletKit_PublishTransaction() Method {
    return Method{
        Name: "PublishTransaction",
        Service: "WalletKit",
        Description: []string{
            "PublishTransaction attempts to publish the passed transaction to the",
            "network. Once this returns without an error, the wallet will continually",
            "attempt to re-broadcast the transaction on start up, until it enters the",
            "chain.",
        },
        Req: mkwalletrpc_Transaction(),
        Res: mkwalletrpc_PublishResponse(),
    }
}
func WalletKit_SendOutputs() Method {
    return Method{
        Name: "SendOutputs",
        Service: "WalletKit",
        Description: []string{
            "SendOutputs is similar to the existing sendmany call in Bitcoind, and",
            "allows the caller to create a transaction that sends to several outputs at",
            "once. This is ideal when wanting to batch create a set of transactions.",
        },
        Req: mkwalletrpc_SendOutputsRequest(),
        Res: mkwalletrpc_SendOutputsResponse(),
    }
}
func WalletKit_EstimateFee() Method {
    return Method{
        Name: "EstimateFee",
        Service: "WalletKit",
        Description: []string{
            "EstimateFee attempts to query the internal fee estimator of the wallet to",
            "determine the fee (in sat/kw) to attach to a transaction in order to",
            "achieve the confirmation target.",
        },
        Req: mkwalletrpc_EstimateFeeRequest(),
        Res: mkwalletrpc_EstimateFeeResponse(),
    }
}
func WalletKit_PendingSweeps() Method {
    return Method{
        Name: "PendingSweeps",
        Service: "WalletKit",
        Description: []string{
            "PendingSweeps returns lists of on-chain outputs that lnd is currently",
            "attempting to sweep within its central batching engine. Outputs with similar",
            "fee rates are batched together in order to sweep them within a single",
            "transaction.",
            "NOTE: Some of the fields within PendingSweepsRequest are not guaranteed to",
            "remain supported. This is an advanced API that depends on the internals of",
            "the UtxoSweeper, so things may change.",
        },
        Req: mkwalletrpc_PendingSweepsRequest(),
        Res: mkwalletrpc_PendingSweepsResponse(),
    }
}
func WalletKit_BumpFee() Method {
    return Method{
        Name: "BumpFee",
        Service: "WalletKit",
        Description: []string{
            "BumpFee bumps the fee of an arbitrary input within a transaction. This RPC",
            "takes a different approach than bitcoind's bumpfee command. lnd has a",
            "central batching engine in which inputs with similar fee rates are batched",
            "together to save on transaction fees. Due to this, we cannot rely on",
            "bumping the fee on a specific transaction, since transactions can change at",
            "any point with the addition of new inputs. The list of inputs that",
            "currently exist within lnd's central batching engine can be retrieved",
            "through the PendingSweeps RPC.",
            "When bumping the fee of an input that currently exists within lnd's central",
            "batching engine, a higher fee transaction will be created that replaces the",
            "lower fee transaction through the Replace-By-Fee (RBF) policy. If it",
            "This RPC also serves useful when wanting to perform a Child-Pays-For-Parent",
            "(CPFP), where the child transaction pays for its parent's fee. This can be",
            "done by specifying an outpoint within the low fee transaction that is under",
            "the control of the wallet.",
            "The fee preference can be expressed either as a specific fee rate or a delta",
            "of blocks in which the output should be swept on-chain within. If a fee",
            "preference is not explicitly specified, then an error is returned.",
            "Note that this RPC currently doesn't perform any validation checks on the",
            "fee preference being provided. For now, the responsibility of ensuring that",
            "the new fee preference is sufficient is delegated to the user.",
        },
        Req: mkwalletrpc_BumpFeeRequest(),
        Res: mkwalletrpc_BumpFeeResponse(),
    }
}
func WalletKit_ListSweeps() Method {
    return Method{
        Name: "ListSweeps",
        Service: "WalletKit",
        Description: []string{
            "ListSweeps returns a list of the sweep transactions our node has produced.",
            "Note that these sweeps may not be confirmed yet, as we record sweeps on",
            "broadcast, not confirmation.",
        },
        Req: mkwalletrpc_ListSweepsRequest(),
        Res: mkwalletrpc_ListSweepsResponse(),
    }
}
func WalletKit_LabelTransaction() Method {
    return Method{
        Name: "LabelTransaction",
        Service: "WalletKit",
        Description: []string{
            "LabelTransaction adds a label to a transaction. If the transaction already",
            "has a label the call will fail unless the overwrite bool is set. This will",
            "overwrite the exiting transaction label. Labels must not be empty, and",
            "cannot exceed 500 characters.",
        },
        Req: mkwalletrpc_LabelTransactionRequest(),
        Res: mkwalletrpc_LabelTransactionResponse(),
    }
}
func WalletKit_FundPsbt() Method {
    return Method{
        Name: "FundPsbt",
        Service: "WalletKit",
        Description: []string{
            "FundPsbt creates a fully populated PSBT that contains enough inputs to fund",
            "the outputs specified in the template. There are two ways of specifying a",
            "template: Either by passing in a PSBT with at least one output declared or",
            "by passing in a raw TxTemplate message.",
            "If there are no inputs specified in the template, coin selection is",
            "performed automatically. If the template does contain any inputs, it is",
            "assumed that full coin selection happened externally and no additional",
            "inputs are added. If the specified inputs aren't enough to fund the outputs",
            "with the given fee rate, an error is returned.",
            "After either selecting or verifying the inputs, all input UTXOs are locked",
            "with an internal app ID.",
            "NOTE: If this method returns without an error, it is the caller's",
            "responsibility to either spend the locked UTXOs (by finalizing and then",
            "publishing the transaction) or to unlock/release the locked UTXOs in case of",
            "an error on the caller's side.",
        },
        Req: mkwalletrpc_FundPsbtRequest(),
        Res: mkwalletrpc_FundPsbtResponse(),
    }
}
func WalletKit_FinalizePsbt() Method {
    return Method{
        Name: "FinalizePsbt",
        Service: "WalletKit",
        Description: []string{
            "FinalizePsbt expects a partial transaction with all inputs and outputs fully",
            "declared and tries to sign all inputs that belong to the wallet. Lnd must be",
            "the last signer of the transaction. That means, if there are any unsigned",
            "non-witness inputs or inputs without UTXO information attached or inputs",
            "without witness data that do not belong to lnd's wallet, this method will",
            "fail. If no error is returned, the PSBT is ready to be extracted and the",
            "final TX within to be broadcast.",
            "NOTE: This method does NOT publish the transaction once finalized. It is the",
            "caller's responsibility to either publish the transaction on success or",
            "unlock/release any locked UTXOs in case of an error in this method.",
        },
        Req: mkwalletrpc_FinalizePsbtRequest(),
        Res: mkwalletrpc_FinalizePsbtResponse(),
    }
}
func Router_SendPaymentV2() Method {
    return Method{
        Name: "SendPaymentV2",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Send a payment over lightning",
        Description: []string{
            "SendPaymentV2 attempts to route a payment described by the passed",
            "PaymentRequest to the final destination. The call returns a stream of",
            "payment updates.",
        },
        Req: mkrouterrpc_SendPaymentRequest(),
        Res: mklnrpc_Payment(),
    }
}
func Router_TrackPaymentV2() Method {
    return Method{
        Name: "TrackPaymentV2",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Track payment",
        Description: []string{
            "TrackPaymentV2 returns an update stream for the payment identified by the",
            "payment hash.",
        },
        Req: mkrouterrpc_TrackPaymentRequest(),
        Res: mklnrpc_Payment(),
    }
}
func Router_EstimateRouteFee() Method {
    return Method{
        Name: "EstimateRouteFee",
        Service: "Router",
        Description: []string{
            "EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it",
            "may cost to send an HTLC to the target end destination.",
        },
        Req: mkrouterrpc_RouteFeeRequest(),
        Res: mkrouterrpc_RouteFeeResponse(),
    }
}
func Router_SendToRoute() Method {
    return Method{
        Name: "SendToRoute",
        Service: "Router",
        Description: []string{
            "Deprecated, use SendToRouteV2. SendToRoute attempts to make a payment via",
            "the specified route. This method differs from SendPayment in that it",
            "allows users to specify a full route manually. This can be used for",
            "things like rebalancing, and atomic swaps. It differs from the newer",
            "SendToRouteV2 in that it doesn't return the full HTLC information.",
        },
        Req: mkrouterrpc_SendToRouteRequest(),
        Res: mkrouterrpc_SendToRouteResponse(),
    }
}
func Router_SendToRouteV2() Method {
    return Method{
        Name: "SendToRouteV2",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Send a payment over a predefined route",
        Description: []string{
            "SendToRouteV2 attempts to make a payment via the specified route. This",
            "method differs from SendPayment in that it allows users to specify a full",
            "route manually. This can be used for things like rebalancing, and atomic",
            "swaps.",
        },
        Req: mkrouterrpc_SendToRouteRequest(),
        Res: mklnrpc_HTLCAttempt(),
    }
}
func Router_ResetMissionControl() Method {
    return Method{
        Name: "ResetMissionControl",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Reset internal mission control state",
        Description: []string{
            "ResetMissionControl clears all mission control state and starts with a clean",
            "slate.",
        },
        Req: mkrouterrpc_ResetMissionControlRequest(),
        Res: mkrouterrpc_ResetMissionControlResponse(),
    }
}
func Router_QueryMissionControl() Method {
    return Method{
        Name: "QueryMissionControl",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Query the internal mission control state",
        Description: []string{
            "QueryMissionControl exposes the internal mission control state to callers.",
            "It is a development feature.",
        },
        Req: mkrouterrpc_QueryMissionControlRequest(),
        Res: mkrouterrpc_QueryMissionControlResponse(),
    }
}
func Router_QueryProbability() Method {
    return Method{
        Name: "QueryProbability",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Estimate a success probability",
        Description: []string{
            "QueryProbability returns the current success probability estimate for a",
            "given node pair and amount.",
        },
        Req: mkrouterrpc_QueryProbabilityRequest(),
        Res: mkrouterrpc_QueryProbabilityResponse(),
    }
}
func Router_BuildRoute() Method {
    return Method{
        Name: "BuildRoute",
        Service: "Router",
        Category: "Payment",
        ShortDescription: "Build a route from a list of hop pubkeys",
        Description: []string{
            "BuildRoute builds a fully specified route based on a list of hop public",
            "keys. It retrieves the relevant channel policies from the graph in order to",
            "calculate the correct fees and time locks.",
        },
        Req: mkrouterrpc_BuildRouteRequest(),
        Res: mkrouterrpc_BuildRouteResponse(),
    }
}
func Router_SubscribeHtlcEvents() Method {
    return Method{
        Name: "SubscribeHtlcEvents",
        Service: "Router",
        Description: []string{
            "SubscribeHtlcEvents creates a uni-directional stream from the server to",
            "the client which delivers a stream of htlc events.",
        },
        Req: mkrouterrpc_SubscribeHtlcEventsRequest(),
        Res: mkrouterrpc_HtlcEvent(),
    }
}
func Router_SendPayment() Method {
    return Method{
        Name: "SendPayment",
        Service: "Router",
        Description: []string{
            "Deprecated, use SendPaymentV2. SendPayment attempts to route a payment",
            "described by the passed PaymentRequest to the final destination. The call",
            "returns a stream of payment status updates.",
        },
        Req: mkrouterrpc_SendPaymentRequest(),
        Res: mkrouterrpc_PaymentStatus(),
    }
}
func Router_TrackPayment() Method {
    return Method{
        Name: "TrackPayment",
        Service: "Router",
        Description: []string{
            "Deprecated, use TrackPaymentV2. TrackPayment returns an update stream for",
            "the payment identified by the payment hash.",
        },
        Req: mkrouterrpc_TrackPaymentRequest(),
        Res: mkrouterrpc_PaymentStatus(),
    }
}
func Router_HtlcInterceptor() Method {
    return Method{
        Name: "HtlcInterceptor",
        Service: "Router",
        Description: []string{
            "HtlcInterceptor dispatches a bi-directional streaming RPC in which",
            "Forwarded HTLC requests are sent to the client and the client responds with",
            "a boolean that tells LND if this htlc should be intercepted.",
            "In case of interception, the htlc can be either settled, cancelled or",
            "resumed later by using the ResolveHoldForward endpoint.",
        },
        Req: mkrouterrpc_ForwardHtlcInterceptResponse(),
        Res: mkrouterrpc_ForwardHtlcInterceptRequest(),
    }
}
func Invoices_SubscribeSingleInvoice() Method {
    return Method{
        Name: "SubscribeSingleInvoice",
        Service: "Invoices",
        Description: []string{
            "SubscribeSingleInvoice returns a uni-directional stream (server -> client)",
            "to notify the client of state transitions of the specified invoice.",
            "Initially the current invoice state is always sent out.",
        },
        Req: mkinvoicesrpc_SubscribeSingleInvoiceRequest(),
        Res: mklnrpc_Invoice(),
    }
}
func Invoices_CancelInvoice() Method {
    return Method{
        Name: "CancelInvoice",
        Service: "Invoices",
        Description: []string{
            "CancelInvoice cancels a currently open invoice. If the invoice is already",
            "canceled, this call will succeed. If the invoice is already settled, it will",
            "fail.",
        },
        Req: mkinvoicesrpc_CancelInvoiceMsg(),
        Res: mkinvoicesrpc_CancelInvoiceResp(),
    }
}
func Invoices_AddHoldInvoice() Method {
    return Method{
        Name: "AddHoldInvoice",
        Service: "Invoices",
        Description: []string{
            "AddHoldInvoice creates a hold invoice. It ties the invoice to the hash",
            "supplied in the request.",
        },
        Req: mkinvoicesrpc_AddHoldInvoiceRequest(),
        Res: mkinvoicesrpc_AddHoldInvoiceResp(),
    }
}
func Invoices_SettleInvoice() Method {
    return Method{
        Name: "SettleInvoice",
        Service: "Invoices",
        Description: []string{
            "SettleInvoice settles an accepted invoice. If the invoice is already",
            "settled, this call will succeed.",
        },
        Req: mkinvoicesrpc_SettleInvoiceMsg(),
        Res: mkinvoicesrpc_SettleInvoiceResp(),
    }
}
func MetaService_GetInfo2() Method {
    return Method{
        Name: "GetInfo2",
        Service: "MetaService",
        Req: mklnrpc_GetInfo2Request(),
        Res: mklnrpc_GetInfo2Response(),
    }
}
func MetaService_ChangePassword() Method {
    return Method{
        Name: "ChangePassword",
        Service: "MetaService",
        Req: mklnrpc_ChangePasswordRequest(),
        Res: mklnrpc_ChangePasswordResponse(),
    }
}
func MetaService_CheckPassword() Method {
    return Method{
        Name: "CheckPassword",
        Service: "MetaService",
        Req: mklnrpc_CheckPasswordRequest(),
        Res: mklnrpc_CheckPasswordResponse(),
    }
}
func MetaService_ForceCrash() Method {
    return Method{
        Name: "ForceCrash",
        Service: "MetaService",
        Req: mklnrpc_CrashRequest(),
        Res: mklnrpc_CrashResponse(),
    }
}
func WalletUnlocker_GenSeed() Method {
    return Method{
        Name: "GenSeed",
        Service: "WalletUnlocker",
        Category: "Seed",
        ShortDescription: "Create a secret seed",
        Description: []string{
            "GenSeed is the first method that should be used to instantiate a new lnd",
            "instance. This method allows a caller to generate a new aezeed cipher seed",
            "given an optional passphrase. If provided, the passphrase will be necessary",
            "to decrypt the cipherseed to expose the internal wallet seed.",
            "Once the cipherseed is obtained and verified by the user, the InitWallet",
            "method should be used to commit the newly generated seed, and create the",
            "wallet.",
        },
        Req: mklnrpc_GenSeedRequest(),
        Res: mklnrpc_GenSeedResponse(),
    }
}
func WalletUnlocker_InitWallet() Method {
    return Method{
        Name: "InitWallet",
        Service: "WalletUnlocker",
        Category: "Wallet",
        ShortDescription: "Initialize a wallet when starting lnd for the first time",
        Description: []string{
            "InitWallet is used when lnd is starting up for the first time to fully",
            "initialize the daemon and its internal wallet. At the very least a wallet",
            "password must be provided. This will be used to encrypt sensitive material",
            "on disk.",
            "In the case of a recovery scenario, the user can also specify their aezeed",
            "mnemonic and passphrase. If set, then the daemon will use this prior state",
            "to initialize its internal wallet.",
            "Alternatively, this can be used along with the GenSeed RPC to obtain a",
            "seed, then present it to the user. Once it has been verified by the user,",
            "the seed can be fed into this RPC in order to commit the new wallet.",
        },
        Req: mklnrpc_InitWalletRequest(),
        Res: mklnrpc_InitWalletResponse(),
    }
}
func WalletUnlocker_UnlockWallet() Method {
    return Method{
        Name: "UnlockWallet",
        Service: "WalletUnlocker",
        Category: "Wallet",
        ShortDescription: "Unlock an encrypted wallet at startup",
        Description: []string{
            "UnlockWallet is used at startup of lnd to provide a password to unlock",
            "the wallet database.",
        },
        Req: mklnrpc_UnlockWalletRequest(),
        Res: mklnrpc_UnlockWalletResponse(),
    }
}
