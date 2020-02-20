package view

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/selassje/passwordGenerator/passwordGenerator"
	"github.com/zserge/webview"
)

const (
	windowWidth  = 270
	windowHeight = 250
)

var indexHTML = `
<!doctype html>
<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<style>
			input[type="checkbox"] {
				hidden:;
				color: blue;
			}
		</style>
	</head>
	<body>
		<script>
			function getCheckbox(id) {
				if (document.getElementById(id).checked)
					return "True"
				return "False"
			}

			function getSettings() {
				return document.getElementById("length").value + ":" +
					   getCheckbox("includeLowerCase") + ":" +
					   getCheckbox("includeUpperCase") + ":" +
					   getCheckbox("includeDigits")    + ":" +
					   getCheckbox("includeSymbols")
			}
		</script>

		<div class="inputLength">
			<p>Length</p>
			<input id="length" type="text" value="10" style="width: 140px;"></input>
		</div>
		<div class="inputCheckBox">
			<input class="inputCheckBox2" type="checkbox" id="includeLowerCase" name="includeLowerCase" value="LowerCaseAllowed" checked>
			<label for="includeLowerCase">Include lower case</label><br>
		</div>
		<div class="inputCheckBox">
			<input type="checkbox" id="includeUpperCase" name="includeUpperCase" value="UpperCaseAllowed" checked>
			<label for="includeUpperCase">Include upper case</label><br>
		</div>
		<div class="inputCheckBox">
			<input type="checkbox" id="includeDigits" name="includeDigits" value="DigitsAllowed" checked>
			<label for="includeDigits">Include digits</label><br>
		</div>
		<div class="inputCheckBox">
			<input type="checkbox" id="includeSymbols" name="includeSymbols" value="SymbolsAllowed" checked>
			<label for="includeSymbols">Include symbols</label><br>
		</div>
		<button onclick="external.invoke('generate:'+ getSettings())">
			Generate Password
		</button>
		<p id ="output"></p>		
	</body>
</html>
`

func updateField(w webview.WebView, e string, v string) {
	jsCode := `document.getElementById("` + e + `").innerHTML ="` + e + `: ` + v + `"`
	w.Eval(jsCode)
}

func stringToBool(s string) bool {
	if s == "True" {
		return true
	}
	return false
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case strings.HasPrefix(data, "generate:"):
		settingsRaw := strings.TrimPrefix(data, "generate:")
		settingRawSlice := strings.Split(settingsRaw,":")
		var s passwordGenerator.Settings
		var password string
		length, err := strconv.Atoi(settingRawSlice[0])
		if err == nil {
			s.Length = length
			s.IncludeLowerCaseLetters = stringToBool(settingRawSlice[1])
			s.IncludeUpperCaseLetters = stringToBool(settingRawSlice[2])
			s.IncludeDigits = stringToBool(settingRawSlice[3])
			s.IncludeSymbols = stringToBool(settingRawSlice[4])
			err = passwordGenerator.ValidateSettings(&s)
			if err == nil {
				password = string(passwordGenerator.GeneratePassword(&s))
			}
		}
		if err != nil {
			updateField(w,"output", err.Error())
		} else {
			updateField(w,"output", password)
		}	
	}
}

func RunGui() {
	w := webview.New(webview.Settings{
		Width:                  windowWidth,
		Height:                 windowHeight,
		Title:                  "PasswordGenerator",
		Resizable:              false,
		URL:                    "data:text/html," + url.PathEscape(indexHTML),
		ExternalInvokeCallback: handleRPC,
		Debug:                  true,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()

}
