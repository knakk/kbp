package marc

import (
	"encoding/json"
	"testing"

	"github.com/knakk/kbp/marc/normarc"
)

func TestCreateRecord(t *testing.T) {
	r := NewRecord().
		SetLeaderPos(
			LeaderRecordStatus, normarc.StatusNy).
		SetLeaderPos(
			LeaderRecordType, normarc.MaterialtypeTekstlig).
		SetLeaderPos(
			LeaderBibliograhicLevel, normarc.KategoriMonografi).
		SetLeaderPos(
			LeaderEncodingLevel, normarc.Nivå1).
		AddControlField(
			NewControlField(Tag001).Set("86000097")).
		AddControlField(
			NewControlField(Tag008).
				SetPos(normarc.PosRegDato, "860708").
				SetPos(normarc.PosSpråk, normarc.SpråkFlerspråklig).
				SetPos(normarc.PosUtgivelsesland, normarc.LandNorge).
				SetPos(normarc.PosOffentligPublikasjon, "a").
				SetPos(normarc.PosKonferanse, "1").
				SetPos(normarc.PosFestskriftkode, "0").
				SetPos(normarc.PosLitterærForm, "0").
				SetPos(normarc.PosMålgruppe, "a")).
		AddDataField(
			NewDataField(Tag020).Add('a', "82-00-09516-9")).
		AddDataField(
			NewDataFieldWithIndicators(Tag041, '0', ' ').Add('a', "nordanswe")).
		AddDataField(
			NewDataFieldWithIndicators(Tag111, ' ', '0').
				Add('c', "Lövånger").
				Add('d', "1979").
				Add('n', "15").
				Add('a', "Nordisk fagkonferanse for historisk metodelære")).
		AddDataField(
			NewDataFieldWithIndicators(Tag245, '1', ' ').
				Add('a', "Teknologi och samhällsförändring").
				Add('b', "föredrag från Nordiska fackkonferansen för historisk metodlära i Lövångers kyrkby 20-23 maj 1979").
				Add('w', "Teknologi och samhællsførændring")).
		AddDataField(
			NewDataField(Tag260).
				Add('a', "Oslo").
				Add('b', "Universitetsforl.").
				Add('c', "cop. 1980")).
		AddDataField(
			NewDataField(Tag300).
				Add('a', "103 s.").
				Add('c', "21 cm")).
		AddDataField(
			NewDataFieldWithIndicators(Tag440, ' ', '0').
				Add('a', "Studier i historisk metode").
				Add('v', "15")).
		AddDataField(
			NewDataField(Tag508).
				Add('a', "Redigert av Jon Elster, med bistand fra Ingrid Åberg")).
		AddDataField(
			NewDataField(Tag546).
				Add('a', "Artikler på norsk, dansk og svensk")).
		AddDataField(
			NewDataFieldWithIndicators(Tag700, '1', '0').
				Add('a', "Elster, Jon").
				Add('e', "red.")).
		AddDataField(
			NewDataFieldWithIndicators(Tag911, ' ', '0').
				Add('a', "Nordiska fackkonferansen för historisk metodlära").
				Add('w', "Nordiska fackkonferansen før historisk metodlæra").
				Add('z', "Nordisk fagkonferanse for historisk metodelære"))

	want := mustDecode(`
*000     nam         1
*00186000097
*008860708         no     a     a10  0 mul
*020  $a82-00-09516-9
*0410 $anordanswe
*111 0$aNordisk fagkonferanse for historisk metodelære$cLövånger$d1979$n15
*2451 $aTeknologi och samhällsförändring$bföredrag från Nordiska fackkonferansen för historisk metodlära i Lövångers kyrkby 20-23 maj 1979$wTeknologi och samhællsførændring
*260  $aOslo$bUniversitetsforl.$ccop. 1980
*300  $a103 s.$c21 cm
*440 0$aStudier i historisk metode$v15
*508  $aRedigert av Jon Elster, med bistand fra Ingrid Åberg
*546  $aArtikler på norsk, dansk og svensk
*70010$aElster, Jon$ered.
*911 0$aNordiska fackkonferansen för historisk metodlära$wNordiska fackkonferansen før historisk metodlæra$zNordisk fagkonferanse for historisk metodelære
^`,
	)

	if !r.Eq(want) {
		t.Errorf("got:\n%+v\nwant:\n%+v", r, want)
	}
}

