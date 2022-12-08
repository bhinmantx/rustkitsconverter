// idToItem
package main

import (
	"fmt"
)

func idToItem(item_id int) (short_name string, err error) {

	//check for invalid names
	idsShortNames := map[int]string{
		550753330:   "ammo.snowballgun",
		-1685290200: "ammo.shotgun",
		-1036635990: "ammo.shotgun.fire",
		-727717969:  "ammo.shotgun.slug",
		174866732:   "weapon.mod.8x.scope",
		349762871:   "ammo.grenadelauncher.he",
		1055319033:  "ammo.grenadelauncher.buckshot",
		915408809:   "ammo.grenadelauncher.smoke",
		-1211166256: "ammo.rifle",
		567235583:   "weapon.mod.small.scope",
		-1215166612: "barrelcostume",
		1840570710:  "abovegroundpool",
		-2124352573: "fun.guitar",
		-1729415579: "radiationresisttea.advanced",
		-2123125470: "healingtea.advanced",
		603811464:   "maxhealthtea.advanced",
		2063916636:  "oretea.advanced",
		2021351233:  "radiationremovetea.advanced",
		524678627:   "scraptea.advanced",
		-541206665:  "woodtea.advanced",
		-2027793839: "xmas.advent",
		342438846:   "fish.anchovy",
		1171735914:  "electric.andswitch",
		-1018587433: "fat.animal",
		-487356515:  "radiationresisttea",
		-1432674913: "antiradpills",
		1548091822:  "apple",
		1107575710:  "hazmatsuit_scientist_arctic",
		-470439097:  "hazmatsuit.arcticsuit",
		1874610722:  "vehicle.1mod.cockpit.armored",
		1353298668:  "door.hinged.toptier",
		1221063409:  "door.double.hinged.toptier",
		-1615281216: "vehicle.1mod.passengers.armored",
		1545779598:  "rifle.ak",
		-1335497659: "rifle.ak.ice",
		2100007442:  "electric.audioalarm",
		-2139580305: "autoturret",
		-2072273936: "bandage",
		-702051347:  "mask.bandana",
		-1622110948: "attire.banditguard",
		1382263453:  "barricade.woodwire",
		1099314009:  "bbq",
		-1022661119: "hat.cap",
		-929092070:  "healingtea",
		-1211268013: "horse.shoes.basic",
		-1184406448: "maxhealthtea",
		1480022580:  "oretea",
		263834859:   "scraptea",
		-649128577:  "woodtea",
		609049394:   "battery.small",
		-321431890:  "beachchair",
		-1621539785: "beachparasol",
		657352755:   "beachtable",
		-8312704:    "beachtowel",
		1840822026:  "grenade.beancan",
		-1273339005: "bed",
		1675639563:  "hat.beenie",
		-1262185308: "tool.binoculars",
		1973165031:  "cakefiveyear",
		1771755747:  "black.berry",
		122783240:   "clone.black.berry",
		1911552868:  "seed.black.berry",
		1931713481:  "black.raspberries",
		1553078977:  "bleach",
		-690968985:  "electric.blocker",
		1776460938:  "blood",
		1112162468:  "blue.berry",
		838831151:   "clone.blue.berry",
		803954639:   "seed.blue.berry",
		1744298439:  "firework.boomer.blue",
		1036321299:  "bluedogtags",
		1601468620:  "jumpsuit.suit.blue",
		-484206264:  "keycard_blue",
		-515830359:  "firework.romancandle.blue",
		-586342290:  "blueberries",
		-996920608:  "blueprintbase",
		1588298435:  "rifle.bolt",
		1746956556:  "bone.armor.suit",
		215754713:   "arrow.bone",
		1711033574:  "bone.club",
		1719978075:  "bone.fragments",
		-1903165497: "deer.skull.mask",
		1814288539:  "knife.bone",
		-1478094705: "boogieboard",
		-1113501606: "boombox",
		-23994173:   "hat.boonie",
		-1549739227: "shoes.boots",
		613961768:   "botabag",
		844440409:   "easter.bronzeegg",
		850280505:   "bucket.helmet",
		1525520776:  "building.planner",
		-1004426654: "attire.bunnyears",
		23391694:    "hat.bunnyhat",
		-1266045928: "attire.bunny.onesie",
		21402876:    "burlap.gloves.new",
		1877339384:  "burlap.headwrap",
		602741290:   "burlap.shirt",
		-761829530:  "burlap.shoes",
		1992974553:  "burlap.trousers",
		-989755543:  "bearmeat.burned",
		1973684065:  "chicken.burned",
		-78533081:   "deermeat.burned",
		1917703890:  "horsemeat.burned",
		-682687162:  "humanmeat.burned",
		1391703481:  "meat.pork.burned",
		1827479659:  "wolfmeat.burned",
		-194509282:  "knife.butcher",
		-1778897469: "electric.button",
		1835946060:  "electric.cabletunnel",
		1783512007:  "cactusflesh",
		-1316706473: "tool.camera",
		1946219319:  "campfire",
		-1040518150: "vehicle.2mod.camper",
		-700591459:  "can.beans",
		-1941646328: "can.tuna",
		-1379036069: "fun.tambourine",
		1714496074:  "hat.candle",
		1121925526:  "candycane",
		1789825282:  "candycaneclub",
		1230691307:  "captainslog",
		946662961:   "car.key",
		3380160:     "movembermoustachecard",
		1081921512:  "cardtable",
		1524980732:  "carvable.pumpkin",
		476066818:   "cassette",
		-912398867:  "cassette.medium",
		1523403414:  "cassette.short",
		-1530414568: "fun.casetterecorder",
		-587989372:  "fish.catfish",
		634478325:   "cctv.camera",
		1142993169:  "ceilinglight",
		-1117626326: "wall.frame.fence",
		1451568081:  "wall.frame.fence.gate",
		1104520648:  "chainsaw",
		1534542921:  "chair",
		1324203999:  "firework.boomer.champagne",
		-1938052175: "charcoal",
		-1916473915: "chineselantern",
		359723196:   "arcade.machine.chippy",
		363467698:   "chocholate",
		2009734114:  "xmasdoorwreath",
		1058261682:  "xmas.lightstring",
		794443127:   "xmas.tree",
		968019378:   "clatter.helmet",
		-858312878:  "cloth",
		204391461:   "coal",
		-1501451746: "vehicle.1mod.cockpit",
		170758448:   "vehicle.1mod.cockpit.with.engine",
		1159991980:  "lock.code",
		-803263829:  "coffeecan.helmet",
		573676040:   "coffin.storage",
		2040726127:  "knife.combat",
		-1488398114: "composter",
		884424049:   "bow.compound",
		-1588628467: "computerstation",
		-1950721390: "barricade.concrete",
		968421290:   "connected.speaker",
		1873897110:  "bearmeat.cooked",
		-1848736516: "chicken.cooked",
		-1509851560: "deermeat.cooked",
		1668129151:  "fish.cooked",
		-1162759543: "horsemeat.cooked",
		1536610005:  "humanmeat.cooked",
		-242084766:  "meat.pork.cooked",
		813023040:   "wolfmeat.cooked",
		1367190888:  "corn",
		-778875547:  "clone.corn",
		998894949:   "seed.corn",
		-216999575:  "electric.counter",
		-1049881973: "fun.cowbell",
		1189981699:  "cratecostume",
		1965232394:  "crossbow",
		-321733511:  "crude.oil",
		1242522330:  "cursedcauldron",
		1796682209:  "smg.2",
		-1667224349: "xmas.decoration.baubels",
		1686524871:  "xmas.decoration.gingerbreadmen",
		-129230242:  "xmas.decoration.pinecone",
		-209869746:  "xmas.decoration.candycanes",
		2106561762:  "xmas.decoration.tinsel",
		-151387974:  "xmas.lightstring.advanced",
		1568388703:  "diesel_barrel",
		1895235349:  "discoball",
		1735402444:  "discofloor.largetiles",
		286648290:   "discofloor",
		296519935:   "diving.fins",
		-113413047:  "diving.mask",
		-2022172587: "diving.tank",
		1223900335:  "dogtagneutral",
		1409529282:  "door.closer",
		-502177121:  "electric.doorcontroller",
		-1112793865: "door.key",
		-765183617:  "shotgun.double",
		1521286012:  "sign.post.double",
		-854270928:  "dragondoorknocker",
		-22883916:   "hat.dragonmask",
		1588492232:  "drone",
		-1519126340: "dropbox",
		1401987718:  "ducttape",
		1015352446:  "submarineduo",
		-979302481:  "easterdoorwreath",
		1856217390:  "easterbasket",
		-747743875:  "attire.egg.suit",
		-394470247:  "sign.egg.suit",
		-629028935:  "fuse",
		-784870360:  "electric.heater",
		-1448252298: "electrical.branch",
		1177596584:  "elevator",
		1655979682:  "can.beans.empty",
		-1673693549: "propanetank",
		-1557377697: "can.tuna.empty",
		1559779253:  "vehicle.1mod.engine",
		-75944661:   "pistol.eoka",
		-1321651331: "ammo.rifle.explosive",
		-592016202:  "explosives",
		2005491391:  "weapon.mod.extendedmags",
		143803535:   "grenade.f1",
		2054391128:  "factorydoor",
		-930193596:  "fertilizer",
		674734128:   "xmas.door.garland",
		-1230433643: "xmas.double.door.garland",
		-1379835144: "xmas.window.garland",
		14241751:    "arrow.fire",
		-1961560162: "lunar.firecrackers",
		-1707425764: "fishing.tackle",
		-1215753368: "flamethrower",
		528668503:   "flameturret",
		304481038:   "flare",
		-936921910:  "grenade.flashbang",
		-939424778:  "electric.flasherlight",
		-196667575:  "flashlight.held",
		-1880231361: "vehicle.1mod.flatbed",
		936496778:   "floor.grill",
		1983621560:  "floor.triangle.grill",
		-265292885:  "fluid.combiner",
		-1166712463: "fluid.splitter",
		443432036:   "fluid.switch",
		-1973785141: "fogmachine",
		1575635062:  "frankensteintable",
		1413014235:  "fridge",
		-1000573653: "boots.frog",
		1186655046:  "vehicle.2mod.fuel.tank",
		-1999722522: "furnace",
		-148794216:  "wall.frame.garagedoor",
		1803831286:  "toolgun",
		1659114910:  "hat.gas.mask",
		479143914:   "gears",
		999690781:   "geiger.counter",
		1770744540:  "vehicle.chassis",
		878301596:   "vehicle.module",
		-1043618880: "ghostsheet",
		-695124222:  "giantcandycanedecor",
		282103175:   "giantlollipops",
		-690276911:  "gloweyes",
		-1899491405: "glue",
		-1002156085: "easter.goldegg",
		-746030907:  "granolabar",
		809199956:   "gravestone",
		-1679267738: "wall.graveyard.fence",
		858486327:   "green.berry",
		-1305326964: "clone.green.berry",
		-1776128552: "seed.green.berry",
		-656349006:  "firework.boomer.green",
		1268178466:  "industrial.wall.light.green",
		37122747:    "keycard_green",
		-1306288356: "firework.romancandle.green",
		-568419968:  "grub",
		-265876753:  "gunpowder",
		-1759188988: "habrepair",
		-888153050:  "halloween.candy",
		200773292:   "hammer",
		1569882109:  "fishingrod.handmade",
		588596902:   "ammo.handmade.shell",
		-1252059217: "hatchet",
		1266491000:  "hazmatsuit",
		-1507239837: "electric.hbhfsensor",
		-1569700847: "twitch.headset",
		-297099594:  "frankensteins.monster.03.head",
		-2024549027: "frankensteins.monster.03.legs",
		1614528785:  "frankensteins.monster.03.torso",
		1181207482:  "heavy.plate.helmet",
		-1102429027: "heavy.plate.jacket",
		-1778159885: "heavy.plate.pants",
		-1772746857: "scientistsuit_heavy",
		-886280491:  "clone.hemp",
		-237809779:  "seed.hemp",
		-1698937385: "fish.herring",
		794356786:   "attire.hide.boots",
		3222790:     "attire.hide.helterneck",
		1722154847:  "attire.hide.pants",
		980333378:   "attire.hide.poncho",
		-1773144852: "attire.hide.skirt",
		196700171:   "attire.hide.vest",
		-691113464:  "gates.external.high.stone",
		-967648160:  "wall.external.high.stone",
		-335089230:  "gates.external.high.wood",
		99588025:    "wall.external.high",
		-985781766:  "wall.external.high.ice",
		656371026:   "carburetor3",
		1158340332:  "crankshaft3",
		1989785143:  "horse.shoes.advanced",
		317398316:   "metal.refined",
		-1982036270: "hq.metal.ore",
		1883981800:  "piston3",
		1072924620:  "sparkplug3",
		-1802083073: "valve3",
		-1023065463: "arrow.hv",
		-1841918730: "ammo.rocket.hv",
		1160881421:  "hitchtroughcombo",
		-1214542497: "hmlmg",
		-1442559428: "hobobarrel",
		442289265:   "weapon.mod.holosight",
		-1663759755: "trap.landmine",
		1751045826:  "hoodie",
		-1579932985: "horsedung",
		-1997543660: "horse.saddle",
		363163265:   "hosetool",
		-143132326:  "sign.wooden.huge",
		996293980:   "skull.human",
		1443579727:  "bow.hunting",
		1712070256:  "ammo.rifle.hv",
		-1691396643: "ammo.pistol.hv",
		-44876289:   "electric.igniter",
		-2012470695: "mask.balaclava",
		605467368:   "ammo.rifle.incendiary",
		51984655:    "ammo.pistol.fire",
		1638322904:  "ammo.rocket.fire",
		818733919:   "door.hinged.industrial.a",
		1623701499:  "industrial.wall.light",
		-697981032:  "innertube",
		185586769:   "innertube.horse",
		2052270186:  "innertube.unicorn",
		-2001260025: "tool.instant_camera",
		1242482355:  "jackolantern.angry",
		-1824943010: "jackolantern.happy",
		-1163532624: "jacket",
		1488979457:  "jackhammer",
		-979951147:  "fun.jerrycanguitar",
		-97459906:   "jumpsuit.suit",
		-1330640246: "drumkit",
		190184021:   "kayak",
		-850982208:  "lock.key",
		-778367295:  "rifle.l96",
		1948067030:  "floor.ladder.hatch",
		1697996440:  "photoframe.landscape",
		-845557339:  "sign.pictureframe.landscape",
		1658229558:  "lantern",
		1643667218:  "sign.neon.xl.animated",
		23352662:    "sign.hanging.banner.large",
		2070189026:  "sign.pole.banner.large",
		-489848205:  "largecandles",
		-44066790:   "vehicle.chassis.4mod",
		-1693832478: "vehicle.2mod.flatbed",
		-1992717673: "furnace.large",
		479292118:   "halloween.lootbag.large",
		254522515:   "largemedkit",
		866332017:   "sign.neon.xl",
		1205084994:  "photoframe.large",
		1581210395:  "planter.large",
		-1622660759: "xmas.present.large",
		553270375:   "electric.battery.rechargable.large",
		2090395347:  "electric.solarpanel.large",
		-1100168350: "water.catcher.large",
		833533164:   "box.wooden.large",
		1153652756:  "sign.wooden.large",
		-798293154:  "electric.laserdetector",
		853471967:   "laserlight",
		1381010055:  "leather",
		1366282552:  "burlap.gloves",
		-134959124:  "frankensteins.monster.01.head",
		106959911:   "frankensteins.monster.01.legs",
		-1624770297: "frankensteins.monster.01.torso",
		-110921842:  "locker",
		-2027988285: "locomotive",
		935692442:   "tshirt.long",
		-1469578201: "longsword",
		-946369541:  "lowgradefuel",
		656371028:   "carburetor1",
		1158340334:  "crankshaft1",
		1883981798:  "piston1",
		-89874794:   "sparkplug1",
		1330084809:  "valve1",
		-1812555177: "rifle.lr300",
		-763071910:  "lumberjack hoodie",
		-2069578888: "lmg.m249",
		28201841:    "rifle.m39",
		-852563019:  "pistol.m92",
		-1966748496: "mace",
		-1137865085: "machete",
		-586784898:  "mailbox",
		1426574435:  "minihelicopter.repair",
		1079279582:  "syringe.medical",
		42535890:    "sign.neon.125x215.animated",
		-44066823:   "vehicle.chassis.3mod",
		-1732475823: "frankensteins.monster.02.head",
		835042040:   "frankensteins.monster.02.legs",
		1491753484:  "frankensteins.monster.02.torso",
		1899610628:  "halloween.lootbag.medium",
		-1423304443: "sign.neon.125x215",
		756517185:   "xmas.present.medium",
		656371027:   "carburetor2",
		1158340331:  "crankshaft2",
		1883981801:  "piston2",
		-493159321:  "sparkplug2",
		926800282:   "valve2",
		2023888403:  "electric.battery.rechargable.medium",
		-1819233322: "sign.wooden.medium",
		-583379016:  "megaphone",
		-746647361:  "electrical.memorycell",
		1655650836:  "barricade.metal",
		1882709339:  "metalblade",
		1110385766:  "metal.plate.torso",
		-194953424:  "metal.facemask",
		69511070:    "metal.fragments",
		-1199897169: "shutter.metal.embrasure.a",
		-4031221:    "metal.ore",
		95950017:    "metalpipe",
		-148229307:  "wall.frame.shopfront.metal",
		-1021495308: "metalspring",
		-1199897172: "shutter.metal.embrasure.b",
		-819720157:  "wall.window.bars.metal",
		39600618:    "microphonestand",
		-1539025626: "hat.miner",
		1052926200:  "mining.quarry",
		-542577259:  "fish.minnows",
		1259919256:  "mixingtable",
		-1449152644: "mlrs",
		343045591:   "aiming.module.mlrs",
		-1843426638: "ammo.rocket.mlrs",
		-20045316:   "mobilephone",
		1696050067:  "modularcarlift",
		1556365900:  "grenade.molotov",
		-2047081330: "movembermoustache",
		1318558775:  "smg.mp5",
		-1123473824: "multiplegrenadelauncher",
		277730763:   "halloween.mummysuit",
		-1962971928: "mushroom",
		-1405508498: "weapon.mod.muzzleboost",
		1478091698:  "weapon.mod.muzzlebrake",
		1953903201:  "pistol.nailgun",
		-2097376851: "ammo.nailgun.nails",
		1081315464:  "attire.nesthat",
		1516985844:  "wall.frame.netting",
		-961457160:  "newyeargong",
		-1518883088: "nightvisiongoggles",
		-1506417026: "attire.ninja.suit",
		491263800:   "hazmatsuit.nomadsuit",
		1414245162:  "note",
		-1832422579: "sign.post.town",
		-1286302544: "electric.orswitch",
		-7270019:    "firework.boomer.orange",
		-1904821376: "fish.orangeroughy",
		1315082560:  "hat.oxmask",
		1491189398:  "paddle",
		-733625651:  "paddlingpool",
		-126305173:  "easter.paintedeggs",
		-2040817543: "fun.flute",
		237239288:   "pants",
		-1779183908: "paper",
		696029452:   "map",
		-575744869:  "partyhat",
		895374329:   "vehicle.2mod.passengers",
		-379734527:  "firework.boomer.pattern",
		62577426:    "photo",
		-1302129395: "pickaxe",
		286193827:   "jar.pickle",
		785728077:   "ammo.pistol",
		1090916276:  "pitchfork",
		-804769727:  "plantfiber",
		273172220:   "fun.trumpet",
		-1651220691: "pookie.bear",
		576509618:   "fun.boomboxportable",
		1729712564:  "photoframe.portrait",
		-1370759135: "sign.pictureframe.portrait",
		-2086926071: "potato",
		1512054436:  "clone.potato",
		-2084071424: "seed.potato",
		-365097295:  "powered.water.purifier",
		-2049214035: "electric.pressurepad",
		-956706906:  "wall.frame.cell.gate",
		-1429456799: "wall.frame.cell",
		-1130709577: "mining.pumpjack",
		795371088:   "shotgun.pump",
		-567909622:  "pumpkin",
		1346158228:  "pumpkinbasket",
		1898094925:  "clone.pumpkin",
		-1511285251: "seed.pumpkin",
		-33009419:   "radiationresisttea.pure",
		-1677315902: "healingtea.pure",
		1712261904:  "maxhealthtea.pure",
		1729374708:  "oretea.pure",
		1905387657:  "radiationremovetea.pure",
		2024467711:  "scraptea.pure",
		-557539629:  "woodtea.pure",
		20489901:    "twitchsunglasses",
		1373971859:  "pistol.python",
		-496584751:  "radiationremovetea",
		492357192:   "electric.random.switch",
		271048478:   "hat.ratmask",
		-1520560807: "bearmeat",
		-1440987069: "chicken.raw",
		1422530437:  "deermeat.raw",
		989925924:   "fish.raw",
		-1130350864: "horsemeat.raw",
		-1709878924: "humanmeat.raw",
		621915341:   "meat.boar",
		-395377963:  "wolfmeat.raw",
		-1736356576: "target.reactive",
		1376065505:  "vehicle.1mod.rear.seats",
		1272194103:  "red.berry",
		2133269020:  "clone.red.berry",
		830839496:   "seed.red.berry",
		-1553999294: "firework.boomer.red",
		-602717596:  "reddogtags",
		-1160621614: "industrial.wall.light.red",
		-1880870149: "keycard_red",
		-1486461488: "firework.romancandle.red",
		-454370658:  "firework.volcano.red",
		-324675402:  "attire.reindeer.headband",
		671706427:   "wall.window.bars.toptier",
		803222026:   "box.repair.bench",
		-544317637:  "researchpaper",
		-1861522751: "research.table",
		649912614:   "pistol.revolver",
		-1044468317: "electric.rf.broadcaster",
		-566907190:  "rf_pager",
		888415708:   "electric.rf.receiver",
		596469572:   "rf.detonator",
		1394042569:  "rhib",
		176787552:   "riflebody",
		671063303:   "riot.helmet",
		-2002277461: "roadsign.jacket",
		1850456855:  "roadsign.kilt",
		1199391518:  "roadsigns",
		-699558439:  "roadsign.gloves",
		60528587:    "horse.armor.roadsign",
		963906841:   "rock",
		-742865266:  "ammo.rocket.basic",
		442886268:   "rocket.launcher",
		-458565393:  "electrical.combiner",
		1414245522:  "rope",
		352130972:   "apple.spoiled",
		1878053256:  "rowboat",
		-1985799200: "rug",
		-1104881824: "rug.bear",
		-173268132:  "rustige_egg_b",
		-173268125:  "rustige_egg_e",
		-173268126:  "rustige_egg_d",
		-173268131:  "rustige_egg_c",
		-173268129:  "rustige_egg_a",
		1400460850:  "horse.saddlebag",
		-851988960:  "fish.salmon",
		-277057363:  "water.salt",
		-262590403:  "axe.salvaged",
		-1978999529: "salvaged.cleaver",
		-1506397857: "hammer.salvaged",
		-1780802565: "icepick.salvaged",
		1950721418:  "shelves",
		1326180354:  "salvaged.sword",
		-384243979:  "ammo.rocket.sam",
		-1009359066: "samsite",
		-559599960:  "barricade.sandbags",
		2126889441:  "santabeard",
		-575483084:  "santahat",
		-1654233406: "fish.sardine",
		-1878475007: "explosive.satchel",
		177226991:   "scarecrow",
		273951840:   "scarecrow.suit",
		809942731:   "scarecrowhead",
		-253079493:  "hazmatsuit_scientist",
		-1958316066: "hazmatsuit_scientist_peacekeeper",
		-932201673:  "scrap",
		-1884328185: "scraptransportheli.repair",
		2087678962:  "searchlight",
		567871954:   "secretlabchair",
		573926264:   "semibody",
		818877484:   "pistol.semiauto",
		-904863145:  "rifle.semiauto",
		1234880403:  "sewingkit",
		-1994909036: "sheetmetal",
		-2067472972: "door.hinged.metal",
		1390353317:  "door.double.hinged.metal",
		-2025184684: "shirt.collared",
		-796583652:  "wall.frame.shopfront",
		1327005675:  "wall.ice.wall",
		-1695367501: "pants.shorts",
		352499047:   "guntrap",
		-2107018088: "fun.bass",
		-1368584029: "sickle",
		-1850571427: "weapon.mod.silencer",
		1757265204:  "easter.silveregg",
		-855748505:  "weapon.mod.simplesight",
		-282113991:  "electric.simplelight",
		1542290441:  "sign.post.single",
		762289806:   "electric.sirenlight",
		1312843609:  "skull",
		-216116642:  "skulldoorknocker",
		553887414:   "skull_fire_pit",
		-25740268:   "skullspikes.candles",
		-1078639462: "skullspikes.pumpkin",
		-1073015016: "skullspikes",
		-769647921:  "skull.trophy",
		-156748077:  "skull.trophy.table",
		-924959988:  "skull.trophy.jar2",
		971362526:   "skull.trophy.jar",
		1819863051:  "skylantern",
		-1770889433: "skylantern.skylantern.green",
		-1824770114: "skylantern.skylantern.orange",
		831955134:   "skylantern.skylantern.purple",
		-1433390281: "skylantern.skylantern.red",
		-135252633:  "sled.xmas",
		-333406828:  "sled",
		-1754948969: "sleepingbag",
		-2058362263: "smallcandles",
		-44066600:   "vehicle.chassis.2mod",
		1849887541:  "electric.fuelgenerator.small",
		1319617282:  "halloween.lootbag.small",
		1305578813:  "sign.neon.125x125",
		-1293296287: "small.oil.refinery",
		1903654061:  "planter.small",
		-722241321:  "xmas.present.small",
		-692338819:  "electric.battery.rechargable.small",
		-1768880890: "fish.smallshark",
		-369760990:  "stash.small",
		1668858301:  "stocking.small",
		-1878764039: "fish.troutsmall",
		-1039528932: "smallwaterbottle",
		-132247350:  "water.catcher.small",
		-1138208076: "sign.wooden.small",
		-695978112:  "smart.alarm",
		988652725:   "smart.switch",
		1230323789:  "smgbody",
		1263920163:  "grenade.smoke",
		-17123659:   "ammo.rocket.smoke",
		-582782051:  "trap.bear",
		-48090175:   "jacket.snow",
		1358643074:  "snowmachine",
		-363689972:  "snowball",
		1103488722:  "snowballgun",
		1629293099:  "snowman",
		-842267147:  "attire.snowman.helmet",
		-1364246987: "snowmobile",
		-555122905:  "sofa",
		782422285:   "sofa.pattern",
		-187031121:  "submarinesolo",
		-343857907:  "soundlight",
		1784406797:  "fun.tuba",
		-560304835:  "hazmatsuit.spacesuit",
		-41440462:   "shotgun.spas12",
		-1517740219: "speargun",
		-1800345240: "speargun.spear",
		882559853:   "spiderweb",
		-1100422738: "spinner.wheel",
		-563624462:  "electric.splitter",
		-751151717:  "chicken.spoiled",
		1272768630:  "humanmeat.spoiled",
		-1167031859: "wolfmeat.spoiled",
		1885488976:  "spookyspeaker",
		-596876839:  "spraycan",
		-1366326648: "spraycandecal",
		-781014061:  "electric.sprinkler",
		-1331212963: "xmas.decoration.star",
		642482233:   "sticks",
		15388698:    "barricade.stone",
		-1535621066: "fireplace.stone",
		-1583967946: "stonehatchet",
		171931394:   "stone.pickaxe",
		1602646136:  "spear.stone",
		-2099697608: "stones",
		1149964039:  "storage.monitor",
		268565518:   "vehicle.1mod.storage",
		-1614955425: "wall.window.glass.reinforced",
		2104517339:  "strobelight",
		-1581843485: "sulfur",
		-1157596551: "sulfur.ore",
		1258768145:  "sunglasses02black",
		-2103694546: "sunglasses02camo",
		1557173737:  "sunglasses02red",
		-176608084:  "sunglasses03black",
		-1997698639: "sunglasses03chrome",
		-1408336705: "sunglasses03gold",
		352321488:   "sunglasses",
		-465682601:  "stocking.large",
		1397052267:  "supply.signal",
		-239306133:  "submarine.torpedo.rising",
		-1785231475: "halloween.surgeonsuit",
		1975934948:  "surveycharge",
		559147458:   "fishtrap.small",
		1951603367:  "electric.switch",
		223891266:   "tshirt",
		593465182:   "table",
		-1108136649: "tactical.gloves",
		121049755:   "sign.pictureframe.tall",
		1608640313:  "shirt.tanktop",
		1523195708:  "targeting.computer",
		2019042823:  "tarp",
		-626174997:  "vehicle.1mod.taxi",
		73681876:    "techparts",
		1234878710:  "telephone",
		1371909803:  "electric.teslacoil",
		-295829489:  "electric.generator.small",
		-1758372725: "smg.thompson",
		709206314:   "hat.tigermask",
		1248356124:  "explosive.timed",
		665332906:   "electric.timer",
		1768112091:  "snowmobiletomaha",
		-97956382:   "cupboard.tool",
		795236088:   "torch",
		-1671551935: "submarine.torpedo.straight",
		1723747470:  "xmas.decoration.lights",
		2041899972:  "floor.triangle.ladder.hatch",
		-1478445584: "tunalight",
		975983052:   "trophy",
		1205607945:  "sign.hanging",
		-1647846966: "sign.hanging.ornate",
		826309791:   "sign.post.town.roof",
		198438816:   "vending.machine",
		-280223496:  "firework.boomer.violet",
		-99886070:   "firework.romancandle.violet",
		-1538109120: "firework.volcano.violet",
		755224797:   "bottle.vodka",
		996757362:   "wagon",
		-463122489:  "watchtower.wood",
		-1779180711: "water",
		-1863559151: "water.barrel",
		1424075905:  "bucket.water",
		722955039:   "gun.water",
		-119235651:  "waterjug",
		-1815301988: "pistol.water",
		-1284169891: "waterpump",
		2114754781:  "water.purifier",
		-1367281941: "shotgun.waterpipe",
		952603248:   "weapon.mod.flashlight",
		-132516482:  "weapon.mod.lasersight",
		-1101924344: "diving.wetsuit",
		1272430949:  "piano",
		854447607:   "white.berry",
		1533551194:  "clone.white.berry",
		-992286106:  "seed.white.berry",
		261913429:   "firework.volcano",
		-1819763926: "generator.wind.scrap",
		-144417939:  "wiretool",
		-1478212975: "hat.wolf",
		2048317869:  "skull.wolf",
		-151838493:  "wood",
		-2094954543: "wood.armor.helmet",
		832133926:   "wood.armor.pants",
		418081930:   "wood.armor.jacket",
		-1336109173: "door.double.hinged.wood",
		-1023374709: "shutter.wood.a",
		-180129657:  "box.wooden",
		-1234735557: "arrow.wooden",
		866889860:   "barricade.wood",
		1373240771:  "barricade.wood.cover",
		699075597:   "woodcross",
		1729120840:  "door.hinged.wood",
		-92759291:   "spikes.floor",
		1659447559:  "horse.armor.wood",
		-316250604:  "ladder.wooden.wall",
		1540934679:  "spear.wooden",
		-1183726687: "wall.window.bars.wood",
		-810326667:  "workcart",
		1524187186:  "workbench1",
		-41896755:   "workbench2",
		-1607980696: "workbench3",
		1770475779:  "worm",
		204970153:   "wrappedgift",
		1094293920:  "wrappingpaper",
		-996185386:  "sign.pictureframe.xl",
		1293102274:  "electric.xorswitch",
		98508942:    "sign.pictureframe.xxl",
		-211235948:  "xylophone",
		1660145984:  "yellow.berry",
		390728933:   "clone.yellow.berry",
		-520133715:  "seed.yellow.berry",
		680234026:   "fish.yellowperch",
	}

	short_name, ok := idsShortNames[item_id]
	if ok {
		return
	} else {
		err = fmt.Errorf("could not find that id!")
	}
	return
}

func needsAmmoWhatType(item_name string) (needs_ammo bool, ammo_name string, amount int, err error) {
	ammo_name = ""
	weaponToAmmoType := map[string]string{
		"rifle.ak": "ammo.rifle",
	}
	weaponToAmmoAmount := map[string]int{
		"rifle.ak": 10,
	}
	var ok bool
	ammo_name, ok = weaponToAmmoType[item_name]
	if ok && ammo_name != "" {
		needs_ammo = true
		amount, ok = weaponToAmmoAmount[item_name]
	} else if ok {
		needs_ammo = false
	} else {
		fmt.Println("could not find that id!")
	}
	return needs_ammo, ammo_name, amount, err
}
