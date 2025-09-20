////////////////////////////////////////////////////////////////////////////////
//	lndcli/main.go  -  Apr-8-2022  -  aldebap
//
//	Entry point for the pld client using the REST APIs
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pkt-cash/pktd/lnd/lnrpc/restrpc/help"
	"github.com/pkt-cash/pktd/lnd/pkthelp"
)

var (
	pldServer string
)

func main() {
	var showRequestPayload bool

	//	parse command line arguments
	flag.StringVar(&pldServer, "pld_server", "http://localhost:8080", "set the pld server URL")
	flag.BoolVar(&showRequestPayload, "show_req_payload", false, "show the request payload before invoke the pld command")

	flag.Parse()

	//	if a protocol is missing from pld_server, assume HTTP as default
	if !strings.HasPrefix(pldServer, "http://") && !strings.HasPrefix(pldServer, "https://") {
		pldServer = "http://" + pldServer
	}

	var err error

	switch len(flag.Args()) {
	//	print the main help if no arguments are available
	case 0:
		err = getMasterHelp()

	//	one or more arguments means the help + command
	//		or command to be executed followed by arguments to build request payload
	default:
		var command = flag.Args()[0]

		//	if the user wants help on a command
		if command == "help" {

			switch len(flag.Args()) {
			case 1:
				err = getMasterHelp()

			case 2:
				err = getCommandHelp(flag.Args()[1])

			default:
				fmt.Fprintf(os.Stderr, "error: unexpected arguments for help on command %v\n", flag.Args()[2:])
			}
			break
		}

		//	first argument is a pld command followed by arguments to build request payload
		var requestPayload string

		requestPayload, err = formatRequestPayload(command, flag.Args()[1:])
		if err != nil {
			break
		}
		//	if necessary, indent the request payload before show it
		if showRequestPayload {
			var requestPayloadMap map[string]interface{}

			err = json.Unmarshal([]byte(requestPayload), &requestPayloadMap)
			if err != nil {
			} else {
				prettyRequestPayload, err := json.MarshalIndent(requestPayloadMap, "", "    ")
				if err != nil {
				} else {
					fmt.Fprintf(os.Stdout, "[trace]: request payload: %s\n", string(prettyRequestPayload))
				}
			}
		}

		//	send the request payload to pld
		err = executeCommand(command, requestPayload)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}

//	based on pld's command path, parse the CLI arguments to build the request payload
func formatRequestPayload(commandPath string, arguments []string) (string, error) {

	//	search all pld commands for help on command path
	var commandHelp pkthelp.Method
	var commandFound bool
	var allowGet bool

	for _, commandInfo := range help.CommandInfoData {
		if (commandPath[0] == '/' && commandInfo.Path == commandPath) || commandInfo.Path == "/"+commandPath {
			if commandInfo.HelpInfo != nil {
				commandHelp = commandInfo.HelpInfo()
				commandFound = true
				allowGet = commandInfo.AllowGet
				break
			}
		}
	}

	if !commandFound {
		return "", errors.New("invalid pld command: " + commandPath)
	}

	//	build request payload based on request's help info hierarchy
	var parsedArgument []bool = make([]bool, len(arguments))
	var requestPayload string

	if len(commandHelp.Req.Fields) > 0 {
		for _, requestField := range commandHelp.Req.Fields {

			formattedField, err := formatRequestField("", &requestField, arguments, &parsedArgument)
			if err != nil {
				return "", errors.New("error parsing arguments: " + err.Error())
			}

			if len(formattedField) > 0 {
				if len(requestPayload) > 0 {
					requestPayload += ", "
				}
				requestPayload += formattedField
			}
		}
	}
	if len(requestPayload) == 0 && allowGet {
		requestPayload = ""
	} else {
		requestPayload = "{ " + requestPayload + " }"
	}

	//	check if there are invalid arguments (not parsed)
	if len(arguments) > 0 {
		for i := 0; i < len(parsedArgument); i++ {
			if !parsedArgument[i] {
				return "", errors.New("invalid command argument: " + arguments[i])
			}
		}
	}

	return requestPayload, nil
}

//	check if there's a CLI argument for a specific payload field,
//	in which case, returs the field formatted accordingly to it's data type
func formatRequestField(fieldHierarchy string, requestField *pkthelp.Field, arguments []string, parsedArgument *[]bool) (string, error) {

	var formattedField string

	if len(requestField.Type.Fields) == 0 {

		var commandOption string

		if len(fieldHierarchy) == 0 {
			commandOption = "--" + requestField.Name
		} else {
			commandOption = "--" + fieldHierarchy + "." + requestField.Name
		}

		switch requestField.Type.Name {
		case "bool":
			for i, argument := range arguments {
				if argument == commandOption {
					formattedField += "\"" + requestField.Name + "\": true"
					(*parsedArgument)[i] = true
				}
			}

		case "[]byte":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": \"" + argument[len(commandOption):] + "\""
					(*parsedArgument)[i] = true
				}
			}

		case "string":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					if !requestField.Repeated {
						formattedField += "\"" + requestField.Name + "\": \"" + argument[len(commandOption):] + "\""
					} else {
						//	make sure the array is delimited by square brackets
						arrayArgument := strings.TrimSpace(argument[len(commandOption):])

						if arrayArgument[0] != '[' || arrayArgument[len(arrayArgument)-1] != ']' {
							return "", errors.New("array argument must be delimitted by square brackets: " + arrayArgument)
						}

						//	each string in the array is comma separated
						var arrayOfStrings string

						for _, stringValue := range strings.Split(arrayArgument[1:len(arrayArgument)-1], ",") {
							stringValue = strings.TrimSpace(stringValue)

							//	make sure the string element is delimited by double quotes
							if stringValue[0] != '"' || stringValue[len(stringValue)-1] != '"' {
								return "", errors.New("array element must be delimitted by double quotes: " + stringValue)
							}

							if len(arrayOfStrings) > 0 {
								arrayOfStrings += ", "
							}
							arrayOfStrings += stringValue
						}

						formattedField += "\"" + requestField.Name + "\": [ " + arrayOfStrings + " ]"
					}
					(*parsedArgument)[i] = true
				}
			}

		//	TODO: to make sure that for integer types pld doen't have arrays (Repeated == true)
		case "uint32":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": " + argument[len(commandOption):]
					(*parsedArgument)[i] = true
				}
			}

		case "int32":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": " + argument[len(commandOption):]
					(*parsedArgument)[i] = true
				}
			}

		case "uint64":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": " + argument[len(commandOption):]
					(*parsedArgument)[i] = true
				}
			}

		case "int64":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": " + argument[len(commandOption):]
					(*parsedArgument)[i] = true
				}
			}

		case "float64":
			commandOption += "="
			for i, argument := range arguments {
				if strings.HasPrefix(argument, commandOption) {
					formattedField += "\"" + requestField.Name + "\": " + argument[len(commandOption):]
					(*parsedArgument)[i] = true
				}
			}

		//	enums are formatted as it's name, because pld is able to unmarshall them
		case "ENUM_VARIENT":
			for i, argument := range arguments {
				if argument == commandOption {
					formattedField += "\"" + requestField.Name + "\""
					(*parsedArgument)[i] = true
				}
			}
		}
	} else {

		//	field composed of sub-fields
		var formattedSubFields string

		for _, requestSubField := range requestField.Type.Fields {
			var formattedSubField string
			var err error

			if len(fieldHierarchy) == 0 {
				formattedSubField, err = formatRequestField(requestField.Name, &requestSubField, arguments, parsedArgument)
			} else {
				formattedSubField, err = formatRequestField(fieldHierarchy+"."+requestField.Name, &requestSubField, arguments, parsedArgument)
			}
			if err != nil {
				return "", err
			}

			if len(formattedSubField) > 0 {
				if len(formattedSubFields) > 0 {
					formattedSubFields += ", "
				}
				formattedSubFields += formattedSubField
			}
		}
		if len(formattedSubFields) > 0 {
			if !requestField.Repeated {
				formattedField += "\"" + requestField.Name + "\": { " + formattedSubFields + " }"
			} else {
				formattedField += "\"" + requestField.Name + "\": [ " + formattedSubFields + " ]"
			}
		}
	}

	return formattedField, nil
}

