#!  /usr/bin/bash

################################################################################
#   smoke tests for pld / pldctl commands
################################################################################

export  PKT_HOME="$( pwd )"
export  PLD="${PKT_HOME}/bin/pld"
export  PLD_OPTIONS=""
export  PLD_OUTPUT_FILE="./pld.out"
export  PLD_PID=
export  PLDCTL="${PKT_HOME}/bin/pldctl"
export  PLDCTL_OPTIONS=""
export  PLDCTL_ERRORS_FILE="./pldctl.err"
export  JSON_OUTPUT=""
export  TARGET_WALLET="pkt1q07ly7r47ss4drsvt2zq9zkcstksrq2dap3x0yw"

#   start pld deamon in background
startPldDeamon() {

    ${PLD} ${PLD_OPTIONS} > ${PLD_OUTPUT_FILE} &

    PLD_PID=$!

    echo ">>> ${PLD} daemon up and running: PID: ${PLD_PID}"

    sleep 10s
}

#   stop pld deamon
stopPldDeamon() {

    executeCommand 'stop'

    kill ${PLD_PID} 2> /dev/null

    sleep 10s
}

#   send a command to create a wallet
createWallet() {

    local OUTPUT=$( perl -w ./test/createWallet.pl )

    if [ -z "$( echo ${OUTPUT} | grep 'pld successfully initialized!' )" ]
    then
        kill ${PLD_PID}
        exit "error: fail attempting to run command create wallet"
    fi

    echo ">>> create: command successfully executed"
}

#   send a command to unlock the wallet
unlockWallet() {

    local OUTPUT=$( perl -w ./test/unlockWallet.pl )

    if [ -z "$( echo ${OUTPUT} | grep 'lnd successfully unlocked!' )" ]
    then
        kill ${PLD_PID}
        exit "error: fail attempting to run command unlock wallet"
    fi

    echo ">>> unlock: command successfully executed"
}

#   use pldctl to execute a command
executeCommand() {
    local COMMAND="${1}"
    local ARGUMENTS="${2}"

    JSON_OUTPUT=$( ${PLDCTL} ${PLDCTL_OPTIONS} ${COMMAND} ${ARGUMENTS} 2>> ${PLDCTL_ERRORS_FILE} )
    if [ $? -eq 0 ]
    then
        echo ">>> ${COMMAND} ${ARGUMENTS}: command successfully executed"
    else
        echo "error: fail attempting to run command \"${COMMAND} ${ARGUMENTS}\": $?"
        return 1
    fi
}

#   splash screen
echo ">>>>> Testing pld and pldctl"
echo

#   check if jq is available

