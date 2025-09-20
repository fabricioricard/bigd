#!  /usr/bin/bash

################################################################################
#   smoke tests for pld REST wallet creation/unlock/changePassphrase endpoints
################################################################################

export  PKT_HOME="$( pwd )"
export  PKT_WALLET_DIR="${HOME}/.pktwallet"
export  PLD="${PKT_HOME}/bin/pld"
export  PLD_OPTIONS=""
export  PLD_OUTPUT_FILE="./pld.out"
export  PLD_PID=
export  PLD_REST_SERVER='http://localhost:8080'
export  REST_ERRORS_FILE='./rest.err'
export  JSON_OUTPUT=''
export  SEED_PASSPHRASE='cGFzc3dvcmQ='
export  WALLET_PASSPHRASE='w4ll3tP@sswd'
export  VERBOSE='false'

#   start pld deamon in background
startPldDeamon() {

    ${PLD} ${PLD_OPTIONS} > ${PLD_OUTPUT_FILE} &

    PLD_PID=$!

    echo "[info] ${PLD} daemon up and running: PID: ${PLD_PID}"

    sleep 10s
}

#   stop pld deamon
stopPldDeamon() {

    executeCommand 'stop' 'GET' '/api/v1/meta/stop'

    kill ${PLD_PID} 2> /dev/null

    sleep 10s
}

#   use curl to execute a command
executeCommand() {
    local COMMAND="${1}"
    local HTTP_METHOD="${2}"
    local URI="${3}"
    local PAYLOAD="${4}"

    RED='\033[0;31m'
    GREEN='\033[0;32m'
    LIGHTGRAY='\033[0;37m'
    NOCOLOR='\033[0m'

    if [ "${HTTP_METHOD}" == "GET" ]
    then
        if [ "${VERBOSE}" == 'true' ]
        then
            echo -e "[trace] ${LIGHTGRAY}curl \"${PLD_REST_SERVER}${URI}\"${NOCOLOR}"
        fi

        JSON_OUTPUT=$( curl "${PLD_REST_SERVER}${URI}" 2>> ${REST_ERRORS_FILE} )
    elif [ "${HTTP_METHOD}" == "POST" ]
    then
        if [ "${VERBOSE}" == 'true' ]
        then
            echo -e "[trace] ${LIGHTGRAY}curl -H \"Content-Type: application/json\" -X POST -d '${PAYLOAD}' \"${PLD_REST_SERVER}${URI}\"${NOCOLOR}"
        fi

        JSON_OUTPUT=$( curl -H "Content-Type: application/json" -X POST -d "${PAYLOAD}" "${PLD_REST_SERVER}${URI}" 2>> ${REST_ERRORS_FILE} )
    else
        echo -e "${RED}error: invalid HTTP method \"${HTTP_METHOD}\"${NOCOLOR}"
        return 1
    fi

    if [ $? -eq 0 ]
    then
        echo -e "[info] ${COMMAND}: ${GREEN}command successfully executed${NOCOLOR}"
    else
        echo -e "${RED}error: fail attempting to run command \"${COMMAND} ${ARGUMENTS}\": $?${NOCOLOR}"
        JSON_OUTPUT=''
        return 1
    fi

    if [ "${VERBOSE}" == 'true' ]
    then
        echo -e "[trace] ${LIGHTGRAY}response: $( echo "${JSON_OUTPUT}" | tr '\t\n' '  ' | tr -s '[:space:]' )${NOCOLOR}"
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
            RESULT=$( echo "${JSON_OUTPUT}" | tr -d '\n' )
        else
            RESULT=$( echo "${JSON_OUTPUT}" | jq "${FILTER}" | tr -d '\n' )
        fi
        echo -e "[info] ${TITLE}: ${RESULT}"
    fi
}

#   use jq to filter results of previously executed command
getCommandResult() {
    local FILTER="${1}"

    if [ ! -z "${JSON_OUTPUT}" ]
    then
        if [ ! -z "${FILTER}" ]
        then
            echo -e "$( echo "${JSON_OUTPUT}" | jq "${FILTER}" | tr -d '\n' )"
        fi
    fi
}

#   send a command to change a seed's paaphrase
changeSeedPassphrase() {
    local SEED_PASSPHRASE="${1}"

    executeCommand 'genseed' 'POST' '/api/v1/util/seed/create' "{ \"seed_passphrase_bin\": \"${SEED_PASSPHRASE}\" }"
    WALLET_SEED="$( getCommandResult '.seed' )"
    showCommandResult 'new wallet enciphered seed' "${WALLET_SEED}"

    executeCommand 'changepassphrase' 'POST' '/api/v1/util/seed/changepassphrase' "{ \"current_seed_passphrase_bin\": \"${SEED_PASSPHRASE}\", \"current_seed\": ${WALLET_SEED}, \"new_seed_passphrase_bin\": \"${SEED_PASSPHRASE}\" }"
    showCommandResult 'result' ''

    echo "[info] changepassphrase: command successfully executed"
}

