#!  /usr/bin/bash

################################################################################
#   smoke tests for pld REST endpoints
################################################################################

export  PLD_REST_SERVER='http://localhost:8080'
export  PLD_REST_CONTEXT='/api/v1'
export  REST_ERRORS_FILE='./rest.err'
export  JSON_OUTPUT=''
export  WALLET_PASSPHRASE='w4ll3tP@sswd'
export  VERBOSE='false'

#   color commands
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
LIGHTGRAY='\033[0;37m'
NOCOLOR='\033[0m'

#   use curl to execute a command
executeCommand() {
    local COMMAND="${1}"
    local HTTP_METHOD="${2}"
    local URI="${3}"
    local PAYLOAD="${4}"


    JSON_OUTPUT=""
    HTTP_RESPONSE_CODE=""

    if [ "${HTTP_METHOD}" == "GET" ]
    then
        if [ "${VERBOSE}" == 'true' ]
        then
            echo -e "[trace] ${LIGHTGRAY}curl \"${PLD_REST_SERVER}${PLD_REST_CONTEXT}${URI}\"${NOCOLOR}"
        fi

        OUTPUT=$( curl --write-out '|%{response_code}' "${PLD_REST_SERVER}${PLD_REST_CONTEXT}${URI}" 2>> ${REST_ERRORS_FILE} | tr -d '\n' )
        JSON_OUTPUT=$( echo "${OUTPUT}" | perl -ne 'if( $_ =~ /^(.+)\|(\d{3})$/ ) { print qq/$1\n/; }' )
        HTTP_RESPONSE_CODE=$( echo "${OUTPUT}" | perl -ne 'if( $_ =~ /^(.+)\|(\d{3})$/ ) { print qq/$2\n/; }' )
    elif [ "${HTTP_METHOD}" == "POST" ]
    then
        if [ "${VERBOSE}" == 'true' ]
        then
            echo -e "[trace] ${LIGHTGRAY}curl -H \"Content-Type: application/json\" -X POST -d '${PAYLOAD}' \"${PLD_REST_SERVER}${PLD_REST_CONTEXT}${URI}\"${NOCOLOR}"
        fi

        OUTPUT=$( curl --write-out '|%{response_code}' -H "Content-Type: application/json" -X POST -d "${PAYLOAD}" "${PLD_REST_SERVER}${PLD_REST_CONTEXT}${URI}" 2>> ${REST_ERRORS_FILE} | tr -d '\n' )
        JSON_OUTPUT=$( echo "${OUTPUT}" | perl -ne 'if( $_ =~ /^(.+)\|(\d{3})$/ ) { print qq/$1\n/; }' )
        HTTP_RESPONSE_CODE=$( echo "${OUTPUT}" | perl -ne 'if( $_ =~ /^(.+)\|(\d{3})$/ ) { print qq/$2\n/; }' )
    else
        echo -e "${RED}error: invalid HTTP method \"${HTTP_METHOD}\"${NOCOLOR}"
        return 1
    fi

    if [ "${VERBOSE}" == 'true' ]
    then
        echo -e "[trace] ${LIGHTGRAY}response: ${JSON_OUTPUT}${NOCOLOR}"
    fi

    if [ "${HTTP_RESPONSE_CODE}" == "200" ]
    then
        echo -e ">>> ${CYAN}[${COMMAND}]${NOCOLOR}: ${GREEN}command successfully executed${NOCOLOR}"
    else
        ERROR_MESSAGE=$( echo "${JSON_OUTPUT}" | jq '.message' )
        echo -e ">>> ${CYAN}[${COMMAND}]${NOCOLOR}: ${RED}error running pld command: HTTP response: ${HTTP_RESPONSE_CODE}\n\t${ERROR_MESSAGE}${NOCOLOR}"
        JSON_OUTPUT=''
        return 1
    fi
}

