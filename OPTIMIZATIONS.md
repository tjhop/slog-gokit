# slog-gokit Optimization Tracker

## Benchmark Notes

All bechmarks are run via the following format:

```shell
# before changes
go test -bench=. -benchmem -count=10 -run='^$' ./... | tee before.benchmark

# after changes
go test -bench=. -benchmem -count=10 -run='^$' ./... | tee after.benchmark
```

Benchmarks are compared with `benchstat`.

## Current Performance (04-04-2026)

- **CPU:** 11th Gen Intel Core i7-11850H @ 2.50GHz
- **OS:** Linux 6.19.8 (Arch)
- **Go:** go1.26.0

<details>

<summary>Current Benchmark Results</summary>

```shell
~/go/src/github.com/tjhop/slog-gokit (perf/more-optimizations [  ]) -> go test -bench=. -benchmem -count=10 -run='^$' ./...
goos: linux
goarch: amd64
pkg: github.com/tjhop/slog-gokit
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkBasicLog/NoAttrs-16              	  636289	      1754 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  665864	      1752 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  637110	      1730 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  633460	      1812 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  654802	      1734 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  662215	      1745 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  615259	      1726 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  632348	      1728 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  626706	      1741 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/NoAttrs-16              	  656272	      1744 ns/op	     568 B/op	       9 allocs/op
BenchmarkBasicLog/2Attrs-16               	  505902	      2094 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  516008	      2074 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  536577	      2073 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  520293	      2066 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  539901	      2068 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  537237	      2087 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  541114	      2159 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  525652	      2072 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  547346	      2080 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/2Attrs-16               	  530852	      2073 ns/op	     809 B/op	      13 allocs/op
BenchmarkBasicLog/5Attrs-16               	  426004	      2538 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  434184	      2555 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  418194	      2567 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  429315	      2546 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  438709	      2545 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  438493	      2549 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  422252	      2551 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  441096	      2559 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  443709	      2549 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/5Attrs-16               	  429924	      2658 ns/op	    1185 B/op	      19 allocs/op
BenchmarkBasicLog/10Attrs-16              	  329006	      3320 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  334419	      3316 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  329641	      3325 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  325256	      3306 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  342499	      3315 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  329143	      3310 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  325910	      3317 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  339332	      3315 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  330642	      3311 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/10Attrs-16              	  342246	      3304 ns/op	    2011 B/op	      30 allocs/op
BenchmarkBasicLog/20Attrs-16              	  225256	      4966 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  235262	      4825 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  245655	      4854 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  236233	      4773 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  233300	      4817 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  239341	      4803 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  231261	      4808 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  220112	      4831 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  227156	      4844 ns/op	    3870 B/op	      50 allocs/op
BenchmarkBasicLog/20Attrs-16              	  227308	      4810 ns/op	    3870 B/op	      50 allocs/op
BenchmarkLogLevels/Debug-16               	  485929	      2264 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  508857	      2280 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  458024	      2265 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  484138	      2269 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  514101	      2315 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  514194	      2275 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  494365	      2290 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  499692	      2294 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  472842	      2281 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Debug-16               	  503733	      2279 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  504186	      2210 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  487221	      2203 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  521533	      2194 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  508917	      2198 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  509956	      2180 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  531730	      2224 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  475010	      2198 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  478305	      2298 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  499063	      2211 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Info-16                	  514771	      2198 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  472125	      2199 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  496244	      2204 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  485616	      2201 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  515882	      2191 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  463312	      2199 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  503876	      2192 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  525422	      2207 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  491055	      2205 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  516418	      2207 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Warn-16                	  513962	      2221 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  501590	      2313 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  509017	      2187 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  511624	      2221 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  495082	      2191 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  513298	      2202 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  501135	      2194 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  549127	      2199 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  511471	      2214 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  501180	      2195 ns/op	     873 B/op	      14 allocs/op
BenchmarkLogLevels/Error-16               	  529693	      2217 ns/op	     873 B/op	      14 allocs/op
BenchmarkDisabledLogs-16                  	208878093	         5.692 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	209237424	         6.165 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	197415277	         6.022 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	199556338	         6.064 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	209454799	         5.655 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	202063815	         5.713 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	212028201	         6.031 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	214898260	         5.993 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	213917360	         5.705 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabledLogs-16                  	202156894	         5.661 ns/op	       0 B/op	       0 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  568863	      2048 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  531822	      2084 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  533874	      2068 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  560713	      2068 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  533586	      2072 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  514795	      2070 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  558682	      2067 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  546544	      2055 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  532340	      2075 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth1-16      	  546745	      2065 ns/op	     769 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  496840	      2235 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  480568	      2226 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  499434	      2303 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  483046	      2358 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  509839	      2239 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  489609	      2240 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  511219	      2207 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  497114	      2208 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  509479	      2248 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth3-16      	  510265	      2230 ns/op	     897 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  471813	      2470 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  451660	      2373 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  478618	      2369 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  439886	      2370 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  479174	      2486 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  462367	      2483 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  461584	      2414 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  428366	      2369 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  457294	      2392 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth5-16      	  447541	      2385 ns/op	    1041 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  392414	      2750 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  442842	      2741 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  393020	      2725 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  387806	      2725 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  393210	      2765 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  399430	      2896 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  387951	      2744 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  408463	      2883 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  389332	      2758 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth10-16     	  382438	      2756 ns/op	    1362 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  331488	      3375 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  324150	      3356 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  305222	      3369 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  329026	      3352 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  324162	      3389 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  322783	      3378 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  319780	      3370 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  321042	      3385 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  333080	      3383 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithAttrsChaining/Depth20-16     	  328910	      3427 ns/op	    2099 B/op	      12 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  579816	      2001 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  562706	      1936 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  542592	      1942 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  586719	      1930 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  558810	      1934 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  566269	      1940 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  582598	      1934 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  556035	      1926 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  596676	      1939 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth1-16       	  555030	      1924 ns/op	     721 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  568413	      1981 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  570718	      1970 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  552758	      1973 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  540135	      1993 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  539262	      2060 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  567922	      1986 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  578620	      1954 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  564310	      1974 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  549669	      1976 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth3-16       	  554486	      1965 ns/op	     729 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  577636	      2010 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  529119	      2004 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  561735	      2044 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  525296	      1992 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  539959	      1992 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  538401	      2006 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  538701	      2176 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  563100	      2078 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  571605	      1991 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth5-16       	  525771	      1997 ns/op	     753 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  505022	      2120 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  516330	      2085 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  541425	      2078 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  532459	      2073 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  540076	      2078 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  531536	      2055 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  545301	      2077 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  520726	      2072 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  528618	      2183 ns/op	     785 B/op	      13 allocs/op
BenchmarkWithGroupNesting/Depth10-16      	  559110	      2078 ns/op	     785 B/op	      13 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  405991	      2755 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  422163	      2639 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  427927	      2661 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  412872	      2664 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  418064	      2630 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  433738	      2627 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  420831	      2671 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  437056	      2726 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  437319	      2651 ns/op	    1129 B/op	      18 allocs/op
BenchmarkMixedWithAttrsAndGroups-16       	  416382	      2660 ns/op	    1129 B/op	      18 allocs/op
BenchmarkAttributeTypes/Strings-16        	  542287	      2008 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  597507	      2021 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  548100	      2093 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  570438	      2007 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  546085	      2043 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  549211	      2010 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  539126	      2004 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  567771	      1999 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  543910	      2032 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Strings-16        	  555088	      2009 ns/op	     944 B/op	      16 allocs/op
BenchmarkAttributeTypes/Ints-16           	  505730	      2248 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  483604	      2254 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  513278	      2266 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  469377	      2245 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  492333	      2272 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  486919	      2348 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  500774	      2254 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  493194	      2262 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  468710	      2258 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Ints-16           	  505465	      2240 ns/op	     968 B/op	      19 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  466428	      2415 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  440230	      2434 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  451138	      2444 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  469898	      2424 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  455983	      2418 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  440834	      2424 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  455941	      2428 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  455265	      2480 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  458919	      2521 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/Mixed-16          	  485618	      2420 ns/op	    1112 B/op	      24 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  297434	      3893 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  295980	      4001 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  291062	      3892 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  303938	      3891 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  299946	      4023 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  296604	      3905 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  289450	      3887 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  298599	      3911 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  296809	      4126 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/LargeStrings-16   	  288294	      3870 ns/op	     857 B/op	      16 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  437026	      2524 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  472524	      2390 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  450650	      2404 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  462609	      2404 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  460315	      2387 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  464936	      2414 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  485206	      2389 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  477424	      2429 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  462834	      2392 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/GroupAttr-16      	  464115	      2390 ns/op	    1289 B/op	      23 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  389221	      2717 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  409635	      2727 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  395338	      2713 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  403644	      2785 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  436465	      2724 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  417690	      2749 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  397772	      2713 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  405403	      2727 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  396518	      2716 ns/op	    1617 B/op	      30 allocs/op
BenchmarkAttributeTypes/NestedGroups-16   	  411830	      2731 ns/op	    1617 B/op	      30 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 2076387	       588.9 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 2016008	       606.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1938274	       615.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1935580	       652.0 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1905973	       630.6 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1858788	       636.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1888951	       643.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1873437	       654.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1851664	       659.3 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale1x-16     	 1841810	       665.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1790110	       674.9 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1783203	       722.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1737643	       693.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1702630	       703.0 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1727936	       690.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1733552	       690.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1658607	       699.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1715029	       703.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1590536	       748.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale2x-16     	 1645753	       735.3 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1710045	       706.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1684431	       710.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1694298	       706.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1690575	       711.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1680806	       714.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1672045	       715.7 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1678998	       715.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1674369	       746.9 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1668589	       724.2 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale4x-16     	 1661484	       718.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1618657	       740.4 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1631506	       733.9 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1637271	       730.7 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1633744	       729.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1644364	       733.8 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1451578	       771.1 ns/op	     915 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1499121	       854.4 ns/op	     915 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1432897	       801.4 ns/op	     915 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1587924	       777.1 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale8x-16     	 1576452	       762.5 ns/op	     915 B/op	      17 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1661754	       714.8 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1648724	       725.2 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1671494	       743.0 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1652570	       722.6 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1652077	       734.2 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1637641	       718.3 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1666488	       717.7 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1668123	       717.6 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1651491	       717.9 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentLogging/Scale16x-16    	 1593117	       772.5 ns/op	     914 B/op	      16 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1346684	       908.0 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1329376	       905.4 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1303089	       903.9 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1335662	       904.3 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1238959	       948.3 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1343706	       952.8 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1304940	       918.7 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1318575	       953.4 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1263150	       924.3 ns/op	    1195 B/op	      23 allocs/op
BenchmarkConcurrentWithAttrsThenLog-16    	 1291402	       891.6 ns/op	    1195 B/op	      23 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  621548	      1877 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  634125	      1859 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  620350	      1872 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  657882	      1849 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  635662	      1826 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  619210	      1888 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  625698	      2047 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  589621	      1981 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  642075	      1942 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted0-16      	  632882	      1842 ns/op	     585 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  538627	      2545 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  500670	      2392 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  488188	      2280 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  506798	      2273 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  427522	      2374 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  530427	      2367 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  507271	      2710 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  485955	      2560 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  491812	      2546 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted5-16      	  493773	      2501 ns/op	     905 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  345099	      3535 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  342180	      3538 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  329340	      3502 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  333759	      3527 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  335913	      3686 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  349573	      3643 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  343947	      3576 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  307962	      3580 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  309565	      3517 ns/op	    2060 B/op	      10 allocs/op
BenchmarkHandleOnly/Preformatted20-16     	  348354	      3572 ns/op	    2060 B/op	      10 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  463736	      2481 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  501129	      2533 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  491654	      2444 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  482959	      2502 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  513524	      2446 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  504313	      2504 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  496884	      2547 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  487916	      2515 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  439197	      2517 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_DebugAtDebug-16         	  464703	      2554 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  432340	      2749 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  447895	      2549 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  477796	      2507 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  477919	      2604 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  470150	      2515 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  477727	      2581 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  487240	      2560 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  473014	      2530 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  461848	      2551 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Enabled_InfoAtInfo-16           	  475530	      2521 ns/op	     737 B/op	      13 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	40920129	        29.49 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	41401628	        29.15 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	41843199	        29.71 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	35948166	        30.60 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	42116060	        29.64 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	38741269	        30.60 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	40297816	        30.15 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	33477055	        30.96 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	42057386	        30.40 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtInfo-16         	41109142	        31.59 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	40852054	        29.90 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	41373584	        30.91 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	33455965	        31.55 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	36459933	        30.30 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	41205819	        29.36 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	40920979	        29.39 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	43216662	        30.67 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	41615523	        30.63 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	38735136	        29.31 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_DebugAtError-16        	41394224	        29.17 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	41830626	        29.41 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	42589275	        29.03 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	40539139	        28.82 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	40492400	        29.06 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	38261516	        28.94 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	43152350	        28.71 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	39696042	        29.02 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	43260518	        28.97 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	36881878	        30.44 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_InfoAtWarn-16          	40045045	        29.64 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	33788697	        30.77 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	40097775	        29.98 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	41151048	        29.29 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	41450527	        29.70 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	40791807	        29.28 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	36968749	        29.04 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	34787270	        31.70 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	38695982	        30.88 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	39610891	        29.17 ns/op	      32 B/op	       1 allocs/op
BenchmarkEnabledCheck/Disabled_WarnAtError-16         	40604050	        30.57 ns/op	      32 B/op	       1 allocs/op
BenchmarkLevelSelection/Debug-16                      	  613176	      1881 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  625533	      1864 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  648075	      1943 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  548901	      1915 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  636237	      1860 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  613423	      1840 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  600132	      1850 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  642915	      1830 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  636932	      1844 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Debug-16                      	  633606	      1835 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  630736	      1837 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  632142	      1893 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  595467	      1822 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  639230	      1868 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  632005	      1856 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  609621	      1925 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  593924	      1835 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  635916	      1860 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  637219	      1863 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Info-16                       	  621210	      1835 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  633060	      1842 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  648882	      1933 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  557544	      1939 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  610010	      1917 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  681656	      1817 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  644724	      1984 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  524487	      2133 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  649251	      1924 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  628285	      1979 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Warn-16                       	  656542	      1840 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  645842	      1762 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  678906	      1816 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  577011	      1984 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  615843	      1806 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  662245	      1797 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  691346	      1860 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  633122	      1978 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  428229	      2399 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  487486	      2079 ns/op	     584 B/op	      10 allocs/op
BenchmarkLevelSelection/Error-16                      	  654103	      1860 ns/op	     584 B/op	      10 allocs/op
PASS
ok  	github.com/tjhop/slog-gokit	589.479s
```
</details>