output=$( which jq 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'jq' is required to run this script"
fi

#   parse CLI arguments
CREATE_WALLET="false"

while [ true ]
do
    ARG=${1}

    if [ -z "${ARG}" ]
    then
        break
    fi

    if [ "${ARG}" == "--createWallet" ]
    then
        CREATE_WALLET="true"
    fi

    shift
done

#   create wallet when requested
if [ "${CREATE_WALLET}" == "true" ]
then
    #   clean things up by removing previous wallet
    rm -rf ~/.lncli ~/.pki ~/.pktd ~/.pktwallet
    rm -rf ${PLD_OUTPUT_FILE}
    rm -rf ${PLDCTL_ERRORS_FILE}

    #   start pld deamon, create a wallet and stop the deamon, because first test is unlock wallet
    startPldDeamon
    createWallet
    stopPldDeamon
fi

#   star pld daemon and test the command to unlock the wallet
startPldDeamon
unlockWallet

#   test commands to get info about the running pld daemon
executeCommand 'getinfo'
if [ $? -eq 0 ]
then
    echo -e "\t#neutrino peers: $( echo ${JSON_OUTPUT} | jq '.neutrino.peers | length' )"
fi

executeCommand 'getrecoveryinfo'
if [ $? -eq 0 ]
then
    echo -e "\trecovery mode: $( echo ${JSON_OUTPUT} | jq '.recovery_mode' )"
fi

executeCommand 'debuglevel' '--level info --show'

executeCommand 'version'
if [ $? -eq 0 ]
then
    echo -e "\tpld version: $( echo ${JSON_OUTPUT} | jq '.pld | .version' )"
    echo -e "\tpldctl version: $( echo ${JSON_OUTPUT} | jq '.pldctl | .version' )"
fi

#   test commands to manage channels
#executeCommand 'openchannel'
#executeCommand 'closechannel'
#executeCommand 'closeallchannels'
#executeCommand 'abandonchannel'

executeCommand 'channelbalance'
if [ $? -eq 0 ]
then
    echo -e "\tchannel balance: $( echo ${JSON_OUTPUT} | jq '.balance' )"
fi

executeCommand 'pendingchannels'
if [ $? -eq 0 ]
then
    echo -e "\tlimbo balance: $( echo ${JSON_OUTPUT} | jq '.total_limbo_balance' )"
fi

executeCommand 'listchannels'
if [ $? -eq 0 ]
then
    echo -e "\t#open channels: $( echo ${JSON_OUTPUT} | jq '.channels | length' )"
fi

executeCommand 'closedchannels'
if [ $? -eq 0 ]
then
    echo -e "\t#closed channels: $( echo ${JSON_OUTPUT} | jq '.channels | length' )"
fi

executeCommand 'getnetworkinfo'
if [ $? -eq 0 ]
then
    echo -e "\t#nodes: $( echo ${JSON_OUTPUT} | jq '.num_nodes' )"
    echo -e "\t#channels: $( echo ${JSON_OUTPUT} | jq '.num_channels' )"
fi

executeCommand 'feereport'
if [ $? -eq 0 ]
then
    echo -e "\tweek fee sum: $( echo ${JSON_OUTPUT} | jq '.week_fee_sum' )"
fi

#executeCommand 'updatechanpolicy' '10 10 20'

executeCommand 'exportchanbackup' '--all'
if [ $? -eq 0 ]
then
    MULTI_BACKUP=$( echo ${JSON_OUTPUT} | jq '.multi_chan_backup.multi_chan_backup' | tr --delete '"' )
    echo -e "\tmulti backup: ${MULTI_BACKUP}"
fi

executeCommand 'verifychanbackup' "--multi_backup=${MULTI_BACKUP}"
executeCommand 'restorechanbackup'

#   test commands to get graph info
executeCommand 'describegraph'
#executeCommand 'getnodemetrics'
#executeCommand 'getchaninfo'
executeCommand 'getnodeinfo'

#   test commands to manage invoices
executeCommand 'addinvoice'
#executeCommand 'lookupinvoice'
executeCommand 'listinvoices'
#executeCommand 'decodepayreq'

#   test commands to manage on-chain transactions
executeCommand 'estimatefee' "{ \"${TARGET_WALLET}\": 10000000 }"
if [ $? -eq 0 ]
then
    echo -e "\tfee sat: $( echo ${JSON_OUTPUT} | jq '.fee_sat' )"
fi

executeCommand 'sendmany' "{ \"${TARGET_WALLET}\": 10000000 }"
if [ $? -eq 0 ]
then
    echo -e "\ttransaction Id: $( echo ${JSON_OUTPUT} | jq '.txid' )"
fi

executeCommand 'sendcoins' "{ \"${TARGET_WALLET}\": 10000000 }"
if [ $? -eq 0 ]
then
    echo -e "\ttransaction Id: $( echo ${JSON_OUTPUT} | jq '.txid' )"
fi

executeCommand 'listunspent'
if [ $? -eq 0 ]
then
    echo -e "\t#utxos: $( echo ${JSON_OUTPUT} | jq '.utxos | length' )"
fi

executeCommand 'listchaintrns'
if [ $? -eq 0 ]
then
    echo -e "\t#transactions: $( echo ${JSON_OUTPUT} | jq '.transactions | length' )"
fi

executeCommand 'setnetworkstewardvote' "" ""

executeCommand 'getnetworkstewardvote'
if [ $? -eq 0 ]
then
    echo -e "\tvote against: $( echo ${JSON_OUTPUT} | jq '.vote_against' )"
    echo -e "\tvote for: $( echo ${JSON_OUTPUT} | jq '.vote_for' )"
fi

executeCommand 'bcasttransaction'
if [ $? -eq 0 ]
then
    echo -e "\ttransaction hash: $( echo ${JSON_OUTPUT} | jq '.txn_hash' )"
fi

#   test commands to manage payments
#executeCommand 'sendpayment' '02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e 100000 1cc616cdeb96016bf278bfea15d55541d31823986b33c0dab38024cb8eff3791'
#executeCommand 'payinvoice' 'lnpkt100u1p3q4r85pp5kecz6ckl97wwe2nnqn6lq5lju30z9sc8uaeacamudxv52kykgdnqdqqcqzpgsp5fa0tpf3j3ecppn3tvmc50n6w7pl6dcs7zvus82splfjs2qevwkxq9qy9qsq4sfdxwzrku87zaphgh6wa3rtc2a8g7rmg6a2dp4myk3qa8c7409sv205xxfsc2n0mzmemcg92ukg7x6q7xlkp5ca9gdwvsqmtpuazccpw25hg9'
#executeCommand 'sendtoroute' ''

executeCommand 'listpayments' ''
if [ $? -eq 0 ]
then
    echo -e "\t#payments: $( echo ${JSON_OUTPUT} | jq '.payments | length' )"
fi

#executeCommand 'queryroutes' '02e28f38ad50869fd3f3d75147d69bc637090aa9b5013ee49a65c0dda2bf0ab51e 1'

executeCommand 'fwdinghistory'
if [ $? -eq 0 ]
then
    echo -e "\t#forwarding events: $( echo ${JSON_OUTPUT} | jq '.forwarding_events | length' )"
fi

executeCommand 'trackpayment' '1cc616cdeb96016bf278bfea15d55541d31823986b33c0dab38024cb8eff3791'

executeCommand 'querymc' ''
if [ $? -eq 0 ]
then
    echo -e "\t#pairs: $( echo ${JSON_OUTPUT} | jq '.pairs | length' )"
fi

#executeCommand 'queryprob' ''
executeCommand 'resetmc' ''
#executeCommand 'buildroute' ''

#   test commands to manage peers
#executeCommand 'connect' ''
#executeCommand 'disconnect' ''

executeCommand 'listpeers' ''
if [ $? -eq 0 ]
then
    echo -e "\t#peers: $( echo ${JSON_OUTPUT} | jq '.peers | length' )"
fi

#   test commands to deal with profile
executeCommand 'profile' 'add pld_test'
executeCommand 'profile' 'list'
executeCommand 'profile' 'setdefault pld_test'
executeCommand 'profile' 'remove pld_test'

#   remove profile file created by profile commands
rm ~/.lncli/profiles.json

#   test commands to manage the wallet
executeCommand 'newaddress' 'p2wkh'
if [ $? -eq 0 ]
then
    echo -e "\taddress: $( echo ${JSON_OUTPUT} | jq '.address' )"
fi

executeCommand 'walletbalance'
if [ $? -eq 0 ]
then
    echo -e "\ttotal balance: $( echo ${JSON_OUTPUT} | jq '.total_balance' )"
fi

executeCommand 'getaddressbalances'
if [ $? -eq 0 ]
then
    echo -e "\t#addresses: $( echo ${JSON_OUTPUT} | jq '.addrs | length' )"
fi

executeCommand 'signmessage' 'pkt1q0tgwuwcg4tmwegmevdfz3g6tw838upqcq8xt8u message'
if [ $? -eq 0 ]
then
    echo -e "\tsignature: $( echo ${JSON_OUTPUT} | jq '.signature' )"
fi

executeCommand 'resync'
executeCommand 'stopresync'
if [ $? -eq 0 ]
then
    echo -e "\tstop sync: $( echo ${JSON_OUTPUT} | jq '.value' )"
fi

executeCommand 'getwalletseed'
if [ $? -eq 0 ]
then
    echo -e "\twallet seed: $( echo ${JSON_OUTPUT} | jq '.seed' )"
fi

executeCommand 'getsecret'
if [ $? -eq 0 ]
then
    echo -e "\tsecret: $( echo ${JSON_OUTPUT} | jq '.secret' )"
fi

executeCommand 'importprivkey' 'cVgcgWwQpwzViWmG7dGyvf545ra6AdT4tV29UtQfE8okvPuznFZi'
if [ $? -eq 0 ]
then
    echo -e "\taddress: $( echo ${JSON_OUTPUT} | jq '.address' )"
fi

executeCommand 'listlockunspent'
if [ $? -eq 0 ]
then
    echo -e "\t#lock unspent: $( echo ${JSON_OUTPUT} | jq '.locked_unspent | length' )"
fi

#executeCommand 'lockunspent' ''
executeCommand 'createtransaction' 'pkt1q85n69mzthdxlwutn6dr6f7kwyd9nv8ulasdaqz'
if [ $? -eq 0 ]
then
    echo -e "\ttransaction: $( echo ${JSON_OUTPUT} | jq '.transaction' )"
fi

executeCommand 'dumpprivkey' 'pkt1q0tgwuwcg4tmwegmevdfz3g6tw838upqcq8xt8u'
if [ $? -eq 0 ]
then
    echo -e "\tprivate key: $( echo ${JSON_OUTPUT} | jq '.private_key' )"
fi

executeCommand 'getnewaddress'
if [ $? -eq 0 ]
then
    echo -e "\taddress: $( echo ${JSON_OUTPUT} | jq '.address' )"
fi

executeCommand 'gettransaction' '934095dc4afa8d4b5d43732a96e78e11c0e88defdaab12d946f525e54478938f'
if [ $? -eq 0 ]
then
    echo -e "\tamount: $( echo ${JSON_OUTPUT} | jq '.transaction.amount' )"
    echo -e "\tfee: $( echo ${JSON_OUTPUT} | jq '.transaction.fee' )"
fi

#executeCommand 'sendfrom' ''

#   test commands to manage watch tower
#executeCommand 'wtclient' 'towers'

#   show any eventual error during command execution
if [ -f "${PLDCTL_ERRORS_FILE}" -a $( stat --format='%s' "${PLDCTL_ERRORS_FILE}" ) -gt 0 ]
then
    echo ">>> errors executing some command "
    echo "+++++++++++++++"
    cat ${PLDCTL_ERRORS_FILE}
fi

#   stop pld daemon
stopPldDeamon

rm -rf ${PLDCTL_ERRORS_FILE}