func TestRecordEquality(t *testing.T) {
	tests := []struct {
		a, b   string
		wantEq bool
	}{
		{ // 0
			a:      "*000\n^",
			b:      "*000\n^",
			wantEq: true,
		},
		{ // 1
			a:      "*000\n^",
			b:      "*000\n*001\n^",
			wantEq: true, // empty controlfield should be disregarded
		},
		{ // 2
			a:      "*000\n*245\n^",
			b:      "*000\n^",
			wantEq: true, // empty datafield should be disregarded
		},
		{ // 3
			a:      "*000\n*245\n^",
			b:      "*000\n*500\n^",
			wantEq: true, // empty datafields should be disregarded
		},
		{ // 4
			a:      "*000\n*245  $aTitle\n^",
			b:      "*000\n*245  $aTitle\n^",
			wantEq: true,
		},
		{ // 5
			a:      "*000\n*245  $aTitle$bSubtitle\n^",
			b:      "*000\n*245  $aTitle$bSubtitle\n^",
			wantEq: true,
		},
		{ // 6
			a:      "*000\n*245  $bSubtitle$aTitle\n^",
			b:      "*000\n*245  $aTitle$bSubtitle\n^",
			wantEq: true, // order of subfields should not matter
		},
		{ // 7
			a:      "*000\n*245  $aTitle$bSubtitle\n^",
			b:      "*000\n*245  $aTitle\n^",
			wantEq: false,
		},
		{ // 8
			a:      "*000\n*245  $aTitle$bSubtitle\n^",
			b:      "*000\n*245 1$aTitle$bSubtitle\n^",
			wantEq: false, // different indicator2
		},
		{ // 9
			a:      "*000\n*24501$aTitle$bSubtitle\n^",
			b:      "*000\n*245 1$aTitle$bSubtitle\n^",
			wantEq: false, // different indicator1
		},
		{ // 10
			a:      "*000\n*650  $aWar$\n*650  $aPeace$\n*650  $aLove$\n^",
			b:      "*000\n*650  $aPeace$\n*650  $aWar$\n*650  $aLove$\n^",
			wantEq: true, // repeated datafields
		},
		{ // 11
			a:      "*000\n*041  $anor$aswe$adan\n^",
			b:      "*000\n*041  $anor$aswe$adan\n^",
			wantEq: true, // repeated subfields
		},
		{ // 12
			a:      "*000\n*041  $anor$aswe$adan\n^",
			b:      "*000\n*041  $aswe$adan$anor\n^",
			wantEq: true, // repeated subfields; order should not matter
		},
	}

	for i, tt := range tests {
		if got := mustDecode(tt.a).Eq(mustDecode(tt.b)); got != tt.wantEq {
			t.Logf("\n\n%v\nshould %s\n\n%v\n", mustDecode(tt.a), boolstr(tt.wantEq), mustDecode(tt.b))
			t.Errorf("equality test %d: got %v; want %v", i, got, tt.wantEq)
		}
	}
}

func boolstr(b bool) string {
	if b {
		return "=="
	}
	return "!="
}