---

## Completed Optimizations

### 1. Pairs slice capacity pre-allocation

**Status:** PR #22 (merged)

**Files:** `handler.go`

**Goal:** Reduce over-allocation and repeated slice growth in the `Handle()` hot path.

**Before:** The pairs slice in `Handle()` was allocated with a rough estimate of `2 * record.NumAttrs()`, which underestimated the actual capacity needed. This caused repeated `append()` growth and reallocation during every log call.

**After:** The capacity formula now accounts for all known components: 2 slots for the timestamp key-value pair, 2 for the message key-value pair, `len(h.preformatted)` for pre-flattened `WithAttrs` data, and `3 * record.NumAttrs()` for record attributes (2 per attr base + 50% buffer for group expansion):

```go
capacity := 4 + len(h.preformatted) + (3 * record.NumAttrs())
pairs := make([]any, 0, capacity)
```

**Impact vs v0.1.5 baseline (benchstat geomean):**
- Latency: **-47%**
- Memory: **-8%**
- Allocations: **-7%**

### 2. Pre-flatten WithAttrs to []any

**Status:** PR #22 (merged)

**Files:** `handler.go`

**Goal:** Eliminate per-call attribute processing overhead by doing the work once at `WithAttrs()` time instead of on every `Handle()` call.

**Before:** The `preformatted` field on `GoKitHandler` stored attributes as `[]slog.Attr`. On every `Handle()` call, each preformatted attribute had to be individually processed -- resolving values, expanding groups, and appending key-value pairs one at a time.

