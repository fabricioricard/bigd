#!/usr/bin/env sh

#   check for the presence of --protoc flag and remove it from ARGS passed to the compile step
BUILD_ARGS=""
PROTOC="false"

while [ true ]
do
    ARG=${1}

    if [ -z "${ARG}" ]
    then
        break
    fi

    if [ "${ARG}" == "--protoc" ]
    then
        PROTOC="true"
    else
        BUILD_ARGS="${BUILD_ARGS} \"${ARG}\""
    fi

    shift
done

#   when requested, generate stubs from gRPC proto files
if [ "${PROTOC}" == "true" ]
then
    cd lnd/lnrpc
    ./gen_protos.sh

    if [ $? -ne 0 ]
    then
        echo "${0}: fail attempting to generate stubs from gRPC proto files"
        exit 1
    fi
    cd ../..
fi

#   build source code
eval go run ./contrib/build/build.go ${BUILD_ARGS}
