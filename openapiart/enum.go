type ArgumentDirection string

var (
	ArgumentDirectionIn    ArgumentDirection = "in"
	ArgumentDirectionOut   ArgumentDirection = "out"
	ArgumentDirectionInOut ArgumentDirection = "inout"
)

type Type string

var (
	TypeBool   Type = "bool"
	TypeInt8   Type = "int8"
	TypeInt16  Type = "int16"
	TypeInt32  Type = "int32"
	TypeInt64  Type = "int64"
	TypeDouble Type = "double"
	TypeString Type = "string"
	TypeOctets Type = "octets"
)

func (t *Type) String() string {
	return string(*t)
}
