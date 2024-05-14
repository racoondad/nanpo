package po

import "github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"

const (
	// 宇视人脸
	Uniview_face = iota + 1
	// 宇视指纹
	Uniview_fp
	// 熵基人脸
	Zkteco_face
	// 熵基指纹
	Zkteco_fp
	// 德服人脸
	Defu_face
	// 德服指纹
	Defu_fp
	// 英泽人脸
	FCARD_face
	// 英泽指纹
	FCARD_fp
)

type FingPrint struct {
	changes map[string]interface{}
	ormtypes.Model
	BioType  int
	BioIndex int
	BioAddr  int
}
