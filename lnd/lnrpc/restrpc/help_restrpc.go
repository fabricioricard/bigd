package restrpc

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/mux"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/restrpc/help"
	"github.com/pkt-cash/pktd/lnd/pkthelp"
	"github.com/pkt-cash/pktd/pktlog/log"
)

const (
	helpURI_prefix = "/help"
)

//	convert	pkthelp.type to REST help proto struct
func convertHelpType(t pkthelp.Type) *help.Type {
	resultType := &help.Type{
		Name:        t.Name,
		Description: t.Description,
	}

	//	convert the array of fields
	for _, field := range t.Fields {
		resultType.Fields = append(resultType.Fields, &help.Field{
			Name:        field.Name,
			Description: field.Description,
			Repeated:    field.Repeated,
			Type:        convertHelpType(field.Type),
		})
	}

	return resultType
}

//	marshal rest help response
func marshalHelp(httpResponse http.ResponseWriter, helpInfo pkthelp.Method) er.R {

	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "   ",
	}

	s, err := marshaler.MarshalToString(&help.RestHelpResponse{
		Name:        helpInfo.Name,
		Service:     helpInfo.Service,
		Description: helpInfo.Description,
		Request:     convertHelpType(helpInfo.Req),
		Response:    convertHelpType(helpInfo.Res),
	})
	if err != nil {
		return er.E(err)
	}

	_, err = io.WriteString(httpResponse, s)
	if err != nil {
		return er.E(err)
	}

	return nil
}

//	add the REST master help HTTP handler
func RestHandlersHelp(router *mux.Router) {
	router.HandleFunc("/", getMainHelp)
	router.HandleFunc(help.URI_prefix, getMainHelp)
}

//	get REST master help handler
func getMainHelp(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fill response payload
	if httpRequest.Method != "GET" {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "405 - Request should be a GET because the help endpoint requires no input", http.StatusMethodNotAllowed)
		return
	}
	//	marshal the REST master help
	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "   ",
	}

	s, err := marshaler.MarshalToString(help.RESTMaster_help())
	if err != nil {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "500 - Internal Error", http.StatusInternalServerError)
		log.Errorf("Error replying to request for [%s] from [%s] - error sending error, giving up: [%s]",
			httpRequest.RequestURI, httpRequest.RemoteAddr, err)

		return
	}

	_, err = io.WriteString(httpResponse, s)
	if err != nil {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "500 - Internal Error", http.StatusInternalServerError)
		log.Errorf("Error replying to request for [%s] from [%s] - error sending error, giving up: [%s]",
			httpRequest.RequestURI, httpRequest.RemoteAddr, err)
	}
}