**After:** The `preformatted` field was changed to `[]any`, storing pre-flattened key-value pairs. `WithAttrs()` now resolves values and expands groups once at setup time. `Handle()` bulk-appends the entire slice in a single operation:

```go
// WithAttrs: pre-flatten at setup time
pairs := make([]any, len(h.preformatted), len(h.preformatted)+(len(attrs)*2))
copy(pairs, h.preformatted)
for _, attr := range attrs {
    pairs = appendPair(pairs, h.group, attr)
}

// Handle: bulk-append (memcopy) instead of per-attr processing
pairs = append(pairs, h.preformatted...)
```

This also made `WithAttrs` chaining allocation-constant -- the alloc count no longer scales with chain depth because each chained handler just copies and extends the flat `[]any` slice.

**Impact vs optimization 1 (benchstat geomean):**
- Latency: **-21%**
- Memory: **-7%**
- Allocations: **-12%**
- WithAttrs allocs became constant regardless of chain depth

### 3. Level logger caching

**Status:** PR #25 (open, targeting v0.1.7)

**Files:** `level.go`, `handler.go`

**Goal:** Eliminate per-call allocations from go-kit's level package by reusing leveled loggers.

**Before:** Every `Handle()` call invoked `level.Debug()`, `level.Info()`, `level.Warn()`, or `level.Error()` to create a leveled logger for the log call. Each of these functions wraps the base logger with a new context, allocating on every call.

