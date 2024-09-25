package dynamic_string_builder

type Stringable interface {
	String() string
}

type DynamicStringBuilder interface {
	Stringable
	Build(format string)
}

type NameBuilder struct {
	keyedValues map[string]string
}
