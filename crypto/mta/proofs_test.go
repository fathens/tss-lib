// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package mta

import (
	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/crypto/paillier"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"math/big"
	"testing"
)

func TestParamN(t *testing.T) {
	curve := tss.EC()
	n := curve.Params().N
	t.Logf("N: %v", n)
}

func TestPointMul(t *testing.T) {
	ec := tss.EC()
	k, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494336", 10)
	point := crypto.ScalarBaseMult(ec, k)
	x := point.X()
	y := point.Y()
	t.Logf("x: %v", x)
	t.Logf("y: %v", y)
}

func TestNewAndVerify(t *testing.T) {
	ec := tss.EC()
	session := []byte("session")
	newBigInt := func(s string) *big.Int {
		b, _ := new(big.Int).SetString(s, 10)
		return b
	}
	pk := paillier.PublicKey{N: newBigInt("24261589004465272731249327803101071103792958802723210985329987798404286346832084145379821062390362511363592469465551036623716847742366801146734075048032103335288069058682991894283824242201941990356426107676283864045055933692254494542670287211729114434164797753179698775293789682407290565454766674083988564005545698061330433346903087394170496980007835487898180646656591606889559906608994560791838505934557241218559133947318663171630480041559848077544934775873477027858500504325199526168809286442772506257956278830238227308912498225476420365415416389640465904626483958374392724811019844436417721814708231141829835974781")}
	ntilde := newBigInt("25107490776052945575790163886980744121852075793230702092031092910315419013111724585107741342302647097816029689069156500419649067226989207335403141846585589456214707140363806918024254341805807847344462552372749802373561411623464018306841140152736878126807643286464707464144491205717529334857128642937311664356950670200785184493082292988908234459722618881044613550904554507333793627844968327344517418351075665978629614435510466378211576459017353838583039397930178040557511540818370302033808216608330168909665648805527673068950251148153088673193641290377199021831923470431364077200419352774733381328839199321622201645277")
	h1 := newBigInt("947268510305326446073634507724913447936734171636912400557401318775427643035322780043344044871778218536295489345747992085537349997385753459769909944243608187249295932620582767525243046024431872134558350124222211815956076009495579000118546531817489783543950708796804986346442485595844139040615169351977594594085460608932273701244091036215057114383266995365365226626217411088112095883376367775475107954293975266374705057036496941779873360807750450088301028537780564210964889218799820623451941121168857520561736570209171665676631521362739174866629364755585577716299287494251706261472512421959632149833106509542229972234")
	h2 := newBigInt("369382535766024782757053511943484023707590301248858510505619543451105355366349475321600848828578055383112252081262740450957242693258711711573898608872557215737850380375149487180022863563616178163440683814662347260503803753150609907077552201623376131096249150783552367189222999632342102603491398593162398739317344334427947844029843540621897547082716967267285286086227255034044222917612280937408214149645699005643727644027239999997789724357422423935120674874708262799420509411969660535187315093553065000790565517535769427338692918882249946664488170641583406635227373502217028982923125561321182147198392699754510926843")
	c1 := newBigInt("400188980994774655609968091936620060124194780418395642324023940691239945255312247610992811612431249936152054069340910452935275205487424854634246828482246352144337182419040538212228852534225304174368343947023564538801615828687299316803241803369288485506946718273947431531973813858828408905751899378227757543579871688504974553109136356994904735037599168566303793641249248025873904193154372979358251360874998196205783452661730986829472793946270963339861788192770799371305531044326953375551834910629477333425892458792662376368650895554603125094896843040198744699366418823291775347326054231834531358977230768426742453712318682737791792820647755652215856663778772847375042277205326803580443908810548115875614952826259458257258493697853793646331037887120817707513503571379434818731866933591449738884942230286964560975261495708985005280778505293251103037083314378974282722058952927812473759779022197254469017388422994057180505540627929700170170320971992028762205724059261013866722334225838238320698563687792441899338330349607215301028623539836617463866851633411657797882767517790720192752726498897268976870434440522723184905076615752830635734406782805309804373897258628489152609956916649503127542627206201799628788404190292045079845695539005")
	c2 := newBigInt("98573216923670029482821115497694158997701755867518675487809563002262969374429334440082681941600401097934556988790183729523580720309003894174143686603097011750180621654263582730257639883377054326574168737200415168109016934596056214113130498173847556339154501370028792265841294296659920063110922540964643954044560558816019775357577881424654115472796241983478183026500401778150566430093557176244961750759228035574132175895475959282423213068981255863315840431505638769681064902513786237796323681661253900688171694813024684962553099575687213643953866670007212555321378763872102032075258716404994287465655087264659390259026908282688806618224771400409708056097235104949067246227149153871512407333953827944360034115190003167379402652418218088558623986757262384661381181534752894246084721691776580316855378370830325218479465488778739115168862573546202079589373938148889660862603212262415001010351315286539400994916127576235285929134852940474957058813438546566643450052998120855803755264759031202653355459212026268818145535080517620144856445898319577455911245421862936535731411912720397377403186779885488966511055963605408874871060573451766442775230526695777998996381124371051298228555236557500794840592492822277944138084212794148491560062194")
	x := newBigInt("20872918044507599492457919043155243332921084283187084425223278835323345316707")
	y := newBigInt("13365803136768827314472359487832734942388024092573771842356564480371273430093788764148177649590147601079991374084731619027142332042824767083157035176876340102099960663493606582967420224789162175044870543176867902664052695037517942038248013816833256471600451521497813841165600589598039533465498062819816581880369639821378040565954040957506700021927229986652022091615738614219190677069635")
	r := newBigInt("21617745718658475423967690137755260260751176145758099656379627827529447465873941447558271433560154929175996297478257677991009360164924302078973592231666628498979499849260499718709195852482305685149481840993520270133375205592174152961925586484451711506799685098224993369966535853381357250207917204040666031452544422143667110446801956906625901747974275844240807779974858102640852085251773464817602345912999959784397499248630936345359919089789880396555926854044371853388456300266725449255410786594189679918897998146392256946271699268332250404359035739120314312721689120835606374439344529292527051773663149198509109589298")
	point, err := crypto.NewECPoint(ec,
		newBigInt("104639075809233846840558005847879026027812042860709851555234497260916054215660"),
		newBigInt("75893047884391164204942577163850331696590636042088436736848678939010438699885"),
	)
	if err != nil {
		t.Fatal(err)
	}
	proofWc, err := ProveBobWC(session, ec, &pk, ntilde, h1, h2, c1, c2, x, y, r, point)
	if err != nil {
		t.Fatal(err)
	}
	isVerified := proofWc.Verify(session, ec, &pk, ntilde, h1, h2, c1, c2, point)
	if !isVerified {
		t.Fatal("proofWc.Verify() returned false")
	}
}
