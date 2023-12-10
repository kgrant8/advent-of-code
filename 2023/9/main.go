package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	histories := ParseInput(puzzleInput)

	start := time.Now()
	leftCh := make(chan int)
	rightCh := make(chan int)
	var wg sync.WaitGroup
	var readWg sync.WaitGroup

	for _, history := range histories {
		wg.Add(1)

		go func(history []int) {
			defer wg.Done()
			zeroHistory := ProcessHistoryToZero(history)
			left := CalcNextNumberInHistoryLeft(zeroHistory)
			right := CalcNextNumberInHistoryRight(zeroHistory)
			leftCh <- left
			rightCh <- right
		}(history)

	}

	left, right := 0, 0
	readWg.Add(1)
	go func() {
		defer readWg.Done()
		for answer := range leftCh {
			left += answer
		}
	}()

	readWg.Add(1)
	go func() {
		defer readWg.Done()
		for answer := range rightCh {
			right += answer
		}
	}()

	wg.Wait()
	close(leftCh)
	close(rightCh)
	readWg.Wait()

	fmt.Println("left:", left)
	fmt.Println("right:", right)
	fmt.Println("totalTime:", time.Since(start))
}

func ProcessHistoryToZero(history []int) [][]int {
	matrix := [][]int{}
	matrix = append(matrix, history)
	currentNumbers := history
	zeros := 0
	for zeros != len(currentNumbers) {
		nextNumbers := []int{}
		zeros = 0
		for i := 0; i < len(currentNumbers)-1; i++ {
			diff := currentNumbers[i+1] - currentNumbers[i]
			nextNumbers = append(nextNumbers, diff)
			if diff == 0 {
				zeros++
			}
		}

		matrix = append(matrix, nextNumbers)
		currentNumbers = nextNumbers

	}

	return matrix
}

func CalcNextNumberInHistoryRight(history [][]int) int {
	nextNumber := 0
	for i := len(history) - 1; i > 0; i-- {
		lastNum := history[i][len(history[i])-1]
		aboveNumber := history[i-1][len(history[i-1])-1]
		nextNumber = lastNum + aboveNumber
		history[i-1] = append(history[i-1], nextNumber)

	}

	return nextNumber
}

func CalcNextNumberInHistoryLeft(history [][]int) int {
	nextNumber := 0
	for i := len(history) - 1; i > 0; i-- {
		firstNumber := history[i][0]
		aboveNumber := history[i-1][0]
		nextNumber = aboveNumber - firstNumber
		history[i-1] = append([]int{nextNumber}, history[i-1]...)

	}

	return nextNumber
}

