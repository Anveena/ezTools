package gb2312

const encode3Low, encode3High = 65072, 65510

var encode3 = [...]uint16{
	65072 - 65072: 0xA955,
	65073 - 65072: 0xA6F2,
	65075 - 65072: 0xA6F4,
	65076 - 65072: 0xA6F5,
	65077 - 65072: 0xA6E0,
	65078 - 65072: 0xA6E1,
	65079 - 65072: 0xA6F0,
	65080 - 65072: 0xA6F1,
	65081 - 65072: 0xA6E2,
	65082 - 65072: 0xA6E3,
	65083 - 65072: 0xA6EE,
	65084 - 65072: 0xA6EF,
	65085 - 65072: 0xA6E6,
	65086 - 65072: 0xA6E7,
	65087 - 65072: 0xA6E4,
	65088 - 65072: 0xA6E5,
	65089 - 65072: 0xA6E8,
	65090 - 65072: 0xA6E9,
	65091 - 65072: 0xA6EA,
	65092 - 65072: 0xA6EB,
	65097 - 65072: 0xA968,
	65098 - 65072: 0xA969,
	65099 - 65072: 0xA96A,
	65100 - 65072: 0xA96B,
	65101 - 65072: 0xA96C,
	65102 - 65072: 0xA96D,
	65103 - 65072: 0xA96E,
	65104 - 65072: 0xA96F,
	65105 - 65072: 0xA970,
	65106 - 65072: 0xA971,
	65108 - 65072: 0xA972,
	65109 - 65072: 0xA973,
	65110 - 65072: 0xA974,
	65111 - 65072: 0xA975,
	65113 - 65072: 0xA976,
	65114 - 65072: 0xA977,
	65115 - 65072: 0xA978,
	65116 - 65072: 0xA979,
	65117 - 65072: 0xA97A,
	65118 - 65072: 0xA97B,
	65119 - 65072: 0xA97C,
	65120 - 65072: 0xA97D,
	65121 - 65072: 0xA97E,
	65122 - 65072: 0xA980,
	65123 - 65072: 0xA981,
	65124 - 65072: 0xA982,
	65125 - 65072: 0xA983,
	65126 - 65072: 0xA984,
	65128 - 65072: 0xA985,
	65129 - 65072: 0xA986,
	65130 - 65072: 0xA987,
	65131 - 65072: 0xA988,
	65281 - 65072: 0xA3A1,
	65282 - 65072: 0xA3A2,
	65283 - 65072: 0xA3A3,
	65284 - 65072: 0xA1E7,
	65285 - 65072: 0xA3A5,
	65286 - 65072: 0xA3A6,
	65287 - 65072: 0xA3A7,
	65288 - 65072: 0xA3A8,
	65289 - 65072: 0xA3A9,
	65290 - 65072: 0xA3AA,
	65291 - 65072: 0xA3AB,
	65292 - 65072: 0xA3AC,
	65293 - 65072: 0xA3AD,
	65294 - 65072: 0xA3AE,
	65295 - 65072: 0xA3AF,
	65296 - 65072: 0xA3B0,
	65297 - 65072: 0xA3B1,
	65298 - 65072: 0xA3B2,
	65299 - 65072: 0xA3B3,
	65300 - 65072: 0xA3B4,
	65301 - 65072: 0xA3B5,
	65302 - 65072: 0xA3B6,
	65303 - 65072: 0xA3B7,
	65304 - 65072: 0xA3B8,
	65305 - 65072: 0xA3B9,
	65306 - 65072: 0xA3BA,
	65307 - 65072: 0xA3BB,
	65308 - 65072: 0xA3BC,
	65309 - 65072: 0xA3BD,
	65310 - 65072: 0xA3BE,
	65311 - 65072: 0xA3BF,
	65312 - 65072: 0xA3C0,
	65313 - 65072: 0xA3C1,
	65314 - 65072: 0xA3C2,
	65315 - 65072: 0xA3C3,
	65316 - 65072: 0xA3C4,
	65317 - 65072: 0xA3C5,
	65318 - 65072: 0xA3C6,
	65319 - 65072: 0xA3C7,
	65320 - 65072: 0xA3C8,
	65321 - 65072: 0xA3C9,
	65322 - 65072: 0xA3CA,
	65323 - 65072: 0xA3CB,
	65324 - 65072: 0xA3CC,
	65325 - 65072: 0xA3CD,
	65326 - 65072: 0xA3CE,
	65327 - 65072: 0xA3CF,
	65328 - 65072: 0xA3D0,
	65329 - 65072: 0xA3D1,
	65330 - 65072: 0xA3D2,
	65331 - 65072: 0xA3D3,
	65332 - 65072: 0xA3D4,
	65333 - 65072: 0xA3D5,
	65334 - 65072: 0xA3D6,
	65335 - 65072: 0xA3D7,
	65336 - 65072: 0xA3D8,
	65337 - 65072: 0xA3D9,
	65338 - 65072: 0xA3DA,
	65339 - 65072: 0xA3DB,
	65340 - 65072: 0xA3DC,
	65341 - 65072: 0xA3DD,
	65342 - 65072: 0xA3DE,
	65343 - 65072: 0xA3DF,
	65344 - 65072: 0xA3E0,
	65345 - 65072: 0xA3E1,
	65346 - 65072: 0xA3E2,
	65347 - 65072: 0xA3E3,
	65348 - 65072: 0xA3E4,
	65349 - 65072: 0xA3E5,
	65350 - 65072: 0xA3E6,
	65351 - 65072: 0xA3E7,
	65352 - 65072: 0xA3E8,
	65353 - 65072: 0xA3E9,
	65354 - 65072: 0xA3EA,
	65355 - 65072: 0xA3EB,
	65356 - 65072: 0xA3EC,
	65357 - 65072: 0xA3ED,
	65358 - 65072: 0xA3EE,
	65359 - 65072: 0xA3EF,
	65360 - 65072: 0xA3F0,
	65361 - 65072: 0xA3F1,
	65362 - 65072: 0xA3F2,
	65363 - 65072: 0xA3F3,
	65364 - 65072: 0xA3F4,
	65365 - 65072: 0xA3F5,
	65366 - 65072: 0xA3F6,
	65367 - 65072: 0xA3F7,
	65368 - 65072: 0xA3F8,
	65369 - 65072: 0xA3F9,
	65370 - 65072: 0xA3FA,
	65371 - 65072: 0xA3FB,
	65372 - 65072: 0xA3FC,
	65373 - 65072: 0xA3FD,
	65374 - 65072: 0xA1AB,
	65504 - 65072: 0xA1E9,
	65505 - 65072: 0xA1EA,
	65506 - 65072: 0xA956,
	65507 - 65072: 0xA3FE,
	65508 - 65072: 0xA957,
	65509 - 65072: 0xA3A4,
}