//	invoke pld's REST endpoint and try to parse error messages eventually returned by the server
func executeCommand(command string, payload string) error {

	var response *http.Response
	var err error

	commandURI := pldServer + help.URI_prefix + "/" + command

	//	if there's no payload, use HTTP GET method to invoke pld command, otherwise use POST method
	if len(payload) == 0 {
		response, err = http.Get(commandURI)
		if err != nil {
			return errors.New("fail executing pld command: " + err.Error())
		}
	} else {
		response, err = http.Post(commandURI, "application/json", strings.NewReader(payload))
		if err != nil {
			return errors.New("fail executing pld command: " + err.Error())
		}
	}
	defer response.Body.Close()

	responsePayload, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail reading command response payload from pld server: %s", err)
		panic(-1)
	}
	err = checkForServerError(responsePayload)
	if err != nil {
		return errors.New(err.Error() + "\nTry \"pldctl help " + command + "\" for more informaton on this command")
	}

	fmt.Fprintf(os.Stdout, "%s\n", responsePayload)

	return nil
}

type pldErrorResponse struct {
	Message string   `json:"message,omitempty"`
	Stack   []string `json:"stack,omitempty"`
}

//	parse a response payload to check if it indicates an error messages returned by the server
func checkForServerError(responsePayload []byte) error {
	var errorResponse pldErrorResponse

	err := json.Unmarshal(responsePayload, &errorResponse)
	if err == nil {
		if len(errorResponse.Message) > 0 && len(errorResponse.Stack) > 0 {
			var stackTrace string

			//	format the stack trance for output
			for _, step := range errorResponse.Stack {
				stackTrace += step + "\n"
			}

			return errors.New("pld returned an error message: " + errorResponse.Message + "\n\npld stack trace: " + stackTrace)
		}
	}

	return nil
}
