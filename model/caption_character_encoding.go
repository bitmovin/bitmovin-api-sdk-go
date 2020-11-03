package model


// CaptionCharacterEncoding : Character encoding of the SRT file
type CaptionCharacterEncoding string

// List of possible CaptionCharacterEncoding values
const (
    CaptionCharacterEncoding_ANSI_X3_4_1968 CaptionCharacterEncoding = "ANSI_X3.4-1968"
    CaptionCharacterEncoding_ANSI_X3_4_1986 CaptionCharacterEncoding = "ANSI_X3.4-1986"
    CaptionCharacterEncoding_ARABIC CaptionCharacterEncoding = "ARABIC"
    CaptionCharacterEncoding_ARMSCII_8 CaptionCharacterEncoding = "ARMSCII-8"
    CaptionCharacterEncoding_ASCII CaptionCharacterEncoding = "ASCII"
    CaptionCharacterEncoding_ASMO_708 CaptionCharacterEncoding = "ASMO-708"
    CaptionCharacterEncoding_BIG_5 CaptionCharacterEncoding = "BIG-5"
    CaptionCharacterEncoding_BIG_FIVE CaptionCharacterEncoding = "BIG-FIVE"
    CaptionCharacterEncoding_BIG5 CaptionCharacterEncoding = "BIG5"
    CaptionCharacterEncoding_BIG5_HKSCS CaptionCharacterEncoding = "BIG5-HKSCS"
    CaptionCharacterEncoding_BIG5_HKSCS_1999 CaptionCharacterEncoding = "BIG5-HKSCS:1999"
    CaptionCharacterEncoding_BIG5_HKSCS_2001 CaptionCharacterEncoding = "BIG5-HKSCS:2001"
    CaptionCharacterEncoding_BIG5_HKSCS_2004 CaptionCharacterEncoding = "BIG5-HKSCS:2004"
    CaptionCharacterEncoding_BIG5_HKSCS_2008 CaptionCharacterEncoding = "BIG5-HKSCS:2008"
    CaptionCharacterEncoding_BIG5HKSCS CaptionCharacterEncoding = "BIG5HKSCS"
    CaptionCharacterEncoding_BIGFIVE CaptionCharacterEncoding = "BIGFIVE"
    CaptionCharacterEncoding_C99 CaptionCharacterEncoding = "C99"
    CaptionCharacterEncoding_CHAR CaptionCharacterEncoding = "CHAR"
    CaptionCharacterEncoding_CHINESE CaptionCharacterEncoding = "CHINESE"
    CaptionCharacterEncoding_CN CaptionCharacterEncoding = "CN"
    CaptionCharacterEncoding_CN_BIG5 CaptionCharacterEncoding = "CN-BIG5"
    CaptionCharacterEncoding_CN_GB CaptionCharacterEncoding = "CN-GB"
    CaptionCharacterEncoding_CN_GB_ISOIR165 CaptionCharacterEncoding = "CN-GB-ISOIR165"
    CaptionCharacterEncoding_CP1131 CaptionCharacterEncoding = "CP1131"
    CaptionCharacterEncoding_CP1133 CaptionCharacterEncoding = "CP1133"
    CaptionCharacterEncoding_CP1250 CaptionCharacterEncoding = "CP1250"
    CaptionCharacterEncoding_CP1251 CaptionCharacterEncoding = "CP1251"
    CaptionCharacterEncoding_CP1252 CaptionCharacterEncoding = "CP1252"
    CaptionCharacterEncoding_CP1253 CaptionCharacterEncoding = "CP1253"
    CaptionCharacterEncoding_CP1254 CaptionCharacterEncoding = "CP1254"
    CaptionCharacterEncoding_CP1255 CaptionCharacterEncoding = "CP1255"
    CaptionCharacterEncoding_CP1256 CaptionCharacterEncoding = "CP1256"
    CaptionCharacterEncoding_CP1257 CaptionCharacterEncoding = "CP1257"
    CaptionCharacterEncoding_CP1258 CaptionCharacterEncoding = "CP1258"
    CaptionCharacterEncoding_CP1361 CaptionCharacterEncoding = "CP1361"
    CaptionCharacterEncoding_CP154 CaptionCharacterEncoding = "CP154"
    CaptionCharacterEncoding_CP367 CaptionCharacterEncoding = "CP367"
    CaptionCharacterEncoding_CP819 CaptionCharacterEncoding = "CP819"
    CaptionCharacterEncoding_CP850 CaptionCharacterEncoding = "CP850"
    CaptionCharacterEncoding_CP862 CaptionCharacterEncoding = "CP862"
    CaptionCharacterEncoding_CP866 CaptionCharacterEncoding = "CP866"
    CaptionCharacterEncoding_CP874 CaptionCharacterEncoding = "CP874"
    CaptionCharacterEncoding_CP932 CaptionCharacterEncoding = "CP932"
    CaptionCharacterEncoding_CP936 CaptionCharacterEncoding = "CP936"
    CaptionCharacterEncoding_CP949 CaptionCharacterEncoding = "CP949"
    CaptionCharacterEncoding_CP950 CaptionCharacterEncoding = "CP950"
    CaptionCharacterEncoding_CSASCII CaptionCharacterEncoding = "CSASCII"
    CaptionCharacterEncoding_CSBIG5 CaptionCharacterEncoding = "CSBIG5"
    CaptionCharacterEncoding_CSEUCKR CaptionCharacterEncoding = "CSEUCKR"
    CaptionCharacterEncoding_CSEUCPKDFMTJAPANESE CaptionCharacterEncoding = "CSEUCPKDFMTJAPANESE"
    CaptionCharacterEncoding_CSEUCTW CaptionCharacterEncoding = "CSEUCTW"
    CaptionCharacterEncoding_CSGB2312 CaptionCharacterEncoding = "CSGB2312"
    CaptionCharacterEncoding_CSHALFWIDTHKATAKANA CaptionCharacterEncoding = "CSHALFWIDTHKATAKANA"
    CaptionCharacterEncoding_CSHPROMAN8 CaptionCharacterEncoding = "CSHPROMAN8"
    CaptionCharacterEncoding_CSIBM866 CaptionCharacterEncoding = "CSIBM866"
    CaptionCharacterEncoding_CSISO14JISC6220RO CaptionCharacterEncoding = "CSISO14JISC6220RO"
    CaptionCharacterEncoding_CSISO159JISX02121990 CaptionCharacterEncoding = "CSISO159JISX02121990"
    CaptionCharacterEncoding_CSISO2022CN CaptionCharacterEncoding = "CSISO2022CN"
    CaptionCharacterEncoding_CSISO2022JP CaptionCharacterEncoding = "CSISO2022JP"
    CaptionCharacterEncoding_CSISO2022JP2 CaptionCharacterEncoding = "CSISO2022JP2"
    CaptionCharacterEncoding_CSISO2022KR CaptionCharacterEncoding = "CSISO2022KR"
    CaptionCharacterEncoding_CSISO57GB1988 CaptionCharacterEncoding = "CSISO57GB1988"
    CaptionCharacterEncoding_CSISO58GB231280 CaptionCharacterEncoding = "CSISO58GB231280"
    CaptionCharacterEncoding_CSISO87JISX0208 CaptionCharacterEncoding = "CSISO87JISX0208"
    CaptionCharacterEncoding_CSISOLATIN1 CaptionCharacterEncoding = "CSISOLATIN1"
    CaptionCharacterEncoding_CSISOLATIN2 CaptionCharacterEncoding = "CSISOLATIN2"
    CaptionCharacterEncoding_CSISOLATIN3 CaptionCharacterEncoding = "CSISOLATIN3"
    CaptionCharacterEncoding_CSISOLATIN4 CaptionCharacterEncoding = "CSISOLATIN4"
    CaptionCharacterEncoding_CSISOLATIN5 CaptionCharacterEncoding = "CSISOLATIN5"
    CaptionCharacterEncoding_CSISOLATIN6 CaptionCharacterEncoding = "CSISOLATIN6"
    CaptionCharacterEncoding_CSISOLATINARABIC CaptionCharacterEncoding = "CSISOLATINARABIC"
    CaptionCharacterEncoding_CSISOLATINCYRILLIC CaptionCharacterEncoding = "CSISOLATINCYRILLIC"
    CaptionCharacterEncoding_CSISOLATINGREEK CaptionCharacterEncoding = "CSISOLATINGREEK"
    CaptionCharacterEncoding_CSISOLATINHEBREW CaptionCharacterEncoding = "CSISOLATINHEBREW"
    CaptionCharacterEncoding_CSKOI8R CaptionCharacterEncoding = "CSKOI8R"
    CaptionCharacterEncoding_CSKSC56011987 CaptionCharacterEncoding = "CSKSC56011987"
    CaptionCharacterEncoding_CSKZ1048 CaptionCharacterEncoding = "CSKZ1048"
    CaptionCharacterEncoding_CSMACINTOSH CaptionCharacterEncoding = "CSMACINTOSH"
    CaptionCharacterEncoding_CSPC850MULTILINGUAL CaptionCharacterEncoding = "CSPC850MULTILINGUAL"
    CaptionCharacterEncoding_CSPC862LATINHEBREW CaptionCharacterEncoding = "CSPC862LATINHEBREW"
    CaptionCharacterEncoding_CSPTCP154 CaptionCharacterEncoding = "CSPTCP154"
    CaptionCharacterEncoding_CSSHIFTJIS CaptionCharacterEncoding = "CSSHIFTJIS"
    CaptionCharacterEncoding_CSUCS4 CaptionCharacterEncoding = "CSUCS4"
    CaptionCharacterEncoding_CSUNICODE CaptionCharacterEncoding = "CSUNICODE"
    CaptionCharacterEncoding_CSUNICODE11 CaptionCharacterEncoding = "CSUNICODE11"
    CaptionCharacterEncoding_CSUNICODE11UTF7 CaptionCharacterEncoding = "CSUNICODE11UTF7"
    CaptionCharacterEncoding_CSVISCII CaptionCharacterEncoding = "CSVISCII"
    CaptionCharacterEncoding_CYRILLIC CaptionCharacterEncoding = "CYRILLIC"
    CaptionCharacterEncoding_CYRILLIC_ASIAN CaptionCharacterEncoding = "CYRILLIC-ASIAN"
    CaptionCharacterEncoding_ECMA_114 CaptionCharacterEncoding = "ECMA-114"
    CaptionCharacterEncoding_ECMA_118 CaptionCharacterEncoding = "ECMA-118"
    CaptionCharacterEncoding_ELOT_928 CaptionCharacterEncoding = "ELOT_928"
    CaptionCharacterEncoding_EUC_CN CaptionCharacterEncoding = "EUC-CN"
    CaptionCharacterEncoding_EUC_JP CaptionCharacterEncoding = "EUC-JP"
    CaptionCharacterEncoding_EUC_KR CaptionCharacterEncoding = "EUC-KR"
    CaptionCharacterEncoding_EUC_TW CaptionCharacterEncoding = "EUC-TW"
    CaptionCharacterEncoding_EUCCN CaptionCharacterEncoding = "EUCCN"
    CaptionCharacterEncoding_EUCJP CaptionCharacterEncoding = "EUCJP"
    CaptionCharacterEncoding_EUCKR CaptionCharacterEncoding = "EUCKR"
    CaptionCharacterEncoding_EUCTW CaptionCharacterEncoding = "EUCTW"
    CaptionCharacterEncoding_EXTENDED_UNIX_CODE_PACKED_FORMAT_FOR_JAPANESE CaptionCharacterEncoding = "EXTENDED_UNIX_CODE_PACKED_FORMAT_FOR_JAPANESE"
    CaptionCharacterEncoding_GB_1988_80 CaptionCharacterEncoding = "GB_1988-80"
    CaptionCharacterEncoding_GB_2312_80 CaptionCharacterEncoding = "GB_2312-80"
    CaptionCharacterEncoding_GB18030 CaptionCharacterEncoding = "GB18030"
    CaptionCharacterEncoding_GB2312 CaptionCharacterEncoding = "GB2312"
    CaptionCharacterEncoding_GBK CaptionCharacterEncoding = "GBK"
    CaptionCharacterEncoding_GEORGIAN_ACADEMY CaptionCharacterEncoding = "GEORGIAN-ACADEMY"
    CaptionCharacterEncoding_GEORGIAN_PS CaptionCharacterEncoding = "GEORGIAN-PS"
    CaptionCharacterEncoding_GREEK CaptionCharacterEncoding = "GREEK"
    CaptionCharacterEncoding_GREEK8 CaptionCharacterEncoding = "GREEK8"
    CaptionCharacterEncoding_HEBREW CaptionCharacterEncoding = "HEBREW"
    CaptionCharacterEncoding_HP_ROMAN8 CaptionCharacterEncoding = "HP-ROMAN8"
    CaptionCharacterEncoding_HZ CaptionCharacterEncoding = "HZ"
    CaptionCharacterEncoding_HZ_GB_2312 CaptionCharacterEncoding = "HZ-GB-2312"
    CaptionCharacterEncoding_IBM_CP1133 CaptionCharacterEncoding = "IBM-CP1133"
    CaptionCharacterEncoding_IBM367 CaptionCharacterEncoding = "IBM367"
    CaptionCharacterEncoding_IBM819 CaptionCharacterEncoding = "IBM819"
    CaptionCharacterEncoding_IBM850 CaptionCharacterEncoding = "IBM850"
    CaptionCharacterEncoding_IBM862 CaptionCharacterEncoding = "IBM862"
    CaptionCharacterEncoding_IBM866 CaptionCharacterEncoding = "IBM866"
    CaptionCharacterEncoding_ISO_10646_UCS_2 CaptionCharacterEncoding = "ISO-10646-UCS-2"
    CaptionCharacterEncoding_ISO_10646_UCS_4 CaptionCharacterEncoding = "ISO-10646-UCS-4"
    CaptionCharacterEncoding_ISO_2022_CN CaptionCharacterEncoding = "ISO-2022-CN"
    CaptionCharacterEncoding_ISO_2022_CN_EXT CaptionCharacterEncoding = "ISO-2022-CN-EXT"
    CaptionCharacterEncoding_ISO_2022_JP CaptionCharacterEncoding = "ISO-2022-JP"
    CaptionCharacterEncoding_ISO_2022_JP_1 CaptionCharacterEncoding = "ISO-2022-JP-1"
    CaptionCharacterEncoding_ISO_2022_JP_2 CaptionCharacterEncoding = "ISO-2022-JP-2"
    CaptionCharacterEncoding_ISO_2022_KR CaptionCharacterEncoding = "ISO-2022-KR"
    CaptionCharacterEncoding_ISO_8859_1 CaptionCharacterEncoding = "ISO-8859-1"
    CaptionCharacterEncoding_ISO_8859_10 CaptionCharacterEncoding = "ISO-8859-10"
    CaptionCharacterEncoding_ISO_8859_11 CaptionCharacterEncoding = "ISO-8859-11"
    CaptionCharacterEncoding_ISO_8859_13 CaptionCharacterEncoding = "ISO-8859-13"
    CaptionCharacterEncoding_ISO_8859_14 CaptionCharacterEncoding = "ISO-8859-14"
    CaptionCharacterEncoding_ISO_8859_15 CaptionCharacterEncoding = "ISO-8859-15"
    CaptionCharacterEncoding_ISO_8859_16 CaptionCharacterEncoding = "ISO-8859-16"
    CaptionCharacterEncoding_ISO_8859_2 CaptionCharacterEncoding = "ISO-8859-2"
    CaptionCharacterEncoding_ISO_8859_3 CaptionCharacterEncoding = "ISO-8859-3"
    CaptionCharacterEncoding_ISO_8859_4 CaptionCharacterEncoding = "ISO-8859-4"
    CaptionCharacterEncoding_ISO_8859_5 CaptionCharacterEncoding = "ISO-8859-5"
    CaptionCharacterEncoding_ISO_8859_6 CaptionCharacterEncoding = "ISO-8859-6"
    CaptionCharacterEncoding_ISO_8859_7 CaptionCharacterEncoding = "ISO-8859-7"
    CaptionCharacterEncoding_ISO_8859_8 CaptionCharacterEncoding = "ISO-8859-8"
    CaptionCharacterEncoding_ISO_8859_9 CaptionCharacterEncoding = "ISO-8859-9"
    CaptionCharacterEncoding_ISO_CELTIC CaptionCharacterEncoding = "ISO-CELTIC"
    CaptionCharacterEncoding_ISO_IR_100 CaptionCharacterEncoding = "ISO-IR-100"
    CaptionCharacterEncoding_ISO_IR_101 CaptionCharacterEncoding = "ISO-IR-101"
    CaptionCharacterEncoding_ISO_IR_109 CaptionCharacterEncoding = "ISO-IR-109"
    CaptionCharacterEncoding_ISO_IR_110 CaptionCharacterEncoding = "ISO-IR-110"
    CaptionCharacterEncoding_ISO_IR_126 CaptionCharacterEncoding = "ISO-IR-126"
    CaptionCharacterEncoding_ISO_IR_127 CaptionCharacterEncoding = "ISO-IR-127"
    CaptionCharacterEncoding_ISO_IR_138 CaptionCharacterEncoding = "ISO-IR-138"
    CaptionCharacterEncoding_ISO_IR_14 CaptionCharacterEncoding = "ISO-IR-14"
    CaptionCharacterEncoding_ISO_IR_144 CaptionCharacterEncoding = "ISO-IR-144"
    CaptionCharacterEncoding_ISO_IR_148 CaptionCharacterEncoding = "ISO-IR-148"
    CaptionCharacterEncoding_ISO_IR_149 CaptionCharacterEncoding = "ISO-IR-149"
    CaptionCharacterEncoding_ISO_IR_157 CaptionCharacterEncoding = "ISO-IR-157"
    CaptionCharacterEncoding_ISO_IR_159 CaptionCharacterEncoding = "ISO-IR-159"
    CaptionCharacterEncoding_ISO_IR_165 CaptionCharacterEncoding = "ISO-IR-165"
    CaptionCharacterEncoding_ISO_IR_166 CaptionCharacterEncoding = "ISO-IR-166"
    CaptionCharacterEncoding_ISO_IR_179 CaptionCharacterEncoding = "ISO-IR-179"
    CaptionCharacterEncoding_ISO_IR_199 CaptionCharacterEncoding = "ISO-IR-199"
    CaptionCharacterEncoding_ISO_IR_203 CaptionCharacterEncoding = "ISO-IR-203"
    CaptionCharacterEncoding_ISO_IR_226 CaptionCharacterEncoding = "ISO-IR-226"
    CaptionCharacterEncoding_ISO_IR_57 CaptionCharacterEncoding = "ISO-IR-57"
    CaptionCharacterEncoding_ISO_IR_58 CaptionCharacterEncoding = "ISO-IR-58"
    CaptionCharacterEncoding_ISO_IR_6 CaptionCharacterEncoding = "ISO-IR-6"
    CaptionCharacterEncoding_ISO_IR_87 CaptionCharacterEncoding = "ISO-IR-87"
    CaptionCharacterEncoding_ISO_646_IRV_1991 CaptionCharacterEncoding = "ISO_646.IRV:1991"
    CaptionCharacterEncoding_ISO_8859_1_1987 CaptionCharacterEncoding = "ISO_8859-1:1987"
    CaptionCharacterEncoding_ISO_8859_10_1992 CaptionCharacterEncoding = "ISO_8859-10:1992"
    CaptionCharacterEncoding_ISO_8859_14_1998 CaptionCharacterEncoding = "ISO_8859-14:1998"
    CaptionCharacterEncoding_ISO_8859_15_1998 CaptionCharacterEncoding = "ISO_8859-15:1998"
    CaptionCharacterEncoding_ISO_8859_16_2001 CaptionCharacterEncoding = "ISO_8859-16:2001"
    CaptionCharacterEncoding_ISO_8859_2_1987 CaptionCharacterEncoding = "ISO_8859-2:1987"
    CaptionCharacterEncoding_ISO_8859_3_1988 CaptionCharacterEncoding = "ISO_8859-3:1988"
    CaptionCharacterEncoding_ISO_8859_4_1988 CaptionCharacterEncoding = "ISO_8859-4:1988"
    CaptionCharacterEncoding_ISO_8859_5_1988 CaptionCharacterEncoding = "ISO_8859-5:1988"
    CaptionCharacterEncoding_ISO_8859_6_1987 CaptionCharacterEncoding = "ISO_8859-6:1987"
    CaptionCharacterEncoding_ISO_8859_7_1987 CaptionCharacterEncoding = "ISO_8859-7:1987"
    CaptionCharacterEncoding_ISO_8859_7_2003 CaptionCharacterEncoding = "ISO_8859-7:2003"
    CaptionCharacterEncoding_ISO_8859_8_1988 CaptionCharacterEncoding = "ISO_8859-8:1988"
    CaptionCharacterEncoding_ISO_8859_9_1989 CaptionCharacterEncoding = "ISO_8859-9:1989"
    CaptionCharacterEncoding_ISO646_CN CaptionCharacterEncoding = "ISO646-CN"
    CaptionCharacterEncoding_ISO646_JP CaptionCharacterEncoding = "ISO646-JP"
    CaptionCharacterEncoding_ISO646_US CaptionCharacterEncoding = "ISO646-US"
    CaptionCharacterEncoding_ISO8859_1 CaptionCharacterEncoding = "ISO8859-1"
    CaptionCharacterEncoding_ISO8859_10 CaptionCharacterEncoding = "ISO8859-10"
    CaptionCharacterEncoding_ISO8859_11 CaptionCharacterEncoding = "ISO8859-11"
    CaptionCharacterEncoding_ISO8859_13 CaptionCharacterEncoding = "ISO8859-13"
    CaptionCharacterEncoding_ISO8859_14 CaptionCharacterEncoding = "ISO8859-14"
    CaptionCharacterEncoding_ISO8859_15 CaptionCharacterEncoding = "ISO8859-15"
    CaptionCharacterEncoding_ISO8859_16 CaptionCharacterEncoding = "ISO8859-16"
    CaptionCharacterEncoding_ISO8859_2 CaptionCharacterEncoding = "ISO8859-2"
    CaptionCharacterEncoding_ISO8859_3 CaptionCharacterEncoding = "ISO8859-3"
    CaptionCharacterEncoding_ISO8859_4 CaptionCharacterEncoding = "ISO8859-4"
    CaptionCharacterEncoding_ISO8859_5 CaptionCharacterEncoding = "ISO8859-5"
    CaptionCharacterEncoding_ISO8859_6 CaptionCharacterEncoding = "ISO8859-6"
    CaptionCharacterEncoding_ISO8859_7 CaptionCharacterEncoding = "ISO8859-7"
    CaptionCharacterEncoding_ISO8859_8 CaptionCharacterEncoding = "ISO8859-8"
    CaptionCharacterEncoding_ISO8859_9 CaptionCharacterEncoding = "ISO8859-9"
    CaptionCharacterEncoding_JAVA CaptionCharacterEncoding = "JAVA"
    CaptionCharacterEncoding_JIS_C6220_1969_RO CaptionCharacterEncoding = "JIS_C6220-1969-RO"
    CaptionCharacterEncoding_JIS_C6226_1983 CaptionCharacterEncoding = "JIS_C6226-1983"
    CaptionCharacterEncoding_JIS_X0201 CaptionCharacterEncoding = "JIS_X0201"
    CaptionCharacterEncoding_JIS_X0208 CaptionCharacterEncoding = "JIS_X0208"
    CaptionCharacterEncoding_JIS_X0208_1983 CaptionCharacterEncoding = "JIS_X0208-1983"
    CaptionCharacterEncoding_JIS_X0208_1990 CaptionCharacterEncoding = "JIS_X0208-1990"
    CaptionCharacterEncoding_JIS_X0212 CaptionCharacterEncoding = "JIS_X0212"
    CaptionCharacterEncoding_JIS_X0212_1990 CaptionCharacterEncoding = "JIS_X0212-1990"
    CaptionCharacterEncoding_JIS_X0212_1990_0 CaptionCharacterEncoding = "JIS_X0212.1990-0"
    CaptionCharacterEncoding_JIS0208 CaptionCharacterEncoding = "JIS0208"
    CaptionCharacterEncoding_JISX0201_1976 CaptionCharacterEncoding = "JISX0201-1976"
    CaptionCharacterEncoding_JOHAB CaptionCharacterEncoding = "JOHAB"
    CaptionCharacterEncoding_JP CaptionCharacterEncoding = "JP"
    CaptionCharacterEncoding_KOI8_R CaptionCharacterEncoding = "KOI8-R"
    CaptionCharacterEncoding_KOI8_RU CaptionCharacterEncoding = "KOI8-RU"
    CaptionCharacterEncoding_KOI8_T CaptionCharacterEncoding = "KOI8-T"
    CaptionCharacterEncoding_KOI8_U CaptionCharacterEncoding = "KOI8-U"
    CaptionCharacterEncoding_KOREAN CaptionCharacterEncoding = "KOREAN"
    CaptionCharacterEncoding_KS_C_5601_1987 CaptionCharacterEncoding = "KS_C_5601-1987"
    CaptionCharacterEncoding_KS_C_5601_1989 CaptionCharacterEncoding = "KS_C_5601-1989"
    CaptionCharacterEncoding_KSC_5601 CaptionCharacterEncoding = "KSC_5601"
    CaptionCharacterEncoding_KZ_1048 CaptionCharacterEncoding = "KZ-1048"
    CaptionCharacterEncoding_L1 CaptionCharacterEncoding = "L1"
    CaptionCharacterEncoding_L10 CaptionCharacterEncoding = "L10"
    CaptionCharacterEncoding_L2 CaptionCharacterEncoding = "L2"
    CaptionCharacterEncoding_L3 CaptionCharacterEncoding = "L3"
    CaptionCharacterEncoding_L4 CaptionCharacterEncoding = "L4"
    CaptionCharacterEncoding_L5 CaptionCharacterEncoding = "L5"
    CaptionCharacterEncoding_L6 CaptionCharacterEncoding = "L6"
    CaptionCharacterEncoding_L7 CaptionCharacterEncoding = "L7"
    CaptionCharacterEncoding_L8 CaptionCharacterEncoding = "L8"
    CaptionCharacterEncoding_LATIN_9 CaptionCharacterEncoding = "LATIN-9"
    CaptionCharacterEncoding_LATIN1 CaptionCharacterEncoding = "LATIN1"
    CaptionCharacterEncoding_LATIN10 CaptionCharacterEncoding = "LATIN10"
    CaptionCharacterEncoding_LATIN2 CaptionCharacterEncoding = "LATIN2"
    CaptionCharacterEncoding_LATIN3 CaptionCharacterEncoding = "LATIN3"
    CaptionCharacterEncoding_LATIN4 CaptionCharacterEncoding = "LATIN4"
    CaptionCharacterEncoding_LATIN5 CaptionCharacterEncoding = "LATIN5"
    CaptionCharacterEncoding_LATIN6 CaptionCharacterEncoding = "LATIN6"
    CaptionCharacterEncoding_LATIN7 CaptionCharacterEncoding = "LATIN7"
    CaptionCharacterEncoding_LATIN8 CaptionCharacterEncoding = "LATIN8"
    CaptionCharacterEncoding_MAC CaptionCharacterEncoding = "MAC"
    CaptionCharacterEncoding_MACARABIC CaptionCharacterEncoding = "MACARABIC"
    CaptionCharacterEncoding_MACCENTRALEUROPE CaptionCharacterEncoding = "MACCENTRALEUROPE"
    CaptionCharacterEncoding_MACCROATIAN CaptionCharacterEncoding = "MACCROATIAN"
    CaptionCharacterEncoding_MACCYRILLIC CaptionCharacterEncoding = "MACCYRILLIC"
    CaptionCharacterEncoding_MACGREEK CaptionCharacterEncoding = "MACGREEK"
    CaptionCharacterEncoding_MACHEBREW CaptionCharacterEncoding = "MACHEBREW"
    CaptionCharacterEncoding_MACICELAND CaptionCharacterEncoding = "MACICELAND"
    CaptionCharacterEncoding_MACINTOSH CaptionCharacterEncoding = "MACINTOSH"
    CaptionCharacterEncoding_MACROMAN CaptionCharacterEncoding = "MACROMAN"
    CaptionCharacterEncoding_MACROMANIA CaptionCharacterEncoding = "MACROMANIA"
    CaptionCharacterEncoding_MACTHAI CaptionCharacterEncoding = "MACTHAI"
    CaptionCharacterEncoding_MACTURKISH CaptionCharacterEncoding = "MACTURKISH"
    CaptionCharacterEncoding_MACUKRAINE CaptionCharacterEncoding = "MACUKRAINE"
    CaptionCharacterEncoding_MS_ANSI CaptionCharacterEncoding = "MS-ANSI"
    CaptionCharacterEncoding_MS_ARAB CaptionCharacterEncoding = "MS-ARAB"
    CaptionCharacterEncoding_MS_CYRL CaptionCharacterEncoding = "MS-CYRL"
    CaptionCharacterEncoding_MS_EE CaptionCharacterEncoding = "MS-EE"
    CaptionCharacterEncoding_MS_GREEK CaptionCharacterEncoding = "MS-GREEK"
    CaptionCharacterEncoding_MS_HEBR CaptionCharacterEncoding = "MS-HEBR"
    CaptionCharacterEncoding_MS_TURK CaptionCharacterEncoding = "MS-TURK"
    CaptionCharacterEncoding_MS_KANJI CaptionCharacterEncoding = "MS_KANJI"
    CaptionCharacterEncoding_MS936 CaptionCharacterEncoding = "MS936"
    CaptionCharacterEncoding_MULELAO_1 CaptionCharacterEncoding = "MULELAO-1"
    CaptionCharacterEncoding_NEXTSTEP CaptionCharacterEncoding = "NEXTSTEP"
    CaptionCharacterEncoding_PT154 CaptionCharacterEncoding = "PT154"
    CaptionCharacterEncoding_PTCP154 CaptionCharacterEncoding = "PTCP154"
    CaptionCharacterEncoding_R8 CaptionCharacterEncoding = "R8"
    CaptionCharacterEncoding_RK1048 CaptionCharacterEncoding = "RK1048"
    CaptionCharacterEncoding_ROMAN8 CaptionCharacterEncoding = "ROMAN8"
    CaptionCharacterEncoding_SHIFT_JIS CaptionCharacterEncoding = "SHIFT-JIS"
    CaptionCharacterEncoding_SJIS CaptionCharacterEncoding = "SJIS"
    CaptionCharacterEncoding_STRK1048_2002 CaptionCharacterEncoding = "STRK1048-2002"
    CaptionCharacterEncoding_TCVN CaptionCharacterEncoding = "TCVN"
    CaptionCharacterEncoding_TCVN_5712 CaptionCharacterEncoding = "TCVN-5712"
    CaptionCharacterEncoding_TCVN5712_1 CaptionCharacterEncoding = "TCVN5712-1"
    CaptionCharacterEncoding_TCVN5712_1_1993 CaptionCharacterEncoding = "TCVN5712-1:1993"
    CaptionCharacterEncoding_TIS_620 CaptionCharacterEncoding = "TIS-620"
    CaptionCharacterEncoding_TIS620 CaptionCharacterEncoding = "TIS620"
    CaptionCharacterEncoding_TIS620_0 CaptionCharacterEncoding = "TIS620-0"
    CaptionCharacterEncoding_TIS620_2529_1 CaptionCharacterEncoding = "TIS620.2529-1"
    CaptionCharacterEncoding_TIS620_2533_0 CaptionCharacterEncoding = "TIS620.2533-0"
    CaptionCharacterEncoding_TIS620_2533_1 CaptionCharacterEncoding = "TIS620.2533-1"
    CaptionCharacterEncoding_UCS_2 CaptionCharacterEncoding = "UCS-2"
    CaptionCharacterEncoding_UCS_2_INTERNAL CaptionCharacterEncoding = "UCS-2-INTERNAL"
    CaptionCharacterEncoding_UCS_2_SWAPPED CaptionCharacterEncoding = "UCS-2-SWAPPED"
    CaptionCharacterEncoding_UCS_2BE CaptionCharacterEncoding = "UCS-2BE"
    CaptionCharacterEncoding_UCS_2LE CaptionCharacterEncoding = "UCS-2LE"
    CaptionCharacterEncoding_UCS_4 CaptionCharacterEncoding = "UCS-4"
    CaptionCharacterEncoding_UCS_4_INTERNAL CaptionCharacterEncoding = "UCS-4-INTERNAL"
    CaptionCharacterEncoding_UCS_4_SWAPPED CaptionCharacterEncoding = "UCS-4-SWAPPED"
    CaptionCharacterEncoding_UCS_4BE CaptionCharacterEncoding = "UCS-4BE"
    CaptionCharacterEncoding_UCS_4LE CaptionCharacterEncoding = "UCS-4LE"
    CaptionCharacterEncoding_UHC CaptionCharacterEncoding = "UHC"
    CaptionCharacterEncoding_UNICODE_1_1 CaptionCharacterEncoding = "UNICODE-1-1"
    CaptionCharacterEncoding_UNICODE_1_1_UTF_7 CaptionCharacterEncoding = "UNICODE-1-1-UTF-7"
    CaptionCharacterEncoding_UNICODEBIG CaptionCharacterEncoding = "UNICODEBIG"
    CaptionCharacterEncoding_UNICODELITTLE CaptionCharacterEncoding = "UNICODELITTLE"
    CaptionCharacterEncoding_US CaptionCharacterEncoding = "US"
    CaptionCharacterEncoding_US_ASCII CaptionCharacterEncoding = "US-ASCII"
    CaptionCharacterEncoding_UTF_16 CaptionCharacterEncoding = "UTF-16"
    CaptionCharacterEncoding_UTF_16BE CaptionCharacterEncoding = "UTF-16BE"
    CaptionCharacterEncoding_UTF_16LE CaptionCharacterEncoding = "UTF-16LE"
    CaptionCharacterEncoding_UTF_32 CaptionCharacterEncoding = "UTF-32"
    CaptionCharacterEncoding_UTF_32BE CaptionCharacterEncoding = "UTF-32BE"
    CaptionCharacterEncoding_UTF_32LE CaptionCharacterEncoding = "UTF-32LE"
    CaptionCharacterEncoding_UTF_7 CaptionCharacterEncoding = "UTF-7"
    CaptionCharacterEncoding_UTF_8 CaptionCharacterEncoding = "UTF-8"
    CaptionCharacterEncoding_VISCII CaptionCharacterEncoding = "VISCII"
    CaptionCharacterEncoding_VISCII1_1_1 CaptionCharacterEncoding = "VISCII1.1-1"
    CaptionCharacterEncoding_WCHAR_T CaptionCharacterEncoding = "WCHAR_T"
    CaptionCharacterEncoding_WINBALTRIM CaptionCharacterEncoding = "WINBALTRIM"
    CaptionCharacterEncoding_WINDOWS_1250 CaptionCharacterEncoding = "WINDOWS-1250"
    CaptionCharacterEncoding_WINDOWS_1251 CaptionCharacterEncoding = "WINDOWS-1251"
    CaptionCharacterEncoding_WINDOWS_1252 CaptionCharacterEncoding = "WINDOWS-1252"
    CaptionCharacterEncoding_WINDOWS_1253 CaptionCharacterEncoding = "WINDOWS-1253"
    CaptionCharacterEncoding_WINDOWS_1254 CaptionCharacterEncoding = "WINDOWS-1254"
    CaptionCharacterEncoding_WINDOWS_1255 CaptionCharacterEncoding = "WINDOWS-1255"
    CaptionCharacterEncoding_WINDOWS_1256 CaptionCharacterEncoding = "WINDOWS-1256"
    CaptionCharacterEncoding_WINDOWS_1257 CaptionCharacterEncoding = "WINDOWS-1257"
    CaptionCharacterEncoding_WINDOWS_1258 CaptionCharacterEncoding = "WINDOWS-1258"
    CaptionCharacterEncoding_WINDOWS_874 CaptionCharacterEncoding = "WINDOWS-874"
    CaptionCharacterEncoding_WINDOWS_936 CaptionCharacterEncoding = "WINDOWS-936"
    CaptionCharacterEncoding_X0201 CaptionCharacterEncoding = "X0201"
    CaptionCharacterEncoding_X0208 CaptionCharacterEncoding = "X0208"
    CaptionCharacterEncoding_X0212 CaptionCharacterEncoding = "X0212"
)


