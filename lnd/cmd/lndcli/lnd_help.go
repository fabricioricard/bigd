////////////////////////////////////////////////////////////////////////////////
//	lndcli/lnd_help.go  -  Apr-12-2022  -  aldebap
//
//	Invoke the pld REST help and show it in a fancy way for the CLI
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pkt-cash/pktd/lnd/lnrpc/restrpc/help"
	"github.com/pkt-cash/pktd/lnd/pkthelp"
)

//	show a fancy output for the master help
func getMasterHelp() error {

	var masterHelp = help.RESTMaster_help()

	fmt.Fprintf(os.Stdout, "NAME:\n    %s\n\n", masterHelp.Name)

	fmt.Fprintf(os.Stdout, "DESCRIPTION:\n")
	for _, line := range masterHelp.Description {
		fmt.Fprintf(os.Stdout, "    %s\n", line)
	}
	fmt.Fprintf(os.Stdout, "\n")

	fmt.Fprintf(os.Stdout, "CATEGORY:\n")
	for _, category := range masterHelp.Category {
		showCategory(category, 1)
	}

	return nil
}

//	show a fancy output for the help on a specific category
func showCategory(category *help.RestCommandCategory, level int) {

	const indentation = "  "
	var levelIndentation string

	for i := 1; i <= level; i++ {
		levelIndentation = levelIndentation + indentation
	}

	fmt.Fprintf(os.Stdout, "%s%s:\n\n", levelIndentation, category.Name)

	fmt.Fprintf(os.Stdout, "%sDESCRIPTION:\n", levelIndentation)
	for _, line := range category.Description {
		fmt.Fprintf(os.Stdout, "%s%s%s\n", levelIndentation, indentation, line)
	}
	fmt.Fprintf(os.Stdout, "\n")

	if len(category.Endpoints) > 0 {
		fmt.Fprintf(os.Stdout, "%sCOMMANDS:\n", levelIndentation)
		for _, endpoint := range category.Endpoints {
			var command = endpoint.URI

			if strings.HasPrefix(command, help.URI_prefix+"/") {
				command = command[len(help.URI_prefix)+1:]
			}
			fmt.Fprintf(os.Stdout, "%s%s%s: %s\n", levelIndentation, indentation, command, endpoint.ShortDescription)
		}
		fmt.Fprintf(os.Stdout, "\n")
	}

	if len(category.Subcategory) > 0 {

		fmt.Fprintf(os.Stdout, "%sSUBCATEGORY:\n", levelIndentation)
		for _, subcategory := range category.Subcategory {
			showCategory(subcategory, level+1)
		}
	}
}

//	show a fancy output for the help on a specific command
func getCommandHelp(commandPath string) error {

	//	search the help for the pld command path
	var commandHelp pkthelp.Method
	var commandFound bool

	for _, commandInfo := range help.CommandInfoData {
		if (commandPath[0] == '/' && commandInfo.Path == commandPath) || commandInfo.Path == "/"+commandPath {
			commandHelp = commandInfo.HelpInfo()
			commandFound = true
			break
		}
	}

	if !commandFound {
		return errors.New("invalid pld command: " + commandPath)
	}

	//	show a fancy output for the command help
	fmt.Fprintf(os.Stdout, "NAME:\n  %s: %s\n\n", commandHelp.Name, commandPath)

	fmt.Fprintf(os.Stdout, "DESCRIPTION:\n")
	for _, line := range commandHelp.Description {
		fmt.Fprintf(os.Stdout, "  %s\n", line)
	}
	fmt.Fprintf(os.Stdout, "\n")

	fmt.Fprintf(os.Stdout, "SERVICE:\n  %s\n\n", commandHelp.Service)

	if len(commandHelp.Req.Fields) > 0 {
		fmt.Fprintf(os.Stdout, "OPTIONS:\n")
		for _, requestField := range commandHelp.Req.Fields {
			showField("", &requestField)
		}
	}

	return nil
}

//	show help line on a specific command CLI argument
func showField(fieldHierarchy string, requestField *pkthelp.Field) {

	if len(requestField.Type.Fields) == 0 {

		var commandOption string

		if len(fieldHierarchy) == 0 {
			commandOption = "--" + requestField.Name
		} else {
			commandOption = "--" + fieldHierarchy + "." + requestField.Name
		}

		switch requestField.Type.Name {
		case "bool":

		case "[]byte":
			commandOption += "=value"
		case "string":
			commandOption += "=value"
		case "uint32":
			commandOption += "=value"
		case "int32":
			commandOption += "=value"
		case "uint64":
			commandOption += "=value"
		case "int64":
			commandOption += "=value"
		}

		if len(requestField.Description) == 0 {
			fmt.Fprintf(os.Stdout, "  %s\n", commandOption)
		} else {
			for i, description := range requestField.Description {
				if i == 0 {
					fmt.Fprintf(os.Stdout, "  %s - %s\n", commandOption, description)
				} else {
					fmt.Fprintf(os.Stdout, "    %s\n", description)
				}
			}
		}
	} else {

		for _, requestSubField := range requestField.Type.Fields {
			if len(fieldHierarchy) == 0 {
				showField(requestField.Name, &requestSubField)
			} else {
				showField(fieldHierarchy+"."+requestField.Name, &requestSubField)
			}
		}
	}
}
