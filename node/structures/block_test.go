package structures

import (
	"testing"
)

func TestCopyBlock(t *testing.T) {

}

/*
func TestDeserialiseBlock(t *testing.T) {
	data := []string{
		"63ff8503010105426c6f636b01ff86000106010954696d657374616d70010400010c5472616e73616374696f6e7301ff9200010d50726576426c6f636b48617368010a00010448617368010a0001054e6f6e63650104000106486569676874010400000029ff910201011a5b5d2a7472616e73616374696f6e2e5472616e73616374696f6e01ff920001ff8800002fff87030102ff8800010401024944010a00010356696e01ff8c000104566f757401ff9000010454696d65010400000024ff8b020101155b5d7472616e73616374696f6e2e5458496e70757401ff8c0001ff8a000040ff89030101075458496e70757401ff8a000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff8f020101165b5d7472616e73616374696f6e2e54584f757470757401ff900001ff8e00002fff8d0301010854584f757470757401ff8e000102010556616c7565010800010a5075624b657948617368010a000000fe01d3ff8601fcb574c05a0102012051d1fc2a106541bed7a2db77feb0a33ca5e757d8c825cfb3788ee97e6a2c04ff010101207a9608cc0988e3102bb059c9ad5b776a56449eb14848e932bdec5d7b439e90870240b271ad43bfa59e361b490d625a88eaf8272d909c3776a91070088229b32dc7708c4543d4794a984a005456eb6075eba0c7f00848c029901f56f2a9e2801fe87d0140a4f3a167f4e02eee7cd047b64f1d0016bf7757390e4f343f63dd8cb3a0fd347b99f2599b79a0a1a579d01a1c44ed3a2a0a00435dfec198203da64b82788af72200010201fe084001149c2e5938b3c22260921e455024270c571eeeea360001fe1c400114b7ec2219011d4085cd0066c605cec79eb4b349480001f82a3f9fbea61e36ea00012059a9ae66a0559f3c4055e652c67b794708927709af3b425ca63565f2305eade80101020102286230346134613130383063343436353834366334343035666434396538366565623730306435613400010101fe24400114b7ec2219011d4085cd0066c605cec79eb4b349480000012000000e30142450d800c409dd9dd6bee62162406f7c55d42b1d94a5ed955ec87001200000538d7a5dfdda87e9f6beaa5f30bd46aedabd88f0352d5ffc59e777482e5001fd022efc010200",
	}

	for _, bs := range data {
		b := Block{}

		bsb, err := hex.DecodeString(bs)

		if err != nil {
			t.Fatalf("Error 1: %s", err.Error())
		}

		err = b.DeserializeBlock(bsb)

		if err != nil {
			t.Fatalf("Error 2: %s", err.Error())
		}

		if len(b.Transactions) != 2 {
			t.Fatalf("Number of transactions is wrong. 2 is expected, got %d", len(b.Transactions))
		}
	}
}
*/
