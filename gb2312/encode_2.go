package gb2312

const encode2Low, encode2High = 164, 1106

var encode2 = [...]uint16{
	164 - 164:  0xA1E8,
	167 - 164:  0xA1EC,
	168 - 164:  0xA1A7,
	176 - 164:  0xA1E3,
	177 - 164:  0xA1C0,
	183 - 164:  0xA1A4,
	215 - 164:  0xA1C1,
	224 - 164:  0xA8A4,
	225 - 164:  0xA8A2,
	232 - 164:  0xA8A8,
	233 - 164:  0xA8A6,
	234 - 164:  0xA8BA,
	236 - 164:  0xA8AC,
	237 - 164:  0xA8AA,
	242 - 164:  0xA8B0,
	243 - 164:  0xA8AE,
	247 - 164:  0xA1C2,
	249 - 164:  0xA8B4,
	250 - 164:  0xA8B2,
	252 - 164:  0xA8B9,
	257 - 164:  0xA8A1,
	275 - 164:  0xA8A5,
	283 - 164:  0xA8A7,
	299 - 164:  0xA8A9,
	324 - 164:  0xA8BD,
	328 - 164:  0xA8BE,
	333 - 164:  0xA8AD,
	363 - 164:  0xA8B1,
	462 - 164:  0xA8A3,
	464 - 164:  0xA8AB,
	466 - 164:  0xA8AF,
	468 - 164:  0xA8B3,
	470 - 164:  0xA8B5,
	472 - 164:  0xA8B6,
	474 - 164:  0xA8B7,
	476 - 164:  0xA8B8,
	505 - 164:  0xA8BF,
	593 - 164:  0xA8BB,
	609 - 164:  0xA8C0,
	711 - 164:  0xA1A6,
	713 - 164:  0xA1A5,
	714 - 164:  0xA840,
	715 - 164:  0xA841,
	729 - 164:  0xA842,
	913 - 164:  0xA6A1,
	914 - 164:  0xA6A2,
	915 - 164:  0xA6A3,
	916 - 164:  0xA6A4,
	917 - 164:  0xA6A5,
	918 - 164:  0xA6A6,
	919 - 164:  0xA6A7,
	920 - 164:  0xA6A8,
	921 - 164:  0xA6A9,
	922 - 164:  0xA6AA,
	923 - 164:  0xA6AB,
	924 - 164:  0xA6AC,
	925 - 164:  0xA6AD,
	926 - 164:  0xA6AE,
	927 - 164:  0xA6AF,
	928 - 164:  0xA6B0,
	929 - 164:  0xA6B1,
	931 - 164:  0xA6B2,
	932 - 164:  0xA6B3,
	933 - 164:  0xA6B4,
	934 - 164:  0xA6B5,
	935 - 164:  0xA6B6,
	936 - 164:  0xA6B7,
	937 - 164:  0xA6B8,
	945 - 164:  0xA6C1,
	946 - 164:  0xA6C2,
	947 - 164:  0xA6C3,
	948 - 164:  0xA6C4,
	949 - 164:  0xA6C5,
	950 - 164:  0xA6C6,
	951 - 164:  0xA6C7,
	952 - 164:  0xA6C8,
	953 - 164:  0xA6C9,
	954 - 164:  0xA6CA,
	955 - 164:  0xA6CB,
	956 - 164:  0xA6CC,
	957 - 164:  0xA6CD,
	958 - 164:  0xA6CE,
	959 - 164:  0xA6CF,
	960 - 164:  0xA6D0,
	961 - 164:  0xA6D1,
	963 - 164:  0xA6D2,
	964 - 164:  0xA6D3,
	965 - 164:  0xA6D4,
	966 - 164:  0xA6D5,
	967 - 164:  0xA6D6,
	968 - 164:  0xA6D7,
	969 - 164:  0xA6D8,
	1025 - 164: 0xA7A7,
	1040 - 164: 0xA7A1,
	1041 - 164: 0xA7A2,
	1042 - 164: 0xA7A3,
	1043 - 164: 0xA7A4,
	1044 - 164: 0xA7A5,
	1045 - 164: 0xA7A6,
	1046 - 164: 0xA7A8,
	1047 - 164: 0xA7A9,
	1048 - 164: 0xA7AA,
	1049 - 164: 0xA7AB,
	1050 - 164: 0xA7AC,
	1051 - 164: 0xA7AD,
	1052 - 164: 0xA7AE,
	1053 - 164: 0xA7AF,
	1054 - 164: 0xA7B0,
	1055 - 164: 0xA7B1,
	1056 - 164: 0xA7B2,
	1057 - 164: 0xA7B3,
	1058 - 164: 0xA7B4,
	1059 - 164: 0xA7B5,
	1060 - 164: 0xA7B6,
	1061 - 164: 0xA7B7,
	1062 - 164: 0xA7B8,
	1063 - 164: 0xA7B9,
	1064 - 164: 0xA7BA,
	1065 - 164: 0xA7BB,
	1066 - 164: 0xA7BC,
	1067 - 164: 0xA7BD,
	1068 - 164: 0xA7BE,
	1069 - 164: 0xA7BF,
	1070 - 164: 0xA7C0,
	1071 - 164: 0xA7C1,
	1072 - 164: 0xA7D1,
	1073 - 164: 0xA7D2,
	1074 - 164: 0xA7D3,
	1075 - 164: 0xA7D4,
	1076 - 164: 0xA7D5,
	1077 - 164: 0xA7D6,
	1078 - 164: 0xA7D8,
	1079 - 164: 0xA7D9,
	1080 - 164: 0xA7DA,
	1081 - 164: 0xA7DB,
	1082 - 164: 0xA7DC,
	1083 - 164: 0xA7DD,
	1084 - 164: 0xA7DE,
	1085 - 164: 0xA7DF,
	1086 - 164: 0xA7E0,
	1087 - 164: 0xA7E1,
	1088 - 164: 0xA7E2,
	1089 - 164: 0xA7E3,
	1090 - 164: 0xA7E4,
	1091 - 164: 0xA7E5,
	1092 - 164: 0xA7E6,
	1093 - 164: 0xA7E7,
	1094 - 164: 0xA7E8,
	1095 - 164: 0xA7E9,
	1096 - 164: 0xA7EA,
	1097 - 164: 0xA7EB,
	1098 - 164: 0xA7EC,
	1099 - 164: 0xA7ED,
	1100 - 164: 0xA7EE,
	1101 - 164: 0xA7EF,
	1102 - 164: 0xA7F0,
	1103 - 164: 0xA7F1,
	1105 - 164: 0xA7D7,
}
