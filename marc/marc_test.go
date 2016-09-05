package marc

import (
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
