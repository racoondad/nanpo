package configs

type WebRtc struct {
	SipHost  string `mapstructure:"sip-host" yaml:"sip-host"`
	SipPort  int    `mapstructure:"sip-port" yaml:"sip-port"`
	RTMPAddr string `mapstructure:"rtmp-host" yaml:"rtmp-host"`
	RTMPPort int    `mapstructure:"rtmp-port" yaml:"rtmp-port"`
	RTSPAddr string `mapstructure:"rtsp-host" yaml:"rtsp-host"`
	RTSPPort int    `mapstructure:"rtsp-port" yaml:"rtsp-port"`
}
