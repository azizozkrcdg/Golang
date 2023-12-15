package main

import "fmt"

func main() {
	horoscopeFeatures()

}

func horoscope() string {
	var month, day int
	var horoscope string
	var isError = false
	fmt.Print("Doğduğunuz Ay (Örn. 6): ")
	fmt.Scan(&month)
	fmt.Print("Doğduğunuz Gün (Örn. 10): ")
	fmt.Scan(&day)

	if 1 <= month && month <= 12 {
		switch month {
		case 1:
			if 1 <= day && day <= 31 {
				if day < 22 {
					horoscope = "Oğlak"
				} else {
					horoscope = "Kova"
				}
			} else {
				isError = true
			}
		case 2:
			if 1 <= day && day <= 28 {
				if day < 20 {
					horoscope = "Kova"
				} else {
					horoscope = "Balık"
				}
			} else {
				isError = true
			}
		case 3:
			if 1 <= day && day <= 31 {
				if day < 21 {
					horoscope = "Balık"
				} else {
					horoscope = "Koç"
				}
			} else {
				isError = true
			}
		case 4:
			if 1 <= day && day <= 30 {
				if day < 21 {
					horoscope = "Koç"
				} else {
					horoscope = "Boğa"
				}
			} else {
				isError = true
			}
		case 5:
			if 1 <= day && day <= 31 {
				if day < 22 {
					horoscope = "Boğa"
				} else {
					horoscope = "İkizler"
				}
			} else {
				isError = true
			}
		case 6:
			if 1 <= day && day <= 30 {
				if day < 23 {
					horoscope = "İkizler"
				} else {
					horoscope = "Yengeç"
				}
			} else {
				isError = true
			}
		case 7:
			if 1 <= day && day <= 31 {
				if day < 23 {
					horoscope = "Yengeç"
				} else {
					horoscope = "Aslan"
				}
			} else {
				isError = true
			}
		case 8:
			if 1 <= day && day <= 31 {
				if day < 23 {
					horoscope = "Aslan"
				} else {
					horoscope = "Başak"
				}
			} else {
				isError = true
			}
		case 9:
			if 1 <= day && day <= 30 {
				if day < 23 {
					horoscope = "Başak"
				} else {
					horoscope = "Terazi"
				}
			} else {
				isError = true
			}
		case 10:
			if 1 <= day && day <= 31 {
				if day < 23 {
					horoscope = "Terazi"
				} else {
					horoscope = "Akrep"
				}
			} else {
				isError = true
			}
		case 11:
			if 1 <= day && day <= 30 {
				if day < 22 {
					horoscope = "Akrep"
				} else {
					horoscope = "Yay"
				}
			} else {
				isError = true
			}
		case 12:
			if 1 <= day && day <= 31 {
				if day < 22 {
					horoscope = "Yay"
				} else {
					horoscope = "Oğlak"
				}
			} else {
				isError = true
			}
		default:
			isError = true
		}
	} else {
		isError = true
	}

	if isError {
		fmt.Println("Lütfen geçerli seçim bir değer girin!")
	} else {
		fmt.Printf("Burcunuz : %s", horoscope)
	}
	return horoscope
}

