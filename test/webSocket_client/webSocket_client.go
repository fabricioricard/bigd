////////////////////////////////////////////////////////////////////////////////
//	test/webSocket_client.go  -  Mar-22-2022  -  aldebap
//
//	simple webSocket_client to test pld REST APIs
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/restrpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	//	connect to pld webSocket URI
	var webSocketHost = "ws://localhost:8080"
	const resyncWebSocketURI = "/api/v1/meta/websocket"

	conn, _, err := websocket.DefaultDialer.Dial(webSocketHost+resyncWebSocketURI, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	type webSocketTestCase struct {
		scenario string
		testFunc func(conn *websocket.Conn, verbose bool) error
	}

	var jsonBasedTests []webSocketTestCase = []webSocketTestCase{
		{
			scenario: "JSon DebugLevel command",
			testFunc: sendJSonDebugLevelCommand,
		},
		{
			scenario: "JSon GetInfo command",
			testFunc: sendJSonGetInfoCommand,
		},
		{
			scenario: "JSon WalletBalance command",
			testFunc: sendJSonGetWalletBalance,
		},
		{
			scenario: "JSon wrong command URI",
			testFunc: sendJSonWrongEndpoint,
		},
		{
			scenario: "JSon missing request payload",
			testFunc: sendJSonMissingRequestPayload,
		},
	}

	var protobufBasedTests []webSocketTestCase = []webSocketTestCase{
		{
			scenario: "Protobuf DebugLevel command",
			testFunc: sendProtobufDebugLevelCommand,
		},
		{
			scenario: "Protobuf GetInfo command",
			testFunc: sendProtobufGetInfoCommand,
		},
		{
			scenario: "Protobuf WalletBalance command",
			testFunc: sendProtobufGetWalletBalance,
		},
		{
			scenario: "Protobuf wrong command URI",
			testFunc: sendProtobufWrongEndpoint,
		},
	}

	//	JSon based test cases
	fmt.Printf("[info] testing JSon based messages\n")

	var verbose = false

	for _, testCase := range jsonBasedTests {

		fmt.Printf("[info] test case: %s\n", testCase.scenario)
		err = testCase.testFunc(conn, verbose)
		if err != nil {
			log.Println(err.Error())
		}
	}

	//	protobuf based test cases
	fmt.Printf("[info] testing Protobuf based messages\n")

	for _, testCase := range protobufBasedTests {

		fmt.Printf("[info] test case: %s\n", testCase.scenario)
		err = testCase.testFunc(conn, verbose)
		if err != nil {
			log.Println(err.Error())
		}
	}

	//	send a close socket message to server
	err = conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Client closing"),
	)
	if err != nil {
		log.Fatal("Error closing the Websocket:", err)
	}

	//	sleep so the server can read the close message before the socket be effectively closed
	time.Sleep(1 * time.Second)
}

