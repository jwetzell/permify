package utils

import (
	"errors"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"

	base "github.com/Permify/permify/pkg/pb/base/v1"
)

// Key -
func Key(v1, v2 string) string {
	var sb strings.Builder
	sb.WriteString(v1)
	sb.WriteString("#")
	sb.WriteString(v2)
	return sb.String()
}

func ArgumentsAsCelEnv(arguments map[string]base.AttributeType) (*cel.Env, error) {
	opts := make([]cel.EnvOption, 0, len(arguments))
	for name, typ := range arguments {
		typ, err := GetCelType(typ)
		if err != nil {
			return nil, err
		}

		opts = append(opts, cel.Variable(name, typ))
	}
	return cel.NewEnv(opts...)
}

func GetCelType(attributeType base.AttributeType) (*types.Type, error) {
	switch attributeType {
	case base.AttributeType_ATTRIBUTE_TYPE_STRING:
		return types.StringType, nil
	case base.AttributeType_ATTRIBUTE_TYPE_STRING_ARRAY:
		return cel.ListType(cel.StringType), nil
	case base.AttributeType_ATTRIBUTE_TYPE_BOOLEAN:
		return types.BoolType, nil
	case base.AttributeType_ATTRIBUTE_TYPE_BOOLEAN_ARRAY:
		return cel.ListType(types.BoolType), nil
	case base.AttributeType_ATTRIBUTE_TYPE_INTEGER:
		return types.IntType, nil
	case base.AttributeType_ATTRIBUTE_TYPE_INTEGER_ARRAY:
		return cel.ListType(types.IntType), nil
	case base.AttributeType_ATTRIBUTE_TYPE_DOUBLE:
		return types.DoubleType, nil
	case base.AttributeType_ATTRIBUTE_TYPE_DOUBLE_ARRAY:
		return cel.ListType(types.DoubleType), nil
	default:
		return nil, errors.New("")
	}
}

func ConvertProtoAnyToInterface(a *anypb.Any) interface{} {
	switch a.GetTypeUrl() {
	case "type.googleapis.com/base.v1.StringValue":
		stringValue := &base.StringValue{}
		if err := anypb.UnmarshalTo(a, stringValue, proto.UnmarshalOptions{}); err != nil {
			return ""
		}
		return stringValue.GetData()
	case "type.googleapis.com/base.v1.BooleanValue":
		boolValue := &base.BooleanValue{}
		if err := anypb.UnmarshalTo(a, boolValue, proto.UnmarshalOptions{}); err != nil {
			return false
		}
		return boolValue.GetData()
	case "type.googleapis.com/base.v1.IntegerValue":
		integerValue := &base.IntegerValue{}
		if err := anypb.UnmarshalTo(a, integerValue, proto.UnmarshalOptions{}); err != nil {
			return 0
		}
		return integerValue.GetData()
	case "type.googleapis.com/base.v1.DoubleValue":
		doubleValue := &base.DoubleValue{}
		if err := anypb.UnmarshalTo(a, doubleValue, proto.UnmarshalOptions{}); err != nil {
			return 0.0
		}
		return doubleValue.GetData()
	case "type.googleapis.com/base.v1.StringArrayValue":
		stringArrayValue := &base.StringArrayValue{}
		if err := anypb.UnmarshalTo(a, stringArrayValue, proto.UnmarshalOptions{}); err != nil {
			return []string{}
		}
		return stringArrayValue.GetData()
	case "type.googleapis.com/base.v1.BooleanArrayValue":
		booleanArrayValue := &base.BooleanArrayValue{}
		if err := anypb.UnmarshalTo(a, booleanArrayValue, proto.UnmarshalOptions{}); err != nil {
			return []bool{}
		}
		return booleanArrayValue.GetData()
	case "type.googleapis.com/base.v1.IntegerArrayValue":
		integerArrayValue := &base.IntegerArrayValue{}
		if err := anypb.UnmarshalTo(a, integerArrayValue, proto.UnmarshalOptions{}); err != nil {
			return []int32{}
		}
		return integerArrayValue.GetData()
	case "type.googleapis.com/base.v1.DoubleArrayValue":
		doubleArrayValue := &base.DoubleArrayValue{}
		if err := anypb.UnmarshalTo(a, doubleArrayValue, proto.UnmarshalOptions{}); err != nil {
			return []float64{}
		}
		return doubleArrayValue.GetData()
	default:
		return ""
	}
}
