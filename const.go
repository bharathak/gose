// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/gpl-3.0.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gose

const (
	// SeAunitToKm as declared in swephexp.h:88
	SeAunitToKm = 1.495978707e+08
	// SeAunitToLightyear as declared in swephexp.h:89
	SeAunitToLightyear = 1.5812507409819728e-05
	// SeAunitToParsec as declared in swephexp.h:90
	SeAunitToParsec = 4.848136811095274e-06
	// SeJulCal as declared in swephexp.h:93
	SeJulCal = 0
	// SeGregCal as declared in swephexp.h:94
	SeGregCal = 1
	// SeEclNut as declared in swephexp.h:99
	SeEclNut = -1
	// SeSun as declared in swephexp.h:101
	SeSun = 0
	// SeMoon as declared in swephexp.h:102
	SeMoon = 1
	// SeMercury as declared in swephexp.h:103
	SeMercury = 2
	// SeVenus as declared in swephexp.h:104
	SeVenus = 3
	// SeMars as declared in swephexp.h:105
	SeMars = 4
	// SeJupiter as declared in swephexp.h:106
	SeJupiter = 5
	// SeSaturn as declared in swephexp.h:107
	SeSaturn = 6
	// SeUranus as declared in swephexp.h:108
	SeUranus = 7
	// SeNeptune as declared in swephexp.h:109
	SeNeptune = 8
	// SePluto as declared in swephexp.h:110
	SePluto = 9
	// SeMeanNode as declared in swephexp.h:111
	SeMeanNode = 10
	// SeTrueNode as declared in swephexp.h:112
	SeTrueNode = 11
	// SeMeanApog as declared in swephexp.h:113
	SeMeanApog = 12
	// SeOscuApog as declared in swephexp.h:114
	SeOscuApog = 13
	// SeEarth as declared in swephexp.h:115
	SeEarth = 14
	// SeChiron as declared in swephexp.h:116
	SeChiron = 15
	// SePholus as declared in swephexp.h:117
	SePholus = 16
	// SeCeres as declared in swephexp.h:118
	SeCeres = 17
	// SePallas as declared in swephexp.h:119
	SePallas = 18
	// SeJuno as declared in swephexp.h:120
	SeJuno = 19
	// SeVesta as declared in swephexp.h:121
	SeVesta = 20
	// SeIntpApog as declared in swephexp.h:122
	SeIntpApog = 21
	// SeIntpPerg as declared in swephexp.h:123
	SeIntpPerg = 22
	// SeNplanets as declared in swephexp.h:125
	SeNplanets = 23
	// SePlmoonOffset as declared in swephexp.h:127
	SePlmoonOffset = 9000
	// SeAstOffset as declared in swephexp.h:128
	SeAstOffset = 10000
	// SeVaruna as declared in swephexp.h:129
	SeVaruna = 30000
	// SeFictOffset as declared in swephexp.h:131
	SeFictOffset = 40
	// SeFictOffset1 as declared in swephexp.h:132
	SeFictOffset1 = 39
	// SeFictMax as declared in swephexp.h:133
	SeFictMax = 999
	// SeNfictElem as declared in swephexp.h:134
	SeNfictElem = 15
	// SeCometOffset as declared in swephexp.h:136
	SeCometOffset = 1000
	// SeNallNatPoints as declared in swephexp.h:138
	SeNallNatPoints = 38
	// SeCupido as declared in swephexp.h:141
	SeCupido = 40
	// SeHades as declared in swephexp.h:142
	SeHades = 41
	// SeZeus as declared in swephexp.h:143
	SeZeus = 42
	// SeKronos as declared in swephexp.h:144
	SeKronos = 43
	// SeApollon as declared in swephexp.h:145
	SeApollon = 44
	// SeAdmetos as declared in swephexp.h:146
	SeAdmetos = 45
	// SeVulkanus as declared in swephexp.h:147
	SeVulkanus = 46
	// SePoseidon as declared in swephexp.h:148
	SePoseidon = 47
	// SeIsis as declared in swephexp.h:150
	SeIsis = 48
	// SeNibiru as declared in swephexp.h:151
	SeNibiru = 49
	// SeHarrington as declared in swephexp.h:152
	SeHarrington = 50
	// SeNeptuneLeverrier as declared in swephexp.h:153
	SeNeptuneLeverrier = 51
	// SeNeptuneAdams as declared in swephexp.h:154
	SeNeptuneAdams = 52
	// SePlutoLowell as declared in swephexp.h:155
	SePlutoLowell = 53
	// SePlutoPickering as declared in swephexp.h:156
	SePlutoPickering = 54
	// SeVulcan as declared in swephexp.h:157
	SeVulcan = 55
	// SeWhiteMoon as declared in swephexp.h:158
	SeWhiteMoon = 56
	// SeProserpina as declared in swephexp.h:159
	SeProserpina = 57
	// SeWaldemath as declared in swephexp.h:160
	SeWaldemath = 58
	// SeFixstar as declared in swephexp.h:162
	SeFixstar = -10
	// SeAsc as declared in swephexp.h:164
	SeAsc = 0
	// SeMc as declared in swephexp.h:165
	SeMc = 1
	// SeArmc as declared in swephexp.h:166
	SeArmc = 2
	// SeVertex as declared in swephexp.h:167
	SeVertex = 3
	// SeEquasc as declared in swephexp.h:168
	SeEquasc = 4
	// SeCoasc1 as declared in swephexp.h:169
	SeCoasc1 = 5
	// SeCoasc2 as declared in swephexp.h:170
	SeCoasc2 = 6
	// SePolasc as declared in swephexp.h:171
	SePolasc = 7
	// SeNascmc as declared in swephexp.h:172
	SeNascmc = 8
	// SeflgJpleph as declared in swephexp.h:186
	SeflgJpleph = 1
	// SeflgSwieph as declared in swephexp.h:187
	SeflgSwieph = 2
	// SeflgMoseph as declared in swephexp.h:188
	SeflgMoseph = 4
	// SeflgHelctr as declared in swephexp.h:190
	SeflgHelctr = 8
	// SeflgTruepos as declared in swephexp.h:191
	SeflgTruepos = 16
	// SeflgJ2000 as declared in swephexp.h:192
	SeflgJ2000 = 32
	// SeflgNonut as declared in swephexp.h:193
	SeflgNonut = 64
	// SeflgSpeed3 as declared in swephexp.h:194
	SeflgSpeed3 = 128
	// SeflgSpeed as declared in swephexp.h:196
	SeflgSpeed = 256
	// SeflgNogdefl as declared in swephexp.h:197
	SeflgNogdefl = 512
	// SeflgNoaberr as declared in swephexp.h:198
	SeflgNoaberr = 1024
	// SeflgAstrometric as declared in swephexp.h:199
	SeflgAstrometric = 1536
	// SeflgEquatorial as declared in swephexp.h:202
	SeflgEquatorial = 2048
	// SeflgXyz as declared in swephexp.h:203
	SeflgXyz = 4096
	// SeflgRadians as declared in swephexp.h:204
	SeflgRadians = 8192
	// SeflgBaryctr as declared in swephexp.h:205
	SeflgBaryctr = 16384
	// SeflgTopoctr as declared in swephexp.h:206
	SeflgTopoctr = 32768
	// SeflgOrbelAa as declared in swephexp.h:207
	SeflgOrbelAa = 32768
	// SeflgTropical as declared in swephexp.h:209
	SeflgTropical = 0
	// SeflgSidereal as declared in swephexp.h:210
	SeflgSidereal = 65536
	// SeflgIcrs as declared in swephexp.h:211
	SeflgIcrs = 131072
	// SeflgDpsideps1980 as declared in swephexp.h:212
	SeflgDpsideps1980 = 262144
	// SeflgJplhor as declared in swephexp.h:214
	SeflgJplhor = 262144
	// SeflgJplhorApprox as declared in swephexp.h:215
	SeflgJplhorApprox = 524288
	// SeflgCenterBody as declared in swephexp.h:216
	SeflgCenterBody = 1048576
	// SeflgTestPlmoon as declared in swephexp.h:218
	SeflgTestPlmoon = 2228280
	// SeSidbits as declared in swephexp.h:221
	SeSidbits = 256
	// SeSidbitEclT0 as declared in swephexp.h:223
	SeSidbitEclT0 = 256
	// SeSidbitSsyPlane as declared in swephexp.h:225
	SeSidbitSsyPlane = 512
	// SeSidbitUserUt as declared in swephexp.h:227
	SeSidbitUserUt = 1024
	// SeSidbitEclDate as declared in swephexp.h:230
	SeSidbitEclDate = 2048
	// SeSidbitNoPrecOffset as declared in swephexp.h:233
	SeSidbitNoPrecOffset = 4096
	// SeSidbitPrecOrig as declared in swephexp.h:235
	SeSidbitPrecOrig = 8192
	// SeSidmFaganBradley as declared in swephexp.h:238
	SeSidmFaganBradley = 0
	// SeSidmLahiri as declared in swephexp.h:239
	SeSidmLahiri = 1
	// SeSidmDeluce as declared in swephexp.h:240
	SeSidmDeluce = 2
	// SeSidmRaman as declared in swephexp.h:241
	SeSidmRaman = 3
	// SeSidmUshashashi as declared in swephexp.h:242
	SeSidmUshashashi = 4
	// SeSidmKrishnamurti as declared in swephexp.h:243
	SeSidmKrishnamurti = 5
	// SeSidmDjwhalKhul as declared in swephexp.h:244
	SeSidmDjwhalKhul = 6
	// SeSidmYukteshwar as declared in swephexp.h:245
	SeSidmYukteshwar = 7
	// SeSidmJnBhasin as declared in swephexp.h:246
	SeSidmJnBhasin = 8
	// SeSidmBabylKugler1 as declared in swephexp.h:247
	SeSidmBabylKugler1 = 9
	// SeSidmBabylKugler2 as declared in swephexp.h:248
	SeSidmBabylKugler2 = 10
	// SeSidmBabylKugler3 as declared in swephexp.h:249
	SeSidmBabylKugler3 = 11
	// SeSidmBabylHuber as declared in swephexp.h:250
	SeSidmBabylHuber = 12
	// SeSidmBabylEtpsc as declared in swephexp.h:251
	SeSidmBabylEtpsc = 13
	// SeSidmAldebaran15tau as declared in swephexp.h:252
	SeSidmAldebaran15tau = 14
	// SeSidmHipparchos as declared in swephexp.h:253
	SeSidmHipparchos = 15
	// SeSidmSassanian as declared in swephexp.h:254
	SeSidmSassanian = 16
	// SeSidmGalcent0sag as declared in swephexp.h:255
	SeSidmGalcent0sag = 17
	// SeSidmJ2000 as declared in swephexp.h:256
	SeSidmJ2000 = 18
	// SeSidmJ1900 as declared in swephexp.h:257
	SeSidmJ1900 = 19
	// SeSidmB1950 as declared in swephexp.h:258
	SeSidmB1950 = 20
	// SeSidmSuryasiddhanta as declared in swephexp.h:259
	SeSidmSuryasiddhanta = 21
	// SeSidmSuryasiddhantaMsun as declared in swephexp.h:260
	SeSidmSuryasiddhantaMsun = 22
	// SeSidmAryabhata as declared in swephexp.h:261
	SeSidmAryabhata = 23
	// SeSidmAryabhataMsun as declared in swephexp.h:262
	SeSidmAryabhataMsun = 24
	// SeSidmSsRevati as declared in swephexp.h:263
	SeSidmSsRevati = 25
	// SeSidmSsCitra as declared in swephexp.h:264
	SeSidmSsCitra = 26
	// SeSidmTrueCitra as declared in swephexp.h:265
	SeSidmTrueCitra = 27
	// SeSidmTrueRevati as declared in swephexp.h:266
	SeSidmTrueRevati = 28
	// SeSidmTruePushya as declared in swephexp.h:267
	SeSidmTruePushya = 29
	// SeSidmGalcentRgilbrand as declared in swephexp.h:268
	SeSidmGalcentRgilbrand = 30
	// SeSidmGalequIau1958 as declared in swephexp.h:269
	SeSidmGalequIau1958 = 31
	// SeSidmGalequTrue as declared in swephexp.h:270
	SeSidmGalequTrue = 32
	// SeSidmGalequMula as declared in swephexp.h:271
	SeSidmGalequMula = 33
	// SeSidmGalalignMardyks as declared in swephexp.h:272
	SeSidmGalalignMardyks = 34
	// SeSidmTrueMula as declared in swephexp.h:273
	SeSidmTrueMula = 35
	// SeSidmGalcentMulaWilhelm as declared in swephexp.h:274
	SeSidmGalcentMulaWilhelm = 36
	// SeSidmAryabhata522 as declared in swephexp.h:275
	SeSidmAryabhata522 = 37
	// SeSidmBabylBritton as declared in swephexp.h:276
	SeSidmBabylBritton = 38
	// SeSidmTrueSheoran as declared in swephexp.h:277
	SeSidmTrueSheoran = 39
	// SeSidmGalcentCochrane as declared in swephexp.h:278
	SeSidmGalcentCochrane = 40
	// SeSidmGalequFiorenza as declared in swephexp.h:279
	SeSidmGalequFiorenza = 41
	// SeSidmValensMoon as declared in swephexp.h:280
	SeSidmValensMoon = 42
	// SeSidmLahiri1940 as declared in swephexp.h:281
	SeSidmLahiri1940 = 43
	// SeSidmLahiriVp285 as declared in swephexp.h:282
	SeSidmLahiriVp285 = 44
	// SeSidmKrishnamurtiVp291 as declared in swephexp.h:283
	SeSidmKrishnamurtiVp291 = 45
	// SeSidmLahiriIcrc as declared in swephexp.h:284
	SeSidmLahiriIcrc = 46
	// SeSidmUser as declared in swephexp.h:286
	SeSidmUser = 255
	// SeNsidmPredef as declared in swephexp.h:288
	SeNsidmPredef = 47
	// SeNodbitMean as declared in swephexp.h:291
	SeNodbitMean = 1
	// SeNodbitOscu as declared in swephexp.h:292
	SeNodbitOscu = 2
	// SeNodbitOscuBar as declared in swephexp.h:293
	SeNodbitOscuBar = 4
	// SeNodbitFopoint as declared in swephexp.h:294
	SeNodbitFopoint = 256
	// SeflgDefaulteph as declared in swephexp.h:297
	SeflgDefaulteph = 2
	// SeMaxStname as declared in swephexp.h:299
	SeMaxStname = 256
	// SeEclCentral as declared in swephexp.h:307
	SeEclCentral = 1
	// SeEclNoncentral as declared in swephexp.h:308
	SeEclNoncentral = 2
	// SeEclTotal as declared in swephexp.h:309
	SeEclTotal = 4
	// SeEclAnnular as declared in swephexp.h:310
	SeEclAnnular = 8
	// SeEclPartial as declared in swephexp.h:311
	SeEclPartial = 16
	// SeEclAnnularTotal as declared in swephexp.h:312
	SeEclAnnularTotal = 32
	// SeEclHybrid as declared in swephexp.h:313
	SeEclHybrid = 32
	// SeEclPenumbral as declared in swephexp.h:314
	SeEclPenumbral = 64
	// SeEclAlltypesSolar as declared in swephexp.h:315
	SeEclAlltypesSolar = 63
	// SeEclAlltypesLunar as declared in swephexp.h:316
	SeEclAlltypesLunar = 84
	// SeEclVisible as declared in swephexp.h:317
	SeEclVisible = 128
	// SeEclMaxVisible as declared in swephexp.h:318
	SeEclMaxVisible = 256
	// SeEcl1stVisible as declared in swephexp.h:319
	SeEcl1stVisible = 512
	// SeEclPartbegVisible as declared in swephexp.h:320
	SeEclPartbegVisible = 512
	// SeEcl2ndVisible as declared in swephexp.h:321
	SeEcl2ndVisible = 1024
	// SeEclTotbegVisible as declared in swephexp.h:322
	SeEclTotbegVisible = 1024
	// SeEcl3rdVisible as declared in swephexp.h:323
	SeEcl3rdVisible = 2048
	// SeEclTotendVisible as declared in swephexp.h:324
	SeEclTotendVisible = 2048
	// SeEcl4thVisible as declared in swephexp.h:325
	SeEcl4thVisible = 4096
	// SeEclPartendVisible as declared in swephexp.h:326
	SeEclPartendVisible = 4096
	// SeEclPenumbbegVisible as declared in swephexp.h:327
	SeEclPenumbbegVisible = 8192
	// SeEclPenumbendVisible as declared in swephexp.h:328
	SeEclPenumbendVisible = 16384
	// SeEclOccBegDaylight as declared in swephexp.h:329
	SeEclOccBegDaylight = 8192
	// SeEclOccEndDaylight as declared in swephexp.h:330
	SeEclOccEndDaylight = 16384
	// SeEclOneTry as declared in swephexp.h:331
	SeEclOneTry = 32768
	// SeCalcRise as declared in swephexp.h:336
	SeCalcRise = 1
	// SeCalcSet as declared in swephexp.h:337
	SeCalcSet = 2
	// SeCalcMtransit as declared in swephexp.h:338
	SeCalcMtransit = 4
	// SeCalcItransit as declared in swephexp.h:339
	SeCalcItransit = 8
	// SeBitDiscCenter as declared in swephexp.h:340
	SeBitDiscCenter = 256
	// SeBitDiscBottom as declared in swephexp.h:343
	SeBitDiscBottom = 8192
	// SeBitGeoctrNoEclLat as declared in swephexp.h:346
	SeBitGeoctrNoEclLat = 128
	// SeBitNoRefraction as declared in swephexp.h:349
	SeBitNoRefraction = 512
	// SeBitCivilTwilight as declared in swephexp.h:351
	SeBitCivilTwilight = 1024
	// SeBitNauticTwilight as declared in swephexp.h:352
	SeBitNauticTwilight = 2048
	// SeBitAstroTwilight as declared in swephexp.h:353
	SeBitAstroTwilight = 4096
	// SeBitFixedDiscSize as declared in swephexp.h:354
	SeBitFixedDiscSize = 16384
	// SeBitForceSlowMethod as declared in swephexp.h:357
	SeBitForceSlowMethod = 32768
	// SeBitHinduRising as declared in swephexp.h:361
	SeBitHinduRising = 896
	// SeEcl2hor as declared in swephexp.h:364
	SeEcl2hor = 0
	// SeEqu2hor as declared in swephexp.h:365
	SeEqu2hor = 1
	// SeHor2ecl as declared in swephexp.h:366
	SeHor2ecl = 0
	// SeHor2equ as declared in swephexp.h:367
	SeHor2equ = 1
	// SeTrueToApp as declared in swephexp.h:370
	SeTrueToApp = 0
	// SeAppToTrue as declared in swephexp.h:371
	SeAppToTrue = 1
	// SeDeNumber as declared in swephexp.h:377
	SeDeNumber = 431
	// SeFnameDe200 as declared in swephexp.h:378
	SeFnameDe200 = "de200.eph"
	// SeFnameDe403 as declared in swephexp.h:379
	SeFnameDe403 = "de403.eph"
	// SeFnameDe404 as declared in swephexp.h:380
	SeFnameDe404 = "de404.eph"
	// SeFnameDe405 as declared in swephexp.h:381
	SeFnameDe405 = "de405.eph"
	// SeFnameDe406 as declared in swephexp.h:382
	SeFnameDe406 = "de406.eph"
	// SeFnameDe431 as declared in swephexp.h:383
	SeFnameDe431 = "de431.eph"
	// SeFnameDft as declared in swephexp.h:384
	SeFnameDft = "de431.eph"
	// SeFnameDft2 as declared in swephexp.h:385
	SeFnameDft2 = "de406.eph"
	// SeStarfileOld as declared in swephexp.h:386
	SeStarfileOld = "fixstars.cat"
	// SeStarfile as declared in swephexp.h:387
	SeStarfile = "sefstars.txt"
	// SeAstnamfile as declared in swephexp.h:388
	SeAstnamfile = "seasnam.txt"
	// SeFictfile as declared in swephexp.h:389
	SeFictfile = "seorbel.txt"
	// SeEphePath as declared in swephexp.h:403
	SeEphePath = ".:/users/ephe2/:/users/ephe/"
	// SeSplitDegRoundSec as declared in swephexp.h:411
	SeSplitDegRoundSec = 1
	// SeSplitDegRoundMin as declared in swephexp.h:412
	SeSplitDegRoundMin = 2
	// SeSplitDegRoundDeg as declared in swephexp.h:413
	SeSplitDegRoundDeg = 4
	// SeSplitDegZodiacal as declared in swephexp.h:414
	SeSplitDegZodiacal = 8
	// SeSplitDegNakshatra as declared in swephexp.h:415
	SeSplitDegNakshatra = 1024
	// SeSplitDegKeepSign as declared in swephexp.h:416
	SeSplitDegKeepSign = 16
	// SeSplitDegKeepDeg as declared in swephexp.h:419
	SeSplitDegKeepDeg = 32
	// SeHeliacalRising as declared in swephexp.h:424
	SeHeliacalRising = 1
	// SeHeliacalSetting as declared in swephexp.h:425
	SeHeliacalSetting = 2
	// SeMorningFirst as declared in swephexp.h:426
	SeMorningFirst = 1
	// SeEveningLast as declared in swephexp.h:427
	SeEveningLast = 2
	// SeEveningFirst as declared in swephexp.h:428
	SeEveningFirst = 3
	// SeMorningLast as declared in swephexp.h:429
	SeMorningLast = 4
	// SeAcronychalRising as declared in swephexp.h:430
	SeAcronychalRising = 5
	// SeAcronychalSetting as declared in swephexp.h:431
	SeAcronychalSetting = 6
	// SeCosmicalSetting as declared in swephexp.h:432
	SeCosmicalSetting = 6
	// SeHelflagLongSearch as declared in swephexp.h:434
	SeHelflagLongSearch = 128
	// SeHelflagHighPrecision as declared in swephexp.h:435
	SeHelflagHighPrecision = 256
	// SeHelflagOpticalParams as declared in swephexp.h:436
	SeHelflagOpticalParams = 512
	// SeHelflagNoDetails as declared in swephexp.h:437
	SeHelflagNoDetails = 1024
	// SeHelflagSearch1Period as declared in swephexp.h:438
	SeHelflagSearch1Period = 2048
	// SeHelflagVislimDark as declared in swephexp.h:439
	SeHelflagVislimDark = 4096
	// SeHelflagVislimNomoon as declared in swephexp.h:440
	SeHelflagVislimNomoon = 8192
	// SeHelflagVislimPhotopic as declared in swephexp.h:442
	SeHelflagVislimPhotopic = 16384
	// SeHelflagVislimScotopic as declared in swephexp.h:443
	SeHelflagVislimScotopic = 32768
	// SeHelflagAv as declared in swephexp.h:444
	SeHelflagAv = 65536
	// SeHelflagAvkindVr as declared in swephexp.h:445
	SeHelflagAvkindVr = 65536
	// SeHelflagAvkindPto as declared in swephexp.h:446
	SeHelflagAvkindPto = 131072
	// SeHelflagAvkindMin7 as declared in swephexp.h:447
	SeHelflagAvkindMin7 = 262144
	// SeHelflagAvkindMin9 as declared in swephexp.h:448
	SeHelflagAvkindMin9 = 524288
	// SeHelflagAvkind as declared in swephexp.h:449
	SeHelflagAvkind = 983040
	// TjdInvalid as declared in swephexp.h:450
	TjdInvalid = 9.9999999e+07
	// SimulateVictorvb as declared in swephexp.h:451
	SimulateVictorvb = 1
	// SePhotopicFlag as declared in swephexp.h:469
	SePhotopicFlag = 0
	// SeScotopicFlag as declared in swephexp.h:470
	SeScotopicFlag = 1
	// SeMixedopicFlag as declared in swephexp.h:471
	SeMixedopicFlag = 2
	// SeTidalDe200 as declared in swephexp.h:478
	SeTidalDe200 = -23.8946
	// SeTidalDe403 as declared in swephexp.h:479
	SeTidalDe403 = -25.58
	// SeTidalDe404 as declared in swephexp.h:480
	SeTidalDe404 = -25.58
	// SeTidalDe405 as declared in swephexp.h:481
	SeTidalDe405 = -25.826
	// SeTidalDe406 as declared in swephexp.h:482
	SeTidalDe406 = -25.826
	// SeTidalDe421 as declared in swephexp.h:483
	SeTidalDe421 = -25.85
	// SeTidalDe422 as declared in swephexp.h:484
	SeTidalDe422 = -25.85
	// SeTidalDe430 as declared in swephexp.h:485
	SeTidalDe430 = -25.82
	// SeTidalDe431 as declared in swephexp.h:486
	SeTidalDe431 = -25.8
	// SeTidalDe441 as declared in swephexp.h:487
	SeTidalDe441 = -25.936
	// SeTidal26 as declared in swephexp.h:488
	SeTidal26 = -26
	// SeTidalStephenson2016 as declared in swephexp.h:489
	SeTidalStephenson2016 = -25.85
	// SeTidalDefault as declared in swephexp.h:490
	SeTidalDefault = -25.8
	// SeTidalAutomatic as declared in swephexp.h:491
	SeTidalAutomatic = 999999
	// SeTidalMoseph as declared in swephexp.h:492
	SeTidalMoseph = -25.58
	// SeTidalSwieph as declared in swephexp.h:493
	SeTidalSwieph = -25.8
	// SeTidalJpleph as declared in swephexp.h:494
	SeTidalJpleph = -25.8
	// SeDeltatAutomatic as declared in swephexp.h:497
	SeDeltatAutomatic = -1e-10
	// SeModelDeltat as declared in swephexp.h:499
	SeModelDeltat = 0
	// SeModelPrecLongterm as declared in swephexp.h:500
	SeModelPrecLongterm = 1
	// SeModelPrecShortterm as declared in swephexp.h:501
	SeModelPrecShortterm = 2
	// SeModelNut as declared in swephexp.h:502
	SeModelNut = 3
	// SeModelBias as declared in swephexp.h:503
	SeModelBias = 4
	// SeModelJplhorMode as declared in swephexp.h:504
	SeModelJplhorMode = 5
	// SeModelJplhoraMode as declared in swephexp.h:505
	SeModelJplhoraMode = 6
	// SeModelSidt as declared in swephexp.h:506
	SeModelSidt = 7
	// NseModels as declared in swephexp.h:507
	NseModels = 8
	// SemodNprec as declared in swephexp.h:510
	SemodNprec = 11
	// SemodPrecIau1976 as declared in swephexp.h:511
	SemodPrecIau1976 = 1
	// SemodPrecLaskar1986 as declared in swephexp.h:512
	SemodPrecLaskar1986 = 2
	// SemodPrecWillEpsLask as declared in swephexp.h:513
	SemodPrecWillEpsLask = 3
	// SemodPrecWilliams1994 as declared in swephexp.h:514
	SemodPrecWilliams1994 = 4
	// SemodPrecSimon1994 as declared in swephexp.h:515
	SemodPrecSimon1994 = 5
	// SemodPrecIau2000 as declared in swephexp.h:516
	SemodPrecIau2000 = 6
	// SemodPrecBretagnon2003 as declared in swephexp.h:517
	SemodPrecBretagnon2003 = 7
	// SemodPrecIau2006 as declared in swephexp.h:518
	SemodPrecIau2006 = 8
	// SemodPrecVondrak2011 as declared in swephexp.h:519
	SemodPrecVondrak2011 = 9
	// SemodPrecOwen1990 as declared in swephexp.h:520
	SemodPrecOwen1990 = 10
	// SemodPrecNewcomb as declared in swephexp.h:521
	SemodPrecNewcomb = 11
	// SemodPrecDefault as declared in swephexp.h:522
	SemodPrecDefault = 9
	// SemodPrecDefaultShort as declared in swephexp.h:527
	SemodPrecDefaultShort = 9
	// SemodNnut as declared in swephexp.h:530
	SemodNnut = 5
	// SemodNutIau1980 as declared in swephexp.h:531
	SemodNutIau1980 = 1
	// SemodNutIauCorr1987 as declared in swephexp.h:532
	SemodNutIauCorr1987 = 2
	// SemodNutIau2000a as declared in swephexp.h:534
	SemodNutIau2000a = 3
	// SemodNutIau2000b as declared in swephexp.h:535
	SemodNutIau2000b = 4
	// SemodNutWoolard as declared in swephexp.h:536
	SemodNutWoolard = 5
	// SemodNutDefault as declared in swephexp.h:537
	SemodNutDefault = 4
	// SemodNsidt as declared in swephexp.h:540
	SemodNsidt = 4
	// SemodSidtIau1976 as declared in swephexp.h:541
	SemodSidtIau1976 = 1
	// SemodSidtIau2006 as declared in swephexp.h:542
	SemodSidtIau2006 = 2
	// SemodSidtIersConv2010 as declared in swephexp.h:543
	SemodSidtIersConv2010 = 3
	// SemodSidtLongterm as declared in swephexp.h:544
	SemodSidtLongterm = 4
	// SemodSidtDefault as declared in swephexp.h:545
	SemodSidtDefault = 4
	// SemodNbias as declared in swephexp.h:549
	SemodNbias = 3
	// SemodBiasNone as declared in swephexp.h:550
	SemodBiasNone = 1
	// SemodBiasIau2000 as declared in swephexp.h:551
	SemodBiasIau2000 = 2
	// SemodBiasIau2006 as declared in swephexp.h:552
	SemodBiasIau2006 = 3
	// SemodBiasDefault as declared in swephexp.h:553
	SemodBiasDefault = 3
	// SemodNjplhor as declared in swephexp.h:557
	SemodNjplhor = 2
	// SemodJplhorLongAgreement as declared in swephexp.h:558
	SemodJplhorLongAgreement = 1
	// SemodJplhorDefault as declared in swephexp.h:562
	SemodJplhorDefault = 1
	// SemodNjplhora as declared in swephexp.h:577
	SemodNjplhora = 3
	// SemodJplhora1 as declared in swephexp.h:578
	SemodJplhora1 = 1
	// SemodJplhora2 as declared in swephexp.h:579
	SemodJplhora2 = 2
	// SemodJplhora3 as declared in swephexp.h:580
	SemodJplhora3 = 3
	// SemodJplhoraDefault as declared in swephexp.h:581
	SemodJplhoraDefault = 3
	// SemodNdeltat as declared in swephexp.h:599
	SemodNdeltat = 5
	// SemodDeltatStephensonMorrison1984 as declared in swephexp.h:600
	SemodDeltatStephensonMorrison1984 = 1
	// SemodDeltatStephenson1997 as declared in swephexp.h:601
	SemodDeltatStephenson1997 = 2
	// SemodDeltatStephensonMorrison2004 as declared in swephexp.h:602
	SemodDeltatStephensonMorrison2004 = 3
	// SemodDeltatEspenakMeeus2006 as declared in swephexp.h:603
	SemodDeltatEspenakMeeus2006 = 4
	// SemodDeltatStephensonEtc2016 as declared in swephexp.h:604
	SemodDeltatStephensonEtc2016 = 5
	// SemodDeltatDefault as declared in swephexp.h:606
	SemodDeltatDefault = 5
	// Malloc as declared in swephexp.h:638
	Malloc = 0
	// Calloc as declared in swephexp.h:639
	Calloc = 0
	// Free as declared in swephexp.h:640
	Free = 0
	// CallConv as declared in swephexp.h:662

	// Exp32 as declared in swephexp.h:663

	// MyTrue as defined in sweodef.h:74
	MyTrue = 1
	// MyFalse as defined in sweodef.h:75
	MyFalse = 0
	// Tls as defined in sweodef.h:84
	Tls = 0
	// Msdos as defined in sweodef.h:157
	Msdos = 0
	// Hpunix as defined in sweodef.h:158
	Hpunix = 1
	// UnixFs as defined in sweodef.h:162
	UnixFs = 1
	// Abs4 as defined in sweodef.h:220
	Abs4 = 0
	// True as defined in sweodef.h:233
	True = 1
	// False as defined in sweodef.h:234
	False = 0
	// Ok as defined in sweodef.h:238
	Ok = 0
	// Err as defined in sweodef.h:239
	Err = -1
	// Ucp as defined in sweodef.h:250
	//Ucp = ( UCHAR * )( )
	// OdegreeString as defined in sweodef.h:253
	OdegreeString = "°"
	// AsMaxch as defined in sweodef.h:266
	AsMaxch = 256
	// Radtodeg as defined in sweodef.h:272
	Radtodeg = 57.29577951308232
	// Degtorad as defined in sweodef.h:273
	Degtorad = 0.017453292519943295
	// Cs as defined in sweodef.h:276
	Cs = 0
	// Csec as defined in sweodef.h:277
	Csec = 0
	// Deg as defined in sweodef.h:279
	Deg = 360000
	// Deg730 as defined in sweodef.h:280
	Deg730 = 2700000
	// Deg15 as defined in sweodef.h:281
	Deg15 = 5400000
	// Deg24 as defined in sweodef.h:282
	Deg24 = 8640000
	// Deg30 as defined in sweodef.h:283
	Deg30 = 10800000
	// Deg60 as defined in sweodef.h:284
	Deg60 = 21600000
	// Deg90 as defined in sweodef.h:285
	Deg90 = 32400000
	// Deg120 as defined in sweodef.h:286
	Deg120 = 43200000
	// Deg150 as defined in sweodef.h:287
	Deg150 = 54000000
	// Deg180 as defined in sweodef.h:288
	Deg180 = 64800000
	// Deg270 as defined in sweodef.h:289
	Deg270 = 97200000
	// Deg360 as defined in sweodef.h:290
	Deg360 = 129600000
	// Cstorad as defined in sweodef.h:294
	Cstorad = 4.84813681109536e-08
	// Radtocs as defined in sweodef.h:295
	Radtocs = 2.0626480624709636e+07
	// Cs2deg as defined in sweodef.h:297
	Cs2deg = 2.777777777777778e-06
	// BfileRAccess as defined in sweodef.h:301
	BfileRAccess = "r"
	// BfileRwAccess as defined in sweodef.h:302
	BfileRwAccess = "r+"
	// BfileWCreate as defined in sweodef.h:303
	BfileWCreate = "w"
	// BfileAAccess as defined in sweodef.h:304
	BfileAAccess = "a+"
	// FileRAccess as defined in sweodef.h:305
	FileRAccess = "r"
	// FileRwAccess as defined in sweodef.h:306
	FileRwAccess = "r+"
	// FileWCreate as defined in sweodef.h:307
	FileWCreate = "w"
	// FileAAccess as defined in sweodef.h:308
	FileAAccess = "a+"
	// OBinary as defined in sweodef.h:309
	OBinary = 0
	// OpenMode as defined in sweodef.h:310
	OpenMode = 438
	// DirGlue as defined in sweodef.h:311
	DirGlue = "/"
	// PathSeparator as defined in sweodef.h:312
	PathSeparator = ";:"

	// SeVersion as defined in sweph.h:65
	SeVersion = "2.10.01"
	// J2000 as defined in sweph.h:67
	J2000 = 2.451545e+06
	// B1950 as defined in sweph.h:68
	B1950 = 2.43328242345905e+06
	// J1900 as defined in sweph.h:69
	J1900 = 2.41502e+06
	// B1850 as defined in sweph.h:70
	B1850 = 2.396758203581e+06
	// MpcCeres as defined in sweph.h:72
	MpcCeres = 1
	// MpcPallas as defined in sweph.h:73
	MpcPallas = 2
	// MpcJuno as defined in sweph.h:74
	MpcJuno = 3
	// MpcVesta as defined in sweph.h:75
	MpcVesta = 4
	// MpcChiron as defined in sweph.h:76
	MpcChiron = 2060
	// MpcPholus as defined in sweph.h:77
	MpcPholus = 5145
	// SeNameSun as defined in sweph.h:79
	SeNameSun = "Sun"
	// SeNameMoon as defined in sweph.h:80
	SeNameMoon = "Moon"
	// SeNameMercury as defined in sweph.h:81
	SeNameMercury = "Mercury"
	// SeNameVenus as defined in sweph.h:82
	SeNameVenus = "Venus"
	// SeNameMars as defined in sweph.h:83
	SeNameMars = "Mars"
	// SeNameJupiter as defined in sweph.h:84
	SeNameJupiter = "Jupiter"
	// SeNameSaturn as defined in sweph.h:85
	SeNameSaturn = "Saturn"
	// SeNameUranus as defined in sweph.h:86
	SeNameUranus = "Uranus"
	// SeNameNeptune as defined in sweph.h:87
	SeNameNeptune = "Neptune"
	// SeNamePluto as defined in sweph.h:88
	SeNamePluto = "Pluto"
	// SeNameMeanNode as defined in sweph.h:89
	SeNameMeanNode = "mean Node"
	// SeNameTrueNode as defined in sweph.h:90
	SeNameTrueNode = "true Node"
	// SeNameMeanApog as defined in sweph.h:91
	SeNameMeanApog = "mean Apogee"
	// SeNameOscuApog as defined in sweph.h:92
	SeNameOscuApog = "osc. Apogee"
	// SeNameIntpApog as defined in sweph.h:93
	SeNameIntpApog = "intp. Apogee"
	// SeNameIntpPerg as defined in sweph.h:94
	SeNameIntpPerg = "intp. Perigee"
	// SeNameEarth as defined in sweph.h:95
	SeNameEarth = "Earth"
	// SeNameCeres as defined in sweph.h:96
	SeNameCeres = "Ceres"
	// SeNamePallas as defined in sweph.h:97
	SeNamePallas = "Pallas"
	// SeNameJuno as defined in sweph.h:98
	SeNameJuno = "Juno"
	// SeNameVesta as defined in sweph.h:99
	SeNameVesta = "Vesta"
	// SeNameChiron as defined in sweph.h:100
	SeNameChiron = "Chiron"
	// SeNamePholus as defined in sweph.h:101
	SeNamePholus = "Pholus"
	// SeNameCupido as defined in sweph.h:104
	SeNameCupido = "Cupido"
	// SeNameHades as defined in sweph.h:105
	SeNameHades = "Hades"
	// SeNameZeus as defined in sweph.h:106
	SeNameZeus = "Zeus"
	// SeNameKronos as defined in sweph.h:107
	SeNameKronos = "Kronos"
	// SeNameApollon as defined in sweph.h:108
	SeNameApollon = "Apollon"
	// SeNameAdmetos as defined in sweph.h:109
	SeNameAdmetos = "Admetos"
	// SeNameVulkanus as defined in sweph.h:110
	SeNameVulkanus = "Vulkanus"
	// SeNamePoseidon as defined in sweph.h:111
	SeNamePoseidon = "Poseidon"
	// SeNameIsis as defined in sweph.h:112
	SeNameIsis = "Isis"
	// SeNameNibiru as defined in sweph.h:113
	SeNameNibiru = "Nibiru"
	// SeNameHarrington as defined in sweph.h:114
	SeNameHarrington = "Harrington"
	// SeNameNeptuneLeverrier as defined in sweph.h:115
	SeNameNeptuneLeverrier = "Leverrier"
	// SeNameNeptuneAdams as defined in sweph.h:116
	SeNameNeptuneAdams = "Adams"
	// SeNamePlutoLowell as defined in sweph.h:117
	SeNamePlutoLowell = "Lowell"
	// SeNamePlutoPickering as defined in sweph.h:118
	SeNamePlutoPickering = "Pickering"
	// SeNameVulcan as defined in sweph.h:119
	SeNameVulcan = "Vulcan"
	// SeNameWhiteMoon as defined in sweph.h:120
	SeNameWhiteMoon = "White Moon"
	// Pi as defined in sweph.h:126
	Pi = 3.141592653589793
	// Twopi as defined in sweph.h:127
	Twopi = 6.283185307179586
	// Endmark as defined in sweph.h:129
	Endmark = -99
	// SeiEpsilon as defined in sweph.h:131
	SeiEpsilon = -2
	// SeiNutation as defined in sweph.h:132
	SeiNutation = -1
	// SeiEmb as defined in sweph.h:133
	SeiEmb = 0
	// SeiEarth as defined in sweph.h:134
	SeiEarth = 0
	// SeiSun as defined in sweph.h:135
	SeiSun = 0
	// SeiMoon as defined in sweph.h:136
	SeiMoon = 1
	// SeiMercury as defined in sweph.h:137
	SeiMercury = 2
	// SeiVenus as defined in sweph.h:138
	SeiVenus = 3
	// SeiMars as defined in sweph.h:139
	SeiMars = 4
	// SeiJupiter as defined in sweph.h:140
	SeiJupiter = 5
	// SeiSaturn as defined in sweph.h:141
	SeiSaturn = 6
	// SeiUranus as defined in sweph.h:142
	SeiUranus = 7
	// SeiNeptune as defined in sweph.h:143
	SeiNeptune = 8
	// SeiPluto as defined in sweph.h:144
	SeiPluto = 9
	// SeiSunbary as defined in sweph.h:145
	SeiSunbary = 10
	// SeiAnybody as defined in sweph.h:146
	SeiAnybody = 11
	// SeiChiron as defined in sweph.h:147
	SeiChiron = 12
	// SeiPholus as defined in sweph.h:148
	SeiPholus = 13
	// SeiCeres as defined in sweph.h:149
	SeiCeres = 14
	// SeiPallas as defined in sweph.h:150
	SeiPallas = 15
	// SeiJuno as defined in sweph.h:151
	SeiJuno = 16
	// SeiVesta as defined in sweph.h:152
	SeiVesta = 17
	// SeiNplanets as defined in sweph.h:154
	SeiNplanets = 18
	// SeiMeanNode as defined in sweph.h:156
	SeiMeanNode = 0
	// SeiTrueNode as defined in sweph.h:157
	SeiTrueNode = 1
	// SeiMeanApog as defined in sweph.h:158
	SeiMeanApog = 2
	// SeiOscuApog as defined in sweph.h:159
	SeiOscuApog = 3
	// SeiIntpApog as defined in sweph.h:160
	SeiIntpApog = 4
	// SeiIntpPerg as defined in sweph.h:161
	SeiIntpPerg = 5
	// SeiNnodeEtc as defined in sweph.h:163
	SeiNnodeEtc = 6
	// SeiFlgHelio as defined in sweph.h:165
	SeiFlgHelio = 1
	// SeiFlgRotate as defined in sweph.h:166
	SeiFlgRotate = 2
	// SeiFlgEllipse as defined in sweph.h:167
	SeiFlgEllipse = 4
	// SeiFlgEmbhel as defined in sweph.h:168
	SeiFlgEmbhel = 8
	// SeiFilePlanet as defined in sweph.h:173
	SeiFilePlanet = 0
	// SeiFileMoon as defined in sweph.h:174
	SeiFileMoon = 1
	// SeiFileMainAst as defined in sweph.h:175
	SeiFileMainAst = 2
	// SeiFileAnyAst as defined in sweph.h:176
	SeiFileAnyAst = 3
	// SeiFileFixstar as defined in sweph.h:177
	SeiFileFixstar = 4
	// SeiFilePlmoon as defined in sweph.h:178
	SeiFilePlmoon = 5
	// SeiFileTestEndian as defined in sweph.h:183
	SeiFileTestEndian = 6382179
	// SeiFileBigendian as defined in sweph.h:184
	SeiFileBigendian = 0
	// SeiFileNoreord as defined in sweph.h:185
	SeiFileNoreord = 0
	// SeiFileLitendian as defined in sweph.h:186
	SeiFileLitendian = 1
	// SeiFileReord as defined in sweph.h:187
	SeiFileReord = 2
	// SeiFileNmaxplan as defined in sweph.h:189
	SeiFileNmaxplan = 50
	// SeiFileEfposbegin as defined in sweph.h:190
	SeiFileEfposbegin = 500
	// SeFileSuffix as defined in sweph.h:192
	SeFileSuffix = "se1"
	// SeiNephfiles as defined in sweph.h:194
	SeiNephfiles = 7
	// SeiCurrFpos as defined in sweph.h:195
	SeiCurrFpos = -1
	// SeiNmodels as defined in sweph.h:196
	SeiNmodels = 8
	// SeiEclGeoaltMax as defined in sweph.h:198
	SeiEclGeoaltMax = 25000
	// SeiEclGeoaltMin as defined in sweph.h:199
	SeiEclGeoaltMin = -500
	// ChironStart as defined in sweph.h:207
	ChironStart = 1.9676015e+06
	// ChironEnd as defined in sweph.h:208
	ChironEnd = 3.4194375e+06
	// PholusStart as defined in sweph.h:216
	PholusStart = 640648.5
	// PholusEnd as defined in sweph.h:217
	PholusEnd = 4.3906175e+06
	// MoshplephStart as defined in sweph.h:219
	MoshplephStart = 625000.5
	// MoshplephEnd as defined in sweph.h:220
	MoshplephEnd = 2.8180005e+06
	// MoshluephStart as defined in sweph.h:221
	MoshluephStart = 625000.5
	// MoshluephEnd as defined in sweph.h:222
	MoshluephEnd = 2.8180005e+06
	// MoshndephStart as defined in sweph.h:225
	MoshndephStart = -3.1000155e+06
	// MoshndephEnd as defined in sweph.h:226
	MoshndephEnd = 8.0000165e+06
	// JplDe431Start as defined in sweph.h:233
	JplDe431Start = -3.0272155e+06
	// JplDe431End as defined in sweph.h:234
	JplDe431End = 7.9301925e+06
	// Maxord as defined in sweph.h:247
	Maxord = 40
	// Ncties as defined in sweph.h:249
	Ncties = 6
	// NotAvailable as defined in sweph.h:253
	NotAvailable = -2
	// BeyondEphLimits as defined in sweph.h:254
	BeyondEphLimits = -3
	// JToJ2000 as defined in sweph.h:256
	JToJ2000 = 1
	// J2000ToJ as defined in sweph.h:257
	J2000ToJ = -1
	// MoonMeanDist as defined in sweph.h:260
	MoonMeanDist = 3.844e+08
	// MoonMeanIncl as defined in sweph.h:261
	MoonMeanIncl = 5.1453964
	// MoonMeanEcc as defined in sweph.h:262
	MoonMeanEcc = 0.054900489
	// SunEarthMrat as defined in sweph.h:264
	SunEarthMrat = 332946.050895
	// EarthMoonMrat as defined in sweph.h:265
	EarthMoonMrat = 81.30055985272827
	// Aunit as defined in sweph.h:273
	Aunit = 1.495978707e+11
	// Clight as defined in sweph.h:274
	Clight = 2.99792458e+08
	// Helgravconst as defined in sweph.h:278
	Helgravconst = 1.32712440017987e+20
	// Geogconst as defined in sweph.h:279
	Geogconst = 3.98600448e+14
	// Kgauss as defined in sweph.h:280
	Kgauss = 0.01720209895
	// SunRadius as defined in sweph.h:281
	SunRadius = 0.00465241752803144
	// EarthRadius as defined in sweph.h:282
	EarthRadius = 6.3781366e+06
	// EarthOblateness as defined in sweph.h:284
	EarthOblateness = 0.003352819697896193
	// EarthRotSpeed as defined in sweph.h:285
	EarthRotSpeed = 6.300387486748799
	// LighttimeAunit as defined in sweph.h:287
	LighttimeAunit = 0.005775518331437499
	// ParsecToAunit as defined in sweph.h:288
	ParsecToAunit = 206264.8062471
	// SsyPlaneNodeE2000 as defined in sweph.h:291
	SsyPlaneNodeE2000 = 1.8776700468039835
	// SsyPlaneNode as defined in sweph.h:293
	SsyPlaneNode = 1.8777793895872257
	// SsyPlaneIncl as defined in sweph.h:295
	SsyPlaneIncl = 0.027553530354526998
	// KmSToAuCty as defined in sweph.h:297
	KmSToAuCty = 21.095
	// MoonSpeedIntv as defined in sweph.h:298
	MoonSpeedIntv = 5e-05
	// PlanSpeedIntv as defined in sweph.h:299
	PlanSpeedIntv = 0.0001
	// MeanNodeSpeedIntv as defined in sweph.h:300
	MeanNodeSpeedIntv = 0.001
	// NodeCalcIntv as defined in sweph.h:301
	NodeCalcIntv = 0.0001
	// NodeCalcIntvMosh as defined in sweph.h:302
	NodeCalcIntvMosh = 0.1
	// NutSpeedIntv as defined in sweph.h:303
	NutSpeedIntv = 0.0001
	// DeflSpeedIntv as defined in sweph.h:304
	DeflSpeedIntv = 5e-07
	// SeLapseRate as defined in sweph.h:306
	SeLapseRate = 0.0065
	// Ndiam as defined in sweph.h:314
	Ndiam = 21
	// Str as defined in sweph.h:666
	Str = 4.84813681109536e-06
	// SwiStarLength as defined in sweph.h:775
	SwiStarLength = 40
	// SweDataDpsiDeps as defined in sweph.h:785
	SweDataDpsiDeps = 36525
)