func ParseInput(input string) [][]int {
	histories := [][]int{}
	for _, history := range strings.Split(input, "\n") {
		entry := []int{}
		for _, num := range strings.Split(history, " ") {
			n := mustAtoi(num)
			entry = append(entry, n)

		}

		histories = append(histories, entry)
	}
	return histories
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

const puzzleInput = `13 33 54 85 159 343 748 1539 2945 5269 8898 14313 22099 32955 47704 67303 92853 125609 166990 218589 282183
15 23 39 84 195 437 933 1921 3845 7482 14095 25579 44529 74102 117465 176514 249411 326313 382455 367494 189719
11 20 41 88 175 316 525 816 1203 1700 2321 3080 3991 5068 6325 7776 9435 11316 13433 15800 18431
5 5 20 65 172 398 835 1631 3047 5609 10470 20179 40166 81398 164845 328621 638937 1206325 2208968 3925405 6779376
12 18 16 2 -28 -78 -152 -254 -388 -558 -768 -1022 -1324 -1678 -2088 -2558 -3092 -3694 -4368 -5118 -5948
20 47 89 157 281 523 1003 1949 3780 7234 13577 25001 45492 82831 153202 291564 574368 1168005 2429517 5110697 10761191
-2 -7 -11 -14 -15 6 133 611 2041 5750 14460 33452 72542 149422 295420 565790 1058760 1951547 3568599 6509170 11880351
9 9 7 11 37 115 314 803 1988 4812 11394 26340 59318 129892 276207 569966 1141307 2218747 4191393 7703217 13792451
-1 8 32 88 216 498 1087 2246 4407 8291 15193 27642 50806 95237 181852 350434 673423 1277362 2375078 4312524 7635196
4 7 9 22 78 245 656 1565 3458 7268 14784 29441 57899 113273 221722 435561 858412 1691521 3315687 6430820 12285616
9 20 38 77 161 317 576 990 1668 2829 4865 8402 14342 23864 38357 59253 87723 124194 167640 214595 257831
10 23 47 83 134 219 407 879 2034 4673 10322 21783 44023 85535 160390 291494 516366 898579 1551681 2689175 4725760
14 19 24 35 59 98 141 161 130 77 247 1494 6180 20069 55987 140296 323349 694840 1404221 2685550 4887180
11 12 15 15 6 -17 -56 -108 -163 -202 -195 -99 144 609 1390 2602 4383 6896 10331 14907 20874
11 18 37 79 164 344 740 1611 3485 7396 15291 30713 59974 114304 214089 397641 739607 1389305 2649267 5133564 10082748
16 25 49 107 225 443 844 1618 3183 6416 13128 27092 56271 117509 246046 514172 1066830 2188239 4424751 8808701 17261601
16 29 56 114 245 528 1083 2069 3688 6233 10273 17165 30235 57190 114624 235876 484000 972229 1895070 3574068 6523337
12 22 35 42 44 67 177 495 1212 2604 5047 9032 15180 24257 37189 55077 79212 111090 152427 205174 271532
9 15 24 45 96 198 364 583 799 885 612 -387 -2658 -6972 -14376 -26249 -44363 -70949 -108768 -161187 -232260
19 25 41 85 191 420 885 1801 3578 6992 13501 25831 49058 92579 173627 323377 597253 1091827 1971753 3511561 6158913
12 17 26 57 144 345 761 1588 3239 6588 13403 27050 53565 103206 192612 347711 607534 1029107 1693608 2713991 4244294
21 32 43 54 65 76 87 98 109 120 131 142 153 164 175 186 197 208 219 230 241
19 35 61 93 133 207 400 915 2160 4863 10207 19964 36599 63352 104503 166667 263685 431751 770237 1541451 3396285
20 27 34 59 134 306 651 1319 2635 5283 10600 21012 40661 76304 138607 243992 417190 694555 1127926 1788281 2767472
15 21 27 33 39 45 51 57 63 69 75 81 87 93 99 105 111 117 123 129 135
8 24 59 127 249 456 797 1373 2432 4580 9196 19187 40274 83039 165943 319378 590444 1047412 1781573 2902163 4517027
-4 -11 -16 -6 38 141 334 654 1144 1853 2836 4154 5874 8069 10818 14206 18324 23269 29144 36058 44126
15 33 78 178 376 734 1347 2375 4101 7023 11988 20376 34342 57124 93425 149877 235595 362829 547722 811182 1179876
-3 -3 -2 9 53 190 561 1468 3523 7923 16930 34647 68183 129313 236840 420274 726694 1235914 2096607 3611954 6434594
-4 -5 -2 17 79 235 580 1309 2857 6211 13553 29524 63643 134857 279962 568890 1131822 2206035 4214652 7895427 14505811
10 5 -6 -11 11 87 251 551 1064 1921 3344 5697 9553 15779 25641 40931 64118 98525 148534 219821 319623
6 25 60 111 178 278 489 1049 2557 6346 15125 34018 72163 145073 278004 510622 903312 1545527 2566634 4149777 6549344
-4 -9 0 46 160 376 727 1251 2018 3188 5110 8477 14575 25753 46539 86711 170027 358369 815481 1967142 4885157
4 -3 -18 -34 -19 98 455 1317 3202 7147 15278 32031 66707 138608 286879 588464 1189360 2356697 4562134 8610652 15833001
7 0 -7 -17 -38 -83 -170 -322 -567 -938 -1473 -2215 -3212 -4517 -6188 -8288 -10885 -14052 -17867 -22413 -27778
22 39 67 110 172 257 369 512 690 907 1167 1474 1832 2245 2717 3252 3854 4527 5275 6102 7012
7 18 29 43 72 148 356 908 2294 5583 13011 29091 62619 130137 261656 509746 963473 1769112 3160097 5499291 9337378
11 17 21 28 64 192 541 1368 3185 6994 14690 29745 58441 112292 213061 403190 766859 1470715 2839119 5487230 10547202
11 22 50 98 169 270 423 701 1329 2919 6955 16738 39192 88294 191540 402024 820933 1639884 3222641 6267959 12142675
15 22 27 43 108 307 820 2007 4537 9564 18949 35523 63382 108201 177550 281191 431331 642802 933135 1322491 1833408
18 26 44 86 177 360 703 1311 2364 4244 7897 15715 33433 73832 163436 353904 740462 1490510 2886490 5388228 9721281
13 21 35 68 147 326 724 1617 3642 8223 18418 40542 87199 182855 373994 747597 1464900 2823458 5371766 10120781 18931336
2 4 15 57 161 376 803 1680 3558 7619 16197 33581 67216 129474 240210 430263 745769 1252480 2038261 3211192 4892352
4 -4 -10 3 62 202 464 893 1536 2440 3650 5207 7146 9494 12268 15473 19100 23124 27502 32171 37046
-3 -5 -6 -11 -32 -84 -163 -192 86 1299 4765 12999 30487 64828 128370 240494 430731 742931 1240740 2014681 3191178
-5 9 45 110 213 367 591 912 1367 2005 2889 4098 5729 7899 10747 14436 19155 25121 32581 41814 53133
17 26 40 77 174 388 796 1505 2697 4761 8611 16369 32724 67494 140289 288866 584145 1157632 2252458 4319553 8199058
-3 5 23 66 174 427 960 1978 3776 6782 11672 19670 33274 57908 105524 202219 403959 828414 1718491 3568007 7369250
12 26 57 117 228 447 898 1812 3589 6929 13147 24907 47801 93482 185458 369192 728858 1414004 2680501 4951545 8906162
29 52 92 166 300 532 929 1643 3045 5990 12280 25406 51664 101754 192985 352223 619733 1054080 1738268 2787310 4357436
22 32 47 73 128 253 519 1029 1921 3405 5929 10681 20811 44014 97467 216575 469568 980718 1965827 3784689 7016466
9 31 60 104 178 301 503 852 1526 2988 6381 14369 32873 74628 166495 364532 784971 1667368 3501792 7281227 14994052
11 23 45 77 119 171 233 305 387 479 581 693 815 947 1089 1241 1403 1575 1757 1949 2151
24 34 40 44 50 68 128 307 771 1830 3993 7988 14675 24724 37851 51299 57114 37594 -41922 -244012 -675264
17 37 67 107 157 217 287 367 457 557 667 787 917 1057 1207 1367 1537 1717 1907 2107 2317
-4 -3 -7 -15 -11 57 305 958 2443 5588 12018 24878 50058 98146 187392 348029 628366 1103143 1884719 3137751 5098115
20 41 68 94 109 100 51 -57 -246 -541 -970 -1564 -2357 -3386 -4691 -6315 -8304 -10707 -13576 -16966 -20935
19 40 75 132 232 417 756 1358 2420 4377 8291 16740 35698 78331 172445 374826 798517 1664370 3398255 6814236 13459032
9 17 37 69 108 151 223 431 1063 2777 6997 16783 38711 86732 189622 404532 840337 1695997 3321001 6307173 11623666
5 10 27 70 164 355 720 1373 2458 4120 6465 9586 13879 21143 37405 79135 185776 444011 1034664 2325491 5060375
19 31 42 67 138 312 692 1465 2957 5711 10618 19178 34040 60066 106280 189192 338119 603247 1067274 1861525 3187414
20 21 13 3 7 59 246 794 2241 5744 13578 29896 61830 121024 225701 403377 694346 1156071 1868627 2941353 4520881
23 42 77 146 279 518 917 1542 2471 3794 5613 8042 11207 15246 20309 26558 34167 43322 54221 67074 82103
12 9 6 16 64 187 434 866 1556 2589 4062 6084 8776 12271 16714 22262 29084 37361 47286 59064 72912
1 -5 -9 9 95 347 949 2208 4589 8743 15523 25983 41355 62999 92321 130654 179097 238307 308239 387829 474615
6 8 5 0 6 48 162 403 886 1896 4115 9026 19566 41112 82896 159957 295750 525544 900753 1494356 2407574
5 9 29 86 215 486 1052 2248 4776 10021 20552 40870 78472 145306 259697 448828 751863 1223801 1940151 3002518 4545189
9 28 62 116 208 381 715 1339 2443 4290 7228 11702 18266 27595 40497 57925 80989 110968 149322 197704 257972
12 32 68 131 238 422 756 1397 2659 5144 10006 19516 38290 75953 152851 312025 643498 1331588 2743084 5584245 11167890
11 25 43 77 149 291 545 963 1607 2549 3871 5665 8033 11087 14949 19751 25635 32753 41267 51349 63181
28 55 103 195 382 758 1488 2855 5342 9789 17713 31980 58234 107932 204694 397252 785040 1567139 3133019 6226081 12231134
7 16 23 32 57 122 271 609 1414 3394 8218 19540 44890 99080 210252 430505 854353 1650316 3114025 5755696 10441143
13 21 32 56 114 260 622 1464 3271 6857 13482 24932 43469 71518 110990 162399 224741 299077 402972 611148 1152656
20 33 46 55 51 32 36 209 935 3076 8399 20304 45012 93425 183931 346496 628462 1102555 1877700 3113341 5038073
22 45 85 148 249 433 807 1578 3084 5791 10202 16572 24230 30156 26221 -5885 -103674 -339723 -849785 -1883600 -3898717
8 22 61 143 288 528 944 1749 3450 7157 15173 32100 66829 135939 269203 518088 968360 1758216 3103865 5335341 8946818
0 15 48 102 186 340 678 1468 3289 7335 15978 33759 69051 136734 262342 488289 882958 1553647 2664612 4461732 7305648
12 23 42 76 132 217 338 502 716 987 1322 1728 2212 2781 3442 4202 5068 6047 7146 8372 9732
-4 0 19 72 194 446 934 1845 3523 6634 12506 23778 45551 87304 165919 310251 567782 1014012 1765365 2996524 4963256
2 2 1 -4 -16 -38 -73 -124 -194 -286 -403 -548 -724 -934 -1181 -1468 -1798 -2174 -2599 -3076 -3608
7 15 35 91 219 474 954 1863 3661 7389 15314 32135 67180 138403 279708 554392 1079583 2069791 3912497 7295551 13415573
11 4 -1 -1 4 13 38 154 636 2282 7105 19729 50099 118592 265394 567197 1165977 2316924 4465547 8367535 15268974
25 44 71 105 150 228 409 870 2008 4661 10544 23111 49268 102799 211244 429721 868710 1749937 3518772 7069640 14194697
1 10 38 97 195 330 487 644 802 1082 1991 5061 14222 38494 96886 226784 497607 1032122 2038548 3857457 7028509
11 22 51 111 221 405 690 1113 1752 2801 4733 8686 17462 38123 88402 211539 508614 1205614 2787254 6256065 13622072
13 19 36 82 190 419 875 1750 3394 6451 12123 22695 42599 80595 154264 299293 588758 1173438 2366521 4821491 9902667
17 15 12 18 45 113 266 598 1287 2630 5064 9147 15459 24366 35570 47345 55333 50745 17780 -69960 -254319
18 35 60 101 170 283 464 760 1290 2389 4978 11406 27183 64272 146982 322143 676486 1366699 2672792 5095419 9536268
9 17 27 52 128 335 822 1833 3745 7173 13301 24811 48185 98879 212127 464297 1013470 2174530 4555932 9306779 18561383
16 31 50 72 100 145 236 447 957 2172 4957 11044 23687 48610 95217 177877 316830 537845 869156 1333360 1930830
-3 -10 -5 25 86 182 343 699 1645 4182 10597 25799 59919 133307 285982 595220 1207936 2400193 4687434 9031704 17237379
17 43 86 161 295 525 905 1535 2641 4761 9137 18499 38589 81059 168848 345873 693945 1361343 2609560 4887501 8946003
17 19 18 22 58 183 495 1141 2319 4271 7264 11556 17344 24691 33429 43035 52477 60027 63038 57682 38646
13 19 45 106 217 393 649 1000 1461 2047 2773 3654 4705 5941 7377 9028 10909 13035 15421 18082 21033
-4 -2 9 44 132 325 716 1487 3017 6098 12356 25097 51083 104358 214483 443879 923146 1920282 3971147 8116302 16316134
7 19 41 74 121 206 417 982 2381 5486 11706 23103 42470 73504 121637 197143 324463 566481 1081899 2251640 4942603
5 17 40 87 177 336 597 1005 1640 2692 4687 9118 20042 47746 116490 279845 649775 1450482 3117479 6481019 13109637
6 28 77 178 380 770 1487 2738 4833 8280 14014 23876 41521 74076 135261 251725 476832 921508 1822479 3688210 7598090
12 30 59 98 146 202 265 334 408 486 567 650 734 818 901 982 1060 1134 1203 1266 1322
10 18 19 17 33 115 360 965 2332 5263 11295 23242 46031 87942 162388 290400 504014 850792 1399747 2248983 3535405
10 6 2 -2 -6 -10 -14 -18 -22 -26 -30 -34 -38 -42 -46 -50 -54 -58 -62 -66 -70
13 26 43 75 138 253 446 748 1195 1828 2693 3841 5328 7215 9568 12458 15961 20158 25135 30983 37798
0 18 51 99 162 240 333 441 564 702 855 1023 1206 1404 1617 1845 2088 2346 2619 2907 3210
12 28 44 66 121 270 620 1333 2626 4748 7920 12268 17948 26126 42576 88042 224331 614307 1650643 4217081 10193946
7 6 8 27 85 212 446 833 1427 2290 3492 5111 7233 9952 13370 17597 22751 28958 36352 45075 55277
23 41 62 86 113 143 176 212 251 293 338 386 437 491 548 608 671 737 806 878 953
28 42 53 62 78 135 335 944 2592 6664 16022 36292 78140 161349 322277 627769 1201522 2274764 4285045 8071228 15261336
15 32 66 140 282 520 890 1467 2424 4119 7205 12753 22373 38313 63511 101570 156621 233034 334932 465458 625740
11 27 64 141 287 540 945 1561 2499 4038 6927 13111 27370 60832 138227 312551 694502 1511750 3229225 6795114 14145736
7 14 18 19 17 12 4 -7 -21 -38 -58 -81 -107 -136 -168 -203 -241 -282 -326 -373 -423
4 19 55 126 247 434 704 1075 1566 2197 2989 3964 5145 6556 8222 10169 12424 15015 17971 21322 25099
17 23 35 67 145 309 615 1137 1969 3227 5051 7607 11089 15721 21759 29493 39249 51391 66323 84491 106385
9 25 45 66 88 123 225 555 1508 3955 9701 22363 49111 104253 216771 446098 914379 1870213 3809861 7699056 15364394
26 43 63 81 90 81 43 -37 -174 -385 -689 -1107 -1662 -2379 -3285 -4409 -5782 -7437 -9409 -11735 -14454
19 33 64 119 214 396 769 1515 2892 5187 8612 13173 18652 25083 34574 56172 116885 284225 709052 1702495 3867806
5 -2 -17 -42 -61 -25 163 663 1713 3644 6895 12028 19743 30893 46499 67765 96093 133098 180623 240754 315835
21 40 74 140 267 498 901 1600 2849 5201 9879 19565 40053 83696 176548 372934 784437 1635778 3368896 6830198 13596383
-1 3 9 33 118 361 967 2352 5324 11376 23128 44956 83845 150500 260744 437225 711445 1126113 1737811 2619947 3865952
9 28 66 147 306 595 1114 2084 3985 7796 15413 30419 59592 115947 224845 435968 846073 1640900 3170188 6078603 11529190
13 20 34 48 55 52 50 95 312 1014 2986 8188 21379 53662 129906 303805 687665 1510038 3225056 6716190 13669969
15 39 75 126 201 313 478 725 1133 1913 3555 7082 14550 30223 63549 136524 300866 675795 1527604 3433068 7611838
11 33 82 172 323 573 994 1705 2883 4805 8023 13893 25861 52167 110975 239386 508355 1046225 2074424 3959858 7289687
17 45 88 152 253 436 808 1592 3213 6442 12664 24430 46658 89268 172867 340733 683573 1389903 2847324 5843600 11959975
-5 -8 -16 -21 1 96 336 821 1680 3072 5187 8245 12488 18156 25433 34343 44569 55160 64082 67559 59139
21 38 73 131 211 307 412 527 693 1103 2428 6632 18814 51120 130743 315906 727357 1610037 3453721 7229388 14852433
11 13 8 -6 -30 -64 -107 -157 -211 -265 -314 -352 -372 -366 -325 -239 -97 113 404 790 1286
6 3 6 21 53 111 221 458 1030 2493 6266 15784 38943 93091 214968 480176 1040782 2196831 4529874 9147043 18116921
3 4 9 32 106 296 708 1504 2961 5658 10954 22055 46184 98687 210352 439841 894131 1761832 3368775 6275950 11464186
-7 5 40 103 203 368 670 1260 2420 4658 8918 17084 33186 66177 136055 286883 613796 1318154 2819045 5976728 12535578
19 33 61 102 150 205 299 544 1211 2848 6450 13725 27588 53202 100224 187469 352051 666283 1267315 2406768 4530602
6 18 38 65 110 210 442 937 1894 3594 6414 10841 17486 27098 40578 58993 83590 115810 157302 209937 275822
22 33 54 89 147 246 420 741 1386 2805 6076 13563 30029 64442 132984 264547 511902 975873 1858082 3573029 6976719
8 10 28 75 164 308 520 813 1200 1694 2308 3055 3948 5000 6224 7633 9240 11058 13100 15379 17908
10 14 26 44 79 173 420 990 2156 4324 8066 14156 23609 37723 58124 86814 126222 179258 249370 340604 457667
11 15 10 0 -4 27 164 561 1536 3749 8586 18936 40662 85222 174094 345932 667882 1253676 2294063 4111953 7269149
8 19 42 106 258 567 1129 2085 3674 6351 11005 19317 34322 61336 109694 196427 354428 650343 1222132 2354032 4618938
6 11 34 101 259 588 1213 2314 4139 7045 11627 19051 31808 55326 101403 195663 393980 818471 1734572 3710573 7935361
8 3 -4 -10 4 100 428 1307 3379 7908 17361 36528 74657 149475 294674 573706 1104978 2107481 3982678 7459902 13851250
24 47 76 104 136 210 424 969 2168 4521 8756 15886 27272 44692 70416 107287 158808 229235 323676 448196 609928
8 16 24 32 40 48 56 64 72 80 88 96 104 112 120 128 136 144 152 160 168
10 25 52 102 210 444 918 1820 3479 6528 12297 23726 47365 97483 204050 427646 886830 1808657 3624068 7153244 13973898
-8 0 28 99 255 571 1176 2289 4302 7981 14921 28504 55804 111207 223028 445187 877144 1696898 3214054 5952901 10779289
-4 -4 -3 4 26 83 228 586 1416 3225 7033 15045 32290 70337 155162 342892 749919 1608432 3365926 6860022 13622901
10 20 25 28 47 130 371 931 2077 4261 8270 15487 28312 50801 89590 155180 263668 439018 715975 1143734 1790485
3 12 35 85 183 359 645 1061 1600 2223 2880 3578 4522 6360 10568 20016 39761 78118 148065 269043 469217
16 17 27 58 122 231 397 632 948 1357 1871 2502 3262 4163 5217 6436 7832 9417 11203 13202 15426
4 10 22 34 36 22 13 101 529 1847 5235 13177 30823 68600 146918 304113 608962 1178988 2205012 3980504 6930526
11 15 21 29 49 104 231 483 931 1660 2764 4412 7265 14039 34139 95590 272947 748181 1936587 4736255 11005081
4 17 46 108 241 512 1025 1929 3426 5779 9320 14458 21687 31594 44867 62303 84816 113445 149362 193880 248461
5 6 18 47 104 210 403 760 1459 2929 6190 13595 30382 67774 148930 320070 671017 1373044 2748702 5400527 10446700
18 25 43 81 151 271 468 781 1264 1989 3049 4561 6669 9547 13402 18477 25054 33457 44055 57265 73555
9 6 4 16 71 231 632 1565 3610 7835 16078 31354 58505 105421 185665 324410 571667 1030464 1913764 3653576 7100287
24 46 77 117 166 224 291 367 452 546 649 761 882 1012 1151 1299 1456 1622 1797 1981 2174
9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
20 37 60 89 124 165 212 265 324 389 460 537 620 709 804 905 1012 1125 1244 1369 1500
17 20 20 14 -2 -33 -85 -165 -281 -442 -658 -940 -1300 -1751 -2307 -2983 -3795 -4760 -5896 -7222 -8758
19 39 71 122 199 311 472 713 1131 2037 4329 10347 25732 63316 150969 346833 767759 1640383 3389567 6786420 13188433
-5 7 37 93 196 391 762 1470 2855 5680 11664 24574 52345 110994 231534 471829 936845 1814393 3442605 6447835 12044316
25 43 64 93 142 226 355 523 692 763 517 -503 -3213 -9374 -22226 -47477 -94781 -179869 -327530 -575675 -980756
9 6 3 0 -3 -6 -9 -12 -15 -18 -21 -24 -27 -30 -33 -36 -39 -42 -45 -48 -51
6 8 22 65 172 400 830 1577 2823 4902 8515 15278 29064 59064 126250 276076 601912 1285988 2667648 5352591 10381608
-8 -13 -6 39 155 376 730 1231 1870 2605 3350 3963 4233 3866 2470 -461 -5564 -13625 -25598 -42625 -66057
15 25 56 114 201 329 555 1051 2232 4975 10975 23315 47414 92741 176205 329251 612941 1149597 2187542 4230004 8285735
15 25 36 63 137 311 682 1442 2972 5992 11788 22588 42324 78422 146110 278331 547133 1106987 2277668 4695181 9575075
7 16 40 87 167 299 523 925 1690 3202 6222 12228 24160 48183 97809 202983 428713 910612 1918228 3960785 7952789
7 31 70 129 216 346 553 909 1561 2830 5489 11484 25642 59425 138710 319176 715579 1555609 3276016 6689451 13266572
23 30 34 29 5 -52 -160 -341 -621 -1030 -1602 -2375 -3391 -4696 -6340 -8377 -10865 -13866 -17446 -21675 -26627
10 12 20 34 47 54 86 288 1067 3353 9063 21960 49284 104832 214619 426917 829424 1577678 2940774 5373210 9625621
23 34 40 51 88 183 379 730 1301 2168 3418 5149 7470 10501 14373 19228 25219 32510 41276 51703 63988
-4 5 35 102 242 536 1150 2396 4825 9368 17546 31775 55797 95273 158579 257851 410330 640063 980021 1474700 2183276
13 17 21 25 29 33 37 41 45 49 53 57 61 65 69 73 77 81 85 89 93
15 30 44 52 60 112 330 967 2473 5574 11364 21410 37870 63624 102418 159021 239395 350878 502380 704592 970208
2 4 17 67 193 447 894 1612 2692 4238 6367 9209 12907 17617 23508 30762 39574 50152 62717 77503 94757
11 21 35 63 130 275 553 1050 1929 3541 6664 12980 25970 52505 105542 208503 402127 754845 1378041 2447933 4236242
-1 0 2 6 27 110 357 963 2251 4686 8847 15363 24908 38570 59407 97089 179830 382567 889831 2130051 5052100
18 31 56 117 268 615 1344 2757 5321 9738 17047 28772 47133 75340 117993 181614 275340 411809 608274 887983 1281866
19 44 81 128 180 228 269 350 685 1903 5507 14649 35354 78357 161751 314681 582359 1032718 1765069 2921174 4699200
4 6 8 18 55 153 369 795 1574 2920 5142 8672 14097 22195 33975 50721 74040 105914 148756 205470 279515
14 25 27 25 30 53 102 195 411 1019 2755 7372 18721 44970 103416 231195 508860 1110505 2405627 5156671 10885448
-10 -8 -4 -3 -10 -22 -2 178 840 2679 7091 16811 37184 78694 161970 327692 656241 1306674 2593628 5137666 10156119
8 6 15 52 142 325 673 1321 2515 4686 8582 15540 28067 51033 93970 175229 329082 617278 1147081 2098444 3763716
12 14 21 39 70 112 162 227 348 641 1357 2956 6177 12083 22126 38577 66629 122165 256050 622262 1654477
-1 4 21 60 143 327 742 1649 3532 7254 14332 27424 51187 93820 170033 307318 560143 1042777 2004115 3993115 8212469
17 35 77 160 302 525 865 1389 2219 3563 5753 9290 14896 23573 36669 55951 83685 122723 176597 249620 346994
18 35 56 89 145 244 430 795 1512 2877 5360 9665 16799 28150 45574 71491 108990 161943 235128 334361 466637
19 40 66 96 129 164 200 236 271 304 334 360 381 396 404 404 395 376 346 304 249
3 12 26 45 84 196 513 1324 3226 7408 16159 33729 67717 131212 245972 446992 788885 1354580 2266928 3703901 5918170
3 4 10 26 63 162 441 1176 2924 6702 14269 28651 55259 104376 196609 374439 725829 1430944 2851992 5701504 11349787
24 36 52 80 140 271 538 1039 1912 3342 5568 8890 13676 20369 29494 41665 57592 78088 104076 136596 176812
13 31 66 144 317 676 1378 2711 5240 10099 19523 37769 72697 138541 260902 485889 896819 1644219 2998372 5438706 9801409
10 28 63 119 200 323 539 971 1893 3894 8200 17269 35826 72551 142652 271578 500424 895139 1567235 2729968 4855364
20 46 85 147 250 420 704 1200 2103 3761 6730 11812 20055 32689 50967 75875 107670 145200 184955 219793 237280
3 9 28 77 187 418 884 1788 3467 6447 11508 19759 32723 52432 81532 123398 182259 263333 372972 518817 709963
20 43 82 150 269 466 764 1170 1664 2195 2692 3100 3453 3998 5386 8948 17076 33731 65102 120442 213109
3 13 36 68 96 89 -6 -252 -677 -1160 -1164 879 8957 32510 92789 235713 557128 1251434 2703952 5664910 11575649
13 22 33 44 46 33 32 155 667 2055 5087 10891 21201 39190 71926 136877 277988 603547 1371956 3179169 7354295
8 22 47 105 246 559 1187 2358 4459 8219 15139 28425 54875 108515 217420 436368 870226 1717008 3346523 6442129 12251716
-10 -6 20 76 164 281 417 557 706 972 1769 4251 11180 28622 69304 159499 353646 765874 1637434 3477370 7346116
4 19 40 58 65 64 91 269 939 2946 8203 20733 48541 106968 224737 454871 894250 1718041 3238910 6006208 10967693
6 1 -7 -18 -32 -49 -69 -92 -118 -147 -179 -214 -252 -293 -337 -384 -434 -487 -543 -602 -664`