**After:** A `levelLoggerCache` struct is created at handler construction time, pre-building all four leveled loggers. `Handle()` retrieves the appropriate cached logger via a switch on the record's level:

```go
// Construction time: pre-build all leveled loggers
type levelLoggerCache struct {
    debugLogger, infoLogger, warnLogger, errorLogger log.Logger
}

// Handle time: retrieve cached logger (no allocation)
logger := h.levelLoggers.get(record.Level)
```

**Impact:** ~2 fewer allocs/op across all benchmarks (e.g., WithAttrsChaining dropped from 14 to 12, ConcurrentLogging from ~20 to 17).

### 4. Package-level boxing of static keys

**Status:** PR #25 (open, targeting v0.1.7)

**Files:** `handler.go`

**Goal:** Eliminate per-call interface boxing of the `slog.TimeKey` and `slog.MessageKey` string constants in the `Handle()` hot path.

**Before:** Every `Handle()` call boxed `slog.TimeKey` and `slog.MessageKey` (both `string` constants) into `any` when appending them to the pairs slice. These are static values that never change, but the boxing conversion (`string` -> `any` interface header + data pointer) was paid on every log call, causing 2 unnecessary heap escapes per invocation.

**After:** The boxing cost is paid once at package init via package-level `any` variables. `Handle()` appends the pre-boxed values directly, avoiding the repeated conversion:

```go
// Package level: box once
var (
    timeKey any = slog.TimeKey
    msgKey  any = slog.MessageKey
)

// Handle: use pre-boxed keys (no per-call boxing)
pairs = append(pairs, timeKey, record.Time)
pairs = append(pairs, msgKey, record.Message)
```

**Impact vs optimization 3 (benchstat geomean):**
- Latency: **-29%** (uniform across all benchmarks, 25-40% range)
- Memory: **no change** (0 B/op delta)
- Allocations: **no change** (0 allocs/op delta)
- Variance also tightened significantly (most benchmarks dropped from 5-14% to 1-5%), suggesting reduced GC pressure jitter

---

## Cumulative Improvement vs v0.1.5 Baseline

| Benchmark | v0.1.5 ns/op | Current ns/op | Change |
|-----------|---------------:|--------------:|-------:|
| BasicLog/NoAttrs | 4,393 | 1,743 | **-60%** |
| BasicLog/20Attrs | 14,363 | 4,821 | **-66%** |
| WithAttrsChaining/Depth10 | 10,459 | 2,753 | **-74%** |
| WithAttrsChaining/Depth20 | 14,628 | 3,377 | **-77%** |
| ConcurrentLogging/Scale1x | 1,656 | 640 | **-61%** |

| Benchmark | v0.1.5 allocs | Current allocs | Change |
|-----------|----------------:|---------------:|-------:|
| BasicLog/NoAttrs | 12 | 9 | **-25%** |
| BasicLog/20Attrs | 53 | 50 | **-6%** |
| WithAttrsChaining/Depth10 | 38 | 12 | **-68%** |
| WithAttrsChaining/Depth20 | 59 | 12 | **-80%** |
| ConcurrentLogging/Scale1x | 21 | 17 | **-19%** |