#   send a command to create a wallet
createWallet() {
    local PASSPHRASE="${1}"

    executeCommand 'genseed' 'POST' '/api/v1/util/seed/create' "{ \"seed_passphrase_bin\": \"${SEED_PASSPHRASE}\" }"
    WALLET_SEED="$( getCommandResult '.seed' )"
    showCommandResult 'new wallet enciphered seed' "${WALLET_SEED}"

    executeCommand 'create_wallet' 'POST' '/api/v1/wallet/create' "{ \"wallet_passphrase\": \"${PASSPHRASE}\", \"wallet_seed\": ${WALLET_SEED}, \"seed_passphrase_bin\": \"${SEED_PASSPHRASE}\" }"
    showCommandResult 'result' ''

    echo "[info] create: command successfully executed"
}

#   send a command to unlock the wallet
unlockWallet() {
    local PASSPHRASE="${1}"

    executeCommand 'unlock_wallet' 'POST' '/api/v1/wallet/unlock' "{ \"wallet_passphrase\": \"${PASSPHRASE}\" }"
    showCommandResult 'result' ''

    echo "[info] unlock: command successfully executed"
}

#   send a command to change wallet's passphrase
changePassphrase() {
    local PASSPHRASE="${1}"
    local NEW_PASSPHRASE="${2}"

    executeCommand 'changePassphrase' 'POST' '/api/v1/wallet/changepassphrase' "{ \"current_passphrase\": \"${PASSPHRASE}\", \"new_passphrase\": \"${NEW_PASSPHRASE}\" }"
    showCommandResult 'result' ''

    echo "[info] changePassphrase: command successfully executed"
}

#   send a command to check wallet's passphrase
checkPassphrase() {
    local PASSPHRASE="${1}"

    executeCommand 'checkPassphrase' 'POST' '/api/v1/wallet/checkpassphrase' "{ \"wallet_passphrase\": \"${PASSPHRASE}\" }"
    showCommandResult 'result' '.validPassphrase'

    echo "[info] checkPassphrase: command successfully executed"
}

#   splash screen
echo ">>>>> Testing pld create wallet REST endpoints"

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

#   save previous wallet and start a new one just for the tests
if [ -d "${PKT_WALLET_DIR}.bkp" ]
then
    YELLOW='\033[0;33m'
    NOCOLOR='\033[0m'

    echo "[warning] ${YELLOW}Wallet backup directory \"${PKT_WALLET_DIR}.bkp\" already exists !${NOCOLOR}"
    exit -1
else
    mv "${PKT_WALLET_DIR}" "${PKT_WALLET_DIR}.bkp"
fi

rm -rf ${PLD_OUTPUT_FILE} ${REST_ERRORS_FILE}

#
#   change seed passphrase test
#

#   start pld deamon, create a seed, change it's passphrase and stop the deamon
echo
echo ">>> scenario 01 - change seed passphrase"

startPldDeamon
changeSeedPassphrase ${SEED_PASSPHRASE}
stopPldDeamon

#
#   check wallet's passphrase before wallet creation test
#

#   start pld deamon, check passphrase stop the deamon
echo
echo ">>> scenario 02 - check wallet's passphrase before wallet creation"

startPldDeamon
checkPassphrase "${WALLET_PASSPHRASE}"
stopPldDeamon

#
#   create wallet test
#

#   start pld deamon, create a wallet and stop the deamon
echo
echo ">>> scenario 03 - create a new wallet"

startPldDeamon
createWallet "${WALLET_PASSPHRASE}"
stopPldDeamon

#
#   unlock wallet test
#

#   start pld deamon, unlock a wallet and stop the deamon
echo
echo ">>> scenario 04 - unlock the wallet"

startPldDeamon
unlockWallet "${WALLET_PASSPHRASE}"
stopPldDeamon

#
#   change wallet passphrase test
#

#   start pld deamon, change wallet's password, unlock it and change it beck before stop the deamon
export  NEW_WALLET_PASSPHRASE='n3wP$sphrz'

echo
echo ">>> scenario 05 - change wallet passphrase and unlock the wallet"

startPldDeamon
changePassphrase "${WALLET_PASSPHRASE}" "${NEW_WALLET_PASSPHRASE}"
unlockWallet "${NEW_WALLET_PASSPHRASE}"
stopPldDeamon

#
#   check the old and new passwords before and after unlock wallet test
#

#   start pld deamon, unlock a wallet again and stop the deamon
echo
echo ">>> scenario 06 - check the old and new passwords before and after unlock wallet with the new passphrase "

startPldDeamon
checkPassphrase "${WALLET_PASSPHRASE}"
checkPassphrase "${NEW_WALLET_PASSPHRASE}"
unlockWallet "${NEW_WALLET_PASSPHRASE}"
checkPassphrase "${WALLET_PASSPHRASE}"
checkPassphrase "${NEW_WALLET_PASSPHRASE}"
stopPldDeamon

rm -rf ${REST_ERRORS_FILE}

#   restore saved wallet
rm -rf "${PKT_WALLET_DIR}"
mv "${PKT_WALLET_DIR}.bkp" "${PKT_WALLET_DIR}"