func sendJSonDebugLevelCommand(conn *websocket.Conn, verbose bool) error {

	var debugLevelReq = &lnrpc.DebugLevelRequest{
		Show:      true,
		LevelSpec: "debug",
	}
	var debugLevelResp = &lnrpc.DebugLevelResponse{}

	err := sendJSonCommand(conn, "/api/v1/meta/debuglevel", debugLevelReq, debugLevelResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> debugLevel response: subsystems: %s\n", debugLevelResp.SubSystems)

	return nil
}

func sendJSonGetInfoCommand(conn *websocket.Conn, verbose bool) error {

	var getInfoResp = &lnrpc.GetInfo2Response{}

	err := sendJSonCommand(conn, "/api/v1/meta/getinfo", nil, getInfoResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> GetInfo response: wallet version: %d\n\tcurrent height: %d\n",
		getInfoResp.Wallet.WalletVersion, getInfoResp.Wallet.CurrentHeight)

	return nil
}

func sendJSonGetWalletBalance(conn *websocket.Conn, verbose bool) error {

	var walletBalanceResp = &lnrpc.WalletBalanceResponse{}

	err := sendJSonCommand(conn, "/api/v1/wallet/balance", nil, walletBalanceResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> walletBalance response: total balance: %d\n\tconfirmed balance: %d\n\tunconfirmed balance: %d\n",
		walletBalanceResp.TotalBalance, walletBalanceResp.ConfirmedBalance, walletBalanceResp.UnconfirmedBalance)

	return nil
}

func sendJSonWrongEndpoint(conn *websocket.Conn, verbose bool) error {

	var walletBalanceResp = &lnrpc.WalletBalanceResponse{}

	err := sendJSonCommand(conn, "/api/v1/wallet/balance/wrongURI", nil, walletBalanceResp, verbose)
	if err != nil {
		return err
	}

	return nil
}

func sendJSonMissingRequestPayload(conn *websocket.Conn, verbose bool) error {

	var transactionDetailsResp = &lnrpc.TransactionDetails{}

	err := sendJSonCommand(conn, "/api/v1/wallet/transaction/query", nil, transactionDetailsResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> transactionDetails response: #transaction: %d\n", len(transactionDetailsResp.Transactions))

	return nil
}

func sendJSonCommand(conn *websocket.Conn, endpoint string, requestPayload interface{}, responsePayload interface{}, verbose bool) error {

	var payload []byte
	var err error

	//	if there's any, marshal the request payload
	if requestPayload != nil {
		payload, err = jsoniter.Marshal(requestPayload)
		if err != nil {
			return errors.New("Fail marshling request payload: " + err.Error())
		}
	}

	//	marshal the request message to a JSon
	var req = restrpc.WebSocketJSonRequest{
		Endpoint:  endpoint,
		RequestId: uuid.New().String(),
		Payload:   payload,
	}

	reqJsonMsg, err := jsoniter.Marshal(req)
	if err != nil {
		return errors.New("Fail marshling webSocker request message: " + err.Error())
	}

	if verbose {
		fmt.Printf("[trace] request message: %s\n", reqJsonMsg)
	}

	//	write the request to the webSocket client
	err = conn.WriteMessage(websocket.TextMessage, reqJsonMsg)
	if err != nil {
		return errors.New("Fail writing message to pld: " + err.Error())
	}

	//	get response message
	messageType, respJsonMsg, err := conn.ReadMessage()
	if err != nil {
		return errors.New("Fail reading message from pld: " + err.Error())
	}

	if messageType != websocket.TextMessage {
		return errors.New("expecting a text based response message from pld")
	}

	if verbose {
		fmt.Printf("[trace] response message: %s\n", respJsonMsg)
	}

	//	unmarshal the response message
	var resp restrpc.WebSocketJSonResponse

	err = jsoniter.Unmarshal(respJsonMsg, &resp)
	if err != nil {
		return errors.New("Fail unmarshling response message: " + err.Error())
	}

	if resp.Error.HttpCode != 0 {
		return errors.New("Error response received from pld: " + resp.Error.Message)
	}

	//	unmarshal the payload within response
	err = jsoniter.Unmarshal([]byte(resp.Payload), responsePayload)
	if err != nil {
		return errors.New("Fail parsing response payload: " + err.Error())
	}

	return nil
}

func sendProtobufDebugLevelCommand(conn *websocket.Conn, verbose bool) error {

	//	create and marshal a debugLevel payload
	var debugLevelReq = &lnrpc.DebugLevelRequest{
		Show:      true,
		LevelSpec: "info",
	}
	var debugLevelResp = &lnrpc.DebugLevelResponse{}

	err := sendProtocCommand(conn, "/api/v1/meta/debuglevel", debugLevelReq, debugLevelResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> debugLevel response: subsystems: %s\n", debugLevelResp.SubSystems)

	return nil
}

func sendProtobufGetInfoCommand(conn *websocket.Conn, verbose bool) error {

	var getInfoResp = &lnrpc.GetInfo2Response{}

	err := sendProtocCommand(conn, "/api/v1/meta/getinfo", nil, getInfoResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> GetInfo response: wallet version: %d\n\tcurrent height: %d\n",
		getInfoResp.Wallet.WalletVersion, getInfoResp.Wallet.CurrentHeight)

	return nil
}

func sendProtobufGetWalletBalance(conn *websocket.Conn, verbose bool) error {

	var walletBalanceResp = &lnrpc.WalletBalanceResponse{}

	err := sendProtocCommand(conn, "/api/v1/wallet/balance", nil, walletBalanceResp, verbose)
	if err != nil {
		return err
	}

	fmt.Printf("--> walletBalance response: total balance: %d\n\tconfirmed balance: %d\n\tunconfirmed balance: %d\n",
		walletBalanceResp.TotalBalance, walletBalanceResp.ConfirmedBalance, walletBalanceResp.UnconfirmedBalance)

	return nil
}

func sendProtobufWrongEndpoint(conn *websocket.Conn, verbose bool) error {

	var walletBalanceResp = &lnrpc.WalletBalanceResponse{}

	err := sendProtocCommand(conn, "/api/v1/wallet/balance/wrongURI", nil, walletBalanceResp, verbose)
	if err != nil {
		return err
	}

	return nil
}

func sendProtocCommand(conn *websocket.Conn, endpoint string, requestPayload proto.Message, responsePayload proto.Message, verbose bool) error {

	var payload []byte
	var err error

	//	if there's any, marshal the request payload
	if requestPayload != nil {
		payload, err = proto.Marshal(requestPayload)
		if err != nil {
			return errors.New("Fail marshling request payload: " + err.Error())
		}
	}

	//	marshal the request message to a Protobuf
	var requestId = uuid.New().String()
	var req = restrpc.WebSocketProtobufRequest{
		Endpoint:  endpoint,
		RequestId: requestId,
	}

	if requestPayload == nil {
		req.Payload = &anypb.Any{
			Value: nil,
		}
	} else {
		req.Payload = &anypb.Any{
			TypeUrl: "github.com/pkt-cash/pktd/lnd/" + reflect.TypeOf(requestPayload).String()[1:],
			Value:   payload,
		}
	}

	requestMessage, err := proto.Marshal(&req)
	if err != nil {
		return errors.New("Fail marshling webSocker request message: " + err.Error())
	}

	if verbose {
		fmt.Printf("[trace] requestId: %s\n", requestId)
		fmt.Printf("[trace] request message: %s\n", requestMessage)
	}

	//	write the request to the webSocket client
	err = conn.WriteMessage(websocket.BinaryMessage, requestMessage)
	if err != nil {
		return errors.New("Fail writing message to pld: " + err.Error())
	}

	//	get response message
	messageType, message, err := conn.ReadMessage()
	if err != nil {
		return errors.New("Fail reading message from pld: " + err.Error())
	}

	if messageType != websocket.BinaryMessage {
		return errors.New("expecting a binary based response message from pld")
	}

	//	unmarshal the response message
	var webSocketResp = &restrpc.WebSocketProtobufResponse{}

	err = proto.Unmarshal(message, webSocketResp)
	if err != nil {
		return errors.New("Fail unmarshaling the response message: " + err.Error())
	}

	if webSocketResp.GetError() != nil {
		return errors.New("Error response received from pld: " + webSocketResp.GetError().GetMessage())
	}

	if verbose {
		fmt.Printf("[trace] Response payload TypeUrl: %s\n", webSocketResp.GetOk().TypeUrl)
		fmt.Printf("[trace] Response payload size: %d\n", len(webSocketResp.GetOk().Value))
	}

	//	unmarshal the payload within response
	err = proto.Unmarshal(webSocketResp.GetOk().GetValue(), responsePayload)
	if err != nil {
		return errors.New("Fail unmarshaling response payload: " + err.Error())
	}

	return nil
}