---

## Attempted Optimizations

These optimizations were prototyped and benchmarked but ultimately abandoned due to regressions, insufficient improvement, or analysis showing they would be counterproductive.

### 1. sync.Pool for pairs slice in Handle()

**Files:** `handler.go` (`Handle()`)

**Goal:** Eliminate the per-call allocation of the pairs slice by reusing slices across log calls via `sync.Pool`.

**Change:** Added a package-level `sync.Pool` that stores `[]any` slices. `Handle()` attempts to retrieve a pooled slice via `Get()`, uses it for the log call, then returns it via `Put()`. Pool entries were capped at 128 capacity to prevent memory bloat from outlier large slices. Undersized slices fell back to `make()`. Slices were `clear()`-ed before returning to pool.

**Benchmark result (benchstat geomean):**
- Latency: **+38.70% regression** (p=0.000 across all benchmarks)
- Memory: **-18.46% improvement**
- Allocations: **-6.56% improvement** (-1 alloc/op across all benchmarks)

**Analysis:** The pool overhead per call -- `sync.Pool.Get()` (interface unboxing + type assertion), `clear(pairs)`, `sync.Pool.Put()` -- exceeds the cost of a single `make([]any, 0, N)` for the typical small-to-medium slice sizes in this workload. Go's runtime allocator handles short-lived, small slices very efficiently (size-class buckets, per-P caches), so the allocation being replaced was already cheap. The pool wins on B/op (avoids GC pressure from the slice header + backing array) but loses badly on wall-clock time.

**Conclusion:** Abandoned. The latency regression is too large to justify the memory savings. `sync.Pool` may become viable if the per-call pair count grows large enough that allocation cost dominates pool overhead (rough threshold: >256 pairs), but for typical logging workloads with <40 attrs this is a net loss.

### 2. Pre-join group prefix with dot separator

**Files:** `handler.go` (`GoKitHandler`, `WithGroup()`, `appendPair()`)

**Goal:** Eliminate the conditional branch in `appendPair()` by pre-computing the group prefix with its trailing dot.

**Change:** Added a `groupDot` field (`group + "."`) to `GoKitHandler`, pre-computed in `WithGroup()`. Changed `appendPair()` to unconditionally concatenate `groupDot + attr.Key` instead of the conditional `if groupPrefix != "" { key = groupPrefix + "." + key }`.

**Benchmark result (benchstat geomean):**
- Latency: **+10.55% regression** (p=0.000 across all benchmarks)
- Memory: no change
- Allocations: no change

**Analysis:** The common case is no group prefix (`group == ""`). The existing code short-circuits with `key = attr.Key` (zero-cost pointer assignment). The unconditional approach concatenates `"" + attr.Key`, which forces a new string allocation in Go's runtime even though the result is identical. The optimization only helps the less common grouped-attrs path, and the cost on the hot no-group path dominates.

**Conclusion:** Abandoned. The conditional branch is cheaper than the unconditional concatenation for the dominant no-group case. Only revisit if profiling shows group-prefix concatenation is a measurable bottleneck in a real workload. This also subsumes the earlier idea of reducing allocations in `appendPair()` for grouped record attrs -- the core concept was the same.

