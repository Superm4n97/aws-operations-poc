package convert

func StringP(st string) *string {
	return &st
}

func StringPSlice(sl []string) []*string {
	var ret []*string
	for i := 0; i < len(sl); i++ {
		ret = append(ret, &sl[i])
	}
	return ret
}
