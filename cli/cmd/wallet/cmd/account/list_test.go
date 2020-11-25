package account_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/eth2-key-manager/cli/cmd"
	"github.com/bloxapp/eth2-key-manager/cli/util/printer"
)

func TestAccountList(t *testing.T) {
	t.Run("Successfully list accounts", func(t *testing.T) {
		var output bytes.Buffer
		cmd.ResultPrinter = printer.New(&output)
		cmd.RootCmd.SetArgs([]string{
			"wallet",
			"account",
			"list",
			"--storage=7b226163636f756e7473223a2237623232333533313338333336353635363633363264333533363335363232643334333436353330326433393337333336313264363536333634333636323332333633383332333833303339323233613762323236323631373336353431363336333666373536653734353036313734363832323361323232663330323232633232363936343232336132323335333133383333363536353636333632643335333633353632326433343334363533303264333933373333363132643635363336343336363233323336333833323338333033393232326332323665363136643635323233613232363136333633366637353665373432643330323232633232373636313663363936343631373436393666366534623635373932323361376232323639363432323361323233383336363333353339363333363633326433373635333736313264333436323631333932643631363633323333326433323336363436353339363436313334363136343333333032323263323237303631373436383232336132323664326633313332333333383331326633333336333033303266333032663330326633303232326332323730373236393736346236353739323233613232333533323635333033363335333933373336333233353336333133383333363233373331363136313335333533393330333636353336333436313631363633313336333633343332333533313338363433343633333633333631363333363636333736313334363336313338333933303335333533363331333336353336333332323764326332323737363937343638363437323631373736313663353037353632346236353739323233613232363133303632333933333332333436343631333836313338363133363339333636333335333333393335333036353339333833343634363533323335363233323339333936333331333233333634333133373632363136323339333733323635363336313331363136333332363333363337333433393336333436333339363633383331333733303334333736323633333633303334333836353636333033373330333536343337363536333336363636313635333636343335363436313336323237643764222c2268696768657374417474223a223762323237333663366637343232336133313263323236333666366436643639373437343635363535663639366536343635373832323361333132633232363236353631363336663665356636323663366636333662356637323666366637343232336132323531353133643364323232633232373336663735373236333635323233613762323236353730366636333638323233613331326332323732366636663734323233613232353135313364336432323764326332323734363137323637363537343232336137623232363537303666363336383232336133323263323237323666366637343232336132323531353133643364323237643764222c226e6574776f726b223a223664363136393665366536353734222c2270726f706f73616c4d656d6f7279223a22376232323339333533303338333733313338333233393333333736363336333933383332363136353339333936363339363233303336363236343331333133363636333433363333363633343331333433353331333333303333333236353333333336313333363433313337333536343339333633363332363536343634363633313336333233313330333136363633363633363633363133323631333936363635363436313634363536343337333436323338333033343337363333353634363336363566333132323361376232323733366336663734323233613331326332323730373236663730366637333635373235663639366536343635373832323361333132633232373036313732363536653734356637323666366637343232336132323531353133643364323232633232373337343631373436353566373236663666373432323361323235313531336433643232326332323632366636343739356637323666366637343232336132323531353133643364323237643764222c2277616c6c6574223a223762323236393634323233613232333336363331333536343339333133373264333733373336363232643334333033393335326433393338333336353264333436363334333436363330363633303336363633343631323232633232363936653634363537383464363137303730363537323232336137623232333933353330333833373331333833323339333333373636333633393338333236313635333933393636333936323330333636323634333133313336363633343336333336363334333133343335333133333330333333323635333333333631333336343331333733353634333933363336333236353634363436363331333633323331333033313636363336363336363336313332363133393636363536343631363436353634333733343632333833303334333736333335363436333636323233613232333533313338333336353635363633363264333533363335363232643334333436353330326433393337333336313264363536333634333636323332333633383332333833303339323237643263323237343739373036353232336132323438343432323764227d",
		})
		err := cmd.RootCmd.Execute()
		actualOutput := output.String()
		require.NotNil(t, actualOutput)
		require.NoError(t, err)
	})

	t.Run("Fail to JSON un-marshal", func(t *testing.T) {
		var output bytes.Buffer
		cmd.ResultPrinter = printer.New(&output)
		cmd.RootCmd.SetArgs([]string{
			"wallet",
			"account",
			"list",
			"--storage=7b226163636f756e7473223a2237623764222c226174744d656d6f7279223a2237623764222c2270726f706f",
		})
		err := cmd.RootCmd.Execute()
		require.Error(t, err)
		require.EqualError(t, err, "failed to JSON un-marshal storage: unexpected end of JSON input")
	})
}