### 3. Replace attr.Equal(slog.Attr{}) with inline zero check

**Files:** `handler.go` (`appendPair()`)

**Goal:** Avoid constructing a temporary `slog.Attr{}` and dispatching through `Equal`/`Value.Equal` on every attribute by using a direct field comparison.

**Change:** Replaced `attr.Equal(slog.Attr{})` with `attr.Key == "" && attr.Value.Any() == nil`, skipping the temporary allocation and method dispatch.

**Benchmark results (two separate runs):**

Run 1 was inconclusive (geomean -1.01%, mixed per-benchmark). Run 2 showed:
- Latency: **+4.90% geomean regression**
- Memory: **no change** (all samples identical)
- Allocations: **no change** (all samples identical)

Notable per-benchmark results from run 2:
- BasicLog/NoAttrs: **-28.92%** (p=0.000) -- suspicious, `appendPair()` is not called in the NoAttrs path
- BasicLog/2Attrs: **-29.31%** (p=0.000) -- similarly suspicious
- ConcurrentLogging/Scale8x-16x: **+30%** (p=0.000)
- ConcurrentWithAttrsThenLog: **+31%** (p=0.000)
- HandleOnly (all preformatted counts): **+18-20%** (p=0.000)
- EnabledCheck (all paths): **+18-20%** (p=0.000)
- LevelSelection (all levels): **+7-23%** (p=0.000)

**Analysis:** The BasicLog low-attr wins are not attributable to the code change (`appendPair` is not called in the NoAttrs path), pointing to system-level noise between runs. The regressions in concurrent, HandleOnly, EnabledCheck, and LevelSelection benchmarks are consistent, large, and statistically significant across two separate runs.

**Conclusion:** Abandoned. The regressions are real and reproducible; the wins are in benchmarks the change shouldn't affect. Reverted to `attr.Equal(slog.Attr{})`.

### 4. Stack-allocated array for pairs slice in Handle()

**Files:** `handler.go` (`Handle()`)

**Goal:** Avoid heap allocation of the pairs slice for small log calls by using a fixed-size stack array, falling back to `make()` for larger calls.

**Change:** Replaced the unconditional `make()` with a fixed-size array and a capacity check:

```go
var buf [16]any // also tested with [32]any
var pairs []any
if capacity <= len(buf) {
    pairs = buf[:0:len(buf)]
} else {
    pairs = make([]any, 0, capacity)
}
```

**Benchmark result (three-way: baseline vs buf[16] vs buf[32]):**

Latency:
- buf[16] geomean: **-2.18%** (mixed -- BasicLog improved, HandleOnly/LevelSelection regressed)
- buf[32] geomean: **-0.76%** (similar split, worse overall)

Memory (B/op):
- buf[16] geomean: **+12.42% regression**
- buf[32] geomean: **+34.06% regression**
- NoAttrs: 568 -> 761 (+34%) for buf[16], 568 -> 1017 (+79%) for buf[32]
- LevelSelection: 584 -> 776 (+33%) for buf[16], 584 -> 1033 (+77%) for buf[32]

Allocations:
- Unchanged for cases that fit in the buffer
- **+1 alloc/op** for cases exceeding buffer size (both `buf` and fallback `make()` allocate)

**Analysis:** The `pairs` slice is passed to `logger.Log(pairs...)` -- an interface method call (`log.Logger`). The compiler's escape analysis can't see through the interface to prove the backing array doesn't outlive the stack frame, so it conservatively heap-allocates `buf`. The result is strictly worse than `make()`: the fixed-size array allocates more bytes than the exact-capacity `make()` (e.g., `[16]any` = 256 bytes vs `make([]any, 0, 4)` = 64 bytes for NoAttrs), and when the buffer is exceeded, both the escaped `buf` and the fallback `make()` allocate.

For comparison, the stdlib's `fmt.FormatString()` can use this pattern successfully because it uses the buffer locally and never passes it through an interface boundary.