func horoscopeFeatures() {
	var horoscope = horoscope()
	var personal_characteristics string
	var physical_features string
	var group string
	var manager_planet string
	var color string
	var lucky_number int
	var lucky_day string
	var anti_horoscope string

	switch horoscope {
	case "Oğlak":
		personal_characteristics = "Geleceği ayrıntılı biçimde planlamaya çalışan Oğlak'lar bu özellikleri nedeniyle sık sık kuruntulara kapılarak, depresyona girerler. İşleri ile aşırı meşgul olduklarından, insanlarla zor ilişki kurarlar. Fakat, hiçbir zaman kendilerini yalnız hissetmezler. Bu nedenle onları tanımlayan sözcük 'Kullanırım' dır. Oğlaklar ciddilikleri, tutuculukları ve güçlü iradeleri ile tanımlanırlar. Çalışkanlıkları ile kolaylıkla başkalarının saygısını kazanırlar. Gerçek bir Oğlak'ın iki temel özelliği vardır. Güvenilirlik ve dürüstlük. Çok gelişmiş bir görev duyguları vardır. Bazen başkalarının başarısızlıklarını abartırlar. Oğlaklar arkadaşları ile ilgili çok iyi sırdaş olurlar, para konusunda dikkatli olmalarına rağmen, eli açık insanlardır. Geleneklere bağlı kişiler olduklarından, duygusal özgürlüğü anlamakta güçlük çekerler. Bu nedenle kendilerinden yaşça küçük kişilerle bazen sürtüşmeleri olur. Sürekli somut konularla uğraşmayı severler, doğada bulunan olaylardan örneklemelerle işlerini başarılı bir şekilde yürütürler. Bu yüzden ruhsal olarak da doyumlu kişilerdir. Olayları organize ederlerken rahatlıklarını gözetirler ve kendilerinin zevklerine uygun olmasına önem verirler. Bu özellikleri evlerine de yansır."
		physical_features = "Kısa boy, ince yapı, uzun ince bir yüz. Uzun belirgin bir çene. Saçlar genelde siyahtır. Beden dar ve uzun bacaklara sahiptirler. Dizleri genelde çıkıntılıdır. Güzel bir yüz yapıları vardır. Fotojenik oldukları için model dünyasında sık sık Oğlaklara rastlarız. Ciltleri yıllar geçtikçe daha genç olur Soğuk görüntüleri onları ulaşılmaz bir hava estirse de, kendilerine güvenen bir tablo da, ortaya çıkarır. Kendinden emin tavırlarıyla herkesi kendine hayran bırakmasını bilir."
		group = "Toprak/kadınsı/öncü"
		manager_planet = "Satürn"
		color = "Kurşuni, koyu kahverengi, nefti yeşil"
		lucky_number = 8
		lucky_day = "Cumartesi"
		anti_horoscope = "Yengeç"
	case "Kova":
		personal_characteristics = "Hayal güçleri sınırsız olmakla birlikte, düşünceleri bulundukları anın ötesinde, akılcı ve sezgiseldir. Kova 'lar dik kafalıdırlar. Kendilerini dinleyenlerin ne demek istediklerini anlamadıklarını sanırlar. Kovaları tanımlayan sözcük 'Biliyorum' dur. Kova burcu insanları sevecen tavırları ile tanınırlar. Bu kişiler bencil değildirler. Irk, cinsiyet ve sosyal durumuna bakmaksızın, herkesin ayni olanaklara sahip olmasını isterler. Modern görünüşlerine karşın, inatçı ve sabit fikirli olurlar. Onlara yaklaşmak çok zordur, çünkü ne kadar dostça davranırlarsa davransınlar, arada her zaman bir mesafe bırakırlar. Kişisel özgürlükleri onlar için o denli önemlidir ki, bu yüzden en yakın ilişkilerini kesip atabileceği gibi, yine özgürlükleri adına her türlü özveride bulunurlar. Bu yüzden Kovalara aile yaşamı biraz zor gelir. Çoğu zaman yeni bir şeyler keşfetmek için uğraşıda bulunurlar. Belli ilkeleri sonuna dek savunmaları onları hiç rahatsız etmez."
		physical_features = "Orta ve uzun bir boy, düzgün bir beden, açık sarı saçlar, koyu renk gözler en belirgin özellikleridir. Özellikle Kova kadınlarının yüzleri çok güzeldir. Yaşlanıncaya kadar kendilerine bakarlar. Metabolizmaları çok hassastır. Doğal ilaçlar vücutlarında daha iyi etki yapar. Biyoenerji, Akapunktur, Homeopati, alternatif tedavi yöntemleri ilgilerini çeker. Bir anda başlayan hızla iyileşen, ve iz bırakmadan ilginç bir şekilde ortadan kaybolan hastalıklara yakalanabilirler. Bilekleri çok hassas olan Kovaların, spor yaparken dikkatli olmaları gerekiyor. Ayakkabı seçerken dikkatli davranmalılar. Aynı anda bir çok şeyi yapmaya kalkarlar. Spor yaşamlarında önemli bir yer tutmalı, yoksa elektriksel enerjilerini boşaltma konusunda zorluk çekebilirler."
		group = "Hava/erkeksi/sabit, Pozitif"
		manager_planet = "Satürn, Uranüs"
		color = "Koyu mavi, gece mavisi, yeşil"
		lucky_number = 4
		lucky_day = "Pazar"
		anti_horoscope = "Aslan"

	case "Balık":
		personal_characteristics = "Sabır, eli açıklık ve duyarlılık bu burçta doğan kişilerin en önemli nitelikleridir. Büyük bir inandırma yetenekleri vardır. Dürüst, vicdan sahibi, sadık ve uysaldırlar. Her çevreye kolayca uyabilirler. Genellikle hayal dünyasında yaşarlar. Yaşam görüşleri ciddi fakat gerçekçi değildirler. Balık burcu, diğer burçlar arasında dış etkenlerden en çok etkilenen kişilerdir. Düş dünyasında, öylesine mutludur ki, bazen onun aptal olduğunu bile düşünebilirsiniz. Ama, zannettiğinizden daha akıllı ve kurnazdır. Sinirli yapısını gizli bir sakinlikle örterken, herkesin seçtiği özel bir kişi olmanın hayallerini kurarlar. Yaşam onun için ürkütücü ve korkunçtur..Balık burcu insanını dışarıdan gözlemleyenler; duygusal yapısının bir aşk acısına dayanamayacağını zannederler ve kötü haberi hemen vermek istemezler. Aslında, kırılgan görüntüsünün altında güçlü bir kişilik yatar. Göz yaşlarını kısa sürede kurutur ve mendillerini bir kenara atarak, yeni bir yaşamın içinde kendini bulur. Onları teselli edecek o kadar kişi bulunur ki; yeni bir aşk acısı kapısını çalıncaya kadar, hayatını istediği gibi yaşar. Duygusal kırıklıklar pes etmelerine yetmez ve bir sürü yarım kalmış aşkıyla birlikte bir dolu anlatılacak anılarıyla yaşamını tamamlar."
		physical_features = "Balık etinde olup kol ve bacakları kısadır. Soluk tenli, açık kahverengi gözlü, uykulu gibi mahmur gözleri vardır. Kilo almaya meyillidirler. Yaşı ilerlediği zaman alın açılır ve şakaklardan beyazlamaya başlar. Her şeye çok çabuk ağladıkları için, gözleri hep nemlidir. Balıklar son derece hassas insanlardır. Kendi dünyalarında yarattıkları fanteziler ve gerçekler arasında sıkışır kalırlar. Zodyak’ın en melankolik burcudur. Kafalarına taktıkları bir problemleri varsa, çevrelerinden kopuk yaşarlar. Sürekli uykusuzluk çekerler. Sinirleri çok hassas oldukları içinde çok çabuk yıpranırlar. Alkol ve uyuşturucuya karşı dayanıksız bir yapıları olmasına rağmen, her fırsatta bu tür şeylere sığınabilirler. Yıllardır yapılan astrolojik araştırmalara göre; intihar vakaların çoğu balık insanlarında rastlanmıştır. Psikologlara en çok giden kişilerin de balık insanları olduğu saptanmıştır. Kış mevsimi balıklar için çok zor geçer. Soğuk algınlığına çabuk yakalanmaları, tüm mevsimi ellerinde mendil ile geçirmelerine neden olur. Ayakları çok hassas olduğu için, en küçük bir üşütmede karın ağrısı çekerler ve çok çabuk yatağa düşerler."
		group = "Su, dişi, negatif, değişken"
		manager_planet = "Jüpiter, Neptün"
		color = "Deniz Yeşili, mavi, yeşil"
		lucky_number = 3
		lucky_day = "Perşembe"
		anti_horoscope = "Başak"

	case "Koç":
		personal_characteristics = "Burçlar kuşağının ilk burcudur. Hareketli ve enerjik oluşları ile tanınırlar. Ben egoları çok fazla gelişmiştir. BEN, onların aynası olmuştur adeta. Bu burçta doğanlar çok pratiktirler. Olaylar karşısında coşkularını gizleyemezler. Yaşam yolunda canlılıklarını ve atılganlıklarını yitirmeden heyecanla ilerlerler. Merak ettikleri konularda olabildiğince yaratıcılardır. Amaçları doğrultusunda ilerlerken, kendilerini eylemleri ile kanıtlamak isterler. Eğer Koç'lar girişimde bulunacakları zaman izleyecekleri rotayı ayrıntıları ile planlarsa, enerjik yapılarının da yardımı ile daha da üretken olabilirler. Bencilliklerinden kaynaklanan sabırsızlıkları ve söz dinlemez yaratılışları yüzünden zaman zaman güç durumlara düştükleri de olur. Böyle anlarda başladıkları işlerini sonuçlandırmadan bırakırlar. Konuşmaları abartılıdır, bazen gerçekleri değiştirerek anlatırlar. Kavrama yetenekleri fazla olan Koç'lar yaşam sahnesinin başrolünde olmayı tercih ederler. Aşırı kıskanç ve bağımsızlıklarına düşkün olurlar."
		physical_features = "Fiziksel yapıları uzun yüzlü ve uzun boyludurlar. Kolları ve bacakları güçlüdür. Esmer ve kıvırcık saçları vardır. Bakışları delici ve keskindir. Sağlıklı bir ciltleri vardır.Spor yapmaya meraklı oldukları için de, bir çok spor salonlarında Koçlara rastlayabilirsiniz. Güzel dişleri olan Koçların; gözlerini dikerek gülümsemesi, çevrelerini etkisi altına alır."
		group = "Ateş/Erkeksi/Değişken"
		manager_planet = "Mars"
		color = "Ateş Kırmızısı, Nar Çiçeği"
		lucky_number = 9
		lucky_day = "Salı"
		anti_horoscope = "Terazi"

	case "Boğa":
		personal_characteristics = "Boğa'lar hedefleri doğrultusunda ilerlerken, tüm dikkatlerini toplayabilme yeteneklerinin yanında maddecilikleri ile tanınırlar. SAHİP olma onların yaşam gerçekleridir. Bireysel ilişkiler konusunda son derece güvenilir olan Boğa'lar, insanlara yardım etmekten hoşlanırlar. Yaşamları boyunca güven ararlar ve bu yüzden kendilerini riske atmazlar. Amaçladıkları işler konusunda gösterdikleri sabır, bazen diğer kişileri çatlatacak boyutlarda olabilir. Bütün bunlara rağmen çevresi tarafından aranılan, sıcakkanlı insanlardır. Sürekli somut konularla uğraşmayı severler, doğada bulunan olaylardan örneklemelerle işlerini başarılı bir şekilde yürütürler. Bu yüzden ruhsal olarak da doyumlu kişilerdir. Olayları organize ederlerken rahatlıklarını gözetirler ve kendilerinin zevklerine uygun olmasına önem verirler. Yaşamdaki isteklerini elde ettikleri zaman, hiçbir koşul onları başka yönlere çekemez. Parayı rahata ulaşmak için bir araç olarak görürler. Mal ve mülk edinme konusunda beceriklidirler. Finans ve yatırım konuları ilgilerini çeker."
		physical_features = "Kısa boylu, sağlam bir boyun yapısına sahiptirler. Kısa bir boyun ve mahmur bakışlıdırlar. Esmer tenli olup, koyu ve sıcak bakışlı kişilerdir. Saçlar gür ve kıvırcık olabilir.Özellikle boğa kadınları genç yaşlarda çok güzel olmalarına karşın, ileri yaşlarda aşırı nişastalı yiyecekler tükettikleri için, yüz hatları hantallaşır ve kalın bir bedene sahip olurlar."
		group = "Toprak/dişi/sabit/negatif"
		manager_planet = "Venüs"
		color = "Ateş pembe, açık mavi, eflatun"
		lucky_number = 6
		lucky_day = "Cuma"
		anti_horoscope = "Akrep"

	case "İkizler":
		personal_characteristics = "İkizler burcu insanları hızlı düşüncelerine uygun çabuk hareket ederler. Ayni anda birkaç işi birden yapabilirler. Onların adapte olamayacakları iş yoktur. Bu yüzden değişik karakterli olmaları ile tanınırlar. Bu yapılarını her zaman görebilmek mümkündür. Son derece neşeli ve mutlu oldukları bir anda, aniden mutsuz olabilirler. Çevreleri tarafından sürekli yanlış anlaşılabilirler. Herhangi bir konuda bilgileri az bile olsa, bunu çok iyi gizlemeyi başarırlar. Aksine; kulaktan dolma duydukları bilgileri öyle ustaca anlatırlar ki, dinleyenler onları o işin uzmanı sanırlar. Pratik zekalarıyla, çekici ve akıllıdırlar. Bu nedenle onları tanımlayan sözcük 'Düşünüyorum' dur. Fakat, bu düşünceleri hep yeni arayışlara doğru yönelmiştir. Bu yüzden uzun soluklu çalışmalar onları yorar. Kendilerini iyi eğitmiş ikizler hoş ve zariflikleri ile yaşamı zevkli kılarlarken, eğitimsiz olanlar da yaşamı o kadar çekilmez hale getirirler. Kendi paralarına karşı tutumlu davranmalarına karşın, başkalarının paralarını kolayca harcayabilirler."
		physical_features = "Güzel bir fiziğe sahiptirler. Uzun boy, biçimli bir beden, uzun kol ve bacaklar, koyu kumral saçlar, kahverengi gözlere sahiptirler. Gözleri keskindir ve uzağı çok iyi görürler. Gözlemcilikleri çok güçlüdür."
		group = "Hava, erkek, pozitif, değişken"
		manager_planet = "Merkür"
		color = "Sarı, gri, mavi, açık mavi, siyah"
		lucky_number = 5
		lucky_day = "Çarşamba"
		anti_horoscope = "Yay"

	case "Yengeç":
		personal_characteristics = "Yaşamlarındaki her konuda aşırı bir şekilde hassas, alıngan ve kuruntulu olan Yengeç'leri tanımlayan sözcük 'Hissederim' dir. Sorumluluk duyguları çok gelişmiştir. Her işte olağanüstü olan ayrıntıcılıkları, işlerinde mükemmeliyetçiliği getirir. Ayni sorumlulukları karşılarındakilerden de beklerler. Yengeç'ler duygusallıkları ve duyarlılıkları ile tanınırlar. Çevresindeki her insandan da ayni hassasiyeti bekledikleri için, kolay geçinilir tipler değildir. İyi günlerinde neşeli, iyi kalpli, yardımsever, düşünceli ve anlayışlıdırlar. Fakat herhangi belirgin bir neden olmadan somurtkan ve alıngan olabilirler. Yakınlarını ve arkadaşlarını çok sevmelerine karşın, bunu pek belli etmezler. Kendilerini herhangi bir şekilde inciten kişileri zor bağışlarlar ve yapılan hareketi asla unutmazlar. Yengeç'ler müziğe ve dinsel konulara karşı ilgilidirler. Sabırlı olan Yengeç'ler tartışmalardan kesinlikle hoşlanmazlar. Duygularını sessiz bir şekilde saklarlar."
		physical_features = "Gerçek Yengeç Burcu insanı fiziksel olarak; soluk, beyaz tenli, orta ve kısa boylu, yuvarlak yüzlüdür. Gözleri genelde gri veya mavi gözlüdür. Saçları mat ve kahverengi olur.Yengeç kadınları tartışmasız çok güzeldirler. Ay gezegenini simgeleyen güzel yüzleri vardır. Vücut yapıları gençliklerinde güzeldirler. Yaşları ilerledikçe, dikkat etmezlerse kilolu tombul bir vücutları olur."
		group = "Su, negatif,dişi"
		manager_planet = "Ay"
		color = "Gümüş, gri, eflatun"
		lucky_number = 2
		lucky_day = "Pazartesi"
		anti_horoscope = "Oğlak"

	case "Aslan":
		personal_characteristics = "Aslan kraldır, önderdir. Başkalarının yaşantılarını da onlar adına düzenlemek isterler. Her şeye karışırlar, kibirlidirler. Bu nedenle onları tanımlayan sözcük 'Yönetirim' dir. Yaşam sahnesinde her zaman parlayarak, odak noktası olmak isterler. Organizasyon güçleri çok fazladır. İsteklerini başkalarına kabul ettirmek, onlar için yaşamlarının 'olmazsa olmaz' şartıdır. İyi zamanlarında etkileyici, güler yüzlü, başkalarına yardım etmeyi seven ve bunu kendine görev sayan Aslan'lar sevimli ve iyimser kişilerdir. Ona karşı hatalı davransanız bile, size olgun bir şekilde tepki verir. Fakat; Sabrı taştıktan sonra, dürüst ve mert, gerektiğinde sert bir şekilde tavır gösterir. Zor günlerinde şansları onlara her zaman yardım eder. Yönetici gezegenleri Güneş onları en karanlık günlerinde aydınlığa çıkarır. Eğitimsiz ve gelişmemiş Aslan tipleri çekilmez olurlar. Her konuda sahip olduklarından daha fazlası varmış gibi davranırlar."
		physical_features = "Boyları genelde uzun ve kemikleri kalındır. Omuzlar geniş, kaslar gelişmiştir. Açık renk saçlı, pembe beyaz tenli, büyük ve yuvarlak başlıdır. Gözleri iyi görür. Kadınları saçlarını çok severler. Dış görünüşlerine çok önem vermeleri yüzünden, onları her an bakımlı görebilirsiniz. Yürüyüşleri farklıdır ve girdikleri yerlerde gözle hemen üstlerinde toplanır. Fiziksel çekiciliğinin farkında oluşu, onu daha gururlu ve ben merkezci yapar."
		group = "Ateş, erkek, pozitif, sabit"
		manager_planet = "Güneş"
		color = "Koyu sarı, altın rengi, turuncu, kehribar"
		lucky_number = 1
		lucky_day = "Pazar"
		anti_horoscope = "Kova"

	case "Başak":
		personal_characteristics = "Yönetici gezegeninizden dolayı hep bilgiyi ararlar. Zekalarını kendilerine yardımcı olan bir hizmetçi gibi görürler. Bu nedenle Başak burcunu tanımlayan sözcük 'İncelerim “ dir. Başak'lar çalışkan ve pratik insanlar olup, yaşamlarındaki en önemli konu İş' tir. Güvendikleri kişilere yardım etmeyi sevmelerine rağmen, inanmadıkları ve tembel olduklarını bildikleri kişilere karşı soğuk davranırlar. Yaşamları boyunca dinlenmeden çalışırlar. Onların dinlenme biçimi bile başkalarına yorucu gelebilir. Başak'ların yaşamda ayrıntılar arasında boğulma riskleri hep vardır. Böyle anlarda bile, kendi yöntemleri ile yaptıkları işlerde gelişigüzel şeyler bulunabileceğini kabul etmezler. Başaklar, genellikle kendilerini hiç kimseye kullandırtmazlar, sınırlarını belirleyerek 'hayır' demesini bilirler. Tutumlulukları bazen pintilik derecesindedir. İçli-dışlı olmayı sevmedikleri için, soğuk ve mesafeli bir görünüşleri vardır."
		physical_features = "Başak burcunun tipik fiziksel özellikleri: Uzun bir boy, ince bir bel yüzü pembe beyaz veya esmerdir. Saçları koyu renk ve genelde uzun olur. Sağlıklı saç tellerine sahiptirler. Gözleri insanın içini okur gibi baktığı için de hemen onun etkisinde kalırsınız Farklı bir karizması vardır. Sistemli ve kuralcı yapıları dış görünüşlerine de yansır. Onları asla dağınık ve gelişigüzel göremezsiniz. Kıyafetleri klasik fakat etkileyicidir. Daima kaliteli yaşamaktan hoşlandıkları için, bütçelerini yansıtan hoş kıyafetleri kullanırlar, fakat abartıya kaçmamak şartıyla."
		group = "Toprak/dişi/değişken"
		manager_planet = "Merkür"
		color = "Açık sarı, açık mavi, turkuvaz"
		lucky_number = 5
		lucky_day = "Çarşamba"
		anti_horoscope = "Balık"

	case "Terazi":
		personal_characteristics = "Kararsızlıkları ile ünlü cazibe sembolü karşınızda duruyor. Nedenini bilmediğiniz bir çekim gücü etkisi altına girdiğinizi fark ettiğinizde iş işten çoktan geçmiş olacaktır. Onların büyülü bir havası vardır. Ritmik hareketleri ile dans eder gibi bir yaşam sürer. Onun yanında kızgınlıklarınız yok olur. En sinirli anınızda bile sizi regüle etme yeteneğine sahiptir. Üstelik, tartışmayı sevmesine rağmen; bu kadar ustalıkla konulara hakim olup, sonrada sesini hiç yükseltmeden zaferini ilen eden çok az kişi vardır. Bir Terazi bunu başarır. Sizin kılıç salladığınız ortamlarda, o; gülücüklerini, zarafetini ve öpücüklerini gönderecektir. Kafasında müthiş bir denge vardır. Onun ne düşündüğünü bilemezsiniz. Uyumlu ve sevecen yapısının altında gizli bir güç vardır. Bulunduğu ortamları çok çabuk inceler ve gizli detayların nerede saklı olduğunu bulmaya çalışır. Konuya hakim olmak için deliler gibi çalışmak ve kendini harap etmek yerine, keskin görüşlerini ve doğru mantığını kullanacağı ortamları seçer. Başarılı olmak için imkansızlığı değil, kendini yüceltecek elle tutulur konular üzerinde uyumlu bir sessizlikle çalışır. Yaptığı işlerin doğruluğuna inandığı halde, onay beklemekten de kendini alamaz. Onu tenkit etme gibi bir şansınız asla olamaz. Böyle bir şeye kalkıştığınız taktirde; size tüm zarafetiyle gülümseyecek, önerilerinizi can kulağıyla dinleyecek ve sonunda kibarca teşekkür ederek sizden uzaklaşacaktır. Fakat bir daha asla sizin fikrinizi sormayacaktır. Gülümseyen çehresinin altında müthiş bir ego gizlidir, aşk ve sanat onun yaşam öğeleridir. Dünyada her şeyden vazgeçebilir, fakat kişiliğini besleyen ve duygusallığını yansıtan soyut kavramlarından vazgeçemez. Akıl, mantık, aşk işte karşınızda her şeyiyle özel bir kişi duruyor. Onun terazilerinin kefesinde yer almak istiyorsanız, belli standart ölçülerin çok üstünde olmanız gerekir. Kolay bir insan olmadığını şimdiden hatırlatırız."
		physical_features = "Terazi burçlarının fiziksel yapıları çok güçlüdür. Zodyak’ın en güzel kişileri diyebiliriz. Uzun boylu, ince vücutlu, düz kahverengi saçlı, yuvarlak yüzlü, beyaz tenli ve koyu renk gözlüdürler. Terazi bir denge burcudur. Önemsiz bir dengesizlik hassas vücutlarını alt üst etmeye yeterlidir. Hastalıkları sırasında yoğun ilgi beklerler. Aradıkları ilgiyi bulamadıkları zaman, huysuzlaşarak daha çok hastalanırlar. İç dünyasında dalgalanmalar yüzünden, baş ağrısı ve aşırı yemek eğilimi gösterirler. Sinir sistemleri midelerini etkilediği için midedeki asit fazlalığını fosfor içerikli yiyeceklerle dengelemelidirler. Terazilerin metabolizması çok yavaş çalışır. Baharatlı, soslu, şekerli gıdalara düşkün olması nedeniyle, çok fazla kilo almaya müsaittir. Yaşam boyunca sürekli egzersiz yapmak zorundadır, aksi taktirde ileri yaşlarda şişmanlık kaçınılmaz olur. Terazi insanları; yapılı, fakat çok çabuk hastalanan kişilerdir. Sağlıklı görünüşlerine karşın, soğuk algınlıklarına çok çabuk yakalanırlar. Terazileri; böbrekler, bel bölgesi, sırt ve omurgaları simgeler. Ayrıca, Venüs etkisinde oldukları için troidler düzensiz çalışabilir. Kilo problemi olan Terazilerin, ara sıra hormon testleri yaptırmaları gerekebilir. İleri yaşlarda romatizmal rahatsızlıklara aday kişilerdir. Terazilerin, herkesten daha çok kendilerine özen göstermeleri gerekiyor ki, zaten çoğu bu konuda duyarlı olurlar. Kendilerine özen gösteren insanlar en çok bu burçtan çıkar."
		group = "Hava/erkeksi/öncü, Pozitif"
		manager_planet = "Hava/erkeksi/öncü, Pozitif"
		color = "Pastel renkler"
		lucky_number = 6
		lucky_day = "Cuma"
		anti_horoscope = "Koç"

	case "Akrep":
		personal_characteristics = "Akrepler kadar yaşamda tutkuyla yaşayan az insan vardır. Bu nedenle onları simgeleyen sözcük 'Arzuluyorum' dur. Hiçbir şeyi yarım bırakmazlar. Akreplerin güçleri gözlerinden okunur. Mimiklerini kontrol altında tutsalar bile, bakışları ile sevgilerini ya da nefretlerini aktarabilirler. Duygularına kapılırlarsa, tehlikeli olabilirler. Akrep'ler ukala ve kendini beğenmiş insanları sevmezler, onları aşağılayarak hadlerini bildirirler. Kendi bildiklerini okuyarak, uzlaşmaya yanaşmazlar. Aşırı bir şekilde kuşkuculardır, kolay inanmazlar ama inandıkları bir konuyu da sonuna kadar inatla savunurlar. Çalışmalarında sabır ve özenle çalışırlarken, gösterişten uzak bir şekilde işlerini yaparlar. Kendilerini yetiştirmemiş Akrep'ler, yaşamın basitliklerine yatkınlıkları ile kendi kendilerinin yok oluşlarına neden olurlar. Cinsellik yaşamlarında önemli bir yer tutar. Ölümü yeni bir başlangıç olarak kabul ettiklerinden, ölümden korkmazlar."
		physical_features = "Akreplerin fiziksel yapıları oldukça ilginçtir; orta yada daha uzun boylu olabilirler. Kalın fakat düzgün bir bedenleri vardır. Yuvarlak yüzleri olur. Saçları koyu renk ve dalgalı, hatta kıvırcıktır. Gözleri inanılma çekicidir. Baktıkları her şeyi adeta eritirler. Bir Akrebi bakışlarından kolayca tanıyabilirsiniz. Zodyak’ın en anlaşılmaz burcudurlar. Çoğu kez eleştiriye maruz kaldıkları için, kendilerini savunmaktan vazgeçmişlerdir. Duyguları çok derindir. Gizemli yapıları onları bulundukları ortamda farklı yerlere getirirler. Duyumsamaları yakıcıdır. Yoğun duyguları, kıskançlıklarını sürekli körükler. Hırslı yapıları yüzünden en zor hastalıklar gelip onları bulur. Bir Akrep hastalandığı zaman, teşhis edilmesi zor tahlillere gerek duyulabilir. Uzun araştırmalar sonucu hastalıkların nedeni genelde psikosomatik kökenlidir. Açığa çıkaramadıkları duyguları yüzünden bedenlerini harap ederler. Akreplerin cinsiyet hormonları herkesten çok daha fazla çalışır. Vücutlarındaki elektriksel potansiyel nedeniyle, bulundukları ortamda insanları kendilerine mıknatıs gibi çekerler.. Hastalık konusunda sabırsızdır. Acıya onun kadar dayanıklı çok az kişi görülür. Alınabilecek en etkili ilaçları hiç çekinmeden alırlar. İradelerine çok güvenir ve beyin gücü ile kendilerini iyileştirirler. Akrepler mide ve üreme organlarını temsil ettiği için, bu organları hassastır fakat her türlü rahatsızlığın üstesinden kolay gelirler."
		group = "Su, negatif, dişi, sabit"
		manager_planet = "Mars, Plüton"
		color = "Canlı kırmızı, siyah"
		lucky_number = 9
		lucky_day = "Salı"
		anti_horoscope = "Boğa"

	case "Yay":
		personal_characteristics = "Kavrama yetenekleri gelişmiş olduğundan becerikliliklerinin de katkısı ile ele aldıkları her işte, özellikle ciddi işlerde ve felsefe konularında başarılı olurlar. Bu nedenle Yay Burcunu tanımlayan sözcük 'Görüyorum' dur. Yay burcu insanları içtenlikleri ve iyimser yaşam görüşleri ile tanınırlar. Gençliklerinde dikkatsiz, heyecanlı ve geleneklere aykırı davranışlar içinde olsalar da, geçmiş yanılgılarından en çok ders alan kişiler bu burçtan çıkar. Özgürlüklerine aşırı düşkün, patavatsızlık derecesinde pratik insanlardır. Yay'ların yaşam çerçeveleri herhangi bir şekilde kısıtlandığı zaman, içsel bir biçimde alt üst olurlar. Akılsızca risklere atılırlar. Yayların hayatı yeniliklere olan merakları ile karakterize edilebilir. Bilmedikleri şeyleri araştırıp, keşfetmeyi severler. Yaylar çok yönlü ve ayni anda ilgilenebilecekleri birkaç konu olduğunda mutlu olan kişilerdir. Kendilerini yorgun hissediyorlarsa, o konudan sıkılmış demektir."
		physical_features = "Uzun boy, ince ve yağsız beden, çekici bir yüz, kahverengi saçlar, pembe-beyaz tenleri vardır. Kendilerine daima güvenen ve iyimser yüzleri bakışlarına yansımıştır. Parlak ve çekici bir gülümseyiş şekilleri vardır. Sıra dışı ve farklı giyinmekten hoşlandıkları için, onları fark etmemek mümkün değildir. Çocuksu yapıları ile asla büyümezler."
		group = "Ateş/erkeksi/değişken"
		manager_planet = "Jupiter, Neptün"
		color = "Turkuaz, Mor"
		lucky_number = 3
		lucky_day = "Perşembe"
		anti_horoscope = "İkizler"

	}

	fmt.Printf("\n\n Grup : %s", group)
	fmt.Printf("\n\n Yönetici Gezegen : %s ", manager_planet)
	fmt.Printf("\n\n Renk : %s", color)
	fmt.Printf("\n\n Şanslı Sayı : %d", lucky_number)
	fmt.Printf("\n\n Şanslı Gün : %s", lucky_day)
	fmt.Printf("\n\n Karşıt Burç : %s", anti_horoscope)
	fmt.Printf("\n\n Kişisel Özellikler : %s", personal_characteristics)
	fmt.Printf("\n\n Fiziksel Özellikler : %s", physical_features)
	fmt.Println()
}