func TestJSONMarshaler(t *testing.T) {
	marc21 := `
	<record xmlns="http://www.loc.gov/MARC21/slim" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.loc.gov/MARC21/slim http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd">
	   <leader>01778cam a2200469 c 4500</leader>
	   <controlfield tag="005">20190308153218.0</controlfield>
	   <controlfield tag="007">ta</controlfield>
	   <controlfield tag="007">cr||||||||||||</controlfield>
	   <controlfield tag="008">140325s2000    no ||||| |||||000|1|nno|</controlfield>
	   <controlfield tag="001">990016186954702201</controlfield>
	   <datafield tag="015" ind1=" " ind2=" ">
		  <subfield code="a">0110150</subfield>
		  <subfield code="2">nbf</subfield>
	   </datafield>
	   <datafield tag="020" ind1=" " ind2=" ">
		  <subfield code="a">8252157319</subfield>
		  <subfield code="q">ib.</subfield>
		  <subfield code="c">Nkr 159.00</subfield>
	   </datafield>
	   <datafield tag="035" ind1=" " ind2=" ">
		  <subfield code="a">001618695-47bibsys_network</subfield>
	   </datafield>
	   <datafield tag="035" ind1=" " ind2=" ">
		  <subfield code="a">(NO-TrBIB)001618695</subfield>
	   </datafield>
	   <datafield tag="035" ind1=" " ind2=" ">
		  <subfield code="a">(NO-TrBIB)112197523</subfield>
	   </datafield>
	   <datafield tag="040" ind1=" " ind2=" ">
		  <subfield code="a">NO-OsNB</subfield>
		  <subfield code="b">nob</subfield>
		  <subfield code="e">katreg</subfield>
	   </datafield>
	   <datafield tag="042" ind1=" " ind2=" ">
		  <subfield code="a">norbibl</subfield>
	   </datafield>
	   <datafield tag="044" ind1=" " ind2=" ">
		  <subfield code="c">no</subfield>
	   </datafield>
	   <datafield tag="080" ind1=" " ind2=" ">
		  <subfield code="a">839.6(08)-1</subfield>
	   </datafield>
	   <datafield tag="082" ind1="7" ind2="4">
		  <subfield code="a">839.821008</subfield>
		  <subfield code="q">NO-OsNB</subfield>
		  <subfield code="2">4/nor</subfield>
	   </datafield>
	   <datafield tag="082" ind1="0" ind2="4">
		  <subfield code="a">839.821008</subfield>
		  <subfield code="2">23</subfield>
	   </datafield>
	   <datafield tag="084" ind1=" " ind2=" ">
		  <subfield code="a">Sc 3</subfield>
		  <subfield code="2">ubtkl/2</subfield>
	   </datafield>
	   <datafield tag="084" ind1=" " ind2=" ">
		  <subfield code="a">S 3c</subfield>
		  <subfield code="2">oosk</subfield>
	   </datafield>
	   <datafield tag="245" ind1="0" ind2="5">
		  <subfield code="a">"Ein blomstervang for din fot-" :</subfield>
		  <subfield code="b">norske kjærleiksdikt</subfield>
		  <subfield code="c">i utval ved Halldis Moren Vesaas</subfield>
	   </datafield>
	   <datafield tag="250" ind1=" " ind2=" ">
		  <subfield code="a">2. utg.</subfield>
	   </datafield>
	   <datafield tag="260" ind1=" " ind2=" ">
		  <subfield code="a">Oslo</subfield>
		  <subfield code="b">Samlaget</subfield>
		  <subfield code="c">2000</subfield>
	   </datafield>
	   <datafield tag="300" ind1=" " ind2=" ">
		  <subfield code="a">108 s.</subfield>
	   </datafield>
	   <datafield tag="500" ind1=" " ind2=" ">
		  <subfield code="a">1. utg. 1994</subfield>
	   </datafield>
	   <datafield tag="506" ind1="0" ind2=" ">
		  <subfield code="f">Åpen tilgang</subfield>
	   </datafield>
	   <datafield tag="533" ind1=" " ind2=" ">
		  <subfield code="a">Elektronisk reproduksjon</subfield>
		  <subfield code="b">[Norge]</subfield>
		  <subfield code="c">Nasjonalbiblioteket Digital</subfield>
		  <subfield code="d">2011-04-07</subfield>
	   </datafield>
	   <datafield tag="650" ind1=" " ind2="7">
		  <subfield code="a">Kjærlighet</subfield>
		  <subfield code="2">humord</subfield>
		  <subfield code="0">(NO-TrBIB)HUME08010</subfield>
	   </datafield>
	   <datafield tag="653" ind1=" " ind2=" ">
		  <subfield code="a">skjønnlitteratur</subfield>
	   </datafield>
	   <datafield tag="653" ind1=" " ind2="6">
		  <subfield code="a">dikt</subfield>
		  <subfield code="a">lyrikk</subfield>
		  <subfield code="a">antologier</subfield>
		  <subfield code="a">antologiar</subfield>
	   </datafield>
	   <datafield tag="655" ind1=" " ind2="4">
		  <subfield code="a">Kjærlighet</subfield>
	   </datafield>
	   <datafield tag="700" ind1="1" ind2=" ">
		  <subfield code="a">Vesaas, Halldis Moren,</subfield>
		  <subfield code="d">1907-1995,</subfield>
		  <subfield code="e">red.</subfield>
		  <subfield code="0">(NO-TrBIB)90061079</subfield>
	   </datafield>
	   <datafield tag="740" ind1="0" ind2=" ">
		  <subfield code="a">Norske kjærleiksdikt</subfield>
	   </datafield>
	   <datafield tag="856" ind1="4" ind2="2">
		  <subfield code="3">Beskrivelse fra forlaget (kort)</subfield>
		  <subfield code="u">http://content.bibsys.no/content/?type=descr_publ_brief&amp;isbn=8252157319</subfield>
	   </datafield>
	   <datafield tag="856" ind1="4" ind2="2">
		  <subfield code="3">Beskrivelse fra Forlagssentralen</subfield>
		  <subfield code="u">http://content.bibsys.no/content/?type=descr_forlagssentr&amp;isbn=8252157319</subfield>
	   </datafield>
	   <datafield tag="856" ind1="4" ind2="2">
		  <subfield code="3">Innholdsfortegnelse</subfield>
		  <subfield code="u">http://content.bibsys.no/content/?type=toc&amp;isbn=8252157319</subfield>
	   </datafield>
	   <datafield tag="856" ind1="4" ind2="0">
		  <subfield code="3">Fulltekst</subfield>
		  <subfield code="u">http://urn.nb.no/URN:NBN:no-nb_digibok_2011031505114</subfield>
		  <subfield code="y">NB digitalisert</subfield>
		  <subfield code="z">Elektronisk reproduksjon. Gratis. For norske IP-adresser</subfield>
	   </datafield>
	   <datafield tag="901" ind1=" " ind2=" ">
		  <subfield code="a">90</subfield>
	   </datafield>
	   <datafield tag="913" ind1=" " ind2=" ">
		  <subfield code="a">Norbok</subfield>
		  <subfield code="b">NB</subfield>
	   </datafield>
	</record>`

	wantJSON := `{"leader":"01778cam a2200469 c 4500","cfields":{"001":"990016186954702201","005":"20190308153218.0","007":"cr||||||||||||","008":"140325s2000    no ||||| |||||000|1|nno|"},"dfields":{"015":[{"ind1":" ","ind2":" ","subfields":{"2":["nbf"],"a":["0110150"]}}],"020":[{"ind1":" ","ind2":" ","subfields":{"a":["8252157319"],"c":["Nkr 159.00"],"q":["ib."]}}],"035":[{"ind1":" ","ind2":" ","subfields":{"a":["001618695-47bibsys_network"]}},{"ind1":" ","ind2":" ","subfields":{"a":["(NO-TrBIB)001618695"]}},{"ind1":" ","ind2":" ","subfields":{"a":["(NO-TrBIB)112197523"]}}],"040":[{"ind1":" ","ind2":" ","subfields":{"a":["NO-OsNB"],"b":["nob"],"e":["katreg"]}}],"042":[{"ind1":" ","ind2":" ","subfields":{"a":["norbibl"]}}],"044":[{"ind1":" ","ind2":" ","subfields":{"c":["no"]}}],"080":[{"ind1":" ","ind2":" ","subfields":{"a":["839.6(08)-1"]}}],"082":[{"ind1":"7","ind2":"4","subfields":{"2":["4/nor"],"a":["839.821008"],"q":["NO-OsNB"]}},{"ind1":"0","ind2":"4","subfields":{"2":["23"],"a":["839.821008"]}}],"084":[{"ind1":" ","ind2":" ","subfields":{"2":["ubtkl/2"],"a":["Sc 3"]}},{"ind1":" ","ind2":" ","subfields":{"2":["oosk"],"a":["S 3c"]}}],"245":[{"ind1":"0","ind2":"5","subfields":{"a":["\"Ein blomstervang for din fot-\" :"],"b":["norske kjærleiksdikt"],"c":["i utval ved Halldis Moren Vesaas"]}}],"250":[{"ind1":" ","ind2":" ","subfields":{"a":["2. utg."]}}],"260":[{"ind1":" ","ind2":" ","subfields":{"a":["Oslo"],"b":["Samlaget"],"c":["2000"]}}],"300":[{"ind1":" ","ind2":" ","subfields":{"a":["108 s."]}}],"500":[{"ind1":" ","ind2":" ","subfields":{"a":["1. utg. 1994"]}}],"506":[{"ind1":"0","ind2":" ","subfields":{"f":["Åpen tilgang"]}}],"533":[{"ind1":" ","ind2":" ","subfields":{"a":["Elektronisk reproduksjon"],"b":["[Norge]"],"c":["Nasjonalbiblioteket Digital"],"d":["2011-04-07"]}}],"650":[{"ind1":" ","ind2":"7","subfields":{"0":["(NO-TrBIB)HUME08010"],"2":["humord"],"a":["Kjærlighet"]}}],"653":[{"ind1":" ","ind2":" ","subfields":{"a":["skjønnlitteratur"]}},{"ind1":" ","ind2":"6","subfields":{"a":["dikt","lyrikk","antologier","antologiar"]}}],"655":[{"ind1":" ","ind2":"4","subfields":{"a":["Kjærlighet"]}}],"700":[{"ind1":"1","ind2":" ","subfields":{"0":["(NO-TrBIB)90061079"],"a":["Vesaas, Halldis Moren,"],"d":["1907-1995,"],"e":["red."]}}],"740":[{"ind1":"0","ind2":" ","subfields":{"a":["Norske kjærleiksdikt"]}}],"856":[{"ind1":"4","ind2":"2","subfields":{"3":["Beskrivelse fra forlaget (kort)"],"u":["http://content.bibsys.no/content/?type=descr_publ_brief\u0026isbn=8252157319"]}},{"ind1":"4","ind2":"2","subfields":{"3":["Beskrivelse fra Forlagssentralen"],"u":["http://content.bibsys.no/content/?type=descr_forlagssentr\u0026isbn=8252157319"]}},{"ind1":"4","ind2":"2","subfields":{"3":["Innholdsfortegnelse"],"u":["http://content.bibsys.no/content/?type=toc\u0026isbn=8252157319"]}},{"ind1":"4","ind2":"0","subfields":{"3":["Fulltekst"],"u":["http://urn.nb.no/URN:NBN:no-nb_digibok_2011031505114"],"y":["NB digitalisert"],"z":["Elektronisk reproduksjon. Gratis. For norske IP-adresser"]}}],"901":[{"ind1":" ","ind2":" ","subfields":{"a":["90"]}}],"913":[{"ind1":" ","ind2":" ","subfields":{"a":["Norbok"],"b":["NB"]}}]}}`

	gotJSON, err := json.Marshal(mustDecode(marc21))
	if err != nil {
		t.Fatal(err)
	}
	if got := string(gotJSON); got != wantJSON {
		t.Errorf("got:\n%v\nwant:\n%v", got, wantJSON)
	}
}