**Conclusion:** Abandoned. The `make([]any, 0, capacity)` in `Handle()` is at the optimization floor for this allocation. It cannot be eliminated without either (a) removing the interface call to `logger.Log()`, which would mean forking go-kit, or (b) a future Go compiler improvement to escape analysis for interface method calls.

### 5. Guard attr.Value.Resolve() with Kind check

**Files:** `handler.go` (`appendPair()`)

**Goal:** Skip the `Resolve()` call for the common case where the attribute value is not a `LogValuer`.

**Change considered:** Adding `if v.Kind() == slog.KindLogValuer` before calling `Resolve()`. Abandoned before benchmarking based on analysis of the stdlib implementation:

1. `Resolve()` already short-circuits internally with `if v.Kind() != KindLogValuer { return v }` as its first operation.
2. The stdlib's own `commonHandler.appendAttr()` calls `Resolve()` unconditionally -- no Kind guard.
3. Adding a guard would mean *two* Kind checks (ours + Resolve's internal one) instead of one, which is strictly worse for the dominant non-LogValuer path.

**Conclusion:** Abandoned without testing. The current unconditional `Resolve()` call is already the correct pattern.

### 6. Make go-kit caller tracking opt-in

**Files:** `handler.go` (`setCaller()`, `NewGoKitHandler()`)

**Goal:** Reduce per-call allocation overhead from `log.Caller(6)`, which was 21% of allocation space in the original v0.1.5 profile.

**Change considered:** Making the `setCaller()` call in `NewGoKitHandler()` opt-in via a constructor option, so users who don't need caller information could skip the overhead entirely.

**Analysis:** Call-site logging is considered a required feature for the current version of this library. Making it opt-in would be a behavioral change for existing users who depend on caller information appearing in their logs. The cost is constrained by go-kit's API -- `log.Caller()` uses `runtime.Callers`/`runtime.CallersFrames` internally, and there is no way to reduce its cost without changes to go-kit itself.

**Conclusion:** Deferred. Out of scope for this optimization pass. May be revisited as a future opt-in constructor option if caller overhead proves problematic for specific workloads.

### 7. String concatenation with strings.Builder in appendPair()

**Files:** `handler.go` (`appendPair()`)

**Goal:** Reduce allocation cost of group prefix concatenation (`prefix + "." + key`) in `appendPair()` by using `strings.Builder` instead of the `+` operator.

**Change:** Prototyped replacing the `+` string concatenation with `strings.Builder` for building the dot-separated key. The builder pre-grows its buffer and writes each segment, avoiding intermediate string allocations from multi-operand concatenation.

**Analysis:** Go's compiler already optimizes small two-operand string concatenations (like `a + "." + b`) efficiently -- the runtime has a fast path for short concatenations that avoids the overhead of `strings.Builder` setup (method calls, internal buffer management). Benchmarks showed no meaningful improvement.

**Conclusion:** Abandoned. The `+` operator is already optimal for the short, fixed-format concatenations in `appendPair()`. `strings.Builder` is better suited for iterative or large-scale string construction, not simple two- or three-operand joins. Revisit only with profiling evidence that string concatenation is a measurable bottleneck.

---

## Known External Allocation Sources

The following sources of per-call allocations are in external dependencies and cannot be optimized within this library. They represent the allocation floor for any `Handle()` call.

- **`log.Caller(6)` (go-kit/log):** Runtime caller lookup via `runtime.Callers`/`runtime.CallersFrames` + string formatting of the file:line result. Was 21% of allocation space in the original v0.1.5 profile and is likely still the single largest per-call allocation source. Called in `setCaller()` at handler construction, but the deferred `log.Caller` valuer executes on every `logger.Log()` call.

- **`(*context).Log` / `WithPrefix` (go-kit/log):** Downstream allocations in go-kit's logger chain when dispatching the final `logger.Log(pairs...)` call. The context logger builds up key-value pairs and the level-filtering wrapper adds interface boxing overhead.

- **logfmt encoding (go-logfmt/logfmt):** The logfmt encoder allocates during output formatting when serializing key-value pairs to the underlying `io.Writer`. This is the final stage of the logging pipeline and is outside the handler's control.
