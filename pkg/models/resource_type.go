package models

type ResourceType uint16

// https://gibberlings3.github.io/iesdp/file_formats/general.htm
const (
	TYPE_UNKNOWN = ResourceType(0x0000)
	TYPE_BMP     = ResourceType(0x0001)
	TYPE_MVE     = ResourceType(0x0002)
	TYPE_WAV     = ResourceType(0x0004)
	TYPE_PLT     = ResourceType(0x0006)
	TYPE_BAM     = ResourceType(0x03e8)
	TYPE_WED     = ResourceType(0x03e9)
	TYPE_CHU     = ResourceType(0x03ea)
	TYPE_TIS     = ResourceType(0x03eb)
	TYPE_MOS     = ResourceType(0x03ec)
	TYPE_ITM     = ResourceType(0x03ed)
	TYPE_SPL     = ResourceType(0x03ee)
	TYPE_BCS     = ResourceType(0x03ef)
	TYPE_IDS     = ResourceType(0x03f0)
	TYPE_CRE     = ResourceType(0x03f1)
	TYPE_ARE     = ResourceType(0x03f2)
	TYPE_DLG     = ResourceType(0x03f3)
	TYPE_2DA     = ResourceType(0x03f4)
	TYPE_GAM     = ResourceType(0x03f5)
	TYPE_STO     = ResourceType(0x03f6)
	TYPE_WMP     = ResourceType(0x03f7)
	TYPE_CHR2    = ResourceType(0x03f8)
	TYPE_EFF     = ResourceType(0x03f8)
	TYPE_BC      = ResourceType(0x03f9)
	TYPE_CHR     = ResourceType(0x03fa)
	TYPE_VVC     = ResourceType(0x03fb)
	TYPE_VEF     = ResourceType(0x03fc)
	TYPE_PRO     = ResourceType(0x03fd)
	TYPE_BIO     = ResourceType(0x03fe)
	TYPE_WBM     = ResourceType(0x03ff)
	TYPE_FNT     = ResourceType(0x0400)
	TYPE_GUI     = ResourceType(0x0402)
	TYPE_SQL     = ResourceType(0x0403)
	TYPE_PVRZ    = ResourceType(0x0404)
	TYPE_GLSL    = ResourceType(0x0405)
	TYPE_MENU    = ResourceType(0x0408)
	TYPE_MENU2   = ResourceType(0x0409)
	TYPE_TTF     = ResourceType(0x040a)
	TYPE_PNG     = ResourceType(0x040b)
	TYPE_BAH     = ResourceType(0x044c)
	TYPE_INI     = ResourceType(0x0802)
	TYPE_SRC     = ResourceType(0x0803)
)