#   use jq to filter results of previously executed command
showCommandResult() {
    local TITLE="${1}"
    local FILTER="${2}"

    if [ ! -z "${JSON_OUTPUT}" ]
    then
        if [ -z "${FILTER}" ]
        then
            RESULT="${JSON_OUTPUT}"
        else
            RESULT=$( echo "${JSON_OUTPUT}" | jq "${FILTER}" )
        fi
        echo -e "    >>> ${TITLE}: ${LIGHTGRAYNOCOLOR}${RESULT}${NOCOLOR}"
    fi
}

#   use jq to filter results of previously executed command
getCommandResult() {
    local FILTER="${1}"

    if [ ! -z "${JSON_OUTPUT}" ]
    then
        if [ ! -z "${FILTER}" ]
        then
            echo -e "$( echo "${JSON_OUTPUT}" | jq "${FILTER}" )"
        fi
    fi
}

#   splash screen
echo ">>>>> Testing pld REST endpoints"
echo

#   check if curl is available
OUTPUT=$( which curl 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'curl' is required to run this script"
fi

#   check if jq is available
OUTPUT=$( which jq 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'jq' is required to run this script"
fi

#   parse CLI arguments
while [ true ]
do
    ARG=${1}

    if [ -z "${ARG}" ]
    then
        break
    fi

    if [ "${ARG}" == "--verbose" ]
    then
        VERBOSE='true'
    fi

    shift
done

#
#   test commands of "meta" group
#

echo -e ">>> Group ${CYAN}[meta]${NOCOLOR} API endpoints which are relevant to the entire pld node, not any specific part"
echo

executeCommand 'debuglevel' 'POST' '/meta/debuglevel' '{ "show": true, "level_spec": "debug" }'
showCommandResult 'subsystems' '.subSystems'

executeCommand 'getinfo' 'GET' '/meta/getinfo'
showCommandResult '#neutrino peers' '.neutrino.peers | length'

#   this needs be the last of all tests
#executeCommand 'stop' 'GET' '/meta/stop'
#showCommandResult 'result' ''

executeCommand 'version' 'GET' '/meta/version'
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "util/seed" group
#

echo -e ">>> Group ${CYAN}[util/seed]${NOCOLOR} Manipulation of mnemonic seed phrases which represent wallet keys"
echo

executeCommand 'changepassphrase' 'POST' '/util/seed/changepassphrase' '{ "current_seed_passphrase_bin": "cGFzc3dvcmQ=", "current_seed": [ "plastic",  "hollow",  "mansion",  "keep",  "into",  "cloth",  "awesome",  "salmon",  "reopen",  "inner",  "replace",  "dice",  "life",  "example",  "around" ], "new_seed_passphrase": "password" }'
showCommandResult 'new ciphered seed' '.seed'

#   this test is meant to fail, since seed creation can only be ordered before wallet's creation
executeCommand 'genseed' 'POST' '/util/seed/create' '{ "seed_passphrase_bin": "cGFzc3dvcmQ=" }'
showCommandResult 'ciphered seed' '.seed'

echo "----------"
echo

#
#   test commands of "Lightning/Channel" group
#

echo -e ">>> Group ${CYAN}[lightning/channel]${NOCOLOR} Management of lightning channels to direct peers of this pld node"
echo

#   fetch a public key for the channels tests
executeCommand 'describegraph' 'POST' '/lightning/graph' '{ "includeUnannounced": true }'
showCommandResult 'public key' '.nodes | .[0] | .pubKey'
PUBLIC_KEY="$( getCommandResult '.nodes | .[0] | .pubKey ' | tr -d '\"' | cut --characters=3- )"

AMOUNT="100000"
executeCommand 'openchannel' 'POST' '/lightning/channel/open' "{ \"node_pubkey\": \"${PUBLIC_KEY}\", \"local_funding_amount\": ${AMOUNT} }"
showCommandResult 'result' ''

CHANNEL_POINT="XPTO"
executeCommand 'closechannel' 'POST' '/lightning/channel/close' "{ \"channel_point\": \"${CHANNEL_POINT}\" }"
showCommandResult 'result' ''

FUNDING_TXID="934095dc4afa8d4b5d43732a96e78e11c0e88defdaab12d946f525e54478938f"
executeCommand 'abandonchannel' 'POST' '/lightning/channel/abandon' "{ \"channelPoint\": { \"funding_txid_str\": \"${FUNDING_TXID}\" } }"
showCommandResult 'result' ''

executeCommand 'channelbalance' 'GET' '/lightning/channel/balance'
showCommandResult 'channel balance' '.balance'

executeCommand 'pendingchannels' 'GET' '/lightning/channel/pending'
showCommandResult '#pending open channels' '.pendingOpenChannels | length'
showCommandResult '#pending closing channels' '.pendingClosingChannels | length'
showCommandResult 'limbo balance' '.totalLimboBalance'

executeCommand 'listchannels' 'POST' '/lightning/channel' '{  }'
showCommandResult '#open channels' '.channels | length'

executeCommand 'closedchannels' 'POST' '/lightning/channel/closed' '{  }'
showCommandResult '#closed channels' '.channels | length'

executeCommand 'getnetworkinfo' 'GET' '/lightning/channel/networkinfo'
showCommandResult 'nodes' '.numNodes'
showCommandResult 'channels' '.numChannels'

executeCommand 'feereport' 'GET' '/lightning/channel/feereport'
showCommandResult '#channel fees' '.channelFees | length'
showCommandResult 'week fee sum' '.weekFeeSum'

executeCommand 'updatechanpolicy' 'POST' '/lightning/channel/policy' '{ "baseFeeMsat": 10, "feeRate": 10, "timeLockDelta": 20, "maxHtlcMsat": 30, "minHtlcMsat": 1, "minHtlcMsatSpecified": false }'
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Lightning/Channel/Backup" group
#

echo -e ">>> Group ${CYAN}[lightning/channel/backup]${NOCOLOR} Backup and recovery of the state of active Lightning channels to and from this pld node"
echo

executeCommand 'exportchanbackup' 'POST' '/lightning/channel/backup/export' "{ \"chanPoint\": { \"funding_txid_str\": \"${FUNDING_TXID}\" } }"
showCommandResult 'result' '.multi_chan_backup.multi_chan_backup'

executeCommand 'restorechanbackup' 'POST' '/lightning/channel/backup/restore' "{ \"chanBackups\": [ { \"chanPoint\": { \"fundingTxidStr\": \"${FUNDING_TXID}\", \"outputIndex\": 1000 }, \"chanBackup\": \"RW5jcnlwdGVkIENoYW4gQmFja3Vw\" } ] }"
showCommandResult 'result' ''

executeCommand 'verifychanbackup' 'POST' '/lightning/channel/backup/verify' "{ \"singleChanBackups\": { \"chanBackups\": [ { \"chanPoint\": { \"fundingTxidStr\": \"${FUNDING_TXID}\", \"outputIndex\": 1000 }, \"chanBackup\": \"RW5jcnlwdGVkIENoYW4gQmFja3Vw\" } ] } }"
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Lightning/Graph" group
#

echo -e ">>> Group ${CYAN}[lightning/graph]${NOCOLOR} Information about the global known Lightning Network"
echo

#   fetch a public key for the graph tests
executeCommand 'describegraph' 'POST' '/lightning/graph' '{ "includeUnannounced": true }'
showCommandResult 'last update' '.nodes | .[0] | .lastUpdate '
showCommandResult 'public key' '.nodes | .[0] | .pubKey'
PUBLIC_KEY="$( getCommandResult '.nodes | .[0] | .pubKey ' | tr -d '\"' )"

executeCommand 'getnodemetrics' 'POST' '/lightning/graph/nodemetrics' '{ "types": [ 0, 1 ] }'
showCommandResult 'betweenness centrality' '.betweennessCentrality'

CHAN_ID=123
executeCommand 'getchaninfo' 'POST' '/lightning/graph/channel' "{ \"chanId\": ${CHAN_ID} }"
showCommandResult 'result' ''

executeCommand 'getnodeinfo' 'POST' '/lightning/graph/nodeinfo' "{ \"pubKey\": \"${PUBLIC_KEY}\", \"includeChannels\": true }"
showCommandResult 'last update' '.node.lastUpdate'
showCommandResult '#channels' '.node.channels | length'

echo "----------"
echo

#
#   test commands of "Lightning/Invoice" group
#

echo -e ">>> Group ${CYAN}[lightning/invoice]${NOCOLOR} Management of invoices which are used to request payment over Lightning"
echo

executeCommand 'addinvoice' 'POST' '/lightning/invoice/create' '{ "memo": "xpto", "value": 10, "expiry": 3600 }'
showCommandResult 'rHash' '.rHash'
showCommandResult 'payment request' '.paymentRequest'
RHASH="$( getCommandResult '.rHash' | tr -d '\"' )"
PAYREQ="$( getCommandResult '.paymentRequest' | tr -d '\"' )"

executeCommand 'lookupinvoice' 'POST' '/lightning/invoice/lookup' "{ \"rHash\": \"${RHASH}\" }"
showCommandResult 'last update' '.lastUpdate'
showCommandResult 'index' '.addIndex'
showCommandResult 'state' '.state'

executeCommand 'listinvoices' 'POST' '/lightning/invoice' '{ "indexOffset": 1, "numMaxInvoices": 10 }'
showCommandResult '#invoices' '.invoices | length'

executeCommand 'decodepayreq' 'POST' '/lightning/invoice/decodepayreq' "{ \"payReq\": \"${PAYREQ}\" }"
showCommandResult 'destination' '.destination'
showCommandResult 'payment Hash' '.paymentHash'
showCommandResult '#satoshis' '.numSatoshis'

echo "----------"
echo

#
#   test commands of "Lightning/Payment" group
#

echo -e ">>> Group ${CYAN}[lightning/payment]${NOCOLOR} Lightning network payments which have been made, or have been forwarded, through this node"
echo

executeCommand 'sendpayment' 'POST' '/lightning/payment/send' "{ \"paymentHash\": \"${RHASH}\", \"amt\": 100000, \"dest\": \"${TARGET_WALLET}\" }"
showCommandResult 'result' ''

executeCommand 'payinvoice' 'POST' '/lightning/payment/payinvoice' "{ \"paymentHash\": \"${RHASH}\", \"amt\": 100000, \"dest\": \"${TARGET_WALLET}\" }"
showCommandResult 'result' ''

executeCommand 'sendtoroute' 'POST' '/lightning/payment/sendtoroute' '{ "paymentHash": "02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e", "route": { "hops": { "chanId": "xpto"} } }'
showCommandResult 'result' ''

executeCommand 'listpayments' 'POST' '/lightning/payment' '{ "indexOffset": 1, "maxPayments": 10, "includeIncomplete": true }'
showCommandResult '#payments' '.payments | length'

executeCommand 'trackpayment' 'POST' '/lightning/payment/track' '{ "indexOffset": 1, "maxPayments": 10, "includeIncomplete": true }'
showCommandResult '#payments' '.payments | length'

executeCommand 'queryroutes' 'POST' '/lightning/payment/queryroutes' "{  }"
showCommandResult 'result' ''

executeCommand 'fwdinghistory' 'POST' '/lightning/payment/fwdinghistory' '{ "indexOffset": 0, "numMaxEvents": 25 }'
showCommandResult '#forwarding events' '.forwardingEvents | length'

executeCommand 'querymc' 'GET' '/lightning/payment/querymc'
showCommandResult 'result' ''

FROM_NODE="01020304"
TO_NODE="02030405"
AMOUNT="100000"
executeCommand 'queryprob' 'POST' '/lightning/payment/queryprob' "{ \"fromNode\": \"${FROM_NODE}\", \"toNode\": \"${TO_NODE}\", \"amtMsat\": \"${AMOUNT}\" }"
showCommandResult 'result' ''

executeCommand 'resetmc' 'GET' '/lightning/payment/resetmc'
showCommandResult 'result' ''

executeCommand 'buildroute' 'POST' '/lightning/payment/buildroute' '{ "amtMsat": 0, "hopPubkeys": [ "01020304", "02030405", "03040506" ] }'
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Lightning/Peer" group
#

echo -e ">>> Group ${CYAN}[lightning/peer]${NOCOLOR} Connections to other nodes in the Lightning Network"
echo

executeCommand 'connect' 'POST' '/lightning/peer/connect' "{ \"addr\": { \"pubkey\": \"${PUBLIC_KEY}\", \"host\": \"192.168.40.1:8080\" } }"
showCommandResult 'result' ''

executeCommand 'disconnect' 'POST' '/lightning/peer/disconnect' "{ \"pubkey\": \"${PUBLIC_KEY}\" }"
showCommandResult 'result' ''

executeCommand 'listpeers' 'GET' '/lightning/peer'
showCommandResult 'result' ''
showCommandResult '#peers' '.peers | length'

echo "----------"
echo

#
#   test commands of "Neutrino" group
#

echo -e ">>> Group ${CYAN}[neutrino]${NOCOLOR} Management of the Neutrino interface which is used to communicate with the p2p nodes in the network"
echo

#   fetch a transaction ID for the neutrino tests
TARGET_WALLET="pkt1q07ly7r47ss4drsvt2zq9zkcstksrq2dap3x0yw"
executeCommand 'sendmany' 'POST' '/wallet/transaction/sendmany' "{ \"AddrToAmount\": { \"${TARGET_WALLET}\": ${AMOUNT} } }"
showCommandResult 'transaction ID' '.txid'
TXID="$( getCommandResult '.txid' | tr -d '\"' )"

executeCommand 'bcasttransaction' 'POST' '/neutrino/bcasttransaction' "{ \"tx\": \"${TXID}\" }"
showCommandResult 'result' ''
#    echo -e "\ttransaction hash: $( echo ${JSON_OUTPUT} | jq '.txn_hash' )"

executeCommand 'estimatefee' 'POST' '/neutrino/estimatefee' "{ \"AddrToAmount\": [ \"${TARGET_WALLET}\": 100000 ] }"
showCommandResult 'fee sat' '.feeSat'

echo "----------"
echo

#
#   test commands of "Wallet" group
#

echo -e ">>> Group ${CYAN}[wallet]${NOCOLOR} APIs for management of on-chain (non-Lightning) payments, seed export and recovery, and on-chain transaction detection"
echo

executeCommand 'walletbalance' 'GET' '/wallet/balance'
echo -e "\ttotal balance: $( echo ${JSON_OUTPUT} | jq '.totalBalance' )"

#   we don't want to change the wallet's passphrase here, since this test is being made by REST_createWalletTest.sh script
#executeCommand 'changePassphrase' 'POST' '/wallet/changepassphrase' "{ \"current_passphrase\": \"${PASSPHRASE}\", \"new_passphrase\": \"${NEW_PASSPHRASE}\" }"
#showCommandResult 'result' ''

WALLET_PASSPHRASE='w4ll3tP@sswd'
executeCommand 'checkPassphrase' 'POST' '/wallet/checkpassphrase' "{ \"wallet_passphrase\": \"${WALLET_PASSPHRASE}\" }"
showCommandResult 'result' '.validPassphrase'

#   this test is meant to fail, since wallet is supposed to have been created already
WALLET_SEED='[ "plastic",  "hollow",  "mansion",  "keep",  "into",  "cloth",  "awesome",  "salmon",  "reopen",  "inner",  "replace",  "dice",  "life",  "example",  "around" ]'
SEED_PASSPHRASE='cGFzc3dvcmQ='
executeCommand 'create_wallet' 'POST' '/wallet/create' "{ \"wallet_passphrase\": \"${PASSPHRASE}\", \"wallet_seed\": ${WALLET_SEED}, \"seed_passphrase_bin\": \"${SEED_PASSPHRASE}\" }"
showCommandResult 'result' ''

executeCommand 'getsecret' 'POST' '/wallet/getsecret' '{ "name": "Isaac Assimov" }'
showCommandResult 'result' '.secret'

executeCommand 'getwalletseed' 'GET' '/wallet/seed'
showCommandResult 'result' '.seed'

#   we don't want to unlock the wallet here, since this test is being made by REST_createWalletTest.sh script
#executeCommand 'unlock_wallet' 'POST' '/wallet/unlock' "{ \"wallet_passphrase\": \"${WALLET_PASSPHRASE}\" }"
#showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Wallet/Address" group
#

echo -e ">>> Group ${CYAN}[wallet/address]${NOCOLOR} Management of individual wallet addresses"
echo

executeCommand 'getaddressbalances' 'POST' '/wallet/address/balances' '{  }'
echo -e "\t#addresses: $( echo ${JSON_OUTPUT} | jq '.addrs | length' )"

executeCommand 'newaddress' 'POST' '/wallet/address/create' '{  }'
showCommandResult 'result' ''

#executeCommand 'dumpprivkey' 'POST' '/wallet/address/dumpprivkey' '{ "address": "pkt1q85n69mzthdxlwutn6dr6f7kwyd9nv8ulasdaqz" }'
executeCommand 'dumpprivkey' 'POST' '/wallet/address/dumpprivkey' "{ \"address\": \"${TARGET_WALLET}\" }"
showCommandResult 'result' ''
#    echo -e "\tprivate key: $( echo ${JSON_OUTPUT} | jq '.private_key' )"

#   we don't want to start a full rescan so, adding a 'xx' sufix to the Wallet PK to force an error
WALLET_PRIVKEY='cVgcgWwQpwzViWmG7dGyvf545ra6AdT4tV29UtQfE8okvPuznFZi'
executeCommand 'importprivkey' 'POST' '/wallet/address/import' "{ \"privateKey\": \"${WALLET_PRIVKEY}xx\", \"rescan\": true }"
showCommandResult 'result' ''
#    echo -e "\taddress: $( echo ${JSON_OUTPUT} | jq '.address' )"

executeCommand 'signmessage' 'POST' '/wallet/address/signmessage' '{ "msg": "testing pld REST endpoints" }'
showCommandResult 'result' ''
#    echo -e "\tsignature: $( echo ${JSON_OUTPUT} | jq '.signature' )"

echo "----------"
echo

#
#   test commands of "Wallet/NetworkStewardVote" group
#

echo -e ">>> Group ${CYAN}[wallet/networkstewardvote]${NOCOLOR} Control how this wallet votes on PKT Network Steward"
echo

executeCommand 'getnetworkstewardvote' 'GET' '/wallet/networkstewardvote'
showCommandResult 'vote against' '.voteAgainst'
showCommandResult 'vote for' '.voteFor'

executeCommand 'setnetworkstewardvote' 'POST' '/wallet/networkstewardvote/set' '{ "voteAgainst": "0", "voteFor": "1" }'
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Wallet/Transaction" group
#

echo -e ">>> Group ${CYAN}[wallet/transaction]${NOCOLOR} Create and manage on-chain transactions with the wallet"
echo

#executeCommand 'gettransaction' '934095dc4afa8d4b5d43732a96e78e11c0e88defdaab12d946f525e54478938f'
executeCommand 'gettransaction' 'POST' '/wallet/transaction' ''
showCommandResult 'result' ''
#    echo -e "\tamount: $( echo ${JSON_OUTPUT} | jq '.transaction.amount' )"
#    echo -e "\tfee: $( echo ${JSON_OUTPUT} | jq '.transaction.fee' )"

TARGET_WALLET="pkt1q07ly7r47ss4drsvt2zq9zkcstksrq2dap3x0yw"
AMOUNT="10000"
executeCommand 'createtransaction' 'POST' '/wallet/transaction/create' "{ \"toAddress\": \"${TARGET_WALLET}\", \"amount\": ${AMOUNT} }"
showCommandResult 'result' ''
#    echo -e "\ttransaction: $( echo ${JSON_OUTPUT} | jq '.transaction' )"

executeCommand 'query' 'POST' '/wallet/transaction/query' '{  }'
showCommandResult 'result' ''

executeCommand 'sendcoins' 'POST' '/wallet/transaction/sendcoins' "{ \"addr\": \"${TARGET_WALLET}\", \"amount\": ${AMOUNT} }"
showCommandResult 'transaction ID' '.txid'

executeCommand 'sendfrom' 'POST' '/wallet/transaction/sendfrom' '{  }'
showCommandResult 'result' ''

executeCommand 'sendmany' 'POST' '/wallet/transaction/sendmany' "{ \"AddrToAmount\": { \"${TARGET_WALLET}\": ${AMOUNT} } }"
showCommandResult 'result' ''
showCommandResult 'transaction ID' '.txid'
TXID="$( getCommandResult '.txid' | tr -d '\"' )"

echo "----------"
echo

#
#   test commands of "Wallet/Unspent" group
#

echo -e ">>> Group ${CYAN}[wallet/unspent]${NOCOLOR} Detected unspent transactions associated with one of our wallet addresses"
echo

executeCommand 'listunspent' 'POST' '/wallet/unspent' '{ "minConfs": 1, "maxConfs": 100 }'
showCommandResult '#utxos' '.utxos | length'

executeCommand 'resync' 'POST' '/wallet/unspent/resync' '{  }'
showCommandResult 'result' ''

executeCommand 'stopresync' 'GET' '/wallet/unspent/stopresync' ''
showCommandResult 'result' ''
#    echo -e "\tstop sync: $( echo ${JSON_OUTPUT} | jq '.value' )"

#
#   test commands of "Wallet/Unspent/Lock" group
#

echo -e ">>> Group ${CYAN}[wallet/unspent/lock]${NOCOLOR} Manipulation of unspent outputs which are 'locked' and therefore will not be used to source funds for any transaction"
echo

executeCommand 'listlockunspent' 'GET' '/wallet/unspent/lock'
showCommandResult 'result' ''
#    echo -e "\t#lock unspent: $( echo ${JSON_OUTPUT} | jq '.locked_unspent | length' )"

executeCommand 'lockunspent' 'POST' '/wallet/unspent/lock/create' '{ "lockname": "secure vault", "unlock": false, "transactions": [ { "txid": "934095dc4afa8d4b5d43732a96e78e11c0e88defdaab12d946f525e54478938f" } ] }' ''
showCommandResult 'result' ''

echo "----------"
echo

#
#   test commands of "Wtclient/Tower" group
#

echo -e ">>> Group ${CYAN}[wtclient/tower]${NOCOLOR} Interact with the watchtower client"
echo

executeCommand 'show' 'POST' '/wtclient/tower' '{  }'
showCommandResult 'result' ''

executeCommand 'show' 'POST' '/wtclient/tower/create' '{  }'
showCommandResult 'result' ''

executeCommand 'show' 'POST' '/wtclient/tower/getinfo' '{  }'
showCommandResult 'result' ''

executeCommand 'show' 'POST' '/wtclient/tower/policy' '{  }'
showCommandResult 'result' ''

executeCommand 'show' 'POST' '/wtclient/tower/remove' '{  }'
showCommandResult 'result' ''

executeCommand 'show' 'POST' '/wtclient/tower/stats' '{  }'
showCommandResult 'result' ''

echo "----------"
echo

rm -rf ${REST_ERRORS_FILE}
